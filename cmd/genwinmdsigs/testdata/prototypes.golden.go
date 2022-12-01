// APIs for Windows.Win32.Storage.FileSystem
//sys	SearchPathW(/*in*//*optional*/lpPath PWSTR, /*in*/lpFileName PWSTR, /*in*//*optional*/lpExtension PWSTR, /*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PWSTR, /*out*//*optional*/lpFilePart *PWSTR) (uint32)
//sys	SearchPathA(/*in*//*optional*/lpPath PSTR, /*in*/lpFileName PSTR, /*in*//*optional*/lpExtension PSTR, /*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PSTR, /*out*//*optional*/lpFilePart *PSTR) (uint32)
//sys	CompareFileTime(/*in*/lpFileTime1 *FILETIME, /*in*/lpFileTime2 *FILETIME) (int32)
//sys	CreateDirectoryA(/*in*/lpPathName PSTR, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES) (BOOL)
//sys	CreateDirectoryW(/*in*/lpPathName PWSTR, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES) (BOOL)
//sys	CreateFileA(/*in*/lpFileName PSTR, /*in*/dwDesiredAccess FILE_ACCESS_FLAGS, /*in*/dwShareMode FILE_SHARE_MODE, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/dwCreationDisposition FILE_CREATION_DISPOSITION, /*in*/dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES, /*in*//*optional*/hTemplateFile HANDLE) (HANDLE)
//sys	CreateFileW(/*in*/lpFileName PWSTR, /*in*/dwDesiredAccess FILE_ACCESS_FLAGS, /*in*/dwShareMode FILE_SHARE_MODE, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/dwCreationDisposition FILE_CREATION_DISPOSITION, /*in*/dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES, /*in*//*optional*/hTemplateFile HANDLE) (HANDLE)
//sys	DefineDosDeviceW(/*in*/dwFlags DEFINE_DOS_DEVICE_FLAGS, /*in*/lpDeviceName PWSTR, /*in*//*optional*/lpTargetPath PWSTR) (BOOL)
//sys	DeleteFileA(/*in*/lpFileName PSTR) (BOOL)
//sys	DeleteFileW(/*in*/lpFileName PWSTR) (BOOL)
//sys	DeleteVolumeMountPointW(/*in*/lpszVolumeMountPoint PWSTR) (BOOL)
//sys	FileTimeToLocalFileTime(/*in*/lpFileTime *FILETIME, /*out*/lpLocalFileTime *FILETIME) (BOOL)
//sys	FindClose(/*in*/hFindFile FindFileHandle) (BOOL)
//sys	FindCloseChangeNotification(/*in*/hChangeHandle FindChangeNotificationHandle) (BOOL)
//sys	FindFirstChangeNotificationA(/*in*/lpPathName PSTR, /*in*/bWatchSubtree BOOL, /*in*/dwNotifyFilter FILE_NOTIFY_CHANGE) (FindChangeNotificationHandle)
//sys	FindFirstChangeNotificationW(/*in*/lpPathName PWSTR, /*in*/bWatchSubtree BOOL, /*in*/dwNotifyFilter FILE_NOTIFY_CHANGE) (FindChangeNotificationHandle)
//sys	FindFirstFileA(/*in*/lpFileName PSTR, /*out*/lpFindFileData *WIN32_FIND_DATAA) (FindFileHandle)
//sys	FindFirstFileW(/*in*/lpFileName PWSTR, /*out*/lpFindFileData *WIN32_FIND_DATAW) (FindFileHandle)
//sys	FindFirstFileExA(/*in*/lpFileName PSTR, /*in*/fInfoLevelId FINDEX_INFO_LEVELS, /*out*/lpFindFileData unsafe.Pointer, /*in*/fSearchOp FINDEX_SEARCH_OPS, /*in*//*out*/lpSearchFilter unsafe.Pointer, /*in*/dwAdditionalFlags FIND_FIRST_EX_FLAGS) (FindFileHandle)
//sys	FindFirstFileExW(/*in*/lpFileName PWSTR, /*in*/fInfoLevelId FINDEX_INFO_LEVELS, /*out*/lpFindFileData unsafe.Pointer, /*in*/fSearchOp FINDEX_SEARCH_OPS, /*in*//*out*/lpSearchFilter unsafe.Pointer, /*in*/dwAdditionalFlags FIND_FIRST_EX_FLAGS) (FindFileHandle)
//sys	FindFirstVolumeW(/*out*/lpszVolumeName PWSTR, /*in*/cchBufferLength uint32) (FindVolumeHandle)
//sys	FindNextChangeNotification(/*in*/hChangeHandle FindChangeNotificationHandle) (BOOL)
//sys	FindNextFileA(/*in*/hFindFile FindFileHandle, /*out*/lpFindFileData *WIN32_FIND_DATAA) (BOOL)
//sys	FindNextFileW(/*in*/hFindFile HANDLE, /*out*/lpFindFileData *WIN32_FIND_DATAW) (BOOL)
//sys	FindNextVolumeW(/*in*/hFindVolume FindVolumeHandle, /*out*/lpszVolumeName PWSTR, /*in*/cchBufferLength uint32) (BOOL)
//sys	FindVolumeClose(/*in*/hFindVolume FindVolumeHandle) (BOOL)
//sys	FlushFileBuffers(/*in*/hFile HANDLE) (BOOL)
//sys	GetDiskFreeSpaceA(/*in*//*optional*/lpRootPathName PSTR, /*out*//*optional*/lpSectorsPerCluster *uint32, /*out*//*optional*/lpBytesPerSector *uint32, /*out*//*optional*/lpNumberOfFreeClusters *uint32, /*out*//*optional*/lpTotalNumberOfClusters *uint32) (BOOL)
//sys	GetDiskFreeSpaceW(/*in*//*optional*/lpRootPathName PWSTR, /*out*//*optional*/lpSectorsPerCluster *uint32, /*out*//*optional*/lpBytesPerSector *uint32, /*out*//*optional*/lpNumberOfFreeClusters *uint32, /*out*//*optional*/lpTotalNumberOfClusters *uint32) (BOOL)
//sys	GetDiskFreeSpaceExA(/*in*//*optional*/lpDirectoryName PSTR, /*out*//*optional*/lpFreeBytesAvailableToCaller *ULARGE_INTEGER, /*out*//*optional*/lpTotalNumberOfBytes *ULARGE_INTEGER, /*out*//*optional*/lpTotalNumberOfFreeBytes *ULARGE_INTEGER) (BOOL)
//sys	GetDiskFreeSpaceExW(/*in*//*optional*/lpDirectoryName PWSTR, /*out*//*optional*/lpFreeBytesAvailableToCaller *ULARGE_INTEGER, /*out*//*optional*/lpTotalNumberOfBytes *ULARGE_INTEGER, /*out*//*optional*/lpTotalNumberOfFreeBytes *ULARGE_INTEGER) (BOOL)
//sys	GetDiskSpaceInformationA(/*in*//*optional*/rootPath PSTR, /*out*/diskSpaceInfo *DISK_SPACE_INFORMATION) (HRESULT)
//sys	GetDiskSpaceInformationW(/*in*//*optional*/rootPath PWSTR, /*out*/diskSpaceInfo *DISK_SPACE_INFORMATION) (HRESULT)
//sys	GetDriveTypeA(/*in*//*optional*/lpRootPathName PSTR) (uint32)
//sys	GetDriveTypeW(/*in*//*optional*/lpRootPathName PWSTR) (uint32)
//sys	GetFileAttributesA(/*in*/lpFileName PSTR) (uint32)
//sys	GetFileAttributesW(/*in*/lpFileName PWSTR) (uint32)
//sys	GetFileAttributesExA(/*in*/lpFileName PSTR, /*in*/fInfoLevelId GET_FILEEX_INFO_LEVELS, /*out*/lpFileInformation unsafe.Pointer) (BOOL)
//sys	GetFileAttributesExW(/*in*/lpFileName PWSTR, /*in*/fInfoLevelId GET_FILEEX_INFO_LEVELS, /*out*/lpFileInformation unsafe.Pointer) (BOOL)
//sys	GetFileInformationByHandle(/*in*/hFile HANDLE, /*out*/lpFileInformation *BY_HANDLE_FILE_INFORMATION) (BOOL)
//sys	GetFileSize(/*in*/hFile HANDLE, /*out*//*optional*/lpFileSizeHigh *uint32) (uint32)
//sys	GetFileSizeEx(/*in*/hFile HANDLE, /*out*/lpFileSize *LARGE_INTEGER) (BOOL)
//sys	GetFileType(/*in*/hFile HANDLE) (uint32)
//sys	GetFinalPathNameByHandleA(/*in*/hFile HANDLE, /*out*/lpszFilePath PSTR, /*in*/cchFilePath uint32, /*in*/dwFlags FILE_NAME) (uint32)
//sys	GetFinalPathNameByHandleW(/*in*/hFile HANDLE, /*out*/lpszFilePath PWSTR, /*in*/cchFilePath uint32, /*in*/dwFlags FILE_NAME) (uint32)
//sys	GetFileTime(/*in*/hFile HANDLE, /*out*//*optional*/lpCreationTime *FILETIME, /*out*//*optional*/lpLastAccessTime *FILETIME, /*out*//*optional*/lpLastWriteTime *FILETIME) (BOOL)
//sys	GetFullPathNameW(/*in*/lpFileName PWSTR, /*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PWSTR, /*out*//*optional*/lpFilePart *PWSTR) (uint32)
//sys	GetFullPathNameA(/*in*/lpFileName PSTR, /*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PSTR, /*out*//*optional*/lpFilePart *PSTR) (uint32)
//sys	GetLogicalDrives() (uint32)
//sys	GetLogicalDriveStringsW(/*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PWSTR) (uint32)
//sys	GetLongPathNameA(/*in*/lpszShortPath PSTR, /*out*//*optional*/lpszLongPath PSTR, /*in*/cchBuffer uint32) (uint32)
//sys	GetLongPathNameW(/*in*/lpszShortPath PWSTR, /*out*//*optional*/lpszLongPath PWSTR, /*in*/cchBuffer uint32) (uint32)
//sys	AreShortNamesEnabled(/*in*/Handle HANDLE, /*out*/Enabled *BOOL) (BOOL)
//sys	GetShortPathNameW(/*in*/lpszLongPath PWSTR, /*out*//*optional*/lpszShortPath PWSTR, /*in*/cchBuffer uint32) (uint32)
//sys	GetTempFileNameW(/*in*/lpPathName PWSTR, /*in*/lpPrefixString PWSTR, /*in*/uUnique uint32, /*out*/lpTempFileName PWSTR) (uint32)
//sys	GetVolumeInformationByHandleW(/*in*/hFile HANDLE, /*out*//*optional*/lpVolumeNameBuffer PWSTR, /*in*/nVolumeNameSize uint32, /*out*//*optional*/lpVolumeSerialNumber *uint32, /*out*//*optional*/lpMaximumComponentLength *uint32, /*out*//*optional*/lpFileSystemFlags *uint32, /*out*//*optional*/lpFileSystemNameBuffer PWSTR, /*in*/nFileSystemNameSize uint32) (BOOL)
//sys	GetVolumeInformationW(/*in*//*optional*/lpRootPathName PWSTR, /*out*//*optional*/lpVolumeNameBuffer PWSTR, /*in*/nVolumeNameSize uint32, /*out*//*optional*/lpVolumeSerialNumber *uint32, /*out*//*optional*/lpMaximumComponentLength *uint32, /*out*//*optional*/lpFileSystemFlags *uint32, /*out*//*optional*/lpFileSystemNameBuffer PWSTR, /*in*/nFileSystemNameSize uint32) (BOOL)
//sys	GetVolumePathNameW(/*in*/lpszFileName PWSTR, /*out*/lpszVolumePathName PWSTR, /*in*/cchBufferLength uint32) (BOOL)
//sys	LocalFileTimeToFileTime(/*in*/lpLocalFileTime *FILETIME, /*out*/lpFileTime *FILETIME) (BOOL)
//sys	LockFile(/*in*/hFile HANDLE, /*in*/dwFileOffsetLow uint32, /*in*/dwFileOffsetHigh uint32, /*in*/nNumberOfBytesToLockLow uint32, /*in*/nNumberOfBytesToLockHigh uint32) (BOOL)
//sys	LockFileEx(/*in*/hFile HANDLE, /*in*/dwFlags LOCK_FILE_FLAGS, /*in*/dwReserved uint32, /*in*/nNumberOfBytesToLockLow uint32, /*in*/nNumberOfBytesToLockHigh uint32, /*in*//*out*/lpOverlapped *OVERLAPPED) (BOOL)
//sys	QueryDosDeviceW(/*in*//*optional*/lpDeviceName PWSTR, /*out*//*optional*/lpTargetPath PWSTR, /*in*/ucchMax uint32) (uint32)
//sys	ReadFile(/*in*/hFile HANDLE, /*out*//*optional*/lpBuffer unsafe.Pointer, /*in*/nNumberOfBytesToRead uint32, /*out*//*optional*/lpNumberOfBytesRead *uint32, /*in*//*out*//*optional*/lpOverlapped *OVERLAPPED) (BOOL)
//sys	ReadFileEx(/*in*/hFile HANDLE, /*out*//*optional*/lpBuffer unsafe.Pointer, /*in*/nNumberOfBytesToRead uint32, /*in*//*out*/lpOverlapped *OVERLAPPED, /*in*/lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE) (BOOL)
//sys	ReadFileScatter(/*in*/hFile HANDLE, /*in*/aSegmentArray *FILE_SEGMENT_ELEMENT, /*in*/nNumberOfBytesToRead uint32, /*in*//*out*/lpReserved *uint32, /*in*//*out*/lpOverlapped *OVERLAPPED) (BOOL)
//sys	RemoveDirectoryA(/*in*/lpPathName PSTR) (BOOL)
//sys	RemoveDirectoryW(/*in*/lpPathName PWSTR) (BOOL)
//sys	SetEndOfFile(/*in*/hFile HANDLE) (BOOL)
//sys	SetFileAttributesA(/*in*/lpFileName PSTR, /*in*/dwFileAttributes FILE_FLAGS_AND_ATTRIBUTES) (BOOL)
//sys	SetFileAttributesW(/*in*/lpFileName PWSTR, /*in*/dwFileAttributes FILE_FLAGS_AND_ATTRIBUTES) (BOOL)
//sys	SetFileInformationByHandle(/*in*/hFile HANDLE, /*in*/FileInformationClass FILE_INFO_BY_HANDLE_CLASS, /*in*/lpFileInformation unsafe.Pointer, /*in*/dwBufferSize uint32) (BOOL)
//sys	SetFilePointer(/*in*/hFile HANDLE, /*in*/lDistanceToMove int32, /*in*//*out*//*optional*/lpDistanceToMoveHigh *int32, /*in*/dwMoveMethod SET_FILE_POINTER_MOVE_METHOD) (uint32)
//sys	SetFilePointerEx(/*in*/hFile HANDLE, /*in*/liDistanceToMove LARGE_INTEGER, /*out*//*optional*/lpNewFilePointer *LARGE_INTEGER, /*in*/dwMoveMethod SET_FILE_POINTER_MOVE_METHOD) (BOOL)
//sys	SetFileTime(/*in*/hFile HANDLE, /*in*//*optional*/lpCreationTime *FILETIME, /*in*//*optional*/lpLastAccessTime *FILETIME, /*in*//*optional*/lpLastWriteTime *FILETIME) (BOOL)
//sys	SetFileValidData(/*in*/hFile HANDLE, /*in*/ValidDataLength int64) (BOOL)
//sys	UnlockFile(/*in*/hFile HANDLE, /*in*/dwFileOffsetLow uint32, /*in*/dwFileOffsetHigh uint32, /*in*/nNumberOfBytesToUnlockLow uint32, /*in*/nNumberOfBytesToUnlockHigh uint32) (BOOL)
//sys	UnlockFileEx(/*in*/hFile HANDLE, /*in*/dwReserved uint32, /*in*/nNumberOfBytesToUnlockLow uint32, /*in*/nNumberOfBytesToUnlockHigh uint32, /*in*//*out*/lpOverlapped *OVERLAPPED) (BOOL)
//sys	WriteFile(/*in*/hFile HANDLE, /*in*//*optional*/lpBuffer unsafe.Pointer, /*in*/nNumberOfBytesToWrite uint32, /*out*//*optional*/lpNumberOfBytesWritten *uint32, /*in*//*out*//*optional*/lpOverlapped *OVERLAPPED) (BOOL)
//sys	WriteFileEx(/*in*/hFile HANDLE, /*in*//*optional*/lpBuffer unsafe.Pointer, /*in*/nNumberOfBytesToWrite uint32, /*in*//*out*/lpOverlapped *OVERLAPPED, /*in*/lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE) (BOOL)
//sys	WriteFileGather(/*in*/hFile HANDLE, /*in*/aSegmentArray *FILE_SEGMENT_ELEMENT, /*in*/nNumberOfBytesToWrite uint32, /*in*//*out*/lpReserved *uint32, /*in*//*out*/lpOverlapped *OVERLAPPED) (BOOL)
//sys	GetTempPathW(/*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PWSTR) (uint32)
//sys	GetVolumeNameForVolumeMountPointW(/*in*/lpszVolumeMountPoint PWSTR, /*out*/lpszVolumeName PWSTR, /*in*/cchBufferLength uint32) (BOOL)
//sys	GetVolumePathNamesForVolumeNameW(/*in*/lpszVolumeName PWSTR, /*out*//*optional*/lpszVolumePathNames PWSTR, /*in*/cchBufferLength uint32, /*out*/lpcchReturnLength *uint32) (BOOL)
//sys	CreateFile2(/*in*/lpFileName PWSTR, /*in*/dwDesiredAccess FILE_ACCESS_FLAGS, /*in*/dwShareMode FILE_SHARE_MODE, /*in*/dwCreationDisposition FILE_CREATION_DISPOSITION, /*in*//*optional*/pCreateExParams *CREATEFILE2_EXTENDED_PARAMETERS) (HANDLE)
//sys	SetFileIoOverlappedRange(/*in*/FileHandle HANDLE, /*in*/OverlappedRangeStart *uint8, /*in*/Length uint32) (BOOL)
//sys	GetCompressedFileSizeA(/*in*/lpFileName PSTR, /*out*//*optional*/lpFileSizeHigh *uint32) (uint32)
//sys	GetCompressedFileSizeW(/*in*/lpFileName PWSTR, /*out*//*optional*/lpFileSizeHigh *uint32) (uint32)
//sys	FindFirstStreamW(/*in*/lpFileName PWSTR, /*in*/InfoLevel STREAM_INFO_LEVELS, /*out*/lpFindStreamData unsafe.Pointer, /*in*/dwFlags uint32) (FindStreamHandle)
//sys	FindNextStreamW(/*in*/hFindStream FindStreamHandle, /*out*/lpFindStreamData unsafe.Pointer) (BOOL)
//sys	AreFileApisANSI() (BOOL)
//sys	GetTempPathA(/*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PSTR) (uint32)
//sys	FindFirstFileNameW(/*in*/lpFileName PWSTR, /*in*/dwFlags uint32, /*in*//*out*/StringLength *uint32, /*out*/LinkName PWSTR) (FindFileNameHandle)
//sys	FindNextFileNameW(/*in*/hFindStream FindFileNameHandle, /*in*//*out*/StringLength *uint32, /*out*/LinkName PWSTR) (BOOL)
//sys	GetVolumeInformationA(/*in*//*optional*/lpRootPathName PSTR, /*out*//*optional*/lpVolumeNameBuffer PSTR, /*in*/nVolumeNameSize uint32, /*out*//*optional*/lpVolumeSerialNumber *uint32, /*out*//*optional*/lpMaximumComponentLength *uint32, /*out*//*optional*/lpFileSystemFlags *uint32, /*out*//*optional*/lpFileSystemNameBuffer PSTR, /*in*/nFileSystemNameSize uint32) (BOOL)
//sys	GetTempFileNameA(/*in*/lpPathName PSTR, /*in*/lpPrefixString PSTR, /*in*/uUnique uint32, /*out*/lpTempFileName PSTR) (uint32)
//sys	SetFileApisToOEM()
//sys	SetFileApisToANSI()
//sys	GetTempPath2W(/*in*/BufferLength uint32, /*out*//*optional*/Buffer PWSTR) (uint32)
//sys	GetTempPath2A(/*in*/BufferLength uint32, /*out*//*optional*/Buffer PSTR) (uint32)
//sys	CopyFileFromAppW(/*in*/lpExistingFileName PWSTR, /*in*/lpNewFileName PWSTR, /*in*/bFailIfExists BOOL) (BOOL) = api-ms-win-core-file-fromapp-l1-1-0.CopyFileFromAppW
//sys	CreateDirectoryFromAppW(/*in*/lpPathName PWSTR, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES) (BOOL) = api-ms-win-core-file-fromapp-l1-1-0.CreateDirectoryFromAppW
//sys	CreateFileFromAppW(/*in*/lpFileName PWSTR, /*in*/dwDesiredAccess uint32, /*in*/dwShareMode uint32, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/dwCreationDisposition uint32, /*in*/dwFlagsAndAttributes uint32, /*in*//*optional*/hTemplateFile HANDLE) (HANDLE) = api-ms-win-core-file-fromapp-l1-1-0.CreateFileFromAppW
//sys	CreateFile2FromAppW(/*in*/lpFileName PWSTR, /*in*/dwDesiredAccess uint32, /*in*/dwShareMode uint32, /*in*/dwCreationDisposition uint32, /*in*//*optional*/pCreateExParams *CREATEFILE2_EXTENDED_PARAMETERS) (HANDLE) = api-ms-win-core-file-fromapp-l1-1-0.CreateFile2FromAppW
//sys	DeleteFileFromAppW(/*in*/lpFileName PWSTR) (BOOL) = api-ms-win-core-file-fromapp-l1-1-0.DeleteFileFromAppW
//sys	FindFirstFileExFromAppW(/*in*/lpFileName PWSTR, /*in*/fInfoLevelId FINDEX_INFO_LEVELS, /*out*/lpFindFileData unsafe.Pointer, /*in*/fSearchOp FINDEX_SEARCH_OPS, /*in*//*out*/lpSearchFilter unsafe.Pointer, /*in*/dwAdditionalFlags uint32) (HANDLE) = api-ms-win-core-file-fromapp-l1-1-0.FindFirstFileExFromAppW
//sys	GetFileAttributesExFromAppW(/*in*/lpFileName PWSTR, /*in*/fInfoLevelId GET_FILEEX_INFO_LEVELS, /*out*/lpFileInformation unsafe.Pointer) (BOOL) = api-ms-win-core-file-fromapp-l1-1-0.GetFileAttributesExFromAppW
//sys	MoveFileFromAppW(/*in*/lpExistingFileName PWSTR, /*in*/lpNewFileName PWSTR) (BOOL) = api-ms-win-core-file-fromapp-l1-1-0.MoveFileFromAppW
//sys	RemoveDirectoryFromAppW(/*in*/lpPathName PWSTR) (BOOL) = api-ms-win-core-file-fromapp-l1-1-0.RemoveDirectoryFromAppW
//sys	ReplaceFileFromAppW(/*in*/lpReplacedFileName PWSTR, /*in*/lpReplacementFileName PWSTR, /*in*//*optional*/lpBackupFileName PWSTR, /*in*/dwReplaceFlags uint32, /*in*//*out*/lpExclude unsafe.Pointer, /*in*//*out*/lpReserved unsafe.Pointer) (BOOL) = api-ms-win-core-file-fromapp-l1-1-0.ReplaceFileFromAppW
//sys	SetFileAttributesFromAppW(/*in*/lpFileName PWSTR, /*in*/dwFileAttributes uint32) (BOOL) = api-ms-win-core-file-fromapp-l1-1-0.SetFileAttributesFromAppW
//sys	VerFindFileA(/*in*/uFlags VER_FIND_FILE_FLAGS, /*in*/szFileName PSTR, /*in*//*optional*/szWinDir PSTR, /*in*/szAppDir PSTR, /*out*/szCurDir PSTR, /*in*//*out*/puCurDirLen *uint32, /*out*/szDestDir PSTR, /*in*//*out*/puDestDirLen *uint32) (VER_FIND_FILE_STATUS) = version.VerFindFileA
//sys	VerFindFileW(/*in*/uFlags VER_FIND_FILE_FLAGS, /*in*/szFileName PWSTR, /*in*//*optional*/szWinDir PWSTR, /*in*/szAppDir PWSTR, /*out*/szCurDir PWSTR, /*in*//*out*/puCurDirLen *uint32, /*out*/szDestDir PWSTR, /*in*//*out*/puDestDirLen *uint32) (VER_FIND_FILE_STATUS) = version.VerFindFileW
//sys	VerInstallFileA(/*in*/uFlags VER_INSTALL_FILE_FLAGS, /*in*/szSrcFileName PSTR, /*in*/szDestFileName PSTR, /*in*/szSrcDir PSTR, /*in*/szDestDir PSTR, /*in*/szCurDir PSTR, /*out*/szTmpFile PSTR, /*in*//*out*/puTmpFileLen *uint32) (VER_INSTALL_FILE_STATUS) = version.VerInstallFileA
//sys	VerInstallFileW(/*in*/uFlags VER_INSTALL_FILE_FLAGS, /*in*/szSrcFileName PWSTR, /*in*/szDestFileName PWSTR, /*in*/szSrcDir PWSTR, /*in*/szDestDir PWSTR, /*in*/szCurDir PWSTR, /*out*/szTmpFile PWSTR, /*in*//*out*/puTmpFileLen *uint32) (VER_INSTALL_FILE_STATUS) = version.VerInstallFileW
//sys	GetFileVersionInfoSizeA(/*in*/lptstrFilename PSTR, /*out*//*optional*/lpdwHandle *uint32) (uint32) = version.GetFileVersionInfoSizeA
//sys	GetFileVersionInfoSizeW(/*in*/lptstrFilename PWSTR, /*out*//*optional*/lpdwHandle *uint32) (uint32) = version.GetFileVersionInfoSizeW
//sys	GetFileVersionInfoA(/*in*/lptstrFilename PSTR, /*in*/dwHandle uint32, /*in*/dwLen uint32, /*out*/lpData unsafe.Pointer) (BOOL) = version.GetFileVersionInfoA
//sys	GetFileVersionInfoW(/*in*/lptstrFilename PWSTR, /*in*/dwHandle uint32, /*in*/dwLen uint32, /*out*/lpData unsafe.Pointer) (BOOL) = version.GetFileVersionInfoW
//sys	GetFileVersionInfoSizeExA(/*in*/dwFlags GET_FILE_VERSION_INFO_FLAGS, /*in*/lpwstrFilename PSTR, /*out*/lpdwHandle *uint32) (uint32) = version.GetFileVersionInfoSizeExA
//sys	GetFileVersionInfoSizeExW(/*in*/dwFlags GET_FILE_VERSION_INFO_FLAGS, /*in*/lpwstrFilename PWSTR, /*out*/lpdwHandle *uint32) (uint32) = version.GetFileVersionInfoSizeExW
//sys	GetFileVersionInfoExA(/*in*/dwFlags GET_FILE_VERSION_INFO_FLAGS, /*in*/lpwstrFilename PSTR, /*in*/dwHandle uint32, /*in*/dwLen uint32, /*out*/lpData unsafe.Pointer) (BOOL) = version.GetFileVersionInfoExA
//sys	GetFileVersionInfoExW(/*in*/dwFlags GET_FILE_VERSION_INFO_FLAGS, /*in*/lpwstrFilename PWSTR, /*in*/dwHandle uint32, /*in*/dwLen uint32, /*out*/lpData unsafe.Pointer) (BOOL) = version.GetFileVersionInfoExW
//sys	VerLanguageNameA(/*in*/wLang uint32, /*out*/szLang PSTR, /*in*/cchLang uint32) (uint32)
//sys	VerLanguageNameW(/*in*/wLang uint32, /*out*/szLang PWSTR, /*in*/cchLang uint32) (uint32)
//sys	VerQueryValueA(/*in*/pBlock unsafe.Pointer, /*in*/lpSubBlock PSTR, /*out*/lplpBuffer *unsafe.Pointer, /*out*/puLen *uint32) (BOOL) = version.VerQueryValueA
//sys	VerQueryValueW(/*in*/pBlock unsafe.Pointer, /*in*/lpSubBlock PWSTR, /*out*/lplpBuffer *unsafe.Pointer, /*out*/puLen *uint32) (BOOL) = version.VerQueryValueW
//sys	LsnEqual(/*in*/plsn1 *CLS_LSN, /*in*/plsn2 *CLS_LSN) (BOOLEAN) = clfsw32.LsnEqual
//sys	LsnLess(/*in*/plsn1 *CLS_LSN, /*in*/plsn2 *CLS_LSN) (BOOLEAN) = clfsw32.LsnLess
//sys	LsnGreater(/*in*/plsn1 *CLS_LSN, /*in*/plsn2 *CLS_LSN) (BOOLEAN) = clfsw32.LsnGreater
//sys	LsnNull(/*in*/plsn *CLS_LSN) (BOOLEAN) = clfsw32.LsnNull
//sys	LsnContainer(/*in*/plsn *CLS_LSN) (uint32) = clfsw32.LsnContainer
//sys	LsnCreate(/*in*/cidContainer uint32, /*in*/offBlock uint32, /*in*/cRecord uint32) (CLS_LSN) = clfsw32.LsnCreate
//sys	LsnBlockOffset(/*in*/plsn *CLS_LSN) (uint32) = clfsw32.LsnBlockOffset
//sys	LsnRecordSequence(/*in*/plsn *CLS_LSN) (uint32) = clfsw32.LsnRecordSequence
//sys	LsnInvalid(/*in*/plsn *CLS_LSN) (BOOLEAN) = clfsw32.LsnInvalid
//sys	LsnIncrement(/*in*/plsn *CLS_LSN) (CLS_LSN) = clfsw32.LsnIncrement
//sys	CreateLogFile(/*in*/pszLogFileName PWSTR, /*in*/fDesiredAccess FILE_ACCESS_FLAGS, /*in*/dwShareMode FILE_SHARE_MODE, /*in*//*out*/psaLogFile *SECURITY_ATTRIBUTES, /*in*/fCreateDisposition FILE_CREATION_DISPOSITION, /*in*/fFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES) (HANDLE) = clfsw32.CreateLogFile
//sys	DeleteLogByHandle(/*in*/hLog HANDLE) (BOOL) = clfsw32.DeleteLogByHandle
//sys	DeleteLogFile(/*in*/pszLogFileName PWSTR, /*in*//*out*/pvReserved unsafe.Pointer) (BOOL) = clfsw32.DeleteLogFile
//sys	AddLogContainer(/*in*/hLog HANDLE, /*in*//*optional*/pcbContainer *uint64, /*in*/pwszContainerPath PWSTR, /*in*//*out*//*optional*/pReserved unsafe.Pointer) (BOOL) = clfsw32.AddLogContainer
//sys	AddLogContainerSet(/*in*/hLog HANDLE, /*in*/cContainer uint16, /*in*//*optional*/pcbContainer *uint64, /*in*/rgwszContainerPath *PWSTR, /*in*//*out*//*optional*/pReserved unsafe.Pointer) (BOOL) = clfsw32.AddLogContainerSet
//sys	RemoveLogContainer(/*in*/hLog HANDLE, /*in*/pwszContainerPath PWSTR, /*in*/fForce BOOL, /*in*//*out*//*optional*/pReserved unsafe.Pointer) (BOOL) = clfsw32.RemoveLogContainer
//sys	RemoveLogContainerSet(/*in*/hLog HANDLE, /*in*/cContainer uint16, /*in*/rgwszContainerPath *PWSTR, /*in*/fForce BOOL, /*in*//*out*//*optional*/pReserved unsafe.Pointer) (BOOL) = clfsw32.RemoveLogContainerSet
//sys	SetLogArchiveTail(/*in*/hLog HANDLE, /*in*//*out*/plsnArchiveTail *CLS_LSN, /*in*//*out*/pReserved unsafe.Pointer) (BOOL) = clfsw32.SetLogArchiveTail
//sys	SetEndOfLog(/*in*/hLog HANDLE, /*in*//*out*/plsnEnd *CLS_LSN, /*in*//*out*/lpOverlapped *OVERLAPPED) (BOOL) = clfsw32.SetEndOfLog
//sys	TruncateLog(/*in*/pvMarshal unsafe.Pointer, /*in*/plsnEnd *CLS_LSN, /*in*//*out*//*optional*/lpOverlapped *OVERLAPPED) (BOOL) = clfsw32.TruncateLog
//sys	CreateLogContainerScanContext(/*in*/hLog HANDLE, /*in*/cFromContainer uint32, /*in*/cContainers uint32, /*in*/eScanMode uint8, /*in*//*out*/pcxScan *CLS_SCAN_CONTEXT, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.CreateLogContainerScanContext
//sys	ScanLogContainers(/*in*//*out*/pcxScan *CLS_SCAN_CONTEXT, /*in*/eScanMode uint8, /*in*//*out*/pReserved unsafe.Pointer) (BOOL) = clfsw32.ScanLogContainers
//sys	AlignReservedLog(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*/cReservedRecords uint32, /*in*//*out*/rgcbReservation *int64, /*in*//*out*/pcbAlignReservation *int64) (BOOL) = clfsw32.AlignReservedLog
//sys	AllocReservedLog(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*/cReservedRecords uint32, /*in*//*out*/pcbAdjustment *int64) (BOOL) = clfsw32.AllocReservedLog
//sys	FreeReservedLog(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*/cReservedRecords uint32, /*in*//*out*/pcbAdjustment *int64) (BOOL) = clfsw32.FreeReservedLog
//sys	GetLogFileInformation(/*in*/hLog HANDLE, /*in*//*out*/pinfoBuffer *CLS_INFORMATION, /*in*//*out*/cbBuffer *uint32) (BOOL) = clfsw32.GetLogFileInformation
//sys	SetLogArchiveMode(/*in*/hLog HANDLE, /*in*/eMode CLFS_LOG_ARCHIVE_MODE) (BOOL) = clfsw32.SetLogArchiveMode
//sys	ReadLogRestartArea(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*//*out*/ppvRestartBuffer *unsafe.Pointer, /*in*//*out*/pcbRestartBuffer *uint32, /*in*//*out*/plsn *CLS_LSN, /*in*//*out*/ppvContext *unsafe.Pointer, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.ReadLogRestartArea
//sys	ReadPreviousLogRestartArea(/*in*//*out*/pvReadContext unsafe.Pointer, /*in*//*out*/ppvRestartBuffer *unsafe.Pointer, /*in*//*out*/pcbRestartBuffer *uint32, /*in*//*out*/plsnRestart *CLS_LSN, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.ReadPreviousLogRestartArea
//sys	WriteLogRestartArea(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*//*out*/pvRestartBuffer unsafe.Pointer, /*in*/cbRestartBuffer uint32, /*in*//*out*/plsnBase *CLS_LSN, /*in*/fFlags CLFS_FLAG, /*in*//*out*/pcbWritten *uint32, /*in*//*out*/plsnNext *CLS_LSN, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.WriteLogRestartArea
//sys	GetLogReservationInfo(/*in*/pvMarshal unsafe.Pointer, /*out*/pcbRecordNumber *uint32, /*out*/pcbUserReservation *int64, /*out*/pcbCommitReservation *int64) (BOOL) = clfsw32.GetLogReservationInfo
//sys	AdvanceLogBase(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*//*out*/plsnBase *CLS_LSN, /*in*/fFlags uint32, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.AdvanceLogBase
//sys	CloseAndResetLogFile(/*in*/hLog HANDLE) (BOOL) = clfsw32.CloseAndResetLogFile
//sys	CreateLogMarshallingArea(/*in*/hLog HANDLE, /*in*/pfnAllocBuffer CLFS_BLOCK_ALLOCATION, /*in*/pfnFreeBuffer CLFS_BLOCK_DEALLOCATION, /*in*//*out*/pvBlockAllocContext unsafe.Pointer, /*in*/cbMarshallingBuffer uint32, /*in*/cMaxWriteBuffers uint32, /*in*/cMaxReadBuffers uint32, /*in*//*out*/ppvMarshal *unsafe.Pointer) (BOOL) = clfsw32.CreateLogMarshallingArea
//sys	DeleteLogMarshallingArea(/*in*//*out*/pvMarshal unsafe.Pointer) (BOOL) = clfsw32.DeleteLogMarshallingArea
//sys	ReserveAndAppendLog(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*//*out*/rgWriteEntries *CLS_WRITE_ENTRY, /*in*/cWriteEntries uint32, /*in*//*out*/plsnUndoNext *CLS_LSN, /*in*//*out*/plsnPrevious *CLS_LSN, /*in*/cReserveRecords uint32, /*in*//*out*/rgcbReservation *int64, /*in*/fFlags CLFS_FLAG, /*in*//*out*/plsn *CLS_LSN, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.ReserveAndAppendLog
//sys	ReserveAndAppendLogAligned(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*//*out*/rgWriteEntries *CLS_WRITE_ENTRY, /*in*/cWriteEntries uint32, /*in*/cbEntryAlignment uint32, /*in*//*out*/plsnUndoNext *CLS_LSN, /*in*//*out*/plsnPrevious *CLS_LSN, /*in*/cReserveRecords uint32, /*in*//*out*/rgcbReservation *int64, /*in*/fFlags CLFS_FLAG, /*in*//*out*/plsn *CLS_LSN, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.ReserveAndAppendLogAligned
//sys	FlushLogBuffers(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.FlushLogBuffers
//sys	FlushLogToLsn(/*in*//*out*/pvMarshalContext unsafe.Pointer, /*in*//*out*/plsnFlush *CLS_LSN, /*in*//*out*/plsnLastFlushed *CLS_LSN, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.FlushLogToLsn
//sys	ReadLogRecord(/*in*//*out*/pvMarshal unsafe.Pointer, /*in*//*out*/plsnFirst *CLS_LSN, /*in*/eContextMode CLFS_CONTEXT_MODE, /*in*//*out*/ppvReadBuffer *unsafe.Pointer, /*in*//*out*/pcbReadBuffer *uint32, /*in*//*out*/peRecordType *uint8, /*in*//*out*/plsnUndoNext *CLS_LSN, /*in*//*out*/plsnPrevious *CLS_LSN, /*in*//*out*/ppvReadContext *unsafe.Pointer, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.ReadLogRecord
//sys	ReadNextLogRecord(/*in*//*out*/pvReadContext unsafe.Pointer, /*in*//*out*/ppvBuffer *unsafe.Pointer, /*in*//*out*/pcbBuffer *uint32, /*in*//*out*/peRecordType *uint8, /*in*//*out*/plsnUser *CLS_LSN, /*in*//*out*/plsnUndoNext *CLS_LSN, /*in*//*out*/plsnPrevious *CLS_LSN, /*in*//*out*/plsnRecord *CLS_LSN, /*in*//*out*/pOverlapped *OVERLAPPED) (BOOL) = clfsw32.ReadNextLogRecord
//sys	TerminateReadLog(/*in*//*out*/pvCursorContext unsafe.Pointer) (BOOL) = clfsw32.TerminateReadLog
//sys	PrepareLogArchive(/*in*/hLog HANDLE, /*in*//*out*/pszBaseLogFileName PWSTR, /*in*/cLen uint32, /*in*//*optional*/plsnLow *CLS_LSN, /*in*//*optional*/plsnHigh *CLS_LSN, /*out*//*optional*/pcActualLength *uint32, /*out*/poffBaseLogFileData *uint64, /*out*/pcbBaseLogFileLength *uint64, /*out*/plsnBase *CLS_LSN, /*out*/plsnLast *CLS_LSN, /*out*/plsnCurrentArchiveTail *CLS_LSN, /*out*/ppvArchiveContext *unsafe.Pointer) (BOOL) = clfsw32.PrepareLogArchive
//sys	ReadLogArchiveMetadata(/*in*//*out*/pvArchiveContext unsafe.Pointer, /*in*/cbOffset uint32, /*in*/cbBytesToRead uint32, /*in*//*out*/pbReadBuffer *uint8, /*in*//*out*/pcbBytesRead *uint32) (BOOL) = clfsw32.ReadLogArchiveMetadata
//sys	GetNextLogArchiveExtent(/*in*//*out*/pvArchiveContext unsafe.Pointer, /*in*//*out*/rgadExtent *CLS_ARCHIVE_DESCRIPTOR, /*in*/cDescriptors uint32, /*in*//*out*/pcDescriptorsReturned *uint32) (BOOL) = clfsw32.GetNextLogArchiveExtent
//sys	TerminateLogArchive(/*in*//*out*/pvArchiveContext unsafe.Pointer) (BOOL) = clfsw32.TerminateLogArchive
//sys	ValidateLog(/*in*/pszLogFileName PWSTR, /*in*//*out*/psaLogFile *SECURITY_ATTRIBUTES, /*in*//*out*/pinfoBuffer *CLS_INFORMATION, /*in*//*out*/pcbBuffer *uint32) (BOOL) = clfsw32.ValidateLog
//sys	GetLogContainerName(/*in*/hLog HANDLE, /*in*/cidLogicalContainer uint32, /*in*/pwstrContainerName PWSTR, /*in*/cLenContainerName uint32, /*in*//*out*/pcActualLenContainerName *uint32) (BOOL) = clfsw32.GetLogContainerName
//sys	GetLogIoStatistics(/*in*/hLog HANDLE, /*in*//*out*/pvStatsBuffer unsafe.Pointer, /*in*/cbStatsBuffer uint32, /*in*/eStatsClass CLFS_IOSTATS_CLASS, /*in*//*out*/pcbStatsWritten *uint32) (BOOL) = clfsw32.GetLogIoStatistics
//sys	RegisterManageableLogClient(/*in*/hLog HANDLE, /*in*//*out*/pCallbacks *LOG_MANAGEMENT_CALLBACKS) (BOOL) = clfsw32.RegisterManageableLogClient
//sys	DeregisterManageableLogClient(/*in*/hLog HANDLE) (BOOL) = clfsw32.DeregisterManageableLogClient
//sys	ReadLogNotification(/*in*/hLog HANDLE, /*in*//*out*/pNotification *CLFS_MGMT_NOTIFICATION, /*in*//*out*/lpOverlapped *OVERLAPPED) (BOOL) = clfsw32.ReadLogNotification
//sys	InstallLogPolicy(/*in*/hLog HANDLE, /*in*//*out*/pPolicy *CLFS_MGMT_POLICY) (BOOL) = clfsw32.InstallLogPolicy
//sys	RemoveLogPolicy(/*in*/hLog HANDLE, /*in*/ePolicyType CLFS_MGMT_POLICY_TYPE) (BOOL) = clfsw32.RemoveLogPolicy
//sys	QueryLogPolicy(/*in*/hLog HANDLE, /*in*/ePolicyType CLFS_MGMT_POLICY_TYPE, /*in*//*out*/pPolicyBuffer *CLFS_MGMT_POLICY, /*in*//*out*/pcbPolicyBuffer *uint32) (BOOL) = clfsw32.QueryLogPolicy
//sys	SetLogFileSizeWithPolicy(/*in*/hLog HANDLE, /*in*//*out*/pDesiredSize *uint64, /*in*//*out*/pResultingSize *uint64) (BOOL) = clfsw32.SetLogFileSizeWithPolicy
//sys	HandleLogFull(/*in*/hLog HANDLE) (BOOL) = clfsw32.HandleLogFull
//sys	LogTailAdvanceFailure(/*in*/hLog HANDLE, /*in*/dwReason uint32) (BOOL) = clfsw32.LogTailAdvanceFailure
//sys	RegisterForLogWriteNotification(/*in*/hLog HANDLE, /*in*/cbThreshold uint32, /*in*/fEnable BOOL) (BOOL) = clfsw32.RegisterForLogWriteNotification
//sys	QueryUsersOnEncryptedFile(/*in*/lpFileName PWSTR, /*out*/pUsers **ENCRYPTION_CERTIFICATE_HASH_LIST) (uint32) = advapi32.QueryUsersOnEncryptedFile
//sys	QueryRecoveryAgentsOnEncryptedFile(/*in*/lpFileName PWSTR, /*out*/pRecoveryAgents **ENCRYPTION_CERTIFICATE_HASH_LIST) (uint32) = advapi32.QueryRecoveryAgentsOnEncryptedFile
//sys	RemoveUsersFromEncryptedFile(/*in*/lpFileName PWSTR, /*in*/pHashes *ENCRYPTION_CERTIFICATE_HASH_LIST) (uint32) = advapi32.RemoveUsersFromEncryptedFile
//sys	AddUsersToEncryptedFile(/*in*/lpFileName PWSTR, /*in*/pEncryptionCertificates *ENCRYPTION_CERTIFICATE_LIST) (uint32) = advapi32.AddUsersToEncryptedFile
//sys	SetUserFileEncryptionKey(/*in*//*optional*/pEncryptionCertificate *ENCRYPTION_CERTIFICATE) (uint32) = advapi32.SetUserFileEncryptionKey
//sys	SetUserFileEncryptionKeyEx(/*in*//*optional*/pEncryptionCertificate *ENCRYPTION_CERTIFICATE, /*in*/dwCapabilities uint32, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer) (uint32) = advapi32.SetUserFileEncryptionKeyEx
//sys	FreeEncryptionCertificateHashList(/*in*/pUsers *ENCRYPTION_CERTIFICATE_HASH_LIST) = advapi32.FreeEncryptionCertificateHashList
//sys	EncryptionDisable(/*in*/DirPath PWSTR, /*in*/Disable BOOL) (BOOL) = advapi32.EncryptionDisable
//sys	DuplicateEncryptionInfoFile(/*in*/SrcFileName PWSTR, /*in*/DstFileName PWSTR, /*in*/dwCreationDistribution uint32, /*in*/dwAttributes uint32, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES) (uint32) = advapi32.DuplicateEncryptionInfoFile
//sys	GetEncryptedFileMetadata(/*in*/lpFileName PWSTR, /*out*/pcbMetadata *uint32, /*out*/ppbMetadata **uint8) (uint32) = advapi32.GetEncryptedFileMetadata
//sys	SetEncryptedFileMetadata(/*in*/lpFileName PWSTR, /*in*//*optional*/pbOldMetadata *uint8, /*in*/pbNewMetadata *uint8, /*in*/pOwnerHash *ENCRYPTION_CERTIFICATE_HASH, /*in*/dwOperation uint32, /*in*//*optional*/pCertificatesAdded *ENCRYPTION_CERTIFICATE_HASH_LIST) (uint32) = advapi32.SetEncryptedFileMetadata
//sys	FreeEncryptedFileMetadata(/*in*/pbMetadata *uint8) = advapi32.FreeEncryptedFileMetadata
//sys	LZStart() (int32)
//sys	LZDone()
//sys	CopyLZFile(/*in*/hfSource int32, /*in*/hfDest int32) (int32)
//sys	LZCopy(/*in*/hfSource int32, /*in*/hfDest int32) (int32)
//sys	LZInit(/*in*/hfSource int32) (int32)
//sys	GetExpandedNameA(/*in*/lpszSource PSTR, /*out*/lpszBuffer PSTR) (int32)
//sys	GetExpandedNameW(/*in*/lpszSource PWSTR, /*out*/lpszBuffer PWSTR) (int32)
//sys	LZOpenFileA(/*in*/lpFileName PSTR, /*in*//*out*/lpReOpenBuf *OFSTRUCT, /*in*/wStyle LZOPENFILE_STYLE) (int32)
//sys	LZOpenFileW(/*in*/lpFileName PWSTR, /*in*//*out*/lpReOpenBuf *OFSTRUCT, /*in*/wStyle LZOPENFILE_STYLE) (int32)
//sys	LZSeek(/*in*/hFile int32, /*in*/lOffset int32, /*in*/iOrigin int32) (int32)
//sys	LZRead(/*in*/hFile int32, /*out*/lpBuffer PSTR, /*in*/cbRead int32) (int32)
//sys	LZClose(/*in*/hFile int32)
//sys	WofShouldCompressBinaries(/*in*/Volume PWSTR, /*out*/Algorithm *uint32) (BOOL) = wofutil.WofShouldCompressBinaries
//sys	WofGetDriverVersion(/*in*/FileOrVolumeHandle HANDLE, /*in*/Provider uint32, /*out*/WofVersion *uint32) (HRESULT) = wofutil.WofGetDriverVersion
//sys	WofSetFileDataLocation(/*in*/FileHandle HANDLE, /*in*/Provider uint32, /*in*/ExternalFileInfo unsafe.Pointer, /*in*/Length uint32) (HRESULT) = wofutil.WofSetFileDataLocation
//sys	WofIsExternalFile(/*in*/FilePath PWSTR, /*out*//*optional*/IsExternalFile *BOOL, /*out*//*optional*/Provider *uint32, /*out*//*optional*/ExternalFileInfo unsafe.Pointer, /*in*//*out*//*optional*/BufferLength *uint32) (HRESULT) = wofutil.WofIsExternalFile
//sys	WofEnumEntries(/*in*/VolumeName PWSTR, /*in*/Provider uint32, /*in*/EnumProc WofEnumEntryProc, /*in*//*optional*/UserData unsafe.Pointer) (HRESULT) = wofutil.WofEnumEntries
//sys	WofWimAddEntry(/*in*/VolumeName PWSTR, /*in*/WimPath PWSTR, /*in*/WimType uint32, /*in*/WimIndex uint32, /*out*/DataSourceId *LARGE_INTEGER) (HRESULT) = wofutil.WofWimAddEntry
//sys	WofWimEnumFiles(/*in*/VolumeName PWSTR, /*in*/DataSourceId LARGE_INTEGER, /*in*/EnumProc WofEnumFilesProc, /*in*//*optional*/UserData unsafe.Pointer) (HRESULT) = wofutil.WofWimEnumFiles
//sys	WofWimSuspendEntry(/*in*/VolumeName PWSTR, /*in*/DataSourceId LARGE_INTEGER) (HRESULT) = wofutil.WofWimSuspendEntry
//sys	WofWimRemoveEntry(/*in*/VolumeName PWSTR, /*in*/DataSourceId LARGE_INTEGER) (HRESULT) = wofutil.WofWimRemoveEntry
//sys	WofWimUpdateEntry(/*in*/VolumeName PWSTR, /*in*/DataSourceId LARGE_INTEGER, /*in*/NewWimPath PWSTR) (HRESULT) = wofutil.WofWimUpdateEntry
//sys	WofFileEnumFiles(/*in*/VolumeName PWSTR, /*in*/Algorithm uint32, /*in*/EnumProc WofEnumFilesProc, /*in*//*optional*/UserData unsafe.Pointer) (HRESULT) = wofutil.WofFileEnumFiles
//sys	TxfLogCreateFileReadContext(/*in*/LogPath PWSTR, /*in*/BeginningLsn CLS_LSN, /*in*/EndingLsn CLS_LSN, /*in*/TxfFileId *TXF_ID, /*out*/TxfLogContext *unsafe.Pointer) (BOOL) = txfw32.TxfLogCreateFileReadContext
//sys	TxfLogCreateRangeReadContext(/*in*/LogPath PWSTR, /*in*/BeginningLsn CLS_LSN, /*in*/EndingLsn CLS_LSN, /*in*/BeginningVirtualClock *LARGE_INTEGER, /*in*/EndingVirtualClock *LARGE_INTEGER, /*in*/RecordTypeMask uint32, /*out*/TxfLogContext *unsafe.Pointer) (BOOL) = txfw32.TxfLogCreateRangeReadContext
//sys	TxfLogDestroyReadContext(/*in*/TxfLogContext unsafe.Pointer) (BOOL) = txfw32.TxfLogDestroyReadContext
//sys	TxfLogReadRecords(/*in*/TxfLogContext unsafe.Pointer, /*in*/BufferLength uint32, /*out*/Buffer unsafe.Pointer, /*out*/BytesUsed *uint32, /*out*/RecordCount *uint32) (BOOL) = txfw32.TxfLogReadRecords
//sys	TxfReadMetadataInfo(/*in*/FileHandle HANDLE, /*out*/TxfFileId *TXF_ID, /*out*/LastLsn *CLS_LSN, /*out*/TransactionState *uint32, /*out*/LockingTransaction *Guid) (BOOL) = txfw32.TxfReadMetadataInfo
//sys	TxfLogRecordGetFileName(/*in*/RecordBuffer unsafe.Pointer, /*in*/RecordBufferLengthInBytes uint32, /*out*/NameBuffer PWSTR, /*in*//*out*/NameBufferLengthInBytes *uint32, /*out*//*optional*/TxfId *TXF_ID) (BOOL) = txfw32.TxfLogRecordGetFileName
//sys	TxfLogRecordGetGenericType(/*in*/RecordBuffer unsafe.Pointer, /*in*/RecordBufferLengthInBytes uint32, /*out*/GenericType *uint32, /*out*//*optional*/VirtualClock *LARGE_INTEGER) (BOOL) = txfw32.TxfLogRecordGetGenericType
//sys	TxfSetThreadMiniVersionForCreate(/*in*/MiniVersion uint16) = txfw32.TxfSetThreadMiniVersionForCreate
//sys	TxfGetThreadMiniVersionForCreate(/*out*/MiniVersion *uint16) = txfw32.TxfGetThreadMiniVersionForCreate
//sys	CreateTransaction(/*in*//*out*/lpTransactionAttributes *SECURITY_ATTRIBUTES, /*in*//*out*/UOW *Guid, /*in*/CreateOptions uint32, /*in*/IsolationLevel uint32, /*in*/IsolationFlags uint32, /*in*/Timeout uint32, /*in*//*optional*/Description PWSTR) (HANDLE) = ktmw32.CreateTransaction
//sys	OpenTransaction(/*in*/dwDesiredAccess uint32, /*in*//*out*/TransactionId *Guid) (HANDLE) = ktmw32.OpenTransaction
//sys	CommitTransaction(/*in*/TransactionHandle HANDLE) (BOOL) = ktmw32.CommitTransaction
//sys	CommitTransactionAsync(/*in*/TransactionHandle HANDLE) (BOOL) = ktmw32.CommitTransactionAsync
//sys	RollbackTransaction(/*in*/TransactionHandle HANDLE) (BOOL) = ktmw32.RollbackTransaction
//sys	RollbackTransactionAsync(/*in*/TransactionHandle HANDLE) (BOOL) = ktmw32.RollbackTransactionAsync
//sys	GetTransactionId(/*in*/TransactionHandle HANDLE, /*in*//*out*/TransactionId *Guid) (BOOL) = ktmw32.GetTransactionId
//sys	GetTransactionInformation(/*in*/TransactionHandle HANDLE, /*in*//*out*/Outcome *uint32, /*in*//*out*/IsolationLevel *uint32, /*in*//*out*/IsolationFlags *uint32, /*in*//*out*/Timeout *uint32, /*in*/BufferLength uint32, /*out*//*optional*/Description PWSTR) (BOOL) = ktmw32.GetTransactionInformation
//sys	SetTransactionInformation(/*in*/TransactionHandle HANDLE, /*in*/IsolationLevel uint32, /*in*/IsolationFlags uint32, /*in*/Timeout uint32, /*in*//*optional*/Description PWSTR) (BOOL) = ktmw32.SetTransactionInformation
//sys	CreateTransactionManager(/*in*//*out*/lpTransactionAttributes *SECURITY_ATTRIBUTES, /*in*/LogFileName PWSTR, /*in*/CreateOptions uint32, /*in*/CommitStrength uint32) (HANDLE) = ktmw32.CreateTransactionManager
//sys	OpenTransactionManager(/*in*/LogFileName PWSTR, /*in*/DesiredAccess uint32, /*in*/OpenOptions uint32) (HANDLE) = ktmw32.OpenTransactionManager
//sys	OpenTransactionManagerById(/*in*/TransactionManagerId *Guid, /*in*/DesiredAccess uint32, /*in*/OpenOptions uint32) (HANDLE) = ktmw32.OpenTransactionManagerById
//sys	RenameTransactionManager(/*in*/LogFileName PWSTR, /*in*//*out*/ExistingTransactionManagerGuid *Guid) (BOOL) = ktmw32.RenameTransactionManager
//sys	RollforwardTransactionManager(/*in*/TransactionManagerHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.RollforwardTransactionManager
//sys	RecoverTransactionManager(/*in*/TransactionManagerHandle HANDLE) (BOOL) = ktmw32.RecoverTransactionManager
//sys	GetCurrentClockTransactionManager(/*in*/TransactionManagerHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.GetCurrentClockTransactionManager
//sys	GetTransactionManagerId(/*in*/TransactionManagerHandle HANDLE, /*in*//*out*/TransactionManagerId *Guid) (BOOL) = ktmw32.GetTransactionManagerId
//sys	CreateResourceManager(/*in*//*out*/lpResourceManagerAttributes *SECURITY_ATTRIBUTES, /*in*//*out*/ResourceManagerId *Guid, /*in*/CreateOptions uint32, /*in*/TmHandle HANDLE, /*in*//*optional*/Description PWSTR) (HANDLE) = ktmw32.CreateResourceManager
//sys	OpenResourceManager(/*in*/dwDesiredAccess uint32, /*in*/TmHandle HANDLE, /*in*//*out*/ResourceManagerId *Guid) (HANDLE) = ktmw32.OpenResourceManager
//sys	RecoverResourceManager(/*in*/ResourceManagerHandle HANDLE) (BOOL) = ktmw32.RecoverResourceManager
//sys	GetNotificationResourceManager(/*in*/ResourceManagerHandle HANDLE, /*in*//*out*/TransactionNotification *TRANSACTION_NOTIFICATION, /*in*/NotificationLength uint32, /*in*/dwMilliseconds uint32, /*in*//*out*/ReturnLength *uint32) (BOOL) = ktmw32.GetNotificationResourceManager
//sys	GetNotificationResourceManagerAsync(/*in*/ResourceManagerHandle HANDLE, /*in*//*out*/TransactionNotification *TRANSACTION_NOTIFICATION, /*in*/TransactionNotificationLength uint32, /*in*//*out*/ReturnLength *uint32, /*in*//*out*/lpOverlapped *OVERLAPPED) (BOOL) = ktmw32.GetNotificationResourceManagerAsync
//sys	SetResourceManagerCompletionPort(/*in*/ResourceManagerHandle HANDLE, /*in*/IoCompletionPortHandle HANDLE, /*in*/CompletionKey uintptr) (BOOL) = ktmw32.SetResourceManagerCompletionPort
//sys	CreateEnlistment(/*in*//*out*/lpEnlistmentAttributes *SECURITY_ATTRIBUTES, /*in*/ResourceManagerHandle HANDLE, /*in*/TransactionHandle HANDLE, /*in*/NotificationMask uint32, /*in*/CreateOptions uint32, /*in*//*out*/EnlistmentKey unsafe.Pointer) (HANDLE) = ktmw32.CreateEnlistment
//sys	OpenEnlistment(/*in*/dwDesiredAccess uint32, /*in*/ResourceManagerHandle HANDLE, /*in*//*out*/EnlistmentId *Guid) (HANDLE) = ktmw32.OpenEnlistment
//sys	RecoverEnlistment(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/EnlistmentKey unsafe.Pointer) (BOOL) = ktmw32.RecoverEnlistment
//sys	GetEnlistmentRecoveryInformation(/*in*/EnlistmentHandle HANDLE, /*in*/BufferSize uint32, /*in*//*out*/Buffer unsafe.Pointer, /*in*//*out*/BufferUsed *uint32) (BOOL) = ktmw32.GetEnlistmentRecoveryInformation
//sys	GetEnlistmentId(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/EnlistmentId *Guid) (BOOL) = ktmw32.GetEnlistmentId
//sys	SetEnlistmentRecoveryInformation(/*in*/EnlistmentHandle HANDLE, /*in*/BufferSize uint32, /*in*//*out*/Buffer unsafe.Pointer) (BOOL) = ktmw32.SetEnlistmentRecoveryInformation
//sys	PrepareEnlistment(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.PrepareEnlistment
//sys	PrePrepareEnlistment(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.PrePrepareEnlistment
//sys	CommitEnlistment(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.CommitEnlistment
//sys	RollbackEnlistment(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.RollbackEnlistment
//sys	PrePrepareComplete(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.PrePrepareComplete
//sys	PrepareComplete(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.PrepareComplete
//sys	ReadOnlyEnlistment(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.ReadOnlyEnlistment
//sys	CommitComplete(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.CommitComplete
//sys	RollbackComplete(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.RollbackComplete
//sys	SinglePhaseReject(/*in*/EnlistmentHandle HANDLE, /*in*//*out*/TmVirtualClock *LARGE_INTEGER) (BOOL) = ktmw32.SinglePhaseReject
//sys	NetShareAdd(/*in*//*optional*/servername PWSTR, /*in*/level uint32, /*in*/buf *uint8, /*out*//*optional*/parm_err *uint32) (uint32) = netapi32.NetShareAdd
//sys	NetShareEnum(/*in*//*optional*/servername PWSTR, /*in*/level uint32, /*out*/bufptr **uint8, /*in*/prefmaxlen uint32, /*out*/entriesread *uint32, /*out*/totalentries *uint32, /*in*//*out*//*optional*/resume_handle *uint32) (uint32) = netapi32.NetShareEnum
//sys	NetShareEnumSticky(/*in*//*optional*/servername PWSTR, /*in*/level uint32, /*out*/bufptr **uint8, /*in*/prefmaxlen uint32, /*out*/entriesread *uint32, /*out*/totalentries *uint32, /*in*//*out*//*optional*/resume_handle *uint32) (uint32) = netapi32.NetShareEnumSticky
//sys	NetShareGetInfo(/*in*//*optional*/servername PWSTR, /*in*/netname PWSTR, /*in*/level uint32, /*out*/bufptr **uint8) (uint32) = netapi32.NetShareGetInfo
//sys	NetShareSetInfo(/*in*//*optional*/servername PWSTR, /*in*/netname PWSTR, /*in*/level uint32, /*in*/buf *uint8, /*out*//*optional*/parm_err *uint32) (uint32) = netapi32.NetShareSetInfo
//sys	NetShareDel(/*in*//*optional*/servername PWSTR, /*in*/netname PWSTR, /*in*/reserved uint32) (uint32) = netapi32.NetShareDel
//sys	NetShareDelSticky(/*in*//*optional*/servername PWSTR, /*in*/netname PWSTR, /*in*/reserved uint32) (uint32) = netapi32.NetShareDelSticky
//sys	NetShareCheck(/*in*//*optional*/servername PWSTR, /*in*/device PWSTR, /*out*/type *uint32) (uint32) = netapi32.NetShareCheck
//sys	NetShareDelEx(/*in*//*optional*/servername PWSTR, /*in*/level uint32, /*in*/buf *uint8) (uint32) = netapi32.NetShareDelEx
//sys	NetServerAliasAdd(/*in*//*optional*/servername PWSTR, /*in*/level uint32, /*in*/buf *uint8) (uint32) = netapi32.NetServerAliasAdd
//sys	NetServerAliasDel(/*in*//*optional*/servername PWSTR, /*in*/level uint32, /*in*/buf *uint8) (uint32) = netapi32.NetServerAliasDel
//sys	NetServerAliasEnum(/*in*//*optional*/servername PWSTR, /*in*/level uint32, /*out*/bufptr **uint8, /*in*/prefmaxlen uint32, /*out*/entriesread *uint32, /*out*/totalentries *uint32, /*in*//*out*//*optional*/resumehandle *uint32) (uint32) = netapi32.NetServerAliasEnum
//sys	NetSessionEnum(/*in*//*optional*/servername PWSTR, /*in*//*optional*/UncClientName PWSTR, /*in*//*optional*/username PWSTR, /*in*/level uint32, /*out*/bufptr **uint8, /*in*/prefmaxlen uint32, /*out*/entriesread *uint32, /*out*/totalentries *uint32, /*in*//*out*//*optional*/resume_handle *uint32) (uint32) = netapi32.NetSessionEnum
//sys	NetSessionDel(/*in*//*optional*/servername PWSTR, /*in*//*optional*/UncClientName PWSTR, /*in*//*optional*/username PWSTR) (uint32) = netapi32.NetSessionDel
//sys	NetSessionGetInfo(/*in*//*optional*/servername PWSTR, /*in*/UncClientName PWSTR, /*in*/username PWSTR, /*in*/level uint32, /*out*/bufptr **uint8) (uint32) = netapi32.NetSessionGetInfo
//sys	NetConnectionEnum(/*in*//*optional*/servername PWSTR, /*in*/qualifier PWSTR, /*in*/level uint32, /*out*/bufptr **uint8, /*in*/prefmaxlen uint32, /*out*/entriesread *uint32, /*out*/totalentries *uint32, /*in*//*out*//*optional*/resume_handle *uint32) (uint32) = netapi32.NetConnectionEnum
//sys	NetFileClose(/*in*//*optional*/servername PWSTR, /*in*/fileid uint32) (uint32) = netapi32.NetFileClose
//sys	NetFileEnum(/*in*//*optional*/servername PWSTR, /*in*//*optional*/basepath PWSTR, /*in*//*optional*/username PWSTR, /*in*/level uint32, /*out*/bufptr **uint8, /*in*/prefmaxlen uint32, /*out*/entriesread *uint32, /*out*/totalentries *uint32, /*in*//*out*//*optional*/resume_handle *uintptr) (uint32) = netapi32.NetFileEnum
//sys	NetFileGetInfo(/*in*//*optional*/servername PWSTR, /*in*/fileid uint32, /*in*/level uint32, /*out*/bufptr **uint8) (uint32) = netapi32.NetFileGetInfo
//sys	NetStatisticsGet(/*in*/ServerName *int8, /*in*/Service *int8, /*in*/Level uint32, /*in*/Options uint32, /*out*/Buffer **uint8) (uint32) = netapi32.NetStatisticsGet
//sys	QueryIoRingCapabilities(/*out*/capabilities *IORING_CAPABILITIES) (HRESULT) = api-ms-win-core-ioring-l1-1-0.QueryIoRingCapabilities
//sys	IsIoRingOpSupported(/*in*/ioRing *HIORING__, /*in*/op IORING_OP_CODE) (BOOL) = api-ms-win-core-ioring-l1-1-0.IsIoRingOpSupported
//sys	CreateIoRing(/*in*/ioringVersion IORING_VERSION, /*in*/flags IORING_CREATE_FLAGS, /*in*/submissionQueueSize uint32, /*in*/completionQueueSize uint32, /*out*/h **HIORING__) (HRESULT) = api-ms-win-core-ioring-l1-1-0.CreateIoRing
//sys	GetIoRingInfo(/*in*/ioRing *HIORING__, /*out*/info *IORING_INFO) (HRESULT) = api-ms-win-core-ioring-l1-1-0.GetIoRingInfo
//sys	SubmitIoRing(/*in*/ioRing *HIORING__, /*in*/waitOperations uint32, /*in*/milliseconds uint32, /*out*//*optional*/submittedEntries *uint32) (HRESULT) = api-ms-win-core-ioring-l1-1-0.SubmitIoRing
//sys	CloseIoRing(/*in*/ioRing *HIORING__) (HRESULT) = api-ms-win-core-ioring-l1-1-0.CloseIoRing
//sys	PopIoRingCompletion(/*in*/ioRing *HIORING__, /*out*/cqe *IORING_CQE) (HRESULT) = api-ms-win-core-ioring-l1-1-0.PopIoRingCompletion
//sys	SetIoRingCompletionEvent(/*in*/ioRing *HIORING__, /*in*/hEvent HANDLE) (HRESULT) = api-ms-win-core-ioring-l1-1-0.SetIoRingCompletionEvent
//sys	BuildIoRingCancelRequest(/*in*/ioRing *HIORING__, /*in*/file IORING_HANDLE_REF, /*in*/opToCancel uintptr, /*in*/userData uintptr) (HRESULT) = api-ms-win-core-ioring-l1-1-0.BuildIoRingCancelRequest
//sys	BuildIoRingReadFile(/*in*/ioRing *HIORING__, /*in*/fileRef IORING_HANDLE_REF, /*in*/dataRef IORING_BUFFER_REF, /*in*/numberOfBytesToRead uint32, /*in*/fileOffset uint64, /*in*/userData uintptr, /*in*/flags IORING_SQE_FLAGS) (HRESULT) = api-ms-win-core-ioring-l1-1-0.BuildIoRingReadFile
//sys	BuildIoRingRegisterFileHandles(/*in*/ioRing *HIORING__, /*in*/count uint32, /*in*/handles *HANDLE, /*in*/userData uintptr) (HRESULT) = api-ms-win-core-ioring-l1-1-0.BuildIoRingRegisterFileHandles
//sys	BuildIoRingRegisterBuffers(/*in*/ioRing *HIORING__, /*in*/count uint32, /*in*/buffers *IORING_BUFFER_INFO, /*in*/userData uintptr) (HRESULT) = api-ms-win-core-ioring-l1-1-0.BuildIoRingRegisterBuffers
//sys	Wow64EnableWow64FsRedirection(/*in*/Wow64FsEnableRedirection BOOLEAN) (BOOLEAN)
//sys	Wow64DisableWow64FsRedirection(/*out*/OldValue *unsafe.Pointer) (BOOL)
//sys	Wow64RevertWow64FsRedirection(/*in*/OlValue unsafe.Pointer) (BOOL)
//sys	GetBinaryTypeA(/*in*/lpApplicationName PSTR, /*out*/lpBinaryType *uint32) (BOOL)
//sys	GetBinaryTypeW(/*in*/lpApplicationName PWSTR, /*out*/lpBinaryType *uint32) (BOOL)
//sys	GetShortPathNameA(/*in*/lpszLongPath PSTR, /*out*//*optional*/lpszShortPath PSTR, /*in*/cchBuffer uint32) (uint32)
//sys	GetLongPathNameTransactedA(/*in*/lpszShortPath PSTR, /*out*//*optional*/lpszLongPath PSTR, /*in*/cchBuffer uint32, /*in*/hTransaction HANDLE) (uint32)
//sys	GetLongPathNameTransactedW(/*in*/lpszShortPath PWSTR, /*out*//*optional*/lpszLongPath PWSTR, /*in*/cchBuffer uint32, /*in*/hTransaction HANDLE) (uint32)
//sys	SetFileCompletionNotificationModes(/*in*/FileHandle HANDLE, /*in*/Flags uint8) (BOOL)
//sys	SetFileShortNameA(/*in*/hFile HANDLE, /*in*/lpShortName PSTR) (BOOL)
//sys	SetFileShortNameW(/*in*/hFile HANDLE, /*in*/lpShortName PWSTR) (BOOL)
//sys	SetTapePosition(/*in*/hDevice HANDLE, /*in*/dwPositionMethod TAPE_POSITION_METHOD, /*in*/dwPartition uint32, /*in*/dwOffsetLow uint32, /*in*/dwOffsetHigh uint32, /*in*/bImmediate BOOL) (uint32)
//sys	GetTapePosition(/*in*/hDevice HANDLE, /*in*/dwPositionType TAPE_POSITION_TYPE, /*out*/lpdwPartition *uint32, /*out*/lpdwOffsetLow *uint32, /*out*/lpdwOffsetHigh *uint32) (uint32)
//sys	PrepareTape(/*in*/hDevice HANDLE, /*in*/dwOperation PREPARE_TAPE_OPERATION, /*in*/bImmediate BOOL) (uint32)
//sys	EraseTape(/*in*/hDevice HANDLE, /*in*/dwEraseType ERASE_TAPE_TYPE, /*in*/bImmediate BOOL) (uint32)
//sys	CreateTapePartition(/*in*/hDevice HANDLE, /*in*/dwPartitionMethod CREATE_TAPE_PARTITION_METHOD, /*in*/dwCount uint32, /*in*/dwSize uint32) (uint32)
//sys	WriteTapemark(/*in*/hDevice HANDLE, /*in*/dwTapemarkType TAPEMARK_TYPE, /*in*/dwTapemarkCount uint32, /*in*/bImmediate BOOL) (uint32)
//sys	GetTapeStatus(/*in*/hDevice HANDLE) (uint32)
//sys	GetTapeParameters(/*in*/hDevice HANDLE, /*in*/dwOperation GET_TAPE_DRIVE_PARAMETERS_OPERATION, /*in*//*out*/lpdwSize *uint32, /*out*/lpTapeInformation unsafe.Pointer) (uint32)
//sys	SetTapeParameters(/*in*/hDevice HANDLE, /*in*/dwOperation TAPE_INFORMATION_TYPE, /*in*/lpTapeInformation unsafe.Pointer) (uint32)
//sys	EncryptFileA(/*in*/lpFileName PSTR) (BOOL) = advapi32.EncryptFileA
//sys	EncryptFileW(/*in*/lpFileName PWSTR) (BOOL) = advapi32.EncryptFileW
//sys	DecryptFileA(/*in*/lpFileName PSTR, /*in*/dwReserved uint32) (BOOL) = advapi32.DecryptFileA
//sys	DecryptFileW(/*in*/lpFileName PWSTR, /*in*/dwReserved uint32) (BOOL) = advapi32.DecryptFileW
//sys	FileEncryptionStatusA(/*in*/lpFileName PSTR, /*out*/lpStatus *uint32) (BOOL) = advapi32.FileEncryptionStatusA
//sys	FileEncryptionStatusW(/*in*/lpFileName PWSTR, /*out*/lpStatus *uint32) (BOOL) = advapi32.FileEncryptionStatusW
//sys	OpenEncryptedFileRawA(/*in*/lpFileName PSTR, /*in*/ulFlags uint32, /*out*/pvContext *unsafe.Pointer) (uint32) = advapi32.OpenEncryptedFileRawA
//sys	OpenEncryptedFileRawW(/*in*/lpFileName PWSTR, /*in*/ulFlags uint32, /*out*/pvContext *unsafe.Pointer) (uint32) = advapi32.OpenEncryptedFileRawW
//sys	ReadEncryptedFileRaw(/*in*/pfExportCallback PFE_EXPORT_FUNC, /*in*//*optional*/pvCallbackContext unsafe.Pointer, /*in*/pvContext unsafe.Pointer) (uint32) = advapi32.ReadEncryptedFileRaw
//sys	WriteEncryptedFileRaw(/*in*/pfImportCallback PFE_IMPORT_FUNC, /*in*//*optional*/pvCallbackContext unsafe.Pointer, /*in*/pvContext unsafe.Pointer) (uint32) = advapi32.WriteEncryptedFileRaw
//sys	CloseEncryptedFileRaw(/*in*/pvContext unsafe.Pointer) = advapi32.CloseEncryptedFileRaw
//sys	OpenFile(/*in*/lpFileName PSTR, /*in*//*out*/lpReOpenBuff *OFSTRUCT, /*in*/uStyle LZOPENFILE_STYLE) (int32)
//sys	BackupRead(/*in*/hFile HANDLE, /*out*/lpBuffer *uint8, /*in*/nNumberOfBytesToRead uint32, /*out*/lpNumberOfBytesRead *uint32, /*in*/bAbort BOOL, /*in*/bProcessSecurity BOOL, /*in*//*out*/lpContext *unsafe.Pointer) (BOOL)
//sys	BackupSeek(/*in*/hFile HANDLE, /*in*/dwLowBytesToSeek uint32, /*in*/dwHighBytesToSeek uint32, /*out*/lpdwLowByteSeeked *uint32, /*out*/lpdwHighByteSeeked *uint32, /*in*//*out*/lpContext *unsafe.Pointer) (BOOL)
//sys	BackupWrite(/*in*/hFile HANDLE, /*in*/lpBuffer *uint8, /*in*/nNumberOfBytesToWrite uint32, /*out*/lpNumberOfBytesWritten *uint32, /*in*/bAbort BOOL, /*in*/bProcessSecurity BOOL, /*in*//*out*/lpContext *unsafe.Pointer) (BOOL)
//sys	GetLogicalDriveStringsA(/*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PSTR) (uint32)
//sys	SetSearchPathMode(/*in*/Flags uint32) (BOOL)
//sys	CreateDirectoryExA(/*in*/lpTemplateDirectory PSTR, /*in*/lpNewDirectory PSTR, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES) (BOOL)
//sys	CreateDirectoryExW(/*in*/lpTemplateDirectory PWSTR, /*in*/lpNewDirectory PWSTR, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES) (BOOL)
//sys	CreateDirectoryTransactedA(/*in*//*optional*/lpTemplateDirectory PSTR, /*in*/lpNewDirectory PSTR, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/hTransaction HANDLE) (BOOL)
//sys	CreateDirectoryTransactedW(/*in*//*optional*/lpTemplateDirectory PWSTR, /*in*/lpNewDirectory PWSTR, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/hTransaction HANDLE) (BOOL)
//sys	RemoveDirectoryTransactedA(/*in*/lpPathName PSTR, /*in*/hTransaction HANDLE) (BOOL)
//sys	RemoveDirectoryTransactedW(/*in*/lpPathName PWSTR, /*in*/hTransaction HANDLE) (BOOL)
//sys	GetFullPathNameTransactedA(/*in*/lpFileName PSTR, /*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PSTR, /*out*//*optional*/lpFilePart *PSTR, /*in*/hTransaction HANDLE) (uint32)
//sys	GetFullPathNameTransactedW(/*in*/lpFileName PWSTR, /*in*/nBufferLength uint32, /*out*//*optional*/lpBuffer PWSTR, /*out*//*optional*/lpFilePart *PWSTR, /*in*/hTransaction HANDLE) (uint32)
//sys	DefineDosDeviceA(/*in*/dwFlags DEFINE_DOS_DEVICE_FLAGS, /*in*/lpDeviceName PSTR, /*in*//*optional*/lpTargetPath PSTR) (BOOL)
//sys	QueryDosDeviceA(/*in*//*optional*/lpDeviceName PSTR, /*out*//*optional*/lpTargetPath PSTR, /*in*/ucchMax uint32) (uint32)
//sys	CreateFileTransactedA(/*in*/lpFileName PSTR, /*in*/dwDesiredAccess uint32, /*in*/dwShareMode FILE_SHARE_MODE, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/dwCreationDisposition FILE_CREATION_DISPOSITION, /*in*/dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES, /*in*//*optional*/hTemplateFile HANDLE, /*in*/hTransaction HANDLE, /*in*//*optional*/pusMiniVersion *TXFS_MINIVERSION, /*in*//*out*/lpExtendedParameter unsafe.Pointer) (HANDLE)
//sys	CreateFileTransactedW(/*in*/lpFileName PWSTR, /*in*/dwDesiredAccess uint32, /*in*/dwShareMode FILE_SHARE_MODE, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/dwCreationDisposition FILE_CREATION_DISPOSITION, /*in*/dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES, /*in*//*optional*/hTemplateFile HANDLE, /*in*/hTransaction HANDLE, /*in*//*optional*/pusMiniVersion *TXFS_MINIVERSION, /*in*//*out*/lpExtendedParameter unsafe.Pointer) (HANDLE)
//sys	ReOpenFile(/*in*/hOriginalFile HANDLE, /*in*/dwDesiredAccess FILE_ACCESS_FLAGS, /*in*/dwShareMode FILE_SHARE_MODE, /*in*/dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES) (HANDLE)
//sys	SetFileAttributesTransactedA(/*in*/lpFileName PSTR, /*in*/dwFileAttributes uint32, /*in*/hTransaction HANDLE) (BOOL)
//sys	SetFileAttributesTransactedW(/*in*/lpFileName PWSTR, /*in*/dwFileAttributes uint32, /*in*/hTransaction HANDLE) (BOOL)
//sys	GetFileAttributesTransactedA(/*in*/lpFileName PSTR, /*in*/fInfoLevelId GET_FILEEX_INFO_LEVELS, /*out*/lpFileInformation unsafe.Pointer, /*in*/hTransaction HANDLE) (BOOL)
//sys	GetFileAttributesTransactedW(/*in*/lpFileName PWSTR, /*in*/fInfoLevelId GET_FILEEX_INFO_LEVELS, /*out*/lpFileInformation unsafe.Pointer, /*in*/hTransaction HANDLE) (BOOL)
//sys	GetCompressedFileSizeTransactedA(/*in*/lpFileName PSTR, /*out*//*optional*/lpFileSizeHigh *uint32, /*in*/hTransaction HANDLE) (uint32)
//sys	GetCompressedFileSizeTransactedW(/*in*/lpFileName PWSTR, /*out*//*optional*/lpFileSizeHigh *uint32, /*in*/hTransaction HANDLE) (uint32)
//sys	DeleteFileTransactedA(/*in*/lpFileName PSTR, /*in*/hTransaction HANDLE) (BOOL)
//sys	DeleteFileTransactedW(/*in*/lpFileName PWSTR, /*in*/hTransaction HANDLE) (BOOL)
//sys	CheckNameLegalDOS8Dot3A(/*in*/lpName PSTR, /*out*//*optional*/lpOemName PSTR, /*in*/OemNameSize uint32, /*out*//*optional*/pbNameContainsSpaces *BOOL, /*out*/pbNameLegal *BOOL) (BOOL)
//sys	CheckNameLegalDOS8Dot3W(/*in*/lpName PWSTR, /*out*//*optional*/lpOemName PSTR, /*in*/OemNameSize uint32, /*out*//*optional*/pbNameContainsSpaces *BOOL, /*out*/pbNameLegal *BOOL) (BOOL)
//sys	FindFirstFileTransactedA(/*in*/lpFileName PSTR, /*in*/fInfoLevelId FINDEX_INFO_LEVELS, /*out*/lpFindFileData unsafe.Pointer, /*in*/fSearchOp FINDEX_SEARCH_OPS, /*in*//*out*/lpSearchFilter unsafe.Pointer, /*in*/dwAdditionalFlags uint32, /*in*/hTransaction HANDLE) (FindFileHandle)
//sys	FindFirstFileTransactedW(/*in*/lpFileName PWSTR, /*in*/fInfoLevelId FINDEX_INFO_LEVELS, /*out*/lpFindFileData unsafe.Pointer, /*in*/fSearchOp FINDEX_SEARCH_OPS, /*in*//*out*/lpSearchFilter unsafe.Pointer, /*in*/dwAdditionalFlags uint32, /*in*/hTransaction HANDLE) (FindFileHandle)
//sys	CopyFileA(/*in*/lpExistingFileName PSTR, /*in*/lpNewFileName PSTR, /*in*/bFailIfExists BOOL) (BOOL)
//sys	CopyFileW(/*in*/lpExistingFileName PWSTR, /*in*/lpNewFileName PWSTR, /*in*/bFailIfExists BOOL) (BOOL)
//sys	CopyFileExA(/*in*/lpExistingFileName PSTR, /*in*/lpNewFileName PSTR, /*in*//*optional*/lpProgressRoutine LPPROGRESS_ROUTINE, /*in*//*optional*/lpData unsafe.Pointer, /*in*//*out*//*optional*/pbCancel *int32, /*in*/dwCopyFlags uint32) (BOOL)
//sys	CopyFileExW(/*in*/lpExistingFileName PWSTR, /*in*/lpNewFileName PWSTR, /*in*//*optional*/lpProgressRoutine LPPROGRESS_ROUTINE, /*in*//*optional*/lpData unsafe.Pointer, /*in*//*out*//*optional*/pbCancel *int32, /*in*/dwCopyFlags uint32) (BOOL)
//sys	CopyFileTransactedA(/*in*/lpExistingFileName PSTR, /*in*/lpNewFileName PSTR, /*in*//*optional*/lpProgressRoutine LPPROGRESS_ROUTINE, /*in*//*optional*/lpData unsafe.Pointer, /*in*//*optional*/pbCancel *int32, /*in*/dwCopyFlags uint32, /*in*/hTransaction HANDLE) (BOOL)
//sys	CopyFileTransactedW(/*in*/lpExistingFileName PWSTR, /*in*/lpNewFileName PWSTR, /*in*//*optional*/lpProgressRoutine LPPROGRESS_ROUTINE, /*in*//*optional*/lpData unsafe.Pointer, /*in*//*optional*/pbCancel *int32, /*in*/dwCopyFlags uint32, /*in*/hTransaction HANDLE) (BOOL)
//sys	CopyFile2(/*in*/pwszExistingFileName PWSTR, /*in*/pwszNewFileName PWSTR, /*in*//*optional*/pExtendedParameters *COPYFILE2_EXTENDED_PARAMETERS) (HRESULT)
//sys	MoveFileA(/*in*/lpExistingFileName PSTR, /*in*/lpNewFileName PSTR) (BOOL)
//sys	MoveFileW(/*in*/lpExistingFileName PWSTR, /*in*/lpNewFileName PWSTR) (BOOL)
//sys	MoveFileExA(/*in*/lpExistingFileName PSTR, /*in*//*optional*/lpNewFileName PSTR, /*in*/dwFlags MOVE_FILE_FLAGS) (BOOL)
//sys	MoveFileExW(/*in*/lpExistingFileName PWSTR, /*in*//*optional*/lpNewFileName PWSTR, /*in*/dwFlags MOVE_FILE_FLAGS) (BOOL)
//sys	MoveFileWithProgressA(/*in*/lpExistingFileName PSTR, /*in*//*optional*/lpNewFileName PSTR, /*in*//*optional*/lpProgressRoutine LPPROGRESS_ROUTINE, /*in*//*optional*/lpData unsafe.Pointer, /*in*/dwFlags MOVE_FILE_FLAGS) (BOOL)
//sys	MoveFileWithProgressW(/*in*/lpExistingFileName PWSTR, /*in*//*optional*/lpNewFileName PWSTR, /*in*//*optional*/lpProgressRoutine LPPROGRESS_ROUTINE, /*in*//*optional*/lpData unsafe.Pointer, /*in*/dwFlags MOVE_FILE_FLAGS) (BOOL)
//sys	MoveFileTransactedA(/*in*/lpExistingFileName PSTR, /*in*//*optional*/lpNewFileName PSTR, /*in*//*optional*/lpProgressRoutine LPPROGRESS_ROUTINE, /*in*//*optional*/lpData unsafe.Pointer, /*in*/dwFlags MOVE_FILE_FLAGS, /*in*/hTransaction HANDLE) (BOOL)
//sys	MoveFileTransactedW(/*in*/lpExistingFileName PWSTR, /*in*//*optional*/lpNewFileName PWSTR, /*in*//*optional*/lpProgressRoutine LPPROGRESS_ROUTINE, /*in*//*optional*/lpData unsafe.Pointer, /*in*/dwFlags MOVE_FILE_FLAGS, /*in*/hTransaction HANDLE) (BOOL)
//sys	ReplaceFileA(/*in*/lpReplacedFileName PSTR, /*in*/lpReplacementFileName PSTR, /*in*//*optional*/lpBackupFileName PSTR, /*in*/dwReplaceFlags REPLACE_FILE_FLAGS, /*in*//*out*/lpExclude unsafe.Pointer, /*in*//*out*/lpReserved unsafe.Pointer) (BOOL)
//sys	ReplaceFileW(/*in*/lpReplacedFileName PWSTR, /*in*/lpReplacementFileName PWSTR, /*in*//*optional*/lpBackupFileName PWSTR, /*in*/dwReplaceFlags REPLACE_FILE_FLAGS, /*in*//*out*/lpExclude unsafe.Pointer, /*in*//*out*/lpReserved unsafe.Pointer) (BOOL)
//sys	CreateHardLinkA(/*in*/lpFileName PSTR, /*in*/lpExistingFileName PSTR, /*in*//*out*/lpSecurityAttributes *SECURITY_ATTRIBUTES) (BOOL)
//sys	CreateHardLinkW(/*in*/lpFileName PWSTR, /*in*/lpExistingFileName PWSTR, /*in*//*out*/lpSecurityAttributes *SECURITY_ATTRIBUTES) (BOOL)
//sys	CreateHardLinkTransactedA(/*in*/lpFileName PSTR, /*in*/lpExistingFileName PSTR, /*in*//*out*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/hTransaction HANDLE) (BOOL)
//sys	CreateHardLinkTransactedW(/*in*/lpFileName PWSTR, /*in*/lpExistingFileName PWSTR, /*in*//*out*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/hTransaction HANDLE) (BOOL)
//sys	FindFirstStreamTransactedW(/*in*/lpFileName PWSTR, /*in*/InfoLevel STREAM_INFO_LEVELS, /*out*/lpFindStreamData unsafe.Pointer, /*in*/dwFlags uint32, /*in*/hTransaction HANDLE) (FindStreamHandle)
//sys	FindFirstFileNameTransactedW(/*in*/lpFileName PWSTR, /*in*/dwFlags uint32, /*in*//*out*/StringLength *uint32, /*out*/LinkName PWSTR, /*in*//*optional*/hTransaction HANDLE) (FindFileNameHandle)
//sys	SetVolumeLabelA(/*in*//*optional*/lpRootPathName PSTR, /*in*//*optional*/lpVolumeName PSTR) (BOOL)
//sys	SetVolumeLabelW(/*in*//*optional*/lpRootPathName PWSTR, /*in*//*optional*/lpVolumeName PWSTR) (BOOL)
//sys	SetFileBandwidthReservation(/*in*/hFile HANDLE, /*in*/nPeriodMilliseconds uint32, /*in*/nBytesPerPeriod uint32, /*in*/bDiscardable BOOL, /*out*/lpTransferSize *uint32, /*out*/lpNumOutstandingRequests *uint32) (BOOL)
//sys	GetFileBandwidthReservation(/*in*/hFile HANDLE, /*out*/lpPeriodMilliseconds *uint32, /*out*/lpBytesPerPeriod *uint32, /*out*/pDiscardable *int32, /*out*/lpTransferSize *uint32, /*out*/lpNumOutstandingRequests *uint32) (BOOL)
//sys	ReadDirectoryChangesW(/*in*/hDirectory HANDLE, /*out*/lpBuffer unsafe.Pointer, /*in*/nBufferLength uint32, /*in*/bWatchSubtree BOOL, /*in*/dwNotifyFilter FILE_NOTIFY_CHANGE, /*out*//*optional*/lpBytesReturned *uint32, /*in*//*out*//*optional*/lpOverlapped *OVERLAPPED, /*in*//*optional*/lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE) (BOOL)
//sys	ReadDirectoryChangesExW(/*in*/hDirectory HANDLE, /*out*/lpBuffer unsafe.Pointer, /*in*/nBufferLength uint32, /*in*/bWatchSubtree BOOL, /*in*/dwNotifyFilter FILE_NOTIFY_CHANGE, /*out*//*optional*/lpBytesReturned *uint32, /*in*//*out*//*optional*/lpOverlapped *OVERLAPPED, /*in*//*optional*/lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE, /*in*/ReadDirectoryNotifyInformationClass READ_DIRECTORY_NOTIFY_INFORMATION_CLASS) (BOOL)
//sys	FindFirstVolumeA(/*out*/lpszVolumeName PSTR, /*in*/cchBufferLength uint32) (FindVolumeHandle)
//sys	FindNextVolumeA(/*in*/hFindVolume FindVolumeHandle, /*out*/lpszVolumeName PSTR, /*in*/cchBufferLength uint32) (BOOL)
//sys	FindFirstVolumeMountPointA(/*in*/lpszRootPathName PSTR, /*out*/lpszVolumeMountPoint PSTR, /*in*/cchBufferLength uint32) (FindVolumeMointPointHandle)
//sys	FindFirstVolumeMountPointW(/*in*/lpszRootPathName PWSTR, /*out*/lpszVolumeMountPoint PWSTR, /*in*/cchBufferLength uint32) (FindVolumeMointPointHandle)
//sys	FindNextVolumeMountPointA(/*in*/hFindVolumeMountPoint FindVolumeMointPointHandle, /*out*/lpszVolumeMountPoint PSTR, /*in*/cchBufferLength uint32) (BOOL)
//sys	FindNextVolumeMountPointW(/*in*/hFindVolumeMountPoint FindVolumeMointPointHandle, /*out*/lpszVolumeMountPoint PWSTR, /*in*/cchBufferLength uint32) (BOOL)
//sys	FindVolumeMountPointClose(/*in*/hFindVolumeMountPoint FindVolumeMointPointHandle) (BOOL)
//sys	SetVolumeMountPointA(/*in*/lpszVolumeMountPoint PSTR, /*in*/lpszVolumeName PSTR) (BOOL)
//sys	SetVolumeMountPointW(/*in*/lpszVolumeMountPoint PWSTR, /*in*/lpszVolumeName PWSTR) (BOOL)
//sys	DeleteVolumeMountPointA(/*in*/lpszVolumeMountPoint PSTR) (BOOL)
//sys	GetVolumeNameForVolumeMountPointA(/*in*/lpszVolumeMountPoint PSTR, /*out*/lpszVolumeName PSTR, /*in*/cchBufferLength uint32) (BOOL)
//sys	GetVolumePathNameA(/*in*/lpszFileName PSTR, /*out*/lpszVolumePathName PSTR, /*in*/cchBufferLength uint32) (BOOL)
//sys	GetVolumePathNamesForVolumeNameA(/*in*/lpszVolumeName PSTR, /*out*//*optional*/lpszVolumePathNames PSTR, /*in*/cchBufferLength uint32, /*out*/lpcchReturnLength *uint32) (BOOL)
//sys	GetFileInformationByHandleEx(/*in*/hFile HANDLE, /*in*/FileInformationClass FILE_INFO_BY_HANDLE_CLASS, /*out*/lpFileInformation unsafe.Pointer, /*in*/dwBufferSize uint32) (BOOL)
//sys	OpenFileById(/*in*/hVolumeHint HANDLE, /*in*/lpFileId *FILE_ID_DESCRIPTOR, /*in*/dwDesiredAccess FILE_ACCESS_FLAGS, /*in*/dwShareMode FILE_SHARE_MODE, /*in*//*optional*/lpSecurityAttributes *SECURITY_ATTRIBUTES, /*in*/dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES) (HANDLE)
//sys	CreateSymbolicLinkA(/*in*/lpSymlinkFileName PSTR, /*in*/lpTargetFileName PSTR, /*in*/dwFlags SYMBOLIC_LINK_FLAGS) (BOOLEAN)
//sys	CreateSymbolicLinkW(/*in*/lpSymlinkFileName PWSTR, /*in*/lpTargetFileName PWSTR, /*in*/dwFlags SYMBOLIC_LINK_FLAGS) (BOOLEAN)
//sys	CreateSymbolicLinkTransactedA(/*in*/lpSymlinkFileName PSTR, /*in*/lpTargetFileName PSTR, /*in*/dwFlags SYMBOLIC_LINK_FLAGS, /*in*/hTransaction HANDLE) (BOOLEAN)
//sys	CreateSymbolicLinkTransactedW(/*in*/lpSymlinkFileName PWSTR, /*in*/lpTargetFileName PWSTR, /*in*/dwFlags SYMBOLIC_LINK_FLAGS, /*in*/hTransaction HANDLE) (BOOLEAN)
//sys	NtCreateFile(/*in*//*out*/FileHandle *HANDLE, /*in*/DesiredAccess uint32, /*in*//*out*/ObjectAttributes *OBJECT_ATTRIBUTES, /*in*//*out*/IoStatusBlock *IO_STATUS_BLOCK, /*in*//*out*/AllocationSize *LARGE_INTEGER, /*in*/FileAttributes uint32, /*in*/ShareAccess FILE_SHARE_MODE, /*in*/CreateDisposition NT_CREATE_FILE_DISPOSITION, /*in*/CreateOptions uint32, /*in*//*out*/EaBuffer unsafe.Pointer, /*in*/EaLength uint32) (NTSTATUS) = ntdll.NtCreateFile

