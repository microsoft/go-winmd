// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package winmd

import (
	"debug/pe"
	"errors"
	"fmt"
	"iter"
)

// A Metadata represents an open Windows Metadata file.
type Metadata struct {
	Version string
	Tables  *Tables
	Strings StringHeap
	US      USHeap
	Blob    BlobHeap
	GUID    GUIDHeap

	layout *layout
}

// Open opens a Windows Metadata file at path and returns
// a Metadata struct that provides access to its contents.
func Open(path string) (*Metadata, error) {
	pefile, err := pe.Open(path)
	if err != nil {
		return nil, err
	}
	defer pefile.Close()
	return New(pefile)
}

// New creates a new Metadata from an underlying PE file.
// The PE file can be closed after calling New, as the
// returned Metadata doesn't keep any reference to it.
// A #~ tables stream is required; uncompressed #- streams are not supported.
// On success, Tables is non-nil, even when all tables are empty.
func New(pefile *pe.File) (*Metadata, error) {
	return newMetadata(pefile)
}

// FieldSignature decodes an entire field signature blob, rejecting trailing data.
// Type nesting is limited to 64 levels.
// Type handle bounds are checked when table metadata is available.
// Generic instantiations and VAR/MVAR parameter numbers are preserved without
// substitution or validation against a declaring type's or method's constraints.
func (m *Metadata) FieldSignature(bytes SigFieldBlob) (SigField, error) {
	r := m.sigReader(bytes)
	sig := r.fieldSig()
	if r.err == nil && len(r.data) != 0 {
		r.err = errors.New("trailing field signature data")
	}
	return sig, r.err
}

// MethodDefSignature decodes an entire method definition signature blob, rejecting trailing data.
// Type nesting is limited to 64 levels per return type or parameter.
// Type handle bounds are checked when table metadata is available.
// Generic instantiations and VAR/MVAR parameter numbers are preserved without
// substitution or validation against a declaring type's or method's constraints.
func (m *Metadata) MethodDefSignature(data SigMethodDefBlob) (SigMethodDef, error) {
	r := m.sigReader(data)
	sig := r.methodDefSig()
	if r.err == nil && len(r.data) != 0 {
		r.err = errors.New("trailing method definition signature data")
	}
	return sig, r.err
}

// MethodRefSignature decodes an entire method reference signature from
// MemberRef.Signature, rejecting trailing data. Fixed and optional VARARG
// parameters are kept separately in Param and VariableParam.
// Type nesting, including function-pointer signatures, is limited to 64 levels.
// Type handle bounds are checked when table metadata is available. Generic
// parameter numbers are preserved without substitution or constraint checking.
func (m *Metadata) MethodRefSignature(data SigMethodRefBlob) (SigMethodRef, error) {
	r := m.sigReader(data)
	sig := r.methodSig(methodSigAllowGeneric|methodSigAllowSentinel, 0).SigMethodRef
	if r.err == nil && len(r.data) != 0 {
		r.err = errors.New("trailing method reference signature data")
	}
	return sig, r.err
}

// StandAloneMethodSignature decodes an entire calli signature from
// StandAloneSig.Signature, rejecting trailing data. Managed and unmanaged
// calling conventions are preserved; VARARG and Cdecl may have optional
// parameters after SENTINEL. Standalone signatures cannot declare generic
// parameters, but their types may refer to enclosing generic parameters.
// Type nesting, including function-pointer signatures, is limited to 64 levels.
// Type handle bounds are checked when table metadata is available.
func (m *Metadata) StandAloneMethodSignature(data SigStandAloneMethodBlob) (SigStandAloneMethod, error) {
	r := m.sigReader(data)
	sig := r.methodSig(methodSigAllowUnmanaged|methodSigAllowSentinel, 0)
	if r.err == nil && len(r.data) != 0 {
		r.err = errors.New("trailing standalone method signature data")
	}
	return sig, r.err
}

// PropertySignature decodes an entire property signature blob, rejecting trailing data.
// The blob is stored in [Property.Type].
// Type nesting is limited to 64 levels per property type or index parameter.
// Type handle bounds are checked when table metadata is available.
// Generic instantiations and VAR/MVAR parameter numbers are preserved without
// substitution or validation against a declaring type's or method's constraints.
func (m *Metadata) PropertySignature(data SigPropertyBlob) (SigProperty, error) {
	r := m.sigReader(data)
	sig := r.propertySig()
	if r.err == nil && len(r.data) != 0 {
		r.err = errors.New("trailing property signature data")
	}
	return sig, r.err
}

// LocalVarsSignature decodes an entire local variable signature from
// StandAloneSig.Signature, rejecting trailing data. It preserves the encoded
// order of custom modifiers and PINNED constraints. There must be 1 to 65534
// locals. Type nesting is limited to 64 levels per local, including nested
// function-pointer signatures. Type handle bounds are checked when table
// metadata is available; pinning eligibility and generic constraints are not.
func (m *Metadata) LocalVarsSignature(data SigLocalVarsBlob) (SigLocalVars, error) {
	r := m.sigReader(data)
	sig := r.localVarsSig()
	if r.err == nil && len(r.data) != 0 {
		r.err = errors.New("trailing local variable signature data")
	}
	return sig, r.err
}

