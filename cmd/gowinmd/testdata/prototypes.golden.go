

// APIs for Windows.Win32.Security.Cryptography
//sys	SystemPrng(pbRandomData *uint8, cbRandomData uintptr) (r BOOL) = bcryptprimitives.SystemPrng
//sys	ProcessPrng(pbData *uint8, cbData uintptr) (r BOOL) = bcryptprimitives.ProcessPrng
//sys	CryptAcquireContextA(phProv *uintptr, szContainer *PSTRElement, szProvider *PSTRElement, dwProvType uint32, dwFlags uint32) (r BOOL, err error) = advapi32.CryptAcquireContextA
//sys	CryptAcquireContextW(phProv *uintptr, szContainer *PWSTRElement, szProvider *PWSTRElement, dwProvType uint32, dwFlags uint32) (r BOOL, err error) = advapi32.CryptAcquireContextW
//sys	CryptReleaseContext(hProv uintptr, dwFlags uint32) (r BOOL, err error) = advapi32.CryptReleaseContext
//sys	CryptGenKey(hProv uintptr, Algid ALG_ID, dwFlags CRYPT_KEY_FLAGS, phKey *uintptr) (r BOOL, err error) = advapi32.CryptGenKey
//sys	CryptDeriveKey(hProv uintptr, Algid ALG_ID, hBaseData uintptr, dwFlags uint32, phKey *uintptr) (r BOOL, err error) = advapi32.CryptDeriveKey
//sys	CryptDestroyKey(hKey uintptr) (r BOOL, err error) = advapi32.CryptDestroyKey
//sys	CryptSetKeyParam(hKey uintptr, dwParam CRYPT_KEY_PARAM_ID, pbData *uint8, dwFlags uint32) (r BOOL, err error) = advapi32.CryptSetKeyParam
//sys	CryptGetKeyParam(hKey uintptr, dwParam CRYPT_KEY_PARAM_ID, pbData *uint8, pdwDataLen *uint32, dwFlags uint32) (r BOOL, err error) = advapi32.CryptGetKeyParam
//sys	CryptSetHashParam(hHash uintptr, dwParam CRYPT_SET_HASH_PARAM, pbData *uint8, dwFlags uint32) (r BOOL, err error) = advapi32.CryptSetHashParam
//sys	CryptGetHashParam(hHash uintptr, dwParam uint32, pbData *uint8, pdwDataLen *uint32, dwFlags uint32) (r BOOL, err error) = advapi32.CryptGetHashParam
//sys	CryptSetProvParam(hProv uintptr, dwParam CRYPT_SET_PROV_PARAM_ID, pbData *uint8, dwFlags uint32) (r BOOL, err error) = advapi32.CryptSetProvParam
//sys	CryptGetProvParam(hProv uintptr, dwParam uint32, pbData *uint8, pdwDataLen *uint32, dwFlags uint32) (r BOOL, err error) = advapi32.CryptGetProvParam
//sys	CryptGenRandom(hProv uintptr, dwLen uint32, pbBuffer *uint8) (r BOOL, err error) = advapi32.CryptGenRandom
//sys	CryptGetUserKey(hProv uintptr, dwKeySpec uint32, phUserKey *uintptr) (r BOOL, err error) = advapi32.CryptGetUserKey
//sys	CryptExportKey(hKey uintptr, hExpKey uintptr, dwBlobType uint32, dwFlags CRYPT_KEY_FLAGS, pbData *uint8, pdwDataLen *uint32) (r BOOL, err error) = advapi32.CryptExportKey
//sys	CryptImportKey(hProv uintptr, pbData *uint8, dwDataLen uint32, hPubKey uintptr, dwFlags CRYPT_KEY_FLAGS, phKey *uintptr) (r BOOL, err error) = advapi32.CryptImportKey
//sys	CryptEncrypt(hKey uintptr, hHash uintptr, Final BOOL, dwFlags uint32, pbData *uint8, pdwDataLen *uint32, dwBufLen uint32) (r BOOL, err error) = advapi32.CryptEncrypt
//sys	CryptDecrypt(hKey uintptr, hHash uintptr, Final BOOL, dwFlags uint32, pbData *uint8, pdwDataLen *uint32) (r BOOL, err error) = advapi32.CryptDecrypt
//sys	CryptCreateHash(hProv uintptr, Algid ALG_ID, hKey uintptr, dwFlags uint32, phHash *uintptr) (r BOOL, err error) = advapi32.CryptCreateHash
//sys	CryptHashData(hHash uintptr, pbData *uint8, dwDataLen uint32, dwFlags uint32) (r BOOL, err error) = advapi32.CryptHashData
//sys	CryptHashSessionKey(hHash uintptr, hKey uintptr, dwFlags uint32) (r BOOL, err error) = advapi32.CryptHashSessionKey
//sys	CryptDestroyHash(hHash uintptr) (r BOOL, err error) = advapi32.CryptDestroyHash
//sys	CryptSignHashA(hHash uintptr, dwKeySpec uint32, szDescription *PSTRElement, dwFlags uint32, pbSignature *uint8, pdwSigLen *uint32) (r BOOL, err error) = advapi32.CryptSignHashA
//sys	CryptSignHashW(hHash uintptr, dwKeySpec uint32, szDescription *PWSTRElement, dwFlags uint32, pbSignature *uint8, pdwSigLen *uint32) (r BOOL, err error) = advapi32.CryptSignHashW
//sys	CryptVerifySignatureA(hHash uintptr, pbSignature *uint8, dwSigLen uint32, hPubKey uintptr, szDescription *PSTRElement, dwFlags uint32) (r BOOL, err error) = advapi32.CryptVerifySignatureA
//sys	CryptVerifySignatureW(hHash uintptr, pbSignature *uint8, dwSigLen uint32, hPubKey uintptr, szDescription *PWSTRElement, dwFlags uint32) (r BOOL, err error) = advapi32.CryptVerifySignatureW
//sys	CryptSetProviderA(pszProvName *PSTRElement, dwProvType uint32) (r BOOL, err error) = advapi32.CryptSetProviderA
//sys	CryptSetProviderW(pszProvName *PWSTRElement, dwProvType uint32) (r BOOL, err error) = advapi32.CryptSetProviderW
//sys	CryptSetProviderExA(pszProvName *PSTRElement, dwProvType uint32, pdwReserved *uint32, dwFlags uint32) (r BOOL, err error) = advapi32.CryptSetProviderExA
//sys	CryptSetProviderExW(pszProvName *PWSTRElement, dwProvType uint32, pdwReserved *uint32, dwFlags uint32) (r BOOL, err error) = advapi32.CryptSetProviderExW
//sys	CryptGetDefaultProviderA(dwProvType uint32, pdwReserved *uint32, dwFlags uint32, pszProvName *PSTRElement, pcbProvName *uint32) (r BOOL, err error) = advapi32.CryptGetDefaultProviderA
//sys	CryptGetDefaultProviderW(dwProvType uint32, pdwReserved *uint32, dwFlags uint32, pszProvName *PWSTRElement, pcbProvName *uint32) (r BOOL, err error) = advapi32.CryptGetDefaultProviderW
//sys	CryptEnumProviderTypesA(dwIndex uint32, pdwReserved *uint32, dwFlags uint32, pdwProvType *uint32, szTypeName *PSTRElement, pcbTypeName *uint32) (r BOOL, err error) = advapi32.CryptEnumProviderTypesA
//sys	CryptEnumProviderTypesW(dwIndex uint32, pdwReserved *uint32, dwFlags uint32, pdwProvType *uint32, szTypeName *PWSTRElement, pcbTypeName *uint32) (r BOOL, err error) = advapi32.CryptEnumProviderTypesW
//sys	CryptEnumProvidersA(dwIndex uint32, pdwReserved *uint32, dwFlags uint32, pdwProvType *uint32, szProvName *PSTRElement, pcbProvName *uint32) (r BOOL, err error) = advapi32.CryptEnumProvidersA
//sys	CryptEnumProvidersW(dwIndex uint32, pdwReserved *uint32, dwFlags uint32, pdwProvType *uint32, szProvName *PWSTRElement, pcbProvName *uint32) (r BOOL, err error) = advapi32.CryptEnumProvidersW
//sys	CryptContextAddRef(hProv uintptr, pdwReserved *uint32, dwFlags uint32) (r BOOL, err error) = advapi32.CryptContextAddRef
//sys	CryptDuplicateKey(hKey uintptr, pdwReserved *uint32, dwFlags uint32, phKey *uintptr) (r BOOL, err error) = advapi32.CryptDuplicateKey
//sys	CryptDuplicateHash(hHash uintptr, pdwReserved *uint32, dwFlags uint32, phHash *uintptr) (r BOOL, err error) = advapi32.CryptDuplicateHash
//sys	BCryptOpenAlgorithmProvider(phAlgorithm *BCRYPT_ALG_HANDLE, pszAlgId *PWSTRElement, pszImplementation *PWSTRElement, dwFlags BCRYPT_OPEN_ALGORITHM_PROVIDER_FLAGS) (r NTSTATUS) = bcrypt.BCryptOpenAlgorithmProvider
//sys	BCryptEnumAlgorithms(dwAlgOperations BCRYPT_OPERATION, pAlgCount *uint32, ppAlgList **BCRYPT_ALGORITHM_IDENTIFIER, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptEnumAlgorithms
//sys	BCryptEnumProviders(pszAlgId *PWSTRElement, pImplCount *uint32, ppImplList **BCRYPT_PROVIDER_NAME, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptEnumProviders
//sys	BCryptGetProperty(hObject BCRYPT_HANDLE, pszProperty *PWSTRElement, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptGetProperty
//sys	BCryptSetProperty(hObject BCRYPT_HANDLE, pszProperty *PWSTRElement, pbInput *uint8, cbInput uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptSetProperty
//sys	BCryptCloseAlgorithmProvider(hAlgorithm BCRYPT_ALG_HANDLE, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptCloseAlgorithmProvider
//sys	BCryptFreeBuffer(pvBuffer unsafe.Pointer) = bcrypt.BCryptFreeBuffer
//sys	BCryptGenerateSymmetricKey(hAlgorithm BCRYPT_ALG_HANDLE, phKey *BCRYPT_KEY_HANDLE, pbKeyObject *uint8, cbKeyObject uint32, pbSecret *uint8, cbSecret uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptGenerateSymmetricKey
//sys	BCryptGenerateKeyPair(hAlgorithm BCRYPT_ALG_HANDLE, phKey *BCRYPT_KEY_HANDLE, dwLength uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptGenerateKeyPair
//sys	BCryptEncrypt(hKey BCRYPT_KEY_HANDLE, pbInput *uint8, cbInput uint32, pPaddingInfo unsafe.Pointer, pbIV *uint8, cbIV uint32, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags BCRYPT_FLAGS) (r NTSTATUS) = bcrypt.BCryptEncrypt
//sys	BCryptDecrypt(hKey BCRYPT_KEY_HANDLE, pbInput *uint8, cbInput uint32, pPaddingInfo unsafe.Pointer, pbIV *uint8, cbIV uint32, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags BCRYPT_FLAGS) (r NTSTATUS) = bcrypt.BCryptDecrypt
//sys	BCryptExportKey(hKey BCRYPT_KEY_HANDLE, hExportKey BCRYPT_KEY_HANDLE, pszBlobType *PWSTRElement, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptExportKey
//sys	BCryptImportKey(hAlgorithm BCRYPT_ALG_HANDLE, hImportKey BCRYPT_KEY_HANDLE, pszBlobType *PWSTRElement, phKey *BCRYPT_KEY_HANDLE, pbKeyObject *uint8, cbKeyObject uint32, pbInput *uint8, cbInput uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptImportKey
//sys	BCryptImportKeyPair(hAlgorithm BCRYPT_ALG_HANDLE, hImportKey BCRYPT_KEY_HANDLE, pszBlobType *PWSTRElement, phKey *BCRYPT_KEY_HANDLE, pbInput *uint8, cbInput uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptImportKeyPair
//sys	BCryptDuplicateKey(hKey BCRYPT_KEY_HANDLE, phNewKey *BCRYPT_KEY_HANDLE, pbKeyObject *uint8, cbKeyObject uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptDuplicateKey
//sys	BCryptFinalizeKeyPair(hKey BCRYPT_KEY_HANDLE, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptFinalizeKeyPair
//sys	BCryptDestroyKey(hKey BCRYPT_KEY_HANDLE) (r NTSTATUS) = bcrypt.BCryptDestroyKey
//sys	BCryptDestroySecret(hSecret BCRYPT_SECRET_HANDLE) (r NTSTATUS) = bcrypt.BCryptDestroySecret
//sys	BCryptSignHash(hKey BCRYPT_KEY_HANDLE, pPaddingInfo unsafe.Pointer, pbInput *uint8, cbInput uint32, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags BCRYPT_FLAGS) (r NTSTATUS) = bcrypt.BCryptSignHash
//sys	BCryptVerifySignature(hKey BCRYPT_KEY_HANDLE, pPaddingInfo unsafe.Pointer, pbHash *uint8, cbHash uint32, pbSignature *uint8, cbSignature uint32, dwFlags BCRYPT_FLAGS) (r NTSTATUS) = bcrypt.BCryptVerifySignature
//sys	BCryptSecretAgreement(hPrivKey BCRYPT_KEY_HANDLE, hPubKey BCRYPT_KEY_HANDLE, phAgreedSecret *BCRYPT_SECRET_HANDLE, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptSecretAgreement
//sys	BCryptDeriveKey(hSharedSecret BCRYPT_SECRET_HANDLE, pwszKDF *PWSTRElement, pParameterList *BCryptBufferDesc, pbDerivedKey *uint8, cbDerivedKey uint32, pcbResult *uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptDeriveKey
//sys	BCryptKeyDerivation(hKey BCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, pbDerivedKey *uint8, cbDerivedKey uint32, pcbResult *uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptKeyDerivation
//sys	BCryptCreateHash(hAlgorithm BCRYPT_ALG_HANDLE, phHash *BCRYPT_HASH_HANDLE, pbHashObject *uint8, cbHashObject uint32, pbSecret *uint8, cbSecret uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptCreateHash
//sys	BCryptHashData(hHash BCRYPT_HASH_HANDLE, pbInput *uint8, cbInput uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptHashData
//sys	BCryptFinishHash(hHash BCRYPT_HASH_HANDLE, pbOutput *uint8, cbOutput uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptFinishHash
//sys	BCryptCreateMultiHash(hAlgorithm BCRYPT_ALG_HANDLE, phHash *BCRYPT_HASH_HANDLE, nHashes uint32, pbHashObject *uint8, cbHashObject uint32, pbSecret *uint8, cbSecret uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptCreateMultiHash
//sys	BCryptProcessMultiOperations(hObject BCRYPT_HANDLE, operationType BCRYPT_MULTI_OPERATION_TYPE, pOperations unsafe.Pointer, cbOperations uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptProcessMultiOperations
//sys	BCryptDuplicateHash(hHash BCRYPT_HASH_HANDLE, phNewHash *BCRYPT_HASH_HANDLE, pbHashObject *uint8, cbHashObject uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptDuplicateHash
//sys	BCryptDestroyHash(hHash BCRYPT_HASH_HANDLE) (r NTSTATUS) = bcrypt.BCryptDestroyHash
//sys	BCryptHash(hAlgorithm BCRYPT_ALG_HANDLE, pbSecret *uint8, cbSecret uint32, pbInput *uint8, cbInput uint32, pbOutput *uint8, cbOutput uint32) (r NTSTATUS) = bcrypt.BCryptHash
//sys	BCryptGenRandom(hAlgorithm BCRYPT_ALG_HANDLE, pbBuffer *uint8, cbBuffer uint32, dwFlags BCRYPTGENRANDOM_FLAGS) (r NTSTATUS) = bcrypt.BCryptGenRandom
//sys	BCryptDeriveKeyCapi(hHash BCRYPT_HASH_HANDLE, hTargetAlg BCRYPT_ALG_HANDLE, pbDerivedKey *uint8, cbDerivedKey uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptDeriveKeyCapi
//sys	BCryptDeriveKeyPBKDF2(hPrf BCRYPT_ALG_HANDLE, pbPassword *uint8, cbPassword uint32, pbSalt *uint8, cbSalt uint32, cIterations uint64, pbDerivedKey *uint8, cbDerivedKey uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptDeriveKeyPBKDF2
//sys	BCryptEncapsulate(hKey BCRYPT_KEY_HANDLE, pbSecretKey *uint8, cbSecretKey uint32, pcbSecretKey *uint32, pbCipherText *uint8, cbCipherText uint32, pcbCipherText *uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptEncapsulate
//sys	BCryptDecapsulate(hKey BCRYPT_KEY_HANDLE, pbCipherText *uint8, cbCipherText uint32, pbSecretKey *uint8, cbSecretKey uint32, pcbSecretKey *uint32, dwFlags uint32) (r NTSTATUS) = bcrypt.BCryptDecapsulate
//sys	BCryptQueryProviderRegistration(pszProvider *PWSTRElement, dwMode BCRYPT_QUERY_PROVIDER_MODE, dwInterface BCRYPT_INTERFACE, pcbBuffer *uint32, ppBuffer **CRYPT_PROVIDER_REG) (r NTSTATUS) = bcrypt.BCryptQueryProviderRegistration
//sys	BCryptEnumRegisteredProviders(pcbBuffer *uint32, ppBuffer **CRYPT_PROVIDERS) (r NTSTATUS) = bcrypt.BCryptEnumRegisteredProviders
//sys	BCryptCreateContext(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, pConfig *CRYPT_CONTEXT_CONFIG) (r NTSTATUS) = bcrypt.BCryptCreateContext
//sys	BCryptDeleteContext(dwTable BCRYPT_TABLE, pszContext *PWSTRElement) (r NTSTATUS) = bcrypt.BCryptDeleteContext
//sys	BCryptEnumContexts(dwTable BCRYPT_TABLE, pcbBuffer *uint32, ppBuffer **CRYPT_CONTEXTS) (r NTSTATUS) = bcrypt.BCryptEnumContexts
//sys	BCryptConfigureContext(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, pConfig *CRYPT_CONTEXT_CONFIG) (r NTSTATUS) = bcrypt.BCryptConfigureContext
//sys	BCryptQueryContextConfiguration(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, pcbBuffer *uint32, ppBuffer **CRYPT_CONTEXT_CONFIG) (r NTSTATUS) = bcrypt.BCryptQueryContextConfiguration
//sys	BCryptAddContextFunction(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, dwInterface BCRYPT_INTERFACE, pszFunction *PWSTRElement, dwPosition uint32) (r NTSTATUS) = bcrypt.BCryptAddContextFunction
//sys	BCryptRemoveContextFunction(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, dwInterface BCRYPT_INTERFACE, pszFunction *PWSTRElement) (r NTSTATUS) = bcrypt.BCryptRemoveContextFunction
//sys	BCryptEnumContextFunctions(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, dwInterface BCRYPT_INTERFACE, pcbBuffer *uint32, ppBuffer **CRYPT_CONTEXT_FUNCTIONS) (r NTSTATUS) = bcrypt.BCryptEnumContextFunctions
//sys	BCryptConfigureContextFunction(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, dwInterface BCRYPT_INTERFACE, pszFunction *PWSTRElement, pConfig *CRYPT_CONTEXT_FUNCTION_CONFIG) (r NTSTATUS) = bcrypt.BCryptConfigureContextFunction
//sys	BCryptQueryContextFunctionConfiguration(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, dwInterface BCRYPT_INTERFACE, pszFunction *PWSTRElement, pcbBuffer *uint32, ppBuffer **CRYPT_CONTEXT_FUNCTION_CONFIG) (r NTSTATUS) = bcrypt.BCryptQueryContextFunctionConfiguration
//sys	BCryptEnumContextFunctionProviders(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, dwInterface BCRYPT_INTERFACE, pszFunction *PWSTRElement, pcbBuffer *uint32, ppBuffer **CRYPT_CONTEXT_FUNCTION_PROVIDERS) (r NTSTATUS) = bcrypt.BCryptEnumContextFunctionProviders
//sys	BCryptSetContextFunctionProperty(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, dwInterface BCRYPT_INTERFACE, pszFunction *PWSTRElement, pszProperty *PWSTRElement, cbValue uint32, pbValue *uint8) (r NTSTATUS) = bcrypt.BCryptSetContextFunctionProperty
//sys	BCryptQueryContextFunctionProperty(dwTable BCRYPT_TABLE, pszContext *PWSTRElement, dwInterface BCRYPT_INTERFACE, pszFunction *PWSTRElement, pszProperty *PWSTRElement, pcbValue *uint32, ppbValue **uint8) (r NTSTATUS) = bcrypt.BCryptQueryContextFunctionProperty
//sys	BCryptRegisterConfigChangeNotify(phEvent *HANDLE) (r NTSTATUS) = bcrypt.BCryptRegisterConfigChangeNotify
//sys	BCryptUnregisterConfigChangeNotify(hEvent HANDLE) (r NTSTATUS) = bcrypt.BCryptUnregisterConfigChangeNotify
//sys	BCryptResolveProviders(pszContext *PWSTRElement, dwInterface uint32, pszFunction *PWSTRElement, pszProvider *PWSTRElement, dwMode BCRYPT_QUERY_PROVIDER_MODE, dwFlags BCRYPT_RESOLVE_PROVIDERS_FLAGS, pcbBuffer *uint32, ppBuffer **CRYPT_PROVIDER_REFS) (r NTSTATUS) = bcrypt.BCryptResolveProviders
//sys	BCryptGetFipsAlgorithmMode(pfEnabled *uint8) (r NTSTATUS) = bcrypt.BCryptGetFipsAlgorithmMode
//sys	NCryptOpenStorageProvider(phProvider *NCRYPT_PROV_HANDLE, pszProviderName *PWSTRElement, dwFlags uint32) (r HRESULT) = ncrypt.NCryptOpenStorageProvider
//sys	NCryptEnumAlgorithms(hProvider NCRYPT_PROV_HANDLE, dwAlgOperations NCRYPT_OPERATION, pdwAlgCount *uint32, ppAlgList **NCryptAlgorithmName, dwFlags uint32) (r HRESULT) = ncrypt.NCryptEnumAlgorithms
//sys	NCryptIsAlgSupported(hProvider NCRYPT_PROV_HANDLE, pszAlgId *PWSTRElement, dwFlags uint32) (r HRESULT) = ncrypt.NCryptIsAlgSupported
//sys	NCryptEnumKeys(hProvider NCRYPT_PROV_HANDLE, pszScope *PWSTRElement, ppKeyName **NCryptKeyName, ppEnumState *unsafe.Pointer, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptEnumKeys
//sys	NCryptEnumStorageProviders(pdwProviderCount *uint32, ppProviderList **NCryptProviderName, dwFlags uint32) (r HRESULT) = ncrypt.NCryptEnumStorageProviders
//sys	NCryptFreeBuffer(pvInput unsafe.Pointer) (r HRESULT) = ncrypt.NCryptFreeBuffer
//sys	NCryptOpenKey(hProvider NCRYPT_PROV_HANDLE, phKey *NCRYPT_KEY_HANDLE, pszKeyName *PWSTRElement, dwLegacyKeySpec CERT_KEY_SPEC, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptOpenKey
//sys	NCryptCreatePersistedKey(hProvider NCRYPT_PROV_HANDLE, phKey *NCRYPT_KEY_HANDLE, pszAlgId *PWSTRElement, pszKeyName *PWSTRElement, dwLegacyKeySpec CERT_KEY_SPEC, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptCreatePersistedKey
//sys	NCryptGetProperty(hObject NCRYPT_HANDLE, pszProperty *PWSTRElement, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags OBJECT_SECURITY_INFORMATION) (r HRESULT) = ncrypt.NCryptGetProperty
//sys	NCryptSetProperty(hObject NCRYPT_HANDLE, pszProperty *PWSTRElement, pbInput *uint8, cbInput uint32, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptSetProperty
//sys	NCryptFinalizeKey(hKey NCRYPT_KEY_HANDLE, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptFinalizeKey
//sys	NCryptEncrypt(hKey NCRYPT_KEY_HANDLE, pbInput *uint8, cbInput uint32, pPaddingInfo unsafe.Pointer, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptEncrypt
//sys	NCryptDecrypt(hKey NCRYPT_KEY_HANDLE, pbInput *uint8, cbInput uint32, pPaddingInfo unsafe.Pointer, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptDecrypt
//sys	NCryptEncapsulate(hKey NCRYPT_KEY_HANDLE, pbSecretKey *uint8, cbSecretKey uint32, pcbSecretKey *uint32, pbCipherText *uint8, cbCipherText uint32, pcbCipherText *uint32, dwFlags uint32) (r HRESULT) = ncrypt.NCryptEncapsulate
//sys	NCryptDecapsulate(hKey NCRYPT_KEY_HANDLE, pbCipherText *uint8, cbCipherText uint32, pbSecretKey *uint8, cbSecretKey uint32, pcbSecretKey *uint32, dwFlags uint32) (r HRESULT) = ncrypt.NCryptDecapsulate
//sys	NCryptImportKey(hProvider NCRYPT_PROV_HANDLE, hImportKey NCRYPT_KEY_HANDLE, pszBlobType *PWSTRElement, pParameterList *BCryptBufferDesc, phKey *NCRYPT_KEY_HANDLE, pbData *uint8, cbData uint32, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptImportKey
//sys	NCryptExportKey(hKey NCRYPT_KEY_HANDLE, hExportKey NCRYPT_KEY_HANDLE, pszBlobType *PWSTRElement, pParameterList *BCryptBufferDesc, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptExportKey
//sys	NCryptSignHash(hKey NCRYPT_KEY_HANDLE, pPaddingInfo unsafe.Pointer, pbHashValue *uint8, cbHashValue uint32, pbSignature *uint8, cbSignature uint32, pcbResult *uint32, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptSignHash
//sys	NCryptVerifySignature(hKey NCRYPT_KEY_HANDLE, pPaddingInfo unsafe.Pointer, pbHashValue *uint8, cbHashValue uint32, pbSignature *uint8, cbSignature uint32, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptVerifySignature
//sys	NCryptDeleteKey(hKey NCRYPT_KEY_HANDLE, dwFlags uint32) (r HRESULT) = ncrypt.NCryptDeleteKey
//sys	NCryptFreeObject(hObject NCRYPT_HANDLE) (r HRESULT) = ncrypt.NCryptFreeObject
//sys	NCryptIsKeyHandle(hKey NCRYPT_KEY_HANDLE) (r BOOL) = ncrypt.NCryptIsKeyHandle
//sys	NCryptTranslateHandle(phProvider *NCRYPT_PROV_HANDLE, phKey *NCRYPT_KEY_HANDLE, hLegacyProv uintptr, hLegacyKey uintptr, dwLegacyKeySpec CERT_KEY_SPEC, dwFlags uint32) (r HRESULT) = ncrypt.NCryptTranslateHandle
//sys	NCryptNotifyChangeKey(hProvider NCRYPT_PROV_HANDLE, phEvent *HANDLE, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptNotifyChangeKey
//sys	NCryptSecretAgreement(hPrivKey NCRYPT_KEY_HANDLE, hPubKey NCRYPT_KEY_HANDLE, phAgreedSecret *NCRYPT_SECRET_HANDLE, dwFlags NCRYPT_FLAGS) (r HRESULT) = ncrypt.NCryptSecretAgreement
//sys	NCryptDeriveKey(hSharedSecret NCRYPT_SECRET_HANDLE, pwszKDF *PWSTRElement, pParameterList *BCryptBufferDesc, pbDerivedKey *uint8, cbDerivedKey uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.NCryptDeriveKey
//sys	NCryptKeyDerivation(hKey NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, pbDerivedKey *uint8, cbDerivedKey uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.NCryptKeyDerivation
//sys	NCryptCreateClaim(hSubjectKey NCRYPT_KEY_HANDLE, hAuthorityKey NCRYPT_KEY_HANDLE, dwClaimType uint32, pParameterList *BCryptBufferDesc, pbClaimBlob *uint8, cbClaimBlob uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.NCryptCreateClaim
//sys	NCryptVerifyClaim(hSubjectKey NCRYPT_KEY_HANDLE, hAuthorityKey NCRYPT_KEY_HANDLE, dwClaimType uint32, pParameterList *BCryptBufferDesc, pbClaimBlob *uint8, cbClaimBlob uint32, pOutput *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.NCryptVerifyClaim
//sys	CryptFormatObject(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, dwFormatType uint32, dwFormatStrType uint32, pFormatStruct unsafe.Pointer, lpszStructType *PSTRElement, pbEncoded *uint8, cbEncoded uint32, pbFormat unsafe.Pointer, pcbFormat *uint32) (r BOOL, err error) = crypt32.CryptFormatObject
//sys	CryptEncodeObjectEx(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, lpszStructType *PSTRElement, pvStructInfo unsafe.Pointer, dwFlags CRYPT_ENCODE_OBJECT_FLAGS, pEncodePara *CRYPT_ENCODE_PARA, pvEncoded unsafe.Pointer, pcbEncoded *uint32) (r BOOL, err error) = crypt32.CryptEncodeObjectEx
//sys	CryptEncodeObject(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, lpszStructType *PSTRElement, pvStructInfo unsafe.Pointer, pbEncoded *uint8, pcbEncoded *uint32) (r BOOL, err error) = crypt32.CryptEncodeObject
//sys	CryptDecodeObjectEx(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, lpszStructType *PSTRElement, pbEncoded *uint8, cbEncoded uint32, dwFlags uint32, pDecodePara *CRYPT_DECODE_PARA, pvStructInfo unsafe.Pointer, pcbStructInfo *uint32) (r BOOL, err error) = crypt32.CryptDecodeObjectEx
//sys	CryptDecodeObject(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, lpszStructType *PSTRElement, pbEncoded *uint8, cbEncoded uint32, dwFlags uint32, pvStructInfo unsafe.Pointer, pcbStructInfo *uint32) (r BOOL, err error) = crypt32.CryptDecodeObject
//sys	CryptInstallOIDFunctionAddress(hModule HMODULE, dwEncodingType uint32, pszFuncName *PSTRElement, cFuncEntry uint32, rgFuncEntry *CRYPT_OID_FUNC_ENTRY, dwFlags uint32) (r BOOL) = crypt32.CryptInstallOIDFunctionAddress
//sys	CryptInitOIDFunctionSet(pszFuncName *PSTRElement, dwFlags uint32) (r unsafe.Pointer) = crypt32.CryptInitOIDFunctionSet
//sys	CryptGetOIDFunctionAddress(hFuncSet unsafe.Pointer, dwEncodingType uint32, pszOID *PSTRElement, dwFlags uint32, ppvFuncAddr *unsafe.Pointer, phFuncAddr *unsafe.Pointer) (r BOOL, err error) = crypt32.CryptGetOIDFunctionAddress
//sys	CryptGetDefaultOIDDllList(hFuncSet unsafe.Pointer, dwEncodingType uint32, pwszDllList *PWSTRElement, pcchDllList *uint32) (r BOOL, err error) = crypt32.CryptGetDefaultOIDDllList
//sys	CryptGetDefaultOIDFunctionAddress(hFuncSet unsafe.Pointer, dwEncodingType uint32, pwszDll *PWSTRElement, dwFlags uint32, ppvFuncAddr *unsafe.Pointer, phFuncAddr *unsafe.Pointer) (r BOOL) = crypt32.CryptGetDefaultOIDFunctionAddress
//sys	CryptFreeOIDFunctionAddress(hFuncAddr unsafe.Pointer, dwFlags uint32) (r BOOL) = crypt32.CryptFreeOIDFunctionAddress
//sys	CryptRegisterOIDFunction(dwEncodingType uint32, pszFuncName *PSTRElement, pszOID *PSTRElement, pwszDll *PWSTRElement, pszOverrideFuncName *PSTRElement) (r BOOL) = crypt32.CryptRegisterOIDFunction
//sys	CryptUnregisterOIDFunction(dwEncodingType uint32, pszFuncName *PSTRElement, pszOID *PSTRElement) (r BOOL) = crypt32.CryptUnregisterOIDFunction
//sys	CryptRegisterDefaultOIDFunction(dwEncodingType uint32, pszFuncName *PSTRElement, dwIndex uint32, pwszDll *PWSTRElement) (r BOOL) = crypt32.CryptRegisterDefaultOIDFunction
//sys	CryptUnregisterDefaultOIDFunction(dwEncodingType uint32, pszFuncName *PSTRElement, pwszDll *PWSTRElement) (r BOOL) = crypt32.CryptUnregisterDefaultOIDFunction
//sys	CryptSetOIDFunctionValue(dwEncodingType uint32, pszFuncName *PSTRElement, pszOID *PSTRElement, pwszValueName *PWSTRElement, dwValueType REG_VALUE_TYPE, pbValueData *uint8, cbValueData uint32) (r BOOL) = crypt32.CryptSetOIDFunctionValue
//sys	CryptGetOIDFunctionValue(dwEncodingType uint32, pszFuncName *PSTRElement, pszOID *PSTRElement, pwszValueName *PWSTRElement, pdwValueType *uint32, pbValueData *uint8, pcbValueData *uint32) (r BOOL, err error) = crypt32.CryptGetOIDFunctionValue
//sys	CryptEnumOIDFunction(dwEncodingType uint32, pszFuncName *PSTRElement, pszOID *PSTRElement, dwFlags uint32, pvArg unsafe.Pointer, pfnEnumOIDFunc PFN_CRYPT_ENUM_OID_FUNC) (r BOOL, err error) = crypt32.CryptEnumOIDFunction
//sys	CryptFindOIDInfo(dwKeyType uint32, pvKey unsafe.Pointer, dwGroupId uint32) (r *CRYPT_OID_INFO) = crypt32.CryptFindOIDInfo
//sys	CryptRegisterOIDInfo(pInfo *CRYPT_OID_INFO, dwFlags uint32) (r BOOL) = crypt32.CryptRegisterOIDInfo
//sys	CryptUnregisterOIDInfo(pInfo *CRYPT_OID_INFO) (r BOOL) = crypt32.CryptUnregisterOIDInfo
//sys	CryptEnumOIDInfo(dwGroupId uint32, dwFlags uint32, pvArg unsafe.Pointer, pfnEnumOIDInfo PFN_CRYPT_ENUM_OID_INFO) (r BOOL) = crypt32.CryptEnumOIDInfo
//sys	CryptFindLocalizedName(pwszCryptName *PWSTRElement) (r *PWSTRElement) = crypt32.CryptFindLocalizedName
//sys	CryptMsgOpenToEncode(dwMsgEncodingType uint32, dwFlags uint32, dwMsgType CRYPT_MSG_TYPE, pvMsgEncodeInfo unsafe.Pointer, pszInnerContentObjID *PSTRElement, pStreamInfo *CMSG_STREAM_INFO) (r unsafe.Pointer, err error) = crypt32.CryptMsgOpenToEncode
//sys	CryptMsgCalculateEncodedLength(dwMsgEncodingType uint32, dwFlags uint32, dwMsgType uint32, pvMsgEncodeInfo unsafe.Pointer, pszInnerContentObjID *PSTRElement, cbData uint32) (r uint32, err error) = crypt32.CryptMsgCalculateEncodedLength
//sys	CryptMsgOpenToDecode(dwMsgEncodingType uint32, dwFlags uint32, dwMsgType uint32, hCryptProv HCRYPTPROV_LEGACY, pRecipientInfo *CERT_INFO, pStreamInfo *CMSG_STREAM_INFO) (r unsafe.Pointer, err error) = crypt32.CryptMsgOpenToDecode
//sys	CryptMsgDuplicate(hCryptMsg unsafe.Pointer) (r unsafe.Pointer) = crypt32.CryptMsgDuplicate
//sys	CryptMsgClose(hCryptMsg unsafe.Pointer) (r BOOL, err error) = crypt32.CryptMsgClose
//sys	CryptMsgUpdate(hCryptMsg unsafe.Pointer, pbData *uint8, cbData uint32, fFinal BOOL) (r BOOL, err error) = crypt32.CryptMsgUpdate
//sys	CryptMsgGetParam(hCryptMsg unsafe.Pointer, dwParamType uint32, dwIndex uint32, pvData unsafe.Pointer, pcbData *uint32) (r BOOL, err error) = crypt32.CryptMsgGetParam
//sys	CryptMsgControl(hCryptMsg unsafe.Pointer, dwFlags uint32, dwCtrlType uint32, pvCtrlPara unsafe.Pointer) (r BOOL, err error) = crypt32.CryptMsgControl
//sys	CryptMsgVerifyCountersignatureEncoded(hCryptProv HCRYPTPROV_LEGACY, dwEncodingType uint32, pbSignerInfo *uint8, cbSignerInfo uint32, pbSignerInfoCountersignature *uint8, cbSignerInfoCountersignature uint32, pciCountersigner *CERT_INFO) (r BOOL, err error) = crypt32.CryptMsgVerifyCountersignatureEncoded
//sys	CryptMsgVerifyCountersignatureEncodedEx(hCryptProv HCRYPTPROV_LEGACY, dwEncodingType uint32, pbSignerInfo *uint8, cbSignerInfo uint32, pbSignerInfoCountersignature *uint8, cbSignerInfoCountersignature uint32, dwSignerType uint32, pvSigner unsafe.Pointer, dwFlags uint32, pvExtra unsafe.Pointer) (r BOOL, err error) = crypt32.CryptMsgVerifyCountersignatureEncodedEx
//sys	CryptMsgCountersign(hCryptMsg unsafe.Pointer, dwIndex uint32, cCountersigners uint32, rgCountersigners *CMSG_SIGNER_ENCODE_INFO) (r BOOL, err error) = crypt32.CryptMsgCountersign
//sys	CryptMsgCountersignEncoded(dwEncodingType uint32, pbSignerInfo *uint8, cbSignerInfo uint32, cCountersigners uint32, rgCountersigners *CMSG_SIGNER_ENCODE_INFO, pbCountersignature *uint8, pcbCountersignature *uint32) (r BOOL, err error) = crypt32.CryptMsgCountersignEncoded
//sys	CertOpenStore(lpszStoreProvider *PSTRElement, dwEncodingType CERT_QUERY_ENCODING_TYPE, hCryptProv HCRYPTPROV_LEGACY, dwFlags CERT_OPEN_STORE_FLAGS, pvPara unsafe.Pointer) (r HCERTSTORE, err error) = crypt32.CertOpenStore
//sys	CertDuplicateStore(hCertStore HCERTSTORE) (r HCERTSTORE) = crypt32.CertDuplicateStore
//sys	CertSaveStore(hCertStore HCERTSTORE, dwEncodingType CERT_QUERY_ENCODING_TYPE, dwSaveAs CERT_STORE_SAVE_AS, dwSaveTo CERT_STORE_SAVE_TO, pvSaveToPara unsafe.Pointer, dwFlags uint32) (r BOOL, err error) = crypt32.CertSaveStore
//sys	CertCloseStore(hCertStore HCERTSTORE, dwFlags uint32) (r BOOL, err error) = crypt32.CertCloseStore
//sys	CertGetSubjectCertificateFromStore(hCertStore HCERTSTORE, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pCertId *CERT_INFO) (r *CERT_CONTEXT, err error) = crypt32.CertGetSubjectCertificateFromStore
//sys	CertEnumCertificatesInStore(hCertStore HCERTSTORE, pPrevCertContext *CERT_CONTEXT) (r *CERT_CONTEXT, err error) = crypt32.CertEnumCertificatesInStore
//sys	CertFindCertificateInStore(hCertStore HCERTSTORE, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, dwFindFlags uint32, dwFindType CERT_FIND_FLAGS, pvFindPara unsafe.Pointer, pPrevCertContext *CERT_CONTEXT) (r *CERT_CONTEXT, err error) = crypt32.CertFindCertificateInStore
//sys	CertGetIssuerCertificateFromStore(hCertStore HCERTSTORE, pSubjectContext *CERT_CONTEXT, pPrevIssuerContext *CERT_CONTEXT, pdwFlags *uint32) (r *CERT_CONTEXT, err error) = crypt32.CertGetIssuerCertificateFromStore
//sys	CertVerifySubjectCertificateContext(pSubject *CERT_CONTEXT, pIssuer *CERT_CONTEXT, pdwFlags *uint32) (r BOOL, err error) = crypt32.CertVerifySubjectCertificateContext
//sys	CertDuplicateCertificateContext(pCertContext *CERT_CONTEXT) (r *CERT_CONTEXT) = crypt32.CertDuplicateCertificateContext
//sys	CertCreateCertificateContext(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pbCertEncoded *uint8, cbCertEncoded uint32) (r *CERT_CONTEXT, err error) = crypt32.CertCreateCertificateContext
//sys	CertFreeCertificateContext(pCertContext *CERT_CONTEXT) (r BOOL) = crypt32.CertFreeCertificateContext
//sys	CertSetCertificateContextProperty(pCertContext *CERT_CONTEXT, dwPropId uint32, dwFlags uint32, pvData unsafe.Pointer) (r BOOL, err error) = crypt32.CertSetCertificateContextProperty
//sys	CertGetCertificateContextProperty(pCertContext *CERT_CONTEXT, dwPropId uint32, pvData unsafe.Pointer, pcbData *uint32) (r BOOL, err error) = crypt32.CertGetCertificateContextProperty
//sys	CertEnumCertificateContextProperties(pCertContext *CERT_CONTEXT, dwPropId uint32) (r uint32) = crypt32.CertEnumCertificateContextProperties
//sys	CertCreateCTLEntryFromCertificateContextProperties(pCertContext *CERT_CONTEXT, cOptAttr uint32, rgOptAttr *CRYPT_ATTRIBUTE, dwFlags uint32, pvReserved unsafe.Pointer, pCtlEntry *CTL_ENTRY, pcbCtlEntry *uint32) (r BOOL, err error) = crypt32.CertCreateCTLEntryFromCertificateContextProperties
//sys	CertSetCertificateContextPropertiesFromCTLEntry(pCertContext *CERT_CONTEXT, pCtlEntry *CTL_ENTRY, dwFlags uint32) (r BOOL, err error) = crypt32.CertSetCertificateContextPropertiesFromCTLEntry
//sys	CertGetCRLFromStore(hCertStore HCERTSTORE, pIssuerContext *CERT_CONTEXT, pPrevCrlContext *CRL_CONTEXT, pdwFlags *uint32) (r *CRL_CONTEXT, err error) = crypt32.CertGetCRLFromStore
//sys	CertEnumCRLsInStore(hCertStore HCERTSTORE, pPrevCrlContext *CRL_CONTEXT) (r *CRL_CONTEXT, err error) = crypt32.CertEnumCRLsInStore
//sys	CertFindCRLInStore(hCertStore HCERTSTORE, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, dwFindFlags uint32, dwFindType uint32, pvFindPara unsafe.Pointer, pPrevCrlContext *CRL_CONTEXT) (r *CRL_CONTEXT, err error) = crypt32.CertFindCRLInStore
//sys	CertDuplicateCRLContext(pCrlContext *CRL_CONTEXT) (r *CRL_CONTEXT) = crypt32.CertDuplicateCRLContext
//sys	CertCreateCRLContext(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pbCrlEncoded *uint8, cbCrlEncoded uint32) (r *CRL_CONTEXT, err error) = crypt32.CertCreateCRLContext
//sys	CertFreeCRLContext(pCrlContext *CRL_CONTEXT) (r BOOL) = crypt32.CertFreeCRLContext
//sys	CertSetCRLContextProperty(pCrlContext *CRL_CONTEXT, dwPropId uint32, dwFlags uint32, pvData unsafe.Pointer) (r BOOL, err error) = crypt32.CertSetCRLContextProperty
//sys	CertGetCRLContextProperty(pCrlContext *CRL_CONTEXT, dwPropId uint32, pvData unsafe.Pointer, pcbData *uint32) (r BOOL, err error) = crypt32.CertGetCRLContextProperty
//sys	CertEnumCRLContextProperties(pCrlContext *CRL_CONTEXT, dwPropId uint32) (r uint32) = crypt32.CertEnumCRLContextProperties
//sys	CertFindCertificateInCRL(pCert *CERT_CONTEXT, pCrlContext *CRL_CONTEXT, dwFlags uint32, pvReserved unsafe.Pointer, ppCrlEntry **CRL_ENTRY) (r BOOL) = crypt32.CertFindCertificateInCRL
//sys	CertIsValidCRLForCertificate(pCert *CERT_CONTEXT, pCrl *CRL_CONTEXT, dwFlags uint32, pvReserved unsafe.Pointer) (r BOOL) = crypt32.CertIsValidCRLForCertificate
//sys	CertAddEncodedCertificateToStore(hCertStore HCERTSTORE, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pbCertEncoded *uint8, cbCertEncoded uint32, dwAddDisposition uint32, ppCertContext **CERT_CONTEXT) (r BOOL, err error) = crypt32.CertAddEncodedCertificateToStore
//sys	CertAddCertificateContextToStore(hCertStore HCERTSTORE, pCertContext *CERT_CONTEXT, dwAddDisposition uint32, ppStoreContext **CERT_CONTEXT) (r BOOL, err error) = crypt32.CertAddCertificateContextToStore
//sys	CertAddSerializedElementToStore(hCertStore HCERTSTORE, pbElement *uint8, cbElement uint32, dwAddDisposition uint32, dwFlags uint32, dwContextTypeFlags uint32, pdwContextType *uint32, ppvContext *unsafe.Pointer) (r BOOL, err error) = crypt32.CertAddSerializedElementToStore
//sys	CertDeleteCertificateFromStore(pCertContext *CERT_CONTEXT) (r BOOL, err error) = crypt32.CertDeleteCertificateFromStore
//sys	CertAddEncodedCRLToStore(hCertStore HCERTSTORE, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pbCrlEncoded *uint8, cbCrlEncoded uint32, dwAddDisposition uint32, ppCrlContext **CRL_CONTEXT) (r BOOL, err error) = crypt32.CertAddEncodedCRLToStore
//sys	CertAddCRLContextToStore(hCertStore HCERTSTORE, pCrlContext *CRL_CONTEXT, dwAddDisposition uint32, ppStoreContext **CRL_CONTEXT) (r BOOL, err error) = crypt32.CertAddCRLContextToStore
//sys	CertDeleteCRLFromStore(pCrlContext *CRL_CONTEXT) (r BOOL, err error) = crypt32.CertDeleteCRLFromStore
//sys	CertSerializeCertificateStoreElement(pCertContext *CERT_CONTEXT, dwFlags uint32, pbElement *uint8, pcbElement *uint32) (r BOOL, err error) = crypt32.CertSerializeCertificateStoreElement
//sys	CertSerializeCRLStoreElement(pCrlContext *CRL_CONTEXT, dwFlags uint32, pbElement *uint8, pcbElement *uint32) (r BOOL, err error) = crypt32.CertSerializeCRLStoreElement
//sys	CertDuplicateCTLContext(pCtlContext *CTL_CONTEXT) (r *CTL_CONTEXT) = crypt32.CertDuplicateCTLContext
//sys	CertCreateCTLContext(dwMsgAndCertEncodingType uint32, pbCtlEncoded *uint8, cbCtlEncoded uint32) (r *CTL_CONTEXT, err error) = crypt32.CertCreateCTLContext
//sys	CertFreeCTLContext(pCtlContext *CTL_CONTEXT) (r BOOL) = crypt32.CertFreeCTLContext
//sys	CertSetCTLContextProperty(pCtlContext *CTL_CONTEXT, dwPropId uint32, dwFlags uint32, pvData unsafe.Pointer) (r BOOL, err error) = crypt32.CertSetCTLContextProperty
//sys	CertGetCTLContextProperty(pCtlContext *CTL_CONTEXT, dwPropId uint32, pvData unsafe.Pointer, pcbData *uint32) (r BOOL, err error) = crypt32.CertGetCTLContextProperty
//sys	CertEnumCTLContextProperties(pCtlContext *CTL_CONTEXT, dwPropId uint32) (r uint32) = crypt32.CertEnumCTLContextProperties
//sys	CertEnumCTLsInStore(hCertStore HCERTSTORE, pPrevCtlContext *CTL_CONTEXT) (r *CTL_CONTEXT, err error) = crypt32.CertEnumCTLsInStore
//sys	CertFindSubjectInCTL(dwEncodingType uint32, dwSubjectType uint32, pvSubject unsafe.Pointer, pCtlContext *CTL_CONTEXT, dwFlags uint32) (r *CTL_ENTRY, err error) = crypt32.CertFindSubjectInCTL
//sys	CertFindCTLInStore(hCertStore HCERTSTORE, dwMsgAndCertEncodingType uint32, dwFindFlags uint32, dwFindType CERT_FIND_TYPE, pvFindPara unsafe.Pointer, pPrevCtlContext *CTL_CONTEXT) (r *CTL_CONTEXT, err error) = crypt32.CertFindCTLInStore
//sys	CertAddEncodedCTLToStore(hCertStore HCERTSTORE, dwMsgAndCertEncodingType CERT_QUERY_ENCODING_TYPE, pbCtlEncoded *uint8, cbCtlEncoded uint32, dwAddDisposition uint32, ppCtlContext **CTL_CONTEXT) (r BOOL, err error) = crypt32.CertAddEncodedCTLToStore
//sys	CertAddCTLContextToStore(hCertStore HCERTSTORE, pCtlContext *CTL_CONTEXT, dwAddDisposition uint32, ppStoreContext **CTL_CONTEXT) (r BOOL, err error) = crypt32.CertAddCTLContextToStore
//sys	CertSerializeCTLStoreElement(pCtlContext *CTL_CONTEXT, dwFlags uint32, pbElement *uint8, pcbElement *uint32) (r BOOL, err error) = crypt32.CertSerializeCTLStoreElement
//sys	CertDeleteCTLFromStore(pCtlContext *CTL_CONTEXT) (r BOOL, err error) = crypt32.CertDeleteCTLFromStore
//sys	CertAddCertificateLinkToStore(hCertStore HCERTSTORE, pCertContext *CERT_CONTEXT, dwAddDisposition uint32, ppStoreContext **CERT_CONTEXT) (r BOOL, err error) = crypt32.CertAddCertificateLinkToStore
//sys	CertAddCRLLinkToStore(hCertStore HCERTSTORE, pCrlContext *CRL_CONTEXT, dwAddDisposition uint32, ppStoreContext **CRL_CONTEXT) (r BOOL, err error) = crypt32.CertAddCRLLinkToStore
//sys	CertAddCTLLinkToStore(hCertStore HCERTSTORE, pCtlContext *CTL_CONTEXT, dwAddDisposition uint32, ppStoreContext **CTL_CONTEXT) (r BOOL, err error) = crypt32.CertAddCTLLinkToStore
//sys	CertAddStoreToCollection(hCollectionStore HCERTSTORE, hSiblingStore HCERTSTORE, dwUpdateFlags uint32, dwPriority uint32) (r BOOL) = crypt32.CertAddStoreToCollection
//sys	CertRemoveStoreFromCollection(hCollectionStore HCERTSTORE, hSiblingStore HCERTSTORE) = crypt32.CertRemoveStoreFromCollection
//sys	CertControlStore(hCertStore HCERTSTORE, dwFlags CERT_CONTROL_STORE_FLAGS, dwCtrlType uint32, pvCtrlPara unsafe.Pointer) (r BOOL, err error) = crypt32.CertControlStore
//sys	CertSetStoreProperty(hCertStore HCERTSTORE, dwPropId uint32, dwFlags uint32, pvData unsafe.Pointer) (r BOOL) = crypt32.CertSetStoreProperty
//sys	CertGetStoreProperty(hCertStore HCERTSTORE, dwPropId uint32, pvData unsafe.Pointer, pcbData *uint32) (r BOOL, err error) = crypt32.CertGetStoreProperty
//sys	CertCreateContext(dwContextType uint32, dwEncodingType uint32, pbEncoded *uint8, cbEncoded uint32, dwFlags uint32, pCreatePara *CERT_CREATE_CONTEXT_PARA) (r unsafe.Pointer, err error) = crypt32.CertCreateContext
//sys	CertRegisterSystemStore(pvSystemStore unsafe.Pointer, dwFlags uint32, pStoreInfo *CERT_SYSTEM_STORE_INFO, pvReserved unsafe.Pointer) (r BOOL) = crypt32.CertRegisterSystemStore
//sys	CertRegisterPhysicalStore(pvSystemStore unsafe.Pointer, dwFlags uint32, pwszStoreName *PWSTRElement, pStoreInfo *CERT_PHYSICAL_STORE_INFO, pvReserved unsafe.Pointer) (r BOOL) = crypt32.CertRegisterPhysicalStore
//sys	CertUnregisterSystemStore(pvSystemStore unsafe.Pointer, dwFlags uint32) (r BOOL) = crypt32.CertUnregisterSystemStore
//sys	CertUnregisterPhysicalStore(pvSystemStore unsafe.Pointer, dwFlags uint32, pwszStoreName *PWSTRElement) (r BOOL) = crypt32.CertUnregisterPhysicalStore
//sys	CertEnumSystemStoreLocation(dwFlags uint32, pvArg unsafe.Pointer, pfnEnum PFN_CERT_ENUM_SYSTEM_STORE_LOCATION) (r BOOL) = crypt32.CertEnumSystemStoreLocation
//sys	CertEnumSystemStore(dwFlags uint32, pvSystemStoreLocationPara unsafe.Pointer, pvArg unsafe.Pointer, pfnEnum PFN_CERT_ENUM_SYSTEM_STORE) (r BOOL) = crypt32.CertEnumSystemStore
//sys	CertEnumPhysicalStore(pvSystemStore unsafe.Pointer, dwFlags uint32, pvArg unsafe.Pointer, pfnEnum PFN_CERT_ENUM_PHYSICAL_STORE) (r BOOL, err error) = crypt32.CertEnumPhysicalStore
//sys	CertGetEnhancedKeyUsage(pCertContext *CERT_CONTEXT, dwFlags uint32, pUsage *CTL_USAGE, pcbUsage *uint32) (r BOOL, err error) = crypt32.CertGetEnhancedKeyUsage
//sys	CertSetEnhancedKeyUsage(pCertContext *CERT_CONTEXT, pUsage *CTL_USAGE) (r BOOL, err error) = crypt32.CertSetEnhancedKeyUsage
//sys	CertAddEnhancedKeyUsageIdentifier(pCertContext *CERT_CONTEXT, pszUsageIdentifier *PSTRElement) (r BOOL, err error) = crypt32.CertAddEnhancedKeyUsageIdentifier
//sys	CertRemoveEnhancedKeyUsageIdentifier(pCertContext *CERT_CONTEXT, pszUsageIdentifier *PSTRElement) (r BOOL, err error) = crypt32.CertRemoveEnhancedKeyUsageIdentifier
//sys	CertGetValidUsages(cCerts uint32, rghCerts **CERT_CONTEXT, cNumOIDs *int32, rghOIDs **PSTRElement, pcbOIDs *uint32) (r BOOL, err error) = crypt32.CertGetValidUsages
//sys	CryptMsgGetAndVerifySigner(hCryptMsg unsafe.Pointer, cSignerStore uint32, rghSignerStore *HCERTSTORE, dwFlags uint32, ppSigner **CERT_CONTEXT, pdwSignerIndex *uint32) (r BOOL, err error) = crypt32.CryptMsgGetAndVerifySigner
//sys	CryptMsgSignCTL(dwMsgEncodingType uint32, pbCtlContent *uint8, cbCtlContent uint32, pSignInfo *CMSG_SIGNED_ENCODE_INFO, dwFlags uint32, pbEncoded *uint8, pcbEncoded *uint32) (r BOOL, err error) = crypt32.CryptMsgSignCTL
//sys	CryptMsgEncodeAndSignCTL(dwMsgEncodingType uint32, pCtlInfo *CTL_INFO, pSignInfo *CMSG_SIGNED_ENCODE_INFO, dwFlags uint32, pbEncoded *uint8, pcbEncoded *uint32) (r BOOL, err error) = crypt32.CryptMsgEncodeAndSignCTL
//sys	CertFindSubjectInSortedCTL(pSubjectIdentifier *CRYPT_INTEGER_BLOB, pCtlContext *CTL_CONTEXT, dwFlags uint32, pvReserved unsafe.Pointer, pEncodedAttributes *CRYPT_INTEGER_BLOB) (r BOOL) = crypt32.CertFindSubjectInSortedCTL
//sys	CertEnumSubjectInSortedCTL(pCtlContext *CTL_CONTEXT, ppvNextSubject *unsafe.Pointer, pSubjectIdentifier *CRYPT_INTEGER_BLOB, pEncodedAttributes *CRYPT_INTEGER_BLOB) (r BOOL) = crypt32.CertEnumSubjectInSortedCTL
//sys	CertVerifyCTLUsage(dwEncodingType uint32, dwSubjectType uint32, pvSubject unsafe.Pointer, pSubjectUsage *CTL_USAGE, dwFlags uint32, pVerifyUsagePara *CTL_VERIFY_USAGE_PARA, pVerifyUsageStatus *CTL_VERIFY_USAGE_STATUS) (r BOOL, err error) = crypt32.CertVerifyCTLUsage
//sys	CertVerifyRevocation(dwEncodingType uint32, dwRevType uint32, cContext uint32, rgpvContext *unsafe.Pointer, dwFlags uint32, pRevPara *CERT_REVOCATION_PARA, pRevStatus *CERT_REVOCATION_STATUS) (r BOOL, err error) = crypt32.CertVerifyRevocation
//sys	CertCompareIntegerBlob(pInt1 *CRYPT_INTEGER_BLOB, pInt2 *CRYPT_INTEGER_BLOB) (r BOOL, err error) = crypt32.CertCompareIntegerBlob
//sys	CertCompareCertificate(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pCertId1 *CERT_INFO, pCertId2 *CERT_INFO) (r BOOL) = crypt32.CertCompareCertificate
//sys	CertCompareCertificateName(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pCertName1 *CRYPT_INTEGER_BLOB, pCertName2 *CRYPT_INTEGER_BLOB) (r BOOL) = crypt32.CertCompareCertificateName
//sys	CertIsRDNAttrsInCertificateName(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, dwFlags uint32, pCertName *CRYPT_INTEGER_BLOB, pRDN *CERT_RDN) (r BOOL, err error) = crypt32.CertIsRDNAttrsInCertificateName
//sys	CertComparePublicKeyInfo(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pPublicKey1 *CERT_PUBLIC_KEY_INFO, pPublicKey2 *CERT_PUBLIC_KEY_INFO) (r BOOL) = crypt32.CertComparePublicKeyInfo
//sys	CertGetPublicKeyLength(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pPublicKey *CERT_PUBLIC_KEY_INFO) (r uint32, err error) = crypt32.CertGetPublicKeyLength
//sys	CryptVerifyCertificateSignature(hCryptProv HCRYPTPROV_LEGACY, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pbEncoded *uint8, cbEncoded uint32, pPublicKey *CERT_PUBLIC_KEY_INFO) (r BOOL, err error) = crypt32.CryptVerifyCertificateSignature
//sys	CryptVerifyCertificateSignatureEx(hCryptProv HCRYPTPROV_LEGACY, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, dwSubjectType uint32, pvSubject unsafe.Pointer, dwIssuerType uint32, pvIssuer unsafe.Pointer, dwFlags CRYPT_VERIFY_CERT_FLAGS, pvExtra unsafe.Pointer) (r BOOL, err error) = crypt32.CryptVerifyCertificateSignatureEx
//sys	CertIsStrongHashToSign(pStrongSignPara *CERT_STRONG_SIGN_PARA, pwszCNGHashAlgid *PWSTRElement, pSigningCert *CERT_CONTEXT) (r BOOL, err error) = crypt32.CertIsStrongHashToSign
//sys	CryptHashToBeSigned(hCryptProv HCRYPTPROV_LEGACY, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pbEncoded *uint8, cbEncoded uint32, pbComputedHash *uint8, pcbComputedHash *uint32) (r BOOL, err error) = crypt32.CryptHashToBeSigned
//sys	CryptHashCertificate(hCryptProv HCRYPTPROV_LEGACY, Algid ALG_ID, dwFlags uint32, pbEncoded *uint8, cbEncoded uint32, pbComputedHash *uint8, pcbComputedHash *uint32) (r BOOL, err error) = crypt32.CryptHashCertificate
//sys	CryptHashCertificate2(pwszCNGHashAlgid *PWSTRElement, dwFlags uint32, pvReserved unsafe.Pointer, pbEncoded *uint8, cbEncoded uint32, pbComputedHash *uint8, pcbComputedHash *uint32) (r BOOL, err error) = crypt32.CryptHashCertificate2
//sys	CryptSignCertificate(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, dwKeySpec uint32, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pbEncodedToBeSigned *uint8, cbEncodedToBeSigned uint32, pSignatureAlgorithm *CRYPT_ALGORITHM_IDENTIFIER, pvHashAuxInfo unsafe.Pointer, pbSignature *uint8, pcbSignature *uint32) (r BOOL, err error) = crypt32.CryptSignCertificate
//sys	CryptSignAndEncodeCertificate(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, dwKeySpec CERT_KEY_SPEC, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, lpszStructType *PSTRElement, pvStructInfo unsafe.Pointer, pSignatureAlgorithm *CRYPT_ALGORITHM_IDENTIFIER, pvHashAuxInfo unsafe.Pointer, pbEncoded *uint8, pcbEncoded *uint32) (r BOOL, err error) = crypt32.CryptSignAndEncodeCertificate
//sys	CertVerifyTimeValidity(pTimeToVerify *FILETIME, pCertInfo *CERT_INFO) (r int32) = crypt32.CertVerifyTimeValidity
//sys	CertVerifyCRLTimeValidity(pTimeToVerify *FILETIME, pCrlInfo *CRL_INFO) (r int32) = crypt32.CertVerifyCRLTimeValidity
//sys	CertVerifyValidityNesting(pSubjectInfo *CERT_INFO, pIssuerInfo *CERT_INFO) (r BOOL) = crypt32.CertVerifyValidityNesting
//sys	CertVerifyCRLRevocation(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pCertId *CERT_INFO, cCrlInfo uint32, rgpCrlInfo **CRL_INFO) (r BOOL) = crypt32.CertVerifyCRLRevocation
//sys	CertAlgIdToOID(dwAlgId uint32) (r *PSTRElement) = crypt32.CertAlgIdToOID
//sys	CertOIDToAlgId(pszObjId *PSTRElement) (r uint32) = crypt32.CertOIDToAlgId
//sys	CertFindExtension(pszObjId *PSTRElement, cExtensions uint32, rgExtensions *CERT_EXTENSION) (r *CERT_EXTENSION) = crypt32.CertFindExtension
//sys	CertFindAttribute(pszObjId *PSTRElement, cAttr uint32, rgAttr *CRYPT_ATTRIBUTE) (r *CRYPT_ATTRIBUTE) = crypt32.CertFindAttribute
//sys	CertFindRDNAttr(pszObjId *PSTRElement, pName *CERT_NAME_INFO) (r *CERT_RDN_ATTR) = crypt32.CertFindRDNAttr
//sys	CertGetIntendedKeyUsage(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pCertInfo *CERT_INFO, pbKeyUsage *uint8, cbKeyUsage uint32) (r BOOL, err error) = crypt32.CertGetIntendedKeyUsage
//sys	CryptInstallDefaultContext(hCryptProv uintptr, dwDefaultType CRYPT_DEFAULT_CONTEXT_TYPE, pvDefaultPara unsafe.Pointer, dwFlags CRYPT_DEFAULT_CONTEXT_FLAGS, pvReserved unsafe.Pointer, phDefaultContext *unsafe.Pointer) (r BOOL, err error) = crypt32.CryptInstallDefaultContext
//sys	CryptUninstallDefaultContext(hDefaultContext unsafe.Pointer, dwFlags uint32, pvReserved unsafe.Pointer) (r BOOL, err error) = crypt32.CryptUninstallDefaultContext
//sys	CryptExportPublicKeyInfo(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, dwKeySpec uint32, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pInfo *CERT_PUBLIC_KEY_INFO, pcbInfo *uint32) (r BOOL, err error) = crypt32.CryptExportPublicKeyInfo
//sys	CryptExportPublicKeyInfoEx(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, dwKeySpec uint32, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pszPublicKeyObjId *PSTRElement, dwFlags uint32, pvAuxInfo unsafe.Pointer, pInfo *CERT_PUBLIC_KEY_INFO, pcbInfo *uint32) (r BOOL, err error) = crypt32.CryptExportPublicKeyInfoEx
//sys	CryptExportPublicKeyInfoFromBCryptKeyHandle(hBCryptKey BCRYPT_KEY_HANDLE, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pszPublicKeyObjId *PSTRElement, dwFlags uint32, pvAuxInfo unsafe.Pointer, pInfo *CERT_PUBLIC_KEY_INFO, pcbInfo *uint32) (r BOOL) = crypt32.CryptExportPublicKeyInfoFromBCryptKeyHandle
//sys	CryptImportPublicKeyInfo(hCryptProv uintptr, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pInfo *CERT_PUBLIC_KEY_INFO, phKey *uintptr) (r BOOL, err error) = crypt32.CryptImportPublicKeyInfo
//sys	CryptImportPublicKeyInfoEx(hCryptProv uintptr, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pInfo *CERT_PUBLIC_KEY_INFO, aiKeyAlg ALG_ID, dwFlags uint32, pvAuxInfo unsafe.Pointer, phKey *uintptr) (r BOOL, err error) = crypt32.CryptImportPublicKeyInfoEx
//sys	CryptImportPublicKeyInfoEx2(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pInfo *CERT_PUBLIC_KEY_INFO, dwFlags CRYPT_IMPORT_PUBLIC_KEY_FLAGS, pvAuxInfo unsafe.Pointer, phKey *BCRYPT_KEY_HANDLE) (r BOOL, err error) = crypt32.CryptImportPublicKeyInfoEx2
//sys	CryptAcquireCertificatePrivateKey(pCert *CERT_CONTEXT, dwFlags CRYPT_ACQUIRE_FLAGS, pvParameters unsafe.Pointer, phCryptProvOrNCryptKey *HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, pdwKeySpec *CERT_KEY_SPEC, pfCallerFreeProvOrNCryptKey *BOOL) (r BOOL, err error) = crypt32.CryptAcquireCertificatePrivateKey
//sys	CryptFindCertificateKeyProvInfo(pCert *CERT_CONTEXT, dwFlags CRYPT_FIND_FLAGS, pvReserved unsafe.Pointer) (r BOOL, err error) = crypt32.CryptFindCertificateKeyProvInfo
//sys	CryptImportPKCS8(sPrivateKeyAndParams CRYPT_PKCS8_IMPORT_PARAMS, dwFlags CRYPT_KEY_FLAGS, phCryptProv *uintptr, pvAuxInfo unsafe.Pointer) (r BOOL, err error) = crypt32.CryptImportPKCS8
//sys	CryptExportPKCS8(hCryptProv uintptr, dwKeySpec uint32, pszPrivateKeyObjId *PSTRElement, dwFlags uint32, pvAuxInfo unsafe.Pointer, pbPrivateKeyBlob *uint8, pcbPrivateKeyBlob *uint32) (r BOOL, err error) = crypt32.CryptExportPKCS8
//sys	CryptHashPublicKeyInfo(hCryptProv HCRYPTPROV_LEGACY, Algid ALG_ID, dwFlags uint32, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pInfo *CERT_PUBLIC_KEY_INFO, pbComputedHash *uint8, pcbComputedHash *uint32) (r BOOL, err error) = crypt32.CryptHashPublicKeyInfo
//sys	CertRDNValueToStrA(dwValueType uint32, pValue *CRYPT_INTEGER_BLOB, psz *PSTRElement, csz uint32) (r uint32) = crypt32.CertRDNValueToStrA
//sys	CertRDNValueToStrW(dwValueType uint32, pValue *CRYPT_INTEGER_BLOB, psz *PWSTRElement, csz uint32) (r uint32) = crypt32.CertRDNValueToStrW
//sys	CertNameToStrA(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pName *CRYPT_INTEGER_BLOB, dwStrType CERT_STRING_TYPE, psz *PSTRElement, csz uint32) (r uint32) = crypt32.CertNameToStrA
//sys	CertNameToStrW(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pName *CRYPT_INTEGER_BLOB, dwStrType CERT_STRING_TYPE, psz *PWSTRElement, csz uint32) (r uint32) = crypt32.CertNameToStrW
//sys	CertStrToNameA(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pszX500 *PSTRElement, dwStrType CERT_STRING_TYPE, pvReserved unsafe.Pointer, pbEncoded *uint8, pcbEncoded *uint32, ppszError **PSTRElement) (r BOOL, err error) = crypt32.CertStrToNameA
//sys	CertStrToNameW(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pszX500 *PWSTRElement, dwStrType CERT_STRING_TYPE, pvReserved unsafe.Pointer, pbEncoded *uint8, pcbEncoded *uint32, ppszError **PWSTRElement) (r BOOL, err error) = crypt32.CertStrToNameW
//sys	CertGetNameStringA(pCertContext *CERT_CONTEXT, dwType uint32, dwFlags uint32, pvTypePara unsafe.Pointer, pszNameString *PSTRElement, cchNameString uint32) (r uint32) = crypt32.CertGetNameStringA
//sys	CertGetNameStringW(pCertContext *CERT_CONTEXT, dwType uint32, dwFlags uint32, pvTypePara unsafe.Pointer, pszNameString *PWSTRElement, cchNameString uint32) (r uint32) = crypt32.CertGetNameStringW
//sys	CryptSignMessage(pSignPara *CRYPT_SIGN_MESSAGE_PARA, fDetachedSignature BOOL, cToBeSigned uint32, rgpbToBeSigned **uint8, rgcbToBeSigned *uint32, pbSignedBlob *uint8, pcbSignedBlob *uint32) (r BOOL, err error) = crypt32.CryptSignMessage
//sys	CryptVerifyMessageSignature(pVerifyPara *CRYPT_VERIFY_MESSAGE_PARA, dwSignerIndex uint32, pbSignedBlob *uint8, cbSignedBlob uint32, pbDecoded *uint8, pcbDecoded *uint32, ppSignerCert **CERT_CONTEXT) (r BOOL, err error) = crypt32.CryptVerifyMessageSignature
//sys	CryptGetMessageSignerCount(dwMsgEncodingType uint32, pbSignedBlob *uint8, cbSignedBlob uint32) (r int32, err error) = crypt32.CryptGetMessageSignerCount
//sys	CryptGetMessageCertificates(dwMsgAndCertEncodingType uint32, hCryptProv HCRYPTPROV_LEGACY, dwFlags uint32, pbSignedBlob *uint8, cbSignedBlob uint32) (r HCERTSTORE, err error) = crypt32.CryptGetMessageCertificates
//sys	CryptVerifyDetachedMessageSignature(pVerifyPara *CRYPT_VERIFY_MESSAGE_PARA, dwSignerIndex uint32, pbDetachedSignBlob *uint8, cbDetachedSignBlob uint32, cToBeSigned uint32, rgpbToBeSigned **uint8, rgcbToBeSigned *uint32, ppSignerCert **CERT_CONTEXT) (r BOOL, err error) = crypt32.CryptVerifyDetachedMessageSignature
//sys	CryptEncryptMessage(pEncryptPara *CRYPT_ENCRYPT_MESSAGE_PARA, cRecipientCert uint32, rgpRecipientCert **CERT_CONTEXT, pbToBeEncrypted *uint8, cbToBeEncrypted uint32, pbEncryptedBlob *uint8, pcbEncryptedBlob *uint32) (r BOOL, err error) = crypt32.CryptEncryptMessage
//sys	CryptDecryptMessage(pDecryptPara *CRYPT_DECRYPT_MESSAGE_PARA, pbEncryptedBlob *uint8, cbEncryptedBlob uint32, pbDecrypted *uint8, pcbDecrypted *uint32, ppXchgCert **CERT_CONTEXT) (r BOOL, err error) = crypt32.CryptDecryptMessage
//sys	CryptSignAndEncryptMessage(pSignPara *CRYPT_SIGN_MESSAGE_PARA, pEncryptPara *CRYPT_ENCRYPT_MESSAGE_PARA, cRecipientCert uint32, rgpRecipientCert **CERT_CONTEXT, pbToBeSignedAndEncrypted *uint8, cbToBeSignedAndEncrypted uint32, pbSignedAndEncryptedBlob *uint8, pcbSignedAndEncryptedBlob *uint32) (r BOOL, err error) = crypt32.CryptSignAndEncryptMessage
//sys	CryptDecryptAndVerifyMessageSignature(pDecryptPara *CRYPT_DECRYPT_MESSAGE_PARA, pVerifyPara *CRYPT_VERIFY_MESSAGE_PARA, dwSignerIndex uint32, pbEncryptedBlob *uint8, cbEncryptedBlob uint32, pbDecrypted *uint8, pcbDecrypted *uint32, ppXchgCert **CERT_CONTEXT, ppSignerCert **CERT_CONTEXT) (r BOOL, err error) = crypt32.CryptDecryptAndVerifyMessageSignature
//sys	CryptDecodeMessage(dwMsgTypeFlags uint32, pDecryptPara *CRYPT_DECRYPT_MESSAGE_PARA, pVerifyPara *CRYPT_VERIFY_MESSAGE_PARA, dwSignerIndex uint32, pbEncodedBlob *uint8, cbEncodedBlob uint32, dwPrevInnerContentType uint32, pdwMsgType *uint32, pdwInnerContentType *uint32, pbDecoded *uint8, pcbDecoded *uint32, ppXchgCert **CERT_CONTEXT, ppSignerCert **CERT_CONTEXT) (r BOOL, err error) = crypt32.CryptDecodeMessage
//sys	CryptHashMessage(pHashPara *CRYPT_HASH_MESSAGE_PARA, fDetachedHash BOOL, cToBeHashed uint32, rgpbToBeHashed **uint8, rgcbToBeHashed *uint32, pbHashedBlob *uint8, pcbHashedBlob *uint32, pbComputedHash *uint8, pcbComputedHash *uint32) (r BOOL, err error) = crypt32.CryptHashMessage
//sys	CryptVerifyMessageHash(pHashPara *CRYPT_HASH_MESSAGE_PARA, pbHashedBlob *uint8, cbHashedBlob uint32, pbToBeHashed *uint8, pcbToBeHashed *uint32, pbComputedHash *uint8, pcbComputedHash *uint32) (r BOOL, err error) = crypt32.CryptVerifyMessageHash
//sys	CryptVerifyDetachedMessageHash(pHashPara *CRYPT_HASH_MESSAGE_PARA, pbDetachedHashBlob *uint8, cbDetachedHashBlob uint32, cToBeHashed uint32, rgpbToBeHashed **uint8, rgcbToBeHashed *uint32, pbComputedHash *uint8, pcbComputedHash *uint32) (r BOOL, err error) = crypt32.CryptVerifyDetachedMessageHash
//sys	CryptSignMessageWithKey(pSignPara *CRYPT_KEY_SIGN_MESSAGE_PARA, pbToBeSigned *uint8, cbToBeSigned uint32, pbSignedBlob *uint8, pcbSignedBlob *uint32) (r BOOL, err error) = crypt32.CryptSignMessageWithKey
//sys	CryptVerifyMessageSignatureWithKey(pVerifyPara *CRYPT_KEY_VERIFY_MESSAGE_PARA, pPublicKeyInfo *CERT_PUBLIC_KEY_INFO, pbSignedBlob *uint8, cbSignedBlob uint32, pbDecoded *uint8, pcbDecoded *uint32) (r BOOL, err error) = crypt32.CryptVerifyMessageSignatureWithKey
//sys	CertOpenSystemStoreA(hProv HCRYPTPROV_LEGACY, szSubsystemProtocol *PSTRElement) (r HCERTSTORE, err error) = crypt32.CertOpenSystemStoreA
//sys	CertOpenSystemStoreW(hProv HCRYPTPROV_LEGACY, szSubsystemProtocol *PWSTRElement) (r HCERTSTORE, err error) = crypt32.CertOpenSystemStoreW
//sys	CertAddEncodedCertificateToSystemStoreA(szCertStoreName *PSTRElement, pbCertEncoded *uint8, cbCertEncoded uint32) (r BOOL, err error) = crypt32.CertAddEncodedCertificateToSystemStoreA
//sys	CertAddEncodedCertificateToSystemStoreW(szCertStoreName *PWSTRElement, pbCertEncoded *uint8, cbCertEncoded uint32) (r BOOL, err error) = crypt32.CertAddEncodedCertificateToSystemStoreW
//sys	FindCertsByIssuer(pCertChains *CERT_CHAIN, pcbCertChains *uint32, pcCertChains *uint32, pbEncodedIssuerName *uint8, cbEncodedIssuerName uint32, pwszPurpose *PWSTRElement, dwKeySpec uint32) (r HRESULT) = wintrust.FindCertsByIssuer
//sys	CryptQueryObject(dwObjectType CERT_QUERY_OBJECT_TYPE, pvObject unsafe.Pointer, dwExpectedContentTypeFlags CERT_QUERY_CONTENT_TYPE_FLAGS, dwExpectedFormatTypeFlags CERT_QUERY_FORMAT_TYPE_FLAGS, dwFlags uint32, pdwMsgAndCertEncodingType *CERT_QUERY_ENCODING_TYPE, pdwContentType *CERT_QUERY_CONTENT_TYPE, pdwFormatType *CERT_QUERY_FORMAT_TYPE, phCertStore *HCERTSTORE, phMsg *unsafe.Pointer, ppvContext *unsafe.Pointer) (r BOOL, err error) = crypt32.CryptQueryObject
//sys	CryptMemAlloc(cbSize uint32) (r unsafe.Pointer) = crypt32.CryptMemAlloc
//sys	CryptMemRealloc(pv unsafe.Pointer, cbSize uint32) (r unsafe.Pointer) = crypt32.CryptMemRealloc
//sys	CryptMemFree(pv unsafe.Pointer) = crypt32.CryptMemFree
//sys	CryptCreateAsyncHandle(dwFlags uint32, phAsync *HCRYPTASYNC) (r BOOL) = crypt32.CryptCreateAsyncHandle
//sys	CryptSetAsyncParam(hAsync HCRYPTASYNC, pszParamOid *PSTRElement, pvParam unsafe.Pointer, pfnFree PFN_CRYPT_ASYNC_PARAM_FREE_FUNC) (r BOOL) = crypt32.CryptSetAsyncParam
//sys	CryptGetAsyncParam(hAsync HCRYPTASYNC, pszParamOid *PSTRElement, ppvParam *unsafe.Pointer, ppfnFree *PFN_CRYPT_ASYNC_PARAM_FREE_FUNC) (r BOOL) = crypt32.CryptGetAsyncParam
//sys	CryptCloseAsyncHandle(hAsync HCRYPTASYNC) (r BOOL) = crypt32.CryptCloseAsyncHandle
//sys	CryptRetrieveObjectByUrlA(pszUrl *PSTRElement, pszObjectOid *PSTRElement, dwRetrievalFlags uint32, dwTimeout uint32, ppvObject *unsafe.Pointer, hAsyncRetrieve HCRYPTASYNC, pCredentials *CRYPT_CREDENTIALS, pvVerify unsafe.Pointer, pAuxInfo *CRYPT_RETRIEVE_AUX_INFO) (r BOOL) = cryptnet.CryptRetrieveObjectByUrlA
//sys	CryptRetrieveObjectByUrlW(pszUrl *PWSTRElement, pszObjectOid *PSTRElement, dwRetrievalFlags uint32, dwTimeout uint32, ppvObject *unsafe.Pointer, hAsyncRetrieve HCRYPTASYNC, pCredentials *CRYPT_CREDENTIALS, pvVerify unsafe.Pointer, pAuxInfo *CRYPT_RETRIEVE_AUX_INFO) (r BOOL) = cryptnet.CryptRetrieveObjectByUrlW
//sys	CryptInstallCancelRetrieval(pfnCancel PFN_CRYPT_CANCEL_RETRIEVAL, pvArg unsafe.Pointer, dwFlags uint32, pvReserved unsafe.Pointer) (r BOOL) = cryptnet.CryptInstallCancelRetrieval
//sys	CryptUninstallCancelRetrieval(dwFlags uint32, pvReserved unsafe.Pointer) (r BOOL) = cryptnet.CryptUninstallCancelRetrieval
//sys	CryptGetObjectUrl(pszUrlOid *PSTRElement, pvPara unsafe.Pointer, dwFlags CRYPT_GET_URL_FLAGS, pUrlArray *CRYPT_URL_ARRAY, pcbUrlArray *uint32, pUrlInfo *CRYPT_URL_INFO, pcbUrlInfo *uint32, pvReserved unsafe.Pointer) (r BOOL, err error) = cryptnet.CryptGetObjectUrl
//sys	CertCreateSelfSignCertificate(hCryptProvOrNCryptKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, pSubjectIssuerBlob *CRYPT_INTEGER_BLOB, dwFlags CERT_CREATE_SELFSIGN_FLAGS, pKeyProvInfo *CRYPT_KEY_PROV_INFO, pSignatureAlgorithm *CRYPT_ALGORITHM_IDENTIFIER, pStartTime *SYSTEMTIME, pEndTime *SYSTEMTIME, pExtensions *CERT_EXTENSIONS) (r *CERT_CONTEXT, err error) = crypt32.CertCreateSelfSignCertificate
//sys	CryptGetKeyIdentifierProperty(pKeyIdentifier *CRYPT_INTEGER_BLOB, dwPropId uint32, dwFlags uint32, pwszComputerName *PWSTRElement, pvReserved unsafe.Pointer, pvData unsafe.Pointer, pcbData *uint32) (r BOOL, err error) = crypt32.CryptGetKeyIdentifierProperty
//sys	CryptSetKeyIdentifierProperty(pKeyIdentifier *CRYPT_INTEGER_BLOB, dwPropId uint32, dwFlags uint32, pwszComputerName *PWSTRElement, pvReserved unsafe.Pointer, pvData unsafe.Pointer) (r BOOL, err error) = crypt32.CryptSetKeyIdentifierProperty
//sys	CryptEnumKeyIdentifierProperties(pKeyIdentifier *CRYPT_INTEGER_BLOB, dwPropId uint32, dwFlags uint32, pwszComputerName *PWSTRElement, pvReserved unsafe.Pointer, pvArg unsafe.Pointer, pfnEnum PFN_CRYPT_ENUM_KEYID_PROP) (r BOOL, err error) = crypt32.CryptEnumKeyIdentifierProperties
//sys	CryptCreateKeyIdentifierFromCSP(dwCertEncodingType CERT_QUERY_ENCODING_TYPE, pszPubKeyOID *PSTRElement, pPubKeyStruc *PUBLICKEYSTRUC, cbPubKeyStruc uint32, dwFlags uint32, pvReserved unsafe.Pointer, pbHash *uint8, pcbHash *uint32) (r BOOL, err error) = crypt32.CryptCreateKeyIdentifierFromCSP
//sys	CertCreateCertificateChainEngine(pConfig *CERT_CHAIN_ENGINE_CONFIG, phChainEngine *HCERTCHAINENGINE) (r BOOL, err error) = crypt32.CertCreateCertificateChainEngine
//sys	CertFreeCertificateChainEngine(hChainEngine HCERTCHAINENGINE) = crypt32.CertFreeCertificateChainEngine
//sys	CertResyncCertificateChainEngine(hChainEngine HCERTCHAINENGINE) (r BOOL, err error) = crypt32.CertResyncCertificateChainEngine
//sys	CertGetCertificateChain(hChainEngine HCERTCHAINENGINE, pCertContext *CERT_CONTEXT, pTime *FILETIME, hAdditionalStore HCERTSTORE, pChainPara *CERT_CHAIN_PARA, dwFlags uint32, pvReserved unsafe.Pointer, ppChainContext **CERT_CHAIN_CONTEXT) (r BOOL, err error) = crypt32.CertGetCertificateChain
//sys	CertFreeCertificateChain(pChainContext *CERT_CHAIN_CONTEXT) = crypt32.CertFreeCertificateChain
//sys	CertDuplicateCertificateChain(pChainContext *CERT_CHAIN_CONTEXT) (r *CERT_CHAIN_CONTEXT) = crypt32.CertDuplicateCertificateChain
//sys	CertFindChainInStore(hCertStore HCERTSTORE, dwCertEncodingType CERT_QUERY_ENCODING_TYPE, dwFindFlags CERT_FIND_CHAIN_IN_STORE_FLAGS, dwFindType uint32, pvFindPara unsafe.Pointer, pPrevChainContext *CERT_CHAIN_CONTEXT) (r *CERT_CHAIN_CONTEXT) = crypt32.CertFindChainInStore
//sys	CertVerifyCertificateChainPolicy(pszPolicyOID *PSTRElement, pChainContext *CERT_CHAIN_CONTEXT, pPolicyPara *CERT_CHAIN_POLICY_PARA, pPolicyStatus *CERT_CHAIN_POLICY_STATUS) (r BOOL) = crypt32.CertVerifyCertificateChainPolicy
//sys	CryptStringToBinaryA(pszString *PSTRElement, cchString uint32, dwFlags CRYPT_STRING, pbBinary *uint8, pcbBinary *uint32, pdwSkip *uint32, pdwFlags *uint32) (r BOOL, err error) = crypt32.CryptStringToBinaryA
//sys	CryptStringToBinaryW(pszString *PWSTRElement, cchString uint32, dwFlags CRYPT_STRING, pbBinary *uint8, pcbBinary *uint32, pdwSkip *uint32, pdwFlags *uint32) (r BOOL, err error) = crypt32.CryptStringToBinaryW
//sys	CryptBinaryToStringA(pbBinary *uint8, cbBinary uint32, dwFlags CRYPT_STRING, pszString *PSTRElement, pcchString *uint32) (r BOOL) = crypt32.CryptBinaryToStringA
//sys	CryptBinaryToStringW(pbBinary *uint8, cbBinary uint32, dwFlags CRYPT_STRING, pszString *PWSTRElement, pcchString *uint32) (r BOOL) = crypt32.CryptBinaryToStringW
//sys	PFXImportCertStore(pPFX *CRYPT_INTEGER_BLOB, szPassword *PWSTRElement, dwFlags CRYPT_KEY_FLAGS) (r HCERTSTORE, err error) = crypt32.PFXImportCertStore
//sys	PFXIsPFXBlob(pPFX *CRYPT_INTEGER_BLOB) (r BOOL) = crypt32.PFXIsPFXBlob
//sys	PFXVerifyPassword(pPFX *CRYPT_INTEGER_BLOB, szPassword *PWSTRElement, dwFlags uint32) (r BOOL) = crypt32.PFXVerifyPassword
//sys	PFXExportCertStoreEx(hStore HCERTSTORE, pPFX *CRYPT_INTEGER_BLOB, szPassword *PWSTRElement, pvPara unsafe.Pointer, dwFlags uint32) (r BOOL, err error) = crypt32.PFXExportCertStoreEx
//sys	PFXExportCertStore(hStore HCERTSTORE, pPFX *CRYPT_INTEGER_BLOB, szPassword *PWSTRElement, dwFlags uint32) (r BOOL, err error) = crypt32.PFXExportCertStore
//sys	CertOpenServerOcspResponse(pChainContext *CERT_CHAIN_CONTEXT, dwFlags uint32, pOpenPara *CERT_SERVER_OCSP_RESPONSE_OPEN_PARA) (r unsafe.Pointer, err error) = crypt32.CertOpenServerOcspResponse
//sys	CertAddRefServerOcspResponse(hServerOcspResponse unsafe.Pointer) = crypt32.CertAddRefServerOcspResponse
//sys	CertCloseServerOcspResponse(hServerOcspResponse unsafe.Pointer, dwFlags uint32) = crypt32.CertCloseServerOcspResponse
//sys	CertGetServerOcspResponseContext(hServerOcspResponse unsafe.Pointer, dwFlags uint32, pvReserved unsafe.Pointer) (r *CERT_SERVER_OCSP_RESPONSE_CONTEXT) = crypt32.CertGetServerOcspResponseContext
//sys	CertAddRefServerOcspResponseContext(pServerOcspResponseContext *CERT_SERVER_OCSP_RESPONSE_CONTEXT) = crypt32.CertAddRefServerOcspResponseContext
//sys	CertFreeServerOcspResponseContext(pServerOcspResponseContext *CERT_SERVER_OCSP_RESPONSE_CONTEXT) = crypt32.CertFreeServerOcspResponseContext
//sys	CertRetrieveLogoOrBiometricInfo(pCertContext *CERT_CONTEXT, lpszLogoOrBiometricType *PSTRElement, dwRetrievalFlags uint32, dwTimeout uint32, dwFlags uint32, pvReserved unsafe.Pointer, ppbData **uint8, pcbData *uint32, ppwszMimeType **PWSTRElement) (r BOOL, err error) = crypt32.CertRetrieveLogoOrBiometricInfo
//sys	CertSelectCertificateChains(pSelectionContext *Guid, dwFlags uint32, pChainParameters *CERT_SELECT_CHAIN_PARA, cCriteria uint32, rgpCriteria *CERT_SELECT_CRITERIA, hStore HCERTSTORE, pcSelection *uint32, pprgpSelection ***CERT_CHAIN_CONTEXT) (r BOOL, err error) = crypt32.CertSelectCertificateChains
//sys	CertFreeCertificateChainList(prgpSelection **CERT_CHAIN_CONTEXT) = crypt32.CertFreeCertificateChainList
//sys	CryptRetrieveTimeStamp(wszUrl *PWSTRElement, dwRetrievalFlags uint32, dwTimeout uint32, pszHashId *PSTRElement, pPara *CRYPT_TIMESTAMP_PARA, pbData *uint8, cbData uint32, ppTsContext **CRYPT_TIMESTAMP_CONTEXT, ppTsSigner **CERT_CONTEXT, phStore *HCERTSTORE) (r BOOL, err error) = crypt32.CryptRetrieveTimeStamp
//sys	CryptVerifyTimeStampSignature(pbTSContentInfo *uint8, cbTSContentInfo uint32, pbData *uint8, cbData uint32, hAdditionalStore HCERTSTORE, ppTsContext **CRYPT_TIMESTAMP_CONTEXT, ppTsSigner **CERT_CONTEXT, phStore *HCERTSTORE) (r BOOL, err error) = crypt32.CryptVerifyTimeStampSignature
//sys	CertIsWeakHash(dwHashUseType uint32, pwszCNGHashAlgid *PWSTRElement, dwChainFlags uint32, pSignerChainContext *CERT_CHAIN_CONTEXT, pTimeStamp *FILETIME, pwszFileName *PWSTRElement) (r BOOL) = crypt32.CertIsWeakHash
//sys	CryptProtectData(pDataIn *CRYPT_INTEGER_BLOB, szDataDescr *PWSTRElement, pOptionalEntropy *CRYPT_INTEGER_BLOB, pvReserved unsafe.Pointer, pPromptStruct *CRYPTPROTECT_PROMPTSTRUCT, dwFlags uint32, pDataOut *CRYPT_INTEGER_BLOB) (r BOOL, err error) = crypt32.CryptProtectData
//sys	CryptUnprotectData(pDataIn *CRYPT_INTEGER_BLOB, ppszDataDescr **PWSTRElement, pOptionalEntropy *CRYPT_INTEGER_BLOB, pvReserved unsafe.Pointer, pPromptStruct *CRYPTPROTECT_PROMPTSTRUCT, dwFlags uint32, pDataOut *CRYPT_INTEGER_BLOB) (r BOOL, err error) = crypt32.CryptUnprotectData
//sys	CryptUpdateProtectedState(pOldSid PSID, pwszOldPassword *PWSTRElement, dwFlags uint32, pdwSuccessCount *uint32, pdwFailureCount *uint32) (r BOOL, err error) = crypt32.CryptUpdateProtectedState
//sys	CryptProtectMemory(pDataIn unsafe.Pointer, cbDataIn uint32, dwFlags uint32) (r BOOL, err error) = crypt32.CryptProtectMemory
//sys	CryptUnprotectMemory(pDataIn unsafe.Pointer, cbDataIn uint32, dwFlags uint32) (r BOOL, err error) = crypt32.CryptUnprotectMemory
//sys	NCryptRegisterProtectionDescriptorName(pwszName *PWSTRElement, pwszDescriptorString *PWSTRElement, dwFlags uint32) (r HRESULT) = ncrypt.NCryptRegisterProtectionDescriptorName
//sys	NCryptQueryProtectionDescriptorName(pwszName *PWSTRElement, pwszDescriptorString *PWSTRElement, pcDescriptorString *uintptr, dwFlags uint32) (r HRESULT) = ncrypt.NCryptQueryProtectionDescriptorName
//sys	NCryptCreateProtectionDescriptor(pwszDescriptorString *PWSTRElement, dwFlags uint32, phDescriptor *NCRYPT_DESCRIPTOR_HANDLE) (r HRESULT) = ncrypt.NCryptCreateProtectionDescriptor
//sys	NCryptCloseProtectionDescriptor(hDescriptor NCRYPT_DESCRIPTOR_HANDLE) (r HRESULT) = ncrypt.NCryptCloseProtectionDescriptor
//sys	NCryptGetProtectionDescriptorInfo(hDescriptor NCRYPT_DESCRIPTOR_HANDLE, pMemPara *NCRYPT_ALLOC_PARA, dwInfoType uint32, ppvInfo *unsafe.Pointer) (r HRESULT) = ncrypt.NCryptGetProtectionDescriptorInfo
//sys	NCryptProtectSecret(hDescriptor NCRYPT_DESCRIPTOR_HANDLE, dwFlags uint32, pbData *uint8, cbData uint32, pMemPara *NCRYPT_ALLOC_PARA, hWnd HWND, ppbProtectedBlob **uint8, pcbProtectedBlob *uint32) (r HRESULT) = ncrypt.NCryptProtectSecret
//sys	NCryptUnprotectSecret(phDescriptor *NCRYPT_DESCRIPTOR_HANDLE, dwFlags NCRYPT_FLAGS, pbProtectedBlob *uint8, cbProtectedBlob uint32, pMemPara *NCRYPT_ALLOC_PARA, hWnd HWND, ppbData **uint8, pcbData *uint32) (r HRESULT) = ncrypt.NCryptUnprotectSecret
//sys	NCryptStreamOpenToProtect(hDescriptor NCRYPT_DESCRIPTOR_HANDLE, dwFlags uint32, hWnd HWND, pStreamInfo *NCRYPT_PROTECT_STREAM_INFO, phStream *NCRYPT_STREAM_HANDLE) (r HRESULT) = ncrypt.NCryptStreamOpenToProtect
//sys	NCryptStreamOpenToUnprotect(pStreamInfo *NCRYPT_PROTECT_STREAM_INFO, dwFlags uint32, hWnd HWND, phStream *NCRYPT_STREAM_HANDLE) (r HRESULT) = ncrypt.NCryptStreamOpenToUnprotect
//sys	NCryptStreamOpenToUnprotectEx(pStreamInfo *NCRYPT_PROTECT_STREAM_INFO_EX, dwFlags uint32, hWnd HWND, phStream *NCRYPT_STREAM_HANDLE) (r HRESULT) = ncrypt.NCryptStreamOpenToUnprotectEx
//sys	NCryptStreamUpdate(hStream NCRYPT_STREAM_HANDLE, pbData *uint8, cbData uintptr, fFinal BOOL) (r HRESULT) = ncrypt.NCryptStreamUpdate
//sys	NCryptStreamClose(hStream NCRYPT_STREAM_HANDLE) (r HRESULT) = ncrypt.NCryptStreamClose
//sys	SignError() (r HRESULT) = mssign32.SignError
//sys	SignerFreeSignerContext(pSignerContext *SIGNER_CONTEXT) (r HRESULT) = mssign32.SignerFreeSignerContext
//sys	SignerSign(pSubjectInfo *SIGNER_SUBJECT_INFO, pSignerCert *SIGNER_CERT, pSignatureInfo *SIGNER_SIGNATURE_INFO, pProviderInfo *SIGNER_PROVIDER_INFO, pwszHttpTimeStamp *PWSTRElement, psRequest *CRYPT_ATTRIBUTES, pSipData unsafe.Pointer) (r HRESULT) = mssign32.SignerSign
//sys	SignerSignEx(dwFlags SIGNER_SIGN_FLAGS, pSubjectInfo *SIGNER_SUBJECT_INFO, pSignerCert *SIGNER_CERT, pSignatureInfo *SIGNER_SIGNATURE_INFO, pProviderInfo *SIGNER_PROVIDER_INFO, pwszHttpTimeStamp *PWSTRElement, psRequest *CRYPT_ATTRIBUTES, pSipData unsafe.Pointer, ppSignerContext **SIGNER_CONTEXT) (r HRESULT) = mssign32.SignerSignEx
//sys	SignerSignEx2(dwFlags SIGNER_SIGN_FLAGS, pSubjectInfo *SIGNER_SUBJECT_INFO, pSignerCert *SIGNER_CERT, pSignatureInfo *SIGNER_SIGNATURE_INFO, pProviderInfo *SIGNER_PROVIDER_INFO, dwTimestampFlags SIGNER_TIMESTAMP_FLAGS, pszTimestampAlgorithmOid *PSTRElement, pwszHttpTimeStamp *PWSTRElement, psRequest *CRYPT_ATTRIBUTES, pSipData unsafe.Pointer, ppSignerContext **SIGNER_CONTEXT, pCryptoPolicy *CERT_STRONG_SIGN_PARA, pReserved unsafe.Pointer) (r HRESULT) = mssign32.SignerSignEx2
//sys	SignerSignEx3(dwFlags SIGNER_SIGN_FLAGS, pSubjectInfo *SIGNER_SUBJECT_INFO, pSignerCert *SIGNER_CERT, pSignatureInfo *SIGNER_SIGNATURE_INFO, pProviderInfo *SIGNER_PROVIDER_INFO, dwTimestampFlags SIGNER_TIMESTAMP_FLAGS, pszTimestampAlgorithmOid *PSTRElement, pwszHttpTimeStamp *PWSTRElement, psRequest *CRYPT_ATTRIBUTES, pSipData unsafe.Pointer, ppSignerContext **SIGNER_CONTEXT, pCryptoPolicy *CERT_STRONG_SIGN_PARA, pDigestSignInfo *SIGNER_DIGEST_SIGN_INFO, pReserved unsafe.Pointer) (r HRESULT) = mssign32.SignerSignEx3
//sys	SignerTimeStamp(pSubjectInfo *SIGNER_SUBJECT_INFO, pwszHttpTimeStamp *PWSTRElement, psRequest *CRYPT_ATTRIBUTES, pSipData unsafe.Pointer) (r HRESULT) = mssign32.SignerTimeStamp
//sys	SignerTimeStampEx(dwFlags uint32, pSubjectInfo *SIGNER_SUBJECT_INFO, pwszHttpTimeStamp *PWSTRElement, psRequest *CRYPT_ATTRIBUTES, pSipData unsafe.Pointer, ppSignerContext **SIGNER_CONTEXT) (r HRESULT) = mssign32.SignerTimeStampEx
//sys	SignerTimeStampEx2(dwFlags SIGNER_TIMESTAMP_FLAGS, pSubjectInfo *SIGNER_SUBJECT_INFO, pwszHttpTimeStamp *PWSTRElement, dwAlgId ALG_ID, psRequest *CRYPT_ATTRIBUTES, pSipData unsafe.Pointer, ppSignerContext **SIGNER_CONTEXT) (r HRESULT) = mssign32.SignerTimeStampEx2
//sys	SignerTimeStampEx3(dwFlags SIGNER_TIMESTAMP_FLAGS, dwIndex uint32, pSubjectInfo *SIGNER_SUBJECT_INFO, pwszHttpTimeStamp *PWSTRElement, pszAlgorithmOid *PWSTRElement, psRequest *CRYPT_ATTRIBUTES, pSipData unsafe.Pointer, ppSignerContext **SIGNER_CONTEXT, pCryptoPolicy *CERT_STRONG_SIGN_PARA, pReserved unsafe.Pointer) (r HRESULT) = mssign32.SignerTimeStampEx3
//sys	CryptXmlClose(hCryptXml unsafe.Pointer) (r HRESULT) = cryptxml.CryptXmlClose
//sys	CryptXmlGetTransforms(ppConfig **CRYPT_XML_TRANSFORM_CHAIN_CONFIG) (r HRESULT) = cryptxml.CryptXmlGetTransforms
//sys	CryptXmlOpenToEncode(pConfig *CRYPT_XML_TRANSFORM_CHAIN_CONFIG, dwFlags CRYPT_XML_FLAGS, wszId *PWSTRElement, rgProperty *CRYPT_XML_PROPERTY, cProperty uint32, pEncoded *CRYPT_XML_BLOB, phSignature *unsafe.Pointer) (r HRESULT) = cryptxml.CryptXmlOpenToEncode
//sys	CryptXmlOpenToDecode(pConfig *CRYPT_XML_TRANSFORM_CHAIN_CONFIG, dwFlags CRYPT_XML_FLAGS, rgProperty *CRYPT_XML_PROPERTY, cProperty uint32, pEncoded *CRYPT_XML_BLOB, phCryptXml *unsafe.Pointer) (r HRESULT) = cryptxml.CryptXmlOpenToDecode
//sys	CryptXmlAddObject(hSignatureOrObject unsafe.Pointer, dwFlags uint32, rgProperty *CRYPT_XML_PROPERTY, cProperty uint32, pEncoded *CRYPT_XML_BLOB, ppObject **CRYPT_XML_OBJECT) (r HRESULT) = cryptxml.CryptXmlAddObject
//sys	CryptXmlCreateReference(hCryptXml unsafe.Pointer, dwFlags uint32, wszId *PWSTRElement, wszURI *PWSTRElement, wszType *PWSTRElement, pDigestMethod *CRYPT_XML_ALGORITHM, cTransform uint32, rgTransform *CRYPT_XML_ALGORITHM, phReference *unsafe.Pointer) (r HRESULT) = cryptxml.CryptXmlCreateReference
//sys	CryptXmlDigestReference(hReference unsafe.Pointer, dwFlags uint32, pDataProviderIn *CRYPT_XML_DATA_PROVIDER) (r HRESULT) = cryptxml.CryptXmlDigestReference
//sys	CryptXmlSetHMACSecret(hSignature unsafe.Pointer, pbSecret *uint8, cbSecret uint32) (r HRESULT) = cryptxml.CryptXmlSetHMACSecret
//sys	CryptXmlSign(hSignature unsafe.Pointer, hKey HCRYPTPROV_OR_NCRYPT_KEY_HANDLE, dwKeySpec CERT_KEY_SPEC, dwFlags CRYPT_XML_FLAGS, dwKeyInfoSpec CRYPT_XML_KEYINFO_SPEC, pvKeyInfoSpec unsafe.Pointer, pSignatureMethod *CRYPT_XML_ALGORITHM, pCanonicalization *CRYPT_XML_ALGORITHM) (r HRESULT) = cryptxml.CryptXmlSign
//sys	CryptXmlImportPublicKey(dwFlags CRYPT_XML_FLAGS, pKeyValue *CRYPT_XML_KEY_VALUE, phKey *BCRYPT_KEY_HANDLE) (r HRESULT) = cryptxml.CryptXmlImportPublicKey
//sys	CryptXmlVerifySignature(hSignature unsafe.Pointer, hKey BCRYPT_KEY_HANDLE, dwFlags CRYPT_XML_FLAGS) (r HRESULT) = cryptxml.CryptXmlVerifySignature
//sys	CryptXmlGetDocContext(hCryptXml unsafe.Pointer, ppStruct **CRYPT_XML_DOC_CTXT) (r HRESULT) = cryptxml.CryptXmlGetDocContext
//sys	CryptXmlGetSignature(hCryptXml unsafe.Pointer, ppStruct **CRYPT_XML_SIGNATURE) (r HRESULT) = cryptxml.CryptXmlGetSignature
//sys	CryptXmlGetReference(hCryptXml unsafe.Pointer, ppStruct **CRYPT_XML_REFERENCE) (r HRESULT) = cryptxml.CryptXmlGetReference
//sys	CryptXmlGetStatus(hCryptXml unsafe.Pointer, pStatus *CRYPT_XML_STATUS) (r HRESULT) = cryptxml.CryptXmlGetStatus
//sys	CryptXmlEncode(hCryptXml unsafe.Pointer, dwCharset CRYPT_XML_CHARSET, rgProperty *CRYPT_XML_PROPERTY, cProperty uint32, pvCallbackState unsafe.Pointer, pfnWrite PFN_CRYPT_XML_WRITE_CALLBACK) (r HRESULT) = cryptxml.CryptXmlEncode
//sys	CryptXmlGetAlgorithmInfo(pXmlAlgorithm *CRYPT_XML_ALGORITHM, dwFlags CRYPT_XML_FLAGS, ppAlgInfo **CRYPT_XML_ALGORITHM_INFO) (r HRESULT) = cryptxml.CryptXmlGetAlgorithmInfo
//sys	CryptXmlFindAlgorithmInfo(dwFindByType uint32, pvFindBy unsafe.Pointer, dwGroupId uint32, dwFlags uint32) (r *CRYPT_XML_ALGORITHM_INFO) = cryptxml.CryptXmlFindAlgorithmInfo
//sys	CryptXmlEnumAlgorithmInfo(dwGroupId uint32, dwFlags uint32, pvArg unsafe.Pointer, pfnEnumAlgInfo PFN_CRYPT_XML_ENUM_ALG_INFO) (r HRESULT) = cryptxml.CryptXmlEnumAlgorithmInfo
//sys	GetToken(cPolicyChain uint32, pPolicyChain *POLICY_ELEMENT, securityToken **GENERIC_XML_TOKEN, phProofTokenCrypto **INFORMATIONCARD_CRYPTO_HANDLE) (r HRESULT) = infocardapi.GetToken
//sys	ManageCardSpace() (r HRESULT) = infocardapi.ManageCardSpace
//sys	ImportInformationCard(fileName *PWSTRElement) (r HRESULT) = infocardapi.ImportInformationCard
//sys	Encrypt(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, fOAEP BOOL, cbInData uint32, pInData *uint8, pcbOutData *uint32, ppOutData **uint8) (r HRESULT) = infocardapi.Encrypt
//sys	Decrypt(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, fOAEP BOOL, cbInData uint32, pInData *uint8, pcbOutData *uint32, ppOutData **uint8) (r HRESULT) = infocardapi.Decrypt
//sys	SignHash(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, cbHash uint32, pHash *uint8, hashAlgOid *PWSTRElement, pcbSig *uint32, ppSig **uint8) (r HRESULT) = infocardapi.SignHash
//sys	VerifyHash(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, cbHash uint32, pHash *uint8, hashAlgOid *PWSTRElement, cbSig uint32, pSig *uint8, pfVerified *BOOL) (r HRESULT) = infocardapi.VerifyHash
//sys	GetCryptoTransform(hSymmetricCrypto *INFORMATIONCARD_CRYPTO_HANDLE, mode uint32, padding PaddingMode, feedbackSize uint32, direction Direction, cbIV uint32, pIV *uint8, pphTransform **INFORMATIONCARD_CRYPTO_HANDLE) (r HRESULT) = infocardapi.GetCryptoTransform
//sys	GetKeyedHash(hSymmetricCrypto *INFORMATIONCARD_CRYPTO_HANDLE, pphHash **INFORMATIONCARD_CRYPTO_HANDLE) (r HRESULT) = infocardapi.GetKeyedHash
//sys	TransformBlock(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, cbInData uint32, pInData *uint8, pcbOutData *uint32, ppOutData **uint8) (r HRESULT) = infocardapi.TransformBlock
//sys	TransformFinalBlock(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, cbInData uint32, pInData *uint8, pcbOutData *uint32, ppOutData **uint8) (r HRESULT) = infocardapi.TransformFinalBlock
//sys	HashCore(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, cbInData uint32, pInData *uint8) (r HRESULT) = infocardapi.HashCore
//sys	HashFinal(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, cbInData uint32, pInData *uint8, pcbOutData *uint32, ppOutData **uint8) (r HRESULT) = infocardapi.HashFinal
//sys	FreeToken(pAllocMemory *GENERIC_XML_TOKEN) (r BOOL) = infocardapi.FreeToken
//sys	CloseCryptoHandle(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE) (r HRESULT) = infocardapi.CloseCryptoHandle
//sys	GenerateDerivedKey(hCrypto *INFORMATIONCARD_CRYPTO_HANDLE, cbLabel uint32, pLabel *uint8, cbNonce uint32, pNonce *uint8, derivedKeyLength uint32, offset uint32, algId *PWSTRElement, pcbKey *uint32, ppKey **uint8) (r HRESULT) = infocardapi.GenerateDerivedKey
//sys	GetBrowserToken(dwParamType uint32, pParam unsafe.Pointer, pcbToken *uint32, ppToken **uint8) (r HRESULT) = infocardapi.GetBrowserToken
//sys	GetCipherInterface(pszProviderName *PWSTRElement, pszAlgId *PWSTRElement, ppFunctionTable **BCRYPT_CIPHER_FUNCTION_TABLE, dwFlags uint32) (r NTSTATUS) = bcryptprimitives.GetCipherInterface
//sys	GetHashInterface(pszProviderName *PWSTRElement, pszAlgId *PWSTRElement, ppFunctionTable **BCRYPT_HASH_FUNCTION_TABLE, dwFlags uint32) (r NTSTATUS) = bcryptprimitives.GetHashInterface
//sys	GetAsymmetricEncryptionInterface(pszProviderName *PWSTRElement, pszAlgId *PWSTRElement, ppFunctionTable **BCRYPT_ASYMMETRIC_ENCRYPTION_FUNCTION_TABLE, dwFlags uint32) (r NTSTATUS) = bcryptprimitives.GetAsymmetricEncryptionInterface
//sys	GetSecretAgreementInterface(pszProviderName *PWSTRElement, pszAlgId *PWSTRElement, ppFunctionTable **BCRYPT_SECRET_AGREEMENT_FUNCTION_TABLE, dwFlags uint32) (r NTSTATUS) = bcryptprimitives.GetSecretAgreementInterface
//sys	GetSignatureInterface(pszProviderName *PWSTRElement, pszAlgId *PWSTRElement, ppFunctionTable **BCRYPT_SIGNATURE_FUNCTION_TABLE, dwFlags uint32) (r NTSTATUS) = bcryptprimitives.GetSignatureInterface
//sys	GetRngInterface(pszProviderName *PWSTRElement, ppFunctionTable **BCRYPT_RNG_FUNCTION_TABLE, dwFlags uint32) (r NTSTATUS) = bcryptprimitives.GetRngInterface
//sys	GetKeyDerivationInterface(pszProviderName *PWSTRElement, pszAlgId *PWSTRElement, ppFunctionTable **BCRYPT_KEY_DERIVATION_FUNCTION_TABLE, dwFlags uint32) (r NTSTATUS) = bcryptprimitives.GetKeyDerivationInterface
//sys	BCryptRegisterProvider(pszProvider *PWSTRElement, dwFlags uint32, pReg *CRYPT_PROVIDER_REG) (r NTSTATUS) = bcrypt.BCryptRegisterProvider
//sys	BCryptUnregisterProvider(pszProvider *PWSTRElement) (r NTSTATUS) = bcrypt.BCryptUnregisterProvider
//sys	BCryptAddContextFunctionProvider(dwTable uint32, pszContext *PWSTRElement, dwInterface uint32, pszFunction *PWSTRElement, pszProvider *PWSTRElement, dwPosition uint32) (r NTSTATUS) = bcrypt.BCryptAddContextFunctionProvider
//sys	BCryptRemoveContextFunctionProvider(dwTable uint32, pszContext *PWSTRElement, dwInterface uint32, pszFunction *PWSTRElement, pszProvider *PWSTRElement) (r NTSTATUS) = bcrypt.BCryptRemoveContextFunctionProvider
//sys	GetKeyStorageInterface(pszProviderName *PWSTRElement, ppFunctionTable **NCRYPT_KEY_STORAGE_FUNCTION_TABLE, dwFlags uint32) (r NTSTATUS) = ncrypt.GetKeyStorageInterface
//sys	SslChangeNotify(hEvent HANDLE, dwFlags uint32) (r HRESULT) = ncrypt.SslChangeNotify
//sys	SslComputeClientAuthHash(hSslProvider NCRYPT_PROV_HANDLE, hMasterKey NCRYPT_KEY_HANDLE, hHandshakeHash NCRYPT_HASH_HANDLE, pszAlgId *PWSTRElement, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslComputeClientAuthHash
//sys	SslComputeEapKeyBlock(hSslProvider NCRYPT_PROV_HANDLE, hMasterKey NCRYPT_KEY_HANDLE, pbRandoms *uint8, cbRandoms uint32, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslComputeEapKeyBlock
//sys	SslComputeFinishedHash(hSslProvider NCRYPT_PROV_HANDLE, hMasterKey NCRYPT_KEY_HANDLE, hHandshakeHash NCRYPT_HASH_HANDLE, pbOutput *uint8, cbOutput uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslComputeFinishedHash
//sys	SslCreateEphemeralKey(hSslProvider NCRYPT_PROV_HANDLE, phEphemeralKey *NCRYPT_KEY_HANDLE, dwProtocol uint32, dwCipherSuite uint32, dwKeyType uint32, dwKeyBitLen uint32, pbParams *uint8, cbParams uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslCreateEphemeralKey
//sys	SslCreateHandshakeHash(hSslProvider NCRYPT_PROV_HANDLE, phHandshakeHash *NCRYPT_HASH_HANDLE, dwProtocol uint32, dwCipherSuite uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslCreateHandshakeHash
//sys	SslDecryptPacket(hSslProvider NCRYPT_PROV_HANDLE, hKey NCRYPT_KEY_HANDLE, pbInput *uint8, cbInput uint32, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, SequenceNumber uint64, dwFlags uint32) (r HRESULT) = ncrypt.SslDecryptPacket
//sys	SslEncryptPacket(hSslProvider NCRYPT_PROV_HANDLE, hKey NCRYPT_KEY_HANDLE, pbInput *uint8, cbInput uint32, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, SequenceNumber uint64, dwContentType uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslEncryptPacket
//sys	SslEnumCipherSuites(hSslProvider NCRYPT_PROV_HANDLE, hPrivateKey NCRYPT_KEY_HANDLE, ppCipherSuite **NCRYPT_SSL_CIPHER_SUITE, ppEnumState *unsafe.Pointer, dwFlags uint32) (r HRESULT) = ncrypt.SslEnumCipherSuites
//sys	SslEnumCipherSuitesEx(hSslProvider NCRYPT_PROV_HANDLE, hPrivateKey NCRYPT_KEY_HANDLE, ppCipherSuite **NCRYPT_SSL_CIPHER_SUITE_EX, ppEnumState *unsafe.Pointer, dwFlags uint32) (r HRESULT) = ncrypt.SslEnumCipherSuitesEx
//sys	SslEnumEccCurves(hSslProvider NCRYPT_PROV_HANDLE, pEccCurveCount *uint32, ppEccCurve **NCRYPT_SSL_ECC_CURVE, dwFlags uint32) (r HRESULT) = ncrypt.SslEnumEccCurves
//sys	SslEnumProtocolProviders(pdwProviderCount *uint32, ppProviderList **NCryptProviderName, dwFlags uint32) (r HRESULT) = ncrypt.SslEnumProtocolProviders
//sys	SslExportKey(hSslProvider NCRYPT_PROV_HANDLE, hKey NCRYPT_KEY_HANDLE, pszBlobType *PWSTRElement, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslExportKey
//sys	SslFreeBuffer(pvInput unsafe.Pointer) (r HRESULT) = ncrypt.SslFreeBuffer
//sys	SslFreeObject(hObject NCRYPT_HANDLE, dwFlags uint32) (r HRESULT) = ncrypt.SslFreeObject
//sys	SslGenerateMasterKey(hSslProvider NCRYPT_PROV_HANDLE, hPrivateKey NCRYPT_KEY_HANDLE, hPublicKey NCRYPT_KEY_HANDLE, phMasterKey *NCRYPT_KEY_HANDLE, dwProtocol uint32, dwCipherSuite uint32, pParameterList *BCryptBufferDesc, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslGenerateMasterKey
//sys	SslGenerateSessionKeys(hSslProvider NCRYPT_PROV_HANDLE, hMasterKey NCRYPT_KEY_HANDLE, phReadKey *NCRYPT_KEY_HANDLE, phWriteKey *NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslGenerateSessionKeys
//sys	SslGetKeyProperty(hKey NCRYPT_KEY_HANDLE, pszProperty *PWSTRElement, ppbOutput **uint8, pcbOutput *uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslGetKeyProperty
//sys	SslGetProviderProperty(hSslProvider NCRYPT_PROV_HANDLE, pszProperty *PWSTRElement, ppbOutput **uint8, pcbOutput *uint32, ppEnumState *unsafe.Pointer, dwFlags uint32) (r HRESULT) = ncrypt.SslGetProviderProperty
//sys	SslHashHandshake(hSslProvider NCRYPT_PROV_HANDLE, hHandshakeHash NCRYPT_HASH_HANDLE, pbInput *uint8, cbInput uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslHashHandshake
//sys	SslImportKey(hSslProvider NCRYPT_PROV_HANDLE, phKey *NCRYPT_KEY_HANDLE, pszBlobType *PWSTRElement, pbKeyBlob *uint8, cbKeyBlob uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslImportKey
//sys	SslImportMasterKey(hSslProvider NCRYPT_PROV_HANDLE, hPrivateKey NCRYPT_KEY_HANDLE, phMasterKey *NCRYPT_KEY_HANDLE, dwProtocol uint32, dwCipherSuite uint32, pParameterList *BCryptBufferDesc, pbEncryptedKey *uint8, cbEncryptedKey uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslImportMasterKey
//sys	SslLookupCipherSuiteInfo(hSslProvider NCRYPT_PROV_HANDLE, dwProtocol uint32, dwCipherSuite uint32, dwKeyType uint32, pCipherSuite *NCRYPT_SSL_CIPHER_SUITE, dwFlags uint32) (r HRESULT) = ncrypt.SslLookupCipherSuiteInfo
//sys	SslOpenPrivateKey(hSslProvider NCRYPT_PROV_HANDLE, phPrivateKey *NCRYPT_KEY_HANDLE, pCertContext *CERT_CONTEXT, dwFlags uint32) (r HRESULT) = ncrypt.SslOpenPrivateKey
//sys	SslOpenProvider(phSslProvider *NCRYPT_PROV_HANDLE, pszProviderName *PWSTRElement, dwFlags uint32) (r HRESULT) = ncrypt.SslOpenProvider
//sys	SslSignHash(hSslProvider NCRYPT_PROV_HANDLE, hPrivateKey NCRYPT_KEY_HANDLE, pbHashValue *uint8, cbHashValue uint32, pbSignature *uint8, cbSignature uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslSignHash
//sys	SslVerifySignature(hSslProvider NCRYPT_PROV_HANDLE, hPublicKey NCRYPT_KEY_HANDLE, pbHashValue *uint8, cbHashValue uint32, pbSignature *uint8, cbSignature uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslVerifySignature
//sys	SslLookupCipherLengths(hSslProvider NCRYPT_PROV_HANDLE, dwProtocol uint32, dwCipherSuite uint32, dwKeyType uint32, pCipherLengths *NCRYPT_SSL_CIPHER_LENGTHS, cbCipherLengths uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslLookupCipherLengths
//sys	SslCreateClientAuthHash(hSslProvider NCRYPT_PROV_HANDLE, phHandshakeHash *NCRYPT_HASH_HANDLE, dwProtocol uint32, dwCipherSuite uint32, pszHashAlgId *PWSTRElement, dwFlags uint32) (r HRESULT) = ncrypt.SslCreateClientAuthHash
//sys	SslGetCipherSuitePRFHashAlgorithm(hSslProvider NCRYPT_PROV_HANDLE, dwProtocol uint32, dwCipherSuite uint32, dwKeyType uint32, szPRFHash *PWSTRElement, dwFlags uint32) (r HRESULT) = ncrypt.SslGetCipherSuitePRFHashAlgorithm
//sys	SslComputeSessionHash(hSslProvider NCRYPT_PROV_HANDLE, hHandshakeHash NCRYPT_HASH_HANDLE, dwProtocol uint32, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslComputeSessionHash
//sys	SslGeneratePreMasterKey(hSslProvider NCRYPT_PROV_HANDLE, hPublicKey NCRYPT_KEY_HANDLE, phPreMasterKey *NCRYPT_KEY_HANDLE, dwProtocol uint32, dwCipherSuite uint32, pParameterList *BCryptBufferDesc, pbOutput *uint8, cbOutput uint32, pcbResult *uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslGeneratePreMasterKey
//sys	SslExportKeyingMaterial(hSslProvider NCRYPT_PROV_HANDLE, hMasterKey NCRYPT_KEY_HANDLE, sLabel *PSTRElement, pbRandoms *uint8, cbRandoms uint32, pbContextValue *uint8, cbContextValue uint16, pbOutput *uint8, cbOutput uint32, dwFlags uint32) (r HRESULT) = ncrypt.SslExportKeyingMaterial
//sys	SslExtractEarlyKey(hSslProvider NCRYPT_PROV_HANDLE, hPreSharedKey NCRYPT_KEY_HANDLE, phEarlyKey *NCRYPT_KEY_HANDLE, dwProtocol uint32, dwCipherSuite uint32, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslExtractEarlyKey
//sys	SslExtractHandshakeKey(hSslProvider NCRYPT_PROV_HANDLE, hPrivateKey NCRYPT_KEY_HANDLE, hPublicKey NCRYPT_KEY_HANDLE, hEarlyKey NCRYPT_KEY_HANDLE, phHandshakeKey *NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslExtractHandshakeKey
//sys	SslExtractMasterKey(hSslProvider NCRYPT_PROV_HANDLE, hHandshakeKey NCRYPT_KEY_HANDLE, phMasterKey *NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslExtractMasterKey
//sys	SslExpandTrafficKeys(hSslProvider NCRYPT_PROV_HANDLE, hBaseKey NCRYPT_KEY_HANDLE, hHashValue NCRYPT_HASH_HANDLE, phClientTrafficKey *NCRYPT_KEY_HANDLE, phServerTrafficKey *NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslExpandTrafficKeys
//sys	SslExpandWriteKey(hSslProvider NCRYPT_PROV_HANDLE, hBaseTrafficKey NCRYPT_KEY_HANDLE, phWriteKey *NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslExpandWriteKey
//sys	SslExpandExporterMasterKey(hSslProvider NCRYPT_PROV_HANDLE, hBaseKey NCRYPT_KEY_HANDLE, hHashValue NCRYPT_HASH_HANDLE, phExporterMasterKey *NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslExpandExporterMasterKey
//sys	SslExpandResumptionMasterKey(hSslProvider NCRYPT_PROV_HANDLE, hMasterKey NCRYPT_KEY_HANDLE, hHashValue NCRYPT_HASH_HANDLE, phResumptionMasterKey *NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslExpandResumptionMasterKey
//sys	SslDuplicateTranscriptHash(hSslProvider NCRYPT_PROV_HANDLE, hTranscriptHash NCRYPT_HASH_HANDLE, phTranscriptHash *NCRYPT_HASH_HANDLE, dwFlags uint32) (r HRESULT) = ncrypt.SslDuplicateTranscriptHash
//sys	SslExpandBinderKey(hSslProvider NCRYPT_PROV_HANDLE, hEarlyKey NCRYPT_KEY_HANDLE, phBinderKey *NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslExpandBinderKey
//sys	SslExpandPreSharedKey(hSslProvider NCRYPT_PROV_HANDLE, hResumptionMasterKey NCRYPT_KEY_HANDLE, pbTicketNonce *uint8, cbTicketNonce uint32, phPreSharedKey *NCRYPT_KEY_HANDLE, pParameterList *BCryptBufferDesc, dwFlags uint32) (r HRESULT) = ncrypt.SslExpandPreSharedKey
//sys	GetSChannelInterface(pszProviderName *PWSTRElement, ppFunctionTable **NCRYPT_SSL_FUNCTION_TABLE, dwFlags uint32) (r NTSTATUS) = ncrypt.GetSChannelInterface
//sys	SslIncrementProviderReferenceCount(hSslProvider NCRYPT_PROV_HANDLE) (r HRESULT) = ncrypt.SslIncrementProviderReferenceCount
//sys	SslDecrementProviderReferenceCount(hSslProvider NCRYPT_PROV_HANDLE) (r HRESULT) = ncrypt.SslDecrementProviderReferenceCount

// APIs for Windows.Win32.System.Diagnostics.Debug
//sys	ReadProcessMemory(hProcess HANDLE, lpBaseAddress unsafe.Pointer, lpBuffer unsafe.Pointer, nSize uintptr, lpNumberOfBytesRead *uintptr) (r BOOL, err error)
//sys	WriteProcessMemory(hProcess HANDLE, lpBaseAddress unsafe.Pointer, lpBuffer unsafe.Pointer, nSize uintptr, lpNumberOfBytesWritten *uintptr) (r BOOL, err error)
//sys	GetThreadContext(hThread HANDLE, lpContext *CONTEXT) (r BOOL, err error)
//sys	SetThreadContext(hThread HANDLE, lpContext *CONTEXT) (r BOOL, err error)
//sys	FlushInstructionCache(hProcess HANDLE, lpBaseAddress unsafe.Pointer, dwSize uintptr) (r BOOL, err error)
//sys	Wow64GetThreadContext(hThread HANDLE, lpContext *WOW64_CONTEXT) (r BOOL, err error)
//sys	Wow64SetThreadContext(hThread HANDLE, lpContext *WOW64_CONTEXT) (r BOOL, err error)
//sys	RtlCaptureStackBackTrace(FramesToSkip uint32, FramesToCapture uint32, BackTrace *unsafe.Pointer, BackTraceHash *uint32) (r uint16)
//sys	RtlCaptureContext(ContextRecord *CONTEXT)
//sys	RtlUnwind(TargetFrame unsafe.Pointer, TargetIp unsafe.Pointer, ExceptionRecord *EXCEPTION_RECORD, ReturnValue unsafe.Pointer)
//sys	RtlRestoreContext(ContextRecord *CONTEXT, ExceptionRecord *EXCEPTION_RECORD)
//sys	RtlRaiseException(ExceptionRecord *EXCEPTION_RECORD)
//sys	RtlPcToFileHeader(PcValue unsafe.Pointer, BaseOfImage *unsafe.Pointer) (r unsafe.Pointer)
//sys	IsDebuggerPresent() (r BOOL)
//sys	DebugBreak()
//sys	OutputDebugStringA(lpOutputString *PSTRElement)
//sys	OutputDebugStringW(lpOutputString *PWSTRElement)
//sys	ContinueDebugEvent(dwProcessId uint32, dwThreadId uint32, dwContinueStatus NTSTATUS) (r BOOL, err error)
//sys	WaitForDebugEvent(lpDebugEvent *DEBUG_EVENT, dwMilliseconds uint32) (r BOOL, err error)
//sys	DebugActiveProcess(dwProcessId uint32) (r BOOL, err error)
//sys	DebugActiveProcessStop(dwProcessId uint32) (r BOOL, err error)
//sys	CheckRemoteDebuggerPresent(hProcess HANDLE, pbDebuggerPresent *BOOL) (r BOOL, err error)
//sys	WaitForDebugEventEx(lpDebugEvent *DEBUG_EVENT, dwMilliseconds uint32) (r BOOL, err error)
//sys	EncodePointer(Ptr unsafe.Pointer) (r unsafe.Pointer)
//sys	DecodePointer(Ptr unsafe.Pointer) (r unsafe.Pointer)
//sys	EncodeSystemPointer(Ptr unsafe.Pointer) (r unsafe.Pointer)
//sys	DecodeSystemPointer(Ptr unsafe.Pointer) (r unsafe.Pointer)
//sys	EncodeRemotePointer(ProcessHandle HANDLE, Ptr unsafe.Pointer, EncodedPtr *unsafe.Pointer) (r HRESULT) = api-ms-win-core-util-l1-1-1.EncodeRemotePointer
//sys	DecodeRemotePointer(ProcessHandle HANDLE, Ptr unsafe.Pointer, DecodedPtr *unsafe.Pointer) (r HRESULT) = api-ms-win-core-util-l1-1-1.DecodeRemotePointer
//sys	Beep(dwFreq uint32, dwDuration uint32) (r BOOL, err error)
//sys	RaiseException(dwExceptionCode uint32, dwExceptionFlags uint32, nNumberOfArguments uint32, lpArguments *uintptr)
//sys	UnhandledExceptionFilter(ExceptionInfo *EXCEPTION_POINTERS) (r int32)
//sys	SetUnhandledExceptionFilter(lpTopLevelExceptionFilter LPTOP_LEVEL_EXCEPTION_FILTER) (r LPTOP_LEVEL_EXCEPTION_FILTER)
//sys	GetErrorMode() (r uint32)
//sys	SetErrorMode(uMode THREAD_ERROR_MODE) (r THREAD_ERROR_MODE)
//sys	AddVectoredExceptionHandler(First uint32, Handler PVECTORED_EXCEPTION_HANDLER) (r unsafe.Pointer)
//sys	RemoveVectoredExceptionHandler(Handle unsafe.Pointer) (r uint32)
//sys	AddVectoredContinueHandler(First uint32, Handler PVECTORED_EXCEPTION_HANDLER) (r unsafe.Pointer)
//sys	RemoveVectoredContinueHandler(Handle unsafe.Pointer) (r uint32)
//sys	RaiseFailFastException(pExceptionRecord *EXCEPTION_RECORD, pContextRecord *CONTEXT, dwFlags uint32)
//sys	FatalAppExitA(uAction uint32, lpMessageText *PSTRElement)
//sys	FatalAppExitW(uAction uint32, lpMessageText *PWSTRElement)
//sys	GetThreadErrorMode() (r uint32)
//sys	SetThreadErrorMode(dwNewMode THREAD_ERROR_MODE, lpOldMode *THREAD_ERROR_MODE) (r BOOL, err error)
//sys	TerminateProcessOnMemoryExhaustion(FailedAllocationSize uintptr) = api-ms-win-core-errorhandling-l1-1-3.TerminateProcessOnMemoryExhaustion
//sys	OpenThreadWaitChainSession(Flags OPEN_THREAD_WAIT_CHAIN_SESSION_FLAGS, callback PWAITCHAINCALLBACK) (r unsafe.Pointer, err error) = advapi32.OpenThreadWaitChainSession
//sys	CloseThreadWaitChainSession(WctHandle unsafe.Pointer) = advapi32.CloseThreadWaitChainSession
//sys	GetThreadWaitChain(WctHandle unsafe.Pointer, Context uintptr, Flags WAIT_CHAIN_THREAD_OPTIONS, ThreadId uint32, NodeCount *uint32, NodeInfoArray *WAITCHAIN_NODE_INFO, IsCycle *BOOL) (r BOOL, err error) = advapi32.GetThreadWaitChain
//sys	RegisterWaitChainCOMCallback(CallStateCallback PCOGETCALLSTATE, ActivationStateCallback PCOGETACTIVATIONSTATE) = advapi32.RegisterWaitChainCOMCallback
//sys	MiniDumpWriteDump(hProcess HANDLE, ProcessId uint32, hFile HANDLE, DumpType MINIDUMP_TYPE, ExceptionParam *MINIDUMP_EXCEPTION_INFORMATION, UserStreamParam *MINIDUMP_USER_STREAM_INFORMATION, CallbackParam *MINIDUMP_CALLBACK_INFORMATION) (r BOOL, err error) = dbghelp.MiniDumpWriteDump
//sys	MiniDumpReadDumpStream(BaseOfDump unsafe.Pointer, StreamNumber uint32, Dir **MINIDUMP_DIRECTORY, StreamPointer *unsafe.Pointer, StreamSize *uint32) (r BOOL) = dbghelp.MiniDumpReadDumpStream
//sys	BindImage(ImageName *PSTRElement, DllPath *PSTRElement, SymbolPath *PSTRElement) (r BOOL, err error) = imagehlp.BindImage
//sys	BindImageEx(Flags uint32, ImageName *PSTRElement, DllPath *PSTRElement, SymbolPath *PSTRElement, StatusRoutine PIMAGEHLP_STATUS_ROUTINE) (r BOOL, err error) = imagehlp.BindImageEx
//sys	ReBaseImage(CurrentImageName *PSTRElement, SymbolPath *PSTRElement, fReBase BOOL, fRebaseSysfileOk BOOL, fGoingDown BOOL, CheckImageSize uint32, OldImageSize *uint32, OldImageBase *uintptr, NewImageSize *uint32, NewImageBase *uintptr, TimeStamp uint32) (r BOOL, err error) = imagehlp.ReBaseImage
//sys	ReBaseImage64(CurrentImageName *PSTRElement, SymbolPath *PSTRElement, fReBase BOOL, fRebaseSysfileOk BOOL, fGoingDown BOOL, CheckImageSize uint32, OldImageSize *uint32, OldImageBase *uint64, NewImageSize *uint32, NewImageBase *uint64, TimeStamp uint32) (r BOOL, err error) = imagehlp.ReBaseImage64
//sys	MapFileAndCheckSumA(Filename *PSTRElement, HeaderSum *uint32, CheckSum *uint32) (r uint32) = imagehlp.MapFileAndCheckSumA
//sys	MapFileAndCheckSumW(Filename *PWSTRElement, HeaderSum *uint32, CheckSum *uint32) (r uint32) = imagehlp.MapFileAndCheckSumW
//sys	GetImageUnusedHeaderBytes(LoadedImage *LOADED_IMAGE, SizeUnusedHeaderBytes *uint32) (r uint32, err error) = imagehlp.GetImageUnusedHeaderBytes
//sys	ImageGetDigestStream(FileHandle HANDLE, DigestLevel uint32, DigestFunction DIGEST_FUNCTION, DigestHandle unsafe.Pointer) (r BOOL, err error) = imagehlp.ImageGetDigestStream
//sys	ImageAddCertificate(FileHandle HANDLE, Certificate *WIN_CERTIFICATE, Index *uint32) (r BOOL, err error) = imagehlp.ImageAddCertificate
//sys	ImageRemoveCertificate(FileHandle HANDLE, Index uint32) (r BOOL, err error) = imagehlp.ImageRemoveCertificate
//sys	ImageEnumerateCertificates(FileHandle HANDLE, TypeFilter uint16, CertificateCount *uint32, Indices *uint32, IndexCount uint32) (r BOOL, err error) = imagehlp.ImageEnumerateCertificates
//sys	ImageGetCertificateData(FileHandle HANDLE, CertificateIndex uint32, Certificate *WIN_CERTIFICATE, RequiredLength *uint32) (r BOOL, err error) = imagehlp.ImageGetCertificateData
//sys	ImageGetCertificateHeader(FileHandle HANDLE, CertificateIndex uint32, Certificateheader *WIN_CERTIFICATE) (r BOOL, err error) = imagehlp.ImageGetCertificateHeader
//sys	ImageLoad(DllName *PSTRElement, DllPath *PSTRElement) (r *LOADED_IMAGE, err error) = imagehlp.ImageLoad
//sys	ImageUnload(LoadedImage *LOADED_IMAGE) (r BOOL, err error) = imagehlp.ImageUnload
//sys	MapAndLoad(ImageName *PSTRElement, DllPath *PSTRElement, LoadedImage *LOADED_IMAGE, DotDll BOOL, ReadOnly BOOL) (r BOOL, err error) = imagehlp.MapAndLoad
//sys	UnMapAndLoad(LoadedImage *LOADED_IMAGE) (r BOOL, err error) = imagehlp.UnMapAndLoad
//sys	TouchFileTimes(FileHandle HANDLE, pSystemTime *SYSTEMTIME) (r BOOL, err error) = imagehlp.TouchFileTimes
//sys	UpdateDebugInfoFile(ImageFileName *PSTRElement, SymbolPath *PSTRElement, DebugFilePath *PSTRElement, NtHeaders *IMAGE_NT_HEADERS32) (r BOOL, err error) = imagehlp.UpdateDebugInfoFile
//sys	UpdateDebugInfoFileEx(ImageFileName *PSTRElement, SymbolPath *PSTRElement, DebugFilePath *PSTRElement, NtHeaders *IMAGE_NT_HEADERS32, OldCheckSum uint32) (r BOOL) = imagehlp.UpdateDebugInfoFileEx
//sys	SymFindDebugInfoFile(hProcess HANDLE, FileName *PSTRElement, DebugFilePath *PSTRElement, Callback PFIND_DEBUG_FILE_CALLBACK, CallerData unsafe.Pointer) (r HANDLE, err error) = dbghelp.SymFindDebugInfoFile
//sys	SymFindDebugInfoFileW(hProcess HANDLE, FileName *PWSTRElement, DebugFilePath *PWSTRElement, Callback PFIND_DEBUG_FILE_CALLBACKW, CallerData unsafe.Pointer) (r HANDLE, err error) = dbghelp.SymFindDebugInfoFileW
//sys	FindDebugInfoFile(FileName *PSTRElement, SymbolPath *PSTRElement, DebugFilePath *PSTRElement) (r HANDLE, err error) = dbghelp.FindDebugInfoFile
//sys	FindDebugInfoFileEx(FileName *PSTRElement, SymbolPath *PSTRElement, DebugFilePath *PSTRElement, Callback PFIND_DEBUG_FILE_CALLBACK, CallerData unsafe.Pointer) (r HANDLE, err error) = dbghelp.FindDebugInfoFileEx
//sys	FindDebugInfoFileExW(FileName *PWSTRElement, SymbolPath *PWSTRElement, DebugFilePath *PWSTRElement, Callback PFIND_DEBUG_FILE_CALLBACKW, CallerData unsafe.Pointer) (r HANDLE, err error) = dbghelp.FindDebugInfoFileExW
//sys	SymFindFileInPath(hprocess HANDLE, SearchPathA *PSTRElement, FileName *PSTRElement, id unsafe.Pointer, two uint32, three uint32, flags SYM_FIND_ID_OPTION, FoundFile *PSTRElement, callback PFINDFILEINPATHCALLBACK, context unsafe.Pointer) (r BOOL, err error) = dbghelp.SymFindFileInPath
//sys	SymFindFileInPathW(hprocess HANDLE, SearchPathA *PWSTRElement, FileName *PWSTRElement, id unsafe.Pointer, two uint32, three uint32, flags SYM_FIND_ID_OPTION, FoundFile *PWSTRElement, callback PFINDFILEINPATHCALLBACKW, context unsafe.Pointer) (r BOOL, err error) = dbghelp.SymFindFileInPathW
//sys	SymFindExecutableImage(hProcess HANDLE, FileName *PSTRElement, ImageFilePath *PSTRElement, Callback PFIND_EXE_FILE_CALLBACK, CallerData unsafe.Pointer) (r HANDLE, err error) = dbghelp.SymFindExecutableImage
//sys	SymFindExecutableImageW(hProcess HANDLE, FileName *PWSTRElement, ImageFilePath *PWSTRElement, Callback PFIND_EXE_FILE_CALLBACKW, CallerData unsafe.Pointer) (r HANDLE, err error) = dbghelp.SymFindExecutableImageW
//sys	FindExecutableImage(FileName *PSTRElement, SymbolPath *PSTRElement, ImageFilePath *PSTRElement) (r HANDLE, err error) = dbghelp.FindExecutableImage
//sys	FindExecutableImageEx(FileName *PSTRElement, SymbolPath *PSTRElement, ImageFilePath *PSTRElement, Callback PFIND_EXE_FILE_CALLBACK, CallerData unsafe.Pointer) (r HANDLE, err error) = dbghelp.FindExecutableImageEx
//sys	FindExecutableImageExW(FileName *PWSTRElement, SymbolPath *PWSTRElement, ImageFilePath *PWSTRElement, Callback PFIND_EXE_FILE_CALLBACKW, CallerData unsafe.Pointer) (r HANDLE, err error) = dbghelp.FindExecutableImageExW
//sys	ImageDirectoryEntryToDataEx(Base unsafe.Pointer, MappedAsImage BOOLEAN, DirectoryEntry IMAGE_DIRECTORY_ENTRY, Size *uint32, FoundHeader **IMAGE_SECTION_HEADER) (r unsafe.Pointer, err error) = dbghelp.ImageDirectoryEntryToDataEx
//sys	ImageDirectoryEntryToData(Base unsafe.Pointer, MappedAsImage BOOLEAN, DirectoryEntry IMAGE_DIRECTORY_ENTRY, Size *uint32) (r unsafe.Pointer, err error) = dbghelp.ImageDirectoryEntryToData
//sys	SearchTreeForFile(RootPath *PSTRElement, InputPathName *PSTRElement, OutputPathBuffer *PSTRElement) (r BOOL, err error) = dbghelp.SearchTreeForFile
//sys	SearchTreeForFileW(RootPath *PWSTRElement, InputPathName *PWSTRElement, OutputPathBuffer *PWSTRElement) (r BOOL, err error) = dbghelp.SearchTreeForFileW
//sys	EnumDirTree(hProcess HANDLE, RootPath *PSTRElement, InputPathName *PSTRElement, OutputPathBuffer *PSTRElement, cb PENUMDIRTREE_CALLBACK, data unsafe.Pointer) (r BOOL, err error) = dbghelp.EnumDirTree
//sys	EnumDirTreeW(hProcess HANDLE, RootPath *PWSTRElement, InputPathName *PWSTRElement, OutputPathBuffer *PWSTRElement, cb PENUMDIRTREE_CALLBACKW, data unsafe.Pointer) (r BOOL, err error) = dbghelp.EnumDirTreeW
//sys	MakeSureDirectoryPathExists(DirPath *PSTRElement) (r BOOL, err error) = dbghelp.MakeSureDirectoryPathExists
//sys	UnDecorateSymbolName(name *PSTRElement, outputString *PSTRElement, maxStringLength uint32, flags uint32) (r uint32, err error) = dbghelp.UnDecorateSymbolName
//sys	UnDecorateSymbolNameW(name *PWSTRElement, outputString *PWSTRElement, maxStringLength uint32, flags uint32) (r uint32, err error) = dbghelp.UnDecorateSymbolNameW
//sys	StackWalk64(MachineType uint32, hProcess HANDLE, hThread HANDLE, StackFrame *STACKFRAME64, ContextRecord unsafe.Pointer, ReadMemoryRoutine PREAD_PROCESS_MEMORY_ROUTINE64, FunctionTableAccessRoutine PFUNCTION_TABLE_ACCESS_ROUTINE64, GetModuleBaseRoutine PGET_MODULE_BASE_ROUTINE64, TranslateAddress PTRANSLATE_ADDRESS_ROUTINE64) (r BOOL) = dbghelp.StackWalk64
//sys	StackWalkEx(MachineType uint32, hProcess HANDLE, hThread HANDLE, StackFrame *STACKFRAME_EX, ContextRecord unsafe.Pointer, ReadMemoryRoutine PREAD_PROCESS_MEMORY_ROUTINE64, FunctionTableAccessRoutine PFUNCTION_TABLE_ACCESS_ROUTINE64, GetModuleBaseRoutine PGET_MODULE_BASE_ROUTINE64, TranslateAddress PTRANSLATE_ADDRESS_ROUTINE64, Flags uint32) (r BOOL) = dbghelp.StackWalkEx
//sys	StackWalk2(MachineType uint32, hProcess HANDLE, hThread HANDLE, StackFrame *STACKFRAME_EX, ContextRecord unsafe.Pointer, ReadMemoryRoutine PREAD_PROCESS_MEMORY_ROUTINE64, FunctionTableAccessRoutine PFUNCTION_TABLE_ACCESS_ROUTINE64, GetModuleBaseRoutine PGET_MODULE_BASE_ROUTINE64, TranslateAddress PTRANSLATE_ADDRESS_ROUTINE64, GetTargetAttributeValue PGET_TARGET_ATTRIBUTE_VALUE64, Flags uint32) (r BOOL) = dbghelp.StackWalk2
//sys	ImagehlpApiVersion() (r *API_VERSION) = dbghelp.ImagehlpApiVersion
//sys	ImagehlpApiVersionEx(AppVersion *API_VERSION) (r *API_VERSION) = dbghelp.ImagehlpApiVersionEx
//sys	GetTimestampForLoadedLibrary(Module HMODULE) (r uint32, err error) = dbghelp.GetTimestampForLoadedLibrary
//sys	SymSetParentWindow(hwnd HWND) (r BOOL, err error) = dbghelp.SymSetParentWindow
//sys	SymGetParentWindow(pHwnd *HWND) (r BOOL) = dbghelp.SymGetParentWindow
//sys	SymSetHomeDirectory(hProcess HANDLE, dir *PSTRElement) (r *PSTRElement, err error) = dbghelp.SymSetHomeDirectory
//sys	SymSetHomeDirectoryW(hProcess HANDLE, dir *PWSTRElement) (r *PWSTRElement, err error) = dbghelp.SymSetHomeDirectoryW
//sys	SymGetHomeDirectory(typeParam uint32, dir *PSTRElement, size uintptr) (r *PSTRElement, err error) = dbghelp.SymGetHomeDirectory
//sys	SymGetHomeDirectoryW(typeParam uint32, dir *PWSTRElement, size uintptr) (r *PWSTRElement, err error) = dbghelp.SymGetHomeDirectoryW
//sys	SymGetOmaps(hProcess HANDLE, BaseOfDll uint64, OmapTo **OMAP, cOmapTo *uint64, OmapFrom **OMAP, cOmapFrom *uint64) (r BOOL, err error) = dbghelp.SymGetOmaps
//sys	SymSetOptions(SymOptions uint32) (r uint32) = dbghelp.SymSetOptions
//sys	SymGetOptions() (r uint32) = dbghelp.SymGetOptions
//sys	SymCleanup(hProcess HANDLE) (r BOOL, err error) = dbghelp.SymCleanup
//sys	SymGetExtendedOption(option IMAGEHLP_EXTENDED_OPTIONS) (r BOOL) = dbghelp.SymGetExtendedOption
//sys	SymSetExtendedOption(option IMAGEHLP_EXTENDED_OPTIONS, value BOOL) (r BOOL) = dbghelp.SymSetExtendedOption
//sys	SymMatchString(string *PSTRElement, expression *PSTRElement, fCase BOOL) (r BOOL, err error) = dbghelp.SymMatchString
//sys	SymMatchStringA(string *PSTRElement, expression *PSTRElement, fCase BOOL) (r BOOL) = dbghelp.SymMatchStringA
//sys	SymMatchStringW(string *PWSTRElement, expression *PWSTRElement, fCase BOOL) (r BOOL, err error) = dbghelp.SymMatchStringW
//sys	SymEnumSourceFiles(hProcess HANDLE, ModBase uint64, Mask *PSTRElement, cbSrcFiles PSYM_ENUMSOURCEFILES_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumSourceFiles
//sys	SymEnumSourceFilesW(hProcess HANDLE, ModBase uint64, Mask *PWSTRElement, cbSrcFiles PSYM_ENUMSOURCEFILES_CALLBACKW, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumSourceFilesW
//sys	SymEnumerateModules64(hProcess HANDLE, EnumModulesCallback PSYM_ENUMMODULES_CALLBACK64, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumerateModules64
//sys	SymEnumerateModulesW64(hProcess HANDLE, EnumModulesCallback PSYM_ENUMMODULES_CALLBACKW64, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumerateModulesW64
//sys	EnumerateLoadedModulesEx(hProcess HANDLE, EnumLoadedModulesCallback PENUMLOADED_MODULES_CALLBACK64, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.EnumerateLoadedModulesEx
//sys	EnumerateLoadedModulesExW(hProcess HANDLE, EnumLoadedModulesCallback PENUMLOADED_MODULES_CALLBACKW64, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.EnumerateLoadedModulesExW
//sys	EnumerateLoadedModules64(hProcess HANDLE, EnumLoadedModulesCallback PENUMLOADED_MODULES_CALLBACK64, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.EnumerateLoadedModules64
//sys	EnumerateLoadedModulesW64(hProcess HANDLE, EnumLoadedModulesCallback PENUMLOADED_MODULES_CALLBACKW64, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.EnumerateLoadedModulesW64
//sys	SymFunctionTableAccess64(hProcess HANDLE, AddrBase uint64) (r unsafe.Pointer, err error) = dbghelp.SymFunctionTableAccess64
//sys	SymFunctionTableAccess64AccessRoutines(hProcess HANDLE, AddrBase uint64, ReadMemoryRoutine PREAD_PROCESS_MEMORY_ROUTINE64, GetModuleBaseRoutine PGET_MODULE_BASE_ROUTINE64) (r unsafe.Pointer) = dbghelp.SymFunctionTableAccess64AccessRoutines
//sys	SymGetUnwindInfo(hProcess HANDLE, Address uint64, Buffer unsafe.Pointer, Size *uint32) (r BOOL) = dbghelp.SymGetUnwindInfo
//sys	SymGetModuleInfo64(hProcess HANDLE, qwAddr uint64, ModuleInfo *IMAGEHLP_MODULE64) (r BOOL, err error) = dbghelp.SymGetModuleInfo64
//sys	SymGetModuleInfoW64(hProcess HANDLE, qwAddr uint64, ModuleInfo *IMAGEHLP_MODULEW64) (r BOOL, err error) = dbghelp.SymGetModuleInfoW64
//sys	SymGetModuleBase64(hProcess HANDLE, qwAddr uint64) (r uint64, err error) = dbghelp.SymGetModuleBase64
//sys	SymEnumLines(hProcess HANDLE, Base uint64, Obj *PSTRElement, File *PSTRElement, EnumLinesCallback PSYM_ENUMLINES_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumLines
//sys	SymEnumLinesW(hProcess HANDLE, Base uint64, Obj *PWSTRElement, File *PWSTRElement, EnumLinesCallback PSYM_ENUMLINES_CALLBACKW, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumLinesW
//sys	SymGetLineFromAddr64(hProcess HANDLE, qwAddr uint64, pdwDisplacement *uint32, Line64 *IMAGEHLP_LINE64) (r BOOL, err error) = dbghelp.SymGetLineFromAddr64
//sys	SymGetLineFromAddrW64(hProcess HANDLE, dwAddr uint64, pdwDisplacement *uint32, Line *IMAGEHLP_LINEW64) (r BOOL, err error) = dbghelp.SymGetLineFromAddrW64
//sys	SymGetLineFromInlineContext(hProcess HANDLE, qwAddr uint64, InlineContext uint32, qwModuleBaseAddress uint64, pdwDisplacement *uint32, Line64 *IMAGEHLP_LINE64) (r BOOL, err error) = dbghelp.SymGetLineFromInlineContext
//sys	SymGetLineFromInlineContextW(hProcess HANDLE, dwAddr uint64, InlineContext uint32, qwModuleBaseAddress uint64, pdwDisplacement *uint32, Line *IMAGEHLP_LINEW64) (r BOOL, err error) = dbghelp.SymGetLineFromInlineContextW
//sys	SymEnumSourceLines(hProcess HANDLE, Base uint64, Obj *PSTRElement, File *PSTRElement, Line uint32, Flags uint32, EnumLinesCallback PSYM_ENUMLINES_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumSourceLines
//sys	SymEnumSourceLinesW(hProcess HANDLE, Base uint64, Obj *PWSTRElement, File *PWSTRElement, Line uint32, Flags uint32, EnumLinesCallback PSYM_ENUMLINES_CALLBACKW, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumSourceLinesW
//sys	SymAddrIncludeInlineTrace(hProcess HANDLE, Address uint64) (r uint32) = dbghelp.SymAddrIncludeInlineTrace
//sys	SymCompareInlineTrace(hProcess HANDLE, Address1 uint64, InlineContext1 uint32, RetAddress1 uint64, Address2 uint64, RetAddress2 uint64) (r uint32) = dbghelp.SymCompareInlineTrace
//sys	SymQueryInlineTrace(hProcess HANDLE, StartAddress uint64, StartContext uint32, StartRetAddress uint64, CurAddress uint64, CurContext *uint32, CurFrameIndex *uint32) (r BOOL, err error) = dbghelp.SymQueryInlineTrace
//sys	SymGetLineFromName64(hProcess HANDLE, ModuleName *PSTRElement, FileName *PSTRElement, dwLineNumber uint32, plDisplacement *int32, Line *IMAGEHLP_LINE64) (r BOOL, err error) = dbghelp.SymGetLineFromName64
//sys	SymGetLineFromNameW64(hProcess HANDLE, ModuleName *PWSTRElement, FileName *PWSTRElement, dwLineNumber uint32, plDisplacement *int32, Line *IMAGEHLP_LINEW64) (r BOOL, err error) = dbghelp.SymGetLineFromNameW64
//sys	SymGetLineNext64(hProcess HANDLE, Line *IMAGEHLP_LINE64) (r BOOL, err error) = dbghelp.SymGetLineNext64
//sys	SymGetLineNextW64(hProcess HANDLE, Line *IMAGEHLP_LINEW64) (r BOOL, err error) = dbghelp.SymGetLineNextW64
//sys	SymGetLinePrev64(hProcess HANDLE, Line *IMAGEHLP_LINE64) (r BOOL, err error) = dbghelp.SymGetLinePrev64
//sys	SymGetLinePrevW64(hProcess HANDLE, Line *IMAGEHLP_LINEW64) (r BOOL, err error) = dbghelp.SymGetLinePrevW64
//sys	SymGetFileLineOffsets64(hProcess HANDLE, ModuleName *PSTRElement, FileName *PSTRElement, Buffer *uint64, BufferLines uint32) (r uint32, err error) = dbghelp.SymGetFileLineOffsets64
//sys	SymMatchFileName(FileName *PSTRElement, Match *PSTRElement, FileNameStop **PSTRElement, MatchStop **PSTRElement) (r BOOL, err error) = dbghelp.SymMatchFileName
//sys	SymMatchFileNameW(FileName *PWSTRElement, Match *PWSTRElement, FileNameStop **PWSTRElement, MatchStop **PWSTRElement) (r BOOL, err error) = dbghelp.SymMatchFileNameW
//sys	SymGetSourceFile(hProcess HANDLE, Base uint64, Params *PSTRElement, FileSpec *PSTRElement, FilePath *PSTRElement, Size uint32) (r BOOL, err error) = dbghelp.SymGetSourceFile
//sys	SymGetSourceFileW(hProcess HANDLE, Base uint64, Params *PWSTRElement, FileSpec *PWSTRElement, FilePath *PWSTRElement, Size uint32) (r BOOL, err error) = dbghelp.SymGetSourceFileW
//sys	SymGetSourceFileToken(hProcess HANDLE, Base uint64, FileSpec *PSTRElement, Token *unsafe.Pointer, Size *uint32) (r BOOL, err error) = dbghelp.SymGetSourceFileToken
//sys	SymGetSourceFileTokenByTokenName(hProcess HANDLE, Base uint64, FileSpec *PSTRElement, TokenName *PSTRElement, TokenParameters *PSTRElement, Token *unsafe.Pointer, Size *uint32) (r BOOL) = dbghelp.SymGetSourceFileTokenByTokenName
//sys	SymGetSourceFileChecksumW(hProcess HANDLE, Base uint64, FileSpec *PWSTRElement, pCheckSumType *uint32, pChecksum *uint8, checksumSize uint32, pActualBytesWritten *uint32) (r BOOL, err error) = dbghelp.SymGetSourceFileChecksumW
//sys	SymGetSourceFileChecksum(hProcess HANDLE, Base uint64, FileSpec *PSTRElement, pCheckSumType *uint32, pChecksum *uint8, checksumSize uint32, pActualBytesWritten *uint32) (r BOOL, err error) = dbghelp.SymGetSourceFileChecksum
//sys	SymGetSourceFileTokenW(hProcess HANDLE, Base uint64, FileSpec *PWSTRElement, Token *unsafe.Pointer, Size *uint32) (r BOOL, err error) = dbghelp.SymGetSourceFileTokenW
//sys	SymGetSourceFileTokenByTokenNameW(hProcess HANDLE, Base uint64, FileSpec *PWSTRElement, TokenName *PWSTRElement, TokenParameters *PWSTRElement, Token *unsafe.Pointer, Size *uint32) (r BOOL) = dbghelp.SymGetSourceFileTokenByTokenNameW
//sys	SymGetSourceFileFromToken(hProcess HANDLE, Token unsafe.Pointer, Params *PSTRElement, FilePath *PSTRElement, Size uint32) (r BOOL, err error) = dbghelp.SymGetSourceFileFromToken
//sys	SymGetSourceFileFromTokenByTokenName(hProcess HANDLE, Token unsafe.Pointer, TokenName *PSTRElement, Params *PSTRElement, FilePath *PSTRElement, Size uint32) (r BOOL) = dbghelp.SymGetSourceFileFromTokenByTokenName
//sys	SymGetSourceFileFromTokenW(hProcess HANDLE, Token unsafe.Pointer, Params *PWSTRElement, FilePath *PWSTRElement, Size uint32) (r BOOL, err error) = dbghelp.SymGetSourceFileFromTokenW
//sys	SymGetSourceFileFromTokenByTokenNameW(hProcess HANDLE, Token unsafe.Pointer, TokenName *PWSTRElement, Params *PWSTRElement, FilePath *PWSTRElement, Size uint32) (r BOOL) = dbghelp.SymGetSourceFileFromTokenByTokenNameW
//sys	SymGetSourceVarFromToken(hProcess HANDLE, Token unsafe.Pointer, Params *PSTRElement, VarName *PSTRElement, Value *PSTRElement, Size uint32) (r BOOL, err error) = dbghelp.SymGetSourceVarFromToken
//sys	SymGetSourceVarFromTokenW(hProcess HANDLE, Token unsafe.Pointer, Params *PWSTRElement, VarName *PWSTRElement, Value *PWSTRElement, Size uint32) (r BOOL, err error) = dbghelp.SymGetSourceVarFromTokenW
//sys	SymEnumSourceFileTokens(hProcess HANDLE, Base uint64, Callback PENUMSOURCEFILETOKENSCALLBACK) (r BOOL, err error) = dbghelp.SymEnumSourceFileTokens
//sys	SymInitialize(hProcess HANDLE, UserSearchPath *PSTRElement, fInvadeProcess BOOL) (r BOOL, err error) = dbghelp.SymInitialize
//sys	SymInitializeW(hProcess HANDLE, UserSearchPath *PWSTRElement, fInvadeProcess BOOL) (r BOOL, err error) = dbghelp.SymInitializeW
//sys	SymGetSearchPath(hProcess HANDLE, SearchPathA *PSTRElement, SearchPathLength uint32) (r BOOL, err error) = dbghelp.SymGetSearchPath
//sys	SymGetSearchPathW(hProcess HANDLE, SearchPathA *PWSTRElement, SearchPathLength uint32) (r BOOL, err error) = dbghelp.SymGetSearchPathW
//sys	SymSetSearchPath(hProcess HANDLE, SearchPathA *PSTRElement) (r BOOL, err error) = dbghelp.SymSetSearchPath
//sys	SymSetSearchPathW(hProcess HANDLE, SearchPathA *PWSTRElement) (r BOOL, err error) = dbghelp.SymSetSearchPathW
//sys	SymLoadModuleEx(hProcess HANDLE, hFile HANDLE, ImageName *PSTRElement, ModuleName *PSTRElement, BaseOfDll uint64, DllSize uint32, Data *MODLOAD_DATA, Flags SYM_LOAD_FLAGS) (r uint64, err error) = dbghelp.SymLoadModuleEx
//sys	SymLoadModuleExW(hProcess HANDLE, hFile HANDLE, ImageName *PWSTRElement, ModuleName *PWSTRElement, BaseOfDll uint64, DllSize uint32, Data *MODLOAD_DATA, Flags SYM_LOAD_FLAGS) (r uint64, err error) = dbghelp.SymLoadModuleExW
//sys	SymUnloadModule64(hProcess HANDLE, BaseOfDll uint64) (r BOOL, err error) = dbghelp.SymUnloadModule64
//sys	SymUnDName64(sym *IMAGEHLP_SYMBOL64, UnDecName *PSTRElement, UnDecNameLength uint32) (r BOOL, err error) = dbghelp.SymUnDName64
//sys	SymRegisterCallback64(hProcess HANDLE, CallbackFunction PSYMBOL_REGISTERED_CALLBACK64, UserContext uint64) (r BOOL, err error) = dbghelp.SymRegisterCallback64
//sys	SymRegisterCallbackW64(hProcess HANDLE, CallbackFunction PSYMBOL_REGISTERED_CALLBACK64, UserContext uint64) (r BOOL, err error) = dbghelp.SymRegisterCallbackW64
//sys	SymRegisterFunctionEntryCallback64(hProcess HANDLE, CallbackFunction PSYMBOL_FUNCENTRY_CALLBACK64, UserContext uint64) (r BOOL, err error) = dbghelp.SymRegisterFunctionEntryCallback64
//sys	SymSetContext(hProcess HANDLE, StackFrame *IMAGEHLP_STACK_FRAME, Context unsafe.Pointer) (r BOOL, err error) = dbghelp.SymSetContext
//sys	SymSetScopeFromAddr(hProcess HANDLE, Address uint64) (r BOOL, err error) = dbghelp.SymSetScopeFromAddr
//sys	SymSetScopeFromInlineContext(hProcess HANDLE, Address uint64, InlineContext uint32) (r BOOL, err error) = dbghelp.SymSetScopeFromInlineContext
//sys	SymSetScopeFromIndex(hProcess HANDLE, BaseOfDll uint64, Index uint32) (r BOOL, err error) = dbghelp.SymSetScopeFromIndex
//sys	SymEnumProcesses(EnumProcessesCallback PSYM_ENUMPROCESSES_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumProcesses
//sys	SymFromAddr(hProcess HANDLE, Address uint64, Displacement *uint64, Symbol *SYMBOL_INFO) (r BOOL, err error) = dbghelp.SymFromAddr
//sys	SymFromAddrW(hProcess HANDLE, Address uint64, Displacement *uint64, Symbol *SYMBOL_INFOW) (r BOOL, err error) = dbghelp.SymFromAddrW
//sys	SymFromInlineContext(hProcess HANDLE, Address uint64, InlineContext uint32, Displacement *uint64, Symbol *SYMBOL_INFO) (r BOOL, err error) = dbghelp.SymFromInlineContext
//sys	SymFromInlineContextW(hProcess HANDLE, Address uint64, InlineContext uint32, Displacement *uint64, Symbol *SYMBOL_INFOW) (r BOOL, err error) = dbghelp.SymFromInlineContextW
//sys	SymFromToken(hProcess HANDLE, Base uint64, Token uint32, Symbol *SYMBOL_INFO) (r BOOL, err error) = dbghelp.SymFromToken
//sys	SymFromTokenW(hProcess HANDLE, Base uint64, Token uint32, Symbol *SYMBOL_INFOW) (r BOOL, err error) = dbghelp.SymFromTokenW
//sys	SymNext(hProcess HANDLE, si *SYMBOL_INFO) (r BOOL, err error) = dbghelp.SymNext
//sys	SymNextW(hProcess HANDLE, siw *SYMBOL_INFOW) (r BOOL, err error) = dbghelp.SymNextW
//sys	SymPrev(hProcess HANDLE, si *SYMBOL_INFO) (r BOOL, err error) = dbghelp.SymPrev
//sys	SymPrevW(hProcess HANDLE, siw *SYMBOL_INFOW) (r BOOL, err error) = dbghelp.SymPrevW
//sys	SymFromName(hProcess HANDLE, Name *PSTRElement, Symbol *SYMBOL_INFO) (r BOOL, err error) = dbghelp.SymFromName
//sys	SymFromNameW(hProcess HANDLE, Name *PWSTRElement, Symbol *SYMBOL_INFOW) (r BOOL, err error) = dbghelp.SymFromNameW
//sys	SymEnumSymbols(hProcess HANDLE, BaseOfDll uint64, Mask *PSTRElement, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumSymbols
//sys	SymEnumSymbolsEx(hProcess HANDLE, BaseOfDll uint64, Mask *PSTRElement, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACK, UserContext unsafe.Pointer, Options uint32) (r BOOL, err error) = dbghelp.SymEnumSymbolsEx
//sys	SymEnumSymbolsW(hProcess HANDLE, BaseOfDll uint64, Mask *PWSTRElement, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACKW, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumSymbolsW
//sys	SymEnumSymbolsExW(hProcess HANDLE, BaseOfDll uint64, Mask *PWSTRElement, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACKW, UserContext unsafe.Pointer, Options uint32) (r BOOL, err error) = dbghelp.SymEnumSymbolsExW
//sys	SymEnumSymbolsForAddr(hProcess HANDLE, Address uint64, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumSymbolsForAddr
//sys	SymEnumSymbolsForAddrW(hProcess HANDLE, Address uint64, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACKW, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumSymbolsForAddrW
//sys	SymSearch(hProcess HANDLE, BaseOfDll uint64, Index uint32, SymTag uint32, Mask *PSTRElement, Address uint64, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACK, UserContext unsafe.Pointer, Options uint32) (r BOOL, err error) = dbghelp.SymSearch
//sys	SymSearchW(hProcess HANDLE, BaseOfDll uint64, Index uint32, SymTag uint32, Mask *PWSTRElement, Address uint64, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACKW, UserContext unsafe.Pointer, Options uint32) (r BOOL, err error) = dbghelp.SymSearchW
//sys	SymGetScope(hProcess HANDLE, BaseOfDll uint64, Index uint32, Symbol *SYMBOL_INFO) (r BOOL, err error) = dbghelp.SymGetScope
//sys	SymGetScopeW(hProcess HANDLE, BaseOfDll uint64, Index uint32, Symbol *SYMBOL_INFOW) (r BOOL, err error) = dbghelp.SymGetScopeW
//sys	SymFromIndex(hProcess HANDLE, BaseOfDll uint64, Index uint32, Symbol *SYMBOL_INFO) (r BOOL, err error) = dbghelp.SymFromIndex
//sys	SymFromIndexW(hProcess HANDLE, BaseOfDll uint64, Index uint32, Symbol *SYMBOL_INFOW) (r BOOL, err error) = dbghelp.SymFromIndexW
//sys	SymGetTypeInfo(hProcess HANDLE, ModBase uint64, TypeId uint32, GetType IMAGEHLP_SYMBOL_TYPE_INFO, pInfo unsafe.Pointer) (r BOOL, err error) = dbghelp.SymGetTypeInfo
//sys	SymGetTypeInfoEx(hProcess HANDLE, ModBase uint64, Params *IMAGEHLP_GET_TYPE_INFO_PARAMS) (r BOOL, err error) = dbghelp.SymGetTypeInfoEx
//sys	SymEnumTypes(hProcess HANDLE, BaseOfDll uint64, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumTypes
//sys	SymEnumTypesW(hProcess HANDLE, BaseOfDll uint64, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACKW, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumTypesW
//sys	SymEnumTypesByName(hProcess HANDLE, BaseOfDll uint64, mask *PSTRElement, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACK, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumTypesByName
//sys	SymEnumTypesByNameW(hProcess HANDLE, BaseOfDll uint64, mask *PWSTRElement, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACKW, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumTypesByNameW
//sys	SymGetTypeFromName(hProcess HANDLE, BaseOfDll uint64, Name *PSTRElement, Symbol *SYMBOL_INFO) (r BOOL, err error) = dbghelp.SymGetTypeFromName
//sys	SymGetTypeFromNameW(hProcess HANDLE, BaseOfDll uint64, Name *PWSTRElement, Symbol *SYMBOL_INFOW) (r BOOL, err error) = dbghelp.SymGetTypeFromNameW
//sys	SymAddSymbol(hProcess HANDLE, BaseOfDll uint64, Name *PSTRElement, Address uint64, Size uint32, Flags uint32) (r BOOL, err error) = dbghelp.SymAddSymbol
//sys	SymAddSymbolW(hProcess HANDLE, BaseOfDll uint64, Name *PWSTRElement, Address uint64, Size uint32, Flags uint32) (r BOOL, err error) = dbghelp.SymAddSymbolW
//sys	SymDeleteSymbol(hProcess HANDLE, BaseOfDll uint64, Name *PSTRElement, Address uint64, Flags uint32) (r BOOL, err error) = dbghelp.SymDeleteSymbol
//sys	SymDeleteSymbolW(hProcess HANDLE, BaseOfDll uint64, Name *PWSTRElement, Address uint64, Flags uint32) (r BOOL, err error) = dbghelp.SymDeleteSymbolW
//sys	SymRefreshModuleList(hProcess HANDLE) (r BOOL, err error) = dbghelp.SymRefreshModuleList
//sys	SymAddSourceStream(hProcess HANDLE, Base uint64, StreamFile *PSTRElement, Buffer *uint8, Size uintptr) (r BOOL, err error) = dbghelp.SymAddSourceStream
//sys	SymAddSourceStreamA(hProcess HANDLE, Base uint64, StreamFile *PSTRElement, Buffer *uint8, Size uintptr) (r BOOL) = dbghelp.SymAddSourceStreamA
//sys	SymAddSourceStreamW(hProcess HANDLE, Base uint64, FileSpec *PWSTRElement, Buffer *uint8, Size uintptr) (r BOOL, err error) = dbghelp.SymAddSourceStreamW
//sys	SymSrvIsStoreW(hProcess HANDLE, path *PWSTRElement) (r BOOL, err error) = dbghelp.SymSrvIsStoreW
//sys	SymSrvIsStore(hProcess HANDLE, path *PSTRElement) (r BOOL, err error) = dbghelp.SymSrvIsStore
//sys	SymSrvDeltaName(hProcess HANDLE, SymPath *PSTRElement, Type *PSTRElement, File1 *PSTRElement, File2 *PSTRElement) (r *PSTRElement, err error) = dbghelp.SymSrvDeltaName
//sys	SymSrvDeltaNameW(hProcess HANDLE, SymPath *PWSTRElement, Type *PWSTRElement, File1 *PWSTRElement, File2 *PWSTRElement) (r *PWSTRElement, err error) = dbghelp.SymSrvDeltaNameW
//sys	SymSrvGetSupplement(hProcess HANDLE, SymPath *PSTRElement, Node *PSTRElement, File *PSTRElement) (r *PSTRElement, err error) = dbghelp.SymSrvGetSupplement
//sys	SymSrvGetSupplementW(hProcess HANDLE, SymPath *PWSTRElement, Node *PWSTRElement, File *PWSTRElement) (r *PWSTRElement, err error) = dbghelp.SymSrvGetSupplementW
//sys	SymSrvGetFileIndexes(File *PSTRElement, Id *Guid, Val1 *uint32, Val2 *uint32, Flags uint32) (r BOOL, err error) = dbghelp.SymSrvGetFileIndexes
//sys	SymSrvGetFileIndexesW(File *PWSTRElement, Id *Guid, Val1 *uint32, Val2 *uint32, Flags uint32) (r BOOL, err error) = dbghelp.SymSrvGetFileIndexesW
//sys	SymSrvGetFileIndexStringW(hProcess HANDLE, SrvPath *PWSTRElement, File *PWSTRElement, Index *PWSTRElement, Size uintptr, Flags uint32) (r BOOL, err error) = dbghelp.SymSrvGetFileIndexStringW
//sys	SymSrvGetFileIndexString(hProcess HANDLE, SrvPath *PSTRElement, File *PSTRElement, Index *PSTRElement, Size uintptr, Flags uint32) (r BOOL, err error) = dbghelp.SymSrvGetFileIndexString
//sys	SymSrvGetFileIndexInfo(File *PSTRElement, Info *SYMSRV_INDEX_INFO, Flags uint32) (r BOOL, err error) = dbghelp.SymSrvGetFileIndexInfo
//sys	SymSrvGetFileIndexInfoW(File *PWSTRElement, Info *SYMSRV_INDEX_INFOW, Flags uint32) (r BOOL, err error) = dbghelp.SymSrvGetFileIndexInfoW
//sys	SymSrvStoreSupplement(hProcess HANDLE, SrvPath *PSTRElement, Node *PSTRElement, File *PSTRElement, Flags uint32) (r *PSTRElement, err error) = dbghelp.SymSrvStoreSupplement
//sys	SymSrvStoreSupplementW(hProcess HANDLE, SymPath *PWSTRElement, Node *PWSTRElement, File *PWSTRElement, Flags uint32) (r *PWSTRElement, err error) = dbghelp.SymSrvStoreSupplementW
//sys	SymSrvStoreFile(hProcess HANDLE, SrvPath *PSTRElement, File *PSTRElement, Flags SYM_SRV_STORE_FILE_FLAGS) (r *PSTRElement, err error) = dbghelp.SymSrvStoreFile
//sys	SymSrvStoreFileW(hProcess HANDLE, SrvPath *PWSTRElement, File *PWSTRElement, Flags SYM_SRV_STORE_FILE_FLAGS) (r *PWSTRElement, err error) = dbghelp.SymSrvStoreFileW
//sys	SymGetSymbolFile(hProcess HANDLE, SymPath *PSTRElement, ImageFile *PSTRElement, Type uint32, SymbolFile *PSTRElement, cSymbolFile uintptr, DbgFile *PSTRElement, cDbgFile uintptr) (r BOOL, err error) = dbghelp.SymGetSymbolFile
//sys	SymGetSymbolFileW(hProcess HANDLE, SymPath *PWSTRElement, ImageFile *PWSTRElement, Type uint32, SymbolFile *PWSTRElement, cSymbolFile uintptr, DbgFile *PWSTRElement, cDbgFile uintptr) (r BOOL, err error) = dbghelp.SymGetSymbolFileW
//sys	DbgHelpCreateUserDump(FileName *PSTRElement, Callback PDBGHELP_CREATE_USER_DUMP_CALLBACK, UserData unsafe.Pointer) (r BOOL) = dbghelp.DbgHelpCreateUserDump
//sys	DbgHelpCreateUserDumpW(FileName *PWSTRElement, Callback PDBGHELP_CREATE_USER_DUMP_CALLBACK, UserData unsafe.Pointer) (r BOOL) = dbghelp.DbgHelpCreateUserDumpW
//sys	SymGetSymFromAddr64(hProcess HANDLE, qwAddr uint64, pdwDisplacement *uint64, Symbol *IMAGEHLP_SYMBOL64) (r BOOL, err error) = dbghelp.SymGetSymFromAddr64
//sys	SymGetSymFromName64(hProcess HANDLE, Name *PSTRElement, Symbol *IMAGEHLP_SYMBOL64) (r BOOL, err error) = dbghelp.SymGetSymFromName64
//sys	FindFileInPath(hprocess HANDLE, SearchPathA *PSTRElement, FileName *PSTRElement, id unsafe.Pointer, two uint32, three uint32, flags uint32, FilePath *PSTRElement) (r BOOL) = dbghelp.FindFileInPath
//sys	FindFileInSearchPath(hprocess HANDLE, SearchPathA *PSTRElement, FileName *PSTRElement, one uint32, two uint32, three uint32, FilePath *PSTRElement) (r BOOL) = dbghelp.FindFileInSearchPath
//sys	SymEnumSym(hProcess HANDLE, BaseOfDll uint64, EnumSymbolsCallback PSYM_ENUMERATESYMBOLS_CALLBACK, UserContext unsafe.Pointer) (r BOOL) = dbghelp.SymEnumSym
//sys	SymEnumerateSymbols64(hProcess HANDLE, BaseOfDll uint64, EnumSymbolsCallback PSYM_ENUMSYMBOLS_CALLBACK64, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumerateSymbols64
//sys	SymEnumerateSymbolsW64(hProcess HANDLE, BaseOfDll uint64, EnumSymbolsCallback PSYM_ENUMSYMBOLS_CALLBACK64W, UserContext unsafe.Pointer) (r BOOL, err error) = dbghelp.SymEnumerateSymbolsW64
//sys	SymLoadModule64(hProcess HANDLE, hFile HANDLE, ImageName *PSTRElement, ModuleName *PSTRElement, BaseOfDll uint64, SizeOfDll uint32) (r uint64, err error) = dbghelp.SymLoadModule64
//sys	SymGetSymNext64(hProcess HANDLE, Symbol *IMAGEHLP_SYMBOL64) (r BOOL, err error) = dbghelp.SymGetSymNext64
//sys	SymGetSymPrev64(hProcess HANDLE, Symbol *IMAGEHLP_SYMBOL64) (r BOOL, err error) = dbghelp.SymGetSymPrev64
//sys	SetCheckUserInterruptShared(lpStartAddress LPCALL_BACK_USER_INTERRUPT_ROUTINE) = dbghelp.SetCheckUserInterruptShared
//sys	GetSymLoadError() (r uint32) = dbghelp.GetSymLoadError
//sys	SetSymLoadError(error uint32) = dbghelp.SetSymLoadError
//sys	ReportSymbolLoadSummary(hProcess HANDLE, pLoadModule *PWSTRElement, pSymbolData *DBGHELP_DATA_REPORT_STRUCT) (r BOOL) = dbghelp.ReportSymbolLoadSummary
//sys	RemoveInvalidModuleList(hProcess HANDLE) = dbghelp.RemoveInvalidModuleList
//sys	RangeMapCreate() (r unsafe.Pointer) = dbghelp.RangeMapCreate
//sys	RangeMapFree(RmapHandle unsafe.Pointer) = dbghelp.RangeMapFree
//sys	RangeMapAddPeImageSections(RmapHandle unsafe.Pointer, ImageName *PWSTRElement, MappedImage unsafe.Pointer, MappingBytes uint32, ImageBase uint64, UserTag uint64, MappingFlags uint32) (r BOOL) = dbghelp.RangeMapAddPeImageSections
//sys	RangeMapRemove(RmapHandle unsafe.Pointer, UserTag uint64) (r BOOL) = dbghelp.RangeMapRemove
//sys	RangeMapRead(RmapHandle unsafe.Pointer, Offset uint64, Buffer unsafe.Pointer, RequestBytes uint32, Flags uint32, DoneBytes *uint32) (r BOOL) = dbghelp.RangeMapRead
//sys	RangeMapWrite(RmapHandle unsafe.Pointer, Offset uint64, Buffer unsafe.Pointer, RequestBytes uint32, Flags uint32, DoneBytes *uint32) (r BOOL) = dbghelp.RangeMapWrite
//sys	MessageBeep(uType MESSAGEBOX_STYLE) (r BOOL, err error) = user32.MessageBeep
//sys	FatalExit(ExitCode int32)
//sys	GetThreadSelectorEntry(hThread HANDLE, dwSelector uint32, lpSelectorEntry *LDT_ENTRY) (r BOOL, err error)
//sys	Wow64GetThreadSelectorEntry(hThread HANDLE, dwSelector uint32, lpSelectorEntry *WOW64_LDT_ENTRY) (r BOOL, err error)
//sys	DebugSetProcessKillOnExit(KillOnExit BOOL) (r BOOL, err error)
//sys	DebugBreakProcess(Process HANDLE) (r BOOL, err error)
//sys	FormatMessageA(dwFlags FORMAT_MESSAGE_OPTIONS, lpSource unsafe.Pointer, dwMessageId uint32, dwLanguageId uint32, lpBuffer *PSTRElement, nSize uint32, Arguments **int8) (r uint32, err error)
//sys	FormatMessageW(dwFlags FORMAT_MESSAGE_OPTIONS, lpSource unsafe.Pointer, dwMessageId uint32, dwLanguageId uint32, lpBuffer *PWSTRElement, nSize uint32, Arguments **int8) (r uint32, err error)
//sys	CopyContext(Destination *CONTEXT, ContextFlags CONTEXT_FLAGS, Source *CONTEXT) (r BOOL, err error)
//sys	InitializeContext(Buffer unsafe.Pointer, ContextFlags CONTEXT_FLAGS, Context **CONTEXT, ContextLength *uint32) (r BOOL, err error)
//sys	InitializeContext2(Buffer unsafe.Pointer, ContextFlags CONTEXT_FLAGS, Context **CONTEXT, ContextLength *uint32, XStateCompactionMask uint64) (r BOOL, err error)
//sys	GetEnabledXStateFeatures() (r uint64)
//sys	GetXStateFeaturesMask(Context *CONTEXT, FeatureMask *uint64) (r BOOL)
//sys	LocateXStateFeature(Context *CONTEXT, FeatureId uint32, Length *uint32) (r unsafe.Pointer)
//sys	SetXStateFeaturesMask(Context *CONTEXT, FeatureMask uint64) (r BOOL)

// APIs for Windows.Win32.Storage.FileSystem
//sys	SearchPathW(lpPath *PWSTRElement, lpFileName *PWSTRElement, lpExtension *PWSTRElement, nBufferLength uint32, lpBuffer *PWSTRElement, lpFilePart **PWSTRElement) (r uint32, err error)
//sys	SearchPathA(lpPath *PSTRElement, lpFileName *PSTRElement, lpExtension *PSTRElement, nBufferLength uint32, lpBuffer *PSTRElement, lpFilePart **PSTRElement) (r uint32, err error)
//sys	CompareFileTime(lpFileTime1 *FILETIME, lpFileTime2 *FILETIME) (r int32)
//sys	CreateDirectoryA(lpPathName *PSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r BOOL, err error)
//sys	CreateDirectoryW(lpPathName *PWSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r BOOL, err error)
//sys	CreateFileA(lpFileName *PSTRElement, dwDesiredAccess uint32, dwShareMode FILE_SHARE_MODE, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwCreationDisposition FILE_CREATION_DISPOSITION, dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES, hTemplateFile HANDLE) (r HANDLE, err error)
//sys	CreateFileW(lpFileName *PWSTRElement, dwDesiredAccess uint32, dwShareMode FILE_SHARE_MODE, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwCreationDisposition FILE_CREATION_DISPOSITION, dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES, hTemplateFile HANDLE) (r HANDLE, err error)
//sys	DefineDosDeviceW(dwFlags DEFINE_DOS_DEVICE_FLAGS, lpDeviceName *PWSTRElement, lpTargetPath *PWSTRElement) (r BOOL, err error)
//sys	DeleteFileA(lpFileName *PSTRElement) (r BOOL, err error)
//sys	DeleteFileW(lpFileName *PWSTRElement) (r BOOL, err error)
//sys	DeleteVolumeMountPointW(lpszVolumeMountPoint *PWSTRElement) (r BOOL, err error)
//sys	FileTimeToLocalFileTime(lpFileTime *FILETIME, lpLocalFileTime *FILETIME) (r BOOL, err error)
//sys	FindClose(hFindFile HANDLE) (r BOOL, err error)
//sys	FindCloseChangeNotification(hChangeHandle HANDLE) (r BOOL, err error)
//sys	FindFirstChangeNotificationA(lpPathName *PSTRElement, bWatchSubtree BOOL, dwNotifyFilter FILE_NOTIFY_CHANGE) (r HANDLE, err error)
//sys	FindFirstChangeNotificationW(lpPathName *PWSTRElement, bWatchSubtree BOOL, dwNotifyFilter FILE_NOTIFY_CHANGE) (r HANDLE, err error)
//sys	FindFirstFileA(lpFileName *PSTRElement, lpFindFileData *WIN32_FIND_DATAA) (r HANDLE, err error)
//sys	FindFirstFileW(lpFileName *PWSTRElement, lpFindFileData *WIN32_FIND_DATAW) (r HANDLE, err error)
//sys	FindFirstFileExA(lpFileName *PSTRElement, fInfoLevelId FINDEX_INFO_LEVELS, lpFindFileData unsafe.Pointer, fSearchOp FINDEX_SEARCH_OPS, lpSearchFilter unsafe.Pointer, dwAdditionalFlags FIND_FIRST_EX_FLAGS) (r HANDLE, err error)
//sys	FindFirstFileExW(lpFileName *PWSTRElement, fInfoLevelId FINDEX_INFO_LEVELS, lpFindFileData unsafe.Pointer, fSearchOp FINDEX_SEARCH_OPS, lpSearchFilter unsafe.Pointer, dwAdditionalFlags FIND_FIRST_EX_FLAGS) (r HANDLE, err error)
//sys	FindFirstVolumeW(lpszVolumeName *PWSTRElement, cchBufferLength uint32) (r HANDLE, err error)
//sys	FindNextChangeNotification(hChangeHandle HANDLE) (r BOOL, err error)
//sys	FindNextFileA(hFindFile HANDLE, lpFindFileData *WIN32_FIND_DATAA) (r BOOL, err error)
//sys	FindNextFileW(hFindFile HANDLE, lpFindFileData *WIN32_FIND_DATAW) (r BOOL, err error)
//sys	FindNextVolumeW(hFindVolume HANDLE, lpszVolumeName *PWSTRElement, cchBufferLength uint32) (r BOOL, err error)
//sys	FindVolumeClose(hFindVolume HANDLE) (r BOOL, err error)
//sys	FlushFileBuffers(hFile HANDLE) (r BOOL, err error)
//sys	GetDiskFreeSpaceA(lpRootPathName *PSTRElement, lpSectorsPerCluster *uint32, lpBytesPerSector *uint32, lpNumberOfFreeClusters *uint32, lpTotalNumberOfClusters *uint32) (r BOOL, err error)
//sys	GetDiskFreeSpaceW(lpRootPathName *PWSTRElement, lpSectorsPerCluster *uint32, lpBytesPerSector *uint32, lpNumberOfFreeClusters *uint32, lpTotalNumberOfClusters *uint32) (r BOOL, err error)
//sys	GetDiskFreeSpaceExA(lpDirectoryName *PSTRElement, lpFreeBytesAvailableToCaller *uint64, lpTotalNumberOfBytes *uint64, lpTotalNumberOfFreeBytes *uint64) (r BOOL, err error)
//sys	GetDiskFreeSpaceExW(lpDirectoryName *PWSTRElement, lpFreeBytesAvailableToCaller *uint64, lpTotalNumberOfBytes *uint64, lpTotalNumberOfFreeBytes *uint64) (r BOOL, err error)
//sys	GetDiskSpaceInformationA(rootPath *PSTRElement, diskSpaceInfo *DISK_SPACE_INFORMATION) (r HRESULT)
//sys	GetDiskSpaceInformationW(rootPath *PWSTRElement, diskSpaceInfo *DISK_SPACE_INFORMATION) (r HRESULT)
//sys	GetDriveTypeA(lpRootPathName *PSTRElement) (r uint32)
//sys	GetDriveTypeW(lpRootPathName *PWSTRElement) (r uint32)
//sys	GetFileAttributesA(lpFileName *PSTRElement) (r uint32, err error)
//sys	GetFileAttributesW(lpFileName *PWSTRElement) (r uint32, err error)
//sys	GetFileAttributesExA(lpFileName *PSTRElement, fInfoLevelId GET_FILEEX_INFO_LEVELS, lpFileInformation unsafe.Pointer) (r BOOL, err error)
//sys	GetFileAttributesExW(lpFileName *PWSTRElement, fInfoLevelId GET_FILEEX_INFO_LEVELS, lpFileInformation unsafe.Pointer) (r BOOL, err error)
//sys	GetFileInformationByHandle(hFile HANDLE, lpFileInformation *BY_HANDLE_FILE_INFORMATION) (r BOOL, err error)
//sys	GetFileSize(hFile HANDLE, lpFileSizeHigh *uint32) (r uint32, err error)
//sys	GetFileSizeEx(hFile HANDLE, lpFileSize *int64) (r BOOL, err error)
//sys	GetFileType(hFile HANDLE) (r FILE_TYPE, err error)
//sys	GetFinalPathNameByHandleA(hFile HANDLE, lpszFilePath *PSTRElement, cchFilePath uint32, dwFlags GETFINALPATHNAMEBYHANDLE_FLAGS) (r uint32, err error)
//sys	GetFinalPathNameByHandleW(hFile HANDLE, lpszFilePath *PWSTRElement, cchFilePath uint32, dwFlags GETFINALPATHNAMEBYHANDLE_FLAGS) (r uint32, err error)
//sys	GetFileTime(hFile HANDLE, lpCreationTime *FILETIME, lpLastAccessTime *FILETIME, lpLastWriteTime *FILETIME) (r BOOL, err error)
//sys	GetFullPathNameW(lpFileName *PWSTRElement, nBufferLength uint32, lpBuffer *PWSTRElement, lpFilePart **PWSTRElement) (r uint32, err error)
//sys	GetFullPathNameA(lpFileName *PSTRElement, nBufferLength uint32, lpBuffer *PSTRElement, lpFilePart **PSTRElement) (r uint32, err error)
//sys	GetLogicalDrives() (r uint32, err error)
//sys	GetLogicalDriveStringsW(nBufferLength uint32, lpBuffer *PWSTRElement) (r uint32, err error)
//sys	GetLongPathNameA(lpszShortPath *PSTRElement, lpszLongPath *PSTRElement, cchBuffer uint32) (r uint32, err error)
//sys	GetLongPathNameW(lpszShortPath *PWSTRElement, lpszLongPath *PWSTRElement, cchBuffer uint32) (r uint32, err error)
//sys	AreShortNamesEnabled(Handle HANDLE, Enabled *BOOL) (r BOOL)
//sys	GetShortPathNameW(lpszLongPath *PWSTRElement, lpszShortPath *PWSTRElement, cchBuffer uint32) (r uint32, err error)
//sys	GetTempFileNameW(lpPathName *PWSTRElement, lpPrefixString *PWSTRElement, uUnique uint32, lpTempFileName *PWSTRElement) (r uint32, err error)
//sys	GetVolumeInformationByHandleW(hFile HANDLE, lpVolumeNameBuffer *PWSTRElement, nVolumeNameSize uint32, lpVolumeSerialNumber *uint32, lpMaximumComponentLength *uint32, lpFileSystemFlags *uint32, lpFileSystemNameBuffer *PWSTRElement, nFileSystemNameSize uint32) (r BOOL, err error)
//sys	GetVolumeInformationW(lpRootPathName *PWSTRElement, lpVolumeNameBuffer *PWSTRElement, nVolumeNameSize uint32, lpVolumeSerialNumber *uint32, lpMaximumComponentLength *uint32, lpFileSystemFlags *uint32, lpFileSystemNameBuffer *PWSTRElement, nFileSystemNameSize uint32) (r BOOL, err error)
//sys	GetVolumePathNameW(lpszFileName *PWSTRElement, lpszVolumePathName *PWSTRElement, cchBufferLength uint32) (r BOOL, err error)
//sys	LocalFileTimeToFileTime(lpLocalFileTime *FILETIME, lpFileTime *FILETIME) (r BOOL, err error)
//sys	LockFile(hFile HANDLE, dwFileOffsetLow uint32, dwFileOffsetHigh uint32, nNumberOfBytesToLockLow uint32, nNumberOfBytesToLockHigh uint32) (r BOOL, err error)
//sys	LockFileEx(hFile HANDLE, dwFlags LOCK_FILE_FLAGS, dwReserved uint32, nNumberOfBytesToLockLow uint32, nNumberOfBytesToLockHigh uint32, lpOverlapped *OVERLAPPED) (r BOOL, err error)
//sys	QueryDosDeviceW(lpDeviceName *PWSTRElement, lpTargetPath *PWSTRElement, ucchMax uint32) (r uint32, err error)
//sys	ReadFile(hFile HANDLE, lpBuffer *uint8, nNumberOfBytesToRead uint32, lpNumberOfBytesRead *uint32, lpOverlapped *OVERLAPPED) (r BOOL, err error)
//sys	ReadFileEx(hFile HANDLE, lpBuffer *uint8, nNumberOfBytesToRead uint32, lpOverlapped *OVERLAPPED, lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE) (r BOOL, err error)
//sys	ReadFileScatter(hFile HANDLE, aSegmentArray *FILE_SEGMENT_ELEMENT, nNumberOfBytesToRead uint32, lpReserved *uint32, lpOverlapped *OVERLAPPED) (r BOOL, err error)
//sys	RemoveDirectoryA(lpPathName *PSTRElement) (r BOOL, err error)
//sys	RemoveDirectoryW(lpPathName *PWSTRElement) (r BOOL, err error)
//sys	SetEndOfFile(hFile HANDLE) (r BOOL, err error)
//sys	SetFileAttributesA(lpFileName *PSTRElement, dwFileAttributes FILE_FLAGS_AND_ATTRIBUTES) (r BOOL, err error)
//sys	SetFileAttributesW(lpFileName *PWSTRElement, dwFileAttributes FILE_FLAGS_AND_ATTRIBUTES) (r BOOL, err error)
//sys	SetFileInformationByHandle(hFile HANDLE, FileInformationClass FILE_INFO_BY_HANDLE_CLASS, lpFileInformation unsafe.Pointer, dwBufferSize uint32) (r BOOL, err error)
//sys	SetFilePointer(hFile HANDLE, lDistanceToMove int32, lpDistanceToMoveHigh *int32, dwMoveMethod SET_FILE_POINTER_MOVE_METHOD) (r uint32, err error)
//sys	SetFilePointerEx(hFile HANDLE, liDistanceToMove int64, lpNewFilePointer *int64, dwMoveMethod SET_FILE_POINTER_MOVE_METHOD) (r BOOL, err error)
//sys	SetFileTime(hFile HANDLE, lpCreationTime *FILETIME, lpLastAccessTime *FILETIME, lpLastWriteTime *FILETIME) (r BOOL, err error)
//sys	SetFileValidData(hFile HANDLE, ValidDataLength int64) (r BOOL, err error)
//sys	UnlockFile(hFile HANDLE, dwFileOffsetLow uint32, dwFileOffsetHigh uint32, nNumberOfBytesToUnlockLow uint32, nNumberOfBytesToUnlockHigh uint32) (r BOOL, err error)
//sys	UnlockFileEx(hFile HANDLE, dwReserved uint32, nNumberOfBytesToUnlockLow uint32, nNumberOfBytesToUnlockHigh uint32, lpOverlapped *OVERLAPPED) (r BOOL, err error)
//sys	WriteFile(hFile HANDLE, lpBuffer *uint8, nNumberOfBytesToWrite uint32, lpNumberOfBytesWritten *uint32, lpOverlapped *OVERLAPPED) (r BOOL, err error)
//sys	WriteFileEx(hFile HANDLE, lpBuffer *uint8, nNumberOfBytesToWrite uint32, lpOverlapped *OVERLAPPED, lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE) (r BOOL, err error)
//sys	WriteFileGather(hFile HANDLE, aSegmentArray *FILE_SEGMENT_ELEMENT, nNumberOfBytesToWrite uint32, lpReserved *uint32, lpOverlapped *OVERLAPPED) (r BOOL, err error)
//sys	GetTempPathW(nBufferLength uint32, lpBuffer *PWSTRElement) (r uint32, err error)
//sys	GetVolumeNameForVolumeMountPointW(lpszVolumeMountPoint *PWSTRElement, lpszVolumeName *PWSTRElement, cchBufferLength uint32) (r BOOL, err error)
//sys	GetVolumePathNamesForVolumeNameW(lpszVolumeName *PWSTRElement, lpszVolumePathNames *PWSTRElement, cchBufferLength uint32, lpcchReturnLength *uint32) (r BOOL, err error)
//sys	CreateFile2(lpFileName *PWSTRElement, dwDesiredAccess uint32, dwShareMode FILE_SHARE_MODE, dwCreationDisposition FILE_CREATION_DISPOSITION, pCreateExParams *CREATEFILE2_EXTENDED_PARAMETERS) (r HANDLE, err error)
//sys	SetFileIoOverlappedRange(FileHandle HANDLE, OverlappedRangeStart *uint8, Length uint32) (r BOOL, err error)
//sys	GetCompressedFileSizeA(lpFileName *PSTRElement, lpFileSizeHigh *uint32) (r uint32, err error)
//sys	GetCompressedFileSizeW(lpFileName *PWSTRElement, lpFileSizeHigh *uint32) (r uint32, err error)
//sys	FindFirstStreamW(lpFileName *PWSTRElement, InfoLevel STREAM_INFO_LEVELS, lpFindStreamData unsafe.Pointer, dwFlags uint32) (r HANDLE, err error)
//sys	FindNextStreamW(hFindStream HANDLE, lpFindStreamData unsafe.Pointer) (r BOOL, err error)
//sys	AreFileApisANSI() (r BOOL)
//sys	GetTempPathA(nBufferLength uint32, lpBuffer *PSTRElement) (r uint32, err error)
//sys	FindFirstFileNameW(lpFileName *PWSTRElement, dwFlags uint32, StringLength *uint32, LinkName *PWSTRElement) (r HANDLE, err error)
//sys	FindNextFileNameW(hFindStream HANDLE, StringLength *uint32, LinkName *PWSTRElement) (r BOOL, err error)
//sys	GetVolumeInformationA(lpRootPathName *PSTRElement, lpVolumeNameBuffer *PSTRElement, nVolumeNameSize uint32, lpVolumeSerialNumber *uint32, lpMaximumComponentLength *uint32, lpFileSystemFlags *uint32, lpFileSystemNameBuffer *PSTRElement, nFileSystemNameSize uint32) (r BOOL, err error)
//sys	GetTempFileNameA(lpPathName *PSTRElement, lpPrefixString *PSTRElement, uUnique uint32, lpTempFileName *PSTRElement) (r uint32, err error)
//sys	SetFileApisToOEM()
//sys	SetFileApisToANSI()
//sys	GetTempPath2W(BufferLength uint32, Buffer *PWSTRElement) (r uint32)
//sys	GetTempPath2A(BufferLength uint32, Buffer *PSTRElement) (r uint32)
//sys	CreateFile3(lpFileName *PWSTRElement, dwDesiredAccess uint32, dwShareMode uint32, dwCreationDisposition uint32, pCreateExParams *CREATEFILE3_EXTENDED_PARAMETERS) (r HANDLE)
//sys	CreateDirectory2A(lpPathName *PSTRElement, dwDesiredAccess uint32, dwShareMode uint32, DirectoryFlags DIRECTORY_FLAGS, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r HANDLE)
//sys	CreateDirectory2W(lpPathName *PWSTRElement, dwDesiredAccess uint32, dwShareMode uint32, DirectoryFlags DIRECTORY_FLAGS, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r HANDLE)
//sys	RemoveDirectory2A(lpPathName *PSTRElement, DirectoryFlags DIRECTORY_FLAGS) (r BOOL)
//sys	RemoveDirectory2W(lpPathName *PWSTRElement, DirectoryFlags DIRECTORY_FLAGS) (r BOOL)
//sys	DeleteFile2A(lpFileName *PSTRElement, Flags uint32) (r BOOL)
//sys	DeleteFile2W(lpFileName *PWSTRElement, Flags uint32) (r BOOL)
//sys	CopyFileFromAppW(lpExistingFileName *PWSTRElement, lpNewFileName *PWSTRElement, bFailIfExists BOOL) (r BOOL) = api-ms-win-core-file-fromapp-l1-1-0.CopyFileFromAppW
//sys	CreateDirectoryFromAppW(lpPathName *PWSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r BOOL) = api-ms-win-core-file-fromapp-l1-1-0.CreateDirectoryFromAppW
//sys	CreateFileFromAppW(lpFileName *PWSTRElement, dwDesiredAccess uint32, dwShareMode uint32, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwCreationDisposition uint32, dwFlagsAndAttributes uint32, hTemplateFile HANDLE) (r HANDLE) = api-ms-win-core-file-fromapp-l1-1-0.CreateFileFromAppW
//sys	CreateFile2FromAppW(lpFileName *PWSTRElement, dwDesiredAccess uint32, dwShareMode uint32, dwCreationDisposition uint32, pCreateExParams *CREATEFILE2_EXTENDED_PARAMETERS) (r HANDLE) = api-ms-win-core-file-fromapp-l1-1-0.CreateFile2FromAppW
//sys	DeleteFileFromAppW(lpFileName *PWSTRElement) (r BOOL) = api-ms-win-core-file-fromapp-l1-1-0.DeleteFileFromAppW
//sys	FindFirstFileExFromAppW(lpFileName *PWSTRElement, fInfoLevelId FINDEX_INFO_LEVELS, lpFindFileData unsafe.Pointer, fSearchOp FINDEX_SEARCH_OPS, lpSearchFilter unsafe.Pointer, dwAdditionalFlags uint32) (r HANDLE) = api-ms-win-core-file-fromapp-l1-1-0.FindFirstFileExFromAppW
//sys	GetFileAttributesExFromAppW(lpFileName *PWSTRElement, fInfoLevelId GET_FILEEX_INFO_LEVELS, lpFileInformation unsafe.Pointer) (r BOOL) = api-ms-win-core-file-fromapp-l1-1-0.GetFileAttributesExFromAppW
//sys	MoveFileFromAppW(lpExistingFileName *PWSTRElement, lpNewFileName *PWSTRElement) (r BOOL) = api-ms-win-core-file-fromapp-l1-1-0.MoveFileFromAppW
//sys	RemoveDirectoryFromAppW(lpPathName *PWSTRElement) (r BOOL) = api-ms-win-core-file-fromapp-l1-1-0.RemoveDirectoryFromAppW
//sys	ReplaceFileFromAppW(lpReplacedFileName *PWSTRElement, lpReplacementFileName *PWSTRElement, lpBackupFileName *PWSTRElement, dwReplaceFlags uint32, lpExclude unsafe.Pointer, lpReserved unsafe.Pointer) (r BOOL) = api-ms-win-core-file-fromapp-l1-1-0.ReplaceFileFromAppW
//sys	SetFileAttributesFromAppW(lpFileName *PWSTRElement, dwFileAttributes uint32) (r BOOL) = api-ms-win-core-file-fromapp-l1-1-0.SetFileAttributesFromAppW
//sys	VerFindFileA(uFlags VER_FIND_FILE_FLAGS, szFileName *PSTRElement, szWinDir *PSTRElement, szAppDir *PSTRElement, szCurDir *PSTRElement, puCurDirLen *uint32, szDestDir *PSTRElement, puDestDirLen *uint32) (r VER_FIND_FILE_STATUS) = version.VerFindFileA
//sys	VerFindFileW(uFlags VER_FIND_FILE_FLAGS, szFileName *PWSTRElement, szWinDir *PWSTRElement, szAppDir *PWSTRElement, szCurDir *PWSTRElement, puCurDirLen *uint32, szDestDir *PWSTRElement, puDestDirLen *uint32) (r VER_FIND_FILE_STATUS) = version.VerFindFileW
//sys	VerInstallFileA(uFlags VER_INSTALL_FILE_FLAGS, szSrcFileName *PSTRElement, szDestFileName *PSTRElement, szSrcDir *PSTRElement, szDestDir *PSTRElement, szCurDir *PSTRElement, szTmpFile *PSTRElement, puTmpFileLen *uint32) (r VER_INSTALL_FILE_STATUS) = version.VerInstallFileA
//sys	VerInstallFileW(uFlags VER_INSTALL_FILE_FLAGS, szSrcFileName *PWSTRElement, szDestFileName *PWSTRElement, szSrcDir *PWSTRElement, szDestDir *PWSTRElement, szCurDir *PWSTRElement, szTmpFile *PWSTRElement, puTmpFileLen *uint32) (r VER_INSTALL_FILE_STATUS) = version.VerInstallFileW
//sys	GetFileVersionInfoSizeA(lptstrFilename *PSTRElement, lpdwHandle *uint32) (r uint32, err error) = version.GetFileVersionInfoSizeA
//sys	GetFileVersionInfoSizeW(lptstrFilename *PWSTRElement, lpdwHandle *uint32) (r uint32, err error) = version.GetFileVersionInfoSizeW
//sys	GetFileVersionInfoA(lptstrFilename *PSTRElement, dwHandle uint32, dwLen uint32, lpData unsafe.Pointer) (r BOOL, err error) = version.GetFileVersionInfoA
//sys	GetFileVersionInfoW(lptstrFilename *PWSTRElement, dwHandle uint32, dwLen uint32, lpData unsafe.Pointer) (r BOOL, err error) = version.GetFileVersionInfoW
//sys	GetFileVersionInfoSizeExA(dwFlags GET_FILE_VERSION_INFO_FLAGS, lpwstrFilename *PSTRElement, lpdwHandle *uint32) (r uint32, err error) = version.GetFileVersionInfoSizeExA
//sys	GetFileVersionInfoSizeExW(dwFlags GET_FILE_VERSION_INFO_FLAGS, lpwstrFilename *PWSTRElement, lpdwHandle *uint32) (r uint32, err error) = version.GetFileVersionInfoSizeExW
//sys	GetFileVersionInfoExA(dwFlags GET_FILE_VERSION_INFO_FLAGS, lpwstrFilename *PSTRElement, dwHandle uint32, dwLen uint32, lpData unsafe.Pointer) (r BOOL, err error) = version.GetFileVersionInfoExA
//sys	GetFileVersionInfoExW(dwFlags GET_FILE_VERSION_INFO_FLAGS, lpwstrFilename *PWSTRElement, dwHandle uint32, dwLen uint32, lpData unsafe.Pointer) (r BOOL, err error) = version.GetFileVersionInfoExW
//sys	VerLanguageNameA(wLang uint32, szLang *PSTRElement, cchLang uint32) (r uint32)
//sys	VerLanguageNameW(wLang uint32, szLang *PWSTRElement, cchLang uint32) (r uint32)
//sys	VerQueryValueA(pBlock unsafe.Pointer, lpSubBlock *PSTRElement, lplpBuffer *unsafe.Pointer, puLen *uint32) (r BOOL) = version.VerQueryValueA
//sys	VerQueryValueW(pBlock unsafe.Pointer, lpSubBlock *PWSTRElement, lplpBuffer *unsafe.Pointer, puLen *uint32) (r BOOL) = version.VerQueryValueW
//sys	LsnEqual(plsn1 *CLS_LSN, plsn2 *CLS_LSN) (r BOOLEAN) = clfsw32.LsnEqual
//sys	LsnLess(plsn1 *CLS_LSN, plsn2 *CLS_LSN) (r BOOLEAN) = clfsw32.LsnLess
//sys	LsnGreater(plsn1 *CLS_LSN, plsn2 *CLS_LSN) (r BOOLEAN) = clfsw32.LsnGreater
//sys	LsnNull(plsn *CLS_LSN) (r BOOLEAN) = clfsw32.LsnNull
//sys	LsnContainer(plsn *CLS_LSN) (r uint32) = clfsw32.LsnContainer
//sys	LsnCreate(cidContainer uint32, offBlock uint32, cRecord uint32) (r CLS_LSN) = clfsw32.LsnCreate
//sys	LsnBlockOffset(plsn *CLS_LSN) (r uint32) = clfsw32.LsnBlockOffset
//sys	LsnRecordSequence(plsn *CLS_LSN) (r uint32) = clfsw32.LsnRecordSequence
//sys	LsnInvalid(plsn *CLS_LSN) (r BOOLEAN) = clfsw32.LsnInvalid
//sys	LsnIncrement(plsn *CLS_LSN) (r CLS_LSN) = clfsw32.LsnIncrement
//sys	CreateLogFile(pszLogFileName *PWSTRElement, fDesiredAccess uint32, dwShareMode FILE_SHARE_MODE, psaLogFile *SECURITY_ATTRIBUTES, fCreateDisposition FILE_CREATION_DISPOSITION, fFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES) (r HANDLE, err error) = clfsw32.CreateLogFile
//sys	DeleteLogByHandle(hLog HANDLE) (r BOOL, err error) = clfsw32.DeleteLogByHandle
//sys	DeleteLogFile(pszLogFileName *PWSTRElement, pvReserved unsafe.Pointer) (r BOOL, err error) = clfsw32.DeleteLogFile
//sys	AddLogContainer(hLog HANDLE, pcbContainer *uint64, pwszContainerPath *PWSTRElement, pReserved unsafe.Pointer) (r BOOL, err error) = clfsw32.AddLogContainer
//sys	AddLogContainerSet(hLog HANDLE, cContainer uint16, pcbContainer *uint64, rgwszContainerPath **PWSTRElement, pReserved unsafe.Pointer) (r BOOL, err error) = clfsw32.AddLogContainerSet
//sys	RemoveLogContainer(hLog HANDLE, pwszContainerPath *PWSTRElement, fForce BOOL, pReserved unsafe.Pointer) (r BOOL, err error) = clfsw32.RemoveLogContainer
//sys	RemoveLogContainerSet(hLog HANDLE, cContainer uint16, rgwszContainerPath **PWSTRElement, fForce BOOL, pReserved unsafe.Pointer) (r BOOL, err error) = clfsw32.RemoveLogContainerSet
//sys	SetLogArchiveTail(hLog HANDLE, plsnArchiveTail *CLS_LSN, pReserved unsafe.Pointer) (r BOOL, err error) = clfsw32.SetLogArchiveTail
//sys	SetEndOfLog(hLog HANDLE, plsnEnd *CLS_LSN, lpOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.SetEndOfLog
//sys	TruncateLog(pvMarshal unsafe.Pointer, plsnEnd *CLS_LSN, lpOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.TruncateLog
//sys	CreateLogContainerScanContext(hLog HANDLE, cFromContainer uint32, cContainers uint32, eScanMode uint8, pcxScan *CLS_SCAN_CONTEXT, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.CreateLogContainerScanContext
//sys	ScanLogContainers(pcxScan *CLS_SCAN_CONTEXT, eScanMode uint8, pReserved unsafe.Pointer) (r BOOL, err error) = clfsw32.ScanLogContainers
//sys	AlignReservedLog(pvMarshal unsafe.Pointer, cReservedRecords uint32, rgcbReservation *int64, pcbAlignReservation *int64) (r BOOL, err error) = clfsw32.AlignReservedLog
//sys	AllocReservedLog(pvMarshal unsafe.Pointer, cReservedRecords uint32, pcbAdjustment *int64) (r BOOL, err error) = clfsw32.AllocReservedLog
//sys	FreeReservedLog(pvMarshal unsafe.Pointer, cReservedRecords uint32, pcbAdjustment *int64) (r BOOL, err error) = clfsw32.FreeReservedLog
//sys	GetLogFileInformation(hLog HANDLE, pinfoBuffer *CLS_INFORMATION, cbBuffer *uint32) (r BOOL, err error) = clfsw32.GetLogFileInformation
//sys	SetLogArchiveMode(hLog HANDLE, eMode CLFS_LOG_ARCHIVE_MODE) (r BOOL, err error) = clfsw32.SetLogArchiveMode
//sys	ReadLogRestartArea(pvMarshal unsafe.Pointer, ppvRestartBuffer *unsafe.Pointer, pcbRestartBuffer *uint32, plsn *CLS_LSN, ppvContext *unsafe.Pointer, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.ReadLogRestartArea
//sys	ReadPreviousLogRestartArea(pvReadContext unsafe.Pointer, ppvRestartBuffer *unsafe.Pointer, pcbRestartBuffer *uint32, plsnRestart *CLS_LSN, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.ReadPreviousLogRestartArea
//sys	WriteLogRestartArea(pvMarshal unsafe.Pointer, pvRestartBuffer unsafe.Pointer, cbRestartBuffer uint32, plsnBase *CLS_LSN, fFlags CLFS_FLAG, pcbWritten *uint32, plsnNext *CLS_LSN, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.WriteLogRestartArea
//sys	GetLogReservationInfo(pvMarshal unsafe.Pointer, pcbRecordNumber *uint32, pcbUserReservation *int64, pcbCommitReservation *int64) (r BOOL) = clfsw32.GetLogReservationInfo
//sys	AdvanceLogBase(pvMarshal unsafe.Pointer, plsnBase *CLS_LSN, fFlags uint32, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.AdvanceLogBase
//sys	CloseAndResetLogFile(hLog HANDLE) (r BOOL, err error) = clfsw32.CloseAndResetLogFile
//sys	CreateLogMarshallingArea(hLog HANDLE, pfnAllocBuffer CLFS_BLOCK_ALLOCATION, pfnFreeBuffer CLFS_BLOCK_DEALLOCATION, pvBlockAllocContext unsafe.Pointer, cbMarshallingBuffer uint32, cMaxWriteBuffers uint32, cMaxReadBuffers uint32, ppvMarshal *unsafe.Pointer) (r BOOL, err error) = clfsw32.CreateLogMarshallingArea
//sys	DeleteLogMarshallingArea(pvMarshal unsafe.Pointer) (r BOOL, err error) = clfsw32.DeleteLogMarshallingArea
//sys	ReserveAndAppendLog(pvMarshal unsafe.Pointer, rgWriteEntries *CLS_WRITE_ENTRY, cWriteEntries uint32, plsnUndoNext *CLS_LSN, plsnPrevious *CLS_LSN, cReserveRecords uint32, rgcbReservation *int64, fFlags CLFS_FLAG, plsn *CLS_LSN, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.ReserveAndAppendLog
//sys	ReserveAndAppendLogAligned(pvMarshal unsafe.Pointer, rgWriteEntries *CLS_WRITE_ENTRY, cWriteEntries uint32, cbEntryAlignment uint32, plsnUndoNext *CLS_LSN, plsnPrevious *CLS_LSN, cReserveRecords uint32, rgcbReservation *int64, fFlags CLFS_FLAG, plsn *CLS_LSN, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.ReserveAndAppendLogAligned
//sys	FlushLogBuffers(pvMarshal unsafe.Pointer, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.FlushLogBuffers
//sys	FlushLogToLsn(pvMarshalContext unsafe.Pointer, plsnFlush *CLS_LSN, plsnLastFlushed *CLS_LSN, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.FlushLogToLsn
//sys	ReadLogRecord(pvMarshal unsafe.Pointer, plsnFirst *CLS_LSN, eContextMode CLFS_CONTEXT_MODE, ppvReadBuffer *unsafe.Pointer, pcbReadBuffer *uint32, peRecordType *uint8, plsnUndoNext *CLS_LSN, plsnPrevious *CLS_LSN, ppvReadContext *unsafe.Pointer, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.ReadLogRecord
//sys	ReadNextLogRecord(pvReadContext unsafe.Pointer, ppvBuffer *unsafe.Pointer, pcbBuffer *uint32, peRecordType *uint8, plsnUser *CLS_LSN, plsnUndoNext *CLS_LSN, plsnPrevious *CLS_LSN, plsnRecord *CLS_LSN, pOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.ReadNextLogRecord
//sys	TerminateReadLog(pvCursorContext unsafe.Pointer) (r BOOL, err error) = clfsw32.TerminateReadLog
//sys	PrepareLogArchive(hLog HANDLE, pszBaseLogFileName *PWSTRElement, cLen uint32, plsnLow *CLS_LSN, plsnHigh *CLS_LSN, pcActualLength *uint32, poffBaseLogFileData *uint64, pcbBaseLogFileLength *uint64, plsnBase *CLS_LSN, plsnLast *CLS_LSN, plsnCurrentArchiveTail *CLS_LSN, ppvArchiveContext *unsafe.Pointer) (r BOOL, err error) = clfsw32.PrepareLogArchive
//sys	ReadLogArchiveMetadata(pvArchiveContext unsafe.Pointer, cbOffset uint32, cbBytesToRead uint32, pbReadBuffer *uint8, pcbBytesRead *uint32) (r BOOL, err error) = clfsw32.ReadLogArchiveMetadata
//sys	GetNextLogArchiveExtent(pvArchiveContext unsafe.Pointer, rgadExtent *CLS_ARCHIVE_DESCRIPTOR, cDescriptors uint32, pcDescriptorsReturned *uint32) (r BOOL, err error) = clfsw32.GetNextLogArchiveExtent
//sys	TerminateLogArchive(pvArchiveContext unsafe.Pointer) (r BOOL, err error) = clfsw32.TerminateLogArchive
//sys	ValidateLog(pszLogFileName *PWSTRElement, psaLogFile *SECURITY_ATTRIBUTES, pinfoBuffer *CLS_INFORMATION, pcbBuffer *uint32) (r BOOL, err error) = clfsw32.ValidateLog
//sys	GetLogContainerName(hLog HANDLE, cidLogicalContainer uint32, pwstrContainerName *PWSTRElement, cLenContainerName uint32, pcActualLenContainerName *uint32) (r BOOL, err error) = clfsw32.GetLogContainerName
//sys	GetLogIoStatistics(hLog HANDLE, pvStatsBuffer unsafe.Pointer, cbStatsBuffer uint32, eStatsClass CLFS_IOSTATS_CLASS, pcbStatsWritten *uint32) (r BOOL, err error) = clfsw32.GetLogIoStatistics
//sys	RegisterManageableLogClient(hLog HANDLE, pCallbacks *LOG_MANAGEMENT_CALLBACKS) (r BOOL, err error) = clfsw32.RegisterManageableLogClient
//sys	DeregisterManageableLogClient(hLog HANDLE) (r BOOL, err error) = clfsw32.DeregisterManageableLogClient
//sys	ReadLogNotification(hLog HANDLE, pNotification *CLFS_MGMT_NOTIFICATION, lpOverlapped *OVERLAPPED) (r BOOL, err error) = clfsw32.ReadLogNotification
//sys	InstallLogPolicy(hLog HANDLE, pPolicy *CLFS_MGMT_POLICY) (r BOOL, err error) = clfsw32.InstallLogPolicy
//sys	RemoveLogPolicy(hLog HANDLE, ePolicyType CLFS_MGMT_POLICY_TYPE) (r BOOL, err error) = clfsw32.RemoveLogPolicy
//sys	QueryLogPolicy(hLog HANDLE, ePolicyType CLFS_MGMT_POLICY_TYPE, pPolicyBuffer *CLFS_MGMT_POLICY, pcbPolicyBuffer *uint32) (r BOOL, err error) = clfsw32.QueryLogPolicy
//sys	SetLogFileSizeWithPolicy(hLog HANDLE, pDesiredSize *uint64, pResultingSize *uint64) (r BOOL, err error) = clfsw32.SetLogFileSizeWithPolicy
//sys	HandleLogFull(hLog HANDLE) (r BOOL, err error) = clfsw32.HandleLogFull
//sys	LogTailAdvanceFailure(hLog HANDLE, dwReason uint32) (r BOOL, err error) = clfsw32.LogTailAdvanceFailure
//sys	RegisterForLogWriteNotification(hLog HANDLE, cbThreshold uint32, fEnable BOOL) (r BOOL, err error) = clfsw32.RegisterForLogWriteNotification
//sys	QueryUsersOnEncryptedFile(lpFileName *PWSTRElement, pUsers **ENCRYPTION_CERTIFICATE_HASH_LIST) (r uint32) = advapi32.QueryUsersOnEncryptedFile
//sys	QueryRecoveryAgentsOnEncryptedFile(lpFileName *PWSTRElement, pRecoveryAgents **ENCRYPTION_CERTIFICATE_HASH_LIST) (r uint32) = advapi32.QueryRecoveryAgentsOnEncryptedFile
//sys	RemoveUsersFromEncryptedFile(lpFileName *PWSTRElement, pHashes *ENCRYPTION_CERTIFICATE_HASH_LIST) (r uint32) = advapi32.RemoveUsersFromEncryptedFile
//sys	AddUsersToEncryptedFile(lpFileName *PWSTRElement, pEncryptionCertificates *ENCRYPTION_CERTIFICATE_LIST) (r uint32) = advapi32.AddUsersToEncryptedFile
//sys	SetUserFileEncryptionKey(pEncryptionCertificate *ENCRYPTION_CERTIFICATE) (r uint32) = advapi32.SetUserFileEncryptionKey
//sys	SetUserFileEncryptionKeyEx(pEncryptionCertificate *ENCRYPTION_CERTIFICATE, dwCapabilities uint32, dwFlags uint32, pvReserved unsafe.Pointer) (r uint32) = advapi32.SetUserFileEncryptionKeyEx
//sys	FreeEncryptionCertificateHashList(pUsers *ENCRYPTION_CERTIFICATE_HASH_LIST) = advapi32.FreeEncryptionCertificateHashList
//sys	EncryptionDisable(DirPath *PWSTRElement, Disable BOOL) (r BOOL, err error) = advapi32.EncryptionDisable
//sys	DuplicateEncryptionInfoFile(SrcFileName *PWSTRElement, DstFileName *PWSTRElement, dwCreationDistribution uint32, dwAttributes uint32, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r uint32) = advapi32.DuplicateEncryptionInfoFile
//sys	GetEncryptedFileMetadata(lpFileName *PWSTRElement, pcbMetadata *uint32, ppbMetadata **uint8) (r uint32) = advapi32.GetEncryptedFileMetadata
//sys	SetEncryptedFileMetadata(lpFileName *PWSTRElement, pbOldMetadata *uint8, pbNewMetadata *uint8, pOwnerHash *ENCRYPTION_CERTIFICATE_HASH, dwOperation uint32, pCertificatesAdded *ENCRYPTION_CERTIFICATE_HASH_LIST) (r uint32) = advapi32.SetEncryptedFileMetadata
//sys	FreeEncryptedFileMetadata(pbMetadata *uint8) = advapi32.FreeEncryptedFileMetadata
//sys	LZStart() (r int32)
//sys	LZDone()
//sys	CopyLZFile(hfSource int32, hfDest int32) (r int32)
//sys	LZCopy(hfSource int32, hfDest int32) (r int32)
//sys	LZInit(hfSource int32) (r int32)
//sys	GetExpandedNameA(lpszSource *PSTRElement, lpszBuffer *PSTRElement) (r int32, err error)
//sys	GetExpandedNameW(lpszSource *PWSTRElement, lpszBuffer *PWSTRElement) (r int32, err error)
//sys	LZOpenFileA(lpFileName *PSTRElement, lpReOpenBuf *OFSTRUCT, wStyle LZOPENFILE_STYLE) (r int32)
//sys	LZOpenFileW(lpFileName *PWSTRElement, lpReOpenBuf *OFSTRUCT, wStyle LZOPENFILE_STYLE) (r int32)
//sys	LZSeek(hFile int32, lOffset int32, iOrigin int32) (r int32)
//sys	LZRead(hFile int32, lpBuffer *PSTRElement, cbRead int32) (r int32)
//sys	LZClose(hFile int32)
//sys	WofShouldCompressBinaries(Volume *PWSTRElement, Algorithm *uint32) (r BOOL) = wofutil.WofShouldCompressBinaries
//sys	WofGetDriverVersion(FileOrVolumeHandle HANDLE, Provider uint32, WofVersion *uint32) (r HRESULT) = wofutil.WofGetDriverVersion
//sys	WofSetFileDataLocation(FileHandle HANDLE, Provider uint32, ExternalFileInfo unsafe.Pointer, Length uint32) (r HRESULT) = wofutil.WofSetFileDataLocation
//sys	WofIsExternalFile(FilePath *PWSTRElement, IsExternalFile *BOOL, Provider *uint32, ExternalFileInfo unsafe.Pointer, BufferLength *uint32) (r HRESULT) = wofutil.WofIsExternalFile
//sys	WofEnumEntries(VolumeName *PWSTRElement, Provider uint32, EnumProc WofEnumEntryProc, UserData unsafe.Pointer) (r HRESULT) = wofutil.WofEnumEntries
//sys	WofWimAddEntry(VolumeName *PWSTRElement, WimPath *PWSTRElement, WimType uint32, WimIndex uint32, DataSourceId *int64) (r HRESULT) = wofutil.WofWimAddEntry
//sys	WofWimEnumFiles(VolumeName *PWSTRElement, DataSourceId int64, EnumProc WofEnumFilesProc, UserData unsafe.Pointer) (r HRESULT) = wofutil.WofWimEnumFiles
//sys	WofWimSuspendEntry(VolumeName *PWSTRElement, DataSourceId int64) (r HRESULT) = wofutil.WofWimSuspendEntry
//sys	WofWimRemoveEntry(VolumeName *PWSTRElement, DataSourceId int64) (r HRESULT) = wofutil.WofWimRemoveEntry
//sys	WofWimUpdateEntry(VolumeName *PWSTRElement, DataSourceId int64, NewWimPath *PWSTRElement) (r HRESULT) = wofutil.WofWimUpdateEntry
//sys	WofFileEnumFiles(VolumeName *PWSTRElement, Algorithm uint32, EnumProc WofEnumFilesProc, UserData unsafe.Pointer) (r HRESULT) = wofutil.WofFileEnumFiles
//sys	TxfLogCreateFileReadContext(LogPath *PWSTRElement, BeginningLsn CLS_LSN, EndingLsn CLS_LSN, TxfFileId *TXF_ID, TxfLogContext *unsafe.Pointer) (r BOOL, err error) = txfw32.TxfLogCreateFileReadContext
//sys	TxfLogCreateRangeReadContext(LogPath *PWSTRElement, BeginningLsn CLS_LSN, EndingLsn CLS_LSN, BeginningVirtualClock *int64, EndingVirtualClock *int64, RecordTypeMask uint32, TxfLogContext *unsafe.Pointer) (r BOOL) = txfw32.TxfLogCreateRangeReadContext
//sys	TxfLogDestroyReadContext(TxfLogContext unsafe.Pointer) (r BOOL, err error) = txfw32.TxfLogDestroyReadContext
//sys	TxfLogReadRecords(TxfLogContext unsafe.Pointer, BufferLength uint32, Buffer unsafe.Pointer, BytesUsed *uint32, RecordCount *uint32) (r BOOL, err error) = txfw32.TxfLogReadRecords
//sys	TxfReadMetadataInfo(FileHandle HANDLE, TxfFileId *TXF_ID, LastLsn *CLS_LSN, TransactionState *uint32, LockingTransaction *Guid) (r BOOL) = txfw32.TxfReadMetadataInfo
//sys	TxfLogRecordGetFileName(RecordBuffer unsafe.Pointer, RecordBufferLengthInBytes uint32, NameBuffer *PWSTRElement, NameBufferLengthInBytes *uint32, TxfId *TXF_ID) (r BOOL) = txfw32.TxfLogRecordGetFileName
//sys	TxfLogRecordGetGenericType(RecordBuffer unsafe.Pointer, RecordBufferLengthInBytes uint32, GenericType *uint32, VirtualClock *int64) (r BOOL) = txfw32.TxfLogRecordGetGenericType
//sys	TxfSetThreadMiniVersionForCreate(MiniVersion uint16) = txfw32.TxfSetThreadMiniVersionForCreate
//sys	TxfGetThreadMiniVersionForCreate(MiniVersion *uint16) = txfw32.TxfGetThreadMiniVersionForCreate
//sys	CreateTransaction(lpTransactionAttributes *SECURITY_ATTRIBUTES, UOW *Guid, CreateOptions uint32, IsolationLevel uint32, IsolationFlags uint32, Timeout uint32, Description *PWSTRElement) (r HANDLE, err error) = ktmw32.CreateTransaction
//sys	OpenTransaction(dwDesiredAccess uint32, TransactionId *Guid) (r HANDLE, err error) = ktmw32.OpenTransaction
//sys	CommitTransaction(TransactionHandle HANDLE) (r BOOL, err error) = ktmw32.CommitTransaction
//sys	CommitTransactionAsync(TransactionHandle HANDLE) (r BOOL, err error) = ktmw32.CommitTransactionAsync
//sys	RollbackTransaction(TransactionHandle HANDLE) (r BOOL, err error) = ktmw32.RollbackTransaction
//sys	RollbackTransactionAsync(TransactionHandle HANDLE) (r BOOL, err error) = ktmw32.RollbackTransactionAsync
//sys	GetTransactionId(TransactionHandle HANDLE, TransactionId *Guid) (r BOOL, err error) = ktmw32.GetTransactionId
//sys	GetTransactionInformation(TransactionHandle HANDLE, Outcome *uint32, IsolationLevel *uint32, IsolationFlags *uint32, Timeout *uint32, BufferLength uint32, Description *PWSTRElement) (r BOOL, err error) = ktmw32.GetTransactionInformation
//sys	SetTransactionInformation(TransactionHandle HANDLE, IsolationLevel uint32, IsolationFlags uint32, Timeout uint32, Description *PWSTRElement) (r BOOL, err error) = ktmw32.SetTransactionInformation
//sys	CreateTransactionManager(lpTransactionAttributes *SECURITY_ATTRIBUTES, LogFileName *PWSTRElement, CreateOptions uint32, CommitStrength uint32) (r HANDLE, err error) = ktmw32.CreateTransactionManager
//sys	OpenTransactionManager(LogFileName *PWSTRElement, DesiredAccess uint32, OpenOptions uint32) (r HANDLE, err error) = ktmw32.OpenTransactionManager
//sys	OpenTransactionManagerById(TransactionManagerId *Guid, DesiredAccess uint32, OpenOptions uint32) (r HANDLE, err error) = ktmw32.OpenTransactionManagerById
//sys	RenameTransactionManager(LogFileName *PWSTRElement, ExistingTransactionManagerGuid *Guid) (r BOOL, err error) = ktmw32.RenameTransactionManager
//sys	RollforwardTransactionManager(TransactionManagerHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.RollforwardTransactionManager
//sys	RecoverTransactionManager(TransactionManagerHandle HANDLE) (r BOOL, err error) = ktmw32.RecoverTransactionManager
//sys	GetCurrentClockTransactionManager(TransactionManagerHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.GetCurrentClockTransactionManager
//sys	GetTransactionManagerId(TransactionManagerHandle HANDLE, TransactionManagerId *Guid) (r BOOL, err error) = ktmw32.GetTransactionManagerId
//sys	CreateResourceManager(lpResourceManagerAttributes *SECURITY_ATTRIBUTES, ResourceManagerId *Guid, CreateOptions uint32, TmHandle HANDLE, Description *PWSTRElement) (r HANDLE, err error) = ktmw32.CreateResourceManager
//sys	OpenResourceManager(dwDesiredAccess uint32, TmHandle HANDLE, ResourceManagerId *Guid) (r HANDLE, err error) = ktmw32.OpenResourceManager
//sys	RecoverResourceManager(ResourceManagerHandle HANDLE) (r BOOL, err error) = ktmw32.RecoverResourceManager
//sys	GetNotificationResourceManager(ResourceManagerHandle HANDLE, TransactionNotification *TRANSACTION_NOTIFICATION, NotificationLength uint32, dwMilliseconds uint32, ReturnLength *uint32) (r BOOL, err error) = ktmw32.GetNotificationResourceManager
//sys	GetNotificationResourceManagerAsync(ResourceManagerHandle HANDLE, TransactionNotification *TRANSACTION_NOTIFICATION, TransactionNotificationLength uint32, ReturnLength *uint32, lpOverlapped *OVERLAPPED) (r BOOL, err error) = ktmw32.GetNotificationResourceManagerAsync
//sys	SetResourceManagerCompletionPort(ResourceManagerHandle HANDLE, IoCompletionPortHandle HANDLE, CompletionKey uintptr) (r BOOL, err error) = ktmw32.SetResourceManagerCompletionPort
//sys	CreateEnlistment(lpEnlistmentAttributes *SECURITY_ATTRIBUTES, ResourceManagerHandle HANDLE, TransactionHandle HANDLE, NotificationMask uint32, CreateOptions uint32, EnlistmentKey unsafe.Pointer) (r HANDLE, err error) = ktmw32.CreateEnlistment
//sys	OpenEnlistment(dwDesiredAccess uint32, ResourceManagerHandle HANDLE, EnlistmentId *Guid) (r HANDLE, err error) = ktmw32.OpenEnlistment
//sys	RecoverEnlistment(EnlistmentHandle HANDLE, EnlistmentKey unsafe.Pointer) (r BOOL, err error) = ktmw32.RecoverEnlistment
//sys	GetEnlistmentRecoveryInformation(EnlistmentHandle HANDLE, BufferSize uint32, Buffer unsafe.Pointer, BufferUsed *uint32) (r BOOL, err error) = ktmw32.GetEnlistmentRecoveryInformation
//sys	GetEnlistmentId(EnlistmentHandle HANDLE, EnlistmentId *Guid) (r BOOL, err error) = ktmw32.GetEnlistmentId
//sys	SetEnlistmentRecoveryInformation(EnlistmentHandle HANDLE, BufferSize uint32, Buffer unsafe.Pointer) (r BOOL, err error) = ktmw32.SetEnlistmentRecoveryInformation
//sys	PrepareEnlistment(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.PrepareEnlistment
//sys	PrePrepareEnlistment(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.PrePrepareEnlistment
//sys	CommitEnlistment(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.CommitEnlistment
//sys	RollbackEnlistment(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.RollbackEnlistment
//sys	PrePrepareComplete(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.PrePrepareComplete
//sys	PrepareComplete(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.PrepareComplete
//sys	ReadOnlyEnlistment(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.ReadOnlyEnlistment
//sys	CommitComplete(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.CommitComplete
//sys	RollbackComplete(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.RollbackComplete
//sys	SinglePhaseReject(EnlistmentHandle HANDLE, TmVirtualClock *int64) (r BOOL, err error) = ktmw32.SinglePhaseReject
//sys	NetShareAdd(servername *PWSTRElement, level uint32, buf *uint8, parm_err *uint32) (r uint32) = netapi32.NetShareAdd
//sys	NetShareEnum(servername *PWSTRElement, level uint32, bufptr **uint8, prefmaxlen uint32, entriesread *uint32, totalentries *uint32, resume_handle *uint32) (r uint32) = netapi32.NetShareEnum
//sys	NetShareEnumSticky(servername *PWSTRElement, level uint32, bufptr **uint8, prefmaxlen uint32, entriesread *uint32, totalentries *uint32, resume_handle *uint32) (r uint32) = netapi32.NetShareEnumSticky
//sys	NetShareGetInfo(servername *PWSTRElement, netname *PWSTRElement, level uint32, bufptr **uint8) (r uint32) = netapi32.NetShareGetInfo
//sys	NetShareSetInfo(servername *PWSTRElement, netname *PWSTRElement, level uint32, buf *uint8, parm_err *uint32) (r uint32) = netapi32.NetShareSetInfo
//sys	NetShareDel(servername *PWSTRElement, netname *PWSTRElement, reserved uint32) (r uint32) = netapi32.NetShareDel
//sys	NetShareDelSticky(servername *PWSTRElement, netname *PWSTRElement, reserved uint32) (r uint32) = netapi32.NetShareDelSticky
//sys	NetShareCheck(servername *PWSTRElement, device *PWSTRElement, typeParam *uint32) (r uint32) = netapi32.NetShareCheck
//sys	NetShareDelEx(servername *PWSTRElement, level uint32, buf *uint8) (r uint32) = netapi32.NetShareDelEx
//sys	NetServerAliasAdd(servername *PWSTRElement, level uint32, buf *uint8) (r uint32) = netapi32.NetServerAliasAdd
//sys	NetServerAliasDel(servername *PWSTRElement, level uint32, buf *uint8) (r uint32) = netapi32.NetServerAliasDel
//sys	NetServerAliasEnum(servername *PWSTRElement, level uint32, bufptr **uint8, prefmaxlen uint32, entriesread *uint32, totalentries *uint32, resumehandle *uint32) (r uint32) = netapi32.NetServerAliasEnum
//sys	NetSessionEnum(servername *PWSTRElement, UncClientName *PWSTRElement, username *PWSTRElement, level uint32, bufptr **uint8, prefmaxlen uint32, entriesread *uint32, totalentries *uint32, resume_handle *uint32) (r uint32) = netapi32.NetSessionEnum
//sys	NetSessionDel(servername *PWSTRElement, UncClientName *PWSTRElement, username *PWSTRElement) (r uint32) = netapi32.NetSessionDel
//sys	NetSessionGetInfo(servername *PWSTRElement, UncClientName *PWSTRElement, username *PWSTRElement, level uint32, bufptr **uint8) (r uint32) = netapi32.NetSessionGetInfo
//sys	NetConnectionEnum(servername *PWSTRElement, qualifier *PWSTRElement, level uint32, bufptr **uint8, prefmaxlen uint32, entriesread *uint32, totalentries *uint32, resume_handle *uint32) (r uint32) = netapi32.NetConnectionEnum
//sys	NetFileClose(servername *PWSTRElement, fileid uint32) (r uint32) = netapi32.NetFileClose
//sys	NetFileEnum(servername *PWSTRElement, basepath *PWSTRElement, username *PWSTRElement, level uint32, bufptr **uint8, prefmaxlen uint32, entriesread *uint32, totalentries *uint32, resume_handle *uintptr) (r uint32) = netapi32.NetFileEnum
//sys	NetFileGetInfo(servername *PWSTRElement, fileid uint32, level uint32, bufptr **uint8) (r uint32) = netapi32.NetFileGetInfo
//sys	NetStatisticsGet(ServerName *int8, Service *int8, Level uint32, Options uint32, Buffer **uint8) (r uint32) = netapi32.NetStatisticsGet
//sys	QueryIoRingCapabilities(capabilities *IORING_CAPABILITIES) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.QueryIoRingCapabilities
//sys	IsIoRingOpSupported(ioRing HIORING, op IORING_OP_CODE) (r BOOL) = api-ms-win-core-ioring-l1-1-0.IsIoRingOpSupported
//sys	CreateIoRing(ioringVersion IORING_VERSION, flags IORING_CREATE_FLAGS, submissionQueueSize uint32, completionQueueSize uint32, h *HIORING) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.CreateIoRing
//sys	GetIoRingInfo(ioRing HIORING, info *IORING_INFO) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.GetIoRingInfo
//sys	SubmitIoRing(ioRing HIORING, waitOperations uint32, milliseconds uint32, submittedEntries *uint32) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.SubmitIoRing
//sys	CloseIoRing(ioRing HIORING) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.CloseIoRing
//sys	PopIoRingCompletion(ioRing HIORING, cqe *IORING_CQE) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.PopIoRingCompletion
//sys	SetIoRingCompletionEvent(ioRing HIORING, hEvent HANDLE) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.SetIoRingCompletionEvent
//sys	BuildIoRingCancelRequest(ioRing HIORING, file IORING_HANDLE_REF, opToCancel uintptr, userData uintptr) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.BuildIoRingCancelRequest
//sys	BuildIoRingReadFile(ioRing HIORING, fileRef IORING_HANDLE_REF, dataRef IORING_BUFFER_REF, numberOfBytesToRead uint32, fileOffset uint64, userData uintptr, sqeFlags IORING_SQE_FLAGS) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.BuildIoRingReadFile
//sys	BuildIoRingRegisterFileHandles(ioRing HIORING, count uint32, handles *HANDLE, userData uintptr) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.BuildIoRingRegisterFileHandles
//sys	BuildIoRingRegisterBuffers(ioRing HIORING, count uint32, buffers *IORING_BUFFER_INFO, userData uintptr) (r HRESULT) = api-ms-win-core-ioring-l1-1-0.BuildIoRingRegisterBuffers
//sys	BuildIoRingWriteFile(ioRing HIORING, fileRef IORING_HANDLE_REF, bufferRef IORING_BUFFER_REF, numberOfBytesToWrite uint32, fileOffset uint64, writeFlags FILE_WRITE_FLAGS, userData uintptr, sqeFlags IORING_SQE_FLAGS) (r HRESULT)
//sys	BuildIoRingFlushFile(ioRing HIORING, fileRef IORING_HANDLE_REF, flushMode FILE_FLUSH_MODE, userData uintptr, sqeFlags IORING_SQE_FLAGS) (r HRESULT)
//sys	BuildIoRingReadFileScatter(ioRing HIORING, fileRef IORING_HANDLE_REF, segmentCount uint32, segmentArray *FILE_SEGMENT_ELEMENT, numberOfBytesToRead uint32, fileOffset uint64, userData uintptr, sqeFlags IORING_SQE_FLAGS) (r HRESULT)
//sys	BuildIoRingWriteFileGather(ioRing HIORING, fileRef IORING_HANDLE_REF, segmentCount uint32, segmentArray *FILE_SEGMENT_ELEMENT, numberOfBytesToWrite uint32, fileOffset uint64, writeFlags FILE_WRITE_FLAGS, userData uintptr, sqeFlags IORING_SQE_FLAGS) (r HRESULT)
//sys	CreateBindLink(virtualPath *PWSTRElement, backingPath *PWSTRElement, createBindLinkFlags CREATE_BIND_LINK_FLAGS, exceptionCount uint32, exceptionPaths **PWSTRElement) (r HRESULT) = bindfltapi.CreateBindLink
//sys	RemoveBindLink(virtualPath *PWSTRElement) (r HRESULT) = bindfltapi.RemoveBindLink
//sys	Wow64EnableWow64FsRedirection(Wow64FsEnableRedirection BOOLEAN) (r BOOLEAN)
//sys	Wow64DisableWow64FsRedirection(OldValue *unsafe.Pointer) (r BOOL, err error)
//sys	Wow64RevertWow64FsRedirection(OlValue unsafe.Pointer) (r BOOL, err error)
//sys	GetBinaryTypeA(lpApplicationName *PSTRElement, lpBinaryType *uint32) (r BOOL, err error)
//sys	GetBinaryTypeW(lpApplicationName *PWSTRElement, lpBinaryType *uint32) (r BOOL, err error)
//sys	GetShortPathNameA(lpszLongPath *PSTRElement, lpszShortPath *PSTRElement, cchBuffer uint32) (r uint32, err error)
//sys	GetLongPathNameTransactedA(lpszShortPath *PSTRElement, lpszLongPath *PSTRElement, cchBuffer uint32, hTransaction HANDLE) (r uint32, err error)
//sys	GetLongPathNameTransactedW(lpszShortPath *PWSTRElement, lpszLongPath *PWSTRElement, cchBuffer uint32, hTransaction HANDLE) (r uint32, err error)
//sys	SetFileCompletionNotificationModes(FileHandle HANDLE, Flags uint8) (r BOOL, err error)
//sys	SetFileShortNameA(hFile HANDLE, lpShortName *PSTRElement) (r BOOL, err error)
//sys	SetFileShortNameW(hFile HANDLE, lpShortName *PWSTRElement) (r BOOL, err error)
//sys	SetTapePosition(hDevice HANDLE, dwPositionMethod TAPE_POSITION_METHOD, dwPartition uint32, dwOffsetLow uint32, dwOffsetHigh uint32, bImmediate BOOL) (r uint32)
//sys	GetTapePosition(hDevice HANDLE, dwPositionType TAPE_POSITION_TYPE, lpdwPartition *uint32, lpdwOffsetLow *uint32, lpdwOffsetHigh *uint32) (r uint32)
//sys	PrepareTape(hDevice HANDLE, dwOperation PREPARE_TAPE_OPERATION, bImmediate BOOL) (r uint32)
//sys	EraseTape(hDevice HANDLE, dwEraseType ERASE_TAPE_TYPE, bImmediate BOOL) (r uint32)
//sys	CreateTapePartition(hDevice HANDLE, dwPartitionMethod CREATE_TAPE_PARTITION_METHOD, dwCount uint32, dwSize uint32) (r uint32)
//sys	WriteTapemark(hDevice HANDLE, dwTapemarkType TAPEMARK_TYPE, dwTapemarkCount uint32, bImmediate BOOL) (r uint32)
//sys	GetTapeStatus(hDevice HANDLE) (r uint32)
//sys	GetTapeParameters(hDevice HANDLE, dwOperation GET_TAPE_DRIVE_PARAMETERS_OPERATION, lpdwSize *uint32, lpTapeInformation unsafe.Pointer) (r uint32)
//sys	SetTapeParameters(hDevice HANDLE, dwOperation TAPE_INFORMATION_TYPE, lpTapeInformation unsafe.Pointer) (r uint32)
//sys	EncryptFileA(lpFileName *PSTRElement) (r BOOL, err error) = advapi32.EncryptFileA
//sys	EncryptFileW(lpFileName *PWSTRElement) (r BOOL, err error) = advapi32.EncryptFileW
//sys	DecryptFileA(lpFileName *PSTRElement, dwReserved uint32) (r BOOL, err error) = advapi32.DecryptFileA
//sys	DecryptFileW(lpFileName *PWSTRElement, dwReserved uint32) (r BOOL, err error) = advapi32.DecryptFileW
//sys	FileEncryptionStatusA(lpFileName *PSTRElement, lpStatus *uint32) (r BOOL, err error) = advapi32.FileEncryptionStatusA
//sys	FileEncryptionStatusW(lpFileName *PWSTRElement, lpStatus *uint32) (r BOOL, err error) = advapi32.FileEncryptionStatusW
//sys	OpenEncryptedFileRawA(lpFileName *PSTRElement, ulFlags uint32, pvContext *unsafe.Pointer) (r uint32) = advapi32.OpenEncryptedFileRawA
//sys	OpenEncryptedFileRawW(lpFileName *PWSTRElement, ulFlags uint32, pvContext *unsafe.Pointer) (r uint32) = advapi32.OpenEncryptedFileRawW
//sys	ReadEncryptedFileRaw(pfExportCallback PFE_EXPORT_FUNC, pvCallbackContext unsafe.Pointer, pvContext unsafe.Pointer) (r uint32) = advapi32.ReadEncryptedFileRaw
//sys	WriteEncryptedFileRaw(pfImportCallback PFE_IMPORT_FUNC, pvCallbackContext unsafe.Pointer, pvContext unsafe.Pointer) (r uint32) = advapi32.WriteEncryptedFileRaw
//sys	CloseEncryptedFileRaw(pvContext unsafe.Pointer) = advapi32.CloseEncryptedFileRaw
//sys	OpenFile(lpFileName *PSTRElement, lpReOpenBuff *OFSTRUCT, uStyle uint32) (r int32, err error)
//sys	BackupRead(hFile HANDLE, lpBuffer *uint8, nNumberOfBytesToRead uint32, lpNumberOfBytesRead *uint32, bAbort BOOL, bProcessSecurity BOOL, lpContext *unsafe.Pointer) (r BOOL, err error)
//sys	BackupSeek(hFile HANDLE, dwLowBytesToSeek uint32, dwHighBytesToSeek uint32, lpdwLowByteSeeked *uint32, lpdwHighByteSeeked *uint32, lpContext *unsafe.Pointer) (r BOOL, err error)
//sys	BackupWrite(hFile HANDLE, lpBuffer *uint8, nNumberOfBytesToWrite uint32, lpNumberOfBytesWritten *uint32, bAbort BOOL, bProcessSecurity BOOL, lpContext *unsafe.Pointer) (r BOOL, err error)
//sys	GetLogicalDriveStringsA(nBufferLength uint32, lpBuffer *PSTRElement) (r uint32, err error)
//sys	SetSearchPathMode(Flags uint32) (r BOOL, err error)
//sys	CreateDirectoryExA(lpTemplateDirectory *PSTRElement, lpNewDirectory *PSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r BOOL, err error)
//sys	CreateDirectoryExW(lpTemplateDirectory *PWSTRElement, lpNewDirectory *PWSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r BOOL, err error)
//sys	CreateDirectoryTransactedA(lpTemplateDirectory *PSTRElement, lpNewDirectory *PSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES, hTransaction HANDLE) (r BOOL, err error)
//sys	CreateDirectoryTransactedW(lpTemplateDirectory *PWSTRElement, lpNewDirectory *PWSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES, hTransaction HANDLE) (r BOOL, err error)
//sys	RemoveDirectoryTransactedA(lpPathName *PSTRElement, hTransaction HANDLE) (r BOOL, err error)
//sys	RemoveDirectoryTransactedW(lpPathName *PWSTRElement, hTransaction HANDLE) (r BOOL, err error)
//sys	GetFullPathNameTransactedA(lpFileName *PSTRElement, nBufferLength uint32, lpBuffer *PSTRElement, lpFilePart **PSTRElement, hTransaction HANDLE) (r uint32, err error)
//sys	GetFullPathNameTransactedW(lpFileName *PWSTRElement, nBufferLength uint32, lpBuffer *PWSTRElement, lpFilePart **PWSTRElement, hTransaction HANDLE) (r uint32, err error)
//sys	DefineDosDeviceA(dwFlags DEFINE_DOS_DEVICE_FLAGS, lpDeviceName *PSTRElement, lpTargetPath *PSTRElement) (r BOOL, err error)
//sys	QueryDosDeviceA(lpDeviceName *PSTRElement, lpTargetPath *PSTRElement, ucchMax uint32) (r uint32, err error)
//sys	CreateFileTransactedA(lpFileName *PSTRElement, dwDesiredAccess uint32, dwShareMode FILE_SHARE_MODE, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwCreationDisposition FILE_CREATION_DISPOSITION, dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES, hTemplateFile HANDLE, hTransaction HANDLE, pusMiniVersion *TXFS_MINIVERSION, lpExtendedParameter unsafe.Pointer) (r HANDLE, err error)
//sys	CreateFileTransactedW(lpFileName *PWSTRElement, dwDesiredAccess uint32, dwShareMode FILE_SHARE_MODE, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwCreationDisposition FILE_CREATION_DISPOSITION, dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES, hTemplateFile HANDLE, hTransaction HANDLE, pusMiniVersion *TXFS_MINIVERSION, lpExtendedParameter unsafe.Pointer) (r HANDLE, err error)
//sys	ReOpenFile(hOriginalFile HANDLE, dwDesiredAccess uint32, dwShareMode FILE_SHARE_MODE, dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES) (r HANDLE, err error)
//sys	SetFileAttributesTransactedA(lpFileName *PSTRElement, dwFileAttributes uint32, hTransaction HANDLE) (r BOOL, err error)
//sys	SetFileAttributesTransactedW(lpFileName *PWSTRElement, dwFileAttributes uint32, hTransaction HANDLE) (r BOOL, err error)
//sys	GetFileAttributesTransactedA(lpFileName *PSTRElement, fInfoLevelId GET_FILEEX_INFO_LEVELS, lpFileInformation unsafe.Pointer, hTransaction HANDLE) (r BOOL, err error)
//sys	GetFileAttributesTransactedW(lpFileName *PWSTRElement, fInfoLevelId GET_FILEEX_INFO_LEVELS, lpFileInformation unsafe.Pointer, hTransaction HANDLE) (r BOOL, err error)
//sys	GetCompressedFileSizeTransactedA(lpFileName *PSTRElement, lpFileSizeHigh *uint32, hTransaction HANDLE) (r uint32, err error)
//sys	GetCompressedFileSizeTransactedW(lpFileName *PWSTRElement, lpFileSizeHigh *uint32, hTransaction HANDLE) (r uint32, err error)
//sys	DeleteFileTransactedA(lpFileName *PSTRElement, hTransaction HANDLE) (r BOOL, err error)
//sys	DeleteFileTransactedW(lpFileName *PWSTRElement, hTransaction HANDLE) (r BOOL, err error)
//sys	CheckNameLegalDOS8Dot3A(lpName *PSTRElement, lpOemName *PSTRElement, OemNameSize uint32, pbNameContainsSpaces *BOOL, pbNameLegal *BOOL) (r BOOL, err error)
//sys	CheckNameLegalDOS8Dot3W(lpName *PWSTRElement, lpOemName *PSTRElement, OemNameSize uint32, pbNameContainsSpaces *BOOL, pbNameLegal *BOOL) (r BOOL, err error)
//sys	FindFirstFileTransactedA(lpFileName *PSTRElement, fInfoLevelId FINDEX_INFO_LEVELS, lpFindFileData unsafe.Pointer, fSearchOp FINDEX_SEARCH_OPS, lpSearchFilter unsafe.Pointer, dwAdditionalFlags uint32, hTransaction HANDLE) (r HANDLE, err error)
//sys	FindFirstFileTransactedW(lpFileName *PWSTRElement, fInfoLevelId FINDEX_INFO_LEVELS, lpFindFileData unsafe.Pointer, fSearchOp FINDEX_SEARCH_OPS, lpSearchFilter unsafe.Pointer, dwAdditionalFlags uint32, hTransaction HANDLE) (r HANDLE, err error)
//sys	CopyFileA(lpExistingFileName *PSTRElement, lpNewFileName *PSTRElement, bFailIfExists BOOL) (r BOOL, err error)
//sys	CopyFileW(lpExistingFileName *PWSTRElement, lpNewFileName *PWSTRElement, bFailIfExists BOOL) (r BOOL, err error)
//sys	CopyFileExA(lpExistingFileName *PSTRElement, lpNewFileName *PSTRElement, lpProgressRoutine LPPROGRESS_ROUTINE, lpData unsafe.Pointer, pbCancel *BOOL, dwCopyFlags COPYFILE_FLAGS) (r BOOL, err error)
//sys	CopyFileExW(lpExistingFileName *PWSTRElement, lpNewFileName *PWSTRElement, lpProgressRoutine LPPROGRESS_ROUTINE, lpData unsafe.Pointer, pbCancel *BOOL, dwCopyFlags COPYFILE_FLAGS) (r BOOL, err error)
//sys	CopyFileTransactedA(lpExistingFileName *PSTRElement, lpNewFileName *PSTRElement, lpProgressRoutine LPPROGRESS_ROUTINE, lpData unsafe.Pointer, pbCancel *BOOL, dwCopyFlags uint32, hTransaction HANDLE) (r BOOL, err error)
//sys	CopyFileTransactedW(lpExistingFileName *PWSTRElement, lpNewFileName *PWSTRElement, lpProgressRoutine LPPROGRESS_ROUTINE, lpData unsafe.Pointer, pbCancel *BOOL, dwCopyFlags uint32, hTransaction HANDLE) (r BOOL, err error)
//sys	CopyFile2(pwszExistingFileName *PWSTRElement, pwszNewFileName *PWSTRElement, pExtendedParameters *COPYFILE2_EXTENDED_PARAMETERS) (r HRESULT)
//sys	MoveFileA(lpExistingFileName *PSTRElement, lpNewFileName *PSTRElement) (r BOOL, err error)
//sys	MoveFileW(lpExistingFileName *PWSTRElement, lpNewFileName *PWSTRElement) (r BOOL, err error)
//sys	MoveFileExA(lpExistingFileName *PSTRElement, lpNewFileName *PSTRElement, dwFlags MOVE_FILE_FLAGS) (r BOOL, err error)
//sys	MoveFileExW(lpExistingFileName *PWSTRElement, lpNewFileName *PWSTRElement, dwFlags MOVE_FILE_FLAGS) (r BOOL, err error)
//sys	MoveFileWithProgressA(lpExistingFileName *PSTRElement, lpNewFileName *PSTRElement, lpProgressRoutine LPPROGRESS_ROUTINE, lpData unsafe.Pointer, dwFlags MOVE_FILE_FLAGS) (r BOOL, err error)
//sys	MoveFileWithProgressW(lpExistingFileName *PWSTRElement, lpNewFileName *PWSTRElement, lpProgressRoutine LPPROGRESS_ROUTINE, lpData unsafe.Pointer, dwFlags MOVE_FILE_FLAGS) (r BOOL, err error)
//sys	MoveFileTransactedA(lpExistingFileName *PSTRElement, lpNewFileName *PSTRElement, lpProgressRoutine LPPROGRESS_ROUTINE, lpData unsafe.Pointer, dwFlags MOVE_FILE_FLAGS, hTransaction HANDLE) (r BOOL, err error)
//sys	MoveFileTransactedW(lpExistingFileName *PWSTRElement, lpNewFileName *PWSTRElement, lpProgressRoutine LPPROGRESS_ROUTINE, lpData unsafe.Pointer, dwFlags MOVE_FILE_FLAGS, hTransaction HANDLE) (r BOOL, err error)
//sys	ReplaceFileA(lpReplacedFileName *PSTRElement, lpReplacementFileName *PSTRElement, lpBackupFileName *PSTRElement, dwReplaceFlags REPLACE_FILE_FLAGS, lpExclude unsafe.Pointer, lpReserved unsafe.Pointer) (r BOOL, err error)
//sys	ReplaceFileW(lpReplacedFileName *PWSTRElement, lpReplacementFileName *PWSTRElement, lpBackupFileName *PWSTRElement, dwReplaceFlags REPLACE_FILE_FLAGS, lpExclude unsafe.Pointer, lpReserved unsafe.Pointer) (r BOOL, err error)
//sys	CreateHardLinkA(lpFileName *PSTRElement, lpExistingFileName *PSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r BOOL, err error)
//sys	CreateHardLinkW(lpFileName *PWSTRElement, lpExistingFileName *PWSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES) (r BOOL, err error)
//sys	CreateHardLinkTransactedA(lpFileName *PSTRElement, lpExistingFileName *PSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES, hTransaction HANDLE) (r BOOL, err error)
//sys	CreateHardLinkTransactedW(lpFileName *PWSTRElement, lpExistingFileName *PWSTRElement, lpSecurityAttributes *SECURITY_ATTRIBUTES, hTransaction HANDLE) (r BOOL, err error)
//sys	FindFirstStreamTransactedW(lpFileName *PWSTRElement, InfoLevel STREAM_INFO_LEVELS, lpFindStreamData unsafe.Pointer, dwFlags uint32, hTransaction HANDLE) (r HANDLE, err error)
//sys	FindFirstFileNameTransactedW(lpFileName *PWSTRElement, dwFlags uint32, StringLength *uint32, LinkName *PWSTRElement, hTransaction HANDLE) (r HANDLE, err error)
//sys	SetVolumeLabelA(lpRootPathName *PSTRElement, lpVolumeName *PSTRElement) (r BOOL, err error)
//sys	SetVolumeLabelW(lpRootPathName *PWSTRElement, lpVolumeName *PWSTRElement) (r BOOL, err error)
//sys	SetFileBandwidthReservation(hFile HANDLE, nPeriodMilliseconds uint32, nBytesPerPeriod uint32, bDiscardable BOOL, lpTransferSize *uint32, lpNumOutstandingRequests *uint32) (r BOOL, err error)
//sys	GetFileBandwidthReservation(hFile HANDLE, lpPeriodMilliseconds *uint32, lpBytesPerPeriod *uint32, pDiscardable *BOOL, lpTransferSize *uint32, lpNumOutstandingRequests *uint32) (r BOOL, err error)
//sys	ReadDirectoryChangesW(hDirectory HANDLE, lpBuffer unsafe.Pointer, nBufferLength uint32, bWatchSubtree BOOL, dwNotifyFilter FILE_NOTIFY_CHANGE, lpBytesReturned *uint32, lpOverlapped *OVERLAPPED, lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE) (r BOOL, err error)
//sys	ReadDirectoryChangesExW(hDirectory HANDLE, lpBuffer unsafe.Pointer, nBufferLength uint32, bWatchSubtree BOOL, dwNotifyFilter FILE_NOTIFY_CHANGE, lpBytesReturned *uint32, lpOverlapped *OVERLAPPED, lpCompletionRoutine LPOVERLAPPED_COMPLETION_ROUTINE, ReadDirectoryNotifyInformationClass READ_DIRECTORY_NOTIFY_INFORMATION_CLASS) (r BOOL, err error)
//sys	FindFirstVolumeA(lpszVolumeName *PSTRElement, cchBufferLength uint32) (r HANDLE, err error)
//sys	FindNextVolumeA(hFindVolume HANDLE, lpszVolumeName *PSTRElement, cchBufferLength uint32) (r BOOL, err error)
//sys	FindFirstVolumeMountPointA(lpszRootPathName *PSTRElement, lpszVolumeMountPoint *PSTRElement, cchBufferLength uint32) (r HANDLE, err error)
//sys	FindFirstVolumeMountPointW(lpszRootPathName *PWSTRElement, lpszVolumeMountPoint *PWSTRElement, cchBufferLength uint32) (r HANDLE, err error)
//sys	FindNextVolumeMountPointA(hFindVolumeMountPoint HANDLE, lpszVolumeMountPoint *PSTRElement, cchBufferLength uint32) (r BOOL, err error)
//sys	FindNextVolumeMountPointW(hFindVolumeMountPoint HANDLE, lpszVolumeMountPoint *PWSTRElement, cchBufferLength uint32) (r BOOL, err error)
//sys	FindVolumeMountPointClose(hFindVolumeMountPoint HANDLE) (r BOOL, err error)
//sys	SetVolumeMountPointA(lpszVolumeMountPoint *PSTRElement, lpszVolumeName *PSTRElement) (r BOOL, err error)
//sys	SetVolumeMountPointW(lpszVolumeMountPoint *PWSTRElement, lpszVolumeName *PWSTRElement) (r BOOL, err error)
//sys	DeleteVolumeMountPointA(lpszVolumeMountPoint *PSTRElement) (r BOOL, err error)
//sys	GetVolumeNameForVolumeMountPointA(lpszVolumeMountPoint *PSTRElement, lpszVolumeName *PSTRElement, cchBufferLength uint32) (r BOOL, err error)
//sys	GetVolumePathNameA(lpszFileName *PSTRElement, lpszVolumePathName *PSTRElement, cchBufferLength uint32) (r BOOL, err error)
//sys	GetVolumePathNamesForVolumeNameA(lpszVolumeName *PSTRElement, lpszVolumePathNames *PSTRElement, cchBufferLength uint32, lpcchReturnLength *uint32) (r BOOL, err error)
//sys	GetFileInformationByHandleEx(hFile HANDLE, FileInformationClass FILE_INFO_BY_HANDLE_CLASS, lpFileInformation unsafe.Pointer, dwBufferSize uint32) (r BOOL, err error)
//sys	GetFileInformationByName(FileName *PWSTRElement, FileInformationClass FILE_INFO_BY_NAME_CLASS, FileInfoBuffer unsafe.Pointer, FileInfoBufferSize uint32) (r BOOL)
//sys	OpenFileById(hVolumeHint HANDLE, lpFileId *FILE_ID_DESCRIPTOR, dwDesiredAccess uint32, dwShareMode FILE_SHARE_MODE, lpSecurityAttributes *SECURITY_ATTRIBUTES, dwFlagsAndAttributes FILE_FLAGS_AND_ATTRIBUTES) (r HANDLE, err error)
//sys	CreateSymbolicLinkA(lpSymlinkFileName *PSTRElement, lpTargetFileName *PSTRElement, dwFlags SYMBOLIC_LINK_FLAGS) (r BOOLEAN, err error)
//sys	CreateSymbolicLinkW(lpSymlinkFileName *PWSTRElement, lpTargetFileName *PWSTRElement, dwFlags SYMBOLIC_LINK_FLAGS) (r BOOLEAN, err error)
//sys	CreateSymbolicLinkTransactedA(lpSymlinkFileName *PSTRElement, lpTargetFileName *PSTRElement, dwFlags SYMBOLIC_LINK_FLAGS, hTransaction HANDLE) (r BOOLEAN, err error)
//sys	CreateSymbolicLinkTransactedW(lpSymlinkFileName *PWSTRElement, lpTargetFileName *PWSTRElement, dwFlags SYMBOLIC_LINK_FLAGS, hTransaction HANDLE) (r BOOLEAN, err error)

// Types used in generated APIs for

type FIND_FIRST_EX_FLAGS uint32

const (
	FIND_FIRST_EX_CASE_SENSITIVE       FIND_FIRST_EX_FLAGS = 0x1
	FIND_FIRST_EX_LARGE_FETCH          FIND_FIRST_EX_FLAGS = 0x2
	FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY FIND_FIRST_EX_FLAGS = 0x4
)

type DEFINE_DOS_DEVICE_FLAGS uint32

const (
	DDD_RAW_TARGET_PATH       DEFINE_DOS_DEVICE_FLAGS = 0x1
	DDD_REMOVE_DEFINITION     DEFINE_DOS_DEVICE_FLAGS = 0x2
	DDD_EXACT_MATCH_ON_REMOVE DEFINE_DOS_DEVICE_FLAGS = 0x4
	DDD_NO_BROADCAST_SYSTEM   DEFINE_DOS_DEVICE_FLAGS = 0x8
	DDD_LUID_BROADCAST_DRIVE  DEFINE_DOS_DEVICE_FLAGS = 0x10
)

type FILE_FLAGS_AND_ATTRIBUTES uint32

const (
	FILE_ATTRIBUTE_READONLY              FILE_FLAGS_AND_ATTRIBUTES = 0x1
	FILE_ATTRIBUTE_HIDDEN                FILE_FLAGS_AND_ATTRIBUTES = 0x2
	FILE_ATTRIBUTE_SYSTEM                FILE_FLAGS_AND_ATTRIBUTES = 0x4
	FILE_ATTRIBUTE_DIRECTORY             FILE_FLAGS_AND_ATTRIBUTES = 0x10
	FILE_ATTRIBUTE_ARCHIVE               FILE_FLAGS_AND_ATTRIBUTES = 0x20
	FILE_ATTRIBUTE_DEVICE                FILE_FLAGS_AND_ATTRIBUTES = 0x40
	FILE_ATTRIBUTE_NORMAL                FILE_FLAGS_AND_ATTRIBUTES = 0x80
	FILE_ATTRIBUTE_TEMPORARY             FILE_FLAGS_AND_ATTRIBUTES = 0x100
	FILE_ATTRIBUTE_SPARSE_FILE           FILE_FLAGS_AND_ATTRIBUTES = 0x200
	FILE_ATTRIBUTE_REPARSE_POINT         FILE_FLAGS_AND_ATTRIBUTES = 0x400
	FILE_ATTRIBUTE_COMPRESSED            FILE_FLAGS_AND_ATTRIBUTES = 0x800
	FILE_ATTRIBUTE_OFFLINE               FILE_FLAGS_AND_ATTRIBUTES = 0x1000
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED   FILE_FLAGS_AND_ATTRIBUTES = 0x2000
	FILE_ATTRIBUTE_ENCRYPTED             FILE_FLAGS_AND_ATTRIBUTES = 0x4000
	FILE_ATTRIBUTE_INTEGRITY_STREAM      FILE_FLAGS_AND_ATTRIBUTES = 0x8000
	FILE_ATTRIBUTE_VIRTUAL               FILE_FLAGS_AND_ATTRIBUTES = 0x10000
	FILE_ATTRIBUTE_NO_SCRUB_DATA         FILE_FLAGS_AND_ATTRIBUTES = 0x20000
	FILE_ATTRIBUTE_EA                    FILE_FLAGS_AND_ATTRIBUTES = 0x40000
	FILE_ATTRIBUTE_PINNED                FILE_FLAGS_AND_ATTRIBUTES = 0x80000
	FILE_ATTRIBUTE_UNPINNED              FILE_FLAGS_AND_ATTRIBUTES = 0x100000
	FILE_ATTRIBUTE_RECALL_ON_OPEN        FILE_FLAGS_AND_ATTRIBUTES = 0x40000
	FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS FILE_FLAGS_AND_ATTRIBUTES = 0x400000
	FILE_FLAG_WRITE_THROUGH              FILE_FLAGS_AND_ATTRIBUTES = 0x80000000
	FILE_FLAG_OVERLAPPED                 FILE_FLAGS_AND_ATTRIBUTES = 0x40000000
	FILE_FLAG_NO_BUFFERING               FILE_FLAGS_AND_ATTRIBUTES = 0x20000000
	FILE_FLAG_RANDOM_ACCESS              FILE_FLAGS_AND_ATTRIBUTES = 0x10000000
	FILE_FLAG_SEQUENTIAL_SCAN            FILE_FLAGS_AND_ATTRIBUTES = 0x8000000
	FILE_FLAG_DELETE_ON_CLOSE            FILE_FLAGS_AND_ATTRIBUTES = 0x4000000
	FILE_FLAG_BACKUP_SEMANTICS           FILE_FLAGS_AND_ATTRIBUTES = 0x2000000
	FILE_FLAG_POSIX_SEMANTICS            FILE_FLAGS_AND_ATTRIBUTES = 0x1000000
	FILE_FLAG_SESSION_AWARE              FILE_FLAGS_AND_ATTRIBUTES = 0x800000
	FILE_FLAG_OPEN_REPARSE_POINT         FILE_FLAGS_AND_ATTRIBUTES = 0x200000
	FILE_FLAG_OPEN_NO_RECALL             FILE_FLAGS_AND_ATTRIBUTES = 0x100000
	FILE_FLAG_FIRST_PIPE_INSTANCE        FILE_FLAGS_AND_ATTRIBUTES = 0x80000
	PIPE_ACCESS_DUPLEX                   FILE_FLAGS_AND_ATTRIBUTES = 0x3
	PIPE_ACCESS_INBOUND                  FILE_FLAGS_AND_ATTRIBUTES = 0x1
	PIPE_ACCESS_OUTBOUND                 FILE_FLAGS_AND_ATTRIBUTES = 0x2
	SECURITY_ANONYMOUS                   FILE_FLAGS_AND_ATTRIBUTES = 0x0
	SECURITY_IDENTIFICATION              FILE_FLAGS_AND_ATTRIBUTES = 0x10000
	SECURITY_IMPERSONATION               FILE_FLAGS_AND_ATTRIBUTES = 0x20000
	SECURITY_DELEGATION                  FILE_FLAGS_AND_ATTRIBUTES = 0x30000
	SECURITY_CONTEXT_TRACKING            FILE_FLAGS_AND_ATTRIBUTES = 0x40000
	SECURITY_EFFECTIVE_ONLY              FILE_FLAGS_AND_ATTRIBUTES = 0x80000
	SECURITY_SQOS_PRESENT                FILE_FLAGS_AND_ATTRIBUTES = 0x100000
	SECURITY_VALID_SQOS_FLAGS            FILE_FLAGS_AND_ATTRIBUTES = 0x1f0000
)

type REG_VALUE_TYPE uint32

const (
	REG_NONE                       REG_VALUE_TYPE = 0x0
	REG_SZ                         REG_VALUE_TYPE = 0x1
	REG_EXPAND_SZ                  REG_VALUE_TYPE = 0x2
	REG_BINARY                     REG_VALUE_TYPE = 0x3
	REG_DWORD                      REG_VALUE_TYPE = 0x4
	REG_DWORD_LITTLE_ENDIAN        REG_VALUE_TYPE = 0x4
	REG_DWORD_BIG_ENDIAN           REG_VALUE_TYPE = 0x5
	REG_LINK                       REG_VALUE_TYPE = 0x6
	REG_MULTI_SZ                   REG_VALUE_TYPE = 0x7
	REG_RESOURCE_LIST              REG_VALUE_TYPE = 0x8
	REG_FULL_RESOURCE_DESCRIPTOR   REG_VALUE_TYPE = 0x9
	REG_RESOURCE_REQUIREMENTS_LIST REG_VALUE_TYPE = 0xa
	REG_QWORD                      REG_VALUE_TYPE = 0xb
	REG_QWORD_LITTLE_ENDIAN        REG_VALUE_TYPE = 0xb
)

type BCRYPT_OPERATION uint32

const (
	BCRYPT_CIPHER_OPERATION                BCRYPT_OPERATION = 0x1
	BCRYPT_HASH_OPERATION                  BCRYPT_OPERATION = 0x2
	BCRYPT_ASYMMETRIC_ENCRYPTION_OPERATION BCRYPT_OPERATION = 0x4
	BCRYPT_SECRET_AGREEMENT_OPERATION      BCRYPT_OPERATION = 0x8
	BCRYPT_SIGNATURE_OPERATION             BCRYPT_OPERATION = 0x10
	BCRYPT_RNG_OPERATION                   BCRYPT_OPERATION = 0x20
)

type NCRYPT_OPERATION uint32

const (
	NCRYPT_CIPHER_OPERATION                NCRYPT_OPERATION = 0x1
	NCRYPT_HASH_OPERATION                  NCRYPT_OPERATION = 0x2
	NCRYPT_ASYMMETRIC_ENCRYPTION_OPERATION NCRYPT_OPERATION = 0x4
	NCRYPT_SECRET_AGREEMENT_OPERATION      NCRYPT_OPERATION = 0x8
	NCRYPT_SIGNATURE_OPERATION             NCRYPT_OPERATION = 0x10
)

type CERT_FIND_FLAGS uint32

const (
	CERT_FIND_ANY                         CERT_FIND_FLAGS = 0x0
	CERT_FIND_CERT_ID                     CERT_FIND_FLAGS = 0x100000
	CERT_FIND_CTL_USAGE                   CERT_FIND_FLAGS = 0xa0000
	CERT_FIND_ENHKEY_USAGE                CERT_FIND_FLAGS = 0xa0000
	CERT_FIND_EXISTING                    CERT_FIND_FLAGS = 0xd0000
	CERT_FIND_HASH                        CERT_FIND_FLAGS = 0x10000
	CERT_FIND_HAS_PRIVATE_KEY             CERT_FIND_FLAGS = 0x150000
	CERT_FIND_ISSUER_ATTR                 CERT_FIND_FLAGS = 0x30004
	CERT_FIND_ISSUER_NAME                 CERT_FIND_FLAGS = 0x20004
	CERT_FIND_ISSUER_OF                   CERT_FIND_FLAGS = 0xc0000
	CERT_FIND_ISSUER_STR                  CERT_FIND_FLAGS = 0x80004
	CERT_FIND_KEY_IDENTIFIER              CERT_FIND_FLAGS = 0xf0000
	CERT_FIND_KEY_SPEC                    CERT_FIND_FLAGS = 0x90000
	CERT_FIND_MD5_HASH                    CERT_FIND_FLAGS = 0x40000
	CERT_FIND_PROPERTY                    CERT_FIND_FLAGS = 0x50000
	CERT_FIND_PUBLIC_KEY                  CERT_FIND_FLAGS = 0x60000
	CERT_FIND_SHA1_HASH                   CERT_FIND_FLAGS = 0x10000
	CERT_FIND_SIGNATURE_HASH              CERT_FIND_FLAGS = 0xe0000
	CERT_FIND_SUBJECT_ATTR                CERT_FIND_FLAGS = 0x30007
	CERT_FIND_SUBJECT_CERT                CERT_FIND_FLAGS = 0xb0000
	CERT_FIND_SUBJECT_NAME                CERT_FIND_FLAGS = 0x20007
	CERT_FIND_SUBJECT_STR                 CERT_FIND_FLAGS = 0x80007
	CERT_FIND_CROSS_CERT_DIST_POINTS      CERT_FIND_FLAGS = 0x110000
	CERT_FIND_PUBKEY_MD5_HASH             CERT_FIND_FLAGS = 0x120000
	CERT_FIND_SHA256_HASH                 CERT_FIND_FLAGS = 0x160000
	CERT_FIND_SHA1_SHA256_HASH            CERT_FIND_FLAGS = 0x170000
	CERT_FIND_SUBJECT_STR_A               CERT_FIND_FLAGS = 0x70007
	CERT_FIND_SUBJECT_STR_W               CERT_FIND_FLAGS = 0x80007
	CERT_FIND_ISSUER_STR_A                CERT_FIND_FLAGS = 0x70004
	CERT_FIND_ISSUER_STR_W                CERT_FIND_FLAGS = 0x80004
	CERT_FIND_SUBJECT_INFO_ACCESS         CERT_FIND_FLAGS = 0x130000
	CERT_FIND_HASH_STR                    CERT_FIND_FLAGS = 0x140000
	CERT_FIND_OPTIONAL_ENHKEY_USAGE_FLAG  CERT_FIND_FLAGS = 0x1
	CERT_FIND_EXT_ONLY_ENHKEY_USAGE_FLAG  CERT_FIND_FLAGS = 0x2
	CERT_FIND_PROP_ONLY_ENHKEY_USAGE_FLAG CERT_FIND_FLAGS = 0x4
	CERT_FIND_NO_ENHKEY_USAGE_FLAG        CERT_FIND_FLAGS = 0x8
	CERT_FIND_OR_ENHKEY_USAGE_FLAG        CERT_FIND_FLAGS = 0x10
	CERT_FIND_VALID_ENHKEY_USAGE_FLAG     CERT_FIND_FLAGS = 0x20
	CERT_FIND_OPTIONAL_CTL_USAGE_FLAG     CERT_FIND_FLAGS = 0x1
	CERT_FIND_EXT_ONLY_CTL_USAGE_FLAG     CERT_FIND_FLAGS = 0x2
	CERT_FIND_PROP_ONLY_CTL_USAGE_FLAG    CERT_FIND_FLAGS = 0x4
	CERT_FIND_NO_CTL_USAGE_FLAG           CERT_FIND_FLAGS = 0x8
	CERT_FIND_OR_CTL_USAGE_FLAG           CERT_FIND_FLAGS = 0x10
	CERT_FIND_VALID_CTL_USAGE_FLAG        CERT_FIND_FLAGS = 0x20
)

type CERT_QUERY_OBJECT_TYPE uint32

const (
	CERT_QUERY_OBJECT_FILE CERT_QUERY_OBJECT_TYPE = 0x1
	CERT_QUERY_OBJECT_BLOB CERT_QUERY_OBJECT_TYPE = 0x2
)

type CERT_QUERY_CONTENT_TYPE uint32

const (
	CERT_QUERY_CONTENT_CERT               CERT_QUERY_CONTENT_TYPE = 0x1
	CERT_QUERY_CONTENT_CTL                CERT_QUERY_CONTENT_TYPE = 0x2
	CERT_QUERY_CONTENT_CRL                CERT_QUERY_CONTENT_TYPE = 0x3
	CERT_QUERY_CONTENT_SERIALIZED_STORE   CERT_QUERY_CONTENT_TYPE = 0x4
	CERT_QUERY_CONTENT_SERIALIZED_CERT    CERT_QUERY_CONTENT_TYPE = 0x5
	CERT_QUERY_CONTENT_SERIALIZED_CTL     CERT_QUERY_CONTENT_TYPE = 0x6
	CERT_QUERY_CONTENT_SERIALIZED_CRL     CERT_QUERY_CONTENT_TYPE = 0x7
	CERT_QUERY_CONTENT_PKCS7_SIGNED       CERT_QUERY_CONTENT_TYPE = 0x8
	CERT_QUERY_CONTENT_PKCS7_UNSIGNED     CERT_QUERY_CONTENT_TYPE = 0x9
	CERT_QUERY_CONTENT_PKCS7_SIGNED_EMBED CERT_QUERY_CONTENT_TYPE = 0xa
	CERT_QUERY_CONTENT_PKCS10             CERT_QUERY_CONTENT_TYPE = 0xb
	CERT_QUERY_CONTENT_PFX                CERT_QUERY_CONTENT_TYPE = 0xc
	CERT_QUERY_CONTENT_CERT_PAIR          CERT_QUERY_CONTENT_TYPE = 0xd
	CERT_QUERY_CONTENT_PFX_AND_LOAD       CERT_QUERY_CONTENT_TYPE = 0xe
)

type CERT_QUERY_CONTENT_TYPE_FLAGS uint32

const (
	CERT_QUERY_CONTENT_FLAG_CERT               CERT_QUERY_CONTENT_TYPE_FLAGS = 0x2
	CERT_QUERY_CONTENT_FLAG_CTL                CERT_QUERY_CONTENT_TYPE_FLAGS = 0x4
	CERT_QUERY_CONTENT_FLAG_CRL                CERT_QUERY_CONTENT_TYPE_FLAGS = 0x8
	CERT_QUERY_CONTENT_FLAG_SERIALIZED_STORE   CERT_QUERY_CONTENT_TYPE_FLAGS = 0x10
	CERT_QUERY_CONTENT_FLAG_SERIALIZED_CERT    CERT_QUERY_CONTENT_TYPE_FLAGS = 0x20
	CERT_QUERY_CONTENT_FLAG_SERIALIZED_CTL     CERT_QUERY_CONTENT_TYPE_FLAGS = 0x40
	CERT_QUERY_CONTENT_FLAG_SERIALIZED_CRL     CERT_QUERY_CONTENT_TYPE_FLAGS = 0x80
	CERT_QUERY_CONTENT_FLAG_PKCS7_SIGNED       CERT_QUERY_CONTENT_TYPE_FLAGS = 0x100
	CERT_QUERY_CONTENT_FLAG_PKCS7_UNSIGNED     CERT_QUERY_CONTENT_TYPE_FLAGS = 0x200
	CERT_QUERY_CONTENT_FLAG_PKCS7_SIGNED_EMBED CERT_QUERY_CONTENT_TYPE_FLAGS = 0x400
	CERT_QUERY_CONTENT_FLAG_PKCS10             CERT_QUERY_CONTENT_TYPE_FLAGS = 0x800
	CERT_QUERY_CONTENT_FLAG_PFX                CERT_QUERY_CONTENT_TYPE_FLAGS = 0x1000
	CERT_QUERY_CONTENT_FLAG_CERT_PAIR          CERT_QUERY_CONTENT_TYPE_FLAGS = 0x2000
	CERT_QUERY_CONTENT_FLAG_PFX_AND_LOAD       CERT_QUERY_CONTENT_TYPE_FLAGS = 0x4000
	CERT_QUERY_CONTENT_FLAG_ALL                CERT_QUERY_CONTENT_TYPE_FLAGS = 0x3ffe
	CERT_QUERY_CONTENT_FLAG_ALL_ISSUER_CERT    CERT_QUERY_CONTENT_TYPE_FLAGS = 0x332
)

type CERT_QUERY_FORMAT_TYPE uint32

const (
	CERT_QUERY_FORMAT_BINARY                CERT_QUERY_FORMAT_TYPE = 0x1
	CERT_QUERY_FORMAT_BASE64_ENCODED        CERT_QUERY_FORMAT_TYPE = 0x2
	CERT_QUERY_FORMAT_ASN_ASCII_HEX_ENCODED CERT_QUERY_FORMAT_TYPE = 0x3
)

type CERT_QUERY_FORMAT_TYPE_FLAGS uint32

const (
	CERT_QUERY_FORMAT_FLAG_BINARY                CERT_QUERY_FORMAT_TYPE_FLAGS = 0x2
	CERT_QUERY_FORMAT_FLAG_BASE64_ENCODED        CERT_QUERY_FORMAT_TYPE_FLAGS = 0x4
	CERT_QUERY_FORMAT_FLAG_ASN_ASCII_HEX_ENCODED CERT_QUERY_FORMAT_TYPE_FLAGS = 0x8
	CERT_QUERY_FORMAT_FLAG_ALL                   CERT_QUERY_FORMAT_TYPE_FLAGS = 0xe
)

type CERT_QUERY_ENCODING_TYPE uint32

const (
	X509_ASN_ENCODING   CERT_QUERY_ENCODING_TYPE = 0x1
	PKCS_7_ASN_ENCODING CERT_QUERY_ENCODING_TYPE = 0x10000
)

type CERT_STRING_TYPE uint32

const (
	CERT_SIMPLE_NAME_STR CERT_STRING_TYPE = 0x1
	CERT_OID_NAME_STR    CERT_STRING_TYPE = 0x2
	CERT_X500_NAME_STR   CERT_STRING_TYPE = 0x3
)

type BCRYPT_TABLE uint32

const (
	CRYPT_LOCAL  BCRYPT_TABLE = 0x1
	CRYPT_DOMAIN BCRYPT_TABLE = 0x2
)

type CERT_KEY_SPEC uint32

const (
	AT_KEYEXCHANGE       CERT_KEY_SPEC = 0x1
	AT_SIGNATURE         CERT_KEY_SPEC = 0x2
	CERT_NCRYPT_KEY_SPEC CERT_KEY_SPEC = 0xffffffff
)

type BCRYPT_INTERFACE uint32

const (
	BCRYPT_ASYMMETRIC_ENCRYPTION_INTERFACE BCRYPT_INTERFACE = 0x3
	BCRYPT_CIPHER_INTERFACE                BCRYPT_INTERFACE = 0x1
	BCRYPT_HASH_INTERFACE                  BCRYPT_INTERFACE = 0x2
	BCRYPT_RNG_INTERFACE                   BCRYPT_INTERFACE = 0x6
	BCRYPT_SECRET_AGREEMENT_INTERFACE      BCRYPT_INTERFACE = 0x4
	BCRYPT_SIGNATURE_INTERFACE             BCRYPT_INTERFACE = 0x5
	NCRYPT_KEY_STORAGE_INTERFACE           BCRYPT_INTERFACE = 0x10001
	NCRYPT_SCHANNEL_INTERFACE              BCRYPT_INTERFACE = 0x10002
	NCRYPT_SCHANNEL_SIGNATURE_INTERFACE    BCRYPT_INTERFACE = 0x10003
)

type BCRYPT_FLAGS uint32

const (
	BCRYPT_BLOCK_PADDING BCRYPT_FLAGS = 0x1
	BCRYPT_PAD_NONE      BCRYPT_FLAGS = 0x1
	BCRYPT_PAD_OAEP      BCRYPT_FLAGS = 0x4
	BCRYPT_PAD_PKCS1     BCRYPT_FLAGS = 0x2
	BCRYPT_PAD_PSS       BCRYPT_FLAGS = 0x8
)

type NCRYPT_FLAGS uint32

const (
	NCRYPT_SILENT_FLAG                    NCRYPT_FLAGS = 0x40
	NCRYPT_NO_PADDING_FLAG                NCRYPT_FLAGS = 0x1
	NCRYPT_PAD_OAEP_FLAG                  NCRYPT_FLAGS = 0x4
	NCRYPT_PAD_PKCS1_FLAG                 NCRYPT_FLAGS = 0x2
	NCRYPT_REGISTER_NOTIFY_FLAG           NCRYPT_FLAGS = 0x1
	NCRYPT_UNREGISTER_NOTIFY_FLAG         NCRYPT_FLAGS = 0x2
	NCRYPT_MACHINE_KEY_FLAG               NCRYPT_FLAGS = 0x20
	NCRYPT_UNPROTECT_NO_DECRYPT           NCRYPT_FLAGS = 0x1
	NCRYPT_OVERWRITE_KEY_FLAG             NCRYPT_FLAGS = 0x80
	NCRYPT_NO_KEY_VALIDATION              NCRYPT_FLAGS = 0x8
	NCRYPT_WRITE_KEY_TO_LEGACY_STORE_FLAG NCRYPT_FLAGS = 0x200
	NCRYPT_PAD_PSS_FLAG                   NCRYPT_FLAGS = 0x8
	NCRYPT_PERSIST_FLAG                   NCRYPT_FLAGS = 0x80000000
	NCRYPT_PERSIST_ONLY_FLAG              NCRYPT_FLAGS = 0x40000000
)

type CRYPT_STRING uint32

const (
	CRYPT_STRING_BASE64HEADER        CRYPT_STRING = 0x0
	CRYPT_STRING_BASE64              CRYPT_STRING = 0x1
	CRYPT_STRING_BINARY              CRYPT_STRING = 0x2
	CRYPT_STRING_BASE64REQUESTHEADER CRYPT_STRING = 0x3
	CRYPT_STRING_HEX                 CRYPT_STRING = 0x4
	CRYPT_STRING_HEXASCII            CRYPT_STRING = 0x5
	CRYPT_STRING_BASE64X509CRLHEADER CRYPT_STRING = 0x9
	CRYPT_STRING_HEXADDR             CRYPT_STRING = 0xa
	CRYPT_STRING_HEXASCIIADDR        CRYPT_STRING = 0xb
	CRYPT_STRING_HEXRAW              CRYPT_STRING = 0xc
	CRYPT_STRING_STRICT              CRYPT_STRING = 0x20000000
	CRYPT_STRING_BASE64_ANY          CRYPT_STRING = 0x6
	CRYPT_STRING_ANY                 CRYPT_STRING = 0x7
	CRYPT_STRING_HEX_ANY             CRYPT_STRING = 0x8
)

type CRYPT_IMPORT_PUBLIC_KEY_FLAGS uint32

const (
	CRYPT_OID_INFO_PUBKEY_SIGN_KEY_FLAG    CRYPT_IMPORT_PUBLIC_KEY_FLAGS = 0x80000000
	CRYPT_OID_INFO_PUBKEY_ENCRYPT_KEY_FLAG CRYPT_IMPORT_PUBLIC_KEY_FLAGS = 0x40000000
)

type CRYPT_XML_FLAGS uint32

const (
	CRYPT_XML_FLAG_DISABLE_EXTENSIONS CRYPT_XML_FLAGS = 0x10000000
	CRYPT_XML_FLAG_NO_SERIALIZE       CRYPT_XML_FLAGS = 0x80000000
	CRYPT_XML_SIGN_ADD_KEYVALUE       CRYPT_XML_FLAGS = 0x1
)

type CRYPT_ENCODE_OBJECT_FLAGS uint32

const (
	CRYPT_ENCODE_ALLOC_FLAG                            CRYPT_ENCODE_OBJECT_FLAGS = 0x8000
	CRYPT_ENCODE_ENABLE_PUNYCODE_FLAG                  CRYPT_ENCODE_OBJECT_FLAGS = 0x20000
	CRYPT_UNICODE_NAME_ENCODE_DISABLE_CHECK_TYPE_FLAG  CRYPT_ENCODE_OBJECT_FLAGS = 0x40000000
	CRYPT_UNICODE_NAME_ENCODE_ENABLE_T61_UNICODE_FLAG  CRYPT_ENCODE_OBJECT_FLAGS = 0x80000000
	CRYPT_UNICODE_NAME_ENCODE_ENABLE_UTF8_UNICODE_FLAG CRYPT_ENCODE_OBJECT_FLAGS = 0x20000000
)

type CRYPT_ACQUIRE_FLAGS uint32

const (
	CRYPT_ACQUIRE_CACHE_FLAG             CRYPT_ACQUIRE_FLAGS = 0x1
	CRYPT_ACQUIRE_USE_PROV_INFO_FLAG     CRYPT_ACQUIRE_FLAGS = 0x2
	CRYPT_ACQUIRE_COMPARE_KEY_FLAG       CRYPT_ACQUIRE_FLAGS = 0x4
	CRYPT_ACQUIRE_NO_HEALING             CRYPT_ACQUIRE_FLAGS = 0x8
	CRYPT_ACQUIRE_SILENT_FLAG            CRYPT_ACQUIRE_FLAGS = 0x40
	CRYPT_ACQUIRE_WINDOW_HANDLE_FLAG     CRYPT_ACQUIRE_FLAGS = 0x80
	CRYPT_ACQUIRE_NCRYPT_KEY_FLAGS_MASK  CRYPT_ACQUIRE_FLAGS = 0x70000
	CRYPT_ACQUIRE_ALLOW_NCRYPT_KEY_FLAG  CRYPT_ACQUIRE_FLAGS = 0x10000
	CRYPT_ACQUIRE_PREFER_NCRYPT_KEY_FLAG CRYPT_ACQUIRE_FLAGS = 0x20000
	CRYPT_ACQUIRE_ONLY_NCRYPT_KEY_FLAG   CRYPT_ACQUIRE_FLAGS = 0x40000
)

type CRYPT_GET_URL_FLAGS uint32

const (
	CRYPT_GET_URL_FROM_PROPERTY         CRYPT_GET_URL_FLAGS = 0x1
	CRYPT_GET_URL_FROM_EXTENSION        CRYPT_GET_URL_FLAGS = 0x2
	CRYPT_GET_URL_FROM_UNAUTH_ATTRIBUTE CRYPT_GET_URL_FLAGS = 0x4
	CRYPT_GET_URL_FROM_AUTH_ATTRIBUTE   CRYPT_GET_URL_FLAGS = 0x8
)

type CERT_STORE_SAVE_AS uint32

const (
	CERT_STORE_SAVE_AS_PKCS7 CERT_STORE_SAVE_AS = 0x2
	CERT_STORE_SAVE_AS_STORE CERT_STORE_SAVE_AS = 0x1
)

type BCRYPT_QUERY_PROVIDER_MODE uint32

const (
	CRYPT_ANY BCRYPT_QUERY_PROVIDER_MODE = 0x4
	CRYPT_UM  BCRYPT_QUERY_PROVIDER_MODE = 0x1
	CRYPT_KM  BCRYPT_QUERY_PROVIDER_MODE = 0x2
	CRYPT_MM  BCRYPT_QUERY_PROVIDER_MODE = 0x3
)

type CERT_FIND_CHAIN_IN_STORE_FLAGS uint32

const (
	CERT_CHAIN_FIND_BY_ISSUER_COMPARE_KEY_FLAG    CERT_FIND_CHAIN_IN_STORE_FLAGS = 0x1
	CERT_CHAIN_FIND_BY_ISSUER_COMPLEX_CHAIN_FLAG  CERT_FIND_CHAIN_IN_STORE_FLAGS = 0x2
	CERT_CHAIN_FIND_BY_ISSUER_CACHE_ONLY_FLAG     CERT_FIND_CHAIN_IN_STORE_FLAGS = 0x8000
	CERT_CHAIN_FIND_BY_ISSUER_CACHE_ONLY_URL_FLAG CERT_FIND_CHAIN_IN_STORE_FLAGS = 0x4
	CERT_CHAIN_FIND_BY_ISSUER_LOCAL_MACHINE_FLAG  CERT_FIND_CHAIN_IN_STORE_FLAGS = 0x8
	CERT_CHAIN_FIND_BY_ISSUER_NO_KEY_FLAG         CERT_FIND_CHAIN_IN_STORE_FLAGS = 0x4000
)

type CERT_CONTROL_STORE_FLAGS uint32

const (
	CERT_STORE_CTRL_COMMIT_FORCE_FLAG             CERT_CONTROL_STORE_FLAGS = 0x1
	CERT_STORE_CTRL_COMMIT_CLEAR_FLAG             CERT_CONTROL_STORE_FLAGS = 0x2
	CERT_STORE_CTRL_INHIBIT_DUPLICATE_HANDLE_FLAG CERT_CONTROL_STORE_FLAGS = 0x1
)

type BCRYPT_OPEN_ALGORITHM_PROVIDER_FLAGS uint32

const (
	BCRYPT_ALG_HANDLE_HMAC_FLAG BCRYPT_OPEN_ALGORITHM_PROVIDER_FLAGS = 0x8
	BCRYPT_PROV_DISPATCH        BCRYPT_OPEN_ALGORITHM_PROVIDER_FLAGS = 0x1
	BCRYPT_HASH_REUSABLE_FLAG   BCRYPT_OPEN_ALGORITHM_PROVIDER_FLAGS = 0x20
)

type CERT_STORE_SAVE_TO uint32

const (
	CERT_STORE_SAVE_TO_FILE       CERT_STORE_SAVE_TO = 0x1
	CERT_STORE_SAVE_TO_FILENAME   CERT_STORE_SAVE_TO = 0x4
	CERT_STORE_SAVE_TO_FILENAME_A CERT_STORE_SAVE_TO = 0x3
	CERT_STORE_SAVE_TO_FILENAME_W CERT_STORE_SAVE_TO = 0x4
	CERT_STORE_SAVE_TO_MEMORY     CERT_STORE_SAVE_TO = 0x2
)

type CRYPT_SET_PROV_PARAM_ID uint32

const (
	PP_CLIENT_HWND            CRYPT_SET_PROV_PARAM_ID = 0x1
	PP_DELETEKEY              CRYPT_SET_PROV_PARAM_ID = 0x18
	PP_KEYEXCHANGE_ALG        CRYPT_SET_PROV_PARAM_ID = 0xe
	PP_KEYEXCHANGE_PIN        CRYPT_SET_PROV_PARAM_ID = 0x20
	PP_KEYEXCHANGE_KEYSIZE    CRYPT_SET_PROV_PARAM_ID = 0xc
	PP_KEYSET_SEC_DESCR       CRYPT_SET_PROV_PARAM_ID = 0x8
	PP_PIN_PROMPT_STRING      CRYPT_SET_PROV_PARAM_ID = 0x2c
	PP_ROOT_CERTSTORE         CRYPT_SET_PROV_PARAM_ID = 0x2e
	PP_SIGNATURE_ALG          CRYPT_SET_PROV_PARAM_ID = 0xf
	PP_SIGNATURE_PIN          CRYPT_SET_PROV_PARAM_ID = 0x21
	PP_SIGNATURE_KEYSIZE      CRYPT_SET_PROV_PARAM_ID = 0xd
	PP_UI_PROMPT              CRYPT_SET_PROV_PARAM_ID = 0x15
	PP_USE_HARDWARE_RNG       CRYPT_SET_PROV_PARAM_ID = 0x26
	PP_USER_CERTSTORE         CRYPT_SET_PROV_PARAM_ID = 0x2a
	PP_SECURE_KEYEXCHANGE_PIN CRYPT_SET_PROV_PARAM_ID = 0x2f
	PP_SECURE_SIGNATURE_PIN   CRYPT_SET_PROV_PARAM_ID = 0x30
	PP_SMARTCARD_READER       CRYPT_SET_PROV_PARAM_ID = 0x2b
)

type CRYPT_KEY_PARAM_ID uint32

const (
	KP_ALGID         CRYPT_KEY_PARAM_ID = 0x7
	KP_CERTIFICATE   CRYPT_KEY_PARAM_ID = 0x1a
	KP_PERMISSIONS   CRYPT_KEY_PARAM_ID = 0x6
	KP_SALT          CRYPT_KEY_PARAM_ID = 0x2
	KP_SALT_EX       CRYPT_KEY_PARAM_ID = 0xa
	KP_BLOCKLEN      CRYPT_KEY_PARAM_ID = 0x8
	KP_GET_USE_COUNT CRYPT_KEY_PARAM_ID = 0x2a
	KP_KEYLEN        CRYPT_KEY_PARAM_ID = 0x9
)

type CRYPT_KEY_FLAGS uint32

const (
	CRYPT_EXPORTABLE                   CRYPT_KEY_FLAGS = 0x1
	CRYPT_USER_PROTECTED               CRYPT_KEY_FLAGS = 0x2
	CRYPT_ARCHIVABLE                   CRYPT_KEY_FLAGS = 0x4000
	CRYPT_CREATE_IV                    CRYPT_KEY_FLAGS = 0x200
	CRYPT_CREATE_SALT                  CRYPT_KEY_FLAGS = 0x4
	CRYPT_DATA_KEY                     CRYPT_KEY_FLAGS = 0x800
	CRYPT_FORCE_KEY_PROTECTION_HIGH    CRYPT_KEY_FLAGS = 0x8000
	CRYPT_KEK                          CRYPT_KEY_FLAGS = 0x400
	CRYPT_INITIATOR                    CRYPT_KEY_FLAGS = 0x40
	CRYPT_NO_SALT                      CRYPT_KEY_FLAGS = 0x10
	CRYPT_ONLINE                       CRYPT_KEY_FLAGS = 0x80
	CRYPT_PREGEN                       CRYPT_KEY_FLAGS = 0x40
	CRYPT_RECIPIENT                    CRYPT_KEY_FLAGS = 0x10
	CRYPT_SF                           CRYPT_KEY_FLAGS = 0x100
	CRYPT_SGCKEY                       CRYPT_KEY_FLAGS = 0x2000
	CRYPT_VOLATILE                     CRYPT_KEY_FLAGS = 0x1000
	CRYPT_MACHINE_KEYSET               CRYPT_KEY_FLAGS = 0x20
	CRYPT_USER_KEYSET                  CRYPT_KEY_FLAGS = 0x1000
	PKCS12_PREFER_CNG_KSP              CRYPT_KEY_FLAGS = 0x100
	PKCS12_ALWAYS_CNG_KSP              CRYPT_KEY_FLAGS = 0x200
	PKCS12_ALLOW_OVERWRITE_KEY         CRYPT_KEY_FLAGS = 0x4000
	PKCS12_NO_PERSIST_KEY              CRYPT_KEY_FLAGS = 0x8000
	PKCS12_INCLUDE_EXTENDED_PROPERTIES CRYPT_KEY_FLAGS = 0x10
	CRYPT_OAEP                         CRYPT_KEY_FLAGS = 0x40
	CRYPT_BLOB_VER3                    CRYPT_KEY_FLAGS = 0x80
	CRYPT_DESTROYKEY                   CRYPT_KEY_FLAGS = 0x4
	CRYPT_SSL2_FALLBACK                CRYPT_KEY_FLAGS = 0x2
	CRYPT_Y_ONLY                       CRYPT_KEY_FLAGS = 0x1
	CRYPT_IPSEC_HMAC_KEY               CRYPT_KEY_FLAGS = 0x100
	CERT_SET_KEY_PROV_HANDLE_PROP_ID   CRYPT_KEY_FLAGS = 0x1
	CERT_SET_KEY_CONTEXT_PROP_ID       CRYPT_KEY_FLAGS = 0x1
)

type CRYPT_MSG_TYPE uint32

const (
	CMSG_DATA                 CRYPT_MSG_TYPE = 0x1
	CMSG_SIGNED               CRYPT_MSG_TYPE = 0x2
	CMSG_ENVELOPED            CRYPT_MSG_TYPE = 0x3
	CMSG_SIGNED_AND_ENVELOPED CRYPT_MSG_TYPE = 0x4
	CMSG_HASHED               CRYPT_MSG_TYPE = 0x5
)

type CERT_OPEN_STORE_FLAGS uint32

const (
	CERT_STORE_BACKUP_RESTORE_FLAG              CERT_OPEN_STORE_FLAGS = 0x800
	CERT_STORE_CREATE_NEW_FLAG                  CERT_OPEN_STORE_FLAGS = 0x2000
	CERT_STORE_DEFER_CLOSE_UNTIL_LAST_FREE_FLAG CERT_OPEN_STORE_FLAGS = 0x4
	CERT_STORE_DELETE_FLAG                      CERT_OPEN_STORE_FLAGS = 0x10
	CERT_STORE_ENUM_ARCHIVED_FLAG               CERT_OPEN_STORE_FLAGS = 0x200
	CERT_STORE_MAXIMUM_ALLOWED_FLAG             CERT_OPEN_STORE_FLAGS = 0x1000
	CERT_STORE_NO_CRYPT_RELEASE_FLAG            CERT_OPEN_STORE_FLAGS = 0x1
	CERT_STORE_OPEN_EXISTING_FLAG               CERT_OPEN_STORE_FLAGS = 0x4000
	CERT_STORE_READONLY_FLAG                    CERT_OPEN_STORE_FLAGS = 0x8000
	CERT_STORE_SET_LOCALIZED_NAME_FLAG          CERT_OPEN_STORE_FLAGS = 0x2
	CERT_STORE_SHARE_CONTEXT_FLAG               CERT_OPEN_STORE_FLAGS = 0x80
	CERT_STORE_UPDATE_KEYID_FLAG                CERT_OPEN_STORE_FLAGS = 0x400
)

type CRYPT_DEFAULT_CONTEXT_FLAGS uint32

const (
	CRYPT_DEFAULT_CONTEXT_AUTO_RELEASE_FLAG CRYPT_DEFAULT_CONTEXT_FLAGS = 0x1
	CRYPT_DEFAULT_CONTEXT_PROCESS_FLAG      CRYPT_DEFAULT_CONTEXT_FLAGS = 0x2
)

type CRYPT_VERIFY_CERT_FLAGS uint32

const (
	CRYPT_VERIFY_CERT_SIGN_DISABLE_MD2_MD4_FLAG          CRYPT_VERIFY_CERT_FLAGS = 0x1
	CRYPT_VERIFY_CERT_SIGN_SET_STRONG_PROPERTIES_FLAG    CRYPT_VERIFY_CERT_FLAGS = 0x2
	CRYPT_VERIFY_CERT_SIGN_RETURN_STRONG_PROPERTIES_FLAG CRYPT_VERIFY_CERT_FLAGS = 0x4
)

type CRYPT_SET_HASH_PARAM uint32

const (
	HP_HMAC_INFO CRYPT_SET_HASH_PARAM = 0x5
	HP_HASHVAL   CRYPT_SET_HASH_PARAM = 0x2
)

type CERT_CREATE_SELFSIGN_FLAGS uint32

const (
	CERT_CREATE_SELFSIGN_NO_KEY_INFO CERT_CREATE_SELFSIGN_FLAGS = 0x2
	CERT_CREATE_SELFSIGN_NO_SIGN     CERT_CREATE_SELFSIGN_FLAGS = 0x1
)

type CRYPT_DEFAULT_CONTEXT_TYPE uint32

const (
	CRYPT_DEFAULT_CONTEXT_CERT_SIGN_OID       CRYPT_DEFAULT_CONTEXT_TYPE = 0x1
	CRYPT_DEFAULT_CONTEXT_MULTI_CERT_SIGN_OID CRYPT_DEFAULT_CONTEXT_TYPE = 0x2
)

type BCRYPT_RESOLVE_PROVIDERS_FLAGS uint32

const (
	CRYPT_ALL_FUNCTIONS BCRYPT_RESOLVE_PROVIDERS_FLAGS = 0x1
	CRYPT_ALL_PROVIDERS BCRYPT_RESOLVE_PROVIDERS_FLAGS = 0x2
)

type CERT_FIND_TYPE uint32

const (
	CTL_FIND_ANY             CERT_FIND_TYPE = 0x0
	CTL_FIND_SHA1_HASH       CERT_FIND_TYPE = 0x1
	CTL_FIND_MD5_HASH        CERT_FIND_TYPE = 0x2
	CTL_FIND_USAGE           CERT_FIND_TYPE = 0x3
	CTL_FIND_SAME_USAGE_FLAG CERT_FIND_TYPE = 0x1
	CTL_FIND_EXISTING        CERT_FIND_TYPE = 0x5
	CTL_FIND_SUBJECT         CERT_FIND_TYPE = 0x4
)

type CRYPT_FIND_FLAGS uint32

const (
	CRYPT_FIND_USER_KEYSET_FLAG    CRYPT_FIND_FLAGS = 0x1
	CRYPT_FIND_MACHINE_KEYSET_FLAG CRYPT_FIND_FLAGS = 0x2
	CRYPT_FIND_SILENT_KEYSET_FLAG  CRYPT_FIND_FLAGS = 0x40
)

type BCRYPTGENRANDOM_FLAGS uint32

const (
	BCRYPT_RNG_USE_ENTROPY_IN_BUFFER BCRYPTGENRANDOM_FLAGS = 0x1
	BCRYPT_USE_SYSTEM_PREFERRED_RNG  BCRYPTGENRANDOM_FLAGS = 0x2
)

type SIGNER_SIGN_FLAGS uint32

const (
	SIG_APPEND                        SIGNER_SIGN_FLAGS = 0x1000
	SPC_INC_PE_RESOURCES_FLAG         SIGNER_SIGN_FLAGS = 0x80
	SPC_INC_PE_DEBUG_INFO_FLAG        SIGNER_SIGN_FLAGS = 0x40
	SPC_INC_PE_IMPORT_ADDR_TABLE_FLAG SIGNER_SIGN_FLAGS = 0x20
	SPC_EXC_PE_PAGE_HASHES_FLAG       SIGNER_SIGN_FLAGS = 0x10
	SPC_INC_PE_PAGE_HASHES_FLAG       SIGNER_SIGN_FLAGS = 0x100
	SPC_DIGEST_GENERATE_FLAG          SIGNER_SIGN_FLAGS = 0x200
	SPC_DIGEST_SIGN_FLAG              SIGNER_SIGN_FLAGS = 0x400
	SPC_DIGEST_SIGN_EX_FLAG           SIGNER_SIGN_FLAGS = 0x4000
)

type SIGNER_TIMESTAMP_FLAGS uint32

const (
	SIGNER_TIMESTAMP_AUTHENTICODE SIGNER_TIMESTAMP_FLAGS = 0x1
	SIGNER_TIMESTAMP_RFC3161      SIGNER_TIMESTAMP_FLAGS = 0x2
)

type ALG_ID uint32

const (
	CALG_MD2                     ALG_ID = 0x8001
	CALG_MD4                     ALG_ID = 0x8002
	CALG_MD5                     ALG_ID = 0x8003
	CALG_SHA                     ALG_ID = 0x8004
	CALG_SHA1                    ALG_ID = 0x8004
	CALG_MAC                     ALG_ID = 0x8005
	CALG_RSA_SIGN                ALG_ID = 0x2400
	CALG_DSS_SIGN                ALG_ID = 0x2200
	CALG_NO_SIGN                 ALG_ID = 0x2000
	CALG_RSA_KEYX                ALG_ID = 0xa400
	CALG_DES                     ALG_ID = 0x6601
	CALG_3DES_112                ALG_ID = 0x6609
	CALG_3DES                    ALG_ID = 0x6603
	CALG_DESX                    ALG_ID = 0x6604
	CALG_RC2                     ALG_ID = 0x6602
	CALG_RC4                     ALG_ID = 0x6801
	CALG_SEAL                    ALG_ID = 0x6802
	CALG_DH_SF                   ALG_ID = 0xaa01
	CALG_DH_EPHEM                ALG_ID = 0xaa02
	CALG_AGREEDKEY_ANY           ALG_ID = 0xaa03
	CALG_KEA_KEYX                ALG_ID = 0xaa04
	CALG_HUGHES_MD5              ALG_ID = 0xa003
	CALG_SKIPJACK                ALG_ID = 0x660a
	CALG_TEK                     ALG_ID = 0x660b
	CALG_CYLINK_MEK              ALG_ID = 0x660c
	CALG_SSL3_SHAMD5             ALG_ID = 0x8008
	CALG_SSL3_MASTER             ALG_ID = 0x4c01
	CALG_SCHANNEL_MASTER_HASH    ALG_ID = 0x4c02
	CALG_SCHANNEL_MAC_KEY        ALG_ID = 0x4c03
	CALG_SCHANNEL_ENC_KEY        ALG_ID = 0x4c07
	CALG_PCT1_MASTER             ALG_ID = 0x4c04
	CALG_SSL2_MASTER             ALG_ID = 0x4c05
	CALG_TLS1_MASTER             ALG_ID = 0x4c06
	CALG_RC5                     ALG_ID = 0x660d
	CALG_HMAC                    ALG_ID = 0x8009
	CALG_TLS1PRF                 ALG_ID = 0x800a
	CALG_HASH_REPLACE_OWF        ALG_ID = 0x800b
	CALG_AES_128                 ALG_ID = 0x660e
	CALG_AES_192                 ALG_ID = 0x660f
	CALG_AES_256                 ALG_ID = 0x6610
	CALG_AES                     ALG_ID = 0x6611
	CALG_SHA_256                 ALG_ID = 0x800c
	CALG_SHA_384                 ALG_ID = 0x800d
	CALG_SHA_512                 ALG_ID = 0x800e
	CALG_ECDH                    ALG_ID = 0xaa05
	CALG_ECDH_EPHEM              ALG_ID = 0xae06
	CALG_ECMQV                   ALG_ID = 0xa001
	CALG_ECDSA                   ALG_ID = 0x2203
	CALG_NULLCIPHER              ALG_ID = 0x6000
	CALG_THIRDPARTY_KEY_EXCHANGE ALG_ID = 0xb000
	CALG_THIRDPARTY_SIGNATURE    ALG_ID = 0x3000
	CALG_THIRDPARTY_CIPHER       ALG_ID = 0x7000
	CALG_THIRDPARTY_HASH         ALG_ID = 0x9000
)

type OBJECT_SECURITY_INFORMATION uint32

const (
	ATTRIBUTE_SECURITY_INFORMATION        OBJECT_SECURITY_INFORMATION = 0x20
	BACKUP_SECURITY_INFORMATION           OBJECT_SECURITY_INFORMATION = 0x10000
	DACL_SECURITY_INFORMATION             OBJECT_SECURITY_INFORMATION = 0x4
	GROUP_SECURITY_INFORMATION            OBJECT_SECURITY_INFORMATION = 0x2
	LABEL_SECURITY_INFORMATION            OBJECT_SECURITY_INFORMATION = 0x10
	OWNER_SECURITY_INFORMATION            OBJECT_SECURITY_INFORMATION = 0x1
	PROTECTED_DACL_SECURITY_INFORMATION   OBJECT_SECURITY_INFORMATION = 0x80000000
	PROTECTED_SACL_SECURITY_INFORMATION   OBJECT_SECURITY_INFORMATION = 0x40000000
	SACL_SECURITY_INFORMATION             OBJECT_SECURITY_INFORMATION = 0x8
	SCOPE_SECURITY_INFORMATION            OBJECT_SECURITY_INFORMATION = 0x40
	UNPROTECTED_DACL_SECURITY_INFORMATION OBJECT_SECURITY_INFORMATION = 0x20000000
	UNPROTECTED_SACL_SECURITY_INFORMATION OBJECT_SECURITY_INFORMATION = 0x10000000
)

type GET_FILE_VERSION_INFO_FLAGS uint32

const (
	FILE_VER_GET_LOCALISED  GET_FILE_VERSION_INFO_FLAGS = 0x1
	FILE_VER_GET_NEUTRAL    GET_FILE_VERSION_INFO_FLAGS = 0x2
	FILE_VER_GET_PREFETCHED GET_FILE_VERSION_INFO_FLAGS = 0x4
)

type VER_FIND_FILE_FLAGS uint32

const (
	VFFF_ISSHAREDFILE VER_FIND_FILE_FLAGS = 0x1
)

type VER_FIND_FILE_STATUS uint32

const (
	VFF_CURNEDEST    VER_FIND_FILE_STATUS = 0x1
	VFF_FILEINUSE    VER_FIND_FILE_STATUS = 0x2
	VFF_BUFFTOOSMALL VER_FIND_FILE_STATUS = 0x4
)

type VER_INSTALL_FILE_FLAGS uint32

const (
	VIFF_FORCEINSTALL  VER_INSTALL_FILE_FLAGS = 0x1
	VIFF_DONTDELETEOLD VER_INSTALL_FILE_FLAGS = 0x2
)

type VER_INSTALL_FILE_STATUS uint32

const (
	VIF_TEMPFILE          VER_INSTALL_FILE_STATUS = 0x1
	VIF_MISMATCH          VER_INSTALL_FILE_STATUS = 0x2
	VIF_SRCOLD            VER_INSTALL_FILE_STATUS = 0x4
	VIF_DIFFLANG          VER_INSTALL_FILE_STATUS = 0x8
	VIF_DIFFCODEPG        VER_INSTALL_FILE_STATUS = 0x10
	VIF_DIFFTYPE          VER_INSTALL_FILE_STATUS = 0x20
	VIF_WRITEPROT         VER_INSTALL_FILE_STATUS = 0x40
	VIF_FILEINUSE         VER_INSTALL_FILE_STATUS = 0x80
	VIF_OUTOFSPACE        VER_INSTALL_FILE_STATUS = 0x100
	VIF_ACCESSVIOLATION   VER_INSTALL_FILE_STATUS = 0x200
	VIF_SHARINGVIOLATION  VER_INSTALL_FILE_STATUS = 0x400
	VIF_CANNOTCREATE      VER_INSTALL_FILE_STATUS = 0x800
	VIF_CANNOTDELETE      VER_INSTALL_FILE_STATUS = 0x1000
	VIF_CANNOTRENAME      VER_INSTALL_FILE_STATUS = 0x2000
	VIF_CANNOTDELETECUR   VER_INSTALL_FILE_STATUS = 0x4000
	VIF_OUTOFMEMORY       VER_INSTALL_FILE_STATUS = 0x8000
	VIF_CANNOTREADSRC     VER_INSTALL_FILE_STATUS = 0x10000
	VIF_CANNOTREADDST     VER_INSTALL_FILE_STATUS = 0x20000
	VIF_BUFFTOOSMALL      VER_INSTALL_FILE_STATUS = 0x40000
	VIF_CANNOTLOADLZ32    VER_INSTALL_FILE_STATUS = 0x80000
	VIF_CANNOTLOADCABINET VER_INSTALL_FILE_STATUS = 0x100000
)

type FILE_CREATION_DISPOSITION uint32

const (
	CREATE_NEW        FILE_CREATION_DISPOSITION = 0x1
	CREATE_ALWAYS     FILE_CREATION_DISPOSITION = 0x2
	OPEN_EXISTING     FILE_CREATION_DISPOSITION = 0x3
	OPEN_ALWAYS       FILE_CREATION_DISPOSITION = 0x4
	TRUNCATE_EXISTING FILE_CREATION_DISPOSITION = 0x5
)

type FILE_SHARE_MODE uint32

const (
	FILE_SHARE_NONE   FILE_SHARE_MODE = 0x0
	FILE_SHARE_DELETE FILE_SHARE_MODE = 0x4
	FILE_SHARE_READ   FILE_SHARE_MODE = 0x1
	FILE_SHARE_WRITE  FILE_SHARE_MODE = 0x2
)

type CLFS_FLAG uint32

const (
	CLFS_FLAG_FORCE_APPEND    CLFS_FLAG = 0x1
	CLFS_FLAG_FORCE_FLUSH     CLFS_FLAG = 0x2
	CLFS_FLAG_NO_FLAGS        CLFS_FLAG = 0x0
	CLFS_FLAG_USE_RESERVATION CLFS_FLAG = 0x4
)

type SET_FILE_POINTER_MOVE_METHOD uint32

const (
	FILE_BEGIN   SET_FILE_POINTER_MOVE_METHOD = 0x0
	FILE_CURRENT SET_FILE_POINTER_MOVE_METHOD = 0x1
	FILE_END     SET_FILE_POINTER_MOVE_METHOD = 0x2
)

type MOVE_FILE_FLAGS uint32

const (
	MOVEFILE_COPY_ALLOWED          MOVE_FILE_FLAGS = 0x2
	MOVEFILE_CREATE_HARDLINK       MOVE_FILE_FLAGS = 0x10
	MOVEFILE_DELAY_UNTIL_REBOOT    MOVE_FILE_FLAGS = 0x4
	MOVEFILE_REPLACE_EXISTING      MOVE_FILE_FLAGS = 0x1
	MOVEFILE_WRITE_THROUGH         MOVE_FILE_FLAGS = 0x8
	MOVEFILE_FAIL_IF_NOT_TRACKABLE MOVE_FILE_FLAGS = 0x20
)

type GETFINALPATHNAMEBYHANDLE_FLAGS uint32

const (
	VOLUME_NAME_DOS      GETFINALPATHNAMEBYHANDLE_FLAGS = 0x0
	VOLUME_NAME_GUID     GETFINALPATHNAMEBYHANDLE_FLAGS = 0x1
	VOLUME_NAME_NT       GETFINALPATHNAMEBYHANDLE_FLAGS = 0x2
	VOLUME_NAME_NONE     GETFINALPATHNAMEBYHANDLE_FLAGS = 0x4
	FILE_NAME_NORMALIZED GETFINALPATHNAMEBYHANDLE_FLAGS = 0x0
	FILE_NAME_OPENED     GETFINALPATHNAMEBYHANDLE_FLAGS = 0x8
)

type LZOPENFILE_STYLE uint16

const (
	OF_CANCEL           LZOPENFILE_STYLE = 0x800
	OF_CREATE           LZOPENFILE_STYLE = 0x1000
	OF_DELETE           LZOPENFILE_STYLE = 0x200
	OF_EXIST            LZOPENFILE_STYLE = 0x4000
	OF_PARSE            LZOPENFILE_STYLE = 0x100
	OF_PROMPT           LZOPENFILE_STYLE = 0x2000
	OF_READ             LZOPENFILE_STYLE = 0x0
	OF_READWRITE        LZOPENFILE_STYLE = 0x2
	OF_REOPEN           LZOPENFILE_STYLE = 0x8000
	OF_SHARE_DENY_NONE  LZOPENFILE_STYLE = 0x40
	OF_SHARE_DENY_READ  LZOPENFILE_STYLE = 0x30
	OF_SHARE_DENY_WRITE LZOPENFILE_STYLE = 0x20
	OF_SHARE_EXCLUSIVE  LZOPENFILE_STYLE = 0x10
	OF_WRITE            LZOPENFILE_STYLE = 0x1
	OF_SHARE_COMPAT     LZOPENFILE_STYLE = 0x0
	OF_VERIFY           LZOPENFILE_STYLE = 0x400
)

type FILE_NOTIFY_CHANGE uint32

const (
	FILE_NOTIFY_CHANGE_FILE_NAME   FILE_NOTIFY_CHANGE = 0x1
	FILE_NOTIFY_CHANGE_DIR_NAME    FILE_NOTIFY_CHANGE = 0x2
	FILE_NOTIFY_CHANGE_ATTRIBUTES  FILE_NOTIFY_CHANGE = 0x4
	FILE_NOTIFY_CHANGE_SIZE        FILE_NOTIFY_CHANGE = 0x8
	FILE_NOTIFY_CHANGE_LAST_WRITE  FILE_NOTIFY_CHANGE = 0x10
	FILE_NOTIFY_CHANGE_LAST_ACCESS FILE_NOTIFY_CHANGE = 0x20
	FILE_NOTIFY_CHANGE_CREATION    FILE_NOTIFY_CHANGE = 0x40
	FILE_NOTIFY_CHANGE_SECURITY    FILE_NOTIFY_CHANGE = 0x100
)

type TXFS_MINIVERSION uint32

const (
	TXFS_MINIVERSION_COMMITTED_VIEW TXFS_MINIVERSION = 0x0
	TXFS_MINIVERSION_DIRTY_VIEW     TXFS_MINIVERSION = 0xffff
	TXFS_MINIVERSION_DEFAULT_VIEW   TXFS_MINIVERSION = 0xfffe
)

type TAPE_POSITION_TYPE uint32

const (
	TAPE_ABSOLUTE_POSITION TAPE_POSITION_TYPE = 0x0
	TAPE_LOGICAL_POSITION  TAPE_POSITION_TYPE = 0x1
)

type CREATE_TAPE_PARTITION_METHOD uint32

const (
	TAPE_FIXED_PARTITIONS     CREATE_TAPE_PARTITION_METHOD = 0x0
	TAPE_INITIATOR_PARTITIONS CREATE_TAPE_PARTITION_METHOD = 0x2
	TAPE_SELECT_PARTITIONS    CREATE_TAPE_PARTITION_METHOD = 0x1
)

type REPLACE_FILE_FLAGS uint32

const (
	REPLACEFILE_WRITE_THROUGH       REPLACE_FILE_FLAGS = 0x1
	REPLACEFILE_IGNORE_MERGE_ERRORS REPLACE_FILE_FLAGS = 0x2
	REPLACEFILE_IGNORE_ACL_ERRORS   REPLACE_FILE_FLAGS = 0x4
)

type TAPEMARK_TYPE uint32

const (
	TAPE_FILEMARKS       TAPEMARK_TYPE = 0x1
	TAPE_LONG_FILEMARKS  TAPEMARK_TYPE = 0x3
	TAPE_SETMARKS        TAPEMARK_TYPE = 0x0
	TAPE_SHORT_FILEMARKS TAPEMARK_TYPE = 0x2
)

type TAPE_POSITION_METHOD uint32

const (
	TAPE_ABSOLUTE_BLOCK        TAPE_POSITION_METHOD = 0x1
	TAPE_LOGICAL_BLOCK         TAPE_POSITION_METHOD = 0x2
	TAPE_REWIND                TAPE_POSITION_METHOD = 0x0
	TAPE_SPACE_END_OF_DATA     TAPE_POSITION_METHOD = 0x4
	TAPE_SPACE_FILEMARKS       TAPE_POSITION_METHOD = 0x6
	TAPE_SPACE_RELATIVE_BLOCKS TAPE_POSITION_METHOD = 0x5
	TAPE_SPACE_SEQUENTIAL_FMKS TAPE_POSITION_METHOD = 0x7
	TAPE_SPACE_SEQUENTIAL_SMKS TAPE_POSITION_METHOD = 0x9
	TAPE_SPACE_SETMARKS        TAPE_POSITION_METHOD = 0x8
)

type TAPE_INFORMATION_TYPE uint32

const (
	SET_TAPE_DRIVE_INFORMATION TAPE_INFORMATION_TYPE = 0x1
	SET_TAPE_MEDIA_INFORMATION TAPE_INFORMATION_TYPE = 0x0
)

type LOCK_FILE_FLAGS uint32

const (
	LOCKFILE_EXCLUSIVE_LOCK   LOCK_FILE_FLAGS = 0x2
	LOCKFILE_FAIL_IMMEDIATELY LOCK_FILE_FLAGS = 0x1
)

type PREPARE_TAPE_OPERATION uint32

const (
	TAPE_FORMAT  PREPARE_TAPE_OPERATION = 0x5
	TAPE_LOAD    PREPARE_TAPE_OPERATION = 0x0
	TAPE_LOCK    PREPARE_TAPE_OPERATION = 0x3
	TAPE_TENSION PREPARE_TAPE_OPERATION = 0x2
	TAPE_UNLOAD  PREPARE_TAPE_OPERATION = 0x1
	TAPE_UNLOCK  PREPARE_TAPE_OPERATION = 0x4
)

type GET_TAPE_DRIVE_PARAMETERS_OPERATION uint32

const (
	GET_TAPE_DRIVE_INFORMATION GET_TAPE_DRIVE_PARAMETERS_OPERATION = 0x1
	GET_TAPE_MEDIA_INFORMATION GET_TAPE_DRIVE_PARAMETERS_OPERATION = 0x0
)

type ERASE_TAPE_TYPE uint32

const (
	TAPE_ERASE_LONG  ERASE_TAPE_TYPE = 0x1
	TAPE_ERASE_SHORT ERASE_TAPE_TYPE = 0x0
)

type SYMBOLIC_LINK_FLAGS uint32

const (
	SYMBOLIC_LINK_FLAG_DIRECTORY                 SYMBOLIC_LINK_FLAGS = 0x1
	SYMBOLIC_LINK_FLAG_ALLOW_UNPRIVILEGED_CREATE SYMBOLIC_LINK_FLAGS = 0x2
)

type FILE_TYPE uint32

const (
	FILE_TYPE_UNKNOWN FILE_TYPE = 0x0
	FILE_TYPE_DISK    FILE_TYPE = 0x1
	FILE_TYPE_CHAR    FILE_TYPE = 0x2
	FILE_TYPE_PIPE    FILE_TYPE = 0x3
	FILE_TYPE_REMOTE  FILE_TYPE = 0x8000
)

type COPYFILE_FLAGS uint32

const (
	COPY_FILE_FAIL_IF_EXISTS              COPYFILE_FLAGS = 0x1
	COPY_FILE_RESTARTABLE                 COPYFILE_FLAGS = 0x2
	COPY_FILE_OPEN_SOURCE_FOR_WRITE       COPYFILE_FLAGS = 0x4
	COPY_FILE_ALLOW_DECRYPTED_DESTINATION COPYFILE_FLAGS = 0x8
	COPY_FILE_COPY_SYMLINK                COPYFILE_FLAGS = 0x800
	COPY_FILE_NO_BUFFERING                COPYFILE_FLAGS = 0x1000
	COPY_FILE_REQUEST_SECURITY_PRIVILEGES COPYFILE_FLAGS = 0x2000
	COPY_FILE_RESUME_FROM_PAUSE           COPYFILE_FLAGS = 0x4000
	COPY_FILE_NO_OFFLOAD                  COPYFILE_FLAGS = 0x40000
	COPY_FILE_IGNORE_EDP_BLOCK            COPYFILE_FLAGS = 0x400000
	COPY_FILE_IGNORE_SOURCE_ENCRYPTION    COPYFILE_FLAGS = 0x800000
	COPY_FILE_DONT_REQUEST_DEST_WRITE_DAC COPYFILE_FLAGS = 0x2000000
	COPY_FILE_REQUEST_COMPRESSED_TRAFFIC  COPYFILE_FLAGS = 0x10000000
	COPY_FILE_OPEN_AND_COPY_REPARSE_POINT COPYFILE_FLAGS = 0x200000
	COPY_FILE_DIRECTORY                   COPYFILE_FLAGS = 0x80
	COPY_FILE_SKIP_ALTERNATE_STREAMS      COPYFILE_FLAGS = 0x8000
	COPY_FILE_DISABLE_PRE_ALLOCATION      COPYFILE_FLAGS = 0x4000000
	COPY_FILE_ENABLE_LOW_FREE_SPACE_MODE  COPYFILE_FLAGS = 0x8000000
	COPY_FILE_ENABLE_SPARSE_COPY          COPYFILE_FLAGS = 0x20000000
	COPY_FILE_DISABLE_SPARSE_COPY         COPYFILE_FLAGS = 0x80000000
)

type SYM_LOAD_FLAGS uint32

const (
	SLMFLAG_NONE       SYM_LOAD_FLAGS = 0x0
	SLMFLAG_VIRTUAL    SYM_LOAD_FLAGS = 0x1
	SLMFLAG_ALT_INDEX  SYM_LOAD_FLAGS = 0x2
	SLMFLAG_NO_SYMBOLS SYM_LOAD_FLAGS = 0x4
)

type THREAD_ERROR_MODE uint32

const (
	SEM_ALL_ERRORS             THREAD_ERROR_MODE = 0x0
	SEM_FAILCRITICALERRORS     THREAD_ERROR_MODE = 0x1
	SEM_NOGPFAULTERRORBOX      THREAD_ERROR_MODE = 0x2
	SEM_NOOPENFILEERRORBOX     THREAD_ERROR_MODE = 0x8000
	SEM_NOALIGNMENTFAULTEXCEPT THREAD_ERROR_MODE = 0x4
)

type FORMAT_MESSAGE_OPTIONS uint32

const (
	FORMAT_MESSAGE_ALLOCATE_BUFFER FORMAT_MESSAGE_OPTIONS = 0x100
	FORMAT_MESSAGE_ARGUMENT_ARRAY  FORMAT_MESSAGE_OPTIONS = 0x2000
	FORMAT_MESSAGE_FROM_HMODULE    FORMAT_MESSAGE_OPTIONS = 0x800
	FORMAT_MESSAGE_FROM_STRING     FORMAT_MESSAGE_OPTIONS = 0x400
	FORMAT_MESSAGE_FROM_SYSTEM     FORMAT_MESSAGE_OPTIONS = 0x1000
	FORMAT_MESSAGE_IGNORE_INSERTS  FORMAT_MESSAGE_OPTIONS = 0x200
)

type RTL_VIRTUAL_UNWIND_HANDLER_TYPE uint32

const (
	UNW_FLAG_NHANDLER  RTL_VIRTUAL_UNWIND_HANDLER_TYPE = 0x0
	UNW_FLAG_EHANDLER  RTL_VIRTUAL_UNWIND_HANDLER_TYPE = 0x1
	UNW_FLAG_UHANDLER  RTL_VIRTUAL_UNWIND_HANDLER_TYPE = 0x2
	UNW_FLAG_CHAININFO RTL_VIRTUAL_UNWIND_HANDLER_TYPE = 0x4
)

type OPEN_THREAD_WAIT_CHAIN_SESSION_FLAGS uint32

const (
	WCT_ASYNC_OPEN_FLAG OPEN_THREAD_WAIT_CHAIN_SESSION_FLAGS = 0x1
)

type SYM_SRV_STORE_FILE_FLAGS uint32

const (
	SYMSTOREOPT_COMPRESS       SYM_SRV_STORE_FILE_FLAGS = 0x1
	SYMSTOREOPT_OVERWRITE      SYM_SRV_STORE_FILE_FLAGS = 0x2
	SYMSTOREOPT_PASS_IF_EXISTS SYM_SRV_STORE_FILE_FLAGS = 0x40
	SYMSTOREOPT_POINTER        SYM_SRV_STORE_FILE_FLAGS = 0x8
	SYMSTOREOPT_RETURNINDEX    SYM_SRV_STORE_FILE_FLAGS = 0x4
)

type IMAGE_DIRECTORY_ENTRY uint16

const (
	IMAGE_DIRECTORY_ENTRY_ARCHITECTURE   IMAGE_DIRECTORY_ENTRY = 0x7
	IMAGE_DIRECTORY_ENTRY_BASERELOC      IMAGE_DIRECTORY_ENTRY = 0x5
	IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT   IMAGE_DIRECTORY_ENTRY = 0xb
	IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR IMAGE_DIRECTORY_ENTRY = 0xe
	IMAGE_DIRECTORY_ENTRY_DEBUG          IMAGE_DIRECTORY_ENTRY = 0x6
	IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT   IMAGE_DIRECTORY_ENTRY = 0xd
	IMAGE_DIRECTORY_ENTRY_EXCEPTION      IMAGE_DIRECTORY_ENTRY = 0x3
	IMAGE_DIRECTORY_ENTRY_EXPORT         IMAGE_DIRECTORY_ENTRY = 0x0
	IMAGE_DIRECTORY_ENTRY_GLOBALPTR      IMAGE_DIRECTORY_ENTRY = 0x8
	IMAGE_DIRECTORY_ENTRY_IAT            IMAGE_DIRECTORY_ENTRY = 0xc
	IMAGE_DIRECTORY_ENTRY_IMPORT         IMAGE_DIRECTORY_ENTRY = 0x1
	IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG    IMAGE_DIRECTORY_ENTRY = 0xa
	IMAGE_DIRECTORY_ENTRY_RESOURCE       IMAGE_DIRECTORY_ENTRY = 0x2
	IMAGE_DIRECTORY_ENTRY_SECURITY       IMAGE_DIRECTORY_ENTRY = 0x4
	IMAGE_DIRECTORY_ENTRY_TLS            IMAGE_DIRECTORY_ENTRY = 0x9
)

type WAIT_CHAIN_THREAD_OPTIONS uint32

const (
	WCT_OUT_OF_PROC_COM_FLAG WAIT_CHAIN_THREAD_OPTIONS = 0x2
	WCT_OUT_OF_PROC_CS_FLAG  WAIT_CHAIN_THREAD_OPTIONS = 0x4
	WCT_OUT_OF_PROC_FLAG     WAIT_CHAIN_THREAD_OPTIONS = 0x1
)

type SYM_FIND_ID_OPTION uint32

const (
	SSRVOPT_DWORD    SYM_FIND_ID_OPTION = 0x2
	SSRVOPT_DWORDPTR SYM_FIND_ID_OPTION = 0x4
	SSRVOPT_GUIDPTR  SYM_FIND_ID_OPTION = 0x8
)

type CONTEXT_FLAGS uint32

const (
	CONTEXT_AMD64                     CONTEXT_FLAGS = 0x100000
	CONTEXT_CONTROL_AMD64             CONTEXT_FLAGS = 0x100001
	CONTEXT_INTEGER_AMD64             CONTEXT_FLAGS = 0x100002
	CONTEXT_SEGMENTS_AMD64            CONTEXT_FLAGS = 0x100004
	CONTEXT_FLOATING_POINT_AMD64      CONTEXT_FLAGS = 0x100008
	CONTEXT_DEBUG_REGISTERS_AMD64     CONTEXT_FLAGS = 0x100010
	CONTEXT_FULL_AMD64                CONTEXT_FLAGS = 0x10000b
	CONTEXT_ALL_AMD64                 CONTEXT_FLAGS = 0x10001f
	CONTEXT_XSTATE_AMD64              CONTEXT_FLAGS = 0x100040
	CONTEXT_KERNEL_CET_AMD64          CONTEXT_FLAGS = 0x100080
	CONTEXT_KERNEL_DEBUGGER_AMD64     CONTEXT_FLAGS = 0x4000000
	CONTEXT_EXCEPTION_ACTIVE_AMD64    CONTEXT_FLAGS = 0x8000000
	CONTEXT_SERVICE_ACTIVE_AMD64      CONTEXT_FLAGS = 0x10000000
	CONTEXT_EXCEPTION_REQUEST_AMD64   CONTEXT_FLAGS = 0x40000000
	CONTEXT_EXCEPTION_REPORTING_AMD64 CONTEXT_FLAGS = 0x80000000
	CONTEXT_UNWOUND_TO_CALL_AMD64     CONTEXT_FLAGS = 0x20000000
	CONTEXT_X86                       CONTEXT_FLAGS = 0x10000
	CONTEXT_CONTROL_X86               CONTEXT_FLAGS = 0x10001
	CONTEXT_INTEGER_X86               CONTEXT_FLAGS = 0x10002
	CONTEXT_SEGMENTS_X86              CONTEXT_FLAGS = 0x10004
	CONTEXT_FLOATING_POINT_X86        CONTEXT_FLAGS = 0x10008
	CONTEXT_DEBUG_REGISTERS_X86       CONTEXT_FLAGS = 0x10010
	CONTEXT_EXTENDED_REGISTERS_X86    CONTEXT_FLAGS = 0x10020
	CONTEXT_FULL_X86                  CONTEXT_FLAGS = 0x10007
	CONTEXT_ALL_X86                   CONTEXT_FLAGS = 0x1003f
	CONTEXT_XSTATE_X86                CONTEXT_FLAGS = 0x10040
	CONTEXT_EXCEPTION_ACTIVE_X86      CONTEXT_FLAGS = 0x8000000
	CONTEXT_SERVICE_ACTIVE_X86        CONTEXT_FLAGS = 0x10000000
	CONTEXT_EXCEPTION_REQUEST_X86     CONTEXT_FLAGS = 0x40000000
	CONTEXT_EXCEPTION_REPORTING_X86   CONTEXT_FLAGS = 0x80000000
	CONTEXT_ARM64                     CONTEXT_FLAGS = 0x400000
	CONTEXT_CONTROL_ARM64             CONTEXT_FLAGS = 0x400001
	CONTEXT_INTEGER_ARM64             CONTEXT_FLAGS = 0x400002
	CONTEXT_FLOATING_POINT_ARM64      CONTEXT_FLAGS = 0x400004
	CONTEXT_DEBUG_REGISTERS_ARM64     CONTEXT_FLAGS = 0x400008
	CONTEXT_X18_ARM64                 CONTEXT_FLAGS = 0x400010
	CONTEXT_FULL_ARM64                CONTEXT_FLAGS = 0x400007
	CONTEXT_ALL_ARM64                 CONTEXT_FLAGS = 0x40001f
	CONTEXT_EXCEPTION_ACTIVE_ARM64    CONTEXT_FLAGS = 0x8000000
	CONTEXT_SERVICE_ACTIVE_ARM64      CONTEXT_FLAGS = 0x10000000
	CONTEXT_EXCEPTION_REQUEST_ARM64   CONTEXT_FLAGS = 0x40000000
	CONTEXT_EXCEPTION_REPORTING_ARM64 CONTEXT_FLAGS = 0x80000000
	CONTEXT_UNWOUND_TO_CALL_ARM64     CONTEXT_FLAGS = 0x20000000
	CONTEXT_RET_TO_GUEST_ARM64        CONTEXT_FLAGS = 0x40000000
	CONTEXT_ARM                       CONTEXT_FLAGS = 0x200000
	CONTEXT_CONTROL_ARM               CONTEXT_FLAGS = 0x200001
	CONTEXT_INTEGER_ARM               CONTEXT_FLAGS = 0x200002
	CONTEXT_FLOATING_POINT_ARM        CONTEXT_FLAGS = 0x200004
	CONTEXT_DEBUG_REGISTERS_ARM       CONTEXT_FLAGS = 0x200008
	CONTEXT_FULL_ARM                  CONTEXT_FLAGS = 0x200007
	CONTEXT_ALL_ARM                   CONTEXT_FLAGS = 0x20000f
	CONTEXT_EXCEPTION_ACTIVE_ARM      CONTEXT_FLAGS = 0x8000000
	CONTEXT_SERVICE_ACTIVE_ARM        CONTEXT_FLAGS = 0x10000000
	CONTEXT_EXCEPTION_REQUEST_ARM     CONTEXT_FLAGS = 0x40000000
	CONTEXT_EXCEPTION_REPORTING_ARM   CONTEXT_FLAGS = 0x80000000
	CONTEXT_UNWOUND_TO_CALL_ARM       CONTEXT_FLAGS = 0x20000000
)

type MESSAGEBOX_STYLE uint32

const (
	MB_ABORTRETRYIGNORE          MESSAGEBOX_STYLE = 0x2
	MB_CANCELTRYCONTINUE         MESSAGEBOX_STYLE = 0x6
	MB_HELP                      MESSAGEBOX_STYLE = 0x4000
	MB_OK                        MESSAGEBOX_STYLE = 0x0
	MB_OKCANCEL                  MESSAGEBOX_STYLE = 0x1
	MB_RETRYCANCEL               MESSAGEBOX_STYLE = 0x5
	MB_YESNO                     MESSAGEBOX_STYLE = 0x4
	MB_YESNOCANCEL               MESSAGEBOX_STYLE = 0x3
	MB_ICONHAND                  MESSAGEBOX_STYLE = 0x10
	MB_ICONQUESTION              MESSAGEBOX_STYLE = 0x20
	MB_ICONEXCLAMATION           MESSAGEBOX_STYLE = 0x30
	MB_ICONASTERISK              MESSAGEBOX_STYLE = 0x40
	MB_USERICON                  MESSAGEBOX_STYLE = 0x80
	MB_ICONWARNING               MESSAGEBOX_STYLE = 0x30
	MB_ICONERROR                 MESSAGEBOX_STYLE = 0x10
	MB_ICONINFORMATION           MESSAGEBOX_STYLE = 0x40
	MB_ICONSTOP                  MESSAGEBOX_STYLE = 0x10
	MB_DEFBUTTON1                MESSAGEBOX_STYLE = 0x0
	MB_DEFBUTTON2                MESSAGEBOX_STYLE = 0x100
	MB_DEFBUTTON3                MESSAGEBOX_STYLE = 0x200
	MB_DEFBUTTON4                MESSAGEBOX_STYLE = 0x300
	MB_APPLMODAL                 MESSAGEBOX_STYLE = 0x0
	MB_SYSTEMMODAL               MESSAGEBOX_STYLE = 0x1000
	MB_TASKMODAL                 MESSAGEBOX_STYLE = 0x2000
	MB_NOFOCUS                   MESSAGEBOX_STYLE = 0x8000
	MB_SETFOREGROUND             MESSAGEBOX_STYLE = 0x10000
	MB_DEFAULT_DESKTOP_ONLY      MESSAGEBOX_STYLE = 0x20000
	MB_TOPMOST                   MESSAGEBOX_STYLE = 0x40000
	MB_RIGHT                     MESSAGEBOX_STYLE = 0x80000
	MB_RTLREADING                MESSAGEBOX_STYLE = 0x100000
	MB_SERVICE_NOTIFICATION      MESSAGEBOX_STYLE = 0x200000
	MB_SERVICE_NOTIFICATION_NT3X MESSAGEBOX_STYLE = 0x40000
	MB_TYPEMASK                  MESSAGEBOX_STYLE = 0xf
	MB_ICONMASK                  MESSAGEBOX_STYLE = 0xf0
	MB_DEFMASK                   MESSAGEBOX_STYLE = 0xf00
	MB_MODEMASK                  MESSAGEBOX_STYLE = 0x3000
	MB_MISCMASK                  MESSAGEBOX_STYLE = 0xc000
)

type HANDLE unsafe.Pointer

type HCRYPTASYNC uintptr

type HCERTCHAINENGINE unsafe.Pointer

type BCRYPT_ALG_HANDLE unsafe.Pointer

type BCRYPT_KEY_HANDLE unsafe.Pointer

type BCRYPT_HASH_HANDLE unsafe.Pointer

type BCRYPT_SECRET_HANDLE unsafe.Pointer

type NCRYPT_PROV_HANDLE uintptr

type NCRYPT_KEY_HANDLE uintptr

type HIORING unsafe.Pointer

type BOOL int32

type BOOLEAN uint8

type HMODULE unsafe.Pointer

type HRESULT int32

type HWND unsafe.Pointer

type NTSTATUS int32

type PSTRElement uint8

type PWSTRElement uint16

type PSID unsafe.Pointer

type NCRYPT_DESCRIPTOR_HANDLE unsafe.Pointer

type NCRYPT_STREAM_HANDLE unsafe.Pointer

type BCRYPT_HANDLE unsafe.Pointer

type NCRYPT_HANDLE uintptr

type NCRYPT_HASH_HANDLE uintptr

type NCRYPT_SECRET_HANDLE uintptr

type HCRYPTPROV_LEGACY uintptr

type HCRYPTPROV_OR_NCRYPT_KEY_HANDLE uintptr

type HCERTSTORE unsafe.Pointer

type SECURITY_ATTRIBUTES struct {
	NLength              uint32
	LpSecurityDescriptor unsafe.Pointer
	BInheritHandle       BOOL
}

type OVERLAPPED struct {
	Internal     uintptr
	InternalHigh uintptr
	Offset       uint32
	OffsetHigh   uint32
	HEvent       HANDLE
}

type LPOVERLAPPED_COMPLETION_ROUTINE uintptr

type SYSTEMTIME struct {
	WYear         uint16
	WMonth        uint16
	WDayOfWeek    uint16
	WDay          uint16
	WHour         uint16
	WMinute       uint16
	WSecond       uint16
	WMilliseconds uint16
}

type WIN32_FIND_DATAA struct {
	DwFileAttributes   uint32
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	NFileSizeHigh      uint32
	NFileSizeLow       uint32
	DwReserved0        uint32
	DwReserved1        uint32
	CFileName          [260]CHAR
	CAlternateFileName [14]CHAR
}

type WIN32_FIND_DATAW struct {
	DwFileAttributes   uint32
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	NFileSizeHigh      uint32
	NFileSizeLow       uint32
	DwReserved0        uint32
	DwReserved1        uint32
	CFileName          [260]uint16
	CAlternateFileName [14]uint16
}

type FINDEX_INFO_LEVELS int32

const (
	FindExInfoStandard     FINDEX_INFO_LEVELS = 0x0
	FindExInfoBasic        FINDEX_INFO_LEVELS = 0x1
	FindExInfoMaxInfoLevel FINDEX_INFO_LEVELS = 0x2
)

type FINDEX_SEARCH_OPS int32

const (
	FindExSearchNameMatch          FINDEX_SEARCH_OPS = 0x0
	FindExSearchLimitToDirectories FINDEX_SEARCH_OPS = 0x1
	FindExSearchLimitToDevices     FINDEX_SEARCH_OPS = 0x2
	FindExSearchMaxSearchOp        FINDEX_SEARCH_OPS = 0x3
)

type READ_DIRECTORY_NOTIFY_INFORMATION_CLASS int32

const (
	ReadDirectoryNotifyInformation         READ_DIRECTORY_NOTIFY_INFORMATION_CLASS = 0x1
	ReadDirectoryNotifyExtendedInformation READ_DIRECTORY_NOTIFY_INFORMATION_CLASS = 0x2
	ReadDirectoryNotifyFullInformation     READ_DIRECTORY_NOTIFY_INFORMATION_CLASS = 0x3
	ReadDirectoryNotifyMaximumInformation  READ_DIRECTORY_NOTIFY_INFORMATION_CLASS = 0x4
)

type GET_FILEEX_INFO_LEVELS int32

const (
	GetFileExInfoStandard GET_FILEEX_INFO_LEVELS = 0x0
	GetFileExMaxInfoLevel GET_FILEEX_INFO_LEVELS = 0x1
)

type FILE_INFO_BY_HANDLE_CLASS int32

const (
	FileBasicInfo                  FILE_INFO_BY_HANDLE_CLASS = 0x0
	FileStandardInfo               FILE_INFO_BY_HANDLE_CLASS = 0x1
	FileNameInfo                   FILE_INFO_BY_HANDLE_CLASS = 0x2
	FileRenameInfo                 FILE_INFO_BY_HANDLE_CLASS = 0x3
	FileDispositionInfo            FILE_INFO_BY_HANDLE_CLASS = 0x4
	FileAllocationInfo             FILE_INFO_BY_HANDLE_CLASS = 0x5
	FileEndOfFileInfo              FILE_INFO_BY_HANDLE_CLASS = 0x6
	FileStreamInfo                 FILE_INFO_BY_HANDLE_CLASS = 0x7
	FileCompressionInfo            FILE_INFO_BY_HANDLE_CLASS = 0x8
	FileAttributeTagInfo           FILE_INFO_BY_HANDLE_CLASS = 0x9
	FileIdBothDirectoryInfo        FILE_INFO_BY_HANDLE_CLASS = 0xa
	FileIdBothDirectoryRestartInfo FILE_INFO_BY_HANDLE_CLASS = 0xb
	FileIoPriorityHintInfo         FILE_INFO_BY_HANDLE_CLASS = 0xc
	FileRemoteProtocolInfo         FILE_INFO_BY_HANDLE_CLASS = 0xd
	FileFullDirectoryInfo          FILE_INFO_BY_HANDLE_CLASS = 0xe
	FileFullDirectoryRestartInfo   FILE_INFO_BY_HANDLE_CLASS = 0xf
	FileStorageInfo                FILE_INFO_BY_HANDLE_CLASS = 0x10
	FileAlignmentInfo              FILE_INFO_BY_HANDLE_CLASS = 0x11
	FileIdInfo                     FILE_INFO_BY_HANDLE_CLASS = 0x12
	FileIdExtdDirectoryInfo        FILE_INFO_BY_HANDLE_CLASS = 0x13
	FileIdExtdDirectoryRestartInfo FILE_INFO_BY_HANDLE_CLASS = 0x14
	FileDispositionInfoEx          FILE_INFO_BY_HANDLE_CLASS = 0x15
	FileRenameInfoEx               FILE_INFO_BY_HANDLE_CLASS = 0x16
	FileCaseSensitiveInfo          FILE_INFO_BY_HANDLE_CLASS = 0x17
	FileNormalizedNameInfo         FILE_INFO_BY_HANDLE_CLASS = 0x18
	MaximumFileInfoByHandleClass   FILE_INFO_BY_HANDLE_CLASS = 0x19
)

type FILE_INFO_BY_NAME_CLASS int32

const (
	FileStatByNameInfo          FILE_INFO_BY_NAME_CLASS = 0x0
	FileStatLxByNameInfo        FILE_INFO_BY_NAME_CLASS = 0x1
	FileCaseSensitiveByNameInfo FILE_INFO_BY_NAME_CLASS = 0x2
	FileStatBasicByNameInfo     FILE_INFO_BY_NAME_CLASS = 0x3
	MaximumFileInfoByNameClass  FILE_INFO_BY_NAME_CLASS = 0x4
)

type DEBUG_EVENT struct {
	DwDebugEventCode DEBUG_EVENT_CODE
	DwProcessId      uint32
	DwThreadId       uint32
	Exception        EXCEPTION_DEBUG_INFO
}

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

type TRANSACTION_NOTIFICATION struct {
	TransactionKey          unsafe.Pointer
	TransactionNotification uint32
	TmVirtualClock          int64
	ArgumentLength          uint32
}

type DISK_SPACE_INFORMATION struct {
	ActualTotalAllocationUnits           uint64
	ActualAvailableAllocationUnits       uint64
	ActualPoolUnavailableAllocationUnits uint64
	CallerTotalAllocationUnits           uint64
	CallerAvailableAllocationUnits       uint64
	CallerPoolUnavailableAllocationUnits uint64
	UsedAllocationUnits                  uint64
	TotalReservedAllocationUnits         uint64
	VolumeStorageReserveAllocationUnits  uint64
	AvailableCommittedAllocationUnits    uint64
	PoolAvailableAllocationUnits         uint64
	SectorsPerAllocationUnit             uint32
	BytesPerSector                       uint32
}

type BY_HANDLE_FILE_INFORMATION struct {
	DwFileAttributes     uint32
	FtCreationTime       FILETIME
	FtLastAccessTime     FILETIME
	FtLastWriteTime      FILETIME
	DwVolumeSerialNumber uint32
	NFileSizeHigh        uint32
	NFileSizeLow         uint32
	NNumberOfLinks       uint32
	NFileIndexHigh       uint32
	NFileIndexLow        uint32
}

type CREATEFILE2_EXTENDED_PARAMETERS struct {
	DwSize               uint32
	DwFileAttributes     uint32
	DwFileFlags          uint32
	DwSecurityQosFlags   uint32
	LpSecurityAttributes *SECURITY_ATTRIBUTES
	HTemplateFile        HANDLE
}

type STREAM_INFO_LEVELS int32

const (
	FindStreamInfoStandard     STREAM_INFO_LEVELS = 0x0
	FindStreamInfoMaxInfoLevel STREAM_INFO_LEVELS = 0x1
)

type CREATEFILE3_EXTENDED_PARAMETERS struct {
	DwSize               uint32
	DwFileAttributes     uint32
	DwFileFlags          uint32
	DwSecurityQosFlags   uint32
	LpSecurityAttributes *SECURITY_ATTRIBUTES
	HTemplateFile        HANDLE
}

type DIRECTORY_FLAGS int32

const (
	DIRECTORY_FLAGS_NONE                    DIRECTORY_FLAGS = 0x0
	DIRECTORY_FLAGS_DISALLOW_PATH_REDIRECTS DIRECTORY_FLAGS = 0x1
)

type CLS_LSN struct {
	Internal uint64
}

type CLFS_CONTEXT_MODE int32

const (
	ClfsContextNone     CLFS_CONTEXT_MODE = 0x0
	ClfsContextUndoNext CLFS_CONTEXT_MODE = 0x1
	ClfsContextPrevious CLFS_CONTEXT_MODE = 0x2
	ClfsContextForward  CLFS_CONTEXT_MODE = 0x3
)

type CLS_WRITE_ENTRY struct {
	Buffer     unsafe.Pointer
	ByteLength uint32
}

type CLS_INFORMATION struct {
	TotalAvailable    int64
	CurrentAvailable  int64
	TotalReservation  int64
	BaseFileSize      uint64
	ContainerSize     uint64
	TotalContainers   uint32
	FreeContainers    uint32
	TotalClients      uint32
	Attributes        uint32
	FlushThreshold    uint32
	SectorSize        uint32
	MinArchiveTailLsn CLS_LSN
	BaseLsn           CLS_LSN
	LastFlushedLsn    CLS_LSN
	LastLsn           CLS_LSN
	RestartLsn        CLS_LSN
	Identity          Guid
}

type CLFS_IOSTATS_CLASS int32

const (
	ClfsIoStatsDefault CLFS_IOSTATS_CLASS = 0x0
	ClfsIoStatsMax     CLFS_IOSTATS_CLASS = 0xffff
)

type CLS_SCAN_CONTEXT struct {
	CidNode             CLFS_NODE_ID
	HLog                HANDLE
	CIndex              uint32
	CContainers         uint32
	CContainersReturned uint32
	EScanMode           uint8
	PinfoContainer      *CLS_CONTAINER_INFORMATION
}

type CLS_ARCHIVE_DESCRIPTOR struct {
	CoffLow       uint64
	CoffHigh      uint64
	InfoContainer CLS_CONTAINER_INFORMATION
}

type CLFS_BLOCK_ALLOCATION uintptr

type CLFS_BLOCK_DEALLOCATION uintptr

type CLFS_LOG_ARCHIVE_MODE int32

const (
	ClfsLogArchiveEnabled  CLFS_LOG_ARCHIVE_MODE = 0x1
	ClfsLogArchiveDisabled CLFS_LOG_ARCHIVE_MODE = 0x2
)

type CLFS_MGMT_POLICY_TYPE int32

const (
	ClfsMgmtPolicyMaximumSize           CLFS_MGMT_POLICY_TYPE = 0x0
	ClfsMgmtPolicyMinimumSize           CLFS_MGMT_POLICY_TYPE = 0x1
	ClfsMgmtPolicyNewContainerSize      CLFS_MGMT_POLICY_TYPE = 0x2
	ClfsMgmtPolicyGrowthRate            CLFS_MGMT_POLICY_TYPE = 0x3
	ClfsMgmtPolicyLogTail               CLFS_MGMT_POLICY_TYPE = 0x4
	ClfsMgmtPolicyAutoShrink            CLFS_MGMT_POLICY_TYPE = 0x5
	ClfsMgmtPolicyAutoGrow              CLFS_MGMT_POLICY_TYPE = 0x6
	ClfsMgmtPolicyNewContainerPrefix    CLFS_MGMT_POLICY_TYPE = 0x7
	ClfsMgmtPolicyNewContainerSuffix    CLFS_MGMT_POLICY_TYPE = 0x8
	ClfsMgmtPolicyNewContainerExtension CLFS_MGMT_POLICY_TYPE = 0x9
	ClfsMgmtPolicyInvalid               CLFS_MGMT_POLICY_TYPE = 0xa
)

type CLFS_MGMT_POLICY struct {
	Version       uint32
	LengthInBytes uint32
	PolicyFlags   uint32
	PolicyType    CLFS_MGMT_POLICY_TYPE
	Containers    uint32
}

type CLFS_MGMT_NOTIFICATION struct {
	Notification CLFS_MGMT_NOTIFICATION_TYPE
	Lsn          CLS_LSN
	LogIsPinned  uint16
}

type LOG_MANAGEMENT_CALLBACKS struct {
	CallbackContext        unsafe.Pointer
	AdvanceTailCallback    PLOG_TAIL_ADVANCE_CALLBACK
	LogFullHandlerCallback PLOG_FULL_HANDLER_CALLBACK
	LogUnpinnedCallback    PLOG_UNPINNED_CALLBACK
}

type ENCRYPTION_CERTIFICATE struct {
	CbTotalLength uint32
	PUserSid      *SID
	PCertBlob     *EFS_CERTIFICATE_BLOB
}

type ENCRYPTION_CERTIFICATE_HASH struct {
	CbTotalLength        uint32
	PUserSid             *SID
	PHash                *EFS_HASH_BLOB
	LpDisplayInformation *PWSTRElement
}

type ENCRYPTION_CERTIFICATE_HASH_LIST struct {
	NCert_Hash uint32
	PUsers     **ENCRYPTION_CERTIFICATE_HASH
}

type ENCRYPTION_CERTIFICATE_LIST struct {
	NUsers uint32
	PUsers **ENCRYPTION_CERTIFICATE
}

type WofEnumEntryProc uintptr

type WofEnumFilesProc uintptr

type TXF_ID struct {
	LowPart  int64
	HighPart int64
}

type IORING_VERSION int32

const (
	IORING_VERSION_INVALID IORING_VERSION = 0x0
	IORING_VERSION_1       IORING_VERSION = 0x1
	IORING_VERSION_2       IORING_VERSION = 0x2
	IORING_VERSION_3       IORING_VERSION = 0x12c
	IORING_VERSION_4       IORING_VERSION = 0x190
)

type IORING_OP_CODE int32

const (
	IORING_OP_NOP              IORING_OP_CODE = 0x0
	IORING_OP_READ             IORING_OP_CODE = 0x1
	IORING_OP_REGISTER_FILES   IORING_OP_CODE = 0x2
	IORING_OP_REGISTER_BUFFERS IORING_OP_CODE = 0x3
	IORING_OP_CANCEL           IORING_OP_CODE = 0x4
	IORING_OP_WRITE            IORING_OP_CODE = 0x5
	IORING_OP_FLUSH            IORING_OP_CODE = 0x6
	IORING_OP_READ_SCATTER     IORING_OP_CODE = 0x7
	IORING_OP_WRITE_GATHER     IORING_OP_CODE = 0x8
)

type IORING_BUFFER_INFO struct {
	Address unsafe.Pointer
	Length  uint32
}

type IORING_SQE_FLAGS int32

const (
	IOSQE_FLAGS_NONE                IORING_SQE_FLAGS = 0x0
	IOSQE_FLAGS_DRAIN_PRECEDING_OPS IORING_SQE_FLAGS = 0x1
)

type IORING_CREATE_FLAGS struct {
	Required IORING_CREATE_REQUIRED_FLAGS
	Advisory IORING_CREATE_ADVISORY_FLAGS
}

type IORING_INFO struct {
	IoRingVersion       IORING_VERSION
	Flags               IORING_CREATE_FLAGS
	SubmissionQueueSize uint32
	CompletionQueueSize uint32
}

type IORING_CAPABILITIES struct {
	MaxVersion             IORING_VERSION
	MaxSubmissionQueueSize uint32
	MaxCompletionQueueSize uint32
	FeatureFlags           IORING_FEATURE_FLAGS
}

type IORING_HANDLE_REF struct {
	Kind   IORING_REF_KIND
	Handle HANDLE
}

type IORING_BUFFER_REF struct {
	Kind    IORING_REF_KIND
	Address unsafe.Pointer
}

type IORING_CQE struct {
	UserData    uintptr
	ResultCode  HRESULT
	Information uintptr
}

type CREATE_BIND_LINK_FLAGS int32

const (
	CREATE_BIND_LINK_FLAG_NONE      CREATE_BIND_LINK_FLAGS = 0x0
	CREATE_BIND_LINK_FLAG_READ_ONLY CREATE_BIND_LINK_FLAGS = 0x1
	CREATE_BIND_LINK_FLAG_MERGED    CREATE_BIND_LINK_FLAGS = 0x2
)

type EXCEPTION_ROUTINE uintptr

type PUBLICKEYSTRUC struct {
	BType    uint8
	BVersion uint8
	Reserved uint16
	AiKeyAlg ALG_ID
}

type CRYPT_INTEGER_BLOB struct {
	CbData uint32
	PbData *uint8
}

type BCryptBufferDesc struct {
	UlVersion uint32
	CBuffers  uint32
	PBuffers  *BCryptBuffer
}

type BCRYPT_MULTI_OPERATION_TYPE int32

const (
	BCRYPT_OPERATION_TYPE_HASH BCRYPT_MULTI_OPERATION_TYPE = 0x1
)

type BCRYPT_ALGORITHM_IDENTIFIER struct {
	PszName *PWSTRElement
	DwClass uint32
	DwFlags uint32
}

type BCRYPT_PROVIDER_NAME struct {
	PszProviderName *PWSTRElement
}

type CRYPT_PROVIDER_REG struct {
	CAliases     uint32
	RgpszAliases **PWSTRElement
	PUM          *CRYPT_IMAGE_REG
	PKM          *CRYPT_IMAGE_REG
}

type CRYPT_PROVIDERS struct {
	CProviders     uint32
	RgpszProviders **PWSTRElement
}

type CRYPT_CONTEXT_CONFIG struct {
	DwFlags    CRYPT_CONTEXT_CONFIG_FLAGS
	DwReserved uint32
}

type CRYPT_CONTEXT_FUNCTION_CONFIG struct {
	DwFlags    uint32
	DwReserved uint32
}

type CRYPT_CONTEXTS struct {
	CContexts     uint32
	RgpszContexts **PWSTRElement
}

type CRYPT_CONTEXT_FUNCTIONS struct {
	CFunctions     uint32
	RgpszFunctions **PWSTRElement
}

type CRYPT_CONTEXT_FUNCTION_PROVIDERS struct {
	CProviders     uint32
	RgpszProviders **PWSTRElement
}

type CRYPT_PROVIDER_REFS struct {
	CProviders   uint32
	RgpProviders **CRYPT_PROVIDER_REF
}

type NCRYPT_ALLOC_PARA struct {
	CbSize   uint32
	PfnAlloc PFN_NCRYPT_ALLOC
	PfnFree  PFN_NCRYPT_FREE
}

type NCryptAlgorithmName struct {
	PszName         *PWSTRElement
	DwClass         NCRYPT_ALGORITHM_NAME_CLASS
	DwAlgOperations NCRYPT_OPERATION
	DwFlags         uint32
}

type NCryptKeyName struct {
	PszName         *PWSTRElement
	PszAlgid        *PWSTRElement
	DwLegacyKeySpec CERT_KEY_SPEC
	DwFlags         uint32
}

type NCryptProviderName struct {
	PszName    *PWSTRElement
	PszComment *PWSTRElement
}

type CRYPT_ALGORITHM_IDENTIFIER struct {
	PszObjId   *PSTRElement
	Parameters CRYPT_INTEGER_BLOB
}

type CERT_EXTENSION struct {
	PszObjId  *PSTRElement
	FCritical BOOL
	Value     CRYPT_INTEGER_BLOB
}

type CRYPT_ATTRIBUTE struct {
	PszObjId *PSTRElement
	CValue   uint32
	RgValue  *CRYPT_INTEGER_BLOB
}

type CRYPT_ATTRIBUTES struct {
	CAttr  uint32
	RgAttr *CRYPT_ATTRIBUTE
}

type CERT_RDN_ATTR struct {
	PszObjId    *PSTRElement
	DwValueType uint32
	Value       CRYPT_INTEGER_BLOB
}

type CERT_RDN struct {
	CRDNAttr  uint32
	RgRDNAttr *CERT_RDN_ATTR
}

type CERT_NAME_INFO struct {
	CRDN  uint32
	RgRDN *CERT_RDN
}

type CERT_PUBLIC_KEY_INFO struct {
	Algorithm CRYPT_ALGORITHM_IDENTIFIER
	PublicKey CRYPT_BIT_BLOB
}

type CRYPT_PKCS8_IMPORT_PARAMS struct {
	PrivateKey             CRYPT_INTEGER_BLOB
	PResolvehCryptProvFunc PCRYPT_RESOLVE_HCRYPTPROV_FUNC
	PVoidResolveFunc       unsafe.Pointer
	PDecryptPrivateKeyFunc PCRYPT_DECRYPT_PRIVATE_KEY_FUNC
	PVoidDecryptFunc       unsafe.Pointer
}

type CERT_INFO struct {
	DwVersion            uint32
	SerialNumber         CRYPT_INTEGER_BLOB
	SignatureAlgorithm   CRYPT_ALGORITHM_IDENTIFIER
	Issuer               CRYPT_INTEGER_BLOB
	NotBefore            FILETIME
	NotAfter             FILETIME
	Subject              CRYPT_INTEGER_BLOB
	SubjectPublicKeyInfo CERT_PUBLIC_KEY_INFO
	IssuerUniqueId       CRYPT_BIT_BLOB
	SubjectUniqueId      CRYPT_BIT_BLOB
	CExtension           uint32
	RgExtension          *CERT_EXTENSION
}

type CRL_ENTRY struct {
	SerialNumber   CRYPT_INTEGER_BLOB
	RevocationDate FILETIME
	CExtension     uint32
	RgExtension    *CERT_EXTENSION
}

type CRL_INFO struct {
	DwVersion          uint32
	SignatureAlgorithm CRYPT_ALGORITHM_IDENTIFIER
	Issuer             CRYPT_INTEGER_BLOB
	ThisUpdate         FILETIME
	NextUpdate         FILETIME
	CCRLEntry          uint32
	RgCRLEntry         *CRL_ENTRY
	CExtension         uint32
	RgExtension        *CERT_EXTENSION
}

type CTL_USAGE struct {
	CUsageIdentifier     uint32
	RgpszUsageIdentifier **PSTRElement
}

type CTL_ENTRY struct {
	SubjectIdentifier CRYPT_INTEGER_BLOB
	CAttribute        uint32
	RgAttribute       *CRYPT_ATTRIBUTE
}

type CTL_INFO struct {
	DwVersion        uint32
	SubjectUsage     CTL_USAGE
	ListIdentifier   CRYPT_INTEGER_BLOB
	SequenceNumber   CRYPT_INTEGER_BLOB
	ThisUpdate       FILETIME
	NextUpdate       FILETIME
	SubjectAlgorithm CRYPT_ALGORITHM_IDENTIFIER
	CCTLEntry        uint32
	RgCTLEntry       *CTL_ENTRY
	CExtension       uint32
	RgExtension      *CERT_EXTENSION
}

type CRYPT_ENCODE_PARA struct {
	CbSize   uint32
	PfnAlloc PFN_CRYPT_ALLOC
	PfnFree  PFN_CRYPT_FREE
}

type CRYPT_DECODE_PARA struct {
	CbSize   uint32
	PfnAlloc PFN_CRYPT_ALLOC
	PfnFree  PFN_CRYPT_FREE
}

type CERT_EXTENSIONS struct {
	CExtension  uint32
	RgExtension *CERT_EXTENSION
}

type CRYPT_OID_FUNC_ENTRY struct {
	PszOID     *PSTRElement
	PvFuncAddr unsafe.Pointer
}

type PFN_CRYPT_ENUM_OID_FUNC uintptr

type CRYPT_OID_INFO struct {
	CbSize    uint32
	PszOID    *PSTRElement
	PwszName  *PWSTRElement
	DwGroupId uint32
	DwValue   uint32
	ExtraInfo CRYPT_INTEGER_BLOB
}

type PFN_CRYPT_ENUM_OID_INFO uintptr

type CERT_STRONG_SIGN_PARA struct {
	CbSize       uint32
	DwInfoChoice uint32
	PvInfo       unsafe.Pointer
}

type CMSG_SIGNER_ENCODE_INFO struct {
	CbSize        uint32
	PCertInfo     *CERT_INFO
	HCryptProv    uintptr
	DwKeySpec     uint32
	HashAlgorithm CRYPT_ALGORITHM_IDENTIFIER
	PvHashAuxInfo unsafe.Pointer
	CAuthAttr     uint32
	RgAuthAttr    *CRYPT_ATTRIBUTE
	CUnauthAttr   uint32
	RgUnauthAttr  *CRYPT_ATTRIBUTE
}

type CMSG_SIGNED_ENCODE_INFO struct {
	CbSize        uint32
	CSigners      uint32
	RgSigners     *CMSG_SIGNER_ENCODE_INFO
	CCertEncoded  uint32
	RgCertEncoded *CRYPT_INTEGER_BLOB
	CCrlEncoded   uint32
	RgCrlEncoded  *CRYPT_INTEGER_BLOB
}

type CMSG_STREAM_INFO struct {
	CbContent       uint32
	PfnStreamOutput PFN_CMSG_STREAM_OUTPUT
	PvArg           unsafe.Pointer
}

type CERT_CONTEXT struct {
	DwCertEncodingType CERT_QUERY_ENCODING_TYPE
	PbCertEncoded      *uint8
	CbCertEncoded      uint32
	PCertInfo          *CERT_INFO
	HCertStore         HCERTSTORE
}

type CRL_CONTEXT struct {
	DwCertEncodingType CERT_QUERY_ENCODING_TYPE
	PbCrlEncoded       *uint8
	CbCrlEncoded       uint32
	PCrlInfo           *CRL_INFO
	HCertStore         HCERTSTORE
}

type CTL_CONTEXT struct {
	DwMsgAndCertEncodingType uint32
	PbCtlEncoded             *uint8
	CbCtlEncoded             uint32
	PCtlInfo                 *CTL_INFO
	HCertStore               HCERTSTORE
	HCryptMsg                unsafe.Pointer
	PbCtlContent             *uint8
	CbCtlContent             uint32
}

type CRYPT_KEY_PROV_INFO struct {
	PwszContainerName *PWSTRElement
	PwszProvName      *PWSTRElement
	DwProvType        uint32
	DwFlags           CRYPT_KEY_FLAGS
	CProvParam        uint32
	RgProvParam       *CRYPT_KEY_PROV_PARAM
	DwKeySpec         uint32
}

type CERT_CREATE_CONTEXT_PARA struct {
	CbSize  uint32
	PfnFree PFN_CRYPT_FREE
	PvFree  unsafe.Pointer
	PfnSort PFN_CERT_CREATE_CONTEXT_SORT_FUNC
	PvSort  unsafe.Pointer
}

type CERT_SYSTEM_STORE_INFO struct {
	CbSize uint32
}

type CERT_PHYSICAL_STORE_INFO struct {
	CbSize               uint32
	PszOpenStoreProvider *PSTRElement
	DwOpenEncodingType   uint32
	DwOpenFlags          uint32
	OpenParameters       CRYPT_INTEGER_BLOB
	DwFlags              uint32
	DwPriority           uint32
}

type PFN_CERT_ENUM_SYSTEM_STORE_LOCATION uintptr

type PFN_CERT_ENUM_SYSTEM_STORE uintptr

type PFN_CERT_ENUM_PHYSICAL_STORE uintptr

type CTL_VERIFY_USAGE_PARA struct {
	CbSize         uint32
	ListIdentifier CRYPT_INTEGER_BLOB
	CCtlStore      uint32
	RghCtlStore    *HCERTSTORE
	CSignerStore   uint32
	RghSignerStore *HCERTSTORE
}

type CTL_VERIFY_USAGE_STATUS struct {
	CbSize          uint32
	DwError         uint32
	DwFlags         uint32
	PpCtl           **CTL_CONTEXT
	DwCtlEntryIndex uint32
	PpSigner        **CERT_CONTEXT
	DwSignerIndex   uint32
}

type CERT_REVOCATION_PARA struct {
	CbSize       uint32
	PIssuerCert  *CERT_CONTEXT
	CCertStore   uint32
	RgCertStore  *HCERTSTORE
	HCrlStore    HCERTSTORE
	PftTimeToUse *FILETIME
}

type CERT_REVOCATION_STATUS struct {
	CbSize            uint32
	DwIndex           uint32
	DwError           uint32
	DwReason          CERT_REVOCATION_STATUS_REASON
	FHasFreshnessTime BOOL
	DwFreshnessTime   uint32
}

type CRYPT_SIGN_MESSAGE_PARA struct {
	CbSize             uint32
	DwMsgEncodingType  uint32
	PSigningCert       *CERT_CONTEXT
	HashAlgorithm      CRYPT_ALGORITHM_IDENTIFIER
	PvHashAuxInfo      unsafe.Pointer
	CMsgCert           uint32
	RgpMsgCert         **CERT_CONTEXT
	CMsgCrl            uint32
	RgpMsgCrl          **CRL_CONTEXT
	CAuthAttr          uint32
	RgAuthAttr         *CRYPT_ATTRIBUTE
	CUnauthAttr        uint32
	RgUnauthAttr       *CRYPT_ATTRIBUTE
	DwFlags            uint32
	DwInnerContentType uint32
}

type CRYPT_VERIFY_MESSAGE_PARA struct {
	CbSize                   uint32
	DwMsgAndCertEncodingType uint32
	HCryptProv               HCRYPTPROV_LEGACY
	PfnGetSignerCertificate  PFN_CRYPT_GET_SIGNER_CERTIFICATE
	PvGetArg                 unsafe.Pointer
}

type CRYPT_ENCRYPT_MESSAGE_PARA struct {
	CbSize                     uint32
	DwMsgEncodingType          uint32
	HCryptProv                 HCRYPTPROV_LEGACY
	ContentEncryptionAlgorithm CRYPT_ALGORITHM_IDENTIFIER
	PvEncryptionAuxInfo        unsafe.Pointer
	DwFlags                    uint32
	DwInnerContentType         uint32
}

type CRYPT_DECRYPT_MESSAGE_PARA struct {
	CbSize                   uint32
	DwMsgAndCertEncodingType uint32
	CCertStore               uint32
	RghCertStore             *HCERTSTORE
}

type CRYPT_HASH_MESSAGE_PARA struct {
	CbSize            uint32
	DwMsgEncodingType uint32
	HCryptProv        HCRYPTPROV_LEGACY
	HashAlgorithm     CRYPT_ALGORITHM_IDENTIFIER
	PvHashAuxInfo     unsafe.Pointer
}

type CRYPT_KEY_SIGN_MESSAGE_PARA struct {
	CbSize                   uint32
	DwMsgAndCertEncodingType CERT_QUERY_ENCODING_TYPE
	HCryptProv               uintptr
	DwKeySpec                CERT_KEY_SPEC
	HashAlgorithm            CRYPT_ALGORITHM_IDENTIFIER
	PvHashAuxInfo            unsafe.Pointer
	PubKeyAlgorithm          CRYPT_ALGORITHM_IDENTIFIER
}

type CRYPT_KEY_VERIFY_MESSAGE_PARA struct {
	CbSize            uint32
	DwMsgEncodingType uint32
	HCryptProv        HCRYPTPROV_LEGACY
}

type CERT_CHAIN struct {
	CCerts         uint32
	Certs          *CRYPT_INTEGER_BLOB
	KeyLocatorInfo CRYPT_KEY_PROV_INFO
}

type PFN_CRYPT_ASYNC_PARAM_FREE_FUNC uintptr

type CRYPT_CREDENTIALS struct {
	CbSize            uint32
	PszCredentialsOid *PSTRElement
	PvCredentials     unsafe.Pointer
}

type CRYPT_RETRIEVE_AUX_INFO struct {
	CbSize                     uint32
	PLastSyncTime              *FILETIME
	DwMaxUrlRetrievalByteCount uint32
	PPreFetchInfo              *CRYPTNET_URL_CACHE_PRE_FETCH_INFO
	PFlushInfo                 *CRYPTNET_URL_CACHE_FLUSH_INFO
	PpResponseInfo             **CRYPTNET_URL_CACHE_RESPONSE_INFO
	PwszCacheFileNamePrefix    *PWSTRElement
	PftCacheResync             *FILETIME
	FProxyCacheRetrieval       BOOL
	DwHttpStatusCode           uint32
	PpwszErrorResponseHeaders  **PWSTRElement
	PpErrorContentBlob         **CRYPT_INTEGER_BLOB
}

type PFN_CRYPT_CANCEL_RETRIEVAL uintptr

type CRYPT_URL_ARRAY struct {
	CUrl     uint32
	RgwszUrl **PWSTRElement
}

type CRYPT_URL_INFO struct {
	CbSize          uint32
	DwSyncDeltaTime uint32
	CGroup          uint32
	RgcGroupEntry   *uint32
}

type PFN_CRYPT_ENUM_KEYID_PROP uintptr

type CERT_CHAIN_ENGINE_CONFIG struct {
	CbSize                    uint32
	HRestrictedRoot           HCERTSTORE
	HRestrictedTrust          HCERTSTORE
	HRestrictedOther          HCERTSTORE
	CAdditionalStore          uint32
	RghAdditionalStore        *HCERTSTORE
	DwFlags                   uint32
	DwUrlRetrievalTimeout     uint32
	MaximumCachedCertificates uint32
	CycleDetectionModulus     uint32
	HExclusiveRoot            HCERTSTORE
	HExclusiveTrustedPeople   HCERTSTORE
	DwExclusiveFlags          uint32
}

type CERT_CHAIN_CONTEXT struct {
	CbSize                      uint32
	TrustStatus                 CERT_TRUST_STATUS
	CChain                      uint32
	RgpChain                    **CERT_SIMPLE_CHAIN
	CLowerQualityChainContext   uint32
	RgpLowerQualityChainContext **CERT_CHAIN_CONTEXT
	FHasRevocationFreshnessTime BOOL
	DwRevocationFreshnessTime   uint32
	DwCreateFlags               uint32
	ChainId                     Guid
}

type CERT_CHAIN_PARA struct {
	CbSize                        uint32
	RequestedUsage                CERT_USAGE_MATCH
	RequestedIssuancePolicy       CERT_USAGE_MATCH
	DwUrlRetrievalTimeout         uint32
	FCheckRevocationFreshnessTime BOOL
	DwRevocationFreshnessTime     uint32
	PftCacheResync                *FILETIME
	PStrongSignPara               *CERT_STRONG_SIGN_PARA
	DwStrongSignFlags             uint32
}

type CERT_CHAIN_POLICY_PARA struct {
	CbSize            uint32
	DwFlags           CERT_CHAIN_POLICY_FLAGS
	PvExtraPolicyPara unsafe.Pointer
}

type CERT_CHAIN_POLICY_STATUS struct {
	CbSize              uint32
	DwError             uint32
	LChainIndex         int32
	LElementIndex       int32
	PvExtraPolicyStatus unsafe.Pointer
}

type CERT_SERVER_OCSP_RESPONSE_CONTEXT struct {
	CbSize                uint32
	PbEncodedOcspResponse *uint8
	CbEncodedOcspResponse uint32
}

type CERT_SERVER_OCSP_RESPONSE_OPEN_PARA struct {
	CbSize              uint32
	DwFlags             uint32
	PcbUsedSize         *uint32
	PwszOcspDirectory   *PWSTRElement
	PfnUpdateCallback   PFN_CERT_SERVER_OCSP_RESPONSE_UPDATE_CALLBACK
	PvUpdateCallbackArg unsafe.Pointer
}

type CERT_SELECT_CHAIN_PARA struct {
	HChainEngine     HCERTCHAINENGINE
	PTime            *FILETIME
	HAdditionalStore HCERTSTORE
	PChainPara       *CERT_CHAIN_PARA
	DwFlags          uint32
}

type CERT_SELECT_CRITERIA struct {
	DwType CERT_SELECT_CRITERIA_TYPE
	CPara  uint32
	PpPara *unsafe.Pointer
}

type CRYPT_TIMESTAMP_CONTEXT struct {
	CbEncoded  uint32
	PbEncoded  *uint8
	PTimeStamp *CRYPT_TIMESTAMP_INFO
}

type CRYPT_TIMESTAMP_PARA struct {
	PszTSAPolicyId *PSTRElement
	FRequestCerts  BOOL
	Nonce          CRYPT_INTEGER_BLOB
	CExtension     uint32
	RgExtension    *CERT_EXTENSION
}

type CRYPTPROTECT_PROMPTSTRUCT struct {
	CbSize        uint32
	DwPromptFlags uint32
	HwndApp       HWND
	SzPrompt      *PWSTRElement
}

type NCRYPT_PROTECT_STREAM_INFO struct {
	PfnStreamOutput PFNCryptStreamOutputCallback
	PvCallbackCtxt  unsafe.Pointer
}

type NCRYPT_PROTECT_STREAM_INFO_EX struct {
	PfnStreamOutput PFNCryptStreamOutputCallbackEx
	PvCallbackCtxt  unsafe.Pointer
}

type SIGNER_CERT struct {
	CbSize       uint32
	DwCertChoice SIGNER_CERT_CHOICE
	PwszSpcFile  *PWSTRElement
	Hwnd         HWND
}

type SIGNER_CONTEXT struct {
	CbSize uint32
	CbBlob uint32
	PbBlob *uint8
}

type SIGNER_DIGEST_SIGN_INFO struct {
	CbSize                    uint32
	DwDigestSignChoice        uint32
	PfnAuthenticodeDigestSign PFN_AUTHENTICODE_DIGEST_SIGN
	PMetadataBlob             *CRYPT_INTEGER_BLOB
	DwReserved                uint32
	DwReserved2               uint32
	DwReserved3               uint32
}

type SIGNER_PROVIDER_INFO struct {
	CbSize           uint32
	PwszProviderName *PWSTRElement
	DwProviderType   uint32
	DwKeySpec        uint32
	DwPvkChoice      SIGNER_PRIVATE_KEY_CHOICE
	PwszPvkFileName  *PWSTRElement
}

type SIGNER_SIGNATURE_INFO struct {
	CbSize            uint32
	AlgidHash         ALG_ID
	DwAttrChoice      SIGNER_SIGNATURE_ATTRIBUTE_CHOICE
	PAttrAuthcode     *SIGNER_ATTR_AUTHCODE
	PsAuthenticated   *CRYPT_ATTRIBUTES
	PsUnauthenticated *CRYPT_ATTRIBUTES
}

type SIGNER_SUBJECT_INFO struct {
	CbSize          uint32
	PdwIndex        *uint32
	DwSubjectChoice SIGNER_SUBJECT_CHOICE
	PSignerFileInfo *SIGNER_FILE_INFO
}

type CRYPT_XML_CHARSET int32

const (
	CRYPT_XML_CHARSET_AUTO    CRYPT_XML_CHARSET = 0x0
	CRYPT_XML_CHARSET_UTF8    CRYPT_XML_CHARSET = 0x1
	CRYPT_XML_CHARSET_UTF16LE CRYPT_XML_CHARSET = 0x2
	CRYPT_XML_CHARSET_UTF16BE CRYPT_XML_CHARSET = 0x3
)

type CRYPT_XML_BLOB struct {
	DwCharset CRYPT_XML_CHARSET
	CbData    uint32
	PbData    *uint8
}

type CRYPT_XML_PROPERTY struct {
	DwPropId CRYPT_XML_PROPERTY_ID
	PvValue  unsafe.Pointer
	CbValue  uint32
}

type PFN_CRYPT_XML_WRITE_CALLBACK uintptr

type CRYPT_XML_DATA_PROVIDER struct {
	PvCallbackState unsafe.Pointer
	CbBufferSize    uint32
	PfnRead         PFN_CRYPT_XML_DATA_PROVIDER_READ
	PfnClose        PFN_CRYPT_XML_DATA_PROVIDER_CLOSE
}

type CRYPT_XML_STATUS struct {
	CbSize        uint32
	DwErrorStatus CRYPT_XML_STATUS_ERROR_STATUS
	DwInfoStatus  CRYPT_XML_STATUS_INFO_STATUS
}

type CRYPT_XML_ALGORITHM struct {
	CbSize       uint32
	WszAlgorithm *PWSTRElement
	Encoded      CRYPT_XML_BLOB
}

type CRYPT_XML_TRANSFORM_CHAIN_CONFIG struct {
	CbSize           uint32
	CTransformInfo   uint32
	RgpTransformInfo **CRYPT_XML_TRANSFORM_INFO
}

type CRYPT_XML_KEY_VALUE struct {
	DwType      CRYPT_XML_KEY_VALUE_TYPE
	DSAKeyValue CRYPT_XML_KEY_DSA_KEY_VALUE
}

type CRYPT_XML_REFERENCE struct {
	CbSize       uint32
	HReference   unsafe.Pointer
	WszId        *PWSTRElement
	WszUri       *PWSTRElement
	WszType      *PWSTRElement
	DigestMethod CRYPT_XML_ALGORITHM
	DigestValue  CRYPT_INTEGER_BLOB
	CTransform   uint32
	RgTransform  *CRYPT_XML_ALGORITHM
}

type CRYPT_XML_OBJECT struct {
	CbSize      uint32
	HObject     unsafe.Pointer
	WszId       *PWSTRElement
	WszMimeType *PWSTRElement
	WszEncoding *PWSTRElement
	Manifest    CRYPT_XML_REFERENCES
	Encoded     CRYPT_XML_BLOB
}

type CRYPT_XML_SIGNATURE struct {
	CbSize         uint32
	HSignature     unsafe.Pointer
	WszId          *PWSTRElement
	SignedInfo     CRYPT_XML_SIGNED_INFO
	SignatureValue CRYPT_INTEGER_BLOB
	PKeyInfo       *CRYPT_XML_KEY_INFO
	CObject        uint32
	RgpObject      **CRYPT_XML_OBJECT
}

type CRYPT_XML_DOC_CTXT struct {
	CbSize            uint32
	HDocCtxt          unsafe.Pointer
	PTransformsConfig *CRYPT_XML_TRANSFORM_CHAIN_CONFIG
	CSignature        uint32
	RgpSignature      **CRYPT_XML_SIGNATURE
}

type CRYPT_XML_KEYINFO_SPEC int32

const (
	CRYPT_XML_KEYINFO_SPEC_NONE    CRYPT_XML_KEYINFO_SPEC = 0x0
	CRYPT_XML_KEYINFO_SPEC_ENCODED CRYPT_XML_KEYINFO_SPEC = 0x1
	CRYPT_XML_KEYINFO_SPEC_PARAM   CRYPT_XML_KEYINFO_SPEC = 0x2
)

type CRYPT_XML_ALGORITHM_INFO struct {
	CbSize           uint32
	WszAlgorithmURI  *PWSTRElement
	WszName          *PWSTRElement
	DwGroupId        CRYPT_XML_GROUP_ID
	WszCNGAlgid      *PWSTRElement
	WszCNGExtraAlgid *PWSTRElement
	DwSignFlags      uint32
	DwVerifyFlags    uint32
	PvPaddingInfo    unsafe.Pointer
	PvExtraInfo      unsafe.Pointer
}

type PFN_CRYPT_XML_ENUM_ALG_INFO uintptr

type PaddingMode int32

const (
	None     PaddingMode = 0x1
	PKCS7    PaddingMode = 0x2
	Zeros    PaddingMode = 0x3
	ANSIX923 PaddingMode = 0x4
	ISO10126 PaddingMode = 0x5
)

type Direction int32

const (
	DirectionEncrypt Direction = 0x1
	DirectionDecrypt Direction = 0x2
)

type INFORMATIONCARD_CRYPTO_HANDLE struct {
	Type             HandleType
	Expiration       int64
	CryptoParameters unsafe.Pointer
}

type GENERIC_XML_TOKEN struct {
	CreateDate             FILETIME
	ExpiryDate             FILETIME
	XmlToken               *PWSTRElement
	InternalTokenReference *PWSTRElement
	ExternalTokenReference *PWSTRElement
}

type POLICY_ELEMENT struct {
	TargetEndpointAddress  *PWSTRElement
	IssuerEndpointAddress  *PWSTRElement
	IssuedTokenParameters  *PWSTRElement
	PrivacyNoticeLink      *PWSTRElement
	PrivacyNoticeVersion   uint32
	UseManagedPresentation BOOL
}

type BCRYPT_CIPHER_FUNCTION_TABLE struct {
	Version                BCRYPT_INTERFACE_VERSION
	OpenAlgorithmProvider  BCryptOpenAlgorithmProviderFn
	GetProperty            BCryptGetPropertyFn
	SetProperty            BCryptSetPropertyFn
	CloseAlgorithmProvider BCryptCloseAlgorithmProviderFn
	GenerateKey            BCryptGenerateSymmetricKeyFn
	Encrypt                BCryptEncryptFn
	Decrypt                BCryptDecryptFn
	ImportKey              BCryptImportKeyFn
	ExportKey              BCryptExportKeyFn
	DuplicateKey           BCryptDuplicateKeyFn
	DestroyKey             BCryptDestroyKeyFn
}

type BCRYPT_HASH_FUNCTION_TABLE struct {
	Version                BCRYPT_INTERFACE_VERSION
	OpenAlgorithmProvider  BCryptOpenAlgorithmProviderFn
	GetProperty            BCryptGetPropertyFn
	SetProperty            BCryptSetPropertyFn
	CloseAlgorithmProvider BCryptCloseAlgorithmProviderFn
	CreateHash             BCryptCreateHashFn
	HashData               BCryptHashDataFn
	FinishHash             BCryptFinishHashFn
	DuplicateHash          BCryptDuplicateHashFn
	DestroyHash            BCryptDestroyHashFn
	CreateMultiHash        BCryptCreateMultiHashFn
	ProcessMultiOperations BCryptProcessMultiOperationsFn
}

type BCRYPT_ASYMMETRIC_ENCRYPTION_FUNCTION_TABLE struct {
	Version                BCRYPT_INTERFACE_VERSION
	OpenAlgorithmProvider  BCryptOpenAlgorithmProviderFn
	GetProperty            BCryptGetPropertyFn
	SetProperty            BCryptSetPropertyFn
	CloseAlgorithmProvider BCryptCloseAlgorithmProviderFn
	GenerateKeyPair        BCryptGenerateKeyPairFn
	FinalizeKeyPair        BCryptFinalizeKeyPairFn
	Encrypt                BCryptEncryptFn
	Decrypt                BCryptDecryptFn
	ImportKeyPair          BCryptImportKeyPairFn
	ExportKey              BCryptExportKeyFn
	DestroyKey             BCryptDestroyKeyFn
	SignHash               BCryptSignHashFn
	VerifySignature        BCryptVerifySignatureFn
}

type BCRYPT_SECRET_AGREEMENT_FUNCTION_TABLE struct {
	Version                BCRYPT_INTERFACE_VERSION
	OpenAlgorithmProvider  BCryptOpenAlgorithmProviderFn
	GetProperty            BCryptGetPropertyFn
	SetProperty            BCryptSetPropertyFn
	CloseAlgorithmProvider BCryptCloseAlgorithmProviderFn
	SecretAgreement        BCryptSecretAgreementFn
	DeriveKey              BCryptDeriveKeyFn
	DestroySecret          BCryptDestroySecretFn
	GenerateKeyPair        BCryptGenerateKeyPairFn
	FinalizeKeyPair        BCryptFinalizeKeyPairFn
	ImportKeyPair          BCryptImportKeyPairFn
	ExportKey              BCryptExportKeyFn
	DestroyKey             BCryptDestroyKeyFn
}

type BCRYPT_SIGNATURE_FUNCTION_TABLE struct {
	Version                BCRYPT_INTERFACE_VERSION
	OpenAlgorithmProvider  BCryptOpenAlgorithmProviderFn
	GetProperty            BCryptGetPropertyFn
	SetProperty            BCryptSetPropertyFn
	CloseAlgorithmProvider BCryptCloseAlgorithmProviderFn
	GenerateKeyPair        BCryptGenerateKeyPairFn
	FinalizeKeyPair        BCryptFinalizeKeyPairFn
	SignHash               BCryptSignHashFn
	VerifySignature        BCryptVerifySignatureFn
	ImportKeyPair          BCryptImportKeyPairFn
	ExportKey              BCryptExportKeyFn
	DestroyKey             BCryptDestroyKeyFn
}

type BCRYPT_RNG_FUNCTION_TABLE struct {
	Version                BCRYPT_INTERFACE_VERSION
	OpenAlgorithmProvider  BCryptOpenAlgorithmProviderFn
	GetProperty            BCryptGetPropertyFn
	SetProperty            BCryptSetPropertyFn
	CloseAlgorithmProvider BCryptCloseAlgorithmProviderFn
	GenRandom              BCryptGenRandomFn
}

type BCRYPT_KEY_DERIVATION_FUNCTION_TABLE struct {
	Version                BCRYPT_INTERFACE_VERSION
	OpenAlgorithmProvider  BCryptOpenAlgorithmProviderFn
	GetProperty            BCryptGetPropertyFn
	SetProperty            BCryptSetPropertyFn
	CloseAlgorithmProvider BCryptCloseAlgorithmProviderFn
	GenerateKey            BCryptGenerateSymmetricKeyFn
	DestroyKey             BCryptDestroyKeyFn
	KeyDerivation          BCryptKeyDerivationFn
	ExportKey              BCryptExportKeyFn
	ImportKey              BCryptImportKeyFn
	DuplicateKey           BCryptDuplicateKeyFn
}

type NCRYPT_KEY_STORAGE_FUNCTION_TABLE struct {
	Version             BCRYPT_INTERFACE_VERSION
	OpenProvider        NCryptOpenStorageProviderFn
	OpenKey             NCryptOpenKeyFn
	CreatePersistedKey  NCryptCreatePersistedKeyFn
	GetProviderProperty NCryptGetProviderPropertyFn
	GetKeyProperty      NCryptGetKeyPropertyFn
	SetProviderProperty NCryptSetProviderPropertyFn
	SetKeyProperty      NCryptSetKeyPropertyFn
	FinalizeKey         NCryptFinalizeKeyFn
	DeleteKey           NCryptDeleteKeyFn
	FreeProvider        NCryptFreeProviderFn
	FreeKey             NCryptFreeKeyFn
	FreeBuffer          NCryptFreeBufferFn
	Encrypt             NCryptEncryptFn
	Decrypt             NCryptDecryptFn
	IsAlgSupported      NCryptIsAlgSupportedFn
	EnumAlgorithms      NCryptEnumAlgorithmsFn
	EnumKeys            NCryptEnumKeysFn
	ImportKey           NCryptImportKeyFn
	ExportKey           NCryptExportKeyFn
	SignHash            NCryptSignHashFn
	VerifySignature     NCryptVerifySignatureFn
	PromptUser          NCryptPromptUserFn
	NotifyChangeKey     NCryptNotifyChangeKeyFn
	SecretAgreement     NCryptSecretAgreementFn
	DeriveKey           NCryptDeriveKeyFn
	FreeSecret          NCryptFreeSecretFn
	KeyDerivation       NCryptKeyDerivationFn
	CreateClaim         NCryptCreateClaimFn
	VerifyClaim         NCryptVerifyClaimFn
}

type NCRYPT_SSL_CIPHER_SUITE struct {
	DwProtocol        uint32
	DwCipherSuite     uint32
	DwBaseCipherSuite uint32
	SzCipherSuite     [64]uint16
	SzCipher          [64]uint16
	DwCipherLen       uint32
	DwCipherBlockLen  uint32
	SzHash            [64]uint16
	DwHashLen         uint32
	SzExchange        [64]uint16
	DwMinExchangeLen  uint32
	DwMaxExchangeLen  uint32
	SzCertificate     [64]uint16
	DwKeyType         uint32
}

type NCRYPT_SSL_CIPHER_SUITE_EX struct {
	DwVersion         uint32
	DwProtocol        uint32
	DwCipherSuite     uint32
	DwBaseCipherSuite uint32
	SzCipherSuite     [64]uint16
	SzCipher          [64]uint16
	DwCipherLen       uint32
	DwCipherBlockLen  uint32
	SzHash            [64]uint16
	DwHashLen         uint32
	SzExchange        [64]uint16
	DwMinExchangeLen  uint32
	DwMaxExchangeLen  uint32
	SzCertificate     [64]uint16
	DwKeyType         uint32
	SzCipherMode      [64]uint16
}

type NCRYPT_SSL_CIPHER_LENGTHS struct {
	CbLength                uint32
	DwHeaderLen             uint32
	DwFixedTrailerLen       uint32
	DwMaxVariableTrailerLen uint32
	DwFlags                 uint32
}

type NCRYPT_SSL_ECC_CURVE struct {
	SzCurveName       [255]uint16
	SzOID             [255]CHAR
	DwPublicKeyLength uint32
	DwCurveType       uint32
	DwFlags           uint32
}

type NCRYPT_SSL_FUNCTION_TABLE struct {
	Version                        BCRYPT_INTERFACE_VERSION
	ComputeClientAuthHash          SslComputeClientAuthHashFn
	ComputeEapKeyBlock             SslComputeEapKeyBlockFn
	ComputeFinishedHash            SslComputeFinishedHashFn
	CreateEphemeralKey             SslCreateEphemeralKeyFn
	CreateHandshakeHash            SslCreateHandshakeHashFn
	DecryptPacket                  SslDecryptPacketFn
	EncryptPacket                  SslEncryptPacketFn
	EnumCipherSuites               SslEnumCipherSuitesFn
	ExportKey                      SslExportKeyFn
	FreeBuffer                     SslFreeBufferFn
	FreeObject                     SslFreeObjectFn
	GenerateMasterKey              SslGenerateMasterKeyFn
	GenerateSessionKeys            SslGenerateSessionKeysFn
	GetKeyProperty                 SslGetKeyPropertyFn
	GetProviderProperty            SslGetProviderPropertyFn
	HashHandshake                  SslHashHandshakeFn
	ImportMasterKey                SslImportMasterKeyFn
	ImportKey                      SslImportKeyFn
	LookupCipherSuiteInfo          SslLookupCipherSuiteInfoFn
	OpenPrivateKey                 SslOpenPrivateKeyFn
	OpenProvider                   SslOpenProviderFn
	SignHash                       SslSignHashFn
	VerifySignature                SslVerifySignatureFn
	LookupCipherLengths            SslLookupCipherLengthsFn
	CreateClientAuthHash           SslCreateClientAuthHashFn
	GetCipherSuitePRFHashAlgorithm SslGetCipherSuitePRFHashAlgorithmFn
	ComputeSessionHash             SslComputeSessionHashFn
	GeneratePreMasterKey           SslGeneratePreMasterKeyFn
	EnumEccCurves                  SslEnumEccCurvesFn
	ExportKeyingMaterial           SslExportKeyingMaterialFn
	ExtractEarlyKey                SslExtractEarlyKeyFn
	ExtractHandshakeKey            SslExtractHandshakeKeyFn
	ExtractMasterKey               SslExtractMasterKeyFn
	ExpandTrafficKeys              SslExpandTrafficKeysFn
	ExpandWriteKey                 SslExpandWriteKeyFn
	ExpandExporterMasterKey        SslExpandExporterMasterKeyFn
	EnumCipherSuitesEx             SslEnumCipherSuitesExFn
	ExpandResumptionMasterKey      SslExpandResumptionMasterKeyFn
	DuplicateTranscriptHash        SslDuplicateTranscriptHashFn
	ExpandBinderKey                SslExpandBinderKeyFn
	ExpandPreSharedKey             SslExpandPreSharedKeyFn
}

type WIN_CERTIFICATE struct {
	DwLength         uint32
	WRevision        uint16
	WCertificateType uint16
	BCertificate     [1]uint8
}

type LDT_ENTRY struct {
	LimitLow uint16
	BaseLow  uint16
	BaseMid  uint8
	Flags1   uint8
	Flags2   uint8
	BaseHi   uint8
}

type WOW64_CONTEXT struct {
	ContextFlags      WOW64_CONTEXT_FLAGS
	Dr0               uint32
	Dr1               uint32
	Dr2               uint32
	Dr3               uint32
	Dr6               uint32
	Dr7               uint32
	FloatSave         WOW64_FLOATING_SAVE_AREA
	SegGs             uint32
	SegFs             uint32
	SegEs             uint32
	SegDs             uint32
	Edi               uint32
	Esi               uint32
	Ebx               uint32
	Edx               uint32
	Ecx               uint32
	Eax               uint32
	Ebp               uint32
	Eip               uint32
	SegCs             uint32
	EFlags            uint32
	Esp               uint32
	SegSs             uint32
	ExtendedRegisters [512]uint8
}

type WOW64_LDT_ENTRY struct {
	LimitLow uint16
	BaseLow  uint16
	BaseMid  uint8
	Flags1   uint8
	Flags2   uint8
	BaseHi   uint8
}

type EXCEPTION_RECORD struct {
	ExceptionCode        NTSTATUS
	ExceptionFlags       uint32
	ExceptionRecord      *EXCEPTION_RECORD
	ExceptionAddress     unsafe.Pointer
	NumberParameters     uint32
	ExceptionInformation [15]uintptr
}

type EXCEPTION_POINTERS struct {
	ExceptionRecord *EXCEPTION_RECORD
	ContextRecord   *CONTEXT
}

type IMAGE_NT_HEADERS64 struct {
	Signature      uint32
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER64
}

type IMAGE_NT_HEADERS32 struct {
	Signature      uint32
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER32
}

type IMAGE_SECTION_HEADER struct {
	Name                 [8]uint8
	PhysicalAddress      uint32
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	PointerToLinenumbers uint32
	NumberOfRelocations  uint16
	NumberOfLinenumbers  uint16
	Characteristics      IMAGE_SECTION_CHARACTERISTICS
}

type IMAGE_LOAD_CONFIG_DIRECTORY32 struct {
	Size                                     uint32
	TimeDateStamp                            uint32
	MajorVersion                             uint16
	MinorVersion                             uint16
	GlobalFlagsClear                         uint32
	GlobalFlagsSet                           uint32
	CriticalSectionDefaultTimeout            uint32
	DeCommitFreeBlockThreshold               uint32
	DeCommitTotalFreeThreshold               uint32
	LockPrefixTable                          uint32
	MaximumAllocationSize                    uint32
	VirtualMemoryThreshold                   uint32
	ProcessHeapFlags                         uint32
	ProcessAffinityMask                      uint32
	CSDVersion                               uint16
	DependentLoadFlags                       uint16
	EditList                                 uint32
	SecurityCookie                           uint32
	SEHandlerTable                           uint32
	SEHandlerCount                           uint32
	GuardCFCheckFunctionPointer              uint32
	GuardCFDispatchFunctionPointer           uint32
	GuardCFFunctionTable                     uint32
	GuardCFFunctionCount                     uint32
	GuardFlags                               uint32
	CodeIntegrity                            IMAGE_LOAD_CONFIG_CODE_INTEGRITY
	GuardAddressTakenIatEntryTable           uint32
	GuardAddressTakenIatEntryCount           uint32
	GuardLongJumpTargetTable                 uint32
	GuardLongJumpTargetCount                 uint32
	DynamicValueRelocTable                   uint32
	CHPEMetadataPointer                      uint32
	GuardRFFailureRoutine                    uint32
	GuardRFFailureRoutineFunctionPointer     uint32
	DynamicValueRelocTableOffset             uint32
	DynamicValueRelocTableSection            uint16
	Reserved2                                uint16
	GuardRFVerifyStackPointerFunctionPointer uint32
	HotPatchTableOffset                      uint32
	Reserved3                                uint32
	EnclaveConfigurationPointer              uint32
	VolatileMetadataPointer                  uint32
	GuardEHContinuationTable                 uint32
	GuardEHContinuationCount                 uint32
	GuardXFGCheckFunctionPointer             uint32
	GuardXFGDispatchFunctionPointer          uint32
	GuardXFGTableDispatchFunctionPointer     uint32
	CastGuardOsDeterminedFailureMode         uint32
	GuardMemcpyFunctionPointer               uint32
	UmaFunctionPointers                      uint32
}

type IMAGE_LOAD_CONFIG_DIRECTORY64 struct {
	Size                                     uint32
	TimeDateStamp                            uint32
	MajorVersion                             uint16
	MinorVersion                             uint16
	GlobalFlagsClear                         uint32
	GlobalFlagsSet                           uint32
	CriticalSectionDefaultTimeout            uint32
	DeCommitFreeBlockThreshold               uint64
	DeCommitTotalFreeThreshold               uint64
	LockPrefixTable                          uint64
	MaximumAllocationSize                    uint64
	VirtualMemoryThreshold                   uint64
	ProcessAffinityMask                      uint64
	ProcessHeapFlags                         uint32
	CSDVersion                               uint16
	DependentLoadFlags                       uint16
	EditList                                 uint64
	SecurityCookie                           uint64
	SEHandlerTable                           uint64
	SEHandlerCount                           uint64
	GuardCFCheckFunctionPointer              uint64
	GuardCFDispatchFunctionPointer           uint64
	GuardCFFunctionTable                     uint64
	GuardCFFunctionCount                     uint64
	GuardFlags                               uint32
	CodeIntegrity                            IMAGE_LOAD_CONFIG_CODE_INTEGRITY
	GuardAddressTakenIatEntryTable           uint64
	GuardAddressTakenIatEntryCount           uint64
	GuardLongJumpTargetTable                 uint64
	GuardLongJumpTargetCount                 uint64
	DynamicValueRelocTable                   uint64
	CHPEMetadataPointer                      uint64
	GuardRFFailureRoutine                    uint64
	GuardRFFailureRoutineFunctionPointer     uint64
	DynamicValueRelocTableOffset             uint32
	DynamicValueRelocTableSection            uint16
	Reserved2                                uint16
	GuardRFVerifyStackPointerFunctionPointer uint64
	HotPatchTableOffset                      uint32
	Reserved3                                uint32
	EnclaveConfigurationPointer              uint64
	VolatileMetadataPointer                  uint64
	GuardEHContinuationTable                 uint64
	GuardEHContinuationCount                 uint64
	GuardXFGCheckFunctionPointer             uint64
	GuardXFGDispatchFunctionPointer          uint64
	GuardXFGTableDispatchFunctionPointer     uint64
	CastGuardOsDeterminedFailureMode         uint64
	GuardMemcpyFunctionPointer               uint64
	UmaFunctionPointers                      uint64
}

type IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY struct {
	BeginAddress uint32
	UnwindData   uint32
}

type IMAGE_RUNTIME_FUNCTION_ENTRY struct {
	BeginAddress      uint32
	EndAddress        uint32
	UnwindInfoAddress uint32
}

type PVECTORED_EXCEPTION_HANDLER uintptr

type FILE_SEGMENT_ELEMENT struct {
	Buffer unsafe.Pointer
}

type LPTOP_LEVEL_EXCEPTION_FILTER uintptr

type WAITCHAIN_NODE_INFO struct {
	ObjectType   WCT_OBJECT_TYPE
	ObjectStatus WCT_OBJECT_STATUS
	ObjectName   [128]uint16
	Timeout      int64
	Alertable    BOOL
}

type PWAITCHAINCALLBACK uintptr

type PCOGETCALLSTATE uintptr

type PCOGETACTIVATIONSTATE uintptr

type MINIDUMP_DIRECTORY struct {
	StreamType uint32
	Location   MINIDUMP_LOCATION_DESCRIPTOR
}

type MINIDUMP_TYPE int32

const (
	MiniDumpNormal                         MINIDUMP_TYPE = 0x0
	MiniDumpWithDataSegs                   MINIDUMP_TYPE = 0x1
	MiniDumpWithFullMemory                 MINIDUMP_TYPE = 0x2
	MiniDumpWithHandleData                 MINIDUMP_TYPE = 0x4
	MiniDumpFilterMemory                   MINIDUMP_TYPE = 0x8
	MiniDumpScanMemory                     MINIDUMP_TYPE = 0x10
	MiniDumpWithUnloadedModules            MINIDUMP_TYPE = 0x20
	MiniDumpWithIndirectlyReferencedMemory MINIDUMP_TYPE = 0x40
	MiniDumpFilterModulePaths              MINIDUMP_TYPE = 0x80
	MiniDumpWithProcessThreadData          MINIDUMP_TYPE = 0x100
	MiniDumpWithPrivateReadWriteMemory     MINIDUMP_TYPE = 0x200
	MiniDumpWithoutOptionalData            MINIDUMP_TYPE = 0x400
	MiniDumpWithFullMemoryInfo             MINIDUMP_TYPE = 0x800
	MiniDumpWithThreadInfo                 MINIDUMP_TYPE = 0x1000
	MiniDumpWithCodeSegs                   MINIDUMP_TYPE = 0x2000
	MiniDumpWithoutAuxiliaryState          MINIDUMP_TYPE = 0x4000
	MiniDumpWithFullAuxiliaryState         MINIDUMP_TYPE = 0x8000
	MiniDumpWithPrivateWriteCopyMemory     MINIDUMP_TYPE = 0x10000
	MiniDumpIgnoreInaccessibleMemory       MINIDUMP_TYPE = 0x20000
	MiniDumpWithTokenInformation           MINIDUMP_TYPE = 0x40000
	MiniDumpWithModuleHeaders              MINIDUMP_TYPE = 0x80000
	MiniDumpFilterTriage                   MINIDUMP_TYPE = 0x100000
	MiniDumpWithAvxXStateContext           MINIDUMP_TYPE = 0x200000
	MiniDumpWithIptTrace                   MINIDUMP_TYPE = 0x400000
	MiniDumpScanInaccessiblePartialPages   MINIDUMP_TYPE = 0x800000
	MiniDumpFilterWriteCombinedMemory      MINIDUMP_TYPE = 0x1000000
	MiniDumpValidTypeFlags                 MINIDUMP_TYPE = 0x1ffffff
	MiniDumpNoIgnoreInaccessibleMemory     MINIDUMP_TYPE = 0x2000000
	MiniDumpValidTypeFlagsEx               MINIDUMP_TYPE = 0x3ffffff
)

type PIMAGEHLP_STATUS_ROUTINE uintptr

type DIGEST_FUNCTION uintptr

type PFIND_DEBUG_FILE_CALLBACK uintptr

type PFIND_DEBUG_FILE_CALLBACKW uintptr

type PFINDFILEINPATHCALLBACK uintptr

type PFINDFILEINPATHCALLBACKW uintptr

type PFIND_EXE_FILE_CALLBACK uintptr

type PFIND_EXE_FILE_CALLBACKW uintptr

type PENUMDIRTREE_CALLBACK uintptr

type PENUMDIRTREE_CALLBACKW uintptr

type MODLOAD_DATA struct {
	Ssize uint32
	Ssig  MODLOAD_DATA_TYPE
	Data  unsafe.Pointer
	Size  uint32
	Flags uint32
}

type STACKFRAME64 struct {
	AddrPC         ADDRESS64
	AddrReturn     ADDRESS64
	AddrFrame      ADDRESS64
	AddrStack      ADDRESS64
	AddrBStore     ADDRESS64
	FuncTableEntry unsafe.Pointer
	Params         [4]uint64
	Far            BOOL
	Virtual        BOOL
	Reserved       [3]uint64
	KdHelp         KDHELP64
}

type STACKFRAME_EX struct {
	AddrPC             ADDRESS64
	AddrReturn         ADDRESS64
	AddrFrame          ADDRESS64
	AddrStack          ADDRESS64
	AddrBStore         ADDRESS64
	FuncTableEntry     unsafe.Pointer
	Params             [4]uint64
	Far                BOOL
	Virtual            BOOL
	Reserved           [3]uint64
	KdHelp             KDHELP64
	StackFrameSize     uint32
	InlineFrameContext uint32
}

type PREAD_PROCESS_MEMORY_ROUTINE64 uintptr

type PFUNCTION_TABLE_ACCESS_ROUTINE64 uintptr

type PGET_MODULE_BASE_ROUTINE64 uintptr

type PTRANSLATE_ADDRESS_ROUTINE64 uintptr

type PGET_TARGET_ATTRIBUTE_VALUE64 uintptr

type API_VERSION struct {
	MajorVersion uint16
	MinorVersion uint16
	Revision     uint16
	Reserved     uint16
}

type PSYM_ENUMMODULES_CALLBACK64 uintptr

type PSYM_ENUMMODULES_CALLBACKW64 uintptr

type PENUMLOADED_MODULES_CALLBACK64 uintptr

type PENUMLOADED_MODULES_CALLBACKW64 uintptr

type PSYM_ENUMSYMBOLS_CALLBACK64 uintptr

type PSYM_ENUMSYMBOLS_CALLBACK64W uintptr

type PSYMBOL_REGISTERED_CALLBACK64 uintptr

type PSYMBOL_FUNCENTRY_CALLBACK uintptr

type PSYMBOL_FUNCENTRY_CALLBACK64 uintptr

type IMAGEHLP_SYMBOL64 struct {
	SizeOfStruct  uint32
	Address       uint64
	Size          uint32
	Flags         uint32
	MaxNameLength uint32
	Name          [1]CHAR
}

type IMAGEHLP_MODULE64 struct {
	SizeOfStruct    uint32
	BaseOfImage     uint64
	ImageSize       uint32
	TimeDateStamp   uint32
	CheckSum        uint32
	NumSyms         uint32
	SymType         SYM_TYPE
	ModuleName      [32]CHAR
	ImageName       [256]CHAR
	LoadedImageName [256]CHAR
	LoadedPdbName   [256]CHAR
	CVSig           uint32
	CVData          [780]CHAR
	PdbSig          uint32
	PdbSig70        Guid
	PdbAge          uint32
	PdbUnmatched    BOOL
	DbgUnmatched    BOOL
	LineNumbers     BOOL
	GlobalSymbols   BOOL
	TypeInfo        BOOL
	SourceIndexed   BOOL
	Publics         BOOL
	MachineType     uint32
	Reserved        uint32
}

type IMAGEHLP_MODULEW64 struct {
	SizeOfStruct    uint32
	BaseOfImage     uint64
	ImageSize       uint32
	TimeDateStamp   uint32
	CheckSum        uint32
	NumSyms         uint32
	SymType         SYM_TYPE
	ModuleName      [32]uint16
	ImageName       [256]uint16
	LoadedImageName [256]uint16
	LoadedPdbName   [256]uint16
	CVSig           uint32
	CVData          [780]uint16
	PdbSig          uint32
	PdbSig70        Guid
	PdbAge          uint32
	PdbUnmatched    BOOL
	DbgUnmatched    BOOL
	LineNumbers     BOOL
	GlobalSymbols   BOOL
	TypeInfo        BOOL
	SourceIndexed   BOOL
	Publics         BOOL
	MachineType     uint32
	Reserved        uint32
}

type IMAGEHLP_LINE64 struct {
	SizeOfStruct uint32
	Key          unsafe.Pointer
	LineNumber   uint32
	FileName     *PSTRElement
	Address      uint64
}

type IMAGEHLP_LINEW64 struct {
	SizeOfStruct uint32
	Key          unsafe.Pointer
	LineNumber   uint32
	FileName     *PWSTRElement
	Address      uint64
}

type OMAP struct {
	Rva   uint32
	RvaTo uint32
}

type IMAGEHLP_EXTENDED_OPTIONS int32

const (
	SYMOPT_EX_DISABLEACCESSTIMEUPDATE IMAGEHLP_EXTENDED_OPTIONS = 0x0
	SYMOPT_EX_LASTVALIDDEBUGDIRECTORY IMAGEHLP_EXTENDED_OPTIONS = 0x1
	SYMOPT_EX_NOIMPLICITPATTERNSEARCH IMAGEHLP_EXTENDED_OPTIONS = 0x2
	SYMOPT_EX_NEVERLOADSYMBOLS        IMAGEHLP_EXTENDED_OPTIONS = 0x3
	SYMOPT_EX_MAX                     IMAGEHLP_EXTENDED_OPTIONS = 0x4
)

type PSYM_ENUMSOURCEFILES_CALLBACK uintptr

type PSYM_ENUMSOURCEFILES_CALLBACKW uintptr

type PSYM_ENUMLINES_CALLBACK uintptr

type PSYM_ENUMLINES_CALLBACKW uintptr

type PENUMSOURCEFILETOKENSCALLBACK uintptr

type SYMBOL_INFO struct {
	SizeOfStruct uint32
	TypeIndex    uint32
	Reserved     [2]uint64
	Index        uint32
	Size         uint32
	ModBase      uint64
	Flags        SYMBOL_INFO_FLAGS
	Value        uint64
	Address      uint64
	Register     uint32
	Scope        uint32
	Tag          uint32
	NameLen      uint32
	MaxNameLen   uint32
	Name         [1]CHAR
}

type SYMBOL_INFOW struct {
	SizeOfStruct uint32
	TypeIndex    uint32
	Reserved     [2]uint64
	Index        uint32
	Size         uint32
	ModBase      uint64
	Flags        SYMBOL_INFO_FLAGS
	Value        uint64
	Address      uint64
	Register     uint32
	Scope        uint32
	Tag          uint32
	NameLen      uint32
	MaxNameLen   uint32
	Name         [1]uint16
}

type IMAGEHLP_STACK_FRAME struct {
	InstructionOffset  uint64
	ReturnOffset       uint64
	FrameOffset        uint64
	StackOffset        uint64
	BackingStoreOffset uint64
	FuncTableEntry     uint64
	Params             [4]uint64
	Reserved           [5]uint64
	Virtual            BOOL
	Reserved2          uint32
}

type PSYM_ENUMPROCESSES_CALLBACK uintptr

type PSYM_ENUMERATESYMBOLS_CALLBACK uintptr

type PSYM_ENUMERATESYMBOLS_CALLBACKW uintptr

type IMAGEHLP_SYMBOL_TYPE_INFO int32

const (
	TI_GET_SYMTAG                             IMAGEHLP_SYMBOL_TYPE_INFO = 0x0
	TI_GET_SYMNAME                            IMAGEHLP_SYMBOL_TYPE_INFO = 0x1
	TI_GET_LENGTH                             IMAGEHLP_SYMBOL_TYPE_INFO = 0x2
	TI_GET_TYPE                               IMAGEHLP_SYMBOL_TYPE_INFO = 0x3
	TI_GET_TYPEID                             IMAGEHLP_SYMBOL_TYPE_INFO = 0x4
	TI_GET_BASETYPE                           IMAGEHLP_SYMBOL_TYPE_INFO = 0x5
	TI_GET_ARRAYINDEXTYPEID                   IMAGEHLP_SYMBOL_TYPE_INFO = 0x6
	TI_FINDCHILDREN                           IMAGEHLP_SYMBOL_TYPE_INFO = 0x7
	TI_GET_DATAKIND                           IMAGEHLP_SYMBOL_TYPE_INFO = 0x8
	TI_GET_ADDRESSOFFSET                      IMAGEHLP_SYMBOL_TYPE_INFO = 0x9
	TI_GET_OFFSET                             IMAGEHLP_SYMBOL_TYPE_INFO = 0xa
	TI_GET_VALUE                              IMAGEHLP_SYMBOL_TYPE_INFO = 0xb
	TI_GET_COUNT                              IMAGEHLP_SYMBOL_TYPE_INFO = 0xc
	TI_GET_CHILDRENCOUNT                      IMAGEHLP_SYMBOL_TYPE_INFO = 0xd
	TI_GET_BITPOSITION                        IMAGEHLP_SYMBOL_TYPE_INFO = 0xe
	TI_GET_VIRTUALBASECLASS                   IMAGEHLP_SYMBOL_TYPE_INFO = 0xf
	TI_GET_VIRTUALTABLESHAPEID                IMAGEHLP_SYMBOL_TYPE_INFO = 0x10
	TI_GET_VIRTUALBASEPOINTEROFFSET           IMAGEHLP_SYMBOL_TYPE_INFO = 0x11
	TI_GET_CLASSPARENTID                      IMAGEHLP_SYMBOL_TYPE_INFO = 0x12
	TI_GET_NESTED                             IMAGEHLP_SYMBOL_TYPE_INFO = 0x13
	TI_GET_SYMINDEX                           IMAGEHLP_SYMBOL_TYPE_INFO = 0x14
	TI_GET_LEXICALPARENT                      IMAGEHLP_SYMBOL_TYPE_INFO = 0x15
	TI_GET_ADDRESS                            IMAGEHLP_SYMBOL_TYPE_INFO = 0x16
	TI_GET_THISADJUST                         IMAGEHLP_SYMBOL_TYPE_INFO = 0x17
	TI_GET_UDTKIND                            IMAGEHLP_SYMBOL_TYPE_INFO = 0x18
	TI_IS_EQUIV_TO                            IMAGEHLP_SYMBOL_TYPE_INFO = 0x19
	TI_GET_CALLING_CONVENTION                 IMAGEHLP_SYMBOL_TYPE_INFO = 0x1a
	TI_IS_CLOSE_EQUIV_TO                      IMAGEHLP_SYMBOL_TYPE_INFO = 0x1b
	TI_GTIEX_REQS_VALID                       IMAGEHLP_SYMBOL_TYPE_INFO = 0x1c
	TI_GET_VIRTUALBASEOFFSET                  IMAGEHLP_SYMBOL_TYPE_INFO = 0x1d
	TI_GET_VIRTUALBASEDISPINDEX               IMAGEHLP_SYMBOL_TYPE_INFO = 0x1e
	TI_GET_IS_REFERENCE                       IMAGEHLP_SYMBOL_TYPE_INFO = 0x1f
	TI_GET_INDIRECTVIRTUALBASECLASS           IMAGEHLP_SYMBOL_TYPE_INFO = 0x20
	TI_GET_VIRTUALBASETABLETYPE               IMAGEHLP_SYMBOL_TYPE_INFO = 0x21
	TI_GET_OBJECTPOINTERTYPE                  IMAGEHLP_SYMBOL_TYPE_INFO = 0x22
	TI_GET_DISCRIMINATEDUNION_TAG_TYPEID      IMAGEHLP_SYMBOL_TYPE_INFO = 0x23
	TI_GET_DISCRIMINATEDUNION_TAG_OFFSET      IMAGEHLP_SYMBOL_TYPE_INFO = 0x24
	TI_GET_DISCRIMINATEDUNION_TAG_RANGESCOUNT IMAGEHLP_SYMBOL_TYPE_INFO = 0x25
	TI_GET_DISCRIMINATEDUNION_TAG_RANGES      IMAGEHLP_SYMBOL_TYPE_INFO = 0x26
	IMAGEHLP_SYMBOL_TYPE_INFO_MAX             IMAGEHLP_SYMBOL_TYPE_INFO = 0x27
)

type IMAGEHLP_GET_TYPE_INFO_PARAMS struct {
	SizeOfStruct   uint32
	Flags          IMAGEHLP_GET_TYPE_INFO_FLAGS
	NumIds         uint32
	TypeIds        *uint32
	TagFilter      uint64
	NumReqs        uint32
	ReqKinds       *IMAGEHLP_SYMBOL_TYPE_INFO
	ReqOffsets     *uintptr
	ReqSizes       *uint32
	ReqStride      uintptr
	BufferSize     uintptr
	Buffer         unsafe.Pointer
	EntriesMatched uint32
	EntriesFilled  uint32
	TagsFound      uint64
	AllReqsValid   uint64
	NumReqsValid   uint32
	ReqsValid      *uint64
}

type SYMSRV_INDEX_INFO struct {
	Sizeofstruct uint32
	File         [261]CHAR
	Stripped     BOOL
	Timestamp    uint32
	Size         uint32
	Dbgfile      [261]CHAR
	Pdbfile      [261]CHAR
	Guid         Guid
	Sig          uint32
	Age          uint32
}

type SYMSRV_INDEX_INFOW struct {
	Sizeofstruct uint32
	File         [261]uint16
	Stripped     BOOL
	Timestamp    uint32
	Size         uint32
	Dbgfile      [261]uint16
	Pdbfile      [261]uint16
	Guid         Guid
	Sig          uint32
	Age          uint32
}

type PDBGHELP_CREATE_USER_DUMP_CALLBACK uintptr

type LPCALL_BACK_USER_INTERRUPT_ROUTINE uintptr

type DBGHELP_DATA_REPORT_STRUCT struct {
	PBinPathNonExist    *PWSTRElement
	PSymbolPathNonExist *PWSTRElement
}

type FILE_WRITE_FLAGS int32

const (
	FILE_WRITE_FLAGS_NONE          FILE_WRITE_FLAGS = 0x0
	FILE_WRITE_FLAGS_WRITE_THROUGH FILE_WRITE_FLAGS = 0x1
)

type FILE_FLUSH_MODE int32

const (
	FILE_FLUSH_DEFAULT      FILE_FLUSH_MODE = 0x0
	FILE_FLUSH_DATA         FILE_FLUSH_MODE = 0x1
	FILE_FLUSH_MIN_METADATA FILE_FLUSH_MODE = 0x2
	FILE_FLUSH_NO_SYNC      FILE_FLUSH_MODE = 0x3
)

type OFSTRUCT struct {
	CBytes     uint8
	FFixedDisk uint8
	NErrCode   uint16
	Reserved1  uint16
	Reserved2  uint16
	SzPathName [128]CHAR
}

type PFE_EXPORT_FUNC uintptr

type PFE_IMPORT_FUNC uintptr

type LPPROGRESS_ROUTINE uintptr

type COPYFILE2_EXTENDED_PARAMETERS struct {
	DwSize            uint32
	DwCopyFlags       COPYFILE_FLAGS
	PfCancel          *BOOL
	PProgressRoutine  PCOPYFILE2_PROGRESS_ROUTINE
	PvCallbackContext unsafe.Pointer
}

type FILE_ID_DESCRIPTOR struct {
	DwSize uint32
	Type   FILE_ID_TYPE
	FileId int64
}

type CERT_REVOCATION_STATUS_REASON uint32

const (
	CRL_REASON_UNSPECIFIED            CERT_REVOCATION_STATUS_REASON = 0x0
	CRL_REASON_KEY_COMPROMISE         CERT_REVOCATION_STATUS_REASON = 0x1
	CRL_REASON_CA_COMPROMISE          CERT_REVOCATION_STATUS_REASON = 0x2
	CRL_REASON_AFFILIATION_CHANGED    CERT_REVOCATION_STATUS_REASON = 0x3
	CRL_REASON_SUPERSEDED             CERT_REVOCATION_STATUS_REASON = 0x4
	CRL_REASON_CESSATION_OF_OPERATION CERT_REVOCATION_STATUS_REASON = 0x5
	CRL_REASON_CERTIFICATE_HOLD       CERT_REVOCATION_STATUS_REASON = 0x6
	CRL_REASON_REMOVE_FROM_CRL        CERT_REVOCATION_STATUS_REASON = 0x8
)

type CRYPT_XML_KEY_VALUE_TYPE uint32

const (
	CRYPT_XML_KEY_VALUE_TYPE_DSA    CRYPT_XML_KEY_VALUE_TYPE = 0x1
	CRYPT_XML_KEY_VALUE_TYPE_RSA    CRYPT_XML_KEY_VALUE_TYPE = 0x2
	CRYPT_XML_KEY_VALUE_TYPE_ECDSA  CRYPT_XML_KEY_VALUE_TYPE = 0x3
	CRYPT_XML_KEY_VALUE_TYPE_CUSTOM CRYPT_XML_KEY_VALUE_TYPE = 0x4
)

type CRYPT_CONTEXT_CONFIG_FLAGS uint32

const (
	CRYPT_EXCLUSIVE CRYPT_CONTEXT_CONFIG_FLAGS = 0x1
	CRYPT_OVERRIDE  CRYPT_CONTEXT_CONFIG_FLAGS = 0x10000
)

type CERT_CHAIN_POLICY_FLAGS uint32

const (
	CERT_CHAIN_POLICY_IGNORE_NOT_TIME_VALID_FLAG                 CERT_CHAIN_POLICY_FLAGS = 0x1
	CERT_CHAIN_POLICY_IGNORE_CTL_NOT_TIME_VALID_FLAG             CERT_CHAIN_POLICY_FLAGS = 0x2
	CERT_CHAIN_POLICY_IGNORE_NOT_TIME_NESTED_FLAG                CERT_CHAIN_POLICY_FLAGS = 0x4
	CERT_CHAIN_POLICY_IGNORE_INVALID_BASIC_CONSTRAINTS_FLAG      CERT_CHAIN_POLICY_FLAGS = 0x8
	CERT_CHAIN_POLICY_IGNORE_ALL_NOT_TIME_VALID_FLAGS            CERT_CHAIN_POLICY_FLAGS = 0x7
	CERT_CHAIN_POLICY_ALLOW_UNKNOWN_CA_FLAG                      CERT_CHAIN_POLICY_FLAGS = 0x10
	CERT_CHAIN_POLICY_IGNORE_WRONG_USAGE_FLAG                    CERT_CHAIN_POLICY_FLAGS = 0x20
	CERT_CHAIN_POLICY_IGNORE_INVALID_NAME_FLAG                   CERT_CHAIN_POLICY_FLAGS = 0x40
	CERT_CHAIN_POLICY_IGNORE_INVALID_POLICY_FLAG                 CERT_CHAIN_POLICY_FLAGS = 0x80
	CERT_CHAIN_POLICY_IGNORE_END_REV_UNKNOWN_FLAG                CERT_CHAIN_POLICY_FLAGS = 0x100
	CERT_CHAIN_POLICY_IGNORE_CTL_SIGNER_REV_UNKNOWN_FLAG         CERT_CHAIN_POLICY_FLAGS = 0x200
	CERT_CHAIN_POLICY_IGNORE_CA_REV_UNKNOWN_FLAG                 CERT_CHAIN_POLICY_FLAGS = 0x400
	CERT_CHAIN_POLICY_IGNORE_ROOT_REV_UNKNOWN_FLAG               CERT_CHAIN_POLICY_FLAGS = 0x800
	CERT_CHAIN_POLICY_IGNORE_ALL_REV_UNKNOWN_FLAGS               CERT_CHAIN_POLICY_FLAGS = 0xf00
	CERT_CHAIN_POLICY_ALLOW_TESTROOT_FLAG                        CERT_CHAIN_POLICY_FLAGS = 0x8000
	CERT_CHAIN_POLICY_TRUST_TESTROOT_FLAG                        CERT_CHAIN_POLICY_FLAGS = 0x4000
	CERT_CHAIN_POLICY_IGNORE_NOT_SUPPORTED_CRITICAL_EXT_FLAG     CERT_CHAIN_POLICY_FLAGS = 0x2000
	CERT_CHAIN_POLICY_IGNORE_PEER_TRUST_FLAG                     CERT_CHAIN_POLICY_FLAGS = 0x1000
	CERT_CHAIN_POLICY_IGNORE_WEAK_SIGNATURE_FLAG                 CERT_CHAIN_POLICY_FLAGS = 0x8000000
	BASIC_CONSTRAINTS_CERT_CHAIN_POLICY_CA_FLAG                  CERT_CHAIN_POLICY_FLAGS = 0x80000000
	BASIC_CONSTRAINTS_CERT_CHAIN_POLICY_END_ENTITY_FLAG          CERT_CHAIN_POLICY_FLAGS = 0x40000000
	MICROSOFT_ROOT_CERT_CHAIN_POLICY_ENABLE_TEST_ROOT_FLAG       CERT_CHAIN_POLICY_FLAGS = 0x10000
	MICROSOFT_ROOT_CERT_CHAIN_POLICY_CHECK_APPLICATION_ROOT_FLAG CERT_CHAIN_POLICY_FLAGS = 0x20000
	MICROSOFT_ROOT_CERT_CHAIN_POLICY_DISABLE_FLIGHT_ROOT_FLAG    CERT_CHAIN_POLICY_FLAGS = 0x40000
)

type CRYPT_XML_STATUS_INFO_STATUS uint32

const (
	CRYPT_XML_STATUS_INTERNAL_REFERENCE CRYPT_XML_STATUS_INFO_STATUS = 0x1
	CRYPT_XML_STATUS_KEY_AVAILABLE      CRYPT_XML_STATUS_INFO_STATUS = 0x2
	CRYPT_XML_STATUS_DIGESTING          CRYPT_XML_STATUS_INFO_STATUS = 0x4
	CRYPT_XML_STATUS_DIGEST_VALID       CRYPT_XML_STATUS_INFO_STATUS = 0x8
	CRYPT_XML_STATUS_SIGNATURE_VALID    CRYPT_XML_STATUS_INFO_STATUS = 0x10000
	CRYPT_XML_STATUS_OPENED_TO_ENCODE   CRYPT_XML_STATUS_INFO_STATUS = 0x80000000
)

type CRYPT_XML_GROUP_ID uint32

const (
	CRYPT_XML_GROUP_ID_HASH CRYPT_XML_GROUP_ID = 0x1
	CRYPT_XML_GROUP_ID_SIGN CRYPT_XML_GROUP_ID = 0x2
)

type CERT_SELECT_CRITERIA_TYPE uint32

const (
	CERT_SELECT_BY_ENHKEY_USAGE      CERT_SELECT_CRITERIA_TYPE = 0x1
	CERT_SELECT_BY_KEY_USAGE         CERT_SELECT_CRITERIA_TYPE = 0x2
	CERT_SELECT_BY_POLICY_OID        CERT_SELECT_CRITERIA_TYPE = 0x3
	CERT_SELECT_BY_PROV_NAME         CERT_SELECT_CRITERIA_TYPE = 0x4
	CERT_SELECT_BY_EXTENSION         CERT_SELECT_CRITERIA_TYPE = 0x5
	CERT_SELECT_BY_SUBJECT_HOST_NAME CERT_SELECT_CRITERIA_TYPE = 0x6
	CERT_SELECT_BY_ISSUER_ATTR       CERT_SELECT_CRITERIA_TYPE = 0x7
	CERT_SELECT_BY_SUBJECT_ATTR      CERT_SELECT_CRITERIA_TYPE = 0x8
	CERT_SELECT_BY_ISSUER_NAME       CERT_SELECT_CRITERIA_TYPE = 0x9
	CERT_SELECT_BY_PUBLIC_KEY        CERT_SELECT_CRITERIA_TYPE = 0xa
	CERT_SELECT_BY_TLS_SIGNATURES    CERT_SELECT_CRITERIA_TYPE = 0xb
)

type CRYPT_XML_STATUS_ERROR_STATUS uint32

const (
	CRYPT_XML_STATUS_ERROR_NOT_RESOLVED            CRYPT_XML_STATUS_ERROR_STATUS = 0x1
	CRYPT_XML_STATUS_ERROR_DIGEST_INVALID          CRYPT_XML_STATUS_ERROR_STATUS = 0x2
	CRYPT_XML_STATUS_ERROR_NOT_SUPPORTED_ALGORITHM CRYPT_XML_STATUS_ERROR_STATUS = 0x5
	CRYPT_XML_STATUS_ERROR_NOT_SUPPORTED_TRANSFORM CRYPT_XML_STATUS_ERROR_STATUS = 0x8
	CRYPT_XML_STATUS_ERROR_SIGNATURE_INVALID       CRYPT_XML_STATUS_ERROR_STATUS = 0x10000
	CRYPT_XML_STATUS_ERROR_KEYINFO_NOT_PARSED      CRYPT_XML_STATUS_ERROR_STATUS = 0x20000
)

type NCRYPT_ALGORITHM_NAME_CLASS uint32

const (
	NCRYPT_ASYMMETRIC_ENCRYPTION_INTERFACE NCRYPT_ALGORITHM_NAME_CLASS = 0x3
	NCRYPT_SECRET_AGREEMENT_INTERFACE      NCRYPT_ALGORITHM_NAME_CLASS = 0x4
	NCRYPT_SIGNATURE_INTERFACE             NCRYPT_ALGORITHM_NAME_CLASS = 0x5
)

type SIGNER_SUBJECT_CHOICE uint32

const (
	SIGNER_SUBJECT_FILE SIGNER_SUBJECT_CHOICE = 0x1
	SIGNER_SUBJECT_BLOB SIGNER_SUBJECT_CHOICE = 0x2
)

type SIGNER_SIGNATURE_ATTRIBUTE_CHOICE uint32

const (
	SIGNER_NO_ATTR       SIGNER_SIGNATURE_ATTRIBUTE_CHOICE = 0x0
	SIGNER_AUTHCODE_ATTR SIGNER_SIGNATURE_ATTRIBUTE_CHOICE = 0x1
)

type SIGNER_PRIVATE_KEY_CHOICE uint32

const (
	PVK_TYPE_FILE_NAME    SIGNER_PRIVATE_KEY_CHOICE = 0x1
	PVK_TYPE_KEYCONTAINER SIGNER_PRIVATE_KEY_CHOICE = 0x2
)

type SIGNER_CERT_CHOICE uint32

const (
	SIGNER_CERT_SPC_FILE  SIGNER_CERT_CHOICE = 0x1
	SIGNER_CERT_STORE     SIGNER_CERT_CHOICE = 0x2
	SIGNER_CERT_SPC_CHAIN SIGNER_CERT_CHOICE = 0x3
)

type SslGetCipherSuitePRFHashAlgorithmFn uintptr

type IMAGE_SECTION_CHARACTERISTICS uint32

const (
	IMAGE_SCN_TYPE_NO_PAD            IMAGE_SECTION_CHARACTERISTICS = 0x8
	IMAGE_SCN_CNT_CODE               IMAGE_SECTION_CHARACTERISTICS = 0x20
	IMAGE_SCN_CNT_INITIALIZED_DATA   IMAGE_SECTION_CHARACTERISTICS = 0x40
	IMAGE_SCN_CNT_UNINITIALIZED_DATA IMAGE_SECTION_CHARACTERISTICS = 0x80
	IMAGE_SCN_LNK_OTHER              IMAGE_SECTION_CHARACTERISTICS = 0x100
	IMAGE_SCN_LNK_INFO               IMAGE_SECTION_CHARACTERISTICS = 0x200
	IMAGE_SCN_LNK_REMOVE             IMAGE_SECTION_CHARACTERISTICS = 0x800
	IMAGE_SCN_LNK_COMDAT             IMAGE_SECTION_CHARACTERISTICS = 0x1000
	IMAGE_SCN_NO_DEFER_SPEC_EXC      IMAGE_SECTION_CHARACTERISTICS = 0x4000
	IMAGE_SCN_GPREL                  IMAGE_SECTION_CHARACTERISTICS = 0x8000
	IMAGE_SCN_MEM_FARDATA            IMAGE_SECTION_CHARACTERISTICS = 0x8000
	IMAGE_SCN_MEM_PURGEABLE          IMAGE_SECTION_CHARACTERISTICS = 0x20000
	IMAGE_SCN_MEM_16BIT              IMAGE_SECTION_CHARACTERISTICS = 0x20000
	IMAGE_SCN_MEM_LOCKED             IMAGE_SECTION_CHARACTERISTICS = 0x40000
	IMAGE_SCN_MEM_PRELOAD            IMAGE_SECTION_CHARACTERISTICS = 0x80000
	IMAGE_SCN_ALIGN_1BYTES           IMAGE_SECTION_CHARACTERISTICS = 0x100000
	IMAGE_SCN_ALIGN_2BYTES           IMAGE_SECTION_CHARACTERISTICS = 0x200000
	IMAGE_SCN_ALIGN_4BYTES           IMAGE_SECTION_CHARACTERISTICS = 0x300000
	IMAGE_SCN_ALIGN_8BYTES           IMAGE_SECTION_CHARACTERISTICS = 0x400000
	IMAGE_SCN_ALIGN_16BYTES          IMAGE_SECTION_CHARACTERISTICS = 0x500000
	IMAGE_SCN_ALIGN_32BYTES          IMAGE_SECTION_CHARACTERISTICS = 0x600000
	IMAGE_SCN_ALIGN_64BYTES          IMAGE_SECTION_CHARACTERISTICS = 0x700000
	IMAGE_SCN_ALIGN_128BYTES         IMAGE_SECTION_CHARACTERISTICS = 0x800000
	IMAGE_SCN_ALIGN_256BYTES         IMAGE_SECTION_CHARACTERISTICS = 0x900000
	IMAGE_SCN_ALIGN_512BYTES         IMAGE_SECTION_CHARACTERISTICS = 0xa00000
	IMAGE_SCN_ALIGN_1024BYTES        IMAGE_SECTION_CHARACTERISTICS = 0xb00000
	IMAGE_SCN_ALIGN_2048BYTES        IMAGE_SECTION_CHARACTERISTICS = 0xc00000
	IMAGE_SCN_ALIGN_4096BYTES        IMAGE_SECTION_CHARACTERISTICS = 0xd00000
	IMAGE_SCN_ALIGN_8192BYTES        IMAGE_SECTION_CHARACTERISTICS = 0xe00000
	IMAGE_SCN_ALIGN_MASK             IMAGE_SECTION_CHARACTERISTICS = 0xf00000
	IMAGE_SCN_LNK_NRELOC_OVFL        IMAGE_SECTION_CHARACTERISTICS = 0x1000000
	IMAGE_SCN_MEM_DISCARDABLE        IMAGE_SECTION_CHARACTERISTICS = 0x2000000
	IMAGE_SCN_MEM_NOT_CACHED         IMAGE_SECTION_CHARACTERISTICS = 0x4000000
	IMAGE_SCN_MEM_NOT_PAGED          IMAGE_SECTION_CHARACTERISTICS = 0x8000000
	IMAGE_SCN_MEM_SHARED             IMAGE_SECTION_CHARACTERISTICS = 0x10000000
	IMAGE_SCN_MEM_EXECUTE            IMAGE_SECTION_CHARACTERISTICS = 0x20000000
	IMAGE_SCN_MEM_READ               IMAGE_SECTION_CHARACTERISTICS = 0x40000000
	IMAGE_SCN_MEM_WRITE              IMAGE_SECTION_CHARACTERISTICS = 0x80000000
	IMAGE_SCN_SCALE_INDEX            IMAGE_SECTION_CHARACTERISTICS = 0x1
)

type IMAGE_FILE_CHARACTERISTICS2 uint32

const (
	IMAGE_FILE_RELOCS_STRIPPED2         IMAGE_FILE_CHARACTERISTICS2 = 0x1
	IMAGE_FILE_EXECUTABLE_IMAGE2        IMAGE_FILE_CHARACTERISTICS2 = 0x2
	IMAGE_FILE_LINE_NUMS_STRIPPED2      IMAGE_FILE_CHARACTERISTICS2 = 0x4
	IMAGE_FILE_LOCAL_SYMS_STRIPPED2     IMAGE_FILE_CHARACTERISTICS2 = 0x8
	IMAGE_FILE_AGGRESIVE_WS_TRIM2       IMAGE_FILE_CHARACTERISTICS2 = 0x10
	IMAGE_FILE_LARGE_ADDRESS_AWARE2     IMAGE_FILE_CHARACTERISTICS2 = 0x20
	IMAGE_FILE_BYTES_REVERSED_LO2       IMAGE_FILE_CHARACTERISTICS2 = 0x80
	IMAGE_FILE_32BIT_MACHINE2           IMAGE_FILE_CHARACTERISTICS2 = 0x100
	IMAGE_FILE_DEBUG_STRIPPED2          IMAGE_FILE_CHARACTERISTICS2 = 0x200
	IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP2 IMAGE_FILE_CHARACTERISTICS2 = 0x400
	IMAGE_FILE_NET_RUN_FROM_SWAP2       IMAGE_FILE_CHARACTERISTICS2 = 0x800
	IMAGE_FILE_SYSTEM_2                 IMAGE_FILE_CHARACTERISTICS2 = 0x1000
	IMAGE_FILE_DLL_2                    IMAGE_FILE_CHARACTERISTICS2 = 0x2000
	IMAGE_FILE_UP_SYSTEM_ONLY_2         IMAGE_FILE_CHARACTERISTICS2 = 0x4000
	IMAGE_FILE_BYTES_REVERSED_HI_2      IMAGE_FILE_CHARACTERISTICS2 = 0x8000
)

type SYMBOL_INFO_FLAGS uint32

const (
	SYMFLAG_CLR_TOKEN    SYMBOL_INFO_FLAGS = 0x40000
	SYMFLAG_CONSTANT     SYMBOL_INFO_FLAGS = 0x100
	SYMFLAG_EXPORT       SYMBOL_INFO_FLAGS = 0x200
	SYMFLAG_FORWARDER    SYMBOL_INFO_FLAGS = 0x400
	SYMFLAG_FRAMEREL     SYMBOL_INFO_FLAGS = 0x20
	SYMFLAG_FUNCTION     SYMBOL_INFO_FLAGS = 0x800
	SYMFLAG_ILREL        SYMBOL_INFO_FLAGS = 0x10000
	SYMFLAG_LOCAL        SYMBOL_INFO_FLAGS = 0x80
	SYMFLAG_METADATA     SYMBOL_INFO_FLAGS = 0x20000
	SYMFLAG_PARAMETER    SYMBOL_INFO_FLAGS = 0x40
	SYMFLAG_REGISTER     SYMBOL_INFO_FLAGS = 0x8
	SYMFLAG_REGREL       SYMBOL_INFO_FLAGS = 0x10
	SYMFLAG_SLOT         SYMBOL_INFO_FLAGS = 0x8000
	SYMFLAG_THUNK        SYMBOL_INFO_FLAGS = 0x2000
	SYMFLAG_TLSREL       SYMBOL_INFO_FLAGS = 0x4000
	SYMFLAG_VALUEPRESENT SYMBOL_INFO_FLAGS = 0x1
	SYMFLAG_VIRTUAL      SYMBOL_INFO_FLAGS = 0x1000
)

type IMAGEHLP_GET_TYPE_INFO_FLAGS uint32

const (
	IMAGEHLP_GET_TYPE_INFO_CHILDREN IMAGEHLP_GET_TYPE_INFO_FLAGS = 0x2
	IMAGEHLP_GET_TYPE_INFO_UNCACHED IMAGEHLP_GET_TYPE_INFO_FLAGS = 0x1
)

type DEBUG_EVENT_CODE uint32

const (
	CREATE_PROCESS_DEBUG_EVENT DEBUG_EVENT_CODE = 0x3
	CREATE_THREAD_DEBUG_EVENT  DEBUG_EVENT_CODE = 0x2
	EXCEPTION_DEBUG_EVENT      DEBUG_EVENT_CODE = 0x1
	EXIT_PROCESS_DEBUG_EVENT   DEBUG_EVENT_CODE = 0x5
	EXIT_THREAD_DEBUG_EVENT    DEBUG_EVENT_CODE = 0x4
	LOAD_DLL_DEBUG_EVENT       DEBUG_EVENT_CODE = 0x6
	OUTPUT_DEBUG_STRING_EVENT  DEBUG_EVENT_CODE = 0x8
	RIP_EVENT                  DEBUG_EVENT_CODE = 0x9
	UNLOAD_DLL_DEBUG_EVENT     DEBUG_EVENT_CODE = 0x7
)

type MODLOAD_DATA_TYPE uint32

const (
	DBHHEADER_DEBUGDIRS MODLOAD_DATA_TYPE = 0x1
	DBHHEADER_CVMISC    MODLOAD_DATA_TYPE = 0x2
)

type WOW64_CONTEXT_FLAGS uint32

const (
	WOW64_CONTEXT_X86                 WOW64_CONTEXT_FLAGS = 0x10000
	WOW64_CONTEXT_CONTROL             WOW64_CONTEXT_FLAGS = 0x10001
	WOW64_CONTEXT_INTEGER             WOW64_CONTEXT_FLAGS = 0x10002
	WOW64_CONTEXT_SEGMENTS            WOW64_CONTEXT_FLAGS = 0x10004
	WOW64_CONTEXT_FLOATING_POINT      WOW64_CONTEXT_FLAGS = 0x10008
	WOW64_CONTEXT_DEBUG_REGISTERS     WOW64_CONTEXT_FLAGS = 0x10010
	WOW64_CONTEXT_EXTENDED_REGISTERS  WOW64_CONTEXT_FLAGS = 0x10020
	WOW64_CONTEXT_FULL                WOW64_CONTEXT_FLAGS = 0x10007
	WOW64_CONTEXT_ALL                 WOW64_CONTEXT_FLAGS = 0x1003f
	WOW64_CONTEXT_XSTATE              WOW64_CONTEXT_FLAGS = 0x10040
	WOW64_CONTEXT_EXCEPTION_ACTIVE    WOW64_CONTEXT_FLAGS = 0x8000000
	WOW64_CONTEXT_SERVICE_ACTIVE      WOW64_CONTEXT_FLAGS = 0x10000000
	WOW64_CONTEXT_EXCEPTION_REQUEST   WOW64_CONTEXT_FLAGS = 0x40000000
	WOW64_CONTEXT_EXCEPTION_REPORTING WOW64_CONTEXT_FLAGS = 0x80000000
)

type CHAR int8

type EXCEPTION_DEBUG_INFO struct {
	ExceptionRecord EXCEPTION_RECORD
	DwFirstChance   uint32
}

type CLFS_NODE_ID struct {
	CType  uint32
	CbNode uint32
}

type CLS_CONTAINER_INFORMATION struct {
	FileAttributes       uint32
	CreationTime         uint64
	LastAccessTime       uint64
	LastWriteTime        uint64
	ContainerSize        int64
	FileNameActualLength uint32
	FileNameLength       uint32
	FileName             [256]uint16
	State                uint32
	PhysicalContainerId  uint32
	LogicalContainerId   uint32
}

type CLFS_MGMT_NOTIFICATION_TYPE int32

const (
	ClfsMgmtAdvanceTailNotification    CLFS_MGMT_NOTIFICATION_TYPE = 0x0
	ClfsMgmtLogFullHandlerNotification CLFS_MGMT_NOTIFICATION_TYPE = 0x1
	ClfsMgmtLogUnpinnedNotification    CLFS_MGMT_NOTIFICATION_TYPE = 0x2
	ClfsMgmtLogWriteNotification       CLFS_MGMT_NOTIFICATION_TYPE = 0x3
)

type PLOG_TAIL_ADVANCE_CALLBACK uintptr

type PLOG_FULL_HANDLER_CALLBACK uintptr

type PLOG_UNPINNED_CALLBACK uintptr

type EFS_CERTIFICATE_BLOB struct {
	DwCertEncodingType uint32
	CbData             uint32
	PbData             *uint8
}

type EFS_HASH_BLOB struct {
	CbData uint32
	PbData *uint8
}

type IORING_FEATURE_FLAGS int32

const (
	IORING_FEATURE_FLAGS_NONE           IORING_FEATURE_FLAGS = 0x0
	IORING_FEATURE_UM_EMULATION         IORING_FEATURE_FLAGS = 0x1
	IORING_FEATURE_SET_COMPLETION_EVENT IORING_FEATURE_FLAGS = 0x2
)

type IORING_CREATE_REQUIRED_FLAGS int32

const (
	IORING_CREATE_REQUIRED_FLAGS_NONE IORING_CREATE_REQUIRED_FLAGS = 0x0
)

type IORING_CREATE_ADVISORY_FLAGS int32

const (
	IORING_CREATE_ADVISORY_FLAGS_NONE       IORING_CREATE_ADVISORY_FLAGS = 0x0
	IORING_CREATE_SKIP_BUILDER_PARAM_CHECKS IORING_CREATE_ADVISORY_FLAGS = 0x1
)

type IORING_REF_KIND int32

const (
	IORING_REF_RAW        IORING_REF_KIND = 0x0
	IORING_REF_REGISTERED IORING_REF_KIND = 0x1
)

type LIST_ENTRY struct {
	Flink *LIST_ENTRY
	Blink *LIST_ENTRY
}

type BCryptBuffer struct {
	CbBuffer   uint32
	BufferType uint32
	PvBuffer   unsafe.Pointer
}

type BCRYPT_INTERFACE_VERSION struct {
	MajorVersion uint16
	MinorVersion uint16
}

type CRYPT_IMAGE_REG struct {
	PszImage      *PWSTRElement
	CInterfaces   uint32
	RgpInterfaces **CRYPT_INTERFACE_REG
}

type CRYPT_PROVIDER_REF struct {
	DwInterface   uint32
	PszFunction   *PWSTRElement
	PszProvider   *PWSTRElement
	CProperties   uint32
	RgpProperties **CRYPT_PROPERTY_REF
	PUM           *CRYPT_IMAGE_REF
	PKM           *CRYPT_IMAGE_REF
}

type PFN_NCRYPT_ALLOC uintptr

type PFN_NCRYPT_FREE uintptr

type CRYPT_BIT_BLOB struct {
	CbData      uint32
	PbData      *uint8
	CUnusedBits uint32
}

type PCRYPT_DECRYPT_PRIVATE_KEY_FUNC uintptr

type PCRYPT_RESOLVE_HCRYPTPROV_FUNC uintptr

type PFN_CRYPT_ALLOC uintptr

type PFN_CRYPT_FREE uintptr

type PFN_CMSG_STREAM_OUTPUT uintptr

type CRYPT_KEY_PROV_PARAM struct {
	DwParam uint32
	PbData  *uint8
	CbData  uint32
	DwFlags uint32
}

type PFN_CERT_CREATE_CONTEXT_SORT_FUNC uintptr

type PFN_CRYPT_GET_SIGNER_CERTIFICATE uintptr

type CRYPTNET_URL_CACHE_PRE_FETCH_INFO struct {
	CbSize         uint32
	DwObjectType   uint32
	DwError        uint32
	DwReserved     uint32
	ThisUpdateTime FILETIME
	NextUpdateTime FILETIME
	PublishTime    FILETIME
}

type CRYPTNET_URL_CACHE_FLUSH_INFO struct {
	CbSize          uint32
	DwExemptSeconds uint32
	ExpireTime      FILETIME
}

type CRYPTNET_URL_CACHE_RESPONSE_INFO struct {
	CbSize           uint32
	WResponseType    uint16
	WResponseFlags   uint16
	LastModifiedTime FILETIME
	DwMaxAge         uint32
	PwszETag         *PWSTRElement
	DwProxyId        uint32
}

type CERT_TRUST_STATUS struct {
	DwErrorStatus uint32
	DwInfoStatus  uint32
}

type CERT_SIMPLE_CHAIN struct {
	CbSize                      uint32
	TrustStatus                 CERT_TRUST_STATUS
	CElement                    uint32
	RgpElement                  **CERT_CHAIN_ELEMENT
	PTrustListInfo              *CERT_TRUST_LIST_INFO
	FHasRevocationFreshnessTime BOOL
	DwRevocationFreshnessTime   uint32
}

type CERT_USAGE_MATCH struct {
	DwType uint32
	Usage  CTL_USAGE
}

type PFN_CERT_SERVER_OCSP_RESPONSE_UPDATE_CALLBACK uintptr

type CRYPT_TIMESTAMP_INFO struct {
	DwVersion      uint32
	PszTSAPolicyId *PSTRElement
	HashAlgorithm  CRYPT_ALGORITHM_IDENTIFIER
	HashedMessage  CRYPT_INTEGER_BLOB
	SerialNumber   CRYPT_INTEGER_BLOB
	FtTime         FILETIME
	PvAccuracy     *CRYPT_TIMESTAMP_ACCURACY
	FOrdering      BOOL
	Nonce          CRYPT_INTEGER_BLOB
	Tsa            CRYPT_INTEGER_BLOB
	CExtension     uint32
	RgExtension    *CERT_EXTENSION
}

type PFNCryptStreamOutputCallback uintptr

type PFNCryptStreamOutputCallbackEx uintptr

type PFN_AUTHENTICODE_DIGEST_SIGN uintptr

type SIGNER_ATTR_AUTHCODE struct {
	CbSize      uint32
	FCommercial BOOL
	FIndividual BOOL
	PwszName    *PWSTRElement
	PwszInfo    *PWSTRElement
}

type SIGNER_FILE_INFO struct {
	CbSize       uint32
	PwszFileName *PWSTRElement
	HFile        HANDLE
}

type CRYPT_XML_PROPERTY_ID int32

const (
	CRYPT_XML_PROPERTY_MAX_HEAP_SIZE      CRYPT_XML_PROPERTY_ID = 0x1
	CRYPT_XML_PROPERTY_SIGNATURE_LOCATION CRYPT_XML_PROPERTY_ID = 0x2
	CRYPT_XML_PROPERTY_MAX_SIGNATURES     CRYPT_XML_PROPERTY_ID = 0x3
	CRYPT_XML_PROPERTY_DOC_DECLARATION    CRYPT_XML_PROPERTY_ID = 0x4
	CRYPT_XML_PROPERTY_XML_OUTPUT_CHARSET CRYPT_XML_PROPERTY_ID = 0x5
)

type PFN_CRYPT_XML_DATA_PROVIDER_READ uintptr

type PFN_CRYPT_XML_DATA_PROVIDER_CLOSE uintptr

type CRYPT_XML_TRANSFORM_INFO struct {
	CbSize             uint32
	WszAlgorithm       *PWSTRElement
	CbBufferSize       uint32
	DwFlags            CRYPT_XML_TRANSFORM_FLAGS
	PfnCreateTransform PFN_CRYPT_XML_CREATE_TRANSFORM
}

type CRYPT_XML_KEY_DSA_KEY_VALUE struct {
	P       CRYPT_XML_DATA_BLOB
	Q       CRYPT_XML_DATA_BLOB
	G       CRYPT_XML_DATA_BLOB
	Y       CRYPT_XML_DATA_BLOB
	J       CRYPT_XML_DATA_BLOB
	Seed    CRYPT_XML_DATA_BLOB
	Counter CRYPT_XML_DATA_BLOB
}

type CRYPT_XML_KEY_INFO struct {
	CbSize     uint32
	WszId      *PWSTRElement
	CKeyInfo   uint32
	RgKeyInfo  *CRYPT_XML_KEY_INFO_ITEM
	HVerifyKey BCRYPT_KEY_HANDLE
}

type CRYPT_XML_REFERENCES struct {
	CReference   uint32
	RgpReference **CRYPT_XML_REFERENCE
}

type CRYPT_XML_SIGNED_INFO struct {
	CbSize           uint32
	WszId            *PWSTRElement
	Canonicalization CRYPT_XML_ALGORITHM
	SignatureMethod  CRYPT_XML_ALGORITHM
	CReference       uint32
	RgpReference     **CRYPT_XML_REFERENCE
	Encoded          CRYPT_XML_BLOB
}

type HandleType int32

const (
	Asymmetric HandleType = 0x1
	Symmetric  HandleType = 0x2
	Transform  HandleType = 0x3
	Hash       HandleType = 0x4
)

type BCryptOpenAlgorithmProviderFn uintptr

type BCryptGetPropertyFn uintptr

type BCryptSetPropertyFn uintptr

type BCryptCloseAlgorithmProviderFn uintptr

type BCryptGenerateSymmetricKeyFn uintptr

type BCryptGenerateKeyPairFn uintptr

type BCryptEncryptFn uintptr

type BCryptDecryptFn uintptr

type BCryptExportKeyFn uintptr

type BCryptImportKeyFn uintptr

type BCryptImportKeyPairFn uintptr

type BCryptDuplicateKeyFn uintptr

type BCryptFinalizeKeyPairFn uintptr

type BCryptDestroyKeyFn uintptr

type BCryptDestroySecretFn uintptr

type BCryptSignHashFn uintptr

type BCryptVerifySignatureFn uintptr

type BCryptSecretAgreementFn uintptr

type BCryptDeriveKeyFn uintptr

type BCryptKeyDerivationFn uintptr

type BCryptCreateHashFn uintptr

type BCryptHashDataFn uintptr

type BCryptFinishHashFn uintptr

type BCryptCreateMultiHashFn uintptr

type BCryptProcessMultiOperationsFn uintptr

type BCryptDuplicateHashFn uintptr

type BCryptDestroyHashFn uintptr

type BCryptGenRandomFn uintptr

type NCryptOpenStorageProviderFn uintptr

type NCryptEnumAlgorithmsFn uintptr

type NCryptIsAlgSupportedFn uintptr

type NCryptEnumKeysFn uintptr

type NCryptFreeBufferFn uintptr

type NCryptOpenKeyFn uintptr

type NCryptCreatePersistedKeyFn uintptr

type NCryptGetProviderPropertyFn uintptr

type NCryptGetKeyPropertyFn uintptr

type NCryptSetProviderPropertyFn uintptr

type NCryptSetKeyPropertyFn uintptr

type NCryptFinalizeKeyFn uintptr

type NCryptEncryptFn uintptr

type NCryptDecryptFn uintptr

type NCryptImportKeyFn uintptr

type NCryptExportKeyFn uintptr

type NCryptSignHashFn uintptr

type NCryptVerifySignatureFn uintptr

type NCryptDeleteKeyFn uintptr

type NCryptFreeProviderFn uintptr

type NCryptFreeKeyFn uintptr

type NCryptFreeSecretFn uintptr

type NCryptPromptUserFn uintptr

type NCryptNotifyChangeKeyFn uintptr

type NCryptSecretAgreementFn uintptr

type NCryptDeriveKeyFn uintptr

type NCryptKeyDerivationFn uintptr

type NCryptCreateClaimFn uintptr

type NCryptVerifyClaimFn uintptr

type SslComputeClientAuthHashFn uintptr

type SslComputeEapKeyBlockFn uintptr

type SslComputeFinishedHashFn uintptr

type SslCreateEphemeralKeyFn uintptr

type SslCreateHandshakeHashFn uintptr

type SslDecryptPacketFn uintptr

type SslEncryptPacketFn uintptr

type SslEnumCipherSuitesFn uintptr

type SslEnumCipherSuitesExFn uintptr

type SslEnumEccCurvesFn uintptr

type SslExportKeyFn uintptr

type SslFreeBufferFn uintptr

type SslFreeObjectFn uintptr

type SslGenerateMasterKeyFn uintptr

type SslGenerateSessionKeysFn uintptr

type SslGetKeyPropertyFn uintptr

type SslGetProviderPropertyFn uintptr

type SslHashHandshakeFn uintptr

type SslImportKeyFn uintptr

type SslImportMasterKeyFn uintptr

type SslLookupCipherSuiteInfoFn uintptr

type SslOpenPrivateKeyFn uintptr

type SslOpenProviderFn uintptr

type SslSignHashFn uintptr

type SslVerifySignatureFn uintptr

type SslLookupCipherLengthsFn uintptr

type SslCreateClientAuthHashFn uintptr

type SslComputeSessionHashFn uintptr

type SslGeneratePreMasterKeyFn uintptr

type SslExportKeyingMaterialFn uintptr

type SslExtractEarlyKeyFn uintptr

type SslExtractHandshakeKeyFn uintptr

type SslExtractMasterKeyFn uintptr

type SslExpandTrafficKeysFn uintptr

type SslExpandWriteKeyFn uintptr

type SslExpandExporterMasterKeyFn uintptr

type SslExpandResumptionMasterKeyFn uintptr

type SslDuplicateTranscriptHashFn uintptr

type SslExpandBinderKeyFn uintptr

type SslExpandPreSharedKeyFn uintptr

type M128A struct {
	Low  uint64
	High int64
}

type ARM64_NT_NEON128 struct {
	Low  uint64
	High int64
}

type WOW64_FLOATING_SAVE_AREA struct {
	ControlWord   uint32
	StatusWord    uint32
	TagWord       uint32
	ErrorOffset   uint32
	ErrorSelector uint32
	DataOffset    uint32
	DataSelector  uint32
	RegisterArea  [80]uint8
	Cr0NpxState   uint32
}

type IMAGE_FILE_HEADER struct {
	Machine              IMAGE_FILE_MACHINE
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      IMAGE_FILE_CHARACTERISTICS
}

type IMAGE_OPTIONAL_HEADER32 struct {
	Magic                       IMAGE_OPTIONAL_HEADER_MAGIC
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   IMAGE_SUBSYSTEM
	DllCharacteristics          IMAGE_DLL_CHARACTERISTICS
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type IMAGE_OPTIONAL_HEADER64 struct {
	Magic                       IMAGE_OPTIONAL_HEADER_MAGIC
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   IMAGE_SUBSYSTEM
	DllCharacteristics          IMAGE_DLL_CHARACTERISTICS
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type IMAGE_LOAD_CONFIG_CODE_INTEGRITY struct {
	Flags         uint16
	Catalog       uint16
	CatalogOffset uint32
	Reserved      uint32
}

type SID struct {
	Revision            uint8
	SubAuthorityCount   uint8
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
	SubAuthority        [1]uint32
}

type WCT_OBJECT_TYPE int32

const (
	WctCriticalSectionType WCT_OBJECT_TYPE = 0x1
	WctSendMessageType     WCT_OBJECT_TYPE = 0x2
	WctMutexType           WCT_OBJECT_TYPE = 0x3
	WctAlpcType            WCT_OBJECT_TYPE = 0x4
	WctComType             WCT_OBJECT_TYPE = 0x5
	WctThreadWaitType      WCT_OBJECT_TYPE = 0x6
	WctProcessWaitType     WCT_OBJECT_TYPE = 0x7
	WctThreadType          WCT_OBJECT_TYPE = 0x8
	WctComActivationType   WCT_OBJECT_TYPE = 0x9
	WctUnknownType         WCT_OBJECT_TYPE = 0xa
	WctSocketIoType        WCT_OBJECT_TYPE = 0xb
	WctSmbIoType           WCT_OBJECT_TYPE = 0xc
	WctMaxType             WCT_OBJECT_TYPE = 0xd
)

type WCT_OBJECT_STATUS int32

const (
	WctStatusNoAccess     WCT_OBJECT_STATUS = 0x1
	WctStatusRunning      WCT_OBJECT_STATUS = 0x2
	WctStatusBlocked      WCT_OBJECT_STATUS = 0x3
	WctStatusPidOnly      WCT_OBJECT_STATUS = 0x4
	WctStatusPidOnlyRpcss WCT_OBJECT_STATUS = 0x5
	WctStatusOwned        WCT_OBJECT_STATUS = 0x6
	WctStatusNotOwned     WCT_OBJECT_STATUS = 0x7
	WctStatusAbandoned    WCT_OBJECT_STATUS = 0x8
	WctStatusUnknown      WCT_OBJECT_STATUS = 0x9
	WctStatusError        WCT_OBJECT_STATUS = 0xa
	WctStatusMax          WCT_OBJECT_STATUS = 0xb
)

type MINIDUMP_LOCATION_DESCRIPTOR struct {
	DataSize uint32
	Rva      uint32
}

type MINIDUMP_CALLBACK_ROUTINE uintptr

type ADDRESS64 struct {
	Offset  uint64
	Segment uint16
	Mode    ADDRESS_MODE
}

type KDHELP64 struct {
	Thread                         uint64
	ThCallbackStack                uint32
	ThCallbackBStore               uint32
	NextCallback                   uint32
	FramePointer                   uint32
	KiCallUserMode                 uint64
	KeUserCallbackDispatcher       uint64
	SystemRangeStart               uint64
	KiUserExceptionDispatcher      uint64
	StackBase                      uint64
	StackLimit                     uint64
	BuildVersion                   uint32
	RetpolineStubFunctionTableSize uint32
	RetpolineStubFunctionTable     uint64
	RetpolineStubOffset            uint32
	RetpolineStubSize              uint32
	Reserved0                      [2]uint64
}

type SYM_TYPE int32

const (
	SymNone     SYM_TYPE = 0x0
	SymCoff     SYM_TYPE = 0x1
	SymCv       SYM_TYPE = 0x2
	SymPdb      SYM_TYPE = 0x3
	SymExport   SYM_TYPE = 0x4
	SymDeferred SYM_TYPE = 0x5
	SymSym      SYM_TYPE = 0x6
	SymDia      SYM_TYPE = 0x7
	SymVirtual  SYM_TYPE = 0x8
	NumSymTypes SYM_TYPE = 0x9
)

type PCOPYFILE2_PROGRESS_ROUTINE uintptr

type FILE_ID_TYPE int32

const (
	FileIdType         FILE_ID_TYPE = 0x0
	ObjectIdType       FILE_ID_TYPE = 0x1
	ExtendedFileIdType FILE_ID_TYPE = 0x2
	MaximumFileIdType  FILE_ID_TYPE = 0x3
)

type CRYPT_XML_TRANSFORM_FLAGS uint32

const (
	CRYPT_XML_TRANSFORM_ON_STREAM        CRYPT_XML_TRANSFORM_FLAGS = 0x1
	CRYPT_XML_TRANSFORM_ON_NODESET       CRYPT_XML_TRANSFORM_FLAGS = 0x2
	CRYPT_XML_TRANSFORM_URI_QUERY_STRING CRYPT_XML_TRANSFORM_FLAGS = 0x3
)

type IMAGE_SUBSYSTEM uint16

const (
	IMAGE_SUBSYSTEM_UNKNOWN                  IMAGE_SUBSYSTEM = 0x0
	IMAGE_SUBSYSTEM_NATIVE                   IMAGE_SUBSYSTEM = 0x1
	IMAGE_SUBSYSTEM_WINDOWS_GUI              IMAGE_SUBSYSTEM = 0x2
	IMAGE_SUBSYSTEM_WINDOWS_CUI              IMAGE_SUBSYSTEM = 0x3
	IMAGE_SUBSYSTEM_OS2_CUI                  IMAGE_SUBSYSTEM = 0x5
	IMAGE_SUBSYSTEM_POSIX_CUI                IMAGE_SUBSYSTEM = 0x7
	IMAGE_SUBSYSTEM_NATIVE_WINDOWS           IMAGE_SUBSYSTEM = 0x8
	IMAGE_SUBSYSTEM_WINDOWS_CE_GUI           IMAGE_SUBSYSTEM = 0x9
	IMAGE_SUBSYSTEM_EFI_APPLICATION          IMAGE_SUBSYSTEM = 0xa
	IMAGE_SUBSYSTEM_EFI_BOOT_SERVICE_DRIVER  IMAGE_SUBSYSTEM = 0xb
	IMAGE_SUBSYSTEM_EFI_RUNTIME_DRIVER       IMAGE_SUBSYSTEM = 0xc
	IMAGE_SUBSYSTEM_EFI_ROM                  IMAGE_SUBSYSTEM = 0xd
	IMAGE_SUBSYSTEM_XBOX                     IMAGE_SUBSYSTEM = 0xe
	IMAGE_SUBSYSTEM_WINDOWS_BOOT_APPLICATION IMAGE_SUBSYSTEM = 0x10
	IMAGE_SUBSYSTEM_XBOX_CODE_CATALOG        IMAGE_SUBSYSTEM = 0x11
)

type IMAGE_DLL_CHARACTERISTICS uint16

const (
	IMAGE_DLLCHARACTERISTICS_HIGH_ENTROPY_VA                               IMAGE_DLL_CHARACTERISTICS = 0x20
	IMAGE_DLLCHARACTERISTICS_DYNAMIC_BASE                                  IMAGE_DLL_CHARACTERISTICS = 0x40
	IMAGE_DLLCHARACTERISTICS_FORCE_INTEGRITY                               IMAGE_DLL_CHARACTERISTICS = 0x80
	IMAGE_DLLCHARACTERISTICS_NX_COMPAT                                     IMAGE_DLL_CHARACTERISTICS = 0x100
	IMAGE_DLLCHARACTERISTICS_NO_ISOLATION                                  IMAGE_DLL_CHARACTERISTICS = 0x200
	IMAGE_DLLCHARACTERISTICS_NO_SEH                                        IMAGE_DLL_CHARACTERISTICS = 0x400
	IMAGE_DLLCHARACTERISTICS_NO_BIND                                       IMAGE_DLL_CHARACTERISTICS = 0x800
	IMAGE_DLLCHARACTERISTICS_APPCONTAINER                                  IMAGE_DLL_CHARACTERISTICS = 0x1000
	IMAGE_DLLCHARACTERISTICS_WDM_DRIVER                                    IMAGE_DLL_CHARACTERISTICS = 0x2000
	IMAGE_DLLCHARACTERISTICS_GUARD_CF                                      IMAGE_DLL_CHARACTERISTICS = 0x4000
	IMAGE_DLLCHARACTERISTICS_TERMINAL_SERVER_AWARE                         IMAGE_DLL_CHARACTERISTICS = 0x8000
	IMAGE_DLLCHARACTERISTICS_EX_CET_COMPAT                                 IMAGE_DLL_CHARACTERISTICS = 0x1
	IMAGE_DLLCHARACTERISTICS_EX_CET_COMPAT_STRICT_MODE                     IMAGE_DLL_CHARACTERISTICS = 0x2
	IMAGE_DLLCHARACTERISTICS_EX_CET_SET_CONTEXT_IP_VALIDATION_RELAXED_MODE IMAGE_DLL_CHARACTERISTICS = 0x4
	IMAGE_DLLCHARACTERISTICS_EX_CET_DYNAMIC_APIS_ALLOW_IN_PROC             IMAGE_DLL_CHARACTERISTICS = 0x8
	IMAGE_DLLCHARACTERISTICS_EX_CET_RESERVED_1                             IMAGE_DLL_CHARACTERISTICS = 0x10
	IMAGE_DLLCHARACTERISTICS_EX_CET_RESERVED_2                             IMAGE_DLL_CHARACTERISTICS = 0x20
	IMAGE_DLLCHARACTERISTICS_EX_FORWARD_CFI_COMPAT                         IMAGE_DLL_CHARACTERISTICS = 0x40
	IMAGE_DLLCHARACTERISTICS_EX_HOTPATCH_COMPATIBLE                        IMAGE_DLL_CHARACTERISTICS = 0x80
)

type IMAGE_OPTIONAL_HEADER_MAGIC uint16

const (
	IMAGE_NT_OPTIONAL_HDR_MAGIC   IMAGE_OPTIONAL_HEADER_MAGIC = 0x20b
	IMAGE_NT_OPTIONAL_HDR32_MAGIC IMAGE_OPTIONAL_HEADER_MAGIC = 0x10b
	IMAGE_NT_OPTIONAL_HDR64_MAGIC IMAGE_OPTIONAL_HEADER_MAGIC = 0x20b
	IMAGE_ROM_OPTIONAL_HDR_MAGIC  IMAGE_OPTIONAL_HEADER_MAGIC = 0x107
)

type IMAGE_FILE_CHARACTERISTICS uint16

const (
	IMAGE_FILE_RELOCS_STRIPPED         IMAGE_FILE_CHARACTERISTICS = 0x1
	IMAGE_FILE_EXECUTABLE_IMAGE        IMAGE_FILE_CHARACTERISTICS = 0x2
	IMAGE_FILE_LINE_NUMS_STRIPPED      IMAGE_FILE_CHARACTERISTICS = 0x4
	IMAGE_FILE_LOCAL_SYMS_STRIPPED     IMAGE_FILE_CHARACTERISTICS = 0x8
	IMAGE_FILE_AGGRESIVE_WS_TRIM       IMAGE_FILE_CHARACTERISTICS = 0x10
	IMAGE_FILE_LARGE_ADDRESS_AWARE     IMAGE_FILE_CHARACTERISTICS = 0x20
	IMAGE_FILE_BYTES_REVERSED_LO       IMAGE_FILE_CHARACTERISTICS = 0x80
	IMAGE_FILE_32BIT_MACHINE           IMAGE_FILE_CHARACTERISTICS = 0x100
	IMAGE_FILE_DEBUG_STRIPPED          IMAGE_FILE_CHARACTERISTICS = 0x200
	IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP IMAGE_FILE_CHARACTERISTICS = 0x400
	IMAGE_FILE_NET_RUN_FROM_SWAP       IMAGE_FILE_CHARACTERISTICS = 0x800
	IMAGE_FILE_SYSTEM                  IMAGE_FILE_CHARACTERISTICS = 0x1000
	IMAGE_FILE_DLL                     IMAGE_FILE_CHARACTERISTICS = 0x2000
	IMAGE_FILE_UP_SYSTEM_ONLY          IMAGE_FILE_CHARACTERISTICS = 0x4000
	IMAGE_FILE_BYTES_REVERSED_HI       IMAGE_FILE_CHARACTERISTICS = 0x8000
)

type IMAGE_FILE_MACHINE uint16

const (
	IMAGE_FILE_MACHINE_AXP64       IMAGE_FILE_MACHINE = 0x284
	IMAGE_FILE_MACHINE_I386        IMAGE_FILE_MACHINE = 0x14c
	IMAGE_FILE_MACHINE_IA64        IMAGE_FILE_MACHINE = 0x200
	IMAGE_FILE_MACHINE_AMD64       IMAGE_FILE_MACHINE = 0x8664
	IMAGE_FILE_MACHINE_UNKNOWN     IMAGE_FILE_MACHINE = 0x0
	IMAGE_FILE_MACHINE_TARGET_HOST IMAGE_FILE_MACHINE = 0x1
	IMAGE_FILE_MACHINE_R3000       IMAGE_FILE_MACHINE = 0x162
	IMAGE_FILE_MACHINE_R4000       IMAGE_FILE_MACHINE = 0x166
	IMAGE_FILE_MACHINE_R10000      IMAGE_FILE_MACHINE = 0x168
	IMAGE_FILE_MACHINE_WCEMIPSV2   IMAGE_FILE_MACHINE = 0x169
	IMAGE_FILE_MACHINE_ALPHA       IMAGE_FILE_MACHINE = 0x184
	IMAGE_FILE_MACHINE_SH3         IMAGE_FILE_MACHINE = 0x1a2
	IMAGE_FILE_MACHINE_SH3DSP      IMAGE_FILE_MACHINE = 0x1a3
	IMAGE_FILE_MACHINE_SH3E        IMAGE_FILE_MACHINE = 0x1a4
	IMAGE_FILE_MACHINE_SH4         IMAGE_FILE_MACHINE = 0x1a6
	IMAGE_FILE_MACHINE_SH5         IMAGE_FILE_MACHINE = 0x1a8
	IMAGE_FILE_MACHINE_ARM         IMAGE_FILE_MACHINE = 0x1c0
	IMAGE_FILE_MACHINE_THUMB       IMAGE_FILE_MACHINE = 0x1c2
	IMAGE_FILE_MACHINE_ARMNT       IMAGE_FILE_MACHINE = 0x1c4
	IMAGE_FILE_MACHINE_AM33        IMAGE_FILE_MACHINE = 0x1d3
	IMAGE_FILE_MACHINE_POWERPC     IMAGE_FILE_MACHINE = 0x1f0
	IMAGE_FILE_MACHINE_POWERPCFP   IMAGE_FILE_MACHINE = 0x1f1
	IMAGE_FILE_MACHINE_MIPS16      IMAGE_FILE_MACHINE = 0x266
	IMAGE_FILE_MACHINE_ALPHA64     IMAGE_FILE_MACHINE = 0x284
	IMAGE_FILE_MACHINE_MIPSFPU     IMAGE_FILE_MACHINE = 0x366
	IMAGE_FILE_MACHINE_MIPSFPU16   IMAGE_FILE_MACHINE = 0x466
	IMAGE_FILE_MACHINE_TRICORE     IMAGE_FILE_MACHINE = 0x520
	IMAGE_FILE_MACHINE_CEF         IMAGE_FILE_MACHINE = 0xcef
	IMAGE_FILE_MACHINE_EBC         IMAGE_FILE_MACHINE = 0xebc
	IMAGE_FILE_MACHINE_M32R        IMAGE_FILE_MACHINE = 0x9041
	IMAGE_FILE_MACHINE_ARM64       IMAGE_FILE_MACHINE = 0xaa64
	IMAGE_FILE_MACHINE_CEE         IMAGE_FILE_MACHINE = 0xc0ee
)

type CRYPT_INTERFACE_REG struct {
	DwInterface    BCRYPT_INTERFACE
	DwFlags        BCRYPT_TABLE
	CFunctions     uint32
	RgpszFunctions **PWSTRElement
}

type CRYPT_PROPERTY_REF struct {
	PszProperty *PWSTRElement
	CbValue     uint32
	PbValue     *uint8
}

type CRYPT_IMAGE_REF struct {
	PszImage *PWSTRElement
	DwFlags  CRYPT_IMAGE_REF_FLAGS
}

type CERT_TRUST_LIST_INFO struct {
	CbSize      uint32
	PCtlEntry   *CTL_ENTRY
	PCtlContext *CTL_CONTEXT
}

type CERT_CHAIN_ELEMENT struct {
	CbSize                uint32
	PCertContext          *CERT_CONTEXT
	TrustStatus           CERT_TRUST_STATUS
	PRevocationInfo       *CERT_REVOCATION_INFO
	PIssuanceUsage        *CTL_USAGE
	PApplicationUsage     *CTL_USAGE
	PwszExtendedErrorInfo *PWSTRElement
}

type CRYPT_TIMESTAMP_ACCURACY struct {
	DwSeconds uint32
	DwMillis  uint32
	DwMicros  uint32
}

type CRYPT_XML_DATA_BLOB struct {
	CbData uint32
	PbData *uint8
}

type PFN_CRYPT_XML_CREATE_TRANSFORM uintptr

type CRYPT_XML_KEY_INFO_ITEM struct {
	DwType     CRYPT_XML_KEYINFO_TYPE
	WszKeyName *PWSTRElement
}

type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress uint32
	Size           uint32
}

type SID_IDENTIFIER_AUTHORITY struct {
	Value [6]uint8
}

type ADDRESS_MODE int32

const (
	AddrMode1616 ADDRESS_MODE = 0x0
	AddrMode1632 ADDRESS_MODE = 0x1
	AddrModeReal ADDRESS_MODE = 0x2
	AddrModeFlat ADDRESS_MODE = 0x3
)

type CRYPT_IMAGE_REF_FLAGS uint32

const (
	CRYPT_MIN_DEPENDENCIES CRYPT_IMAGE_REF_FLAGS = 0x1
	CRYPT_PROCESS_ISOLATE  CRYPT_IMAGE_REF_FLAGS = 0x10000
)

type CRYPT_XML_KEYINFO_TYPE uint32

const (
	CRYPT_XML_KEYINFO_TYPE_KEYNAME   CRYPT_XML_KEYINFO_TYPE = 0x1
	CRYPT_XML_KEYINFO_TYPE_KEYVALUE  CRYPT_XML_KEYINFO_TYPE = 0x2
	CRYPT_XML_KEYINFO_TYPE_RETRIEVAL CRYPT_XML_KEYINFO_TYPE = 0x3
	CRYPT_XML_KEYINFO_TYPE_X509DATA  CRYPT_XML_KEYINFO_TYPE = 0x4
	CRYPT_XML_KEYINFO_TYPE_CUSTOM    CRYPT_XML_KEYINFO_TYPE = 0x5
)

type CERT_REVOCATION_INFO struct {
	CbSize             uint32
	DwRevocationResult uint32
	PszRevocationOid   *PSTRElement
	PvOidSpecificInfo  unsafe.Pointer
	FHasFreshnessTime  BOOL
	DwFreshnessTime    uint32
	PCrlInfo           *CERT_REVOCATION_CRL_INFO
}

type CERT_REVOCATION_CRL_INFO struct {
	CbSize           uint32
	PBaseCrlContext  *CRL_CONTEXT
	PDeltaCrlContext *CRL_CONTEXT
	PCrlEntry        *CRL_ENTRY
	FDeltaCrlEntry   BOOL
}