// APIs for Windows.Win32.Security.Cryptography
//sys	CryptAcquireContextA(/*out*/phProv *uintptr, /*in*//*optional*/szContainer PSTR, /*in*//*optional*/szProvider PSTR, /*in*/dwProvType uint32, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptAcquireContextA
//sys	CryptAcquireContextW(/*out*/phProv *uintptr, /*in*//*optional*/szContainer PWSTR, /*in*//*optional*/szProvider PWSTR, /*in*/dwProvType uint32, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptAcquireContextW
//sys	CryptReleaseContext(/*in*/hProv uintptr, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptReleaseContext
//sys	CryptGenKey(/*in*/hProv uintptr, /*in*/Algid uint32, /*in*/dwFlags CRYPT_KEY_FLAGS, /*out*/phKey *uintptr) (BOOL) = advapi32.CryptGenKey
//sys	CryptDeriveKey(/*in*/hProv uintptr, /*in*/Algid uint32, /*in*/hBaseData uintptr, /*in*/dwFlags uint32, /*out*/phKey *uintptr) (BOOL) = advapi32.CryptDeriveKey
//sys	CryptDestroyKey(/*in*/hKey uintptr) (BOOL) = advapi32.CryptDestroyKey
//sys	CryptSetKeyParam(/*in*/hKey uintptr, /*in*/dwParam CRYPT_KEY_PARAM_ID, /*in*/pbData *uint8, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptSetKeyParam
//sys	CryptGetKeyParam(/*in*/hKey uintptr, /*in*/dwParam CRYPT_KEY_PARAM_ID, /*out*//*optional*/pbData *uint8, /*in*//*out*/pdwDataLen *uint32, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptGetKeyParam
//sys	CryptSetHashParam(/*in*/hHash uintptr, /*in*/dwParam CRYPT_SET_HASH_PARAM, /*in*/pbData *uint8, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptSetHashParam
//sys	CryptGetHashParam(/*in*/hHash uintptr, /*in*/dwParam uint32, /*out*//*optional*/pbData *uint8, /*in*//*out*/pdwDataLen *uint32, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptGetHashParam
//sys	CryptSetProvParam(/*in*/hProv uintptr, /*in*/dwParam CRYPT_SET_PROV_PARAM_ID, /*in*/pbData *uint8, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptSetProvParam
//sys	CryptGetProvParam(/*in*/hProv uintptr, /*in*/dwParam uint32, /*out*//*optional*/pbData *uint8, /*in*//*out*/pdwDataLen *uint32, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptGetProvParam
//sys	CryptGenRandom(/*in*/hProv uintptr, /*in*/dwLen uint32, /*in*//*out*/pbBuffer *uint8) (BOOL) = advapi32.CryptGenRandom
//sys	CryptGetUserKey(/*in*/hProv uintptr, /*in*/dwKeySpec uint32, /*out*/phUserKey *uintptr) (BOOL) = advapi32.CryptGetUserKey
//sys	CryptExportKey(/*in*/hKey uintptr, /*in*/hExpKey uintptr, /*in*/dwBlobType uint32, /*in*/dwFlags CRYPT_KEY_FLAGS, /*out*//*optional*/pbData *uint8, /*in*//*out*/pdwDataLen *uint32) (BOOL) = advapi32.CryptExportKey
//sys	CryptImportKey(/*in*/hProv uintptr, /*in*/pbData *uint8, /*in*/dwDataLen uint32, /*in*/hPubKey uintptr, /*in*/dwFlags CRYPT_KEY_FLAGS, /*out*/phKey *uintptr) (BOOL) = advapi32.CryptImportKey
//sys	CryptEncrypt(/*in*/hKey uintptr, /*in*/hHash uintptr, /*in*/Final BOOL, /*in*/dwFlags uint32, /*out*//*optional*/pbData *uint8, /*in*//*out*/pdwDataLen *uint32, /*in*/dwBufLen uint32) (BOOL) = advapi32.CryptEncrypt
//sys	CryptDecrypt(/*in*/hKey uintptr, /*in*/hHash uintptr, /*in*/Final BOOL, /*in*/dwFlags uint32, /*out*/pbData *uint8, /*in*//*out*/pdwDataLen *uint32) (BOOL) = advapi32.CryptDecrypt
//sys	CryptCreateHash(/*in*/hProv uintptr, /*in*/Algid uint32, /*in*/hKey uintptr, /*in*/dwFlags uint32, /*out*/phHash *uintptr) (BOOL) = advapi32.CryptCreateHash
//sys	CryptHashData(/*in*/hHash uintptr, /*in*/pbData *uint8, /*in*/dwDataLen uint32, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptHashData
//sys	CryptHashSessionKey(/*in*/hHash uintptr, /*in*/hKey uintptr, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptHashSessionKey
//sys	CryptDestroyHash(/*in*/hHash uintptr) (BOOL) = advapi32.CryptDestroyHash
//sys	CryptSignHashA(/*in*/hHash uintptr, /*in*/dwKeySpec uint32, /*in*//*optional*/szDescription PSTR, /*in*/dwFlags uint32, /*out*//*optional*/pbSignature *uint8, /*in*//*out*/pdwSigLen *uint32) (BOOL) = advapi32.CryptSignHashA
//sys	CryptSignHashW(/*in*/hHash uintptr, /*in*/dwKeySpec uint32, /*in*//*optional*/szDescription PWSTR, /*in*/dwFlags uint32, /*out*//*optional*/pbSignature *uint8, /*in*//*out*/pdwSigLen *uint32) (BOOL) = advapi32.CryptSignHashW
//sys	CryptVerifySignatureA(/*in*/hHash uintptr, /*in*/pbSignature *uint8, /*in*/dwSigLen uint32, /*in*/hPubKey uintptr, /*in*//*optional*/szDescription PSTR, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptVerifySignatureA
//sys	CryptVerifySignatureW(/*in*/hHash uintptr, /*in*/pbSignature *uint8, /*in*/dwSigLen uint32, /*in*/hPubKey uintptr, /*in*//*optional*/szDescription PWSTR, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptVerifySignatureW
//sys	CryptSetProviderA(/*in*/pszProvName PSTR, /*in*/dwProvType uint32) (BOOL) = advapi32.CryptSetProviderA
//sys	CryptSetProviderW(/*in*/pszProvName PWSTR, /*in*/dwProvType uint32) (BOOL) = advapi32.CryptSetProviderW
//sys	CryptSetProviderExA(/*in*/pszProvName PSTR, /*in*/dwProvType uint32, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptSetProviderExA
//sys	CryptSetProviderExW(/*in*/pszProvName PWSTR, /*in*/dwProvType uint32, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptSetProviderExW
//sys	CryptGetDefaultProviderA(/*in*/dwProvType uint32, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32, /*out*//*optional*/pszProvName PSTR, /*in*//*out*/pcbProvName *uint32) (BOOL) = advapi32.CryptGetDefaultProviderA
//sys	CryptGetDefaultProviderW(/*in*/dwProvType uint32, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32, /*out*//*optional*/pszProvName PWSTR, /*in*//*out*/pcbProvName *uint32) (BOOL) = advapi32.CryptGetDefaultProviderW
//sys	CryptEnumProviderTypesA(/*in*/dwIndex uint32, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32, /*out*/pdwProvType *uint32, /*out*//*optional*/szTypeName PSTR, /*in*//*out*/pcbTypeName *uint32) (BOOL) = advapi32.CryptEnumProviderTypesA
//sys	CryptEnumProviderTypesW(/*in*/dwIndex uint32, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32, /*out*/pdwProvType *uint32, /*out*//*optional*/szTypeName PWSTR, /*in*//*out*/pcbTypeName *uint32) (BOOL) = advapi32.CryptEnumProviderTypesW
//sys	CryptEnumProvidersA(/*in*/dwIndex uint32, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32, /*out*/pdwProvType *uint32, /*out*//*optional*/szProvName PSTR, /*in*//*out*/pcbProvName *uint32) (BOOL) = advapi32.CryptEnumProvidersA
//sys	CryptEnumProvidersW(/*in*/dwIndex uint32, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32, /*out*/pdwProvType *uint32, /*out*//*optional*/szProvName PWSTR, /*in*//*out*/pcbProvName *uint32) (BOOL) = advapi32.CryptEnumProvidersW
//sys	CryptContextAddRef(/*in*/hProv uintptr, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32) (BOOL) = advapi32.CryptContextAddRef
//sys	CryptDuplicateKey(/*in*/hKey uintptr, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32, /*out*/phKey *uintptr) (BOOL) = advapi32.CryptDuplicateKey
//sys	CryptDuplicateHash(/*in*/hHash uintptr, /*in*//*out*/pdwReserved *uint32, /*in*/dwFlags uint32, /*out*/phHash *uintptr) (BOOL) = advapi32.CryptDuplicateHash
//sys	BCryptOpenAlgorithmProvider(/*out*/phAlgorithm *BCRYPT_ALG_HANDLE, /*in*/pszAlgId PWSTR, /*in*//*optional*/pszImplementation PWSTR, /*in*/dwFlags BCRYPT_OPEN_ALGORITHM_PROVIDER_FLAGS) (NTSTATUS) = bcrypt.BCryptOpenAlgorithmProvider
//sys	BCryptEnumAlgorithms(/*in*/dwAlgOperations BCRYPT_OPERATION, /*out*/pAlgCount *uint32, /*out*/ppAlgList **BCRYPT_ALGORITHM_IDENTIFIER, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptEnumAlgorithms
//sys	BCryptEnumProviders(/*in*/pszAlgId PWSTR, /*out*/pImplCount *uint32, /*out*/ppImplList **BCRYPT_PROVIDER_NAME, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptEnumProviders
//sys	BCryptGetProperty(/*in*/hObject unsafe.Pointer, /*in*/pszProperty PWSTR, /*out*//*optional*/pbOutput *uint8, /*in*/cbOutput uint32, /*out*/pcbResult *uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptGetProperty
//sys	BCryptSetProperty(/*in*//*out*/hObject unsafe.Pointer, /*in*/pszProperty PWSTR, /*in*/pbInput *uint8, /*in*/cbInput uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptSetProperty
//sys	BCryptCloseAlgorithmProvider(/*in*/hAlgorithm BCRYPT_ALG_HANDLE, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptCloseAlgorithmProvider
//sys	BCryptFreeBuffer(/*in*/pvBuffer unsafe.Pointer) = bcrypt.BCryptFreeBuffer
//sys	BCryptGenerateSymmetricKey(/*in*/hAlgorithm BCRYPT_ALG_HANDLE, /*out*/phKey *BCRYPT_KEY_HANDLE, /*out*//*optional*/pbKeyObject *uint8, /*in*/cbKeyObject uint32, /*in*/pbSecret *uint8, /*in*/cbSecret uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptGenerateSymmetricKey
//sys	BCryptGenerateKeyPair(/*in*/hAlgorithm BCRYPT_ALG_HANDLE, /*out*/phKey *BCRYPT_KEY_HANDLE, /*in*/dwLength uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptGenerateKeyPair
//sys	BCryptEncrypt(/*in*/hKey BCRYPT_KEY_HANDLE, /*in*//*optional*/pbInput *uint8, /*in*/cbInput uint32, /*in*//*optional*/pPaddingInfo unsafe.Pointer, /*in*//*out*//*optional*/pbIV *uint8, /*in*/cbIV uint32, /*out*//*optional*/pbOutput *uint8, /*in*/cbOutput uint32, /*out*/pcbResult *uint32, /*in*/dwFlags NCRYPT_FLAGS) (NTSTATUS) = bcrypt.BCryptEncrypt
//sys	BCryptDecrypt(/*in*/hKey BCRYPT_KEY_HANDLE, /*in*//*optional*/pbInput *uint8, /*in*/cbInput uint32, /*in*//*optional*/pPaddingInfo unsafe.Pointer, /*in*//*out*//*optional*/pbIV *uint8, /*in*/cbIV uint32, /*out*//*optional*/pbOutput *uint8, /*in*/cbOutput uint32, /*out*/pcbResult *uint32, /*in*/dwFlags NCRYPT_FLAGS) (NTSTATUS) = bcrypt.BCryptDecrypt
//sys	BCryptExportKey(/*in*/hKey BCRYPT_KEY_HANDLE, /*in*//*optional*/hExportKey BCRYPT_KEY_HANDLE, /*in*/pszBlobType PWSTR, /*out*//*optional*/pbOutput *uint8, /*in*/cbOutput uint32, /*out*/pcbResult *uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptExportKey
//sys	BCryptImportKey(/*in*/hAlgorithm BCRYPT_ALG_HANDLE, /*in*//*optional*/hImportKey BCRYPT_KEY_HANDLE, /*in*/pszBlobType PWSTR, /*out*/phKey *BCRYPT_KEY_HANDLE, /*out*//*optional*/pbKeyObject *uint8, /*in*/cbKeyObject uint32, /*in*/pbInput *uint8, /*in*/cbInput uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptImportKey
//sys	BCryptImportKeyPair(/*in*/hAlgorithm BCRYPT_ALG_HANDLE, /*in*//*optional*/hImportKey BCRYPT_KEY_HANDLE, /*in*/pszBlobType PWSTR, /*out*/phKey *BCRYPT_KEY_HANDLE, /*in*/pbInput *uint8, /*in*/cbInput uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptImportKeyPair
//sys	BCryptDuplicateKey(/*in*/hKey BCRYPT_KEY_HANDLE, /*out*/phNewKey *BCRYPT_KEY_HANDLE, /*out*//*optional*/pbKeyObject *uint8, /*in*/cbKeyObject uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptDuplicateKey
//sys	BCryptFinalizeKeyPair(/*in*/hKey BCRYPT_KEY_HANDLE, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptFinalizeKeyPair
//sys	BCryptDestroyKey(/*in*/hKey BCRYPT_KEY_HANDLE) (NTSTATUS) = bcrypt.BCryptDestroyKey
//sys	BCryptDestroySecret(/*in*//*out*/hSecret unsafe.Pointer) (NTSTATUS) = bcrypt.BCryptDestroySecret
//sys	BCryptSignHash(/*in*/hKey BCRYPT_KEY_HANDLE, /*in*//*optional*/pPaddingInfo unsafe.Pointer, /*in*/pbInput *uint8, /*in*/cbInput uint32, /*out*//*optional*/pbOutput *uint8, /*in*/cbOutput uint32, /*out*/pcbResult *uint32, /*in*/dwFlags NCRYPT_FLAGS) (NTSTATUS) = bcrypt.BCryptSignHash
//sys	BCryptVerifySignature(/*in*/hKey BCRYPT_KEY_HANDLE, /*in*//*optional*/pPaddingInfo unsafe.Pointer, /*in*/pbHash *uint8, /*in*/cbHash uint32, /*in*/pbSignature *uint8, /*in*/cbSignature uint32, /*in*/dwFlags NCRYPT_FLAGS) (NTSTATUS) = bcrypt.BCryptVerifySignature
//sys	BCryptSecretAgreement(/*in*/hPrivKey BCRYPT_KEY_HANDLE, /*in*/hPubKey BCRYPT_KEY_HANDLE, /*out*/phAgreedSecret *unsafe.Pointer, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptSecretAgreement
//sys	BCryptDeriveKey(/*in*/hSharedSecret unsafe.Pointer, /*in*/pwszKDF PWSTR, /*in*//*optional*/pParameterList *BCryptBufferDesc, /*out*//*optional*/pbDerivedKey *uint8, /*in*/cbDerivedKey uint32, /*out*/pcbResult *uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptDeriveKey
//sys	BCryptKeyDerivation(/*in*/hKey BCRYPT_KEY_HANDLE, /*in*//*optional*/pParameterList *BCryptBufferDesc, /*out*/pbDerivedKey *uint8, /*in*/cbDerivedKey uint32, /*out*/pcbResult *uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptKeyDerivation
//sys	BCryptCreateHash(/*in*/hAlgorithm BCRYPT_ALG_HANDLE, /*out*/phHash *unsafe.Pointer, /*out*//*optional*/pbHashObject *uint8, /*in*/cbHashObject uint32, /*in*//*optional*/pbSecret *uint8, /*in*/cbSecret uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptCreateHash
//sys	BCryptHashData(/*in*//*out*/hHash unsafe.Pointer, /*in*/pbInput *uint8, /*in*/cbInput uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptHashData
//sys	BCryptFinishHash(/*in*//*out*/hHash unsafe.Pointer, /*out*/pbOutput *uint8, /*in*/cbOutput uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptFinishHash
//sys	BCryptCreateMultiHash(/*in*/hAlgorithm BCRYPT_ALG_HANDLE, /*out*/phHash *unsafe.Pointer, /*in*/nHashes uint32, /*out*//*optional*/pbHashObject *uint8, /*in*/cbHashObject uint32, /*in*//*optional*/pbSecret *uint8, /*in*/cbSecret uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptCreateMultiHash
//sys	BCryptProcessMultiOperations(/*in*//*out*/hObject unsafe.Pointer, /*in*/operationType BCRYPT_MULTI_OPERATION_TYPE, /*in*/pOperations unsafe.Pointer, /*in*/cbOperations uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptProcessMultiOperations
//sys	BCryptDuplicateHash(/*in*/hHash unsafe.Pointer, /*out*/phNewHash *unsafe.Pointer, /*out*//*optional*/pbHashObject *uint8, /*in*/cbHashObject uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptDuplicateHash
//sys	BCryptDestroyHash(/*in*//*out*/hHash unsafe.Pointer) (NTSTATUS) = bcrypt.BCryptDestroyHash
//sys	BCryptHash(/*in*/hAlgorithm BCRYPT_ALG_HANDLE, /*in*//*optional*/pbSecret *uint8, /*in*/cbSecret uint32, /*in*/pbInput *uint8, /*in*/cbInput uint32, /*out*/pbOutput *uint8, /*in*/cbOutput uint32) (NTSTATUS) = bcrypt.BCryptHash
//sys	BCryptGenRandom(/*in*//*optional*/hAlgorithm BCRYPT_ALG_HANDLE, /*out*/pbBuffer *uint8, /*in*/cbBuffer uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptGenRandom
//sys	BCryptDeriveKeyCapi(/*in*/hHash unsafe.Pointer, /*in*//*optional*/hTargetAlg BCRYPT_ALG_HANDLE, /*out*/pbDerivedKey *uint8, /*in*/cbDerivedKey uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptDeriveKeyCapi
//sys	BCryptDeriveKeyPBKDF2(/*in*/hPrf BCRYPT_ALG_HANDLE, /*in*//*optional*/pbPassword *uint8, /*in*/cbPassword uint32, /*in*//*optional*/pbSalt *uint8, /*in*/cbSalt uint32, /*in*/cIterations uint64, /*out*/pbDerivedKey *uint8, /*in*/cbDerivedKey uint32, /*in*/dwFlags uint32) (NTSTATUS) = bcrypt.BCryptDeriveKeyPBKDF2
//sys	BCryptQueryProviderRegistration(/*in*/pszProvider PWSTR, /*in*/dwMode BCRYPT_QUERY_PROVIDER_MODE, /*in*/dwInterface BCRYPT_INTERFACE, /*in*//*out*/pcbBuffer *uint32, /*in*//*out*/ppBuffer **CRYPT_PROVIDER_REG) (NTSTATUS) = bcrypt.BCryptQueryProviderRegistration
//sys	BCryptEnumRegisteredProviders(/*in*//*out*/pcbBuffer *uint32, /*in*//*out*/ppBuffer **CRYPT_PROVIDERS) (NTSTATUS) = bcrypt.BCryptEnumRegisteredProviders
//sys	BCryptCreateContext(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*//*optional*/pConfig *CRYPT_CONTEXT_CONFIG) (NTSTATUS) = bcrypt.BCryptCreateContext
//sys	BCryptDeleteContext(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR) (NTSTATUS) = bcrypt.BCryptDeleteContext
//sys	BCryptEnumContexts(/*in*/dwTable BCRYPT_TABLE, /*in*//*out*/pcbBuffer *uint32, /*in*//*out*/ppBuffer **CRYPT_CONTEXTS) (NTSTATUS) = bcrypt.BCryptEnumContexts
//sys	BCryptConfigureContext(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*/pConfig *CRYPT_CONTEXT_CONFIG) (NTSTATUS) = bcrypt.BCryptConfigureContext
//sys	BCryptQueryContextConfiguration(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*//*out*/pcbBuffer *uint32, /*in*//*out*/ppBuffer **CRYPT_CONTEXT_CONFIG) (NTSTATUS) = bcrypt.BCryptQueryContextConfiguration
//sys	BCryptAddContextFunction(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*/dwInterface BCRYPT_INTERFACE, /*in*/pszFunction PWSTR, /*in*/dwPosition uint32) (NTSTATUS) = bcrypt.BCryptAddContextFunction
//sys	BCryptRemoveContextFunction(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*/dwInterface BCRYPT_INTERFACE, /*in*/pszFunction PWSTR) (NTSTATUS) = bcrypt.BCryptRemoveContextFunction
//sys	BCryptEnumContextFunctions(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*/dwInterface BCRYPT_INTERFACE, /*in*//*out*/pcbBuffer *uint32, /*in*//*out*/ppBuffer **CRYPT_CONTEXT_FUNCTIONS) (NTSTATUS) = bcrypt.BCryptEnumContextFunctions
//sys	BCryptConfigureContextFunction(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*/dwInterface BCRYPT_INTERFACE, /*in*/pszFunction PWSTR, /*in*/pConfig *CRYPT_CONTEXT_FUNCTION_CONFIG) (NTSTATUS) = bcrypt.BCryptConfigureContextFunction
//sys	BCryptQueryContextFunctionConfiguration(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*/dwInterface BCRYPT_INTERFACE, /*in*/pszFunction PWSTR, /*in*//*out*/pcbBuffer *uint32, /*in*//*out*/ppBuffer **CRYPT_CONTEXT_FUNCTION_CONFIG) (NTSTATUS) = bcrypt.BCryptQueryContextFunctionConfiguration
//sys	BCryptEnumContextFunctionProviders(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*/dwInterface BCRYPT_INTERFACE, /*in*/pszFunction PWSTR, /*in*//*out*/pcbBuffer *uint32, /*in*//*out*/ppBuffer **CRYPT_CONTEXT_FUNCTION_PROVIDERS) (NTSTATUS) = bcrypt.BCryptEnumContextFunctionProviders
//sys	BCryptSetContextFunctionProperty(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*/dwInterface BCRYPT_INTERFACE, /*in*/pszFunction PWSTR, /*in*/pszProperty PWSTR, /*in*/cbValue uint32, /*in*//*optional*/pbValue *uint8) (NTSTATUS) = bcrypt.BCryptSetContextFunctionProperty
//sys	BCryptQueryContextFunctionProperty(/*in*/dwTable BCRYPT_TABLE, /*in*/pszContext PWSTR, /*in*/dwInterface BCRYPT_INTERFACE, /*in*/pszFunction PWSTR, /*in*/pszProperty PWSTR, /*in*//*out*/pcbValue *uint32, /*in*//*out*/ppbValue **uint8) (NTSTATUS) = bcrypt.BCryptQueryContextFunctionProperty
//sys	BCryptRegisterConfigChangeNotify(/*out*/phEvent *HANDLE) (NTSTATUS) = bcrypt.BCryptRegisterConfigChangeNotify
//sys	BCryptUnregisterConfigChangeNotify(/*in*/hEvent HANDLE) (NTSTATUS) = bcrypt.BCryptUnregisterConfigChangeNotify
//sys	BCryptResolveProviders(/*in*//*optional*/pszContext PWSTR, /*in*//*optional*/dwInterface uint32, /*in*//*optional*/pszFunction PWSTR, /*in*//*optional*/pszProvider PWSTR, /*in*/dwMode BCRYPT_QUERY_PROVIDER_MODE, /*in*/dwFlags BCRYPT_RESOLVE_PROVIDERS_FLAGS, /*in*//*out*/pcbBuffer *uint32, /*in*//*out*/ppBuffer **CRYPT_PROVIDER_REFS) (NTSTATUS) = bcrypt.BCryptResolveProviders
//sys	BCryptGetFipsAlgorithmMode(/*out*/pfEnabled *uint8) (NTSTATUS) = bcrypt.BCryptGetFipsAlgorithmMode
//sys	NCryptOpenStorageProvider(/*out*/phProvider *NCRYPT_PROV_HANDLE, /*in*//*optional*/pszProviderName PWSTR, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptOpenStorageProvider
//sys	NCryptEnumAlgorithms(/*in*/hProvider NCRYPT_PROV_HANDLE, /*in*/dwAlgOperations NCRYPT_OPERATION, /*out*/pdwAlgCount *uint32, /*out*/ppAlgList **NCryptAlgorithmName, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptEnumAlgorithms
//sys	NCryptIsAlgSupported(/*in*/hProvider NCRYPT_PROV_HANDLE, /*in*/pszAlgId PWSTR, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptIsAlgSupported
//sys	NCryptEnumKeys(/*in*/hProvider NCRYPT_PROV_HANDLE, /*in*//*optional*/pszScope PWSTR, /*out*/ppKeyName **NCryptKeyName, /*in*//*out*/ppEnumState *unsafe.Pointer, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptEnumKeys
//sys	NCryptEnumStorageProviders(/*out*/pdwProviderCount *uint32, /*out*/ppProviderList **NCryptProviderName, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptEnumStorageProviders
//sys	NCryptFreeBuffer(/*in*//*out*/pvInput unsafe.Pointer) (HRESULT) = ncrypt.NCryptFreeBuffer
//sys	NCryptOpenKey(/*in*/hProvider NCRYPT_PROV_HANDLE, /*out*/phKey *NCRYPT_KEY_HANDLE, /*in*/pszKeyName PWSTR, /*in*//*optional*/dwLegacyKeySpec CERT_KEY_SPEC, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptOpenKey
//sys	NCryptCreatePersistedKey(/*in*/hProvider NCRYPT_PROV_HANDLE, /*out*/phKey *NCRYPT_KEY_HANDLE, /*in*/pszAlgId PWSTR, /*in*//*optional*/pszKeyName PWSTR, /*in*/dwLegacyKeySpec CERT_KEY_SPEC, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptCreatePersistedKey
//sys	NCryptGetProperty(/*in*/hObject NCRYPT_HANDLE, /*in*/pszProperty PWSTR, /*out*//*optional*/pbOutput *uint8, /*in*/cbOutput uint32, /*out*/pcbResult *uint32, /*in*/dwFlags OBJECT_SECURITY_INFORMATION) (HRESULT) = ncrypt.NCryptGetProperty
//sys	NCryptSetProperty(/*in*/hObject NCRYPT_HANDLE, /*in*/pszProperty PWSTR, /*in*/pbInput *uint8, /*in*/cbInput uint32, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptSetProperty
//sys	NCryptFinalizeKey(/*in*/hKey NCRYPT_KEY_HANDLE, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptFinalizeKey
//sys	NCryptEncrypt(/*in*/hKey NCRYPT_KEY_HANDLE, /*in*//*optional*/pbInput *uint8, /*in*/cbInput uint32, /*in*//*optional*/pPaddingInfo unsafe.Pointer, /*out*//*optional*/pbOutput *uint8, /*in*/cbOutput uint32, /*out*/pcbResult *uint32, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptEncrypt
//sys	NCryptDecrypt(/*in*/hKey NCRYPT_KEY_HANDLE, /*in*//*optional*/pbInput *uint8, /*in*/cbInput uint32, /*in*//*optional*/pPaddingInfo unsafe.Pointer, /*out*//*optional*/pbOutput *uint8, /*in*/cbOutput uint32, /*out*/pcbResult *uint32, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptDecrypt
//sys	NCryptImportKey(/*in*/hProvider NCRYPT_PROV_HANDLE, /*in*//*optional*/hImportKey NCRYPT_KEY_HANDLE, /*in*/pszBlobType PWSTR, /*in*//*optional*/pParameterList *BCryptBufferDesc, /*out*/phKey *NCRYPT_KEY_HANDLE, /*in*/pbData *uint8, /*in*/cbData uint32, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptImportKey
//sys	NCryptExportKey(/*in*/hKey NCRYPT_KEY_HANDLE, /*in*//*optional*/hExportKey NCRYPT_KEY_HANDLE, /*in*/pszBlobType PWSTR, /*in*//*optional*/pParameterList *BCryptBufferDesc, /*out*//*optional*/pbOutput *uint8, /*in*/cbOutput uint32, /*out*/pcbResult *uint32, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptExportKey
//sys	NCryptSignHash(/*in*/hKey NCRYPT_KEY_HANDLE, /*in*//*optional*/pPaddingInfo unsafe.Pointer, /*in*/pbHashValue *uint8, /*in*/cbHashValue uint32, /*out*//*optional*/pbSignature *uint8, /*in*/cbSignature uint32, /*out*/pcbResult *uint32, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptSignHash
//sys	NCryptVerifySignature(/*in*/hKey NCRYPT_KEY_HANDLE, /*in*//*optional*/pPaddingInfo unsafe.Pointer, /*in*/pbHashValue *uint8, /*in*/cbHashValue uint32, /*in*/pbSignature *uint8, /*in*/cbSignature uint32, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptVerifySignature
//sys	NCryptDeleteKey(/*in*/hKey NCRYPT_KEY_HANDLE, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptDeleteKey
//sys	NCryptFreeObject(/*in*/hObject NCRYPT_HANDLE) (HRESULT) = ncrypt.NCryptFreeObject
//sys	NCryptIsKeyHandle(/*in*/hKey NCRYPT_KEY_HANDLE) (BOOL) = ncrypt.NCryptIsKeyHandle
//sys	NCryptTranslateHandle(/*out*//*optional*/phProvider *NCRYPT_PROV_HANDLE, /*out*/phKey *NCRYPT_KEY_HANDLE, /*in*/hLegacyProv uintptr, /*in*//*optional*/hLegacyKey uintptr, /*in*//*optional*/dwLegacyKeySpec CERT_KEY_SPEC, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptTranslateHandle
//sys	NCryptNotifyChangeKey(/*in*/hProvider NCRYPT_PROV_HANDLE, /*in*//*out*/phEvent *HANDLE, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptNotifyChangeKey
//sys	NCryptSecretAgreement(/*in*/hPrivKey NCRYPT_KEY_HANDLE, /*in*/hPubKey NCRYPT_KEY_HANDLE, /*out*/phAgreedSecret *NCRYPT_SECRET_HANDLE, /*in*/dwFlags NCRYPT_FLAGS) (HRESULT) = ncrypt.NCryptSecretAgreement
//sys	NCryptDeriveKey(/*in*/hSharedSecret NCRYPT_SECRET_HANDLE, /*in*/pwszKDF PWSTR, /*in*//*optional*/pParameterList *BCryptBufferDesc, /*out*//*optional*/pbDerivedKey *uint8, /*in*/cbDerivedKey uint32, /*out*/pcbResult *uint32, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptDeriveKey
//sys	NCryptKeyDerivation(/*in*/hKey NCRYPT_KEY_HANDLE, /*in*//*optional*/pParameterList *BCryptBufferDesc, /*out*/pbDerivedKey *uint8, /*in*/cbDerivedKey uint32, /*out*/pcbResult *uint32, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptKeyDerivation
//sys	NCryptCreateClaim(/*in*//*optional*/hSubjectKey NCRYPT_KEY_HANDLE, /*in*//*optional*/hAuthorityKey NCRYPT_KEY_HANDLE, /*in*/dwClaimType uint32, /*in*//*optional*/pParameterList *BCryptBufferDesc, /*out*//*optional*/pbClaimBlob *uint8, /*in*/cbClaimBlob uint32, /*out*/pcbResult *uint32, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptCreateClaim
//sys	NCryptVerifyClaim(/*in*/hSubjectKey NCRYPT_KEY_HANDLE, /*in*//*optional*/hAuthorityKey NCRYPT_KEY_HANDLE, /*in*/dwClaimType uint32, /*in*//*optional*/pParameterList *BCryptBufferDesc, /*in*/pbClaimBlob *uint8, /*in*/cbClaimBlob uint32, /*out*/pOutput *BCryptBufferDesc, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptVerifyClaim
//sys	CryptFormatObject(/*in*/dwCertEncodingType uint32, /*in*/dwFormatType uint32, /*in*/dwFormatStrType uint32, /*in*//*optional*/pFormatStruct unsafe.Pointer, /*in*//*optional*/lpszStructType PSTR, /*in*/pbEncoded *uint8, /*in*/cbEncoded uint32, /*out*//*optional*/pbFormat unsafe.Pointer, /*in*//*out*/pcbFormat *uint32) (BOOL) = crypt32.CryptFormatObject
//sys	CryptEncodeObjectEx(/*in*/dwCertEncodingType CERT_QUERY_ENCODING_TYPE, /*in*/lpszStructType PSTR, /*in*/pvStructInfo unsafe.Pointer, /*in*/dwFlags CRYPT_ENCODE_OBJECT_FLAGS, /*in*//*optional*/pEncodePara *CRYPT_ENCODE_PARA, /*out*//*optional*/pvEncoded unsafe.Pointer, /*in*//*out*/pcbEncoded *uint32) (BOOL) = crypt32.CryptEncodeObjectEx
//sys	CryptEncodeObject(/*in*/dwCertEncodingType uint32, /*in*/lpszStructType PSTR, /*in*/pvStructInfo unsafe.Pointer, /*out*//*optional*/pbEncoded *uint8, /*in*//*out*/pcbEncoded *uint32) (BOOL) = crypt32.CryptEncodeObject
//sys	CryptDecodeObjectEx(/*in*/dwCertEncodingType uint32, /*in*/lpszStructType PSTR, /*in*/pbEncoded *uint8, /*in*/cbEncoded uint32, /*in*/dwFlags uint32, /*in*//*optional*/pDecodePara *CRYPT_DECODE_PARA, /*out*//*optional*/pvStructInfo unsafe.Pointer, /*in*//*out*/pcbStructInfo *uint32) (BOOL) = crypt32.CryptDecodeObjectEx
//sys	CryptDecodeObject(/*in*/dwCertEncodingType uint32, /*in*/lpszStructType PSTR, /*in*/pbEncoded *uint8, /*in*/cbEncoded uint32, /*in*/dwFlags uint32, /*out*//*optional*/pvStructInfo unsafe.Pointer, /*in*//*out*/pcbStructInfo *uint32) (BOOL) = crypt32.CryptDecodeObject
//sys	CryptInstallOIDFunctionAddress(/*in*//*optional*/hModule HINSTANCE, /*in*/dwEncodingType uint32, /*in*/pszFuncName PSTR, /*in*/cFuncEntry uint32, /*in*/rgFuncEntry *CRYPT_OID_FUNC_ENTRY, /*in*/dwFlags uint32) (BOOL) = crypt32.CryptInstallOIDFunctionAddress
//sys	CryptInitOIDFunctionSet(/*in*/pszFuncName PSTR, /*in*/dwFlags uint32) (unsafe.Pointer) = crypt32.CryptInitOIDFunctionSet
//sys	CryptGetOIDFunctionAddress(/*in*/hFuncSet unsafe.Pointer, /*in*/dwEncodingType uint32, /*in*/pszOID PSTR, /*in*/dwFlags uint32, /*out*/ppvFuncAddr *unsafe.Pointer, /*out*/phFuncAddr *unsafe.Pointer) (BOOL) = crypt32.CryptGetOIDFunctionAddress
//sys	CryptGetDefaultOIDDllList(/*in*/hFuncSet unsafe.Pointer, /*in*/dwEncodingType uint32, /*out*//*optional*/pwszDllList PWSTR, /*in*//*out*/pcchDllList *uint32) (BOOL) = crypt32.CryptGetDefaultOIDDllList
//sys	CryptGetDefaultOIDFunctionAddress(/*in*/hFuncSet unsafe.Pointer, /*in*/dwEncodingType uint32, /*in*//*optional*/pwszDll PWSTR, /*in*/dwFlags uint32, /*out*/ppvFuncAddr *unsafe.Pointer, /*in*//*out*/phFuncAddr *unsafe.Pointer) (BOOL) = crypt32.CryptGetDefaultOIDFunctionAddress
//sys	CryptFreeOIDFunctionAddress(/*in*/hFuncAddr unsafe.Pointer, /*in*/dwFlags uint32) (BOOL) = crypt32.CryptFreeOIDFunctionAddress
//sys	CryptRegisterOIDFunction(/*in*/dwEncodingType uint32, /*in*/pszFuncName PSTR, /*in*/pszOID PSTR, /*in*//*optional*/pwszDll PWSTR, /*in*//*optional*/pszOverrideFuncName PSTR) (BOOL) = crypt32.CryptRegisterOIDFunction
//sys	CryptUnregisterOIDFunction(/*in*/dwEncodingType uint32, /*in*/pszFuncName PSTR, /*in*/pszOID PSTR) (BOOL) = crypt32.CryptUnregisterOIDFunction
//sys	CryptRegisterDefaultOIDFunction(/*in*/dwEncodingType uint32, /*in*/pszFuncName PSTR, /*in*/dwIndex uint32, /*in*/pwszDll PWSTR) (BOOL) = crypt32.CryptRegisterDefaultOIDFunction
//sys	CryptUnregisterDefaultOIDFunction(/*in*/dwEncodingType uint32, /*in*/pszFuncName PSTR, /*in*/pwszDll PWSTR) (BOOL) = crypt32.CryptUnregisterDefaultOIDFunction
//sys	CryptSetOIDFunctionValue(/*in*/dwEncodingType uint32, /*in*/pszFuncName PSTR, /*in*/pszOID PSTR, /*in*//*optional*/pwszValueName PWSTR, /*in*/dwValueType REG_VALUE_TYPE, /*in*//*optional*/pbValueData *uint8, /*in*/cbValueData uint32) (BOOL) = crypt32.CryptSetOIDFunctionValue
//sys	CryptGetOIDFunctionValue(/*in*/dwEncodingType uint32, /*in*/pszFuncName PSTR, /*in*/pszOID PSTR, /*in*//*optional*/pwszValueName PWSTR, /*out*//*optional*/pdwValueType *uint32, /*out*//*optional*/pbValueData *uint8, /*in*//*out*//*optional*/pcbValueData *uint32) (BOOL) = crypt32.CryptGetOIDFunctionValue
//sys	CryptEnumOIDFunction(/*in*/dwEncodingType uint32, /*in*//*optional*/pszFuncName PSTR, /*in*//*optional*/pszOID PSTR, /*in*/dwFlags uint32, /*in*//*out*//*optional*/pvArg unsafe.Pointer, /*in*/pfnEnumOIDFunc PFN_CRYPT_ENUM_OID_FUNC) (BOOL) = crypt32.CryptEnumOIDFunction
//sys	CryptFindOIDInfo(/*in*/dwKeyType uint32, /*in*/pvKey unsafe.Pointer, /*in*/dwGroupId uint32) (*CRYPT_OID_INFO) = crypt32.CryptFindOIDInfo
//sys	CryptRegisterOIDInfo(/*in*/pInfo *CRYPT_OID_INFO, /*in*/dwFlags uint32) (BOOL) = crypt32.CryptRegisterOIDInfo
//sys	CryptUnregisterOIDInfo(/*in*/pInfo *CRYPT_OID_INFO) (BOOL) = crypt32.CryptUnregisterOIDInfo
//sys	CryptEnumOIDInfo(/*in*/dwGroupId uint32, /*in*/dwFlags uint32, /*in*//*out*//*optional*/pvArg unsafe.Pointer, /*in*/pfnEnumOIDInfo PFN_CRYPT_ENUM_OID_INFO) (BOOL) = crypt32.CryptEnumOIDInfo
//sys	CryptFindLocalizedName(/*in*/pwszCryptName PWSTR) (PWSTR) = crypt32.CryptFindLocalizedName
//sys	CryptMsgOpenToEncode(/*in*/dwMsgEncodingType uint32, /*in*/dwFlags uint32, /*in*/dwMsgType CRYPT_MSG_TYPE, /*in*/pvMsgEncodeInfo unsafe.Pointer, /*in*//*optional*/pszInnerContentObjID PSTR, /*in*//*optional*/pStreamInfo *CMSG_STREAM_INFO) (unsafe.Pointer) = crypt32.CryptMsgOpenToEncode
//sys	CryptMsgCalculateEncodedLength(/*in*/dwMsgEncodingType uint32, /*in*/dwFlags uint32, /*in*/dwMsgType uint32, /*in*/pvMsgEncodeInfo unsafe.Pointer, /*in*//*optional*/pszInnerContentObjID PSTR, /*in*/cbData uint32) (uint32) = crypt32.CryptMsgCalculateEncodedLength
//sys	CryptMsgOpenToDecode(/*in*/dwMsgEncodingType uint32, /*in*/dwFlags uint32, /*in*/dwMsgType uint32, /*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*//*out*/pRecipientInfo *CERT_INFO, /*in*//*optional*/pStreamInfo *CMSG_STREAM_INFO) (unsafe.Pointer) = crypt32.CryptMsgOpenToDecode
//sys	CryptMsgDuplicate(/*in*//*optional*/hCryptMsg unsafe.Pointer) (unsafe.Pointer) = crypt32.CryptMsgDuplicate
//sys	CryptMsgClose(/*in*//*optional*/hCryptMsg unsafe.Pointer) (BOOL) = crypt32.CryptMsgClose
//sys	CryptMsgUpdate(/*in*/hCryptMsg unsafe.Pointer, /*in*//*optional*/pbData *uint8, /*in*/cbData uint32, /*in*/fFinal BOOL) (BOOL) = crypt32.CryptMsgUpdate
//sys	CryptMsgGetParam(/*in*/hCryptMsg unsafe.Pointer, /*in*/dwParamType uint32, /*in*/dwIndex uint32, /*out*//*optional*/pvData unsafe.Pointer, /*in*//*out*/pcbData *uint32) (BOOL) = crypt32.CryptMsgGetParam
//sys	CryptMsgControl(/*in*/hCryptMsg unsafe.Pointer, /*in*/dwFlags uint32, /*in*/dwCtrlType uint32, /*in*//*optional*/pvCtrlPara unsafe.Pointer) (BOOL) = crypt32.CryptMsgControl
//sys	CryptMsgVerifyCountersignatureEncoded(/*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*/dwEncodingType uint32, /*in*/pbSignerInfo *uint8, /*in*/cbSignerInfo uint32, /*in*/pbSignerInfoCountersignature *uint8, /*in*/cbSignerInfoCountersignature uint32, /*in*/pciCountersigner *CERT_INFO) (BOOL) = crypt32.CryptMsgVerifyCountersignatureEncoded
//sys	CryptMsgVerifyCountersignatureEncodedEx(/*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*/dwEncodingType uint32, /*in*/pbSignerInfo *uint8, /*in*/cbSignerInfo uint32, /*in*/pbSignerInfoCountersignature *uint8, /*in*/cbSignerInfoCountersignature uint32, /*in*/dwSignerType uint32, /*in*/pvSigner unsafe.Pointer, /*in*/dwFlags uint32, /*in*//*out*//*optional*/pvExtra unsafe.Pointer) (BOOL) = crypt32.CryptMsgVerifyCountersignatureEncodedEx
//sys	CryptMsgCountersign(/*in*/hCryptMsg unsafe.Pointer, /*in*/dwIndex uint32, /*in*/cCountersigners uint32, /*in*/rgCountersigners *CMSG_SIGNER_ENCODE_INFO) (BOOL) = crypt32.CryptMsgCountersign
//sys	CryptMsgCountersignEncoded(/*in*/dwEncodingType uint32, /*in*/pbSignerInfo *uint8, /*in*/cbSignerInfo uint32, /*in*/cCountersigners uint32, /*in*/rgCountersigners *CMSG_SIGNER_ENCODE_INFO, /*out*//*optional*/pbCountersignature *uint8, /*in*//*out*/pcbCountersignature *uint32) (BOOL) = crypt32.CryptMsgCountersignEncoded
//sys	CertOpenStore(/*in*/lpszStoreProvider PSTR, /*in*/dwEncodingType CERT_QUERY_ENCODING_TYPE, /*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*/dwFlags CERT_OPEN_STORE_FLAGS, /*in*//*optional*/pvPara unsafe.Pointer) (HCERTSTORE) = crypt32.CertOpenStore
//sys	CertDuplicateStore(/*in*/hCertStore HCERTSTORE) (HCERTSTORE) = crypt32.CertDuplicateStore
//sys	CertSaveStore(/*in*/hCertStore HCERTSTORE, /*in*/dwEncodingType CERT_QUERY_ENCODING_TYPE, /*in*/dwSaveAs CERT_STORE_SAVE_AS, /*in*/dwSaveTo CERT_STORE_SAVE_TO, /*in*//*out*/pvSaveToPara unsafe.Pointer, /*in*/dwFlags uint32) (BOOL) = crypt32.CertSaveStore
//sys	CertCloseStore(/*in*//*optional*/hCertStore HCERTSTORE, /*in*/dwFlags uint32) (BOOL) = crypt32.CertCloseStore
//sys	CertGetSubjectCertificateFromStore(/*in*/hCertStore HCERTSTORE, /*in*/dwCertEncodingType uint32, /*in*/pCertId *CERT_INFO) (*CERT_CONTEXT) = crypt32.CertGetSubjectCertificateFromStore
//sys	CertEnumCertificatesInStore(/*in*/hCertStore HCERTSTORE, /*in*//*optional*/pPrevCertContext *CERT_CONTEXT) (*CERT_CONTEXT) = crypt32.CertEnumCertificatesInStore
//sys	CertFindCertificateInStore(/*in*/hCertStore HCERTSTORE, /*in*/dwCertEncodingType uint32, /*in*/dwFindFlags uint32, /*in*/dwFindType CERT_FIND_FLAGS, /*in*//*optional*/pvFindPara unsafe.Pointer, /*in*//*optional*/pPrevCertContext *CERT_CONTEXT) (*CERT_CONTEXT) = crypt32.CertFindCertificateInStore
//sys	CertGetIssuerCertificateFromStore(/*in*/hCertStore HCERTSTORE, /*in*/pSubjectContext *CERT_CONTEXT, /*in*//*optional*/pPrevIssuerContext *CERT_CONTEXT, /*in*//*out*/pdwFlags *uint32) (*CERT_CONTEXT) = crypt32.CertGetIssuerCertificateFromStore
//sys	CertVerifySubjectCertificateContext(/*in*/pSubject *CERT_CONTEXT, /*in*//*optional*/pIssuer *CERT_CONTEXT, /*in*//*out*/pdwFlags *uint32) (BOOL) = crypt32.CertVerifySubjectCertificateContext
//sys	CertDuplicateCertificateContext(/*in*//*optional*/pCertContext *CERT_CONTEXT) (*CERT_CONTEXT) = crypt32.CertDuplicateCertificateContext
//sys	CertCreateCertificateContext(/*in*/dwCertEncodingType uint32, /*in*/pbCertEncoded *uint8, /*in*/cbCertEncoded uint32) (*CERT_CONTEXT) = crypt32.CertCreateCertificateContext
//sys	CertFreeCertificateContext(/*in*//*optional*/pCertContext *CERT_CONTEXT) (BOOL) = crypt32.CertFreeCertificateContext
//sys	CertSetCertificateContextProperty(/*in*/pCertContext *CERT_CONTEXT, /*in*/dwPropId uint32, /*in*/dwFlags uint32, /*in*//*optional*/pvData unsafe.Pointer) (BOOL) = crypt32.CertSetCertificateContextProperty
//sys	CertGetCertificateContextProperty(/*in*/pCertContext *CERT_CONTEXT, /*in*/dwPropId uint32, /*out*//*optional*/pvData unsafe.Pointer, /*in*//*out*/pcbData *uint32) (BOOL) = crypt32.CertGetCertificateContextProperty
//sys	CertEnumCertificateContextProperties(/*in*/pCertContext *CERT_CONTEXT, /*in*/dwPropId uint32) (uint32) = crypt32.CertEnumCertificateContextProperties
//sys	CertCreateCTLEntryFromCertificateContextProperties(/*in*/pCertContext *CERT_CONTEXT, /*in*/cOptAttr uint32, /*in*//*optional*/rgOptAttr *CRYPT_ATTRIBUTE, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer, /*out*//*optional*/pCtlEntry *CTL_ENTRY, /*in*//*out*/pcbCtlEntry *uint32) (BOOL) = crypt32.CertCreateCTLEntryFromCertificateContextProperties
//sys	CertSetCertificateContextPropertiesFromCTLEntry(/*in*/pCertContext *CERT_CONTEXT, /*in*/pCtlEntry *CTL_ENTRY, /*in*/dwFlags uint32) (BOOL) = crypt32.CertSetCertificateContextPropertiesFromCTLEntry
//sys	CertGetCRLFromStore(/*in*/hCertStore HCERTSTORE, /*in*//*optional*/pIssuerContext *CERT_CONTEXT, /*in*//*optional*/pPrevCrlContext *CRL_CONTEXT, /*in*//*out*/pdwFlags *uint32) (*CRL_CONTEXT) = crypt32.CertGetCRLFromStore
//sys	CertEnumCRLsInStore(/*in*/hCertStore HCERTSTORE, /*in*//*optional*/pPrevCrlContext *CRL_CONTEXT) (*CRL_CONTEXT) = crypt32.CertEnumCRLsInStore
//sys	CertFindCRLInStore(/*in*/hCertStore HCERTSTORE, /*in*/dwCertEncodingType uint32, /*in*/dwFindFlags uint32, /*in*/dwFindType uint32, /*in*//*optional*/pvFindPara unsafe.Pointer, /*in*//*optional*/pPrevCrlContext *CRL_CONTEXT) (*CRL_CONTEXT) = crypt32.CertFindCRLInStore
//sys	CertDuplicateCRLContext(/*in*//*optional*/pCrlContext *CRL_CONTEXT) (*CRL_CONTEXT) = crypt32.CertDuplicateCRLContext
//sys	CertCreateCRLContext(/*in*/dwCertEncodingType uint32, /*in*/pbCrlEncoded *uint8, /*in*/cbCrlEncoded uint32) (*CRL_CONTEXT) = crypt32.CertCreateCRLContext
//sys	CertFreeCRLContext(/*in*//*optional*/pCrlContext *CRL_CONTEXT) (BOOL) = crypt32.CertFreeCRLContext
//sys	CertSetCRLContextProperty(/*in*/pCrlContext *CRL_CONTEXT, /*in*/dwPropId uint32, /*in*/dwFlags uint32, /*in*//*optional*/pvData unsafe.Pointer) (BOOL) = crypt32.CertSetCRLContextProperty
//sys	CertGetCRLContextProperty(/*in*/pCrlContext *CRL_CONTEXT, /*in*/dwPropId uint32, /*out*//*optional*/pvData unsafe.Pointer, /*in*//*out*/pcbData *uint32) (BOOL) = crypt32.CertGetCRLContextProperty
//sys	CertEnumCRLContextProperties(/*in*/pCrlContext *CRL_CONTEXT, /*in*/dwPropId uint32) (uint32) = crypt32.CertEnumCRLContextProperties
//sys	CertFindCertificateInCRL(/*in*/pCert *CERT_CONTEXT, /*in*/pCrlContext *CRL_CONTEXT, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer, /*out*//*optional*/ppCrlEntry **CRL_ENTRY) (BOOL) = crypt32.CertFindCertificateInCRL
//sys	CertIsValidCRLForCertificate(/*in*/pCert *CERT_CONTEXT, /*in*/pCrl *CRL_CONTEXT, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer) (BOOL) = crypt32.CertIsValidCRLForCertificate
//sys	CertAddEncodedCertificateToStore(/*in*//*optional*/hCertStore HCERTSTORE, /*in*/dwCertEncodingType uint32, /*in*/pbCertEncoded *uint8, /*in*/cbCertEncoded uint32, /*in*/dwAddDisposition uint32, /*out*//*optional*/ppCertContext **CERT_CONTEXT) (BOOL) = crypt32.CertAddEncodedCertificateToStore
//sys	CertAddCertificateContextToStore(/*in*//*optional*/hCertStore HCERTSTORE, /*in*/pCertContext *CERT_CONTEXT, /*in*/dwAddDisposition uint32, /*out*//*optional*/ppStoreContext **CERT_CONTEXT) (BOOL) = crypt32.CertAddCertificateContextToStore
//sys	CertAddSerializedElementToStore(/*in*//*optional*/hCertStore HCERTSTORE, /*in*/pbElement *uint8, /*in*/cbElement uint32, /*in*/dwAddDisposition uint32, /*in*/dwFlags uint32, /*in*/dwContextTypeFlags uint32, /*out*//*optional*/pdwContextType *uint32, /*out*//*optional*/ppvContext *unsafe.Pointer) (BOOL) = crypt32.CertAddSerializedElementToStore
//sys	CertDeleteCertificateFromStore(/*in*/pCertContext *CERT_CONTEXT) (BOOL) = crypt32.CertDeleteCertificateFromStore
//sys	CertAddEncodedCRLToStore(/*in*//*optional*/hCertStore HCERTSTORE, /*in*/dwCertEncodingType uint32, /*in*/pbCrlEncoded *uint8, /*in*/cbCrlEncoded uint32, /*in*/dwAddDisposition uint32, /*out*//*optional*/ppCrlContext **CRL_CONTEXT) (BOOL) = crypt32.CertAddEncodedCRLToStore
//sys	CertAddCRLContextToStore(/*in*//*optional*/hCertStore HCERTSTORE, /*in*/pCrlContext *CRL_CONTEXT, /*in*/dwAddDisposition uint32, /*out*//*optional*/ppStoreContext **CRL_CONTEXT) (BOOL) = crypt32.CertAddCRLContextToStore
//sys	CertDeleteCRLFromStore(/*in*/pCrlContext *CRL_CONTEXT) (BOOL) = crypt32.CertDeleteCRLFromStore
//sys	CertSerializeCertificateStoreElement(/*in*/pCertContext *CERT_CONTEXT, /*in*/dwFlags uint32, /*out*//*optional*/pbElement *uint8, /*in*//*out*/pcbElement *uint32) (BOOL) = crypt32.CertSerializeCertificateStoreElement
//sys	CertSerializeCRLStoreElement(/*in*/pCrlContext *CRL_CONTEXT, /*in*/dwFlags uint32, /*out*//*optional*/pbElement *uint8, /*in*//*out*/pcbElement *uint32) (BOOL) = crypt32.CertSerializeCRLStoreElement
//sys	CertDuplicateCTLContext(/*in*//*optional*/pCtlContext *CTL_CONTEXT) (*CTL_CONTEXT) = crypt32.CertDuplicateCTLContext
//sys	CertCreateCTLContext(/*in*/dwMsgAndCertEncodingType uint32, /*in*/pbCtlEncoded *uint8, /*in*/cbCtlEncoded uint32) (*CTL_CONTEXT) = crypt32.CertCreateCTLContext
//sys	CertFreeCTLContext(/*in*//*optional*/pCtlContext *CTL_CONTEXT) (BOOL) = crypt32.CertFreeCTLContext
//sys	CertSetCTLContextProperty(/*in*/pCtlContext *CTL_CONTEXT, /*in*/dwPropId uint32, /*in*/dwFlags uint32, /*in*//*optional*/pvData unsafe.Pointer) (BOOL) = crypt32.CertSetCTLContextProperty
//sys	CertGetCTLContextProperty(/*in*/pCtlContext *CTL_CONTEXT, /*in*/dwPropId uint32, /*out*//*optional*/pvData unsafe.Pointer, /*in*//*out*/pcbData *uint32) (BOOL) = crypt32.CertGetCTLContextProperty
//sys	CertEnumCTLContextProperties(/*in*/pCtlContext *CTL_CONTEXT, /*in*/dwPropId uint32) (uint32) = crypt32.CertEnumCTLContextProperties
//sys	CertEnumCTLsInStore(/*in*/hCertStore HCERTSTORE, /*in*//*optional*/pPrevCtlContext *CTL_CONTEXT) (*CTL_CONTEXT) = crypt32.CertEnumCTLsInStore
//sys	CertFindSubjectInCTL(/*in*/dwEncodingType uint32, /*in*/dwSubjectType uint32, /*in*/pvSubject unsafe.Pointer, /*in*/pCtlContext *CTL_CONTEXT, /*in*/dwFlags uint32) (*CTL_ENTRY) = crypt32.CertFindSubjectInCTL
//sys	CertFindCTLInStore(/*in*/hCertStore HCERTSTORE, /*in*/dwMsgAndCertEncodingType uint32, /*in*/dwFindFlags uint32, /*in*/dwFindType CERT_FIND_TYPE, /*in*//*optional*/pvFindPara unsafe.Pointer, /*in*//*optional*/pPrevCtlContext *CTL_CONTEXT) (*CTL_CONTEXT) = crypt32.CertFindCTLInStore
//sys	CertAddEncodedCTLToStore(/*in*//*optional*/hCertStore HCERTSTORE, /*in*/dwMsgAndCertEncodingType uint32, /*in*/pbCtlEncoded *uint8, /*in*/cbCtlEncoded uint32, /*in*/dwAddDisposition uint32, /*out*//*optional*/ppCtlContext **CTL_CONTEXT) (BOOL) = crypt32.CertAddEncodedCTLToStore
//sys	CertAddCTLContextToStore(/*in*//*optional*/hCertStore HCERTSTORE, /*in*/pCtlContext *CTL_CONTEXT, /*in*/dwAddDisposition uint32, /*out*//*optional*/ppStoreContext **CTL_CONTEXT) (BOOL) = crypt32.CertAddCTLContextToStore
//sys	CertSerializeCTLStoreElement(/*in*/pCtlContext *CTL_CONTEXT, /*in*/dwFlags uint32, /*out*//*optional*/pbElement *uint8, /*in*//*out*/pcbElement *uint32) (BOOL) = crypt32.CertSerializeCTLStoreElement
//sys	CertDeleteCTLFromStore(/*in*/pCtlContext *CTL_CONTEXT) (BOOL) = crypt32.CertDeleteCTLFromStore
//sys	CertAddCertificateLinkToStore(/*in*/hCertStore HCERTSTORE, /*in*/pCertContext *CERT_CONTEXT, /*in*/dwAddDisposition uint32, /*out*//*optional*/ppStoreContext **CERT_CONTEXT) (BOOL) = crypt32.CertAddCertificateLinkToStore
//sys	CertAddCRLLinkToStore(/*in*/hCertStore HCERTSTORE, /*in*/pCrlContext *CRL_CONTEXT, /*in*/dwAddDisposition uint32, /*out*//*optional*/ppStoreContext **CRL_CONTEXT) (BOOL) = crypt32.CertAddCRLLinkToStore
//sys	CertAddCTLLinkToStore(/*in*/hCertStore HCERTSTORE, /*in*/pCtlContext *CTL_CONTEXT, /*in*/dwAddDisposition uint32, /*out*//*optional*/ppStoreContext **CTL_CONTEXT) (BOOL) = crypt32.CertAddCTLLinkToStore
//sys	CertAddStoreToCollection(/*in*/hCollectionStore HCERTSTORE, /*in*//*optional*/hSiblingStore HCERTSTORE, /*in*/dwUpdateFlags uint32, /*in*/dwPriority uint32) (BOOL) = crypt32.CertAddStoreToCollection
//sys	CertRemoveStoreFromCollection(/*in*/hCollectionStore HCERTSTORE, /*in*/hSiblingStore HCERTSTORE) = crypt32.CertRemoveStoreFromCollection
//sys	CertControlStore(/*in*/hCertStore HCERTSTORE, /*in*/dwFlags CERT_CONTROL_STORE_FLAGS, /*in*/dwCtrlType uint32, /*in*//*optional*/pvCtrlPara unsafe.Pointer) (BOOL) = crypt32.CertControlStore
//sys	CertSetStoreProperty(/*in*/hCertStore HCERTSTORE, /*in*/dwPropId uint32, /*in*/dwFlags uint32, /*in*//*optional*/pvData unsafe.Pointer) (BOOL) = crypt32.CertSetStoreProperty
//sys	CertGetStoreProperty(/*in*/hCertStore HCERTSTORE, /*in*/dwPropId uint32, /*out*//*optional*/pvData unsafe.Pointer, /*in*//*out*/pcbData *uint32) (BOOL) = crypt32.CertGetStoreProperty
//sys	CertCreateContext(/*in*/dwContextType uint32, /*in*/dwEncodingType uint32, /*in*/pbEncoded *uint8, /*in*/cbEncoded uint32, /*in*/dwFlags uint32, /*in*//*optional*/pCreatePara *CERT_CREATE_CONTEXT_PARA) (unsafe.Pointer) = crypt32.CertCreateContext
//sys	CertRegisterSystemStore(/*in*/pvSystemStore unsafe.Pointer, /*in*/dwFlags uint32, /*in*//*optional*/pStoreInfo *CERT_SYSTEM_STORE_INFO, /*in*//*out*/pvReserved unsafe.Pointer) (BOOL) = crypt32.CertRegisterSystemStore
//sys	CertRegisterPhysicalStore(/*in*/pvSystemStore unsafe.Pointer, /*in*/dwFlags uint32, /*in*/pwszStoreName PWSTR, /*in*/pStoreInfo *CERT_PHYSICAL_STORE_INFO, /*in*//*out*/pvReserved unsafe.Pointer) (BOOL) = crypt32.CertRegisterPhysicalStore
//sys	CertUnregisterSystemStore(/*in*/pvSystemStore unsafe.Pointer, /*in*/dwFlags uint32) (BOOL) = crypt32.CertUnregisterSystemStore
//sys	CertUnregisterPhysicalStore(/*in*/pvSystemStore unsafe.Pointer, /*in*/dwFlags uint32, /*in*/pwszStoreName PWSTR) (BOOL) = crypt32.CertUnregisterPhysicalStore
//sys	CertEnumSystemStoreLocation(/*in*/dwFlags uint32, /*in*//*out*//*optional*/pvArg unsafe.Pointer, /*in*/pfnEnum PFN_CERT_ENUM_SYSTEM_STORE_LOCATION) (BOOL) = crypt32.CertEnumSystemStoreLocation
//sys	CertEnumSystemStore(/*in*/dwFlags uint32, /*in*//*optional*/pvSystemStoreLocationPara unsafe.Pointer, /*in*//*out*//*optional*/pvArg unsafe.Pointer, /*in*/pfnEnum PFN_CERT_ENUM_SYSTEM_STORE) (BOOL) = crypt32.CertEnumSystemStore
//sys	CertEnumPhysicalStore(/*in*/pvSystemStore unsafe.Pointer, /*in*/dwFlags uint32, /*in*//*out*//*optional*/pvArg unsafe.Pointer, /*in*/pfnEnum PFN_CERT_ENUM_PHYSICAL_STORE) (BOOL) = crypt32.CertEnumPhysicalStore
//sys	CertGetEnhancedKeyUsage(/*in*/pCertContext *CERT_CONTEXT, /*in*/dwFlags uint32, /*out*//*optional*/pUsage *CTL_USAGE, /*in*//*out*/pcbUsage *uint32) (BOOL) = crypt32.CertGetEnhancedKeyUsage
//sys	CertSetEnhancedKeyUsage(/*in*/pCertContext *CERT_CONTEXT, /*in*//*optional*/pUsage *CTL_USAGE) (BOOL) = crypt32.CertSetEnhancedKeyUsage
//sys	CertAddEnhancedKeyUsageIdentifier(/*in*/pCertContext *CERT_CONTEXT, /*in*/pszUsageIdentifier PSTR) (BOOL) = crypt32.CertAddEnhancedKeyUsageIdentifier
//sys	CertRemoveEnhancedKeyUsageIdentifier(/*in*/pCertContext *CERT_CONTEXT, /*in*/pszUsageIdentifier PSTR) (BOOL) = crypt32.CertRemoveEnhancedKeyUsageIdentifier
//sys	CertGetValidUsages(/*in*/cCerts uint32, /*in*/rghCerts **CERT_CONTEXT, /*out*/cNumOIDs *int32, /*out*//*optional*/rghOIDs *PSTR, /*in*//*out*/pcbOIDs *uint32) (BOOL) = crypt32.CertGetValidUsages
//sys	CryptMsgGetAndVerifySigner(/*in*/hCryptMsg unsafe.Pointer, /*in*/cSignerStore uint32, /*in*//*optional*/rghSignerStore *HCERTSTORE, /*in*/dwFlags uint32, /*out*//*optional*/ppSigner **CERT_CONTEXT, /*in*//*out*//*optional*/pdwSignerIndex *uint32) (BOOL) = crypt32.CryptMsgGetAndVerifySigner
//sys	CryptMsgSignCTL(/*in*/dwMsgEncodingType uint32, /*in*/pbCtlContent *uint8, /*in*/cbCtlContent uint32, /*in*/pSignInfo *CMSG_SIGNED_ENCODE_INFO, /*in*/dwFlags uint32, /*out*//*optional*/pbEncoded *uint8, /*in*//*out*/pcbEncoded *uint32) (BOOL) = crypt32.CryptMsgSignCTL
//sys	CryptMsgEncodeAndSignCTL(/*in*/dwMsgEncodingType uint32, /*in*/pCtlInfo *CTL_INFO, /*in*/pSignInfo *CMSG_SIGNED_ENCODE_INFO, /*in*/dwFlags uint32, /*out*//*optional*/pbEncoded *uint8, /*in*//*out*/pcbEncoded *uint32) (BOOL) = crypt32.CryptMsgEncodeAndSignCTL
//sys	CertFindSubjectInSortedCTL(/*in*/pSubjectIdentifier *CRYPTOAPI_BLOB, /*in*/pCtlContext *CTL_CONTEXT, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer, /*out*//*optional*/pEncodedAttributes *CRYPTOAPI_BLOB) (BOOL) = crypt32.CertFindSubjectInSortedCTL
//sys	CertEnumSubjectInSortedCTL(/*in*/pCtlContext *CTL_CONTEXT, /*in*//*out*/ppvNextSubject *unsafe.Pointer, /*out*//*optional*/pSubjectIdentifier *CRYPTOAPI_BLOB, /*out*//*optional*/pEncodedAttributes *CRYPTOAPI_BLOB) (BOOL) = crypt32.CertEnumSubjectInSortedCTL
//sys	CertVerifyCTLUsage(/*in*/dwEncodingType uint32, /*in*/dwSubjectType uint32, /*in*/pvSubject unsafe.Pointer, /*in*/pSubjectUsage *CTL_USAGE, /*in*/dwFlags uint32, /*in*//*optional*/pVerifyUsagePara *CTL_VERIFY_USAGE_PARA, /*in*//*out*/pVerifyUsageStatus *CTL_VERIFY_USAGE_STATUS) (BOOL) = crypt32.CertVerifyCTLUsage
//sys	CertVerifyRevocation(/*in*/dwEncodingType uint32, /*in*/dwRevType uint32, /*in*/cContext uint32, /*in*/rgpvContext *unsafe.Pointer, /*in*/dwFlags uint32, /*in*//*optional*/pRevPara *CERT_REVOCATION_PARA, /*in*//*out*/pRevStatus *CERT_REVOCATION_STATUS) (BOOL) = crypt32.CertVerifyRevocation
//sys	CertCompareIntegerBlob(/*in*/pInt1 *CRYPTOAPI_BLOB, /*in*/pInt2 *CRYPTOAPI_BLOB) (BOOL) = crypt32.CertCompareIntegerBlob
//sys	CertCompareCertificate(/*in*/dwCertEncodingType uint32, /*in*/pCertId1 *CERT_INFO, /*in*/pCertId2 *CERT_INFO) (BOOL) = crypt32.CertCompareCertificate
//sys	CertCompareCertificateName(/*in*/dwCertEncodingType uint32, /*in*/pCertName1 *CRYPTOAPI_BLOB, /*in*/pCertName2 *CRYPTOAPI_BLOB) (BOOL) = crypt32.CertCompareCertificateName
//sys	CertIsRDNAttrsInCertificateName(/*in*/dwCertEncodingType uint32, /*in*/dwFlags uint32, /*in*/pCertName *CRYPTOAPI_BLOB, /*in*/pRDN *CERT_RDN) (BOOL) = crypt32.CertIsRDNAttrsInCertificateName
//sys	CertComparePublicKeyInfo(/*in*/dwCertEncodingType uint32, /*in*/pPublicKey1 *CERT_PUBLIC_KEY_INFO, /*in*/pPublicKey2 *CERT_PUBLIC_KEY_INFO) (BOOL) = crypt32.CertComparePublicKeyInfo
//sys	CertGetPublicKeyLength(/*in*/dwCertEncodingType uint32, /*in*/pPublicKey *CERT_PUBLIC_KEY_INFO) (uint32) = crypt32.CertGetPublicKeyLength
//sys	CryptVerifyCertificateSignature(/*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*/dwCertEncodingType uint32, /*in*/pbEncoded *uint8, /*in*/cbEncoded uint32, /*in*/pPublicKey *CERT_PUBLIC_KEY_INFO) (BOOL) = crypt32.CryptVerifyCertificateSignature
//sys	CryptVerifyCertificateSignatureEx(/*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*/dwCertEncodingType uint32, /*in*/dwSubjectType uint32, /*in*/pvSubject unsafe.Pointer, /*in*/dwIssuerType uint32, /*in*//*optional*/pvIssuer unsafe.Pointer, /*in*/dwFlags CRYPT_VERIFY_CERT_FLAGS, /*in*//*out*//*optional*/pvExtra unsafe.Pointer) (BOOL) = crypt32.CryptVerifyCertificateSignatureEx
//sys	CertIsStrongHashToSign(/*in*/pStrongSignPara *CERT_STRONG_SIGN_PARA, /*in*/pwszCNGHashAlgid PWSTR, /*in*//*optional*/pSigningCert *CERT_CONTEXT) (BOOL) = crypt32.CertIsStrongHashToSign
//sys	CryptHashToBeSigned(/*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*/dwCertEncodingType uint32, /*in*/pbEncoded *uint8, /*in*/cbEncoded uint32, /*out*//*optional*/pbComputedHash *uint8, /*in*//*out*/pcbComputedHash *uint32) (BOOL) = crypt32.CryptHashToBeSigned
//sys	CryptHashCertificate(/*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*/Algid uint32, /*in*/dwFlags uint32, /*in*/pbEncoded *uint8, /*in*/cbEncoded uint32, /*out*//*optional*/pbComputedHash *uint8, /*in*//*out*/pcbComputedHash *uint32) (BOOL) = crypt32.CryptHashCertificate
//sys	CryptHashCertificate2(/*in*/pwszCNGHashAlgid PWSTR, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer, /*in*//*optional*/pbEncoded *uint8, /*in*/cbEncoded uint32, /*out*//*optional*/pbComputedHash *uint8, /*in*//*out*/pcbComputedHash *uint32) (BOOL) = crypt32.CryptHashCertificate2
//sys	CryptSignCertificate(/*in*//*optional*/hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, /*in*//*optional*/dwKeySpec uint32, /*in*/dwCertEncodingType uint32, /*in*/pbEncodedToBeSigned *uint8, /*in*/cbEncodedToBeSigned uint32, /*in*/pSignatureAlgorithm *CRYPT_ALGORITHM_IDENTIFIER, /*in*//*optional*/pvHashAuxInfo unsafe.Pointer, /*out*//*optional*/pbSignature *uint8, /*in*//*out*/pcbSignature *uint32) (BOOL) = crypt32.CryptSignCertificate
//sys	CryptSignAndEncodeCertificate(/*in*//*optional*/hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, /*in*//*optional*/dwKeySpec CERT_KEY_SPEC, /*in*/dwCertEncodingType uint32, /*in*/lpszStructType PSTR, /*in*/pvStructInfo unsafe.Pointer, /*in*/pSignatureAlgorithm *CRYPT_ALGORITHM_IDENTIFIER, /*in*//*optional*/pvHashAuxInfo unsafe.Pointer, /*out*//*optional*/pbEncoded *uint8, /*in*//*out*/pcbEncoded *uint32) (BOOL) = crypt32.CryptSignAndEncodeCertificate
//sys	CertVerifyTimeValidity(/*in*//*optional*/pTimeToVerify *FILETIME, /*in*/pCertInfo *CERT_INFO) (int32) = crypt32.CertVerifyTimeValidity
//sys	CertVerifyCRLTimeValidity(/*in*//*optional*/pTimeToVerify *FILETIME, /*in*/pCrlInfo *CRL_INFO) (int32) = crypt32.CertVerifyCRLTimeValidity
//sys	CertVerifyValidityNesting(/*in*/pSubjectInfo *CERT_INFO, /*in*/pIssuerInfo *CERT_INFO) (BOOL) = crypt32.CertVerifyValidityNesting
//sys	CertVerifyCRLRevocation(/*in*/dwCertEncodingType uint32, /*in*/pCertId *CERT_INFO, /*in*/cCrlInfo uint32, /*in*/rgpCrlInfo **CRL_INFO) (BOOL) = crypt32.CertVerifyCRLRevocation
//sys	CertAlgIdToOID(/*in*/dwAlgId uint32) (PSTR) = crypt32.CertAlgIdToOID
//sys	CertOIDToAlgId(/*in*/pszObjId PSTR) (uint32) = crypt32.CertOIDToAlgId
//sys	CertFindExtension(/*in*/pszObjId PSTR, /*in*/cExtensions uint32, /*in*/rgExtensions *CERT_EXTENSION) (*CERT_EXTENSION) = crypt32.CertFindExtension
//sys	CertFindAttribute(/*in*/pszObjId PSTR, /*in*/cAttr uint32, /*in*/rgAttr *CRYPT_ATTRIBUTE) (*CRYPT_ATTRIBUTE) = crypt32.CertFindAttribute
//sys	CertFindRDNAttr(/*in*/pszObjId PSTR, /*in*/pName *CERT_NAME_INFO) (*CERT_RDN_ATTR) = crypt32.CertFindRDNAttr
//sys	CertGetIntendedKeyUsage(/*in*/dwCertEncodingType uint32, /*in*/pCertInfo *CERT_INFO, /*out*/pbKeyUsage *uint8, /*in*/cbKeyUsage uint32) (BOOL) = crypt32.CertGetIntendedKeyUsage
//sys	CryptInstallDefaultContext(/*in*/hCryptProv uintptr, /*in*/dwDefaultType CRYPT_DEFAULT_CONTEXT_TYPE, /*in*//*optional*/pvDefaultPara unsafe.Pointer, /*in*/dwFlags CRYPT_DEFAULT_CONTEXT_FLAGS, /*in*//*out*/pvReserved unsafe.Pointer, /*out*/phDefaultContext *unsafe.Pointer) (BOOL) = crypt32.CryptInstallDefaultContext
//sys	CryptUninstallDefaultContext(/*in*//*optional*/hDefaultContext unsafe.Pointer, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer) (BOOL) = crypt32.CryptUninstallDefaultContext
//sys	CryptExportPublicKeyInfo(/*in*/hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, /*in*//*optional*/dwKeySpec uint32, /*in*/dwCertEncodingType uint32, /*out*//*optional*/pInfo *CERT_PUBLIC_KEY_INFO, /*in*//*out*/pcbInfo *uint32) (BOOL) = crypt32.CryptExportPublicKeyInfo
//sys	CryptExportPublicKeyInfoEx(/*in*/hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, /*in*//*optional*/dwKeySpec uint32, /*in*/dwCertEncodingType uint32, /*in*//*optional*/pszPublicKeyObjId PSTR, /*in*/dwFlags uint32, /*in*//*optional*/pvAuxInfo unsafe.Pointer, /*out*//*optional*/pInfo *CERT_PUBLIC_KEY_INFO, /*in*//*out*/pcbInfo *uint32) (BOOL) = crypt32.CryptExportPublicKeyInfoEx
//sys	CryptExportPublicKeyInfoFromBCryptKeyHandle(/*in*/hBCryptKey BCRYPT_KEY_HANDLE, /*in*/dwCertEncodingType uint32, /*in*//*optional*/pszPublicKeyObjId PSTR, /*in*/dwFlags uint32, /*in*//*optional*/pvAuxInfo unsafe.Pointer, /*out*//*optional*/pInfo *CERT_PUBLIC_KEY_INFO, /*in*//*out*/pcbInfo *uint32) (BOOL) = crypt32.CryptExportPublicKeyInfoFromBCryptKeyHandle
//sys	CryptImportPublicKeyInfo(/*in*/hCryptProv uintptr, /*in*/dwCertEncodingType uint32, /*in*/pInfo *CERT_PUBLIC_KEY_INFO, /*out*/phKey *uintptr) (BOOL) = crypt32.CryptImportPublicKeyInfo
//sys	CryptImportPublicKeyInfoEx(/*in*/hCryptProv uintptr, /*in*/dwCertEncodingType uint32, /*in*/pInfo *CERT_PUBLIC_KEY_INFO, /*in*/aiKeyAlg uint32, /*in*/dwFlags uint32, /*in*//*optional*/pvAuxInfo unsafe.Pointer, /*out*/phKey *uintptr) (BOOL) = crypt32.CryptImportPublicKeyInfoEx
//sys	CryptImportPublicKeyInfoEx2(/*in*/dwCertEncodingType uint32, /*in*/pInfo *CERT_PUBLIC_KEY_INFO, /*in*/dwFlags CRYPT_IMPORT_PUBLIC_KEY_FLAGS, /*in*//*optional*/pvAuxInfo unsafe.Pointer, /*out*/phKey *BCRYPT_KEY_HANDLE) (BOOL) = crypt32.CryptImportPublicKeyInfoEx2
//sys	CryptAcquireCertificatePrivateKey(/*in*/pCert *CERT_CONTEXT, /*in*/dwFlags CRYPT_ACQUIRE_FLAGS, /*in*//*optional*/pvParameters unsafe.Pointer, /*out*/phCryptProvOrNCryptKey *HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, /*out*//*optional*/pdwKeySpec *CERT_KEY_SPEC, /*out*//*optional*/pfCallerFreeProvOrNCryptKey *BOOL) (BOOL) = crypt32.CryptAcquireCertificatePrivateKey
//sys	CryptFindCertificateKeyProvInfo(/*in*/pCert *CERT_CONTEXT, /*in*/dwFlags CRYPT_FIND_FLAGS, /*in*//*out*/pvReserved unsafe.Pointer) (BOOL) = crypt32.CryptFindCertificateKeyProvInfo
//sys	CryptImportPKCS8(/*in*/sPrivateKeyAndParams CRYPT_PKCS8_IMPORT_PARAMS, /*in*/dwFlags CRYPT_KEY_FLAGS, /*out*//*optional*/phCryptProv *uintptr, /*in*//*optional*/pvAuxInfo unsafe.Pointer) (BOOL) = crypt32.CryptImportPKCS8
//sys	CryptExportPKCS8(/*in*/hCryptProv uintptr, /*in*/dwKeySpec uint32, /*in*/pszPrivateKeyObjId PSTR, /*in*/dwFlags uint32, /*in*//*optional*/pvAuxInfo unsafe.Pointer, /*out*//*optional*/pbPrivateKeyBlob *uint8, /*in*//*out*/pcbPrivateKeyBlob *uint32) (BOOL) = crypt32.CryptExportPKCS8
//sys	CryptHashPublicKeyInfo(/*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*/Algid uint32, /*in*/dwFlags uint32, /*in*/dwCertEncodingType uint32, /*in*/pInfo *CERT_PUBLIC_KEY_INFO, /*out*//*optional*/pbComputedHash *uint8, /*in*//*out*/pcbComputedHash *uint32) (BOOL) = crypt32.CryptHashPublicKeyInfo
//sys	CertRDNValueToStrA(/*in*/dwValueType uint32, /*in*/pValue *CRYPTOAPI_BLOB, /*out*//*optional*/psz PSTR, /*in*/csz uint32) (uint32) = crypt32.CertRDNValueToStrA
//sys	CertRDNValueToStrW(/*in*/dwValueType uint32, /*in*/pValue *CRYPTOAPI_BLOB, /*out*//*optional*/psz PWSTR, /*in*/csz uint32) (uint32) = crypt32.CertRDNValueToStrW
//sys	CertNameToStrA(/*in*/dwCertEncodingType uint32, /*in*/pName *CRYPTOAPI_BLOB, /*in*/dwStrType CERT_STRING_TYPE, /*out*//*optional*/psz PSTR, /*in*/csz uint32) (uint32) = crypt32.CertNameToStrA
//sys	CertNameToStrW(/*in*/dwCertEncodingType uint32, /*in*/pName *CRYPTOAPI_BLOB, /*in*/dwStrType CERT_STRING_TYPE, /*out*//*optional*/psz PWSTR, /*in*/csz uint32) (uint32) = crypt32.CertNameToStrW
//sys	CertStrToNameA(/*in*/dwCertEncodingType uint32, /*in*/pszX500 PSTR, /*in*/dwStrType CERT_STRING_TYPE, /*in*//*out*/pvReserved unsafe.Pointer, /*out*//*optional*/pbEncoded *uint8, /*in*//*out*/pcbEncoded *uint32, /*out*//*optional*/ppszError *PSTR) (BOOL) = crypt32.CertStrToNameA
//sys	CertStrToNameW(/*in*/dwCertEncodingType uint32, /*in*/pszX500 PWSTR, /*in*/dwStrType CERT_STRING_TYPE, /*in*//*out*/pvReserved unsafe.Pointer, /*out*//*optional*/pbEncoded *uint8, /*in*//*out*/pcbEncoded *uint32, /*out*//*optional*/ppszError *PWSTR) (BOOL) = crypt32.CertStrToNameW
//sys	CertGetNameStringA(/*in*/pCertContext *CERT_CONTEXT, /*in*/dwType uint32, /*in*/dwFlags uint32, /*in*//*optional*/pvTypePara unsafe.Pointer, /*out*//*optional*/pszNameString PSTR, /*in*/cchNameString uint32) (uint32) = crypt32.CertGetNameStringA
//sys	CertGetNameStringW(/*in*/pCertContext *CERT_CONTEXT, /*in*/dwType uint32, /*in*/dwFlags uint32, /*in*//*optional*/pvTypePara unsafe.Pointer, /*out*//*optional*/pszNameString PWSTR, /*in*/cchNameString uint32) (uint32) = crypt32.CertGetNameStringW
//sys	CryptSignMessage(/*in*/pSignPara *CRYPT_SIGN_MESSAGE_PARA, /*in*/fDetachedSignature BOOL, /*in*/cToBeSigned uint32, /*in*//*optional*/rgpbToBeSigned **uint8, /*in*/rgcbToBeSigned *uint32, /*out*//*optional*/pbSignedBlob *uint8, /*in*//*out*/pcbSignedBlob *uint32) (BOOL) = crypt32.CryptSignMessage
//sys	CryptVerifyMessageSignature(/*in*/pVerifyPara *CRYPT_VERIFY_MESSAGE_PARA, /*in*/dwSignerIndex uint32, /*in*/pbSignedBlob *uint8, /*in*/cbSignedBlob uint32, /*out*//*optional*/pbDecoded *uint8, /*in*//*out*//*optional*/pcbDecoded *uint32, /*out*//*optional*/ppSignerCert **CERT_CONTEXT) (BOOL) = crypt32.CryptVerifyMessageSignature
//sys	CryptGetMessageSignerCount(/*in*/dwMsgEncodingType uint32, /*in*/pbSignedBlob *uint8, /*in*/cbSignedBlob uint32) (int32) = crypt32.CryptGetMessageSignerCount
//sys	CryptGetMessageCertificates(/*in*/dwMsgAndCertEncodingType uint32, /*in*//*optional*/hCryptProv HCRYPTPROV_LEGACY, /*in*/dwFlags uint32, /*in*/pbSignedBlob *uint8, /*in*/cbSignedBlob uint32) (HCERTSTORE) = crypt32.CryptGetMessageCertificates
//sys	CryptVerifyDetachedMessageSignature(/*in*/pVerifyPara *CRYPT_VERIFY_MESSAGE_PARA, /*in*/dwSignerIndex uint32, /*in*/pbDetachedSignBlob *uint8, /*in*/cbDetachedSignBlob uint32, /*in*/cToBeSigned uint32, /*in*/rgpbToBeSigned **uint8, /*in*/rgcbToBeSigned *uint32, /*out*//*optional*/ppSignerCert **CERT_CONTEXT) (BOOL) = crypt32.CryptVerifyDetachedMessageSignature
//sys	CryptEncryptMessage(/*in*/pEncryptPara *CRYPT_ENCRYPT_MESSAGE_PARA, /*in*/cRecipientCert uint32, /*in*/rgpRecipientCert **CERT_CONTEXT, /*in*//*optional*/pbToBeEncrypted *uint8, /*in*/cbToBeEncrypted uint32, /*out*//*optional*/pbEncryptedBlob *uint8, /*in*//*out*/pcbEncryptedBlob *uint32) (BOOL) = crypt32.CryptEncryptMessage
//sys	CryptDecryptMessage(/*in*/pDecryptPara *CRYPT_DECRYPT_MESSAGE_PARA, /*in*/pbEncryptedBlob *uint8, /*in*/cbEncryptedBlob uint32, /*out*//*optional*/pbDecrypted *uint8, /*in*//*out*//*optional*/pcbDecrypted *uint32, /*out*//*optional*/ppXchgCert **CERT_CONTEXT) (BOOL) = crypt32.CryptDecryptMessage
//sys	CryptSignAndEncryptMessage(/*in*/pSignPara *CRYPT_SIGN_MESSAGE_PARA, /*in*/pEncryptPara *CRYPT_ENCRYPT_MESSAGE_PARA, /*in*/cRecipientCert uint32, /*in*/rgpRecipientCert **CERT_CONTEXT, /*in*/pbToBeSignedAndEncrypted *uint8, /*in*/cbToBeSignedAndEncrypted uint32, /*out*//*optional*/pbSignedAndEncryptedBlob *uint8, /*in*//*out*/pcbSignedAndEncryptedBlob *uint32) (BOOL) = crypt32.CryptSignAndEncryptMessage
//sys	CryptDecryptAndVerifyMessageSignature(/*in*/pDecryptPara *CRYPT_DECRYPT_MESSAGE_PARA, /*in*/pVerifyPara *CRYPT_VERIFY_MESSAGE_PARA, /*in*/dwSignerIndex uint32, /*in*/pbEncryptedBlob *uint8, /*in*/cbEncryptedBlob uint32, /*out*//*optional*/pbDecrypted *uint8, /*in*//*out*//*optional*/pcbDecrypted *uint32, /*out*//*optional*/ppXchgCert **CERT_CONTEXT, /*out*//*optional*/ppSignerCert **CERT_CONTEXT) (BOOL) = crypt32.CryptDecryptAndVerifyMessageSignature
//sys	CryptDecodeMessage(/*in*/dwMsgTypeFlags uint32, /*in*//*optional*/pDecryptPara *CRYPT_DECRYPT_MESSAGE_PARA, /*in*//*optional*/pVerifyPara *CRYPT_VERIFY_MESSAGE_PARA, /*in*/dwSignerIndex uint32, /*in*/pbEncodedBlob *uint8, /*in*/cbEncodedBlob uint32, /*in*/dwPrevInnerContentType uint32, /*out*//*optional*/pdwMsgType *uint32, /*out*//*optional*/pdwInnerContentType *uint32, /*out*//*optional*/pbDecoded *uint8, /*in*//*out*//*optional*/pcbDecoded *uint32, /*out*//*optional*/ppXchgCert **CERT_CONTEXT, /*out*//*optional*/ppSignerCert **CERT_CONTEXT) (BOOL) = crypt32.CryptDecodeMessage
//sys	CryptHashMessage(/*in*/pHashPara *CRYPT_HASH_MESSAGE_PARA, /*in*/fDetachedHash BOOL, /*in*/cToBeHashed uint32, /*in*/rgpbToBeHashed **uint8, /*in*/rgcbToBeHashed *uint32, /*out*//*optional*/pbHashedBlob *uint8, /*in*//*out*//*optional*/pcbHashedBlob *uint32, /*out*//*optional*/pbComputedHash *uint8, /*in*//*out*//*optional*/pcbComputedHash *uint32) (BOOL) = crypt32.CryptHashMessage
//sys	CryptVerifyMessageHash(/*in*/pHashPara *CRYPT_HASH_MESSAGE_PARA, /*in*/pbHashedBlob *uint8, /*in*/cbHashedBlob uint32, /*out*//*optional*/pbToBeHashed *uint8, /*in*//*out*//*optional*/pcbToBeHashed *uint32, /*out*//*optional*/pbComputedHash *uint8, /*in*//*out*//*optional*/pcbComputedHash *uint32) (BOOL) = crypt32.CryptVerifyMessageHash
//sys	CryptVerifyDetachedMessageHash(/*in*/pHashPara *CRYPT_HASH_MESSAGE_PARA, /*in*/pbDetachedHashBlob *uint8, /*in*/cbDetachedHashBlob uint32, /*in*/cToBeHashed uint32, /*in*/rgpbToBeHashed **uint8, /*in*/rgcbToBeHashed *uint32, /*out*//*optional*/pbComputedHash *uint8, /*in*//*out*//*optional*/pcbComputedHash *uint32) (BOOL) = crypt32.CryptVerifyDetachedMessageHash
//sys	CryptSignMessageWithKey(/*in*/pSignPara *CRYPT_KEY_SIGN_MESSAGE_PARA, /*in*/pbToBeSigned *uint8, /*in*/cbToBeSigned uint32, /*out*//*optional*/pbSignedBlob *uint8, /*in*//*out*/pcbSignedBlob *uint32) (BOOL) = crypt32.CryptSignMessageWithKey
//sys	CryptVerifyMessageSignatureWithKey(/*in*/pVerifyPara *CRYPT_KEY_VERIFY_MESSAGE_PARA, /*in*//*optional*/pPublicKeyInfo *CERT_PUBLIC_KEY_INFO, /*in*/pbSignedBlob *uint8, /*in*/cbSignedBlob uint32, /*out*//*optional*/pbDecoded *uint8, /*in*//*out*//*optional*/pcbDecoded *uint32) (BOOL) = crypt32.CryptVerifyMessageSignatureWithKey
//sys	CertOpenSystemStoreA(/*in*//*optional*/hProv HCRYPTPROV_LEGACY, /*in*/szSubsystemProtocol PSTR) (HCERTSTORE) = crypt32.CertOpenSystemStoreA
//sys	CertOpenSystemStoreW(/*in*//*optional*/hProv HCRYPTPROV_LEGACY, /*in*/szSubsystemProtocol PWSTR) (HCERTSTORE) = crypt32.CertOpenSystemStoreW
//sys	CertAddEncodedCertificateToSystemStoreA(/*in*/szCertStoreName PSTR, /*in*/pbCertEncoded *uint8, /*in*/cbCertEncoded uint32) (BOOL) = crypt32.CertAddEncodedCertificateToSystemStoreA
//sys	CertAddEncodedCertificateToSystemStoreW(/*in*/szCertStoreName PWSTR, /*in*/pbCertEncoded *uint8, /*in*/cbCertEncoded uint32) (BOOL) = crypt32.CertAddEncodedCertificateToSystemStoreW
//sys	FindCertsByIssuer(/*out*//*optional*/pCertChains *CERT_CHAIN, /*in*//*out*/pcbCertChains *uint32, /*out*/pcCertChains *uint32, /*in*//*optional*/pbEncodedIssuerName *uint8, /*in*/cbEncodedIssuerName uint32, /*in*//*optional*/pwszPurpose PWSTR, /*in*/dwKeySpec uint32) (HRESULT) = wintrust.FindCertsByIssuer
//sys	CryptQueryObject(/*in*/dwObjectType CERT_QUERY_OBJECT_TYPE, /*in*/pvObject unsafe.Pointer, /*in*/dwExpectedContentTypeFlags CERT_QUERY_CONTENT_TYPE_FLAGS, /*in*/dwExpectedFormatTypeFlags CERT_QUERY_FORMAT_TYPE_FLAGS, /*in*/dwFlags uint32, /*out*//*optional*/pdwMsgAndCertEncodingType *CERT_QUERY_ENCODING_TYPE, /*out*//*optional*/pdwContentType *CERT_QUERY_CONTENT_TYPE, /*out*//*optional*/pdwFormatType *CERT_QUERY_FORMAT_TYPE, /*out*//*optional*/phCertStore *HCERTSTORE, /*out*//*optional*/phMsg *unsafe.Pointer, /*out*//*optional*/ppvContext *unsafe.Pointer) (BOOL) = crypt32.CryptQueryObject
//sys	CryptMemAlloc(/*in*/cbSize uint32) (unsafe.Pointer) = crypt32.CryptMemAlloc
//sys	CryptMemRealloc(/*in*//*optional*/pv unsafe.Pointer, /*in*/cbSize uint32) (unsafe.Pointer) = crypt32.CryptMemRealloc
//sys	CryptMemFree(/*in*//*optional*/pv unsafe.Pointer) = crypt32.CryptMemFree
//sys	CryptCreateAsyncHandle(/*in*/dwFlags uint32, /*out*/phAsync *HCRYPTASYNC) (BOOL) = crypt32.CryptCreateAsyncHandle
//sys	CryptSetAsyncParam(/*in*/hAsync HCRYPTASYNC, /*in*/pszParamOid PSTR, /*in*//*optional*/pvParam unsafe.Pointer, /*in*/pfnFree PFN_CRYPT_ASYNC_PARAM_FREE_FUNC) (BOOL) = crypt32.CryptSetAsyncParam
//sys	CryptGetAsyncParam(/*in*/hAsync HCRYPTASYNC, /*in*/pszParamOid PSTR, /*out*//*optional*/ppvParam *unsafe.Pointer, /*out*//*optional*/ppfnFree *PFN_CRYPT_ASYNC_PARAM_FREE_FUNC) (BOOL) = crypt32.CryptGetAsyncParam
//sys	CryptCloseAsyncHandle(/*in*//*optional*/hAsync HCRYPTASYNC) (BOOL) = crypt32.CryptCloseAsyncHandle
//sys	CryptRetrieveObjectByUrlA(/*in*/pszUrl PSTR, /*in*//*optional*/pszObjectOid PSTR, /*in*/dwRetrievalFlags uint32, /*in*/dwTimeout uint32, /*out*/ppvObject *unsafe.Pointer, /*in*//*optional*/hAsyncRetrieve HCRYPTASYNC, /*in*//*optional*/pCredentials *CRYPT_CREDENTIALS, /*in*//*optional*/pvVerify unsafe.Pointer, /*in*//*out*//*optional*/pAuxInfo *CRYPT_RETRIEVE_AUX_INFO) (BOOL) = cryptnet.CryptRetrieveObjectByUrlA
//sys	CryptRetrieveObjectByUrlW(/*in*/pszUrl PWSTR, /*in*//*optional*/pszObjectOid PSTR, /*in*/dwRetrievalFlags uint32, /*in*/dwTimeout uint32, /*out*/ppvObject *unsafe.Pointer, /*in*//*optional*/hAsyncRetrieve HCRYPTASYNC, /*in*//*optional*/pCredentials *CRYPT_CREDENTIALS, /*in*//*optional*/pvVerify unsafe.Pointer, /*in*//*out*//*optional*/pAuxInfo *CRYPT_RETRIEVE_AUX_INFO) (BOOL) = cryptnet.CryptRetrieveObjectByUrlW
//sys	CryptInstallCancelRetrieval(/*in*/pfnCancel PFN_CRYPT_CANCEL_RETRIEVAL, /*in*//*optional*/pvArg unsafe.Pointer, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer) (BOOL) = cryptnet.CryptInstallCancelRetrieval
//sys	CryptUninstallCancelRetrieval(/*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer) (BOOL) = cryptnet.CryptUninstallCancelRetrieval
//sys	CryptGetObjectUrl(/*in*/pszUrlOid PSTR, /*in*/pvPara unsafe.Pointer, /*in*/dwFlags CRYPT_GET_URL_FLAGS, /*out*//*optional*/pUrlArray *CRYPT_URL_ARRAY, /*in*//*out*/pcbUrlArray *uint32, /*out*//*optional*/pUrlInfo *CRYPT_URL_INFO, /*in*//*out*//*optional*/pcbUrlInfo *uint32, /*in*//*out*/pvReserved unsafe.Pointer) (BOOL) = cryptnet.CryptGetObjectUrl
//sys	CertCreateSelfSignCertificate(/*in*//*optional*/hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, /*in*/pSubjectIssuerBlob *CRYPTOAPI_BLOB, /*in*/dwFlags CERT_CREATE_SELFSIGN_FLAGS, /*in*//*optional*/pKeyProvInfo *CRYPT_KEY_PROV_INFO, /*in*//*optional*/pSignatureAlgorithm *CRYPT_ALGORITHM_IDENTIFIER, /*in*//*optional*/pStartTime *SYSTEMTIME, /*in*//*optional*/pEndTime *SYSTEMTIME, /*in*//*optional*/pExtensions *CERT_EXTENSIONS) (*CERT_CONTEXT) = crypt32.CertCreateSelfSignCertificate
//sys	CryptGetKeyIdentifierProperty(/*in*/pKeyIdentifier *CRYPTOAPI_BLOB, /*in*/dwPropId uint32, /*in*/dwFlags uint32, /*in*//*optional*/pwszComputerName PWSTR, /*in*//*out*/pvReserved unsafe.Pointer, /*out*//*optional*/pvData unsafe.Pointer, /*in*//*out*/pcbData *uint32) (BOOL) = crypt32.CryptGetKeyIdentifierProperty
//sys	CryptSetKeyIdentifierProperty(/*in*/pKeyIdentifier *CRYPTOAPI_BLOB, /*in*/dwPropId uint32, /*in*/dwFlags uint32, /*in*//*optional*/pwszComputerName PWSTR, /*in*//*out*/pvReserved unsafe.Pointer, /*in*//*optional*/pvData unsafe.Pointer) (BOOL) = crypt32.CryptSetKeyIdentifierProperty
//sys	CryptEnumKeyIdentifierProperties(/*in*//*optional*/pKeyIdentifier *CRYPTOAPI_BLOB, /*in*/dwPropId uint32, /*in*/dwFlags uint32, /*in*//*optional*/pwszComputerName PWSTR, /*in*//*out*/pvReserved unsafe.Pointer, /*in*//*out*//*optional*/pvArg unsafe.Pointer, /*in*/pfnEnum PFN_CRYPT_ENUM_KEYID_PROP) (BOOL) = crypt32.CryptEnumKeyIdentifierProperties
//sys	CryptCreateKeyIdentifierFromCSP(/*in*/dwCertEncodingType uint32, /*in*//*optional*/pszPubKeyOID PSTR, /*in*/pPubKeyStruc *PUBLICKEYSTRUC, /*in*/cbPubKeyStruc uint32, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer, /*out*//*optional*/pbHash *uint8, /*in*//*out*/pcbHash *uint32) (BOOL) = crypt32.CryptCreateKeyIdentifierFromCSP
//sys	CertCreateCertificateChainEngine(/*in*/pConfig *CERT_CHAIN_ENGINE_CONFIG, /*out*/phChainEngine *HCERTCHAINENGINE) (BOOL) = crypt32.CertCreateCertificateChainEngine
//sys	CertFreeCertificateChainEngine(/*in*//*optional*/hChainEngine HCERTCHAINENGINE) = crypt32.CertFreeCertificateChainEngine
//sys	CertResyncCertificateChainEngine(/*in*//*optional*/hChainEngine HCERTCHAINENGINE) (BOOL) = crypt32.CertResyncCertificateChainEngine
//sys	CertGetCertificateChain(/*in*//*optional*/hChainEngine HCERTCHAINENGINE, /*in*/pCertContext *CERT_CONTEXT, /*in*//*optional*/pTime *FILETIME, /*in*//*optional*/hAdditionalStore HCERTSTORE, /*in*/pChainPara *CERT_CHAIN_PARA, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer, /*out*/ppChainContext **CERT_CHAIN_CONTEXT) (BOOL) = crypt32.CertGetCertificateChain
//sys	CertFreeCertificateChain(/*in*/pChainContext *CERT_CHAIN_CONTEXT) = crypt32.CertFreeCertificateChain
//sys	CertDuplicateCertificateChain(/*in*/pChainContext *CERT_CHAIN_CONTEXT) (*CERT_CHAIN_CONTEXT) = crypt32.CertDuplicateCertificateChain
//sys	CertFindChainInStore(/*in*/hCertStore HCERTSTORE, /*in*/dwCertEncodingType uint32, /*in*/dwFindFlags CERT_FIND_CHAIN_IN_STORE_FLAGS, /*in*/dwFindType uint32, /*in*//*optional*/pvFindPara unsafe.Pointer, /*in*//*optional*/pPrevChainContext *CERT_CHAIN_CONTEXT) (*CERT_CHAIN_CONTEXT) = crypt32.CertFindChainInStore
//sys	CertVerifyCertificateChainPolicy(/*in*/pszPolicyOID PSTR, /*in*/pChainContext *CERT_CHAIN_CONTEXT, /*in*/pPolicyPara *CERT_CHAIN_POLICY_PARA, /*in*//*out*/pPolicyStatus *CERT_CHAIN_POLICY_STATUS) (BOOL) = crypt32.CertVerifyCertificateChainPolicy
//sys	CryptStringToBinaryA(/*in*/pszString PSTR, /*in*/cchString uint32, /*in*/dwFlags CRYPT_STRING, /*out*//*optional*/pbBinary *uint8, /*in*//*out*/pcbBinary *uint32, /*out*//*optional*/pdwSkip *uint32, /*out*//*optional*/pdwFlags *uint32) (BOOL) = crypt32.CryptStringToBinaryA
//sys	CryptStringToBinaryW(/*in*/pszString PWSTR, /*in*/cchString uint32, /*in*/dwFlags CRYPT_STRING, /*out*//*optional*/pbBinary *uint8, /*in*//*out*/pcbBinary *uint32, /*out*//*optional*/pdwSkip *uint32, /*out*//*optional*/pdwFlags *uint32) (BOOL) = crypt32.CryptStringToBinaryW
//sys	CryptBinaryToStringA(/*in*/pbBinary *uint8, /*in*/cbBinary uint32, /*in*/dwFlags CRYPT_STRING, /*out*//*optional*/pszString PSTR, /*in*//*out*/pcchString *uint32) (BOOL) = crypt32.CryptBinaryToStringA
//sys	CryptBinaryToStringW(/*in*/pbBinary *uint8, /*in*/cbBinary uint32, /*in*/dwFlags CRYPT_STRING, /*out*//*optional*/pszString PWSTR, /*in*//*out*/pcchString *uint32) (BOOL) = crypt32.CryptBinaryToStringW
//sys	PFXImportCertStore(/*in*/pPFX *CRYPTOAPI_BLOB, /*in*/szPassword PWSTR, /*in*/dwFlags CRYPT_KEY_FLAGS) (HCERTSTORE) = crypt32.PFXImportCertStore
//sys	PFXIsPFXBlob(/*in*/pPFX *CRYPTOAPI_BLOB) (BOOL) = crypt32.PFXIsPFXBlob
//sys	PFXVerifyPassword(/*in*/pPFX *CRYPTOAPI_BLOB, /*in*/szPassword PWSTR, /*in*/dwFlags uint32) (BOOL) = crypt32.PFXVerifyPassword
//sys	PFXExportCertStoreEx(/*in*/hStore HCERTSTORE, /*in*//*out*/pPFX *CRYPTOAPI_BLOB, /*in*/szPassword PWSTR, /*in*/pvPara unsafe.Pointer, /*in*/dwFlags uint32) (BOOL) = crypt32.PFXExportCertStoreEx
//sys	PFXExportCertStore(/*in*/hStore HCERTSTORE, /*in*//*out*/pPFX *CRYPTOAPI_BLOB, /*in*/szPassword PWSTR, /*in*/dwFlags uint32) (BOOL) = crypt32.PFXExportCertStore
//sys	CertOpenServerOcspResponse(/*in*/pChainContext *CERT_CHAIN_CONTEXT, /*in*/dwFlags uint32, /*in*//*optional*/pOpenPara *CERT_SERVER_OCSP_RESPONSE_OPEN_PARA) (unsafe.Pointer) = crypt32.CertOpenServerOcspResponse
//sys	CertAddRefServerOcspResponse(/*in*//*optional*/hServerOcspResponse unsafe.Pointer) = crypt32.CertAddRefServerOcspResponse
//sys	CertCloseServerOcspResponse(/*in*//*optional*/hServerOcspResponse unsafe.Pointer, /*in*/dwFlags uint32) = crypt32.CertCloseServerOcspResponse
//sys	CertGetServerOcspResponseContext(/*in*/hServerOcspResponse unsafe.Pointer, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer) (*CERT_SERVER_OCSP_RESPONSE_CONTEXT) = crypt32.CertGetServerOcspResponseContext
//sys	CertAddRefServerOcspResponseContext(/*in*//*optional*/pServerOcspResponseContext *CERT_SERVER_OCSP_RESPONSE_CONTEXT) = crypt32.CertAddRefServerOcspResponseContext
//sys	CertFreeServerOcspResponseContext(/*in*//*optional*/pServerOcspResponseContext *CERT_SERVER_OCSP_RESPONSE_CONTEXT) = crypt32.CertFreeServerOcspResponseContext
//sys	CertRetrieveLogoOrBiometricInfo(/*in*/pCertContext *CERT_CONTEXT, /*in*/lpszLogoOrBiometricType PSTR, /*in*/dwRetrievalFlags uint32, /*in*/dwTimeout uint32, /*in*/dwFlags uint32, /*in*//*out*/pvReserved unsafe.Pointer, /*out*/ppbData **uint8, /*out*/pcbData *uint32, /*out*//*optional*/ppwszMimeType *PWSTR) (BOOL) = crypt32.CertRetrieveLogoOrBiometricInfo
//sys	CertSelectCertificateChains(/*in*//*optional*/pSelectionContext *Guid, /*in*/dwFlags uint32, /*in*//*optional*/pChainParameters *CERT_SELECT_CHAIN_PARA, /*in*/cCriteria uint32, /*in*//*optional*/rgpCriteria *CERT_SELECT_CRITERIA, /*in*/hStore HCERTSTORE, /*out*/pcSelection *uint32, /*out*/pprgpSelection ***CERT_CHAIN_CONTEXT) (BOOL) = crypt32.CertSelectCertificateChains
//sys	CertFreeCertificateChainList(/*in*/prgpSelection **CERT_CHAIN_CONTEXT) = crypt32.CertFreeCertificateChainList
//sys	CryptRetrieveTimeStamp(/*in*/wszUrl PWSTR, /*in*/dwRetrievalFlags uint32, /*in*/dwTimeout uint32, /*in*/pszHashId PSTR, /*in*//*optional*/pPara *CRYPT_TIMESTAMP_PARA, /*in*/pbData *uint8, /*in*/cbData uint32, /*out*/ppTsContext **CRYPT_TIMESTAMP_CONTEXT, /*out*//*optional*/ppTsSigner **CERT_CONTEXT, /*out*//*optional*/phStore *HCERTSTORE) (BOOL) = crypt32.CryptRetrieveTimeStamp
//sys	CryptVerifyTimeStampSignature(/*in*/pbTSContentInfo *uint8, /*in*/cbTSContentInfo uint32, /*in*//*optional*/pbData *uint8, /*in*/cbData uint32, /*in*//*optional*/hAdditionalStore HCERTSTORE, /*out*/ppTsContext **CRYPT_TIMESTAMP_CONTEXT, /*out*//*optional*/ppTsSigner **CERT_CONTEXT, /*out*//*optional*/phStore *HCERTSTORE) (BOOL) = crypt32.CryptVerifyTimeStampSignature
//sys	CertIsWeakHash(/*in*/dwHashUseType uint32, /*in*/pwszCNGHashAlgid PWSTR, /*in*/dwChainFlags uint32, /*in*//*optional*/pSignerChainContext *CERT_CHAIN_CONTEXT, /*in*//*optional*/pTimeStamp *FILETIME, /*in*//*optional*/pwszFileName PWSTR) (BOOL) = crypt32.CertIsWeakHash
//sys	CryptProtectData(/*in*/pDataIn *CRYPTOAPI_BLOB, /*in*//*optional*/szDataDescr PWSTR, /*in*//*optional*/pOptionalEntropy *CRYPTOAPI_BLOB, /*in*//*out*/pvReserved unsafe.Pointer, /*in*//*optional*/pPromptStruct *CRYPTPROTECT_PROMPTSTRUCT, /*in*/dwFlags uint32, /*out*/pDataOut *CRYPTOAPI_BLOB) (BOOL) = crypt32.CryptProtectData
//sys	CryptUnprotectData(/*in*/pDataIn *CRYPTOAPI_BLOB, /*out*//*optional*/ppszDataDescr *PWSTR, /*in*//*optional*/pOptionalEntropy *CRYPTOAPI_BLOB, /*in*//*out*/pvReserved unsafe.Pointer, /*in*//*optional*/pPromptStruct *CRYPTPROTECT_PROMPTSTRUCT, /*in*/dwFlags uint32, /*out*/pDataOut *CRYPTOAPI_BLOB) (BOOL) = crypt32.CryptUnprotectData
//sys	CryptUpdateProtectedState(/*in*//*optional*/pOldSid PSID, /*in*//*optional*/pwszOldPassword PWSTR, /*in*/dwFlags uint32, /*out*//*optional*/pdwSuccessCount *uint32, /*out*//*optional*/pdwFailureCount *uint32) (BOOL) = crypt32.CryptUpdateProtectedState
//sys	CryptProtectMemory(/*in*//*out*/pDataIn unsafe.Pointer, /*in*/cbDataIn uint32, /*in*/dwFlags uint32) (BOOL) = crypt32.CryptProtectMemory
//sys	CryptUnprotectMemory(/*in*//*out*/pDataIn unsafe.Pointer, /*in*/cbDataIn uint32, /*in*/dwFlags uint32) (BOOL) = crypt32.CryptUnprotectMemory
//sys	NCryptRegisterProtectionDescriptorName(/*in*/pwszName PWSTR, /*in*//*optional*/pwszDescriptorString PWSTR, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptRegisterProtectionDescriptorName
//sys	NCryptQueryProtectionDescriptorName(/*in*/pwszName PWSTR, /*out*//*optional*/pwszDescriptorString PWSTR, /*in*//*out*/pcDescriptorString *uintptr, /*in*/dwFlags uint32) (HRESULT) = ncrypt.NCryptQueryProtectionDescriptorName
//sys	NCryptCreateProtectionDescriptor(/*in*/pwszDescriptorString PWSTR, /*in*/dwFlags uint32, /*out*/phDescriptor *NCRYPT_DESCRIPTOR_HANDLE) (HRESULT) = ncrypt.NCryptCreateProtectionDescriptor
//sys	NCryptCloseProtectionDescriptor(/*in*/hDescriptor NCRYPT_DESCRIPTOR_HANDLE) (HRESULT) = ncrypt.NCryptCloseProtectionDescriptor
//sys	NCryptGetProtectionDescriptorInfo(/*in*/hDescriptor NCRYPT_DESCRIPTOR_HANDLE, /*in*//*optional*/pMemPara *NCRYPT_ALLOC_PARA, /*in*/dwInfoType uint32, /*out*/ppvInfo *unsafe.Pointer) (HRESULT) = ncrypt.NCryptGetProtectionDescriptorInfo
//sys	NCryptProtectSecret(/*in*/hDescriptor NCRYPT_DESCRIPTOR_HANDLE, /*in*/dwFlags uint32, /*in*/pbData *uint8, /*in*/cbData uint32, /*in*//*optional*/pMemPara *NCRYPT_ALLOC_PARA, /*in*//*optional*/hWnd HWND, /*out*/ppbProtectedBlob **uint8, /*out*/pcbProtectedBlob *uint32) (HRESULT) = ncrypt.NCryptProtectSecret
//sys	NCryptUnprotectSecret(/*out*//*optional*/phDescriptor *NCRYPT_DESCRIPTOR_HANDLE, /*in*/dwFlags NCRYPT_FLAGS, /*in*/pbProtectedBlob *uint8, /*in*/cbProtectedBlob uint32, /*in*//*optional*/pMemPara *NCRYPT_ALLOC_PARA, /*in*//*optional*/hWnd HWND, /*out*/ppbData **uint8, /*out*/pcbData *uint32) (HRESULT) = ncrypt.NCryptUnprotectSecret
//sys	NCryptStreamOpenToProtect(/*in*/hDescriptor NCRYPT_DESCRIPTOR_HANDLE, /*in*/dwFlags uint32, /*in*//*optional*/hWnd HWND, /*in*/pStreamInfo *NCRYPT_PROTECT_STREAM_INFO, /*out*/phStream *NCRYPT_STREAM_HANDLE) (HRESULT) = ncrypt.NCryptStreamOpenToProtect
//sys	NCryptStreamOpenToUnprotect(/*in*/pStreamInfo *NCRYPT_PROTECT_STREAM_INFO, /*in*/dwFlags uint32, /*in*//*optional*/hWnd HWND, /*out*/phStream *NCRYPT_STREAM_HANDLE) (HRESULT) = ncrypt.NCryptStreamOpenToUnprotect
//sys	NCryptStreamOpenToUnprotectEx(/*in*/pStreamInfo *NCRYPT_PROTECT_STREAM_INFO_EX, /*in*/dwFlags uint32, /*in*//*optional*/hWnd HWND, /*out*/phStream *NCRYPT_STREAM_HANDLE) (HRESULT) = ncrypt.NCryptStreamOpenToUnprotectEx
//sys	NCryptStreamUpdate(/*in*/hStream NCRYPT_STREAM_HANDLE, /*in*/pbData *uint8, /*in*/cbData uintptr, /*in*/fFinal BOOL) (HRESULT) = ncrypt.NCryptStreamUpdate
//sys	NCryptStreamClose(/*in*/hStream NCRYPT_STREAM_HANDLE) (HRESULT) = ncrypt.NCryptStreamClose
//sys	CryptXmlClose(/*in*/hCryptXml unsafe.Pointer) (HRESULT) = cryptxml.CryptXmlClose
//sys	CryptXmlGetTransforms(/*out*/ppConfig **CRYPT_XML_TRANSFORM_CHAIN_CONFIG) (HRESULT) = cryptxml.CryptXmlGetTransforms
//sys	CryptXmlOpenToEncode(/*in*//*optional*/pConfig *CRYPT_XML_TRANSFORM_CHAIN_CONFIG, /*in*/dwFlags CRYPT_XML_FLAGS, /*in*//*optional*/wszId PWSTR, /*in*//*optional*/rgProperty *CRYPT_XML_PROPERTY, /*in*/cProperty uint32, /*in*//*optional*/pEncoded *CRYPT_XML_BLOB, /*out*/phSignature *unsafe.Pointer) (HRESULT) = cryptxml.CryptXmlOpenToEncode
//sys	CryptXmlOpenToDecode(/*in*//*optional*/pConfig *CRYPT_XML_TRANSFORM_CHAIN_CONFIG, /*in*/dwFlags CRYPT_XML_FLAGS, /*in*//*optional*/rgProperty *CRYPT_XML_PROPERTY, /*in*/cProperty uint32, /*in*/pEncoded *CRYPT_XML_BLOB, /*out*/phCryptXml *unsafe.Pointer) (HRESULT) = cryptxml.CryptXmlOpenToDecode
//sys	CryptXmlAddObject(/*in*/hSignatureOrObject unsafe.Pointer, /*in*/dwFlags uint32, /*in*//*optional*/rgProperty *CRYPT_XML_PROPERTY, /*in*/cProperty uint32, /*in*/pEncoded *CRYPT_XML_BLOB, /*out*//*optional*/ppObject **CRYPT_XML_OBJECT) (HRESULT) = cryptxml.CryptXmlAddObject
//sys	CryptXmlCreateReference(/*in*/hCryptXml unsafe.Pointer, /*in*/dwFlags uint32, /*in*//*optional*/wszId PWSTR, /*in*//*optional*/wszURI PWSTR, /*in*//*optional*/wszType PWSTR, /*in*/pDigestMethod *CRYPT_XML_ALGORITHM, /*in*/cTransform uint32, /*in*//*optional*/rgTransform *CRYPT_XML_ALGORITHM, /*out*/phReference *unsafe.Pointer) (HRESULT) = cryptxml.CryptXmlCreateReference
//sys	CryptXmlDigestReference(/*in*/hReference unsafe.Pointer, /*in*/dwFlags uint32, /*in*/pDataProviderIn *CRYPT_XML_DATA_PROVIDER) (HRESULT) = cryptxml.CryptXmlDigestReference
//sys	CryptXmlSetHMACSecret(/*in*/hSignature unsafe.Pointer, /*in*/pbSecret *uint8, /*in*/cbSecret uint32) (HRESULT) = cryptxml.CryptXmlSetHMACSecret
//sys	CryptXmlSign(/*in*/hSignature unsafe.Pointer, /*in*//*optional*/hKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, /*in*/dwKeySpec CERT_KEY_SPEC, /*in*/dwFlags CRYPT_XML_FLAGS, /*in*/dwKeyInfoSpec CRYPT_XML_KEYINFO_SPEC, /*in*//*optional*/pvKeyInfoSpec unsafe.Pointer, /*in*/pSignatureMethod *CRYPT_XML_ALGORITHM, /*in*/pCanonicalization *CRYPT_XML_ALGORITHM) (HRESULT) = cryptxml.CryptXmlSign
//sys	CryptXmlImportPublicKey(/*in*/dwFlags CRYPT_XML_FLAGS, /*in*/pKeyValue *CRYPT_XML_KEY_VALUE, /*out*/phKey *BCRYPT_KEY_HANDLE) (HRESULT) = cryptxml.CryptXmlImportPublicKey
//sys	CryptXmlVerifySignature(/*in*/hSignature unsafe.Pointer, /*in*//*optional*/hKey BCRYPT_KEY_HANDLE, /*in*/dwFlags CRYPT_XML_FLAGS) (HRESULT) = cryptxml.CryptXmlVerifySignature
//sys	CryptXmlGetDocContext(/*in*/hCryptXml unsafe.Pointer, /*out*/ppStruct **CRYPT_XML_DOC_CTXT) (HRESULT) = cryptxml.CryptXmlGetDocContext
//sys	CryptXmlGetSignature(/*in*/hCryptXml unsafe.Pointer, /*out*/ppStruct **CRYPT_XML_SIGNATURE) (HRESULT) = cryptxml.CryptXmlGetSignature
//sys	CryptXmlGetReference(/*in*/hCryptXml unsafe.Pointer, /*out*/ppStruct **CRYPT_XML_REFERENCE) (HRESULT) = cryptxml.CryptXmlGetReference
//sys	CryptXmlGetStatus(/*in*/hCryptXml unsafe.Pointer, /*out*/pStatus *CRYPT_XML_STATUS) (HRESULT) = cryptxml.CryptXmlGetStatus
//sys	CryptXmlEncode(/*in*/hCryptXml unsafe.Pointer, /*in*/dwCharset CRYPT_XML_CHARSET, /*in*//*optional*/rgProperty *CRYPT_XML_PROPERTY, /*in*/cProperty uint32, /*in*//*out*/pvCallbackState unsafe.Pointer, /*in*/pfnWrite PFN_CRYPT_XML_WRITE_CALLBACK) (HRESULT) = cryptxml.CryptXmlEncode
//sys	CryptXmlGetAlgorithmInfo(/*in*/pXmlAlgorithm *CRYPT_XML_ALGORITHM, /*in*/dwFlags CRYPT_XML_FLAGS, /*out*/ppAlgInfo **CRYPT_XML_ALGORITHM_INFO) (HRESULT) = cryptxml.CryptXmlGetAlgorithmInfo
//sys	CryptXmlFindAlgorithmInfo(/*in*/dwFindByType uint32, /*in*/pvFindBy unsafe.Pointer, /*in*/dwGroupId uint32, /*in*/dwFlags uint32) (*CRYPT_XML_ALGORITHM_INFO) = cryptxml.CryptXmlFindAlgorithmInfo
//sys	CryptXmlEnumAlgorithmInfo(/*in*/dwGroupId uint32, /*in*/dwFlags uint32, /*in*//*out*//*optional*/pvArg unsafe.Pointer, /*in*/pfnEnumAlgInfo PFN_CRYPT_XML_ENUM_ALG_INFO) (HRESULT) = cryptxml.CryptXmlEnumAlgorithmInfo
//sys	GetToken(/*in*/cPolicyChain uint32, /*in*/pPolicyChain *POLICY_ELEMENT, /*out*/securityToken **GENERIC_XML_TOKEN, /*out*/phProofTokenCrypto **INFORMATIONCARD_CRYPTO_HANDLE) (HRESULT) = infocardapi.GetToken
//sys	ManageCardSpace() (HRESULT) = infocardapi.ManageCardSpace
//sys	ImportInformationCard(/*in*/fileName PWSTR) (HRESULT) = infocardapi.ImportInformationCard
//sys	Encrypt(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/fOAEP BOOL, /*in*/cbInData uint32, /*in*/pInData *uint8, /*out*/pcbOutData *uint32, /*out*/ppOutData **uint8) (HRESULT) = infocardapi.Encrypt
//sys	Decrypt(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/fOAEP BOOL, /*in*/cbInData uint32, /*in*/pInData *uint8, /*out*/pcbOutData *uint32, /*out*/ppOutData **uint8) (HRESULT) = infocardapi.Decrypt
//sys	SignHash(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/cbHash uint32, /*in*/pHash *uint8, /*in*/hashAlgOid PWSTR, /*out*/pcbSig *uint32, /*out*/ppSig **uint8) (HRESULT) = infocardapi.SignHash
//sys	VerifyHash(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/cbHash uint32, /*in*/pHash *uint8, /*in*/hashAlgOid PWSTR, /*in*/cbSig uint32, /*in*/pSig *uint8, /*out*/pfVerified *BOOL) (HRESULT) = infocardapi.VerifyHash
//sys	GetCryptoTransform(/*in*/hSymmetricCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/mode uint32, /*in*/padding PaddingMode, /*in*/feedbackSize uint32, /*in*/direction Direction, /*in*/cbIV uint32, /*in*/pIV *uint8, /*out*/pphTransform **INFORMATIONCARD_CRYPTO_HANDLE) (HRESULT) = infocardapi.GetCryptoTransform
//sys	GetKeyedHash(/*in*/hSymmetricCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*out*/pphHash **INFORMATIONCARD_CRYPTO_HANDLE) (HRESULT) = infocardapi.GetKeyedHash
//sys	TransformBlock(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/cbInData uint32, /*in*/pInData *uint8, /*out*/pcbOutData *uint32, /*out*/ppOutData **uint8) (HRESULT) = infocardapi.TransformBlock
//sys	TransformFinalBlock(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/cbInData uint32, /*in*/pInData *uint8, /*out*/pcbOutData *uint32, /*out*/ppOutData **uint8) (HRESULT) = infocardapi.TransformFinalBlock
//sys	HashCore(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/cbInData uint32, /*in*/pInData *uint8) (HRESULT) = infocardapi.HashCore
//sys	HashFinal(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/cbInData uint32, /*in*/pInData *uint8, /*out*/pcbOutData *uint32, /*out*/ppOutData **uint8) (HRESULT) = infocardapi.HashFinal
//sys	FreeToken(/*in*/pAllocMemory *GENERIC_XML_TOKEN) (BOOL) = infocardapi.FreeToken
//sys	CloseCryptoHandle(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE) (HRESULT) = infocardapi.CloseCryptoHandle
//sys	GenerateDerivedKey(/*in*/hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, /*in*/cbLabel uint32, /*in*/pLabel *uint8, /*in*/cbNonce uint32, /*in*/pNonce *uint8, /*in*/derivedKeyLength uint32, /*in*/offset uint32, /*in*/algId PWSTR, /*out*/pcbKey *uint32, /*out*/ppKey **uint8) (HRESULT) = infocardapi.GenerateDerivedKey
//sys	GetBrowserToken(/*in*/dwParamType uint32, /*in*/pParam unsafe.Pointer, /*out*//*optional*/pcbToken *uint32, /*out*//*optional*/ppToken **uint8) (HRESULT) = infocardapi.GetBrowserToken