// TypeSpecSignature decodes an entire type specification from TypeSpec.Signature,
// rejecting trailing data. There is no calling-convention byte. The type can
// contain open generic parameter references; leading custom modifiers, BYREF,
// TYPEDBYREF, and VOID are not permitted in this context.
// Type nesting, including function-pointer signatures, is limited to 64 levels.
// Type handle bounds are checked when table metadata is available.
func (m *Metadata) TypeSpecSignature(data SigTypeSpecBlob) (SigTypeSpec, error) {
	r := m.sigReader(data)
	typ := r.decodeType(0, 0)
	if r.err == nil && len(r.data) != 0 {
		r.err = errors.New("trailing type specification signature data")
	}
	return SigTypeSpec{Kind: typ.Kind, Value: typ.Value}, r.err
}

// MethodSpecSignature decodes an entire generic method instantiation from
// MethodSpec.Instantiation, rejecting trailing data. At least one type argument
// is required. Type nesting is limited to 64 levels per argument, and type
// handle bounds are checked when table metadata is available. Generic parameter
// numbers are preserved without substitution or checking the target's arity
// or constraints.
func (m *Metadata) MethodSpecSignature(data SigMethodSpecBlob) (SigMethodSpec, error) {
	r := m.sigReader(data)
	sig := r.methodSpecSig()
	if r.err == nil && len(r.data) != 0 {
		r.err = errors.New("trailing method specification signature data")
	}
	return sig, r.err
}

func (m *Metadata) sigReader(data []byte) sigReader {
	return sigReader{
		ecma335Reader{
			data:   data,
			layout: m.layout,
		},
	}
}

// Index indexes a record in a table.
type Index uint32

// CodedIndex indexes a record on any table.
type CodedIndex[T CodedTag] struct {
	Index Index
	Tag   T
}

// String is complete UTF8 string from the #String heap
// It does not contain the null-terminated character.
//
// It is used as an optimization to avoid allocating
// when reading from the #Strings heap.
type String struct {
	// Start is the offset in the #Strings heap where the string starts. This is the parameter that
	// was passed to StringHeap.String to create this String. Equal strings can be stored at
	// different offsets, so Start identifies a heap location, not a unique string value.
	Start uint32
	data  []byte
}

func (s String) String() string {
	return string(s.data)
}

// Slice indexes the range of records [Start,End) on the table T.
type Slice struct {
	Start Index
	End   Index
}

// Len returns the number of records in the slice.
func (s Slice) Len() uint32 {
	if s.End < s.Start {
		return 0
	}
	return uint32(s.End - s.Start)
}

// All returns a sequence of all indices in the slice.
func (s Slice) All() iter.Seq[Index] {
	return func(yield func(Index) bool) {
		for i := s.Start; i < s.End; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Table is a record container as defined in §II.22.
type Table[T any] struct {
	name string
	len  uint32

	decode func(recordReader) (T, string, error)
	width  uint8
	data   []byte
	heaps  *heaps
	layout *layout
}

func newTable[T any](name string, data []byte, hps *heaps, layout *layout, table table, decode func(recordReader) (T, string, error)) Table[T] {
	info := layout.tables[table]
	return Table[T]{
		name:   name,
		len:    info.rowCount,
		decode: decode,
		width:  uint8(info.width),
		data:   data[info.offset : info.offset+int(info.width)*int(info.rowCount)],
		heaps:  hps,
		layout: layout,
	}
}

func (t Table[T]) Indices() iter.Seq[Index] {
	return func(yield func(Index) bool) {
		for i := uint32(0); i < t.len; i++ {
			if !yield(Index(i)) {
				return
			}
		}
	}
}

// Len returns the number of records in the table.
func (t Table[T]) Len() uint32 {
	return t.len
}

// Name returns the metadata table name.
func (t Table[T]) Name() string {
	return t.name
}

// DecodeError describes a failure to read a metadata table row.
// The underlying error is available through errors.Is and errors.As.
type DecodeError struct {
	// Table is the metadata table name. It is empty for a zero-value Table.
	Table string
	// Row is the zero-based index passed to Table.At. For list lookahead errors,
	// this is the row being decoded, not the subsequent row being inspected.
	Row Index
	// Column is the Go field name of the first column that failed to decode.
	// It is empty when no column was read, such as for an out-of-range row.
	Column string
	// Err is the underlying decoding or bounds error.
	Err error
}

func (e *DecodeError) Error() string {
	table := e.Table
	if table == "" {
		table = "table"
	}
	if e.Column != "" {
		return fmt.Sprintf("%s[%d].%s: %v", table, e.Row, e.Column, e.Err)
	}
	return fmt.Sprintf("%s[%d]: %v", table, e.Row, e.Err)
}

// Unwrap returns the underlying error.
func (e *DecodeError) Unwrap() error {
	return e.Err
}

// At returns the record at row.
// Null coded indices are accepted only in columns that permit them.
// On failure, the error is a *DecodeError identifying the table, zero-based row,
// and column, when applicable.
func (t Table[T]) At(row Index) (T, error) {
	var zero T
	if uint32(row) >= t.len {
		return zero, &DecodeError{Table: t.name, Row: row, Err: fmt.Errorf("row %d is beyond the end of the table", row)}
	}
	offset := int(t.width) * int(row)
	r := recordReader{
		ecma335Reader: ecma335Reader{
			data:   t.data[offset:],
			layout: t.layout,
		},
		heaps: t.heaps,
	}
	rec, column, err := t.decode(r)
	if err != nil {
		return rec, &DecodeError{Table: t.name, Row: row, Column: column, Err: err}
	}
	return rec, nil
}
