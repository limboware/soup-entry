//go:build plugin
// +build plugin

package s2sdkv213

import (
	"unsafe"

	"github.com/untrustedmodders/go-plugify"
	"github.com/untrustedmodders/s2sdkv213/s2sdk"
)

var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

//go:linkname s2sdk_Kv1Create __package__/s2sdk._Kv1Create
var s2sdk_Kv1Create func(setName string) uintptr
var S2sdk_Kv1Create = &s2sdk_Kv1Create

//go:linkname s2sdk_Kv1Destroy __package__/s2sdk._Kv1Destroy
var s2sdk_Kv1Destroy func(kv uintptr)
var S2sdk_Kv1Destroy = &s2sdk_Kv1Destroy

//go:linkname s2sdk_Kv1GetName __package__/s2sdk._Kv1GetName
var s2sdk_Kv1GetName func(kv uintptr) string
var S2sdk_Kv1GetName = &s2sdk_Kv1GetName

//go:linkname s2sdk_Kv1SetName __package__/s2sdk._Kv1SetName
var s2sdk_Kv1SetName func(kv uintptr, name string)
var S2sdk_Kv1SetName = &s2sdk_Kv1SetName

//go:linkname s2sdk_Kv1FindKey __package__/s2sdk._Kv1FindKey
var s2sdk_Kv1FindKey func(kv uintptr, keyName string) uintptr
var S2sdk_Kv1FindKey = &s2sdk_Kv1FindKey

//go:linkname s2sdk_Kv1FindOrCreateKey __package__/s2sdk._Kv1FindOrCreateKey
var s2sdk_Kv1FindOrCreateKey func(kv uintptr, keyName string) uintptr
var S2sdk_Kv1FindOrCreateKey = &s2sdk_Kv1FindOrCreateKey

//go:linkname s2sdk_Kv1CreateKey __package__/s2sdk._Kv1CreateKey
var s2sdk_Kv1CreateKey func(kv uintptr, keyName string) uintptr
var S2sdk_Kv1CreateKey = &s2sdk_Kv1CreateKey

//go:linkname s2sdk_Kv1CreateNewKey __package__/s2sdk._Kv1CreateNewKey
var s2sdk_Kv1CreateNewKey func(kv uintptr) uintptr
var S2sdk_Kv1CreateNewKey = &s2sdk_Kv1CreateNewKey

//go:linkname s2sdk_Kv1AddSubKey __package__/s2sdk._Kv1AddSubKey
var s2sdk_Kv1AddSubKey func(kv uintptr, subKey uintptr)
var S2sdk_Kv1AddSubKey = &s2sdk_Kv1AddSubKey

//go:linkname s2sdk_Kv1GetFirstSubKey __package__/s2sdk._Kv1GetFirstSubKey
var s2sdk_Kv1GetFirstSubKey func(kv uintptr) uintptr
var S2sdk_Kv1GetFirstSubKey = &s2sdk_Kv1GetFirstSubKey

//go:linkname s2sdk_Kv1GetNextKey __package__/s2sdk._Kv1GetNextKey
var s2sdk_Kv1GetNextKey func(kv uintptr) uintptr
var S2sdk_Kv1GetNextKey = &s2sdk_Kv1GetNextKey

//go:linkname s2sdk_Kv1GetColor __package__/s2sdk._Kv1GetColor
var s2sdk_Kv1GetColor func(kv uintptr, keyName string, defaultValue plugify.Vector4) plugify.Vector4
var S2sdk_Kv1GetColor = &s2sdk_Kv1GetColor

//go:linkname s2sdk_Kv1SetColor __package__/s2sdk._Kv1SetColor
var s2sdk_Kv1SetColor func(kv uintptr, keyName string, value plugify.Vector4)
var S2sdk_Kv1SetColor = &s2sdk_Kv1SetColor

//go:linkname s2sdk_Kv1GetInt __package__/s2sdk._Kv1GetInt
var s2sdk_Kv1GetInt func(kv uintptr, keyName string, defaultValue int32) int32
var S2sdk_Kv1GetInt = &s2sdk_Kv1GetInt

//go:linkname s2sdk_Kv1SetInt __package__/s2sdk._Kv1SetInt
var s2sdk_Kv1SetInt func(kv uintptr, keyName string, value int32)
var S2sdk_Kv1SetInt = &s2sdk_Kv1SetInt

//go:linkname s2sdk_Kv1GetFloat __package__/s2sdk._Kv1GetFloat
var s2sdk_Kv1GetFloat func(kv uintptr, keyName string, defaultValue float32) float32
var S2sdk_Kv1GetFloat = &s2sdk_Kv1GetFloat

//go:linkname s2sdk_Kv1SetFloat __package__/s2sdk._Kv1SetFloat
var s2sdk_Kv1SetFloat func(kv uintptr, keyName string, value float32)
var S2sdk_Kv1SetFloat = &s2sdk_Kv1SetFloat

//go:linkname s2sdk_Kv1GetString __package__/s2sdk._Kv1GetString
var s2sdk_Kv1GetString func(kv uintptr, keyName string, defaultValue string) string
var S2sdk_Kv1GetString = &s2sdk_Kv1GetString

//go:linkname s2sdk_Kv1SetString __package__/s2sdk._Kv1SetString
var s2sdk_Kv1SetString func(kv uintptr, keyName string, value string)
var S2sdk_Kv1SetString = &s2sdk_Kv1SetString

//go:linkname s2sdk_Kv1GetPtr __package__/s2sdk._Kv1GetPtr
var s2sdk_Kv1GetPtr func(kv uintptr, keyName string, defaultValue uintptr) uintptr
var S2sdk_Kv1GetPtr = &s2sdk_Kv1GetPtr

//go:linkname s2sdk_Kv1SetPtr __package__/s2sdk._Kv1SetPtr
var s2sdk_Kv1SetPtr func(kv uintptr, keyName string, value uintptr)
var S2sdk_Kv1SetPtr = &s2sdk_Kv1SetPtr

//go:linkname s2sdk_Kv1GetBool __package__/s2sdk._Kv1GetBool
var s2sdk_Kv1GetBool func(kv uintptr, keyName string, defaultValue bool) bool
var S2sdk_Kv1GetBool = &s2sdk_Kv1GetBool

//go:linkname s2sdk_Kv1SetBool __package__/s2sdk._Kv1SetBool
var s2sdk_Kv1SetBool func(kv uintptr, keyName string, value bool)
var S2sdk_Kv1SetBool = &s2sdk_Kv1SetBool

//go:linkname s2sdk_Kv1MakeCopy __package__/s2sdk._Kv1MakeCopy
var s2sdk_Kv1MakeCopy func(kv uintptr) uintptr
var S2sdk_Kv1MakeCopy = &s2sdk_Kv1MakeCopy

//go:linkname s2sdk_Kv1Clear __package__/s2sdk._Kv1Clear
var s2sdk_Kv1Clear func(kv uintptr)
var S2sdk_Kv1Clear = &s2sdk_Kv1Clear

//go:linkname s2sdk_Kv1IsEmpty __package__/s2sdk._Kv1IsEmpty
var s2sdk_Kv1IsEmpty func(kv uintptr, keyName string) bool
var S2sdk_Kv1IsEmpty = &s2sdk_Kv1IsEmpty

//go:linkname s2sdk_Kv3Create __package__/s2sdk._Kv3Create
var s2sdk_Kv3Create func(type_ int32, subtype int32) uintptr
var S2sdk_Kv3Create = &s2sdk_Kv3Create

//go:linkname s2sdk_Kv3CreateWithCluster __package__/s2sdk._Kv3CreateWithCluster
var s2sdk_Kv3CreateWithCluster func(cluster_elem int32, type_ int32, subtype int32) uintptr
var S2sdk_Kv3CreateWithCluster = &s2sdk_Kv3CreateWithCluster

//go:linkname s2sdk_Kv3CreateCopy __package__/s2sdk._Kv3CreateCopy
var s2sdk_Kv3CreateCopy func(other uintptr) uintptr
var S2sdk_Kv3CreateCopy = &s2sdk_Kv3CreateCopy

//go:linkname s2sdk_Kv3Destroy __package__/s2sdk._Kv3Destroy
var s2sdk_Kv3Destroy func(kv uintptr)
var S2sdk_Kv3Destroy = &s2sdk_Kv3Destroy

//go:linkname s2sdk_Kv3CopyFrom __package__/s2sdk._Kv3CopyFrom
var s2sdk_Kv3CopyFrom func(kv uintptr, other uintptr)
var S2sdk_Kv3CopyFrom = &s2sdk_Kv3CopyFrom

//go:linkname s2sdk_Kv3OverlayKeysFrom __package__/s2sdk._Kv3OverlayKeysFrom
var s2sdk_Kv3OverlayKeysFrom func(kv uintptr, other uintptr, depth bool)
var S2sdk_Kv3OverlayKeysFrom = &s2sdk_Kv3OverlayKeysFrom

//go:linkname s2sdk_Kv3GetContext __package__/s2sdk._Kv3GetContext
var s2sdk_Kv3GetContext func(kv uintptr) uintptr
var S2sdk_Kv3GetContext = &s2sdk_Kv3GetContext

//go:linkname s2sdk_Kv3GetMetaData __package__/s2sdk._Kv3GetMetaData
var s2sdk_Kv3GetMetaData func(kv uintptr, ppCtx uintptr) uintptr
var S2sdk_Kv3GetMetaData = &s2sdk_Kv3GetMetaData

//go:linkname s2sdk_Kv3HasFlag __package__/s2sdk._Kv3HasFlag
var s2sdk_Kv3HasFlag func(kv uintptr, flag uint8) bool
var S2sdk_Kv3HasFlag = &s2sdk_Kv3HasFlag

//go:linkname s2sdk_Kv3HasAnyFlags __package__/s2sdk._Kv3HasAnyFlags
var s2sdk_Kv3HasAnyFlags func(kv uintptr) bool
var S2sdk_Kv3HasAnyFlags = &s2sdk_Kv3HasAnyFlags

//go:linkname s2sdk_Kv3GetAllFlags __package__/s2sdk._Kv3GetAllFlags
var s2sdk_Kv3GetAllFlags func(kv uintptr) uint8
var S2sdk_Kv3GetAllFlags = &s2sdk_Kv3GetAllFlags

//go:linkname s2sdk_Kv3SetAllFlags __package__/s2sdk._Kv3SetAllFlags
var s2sdk_Kv3SetAllFlags func(kv uintptr, flags uint8)
var S2sdk_Kv3SetAllFlags = &s2sdk_Kv3SetAllFlags

//go:linkname s2sdk_Kv3SetFlag __package__/s2sdk._Kv3SetFlag
var s2sdk_Kv3SetFlag func(kv uintptr, flag uint8, state bool)
var S2sdk_Kv3SetFlag = &s2sdk_Kv3SetFlag

//go:linkname s2sdk_Kv3GetType __package__/s2sdk._Kv3GetType
var s2sdk_Kv3GetType func(kv uintptr) uint8
var S2sdk_Kv3GetType = &s2sdk_Kv3GetType

//go:linkname s2sdk_Kv3GetTypeEx __package__/s2sdk._Kv3GetTypeEx
var s2sdk_Kv3GetTypeEx func(kv uintptr) uint8
var S2sdk_Kv3GetTypeEx = &s2sdk_Kv3GetTypeEx

//go:linkname s2sdk_Kv3GetSubType __package__/s2sdk._Kv3GetSubType
var s2sdk_Kv3GetSubType func(kv uintptr) uint8
var S2sdk_Kv3GetSubType = &s2sdk_Kv3GetSubType

//go:linkname s2sdk_Kv3HasInvalidMemberNames __package__/s2sdk._Kv3HasInvalidMemberNames
var s2sdk_Kv3HasInvalidMemberNames func(kv uintptr) bool
var S2sdk_Kv3HasInvalidMemberNames = &s2sdk_Kv3HasInvalidMemberNames

//go:linkname s2sdk_Kv3SetHasInvalidMemberNames __package__/s2sdk._Kv3SetHasInvalidMemberNames
var s2sdk_Kv3SetHasInvalidMemberNames func(kv uintptr, bValue bool)
var S2sdk_Kv3SetHasInvalidMemberNames = &s2sdk_Kv3SetHasInvalidMemberNames

//go:linkname s2sdk_Kv3GetTypeAsString __package__/s2sdk._Kv3GetTypeAsString
var s2sdk_Kv3GetTypeAsString func(kv uintptr) string
var S2sdk_Kv3GetTypeAsString = &s2sdk_Kv3GetTypeAsString

//go:linkname s2sdk_Kv3GetSubTypeAsString __package__/s2sdk._Kv3GetSubTypeAsString
var s2sdk_Kv3GetSubTypeAsString func(kv uintptr) string
var S2sdk_Kv3GetSubTypeAsString = &s2sdk_Kv3GetSubTypeAsString

//go:linkname s2sdk_Kv3ToString __package__/s2sdk._Kv3ToString
var s2sdk_Kv3ToString func(kv uintptr, flags uint32) string
var S2sdk_Kv3ToString = &s2sdk_Kv3ToString

//go:linkname s2sdk_Kv3IsNull __package__/s2sdk._Kv3IsNull
var s2sdk_Kv3IsNull func(kv uintptr) bool
var S2sdk_Kv3IsNull = &s2sdk_Kv3IsNull

//go:linkname s2sdk_Kv3SetToNull __package__/s2sdk._Kv3SetToNull
var s2sdk_Kv3SetToNull func(kv uintptr)
var S2sdk_Kv3SetToNull = &s2sdk_Kv3SetToNull

//go:linkname s2sdk_Kv3IsArray __package__/s2sdk._Kv3IsArray
var s2sdk_Kv3IsArray func(kv uintptr) bool
var S2sdk_Kv3IsArray = &s2sdk_Kv3IsArray

//go:linkname s2sdk_Kv3IsKV3Array __package__/s2sdk._Kv3IsKV3Array
var s2sdk_Kv3IsKV3Array func(kv uintptr) bool
var S2sdk_Kv3IsKV3Array = &s2sdk_Kv3IsKV3Array

//go:linkname s2sdk_Kv3IsTable __package__/s2sdk._Kv3IsTable
var s2sdk_Kv3IsTable func(kv uintptr) bool
var S2sdk_Kv3IsTable = &s2sdk_Kv3IsTable

//go:linkname s2sdk_Kv3IsString __package__/s2sdk._Kv3IsString
var s2sdk_Kv3IsString func(kv uintptr) bool
var S2sdk_Kv3IsString = &s2sdk_Kv3IsString

//go:linkname s2sdk_Kv3GetBool __package__/s2sdk._Kv3GetBool
var s2sdk_Kv3GetBool func(kv uintptr, defaultValue bool) bool
var S2sdk_Kv3GetBool = &s2sdk_Kv3GetBool

//go:linkname s2sdk_Kv3GetChar __package__/s2sdk._Kv3GetChar
var s2sdk_Kv3GetChar func(kv uintptr, defaultValue int8) int8
var S2sdk_Kv3GetChar = &s2sdk_Kv3GetChar

//go:linkname s2sdk_Kv3GetUChar32 __package__/s2sdk._Kv3GetUChar32
var s2sdk_Kv3GetUChar32 func(kv uintptr, defaultValue uint32) uint32
var S2sdk_Kv3GetUChar32 = &s2sdk_Kv3GetUChar32

//go:linkname s2sdk_Kv3GetInt8 __package__/s2sdk._Kv3GetInt8
var s2sdk_Kv3GetInt8 func(kv uintptr, defaultValue int8) int8
var S2sdk_Kv3GetInt8 = &s2sdk_Kv3GetInt8

//go:linkname s2sdk_Kv3GetUInt8 __package__/s2sdk._Kv3GetUInt8
var s2sdk_Kv3GetUInt8 func(kv uintptr, defaultValue uint8) uint8
var S2sdk_Kv3GetUInt8 = &s2sdk_Kv3GetUInt8

//go:linkname s2sdk_Kv3GetShort __package__/s2sdk._Kv3GetShort
var s2sdk_Kv3GetShort func(kv uintptr, defaultValue int16) int16
var S2sdk_Kv3GetShort = &s2sdk_Kv3GetShort

//go:linkname s2sdk_Kv3GetUShort __package__/s2sdk._Kv3GetUShort
var s2sdk_Kv3GetUShort func(kv uintptr, defaultValue uint16) uint16
var S2sdk_Kv3GetUShort = &s2sdk_Kv3GetUShort

//go:linkname s2sdk_Kv3GetInt __package__/s2sdk._Kv3GetInt
var s2sdk_Kv3GetInt func(kv uintptr, defaultValue int32) int32
var S2sdk_Kv3GetInt = &s2sdk_Kv3GetInt

//go:linkname s2sdk_Kv3GetUInt __package__/s2sdk._Kv3GetUInt
var s2sdk_Kv3GetUInt func(kv uintptr, defaultValue uint32) uint32
var S2sdk_Kv3GetUInt = &s2sdk_Kv3GetUInt

//go:linkname s2sdk_Kv3GetInt64 __package__/s2sdk._Kv3GetInt64
var s2sdk_Kv3GetInt64 func(kv uintptr, defaultValue int64) int64
var S2sdk_Kv3GetInt64 = &s2sdk_Kv3GetInt64

//go:linkname s2sdk_Kv3GetUInt64 __package__/s2sdk._Kv3GetUInt64
var s2sdk_Kv3GetUInt64 func(kv uintptr, defaultValue uint64) uint64
var S2sdk_Kv3GetUInt64 = &s2sdk_Kv3GetUInt64

//go:linkname s2sdk_Kv3GetFloat __package__/s2sdk._Kv3GetFloat
var s2sdk_Kv3GetFloat func(kv uintptr, defaultValue float32) float32
var S2sdk_Kv3GetFloat = &s2sdk_Kv3GetFloat

//go:linkname s2sdk_Kv3GetDouble __package__/s2sdk._Kv3GetDouble
var s2sdk_Kv3GetDouble func(kv uintptr, defaultValue float64) float64
var S2sdk_Kv3GetDouble = &s2sdk_Kv3GetDouble

//go:linkname s2sdk_Kv3SetBool __package__/s2sdk._Kv3SetBool
var s2sdk_Kv3SetBool func(kv uintptr, value bool)
var S2sdk_Kv3SetBool = &s2sdk_Kv3SetBool

//go:linkname s2sdk_Kv3SetChar __package__/s2sdk._Kv3SetChar
var s2sdk_Kv3SetChar func(kv uintptr, value int8)
var S2sdk_Kv3SetChar = &s2sdk_Kv3SetChar

//go:linkname s2sdk_Kv3SetUChar32 __package__/s2sdk._Kv3SetUChar32
var s2sdk_Kv3SetUChar32 func(kv uintptr, value uint32)
var S2sdk_Kv3SetUChar32 = &s2sdk_Kv3SetUChar32

//go:linkname s2sdk_Kv3SetInt8 __package__/s2sdk._Kv3SetInt8
var s2sdk_Kv3SetInt8 func(kv uintptr, value int8)
var S2sdk_Kv3SetInt8 = &s2sdk_Kv3SetInt8

//go:linkname s2sdk_Kv3SetUInt8 __package__/s2sdk._Kv3SetUInt8
var s2sdk_Kv3SetUInt8 func(kv uintptr, value uint8)
var S2sdk_Kv3SetUInt8 = &s2sdk_Kv3SetUInt8

//go:linkname s2sdk_Kv3SetShort __package__/s2sdk._Kv3SetShort
var s2sdk_Kv3SetShort func(kv uintptr, value int16)
var S2sdk_Kv3SetShort = &s2sdk_Kv3SetShort

//go:linkname s2sdk_Kv3SetUShort __package__/s2sdk._Kv3SetUShort
var s2sdk_Kv3SetUShort func(kv uintptr, value uint16)
var S2sdk_Kv3SetUShort = &s2sdk_Kv3SetUShort

//go:linkname s2sdk_Kv3SetInt __package__/s2sdk._Kv3SetInt
var s2sdk_Kv3SetInt func(kv uintptr, value int32)
var S2sdk_Kv3SetInt = &s2sdk_Kv3SetInt

//go:linkname s2sdk_Kv3SetUInt __package__/s2sdk._Kv3SetUInt
var s2sdk_Kv3SetUInt func(kv uintptr, value uint32)
var S2sdk_Kv3SetUInt = &s2sdk_Kv3SetUInt

//go:linkname s2sdk_Kv3SetInt64 __package__/s2sdk._Kv3SetInt64
var s2sdk_Kv3SetInt64 func(kv uintptr, value int64)
var S2sdk_Kv3SetInt64 = &s2sdk_Kv3SetInt64

//go:linkname s2sdk_Kv3SetUInt64 __package__/s2sdk._Kv3SetUInt64
var s2sdk_Kv3SetUInt64 func(kv uintptr, value uint64)
var S2sdk_Kv3SetUInt64 = &s2sdk_Kv3SetUInt64

//go:linkname s2sdk_Kv3SetFloat __package__/s2sdk._Kv3SetFloat
var s2sdk_Kv3SetFloat func(kv uintptr, value float32)
var S2sdk_Kv3SetFloat = &s2sdk_Kv3SetFloat

//go:linkname s2sdk_Kv3SetDouble __package__/s2sdk._Kv3SetDouble
var s2sdk_Kv3SetDouble func(kv uintptr, value float64)
var S2sdk_Kv3SetDouble = &s2sdk_Kv3SetDouble

//go:linkname s2sdk_Kv3GetPointer __package__/s2sdk._Kv3GetPointer
var s2sdk_Kv3GetPointer func(kv uintptr, defaultValue uintptr) uintptr
var S2sdk_Kv3GetPointer = &s2sdk_Kv3GetPointer

//go:linkname s2sdk_Kv3SetPointer __package__/s2sdk._Kv3SetPointer
var s2sdk_Kv3SetPointer func(kv uintptr, ptr uintptr)
var S2sdk_Kv3SetPointer = &s2sdk_Kv3SetPointer

//go:linkname s2sdk_Kv3GetStringToken __package__/s2sdk._Kv3GetStringToken
var s2sdk_Kv3GetStringToken func(kv uintptr, defaultValue uint32) uint32
var S2sdk_Kv3GetStringToken = &s2sdk_Kv3GetStringToken

//go:linkname s2sdk_Kv3SetStringToken __package__/s2sdk._Kv3SetStringToken
var s2sdk_Kv3SetStringToken func(kv uintptr, token uint32)
var S2sdk_Kv3SetStringToken = &s2sdk_Kv3SetStringToken

//go:linkname s2sdk_Kv3GetEHandle __package__/s2sdk._Kv3GetEHandle
var s2sdk_Kv3GetEHandle func(kv uintptr, defaultValue int32) int32
var S2sdk_Kv3GetEHandle = &s2sdk_Kv3GetEHandle

//go:linkname s2sdk_Kv3SetEHandle __package__/s2sdk._Kv3SetEHandle
var s2sdk_Kv3SetEHandle func(kv uintptr, ehandle int32)
var S2sdk_Kv3SetEHandle = &s2sdk_Kv3SetEHandle

//go:linkname s2sdk_Kv3GetString __package__/s2sdk._Kv3GetString
var s2sdk_Kv3GetString func(kv uintptr, defaultValue string) string
var S2sdk_Kv3GetString = &s2sdk_Kv3GetString

//go:linkname s2sdk_Kv3SetString __package__/s2sdk._Kv3SetString
var s2sdk_Kv3SetString func(kv uintptr, str string, subtype uint8)
var S2sdk_Kv3SetString = &s2sdk_Kv3SetString

//go:linkname s2sdk_Kv3SetStringExternal __package__/s2sdk._Kv3SetStringExternal
var s2sdk_Kv3SetStringExternal func(kv uintptr, str string, subtype uint8)
var S2sdk_Kv3SetStringExternal = &s2sdk_Kv3SetStringExternal

//go:linkname s2sdk_Kv3GetBinaryBlob __package__/s2sdk._Kv3GetBinaryBlob
var s2sdk_Kv3GetBinaryBlob func(kv uintptr) []uint8
var S2sdk_Kv3GetBinaryBlob = &s2sdk_Kv3GetBinaryBlob

//go:linkname s2sdk_Kv3GetBinaryBlobSize __package__/s2sdk._Kv3GetBinaryBlobSize
var s2sdk_Kv3GetBinaryBlobSize func(kv uintptr) int32
var S2sdk_Kv3GetBinaryBlobSize = &s2sdk_Kv3GetBinaryBlobSize

//go:linkname s2sdk_Kv3SetToBinaryBlob __package__/s2sdk._Kv3SetToBinaryBlob
var s2sdk_Kv3SetToBinaryBlob func(kv uintptr, blob []uint8)
var S2sdk_Kv3SetToBinaryBlob = &s2sdk_Kv3SetToBinaryBlob

//go:linkname s2sdk_Kv3SetToBinaryBlobExternal __package__/s2sdk._Kv3SetToBinaryBlobExternal
var s2sdk_Kv3SetToBinaryBlobExternal func(kv uintptr, blob []uint8, free_mem bool)
var S2sdk_Kv3SetToBinaryBlobExternal = &s2sdk_Kv3SetToBinaryBlobExternal

//go:linkname s2sdk_Kv3GetColor __package__/s2sdk._Kv3GetColor
var s2sdk_Kv3GetColor func(kv uintptr, defaultValue plugify.Vector4) plugify.Vector4
var S2sdk_Kv3GetColor = &s2sdk_Kv3GetColor

//go:linkname s2sdk_Kv3SetColor __package__/s2sdk._Kv3SetColor
var s2sdk_Kv3SetColor func(kv uintptr, color plugify.Vector4)
var S2sdk_Kv3SetColor = &s2sdk_Kv3SetColor

//go:linkname s2sdk_Kv3GetVector __package__/s2sdk._Kv3GetVector
var s2sdk_Kv3GetVector func(kv uintptr, defaultValue plugify.Vector3) plugify.Vector3
var S2sdk_Kv3GetVector = &s2sdk_Kv3GetVector

//go:linkname s2sdk_Kv3GetVector2D __package__/s2sdk._Kv3GetVector2D
var s2sdk_Kv3GetVector2D func(kv uintptr, defaultValue plugify.Vector2) plugify.Vector2
var S2sdk_Kv3GetVector2D = &s2sdk_Kv3GetVector2D

//go:linkname s2sdk_Kv3GetVector4D __package__/s2sdk._Kv3GetVector4D
var s2sdk_Kv3GetVector4D func(kv uintptr, defaultValue plugify.Vector4) plugify.Vector4
var S2sdk_Kv3GetVector4D = &s2sdk_Kv3GetVector4D

//go:linkname s2sdk_Kv3GetQuaternion __package__/s2sdk._Kv3GetQuaternion
var s2sdk_Kv3GetQuaternion func(kv uintptr, defaultValue plugify.Vector4) plugify.Vector4
var S2sdk_Kv3GetQuaternion = &s2sdk_Kv3GetQuaternion

//go:linkname s2sdk_Kv3GetQAngle __package__/s2sdk._Kv3GetQAngle
var s2sdk_Kv3GetQAngle func(kv uintptr, defaultValue plugify.Vector3) plugify.Vector3
var S2sdk_Kv3GetQAngle = &s2sdk_Kv3GetQAngle

//go:linkname s2sdk_Kv3GetMatrix3x4 __package__/s2sdk._Kv3GetMatrix3x4
var s2sdk_Kv3GetMatrix3x4 func(kv uintptr, defaultValue plugify.Matrix4x4) plugify.Matrix4x4
var S2sdk_Kv3GetMatrix3x4 = &s2sdk_Kv3GetMatrix3x4

//go:linkname s2sdk_Kv3SetVector __package__/s2sdk._Kv3SetVector
var s2sdk_Kv3SetVector func(kv uintptr, vec plugify.Vector3)
var S2sdk_Kv3SetVector = &s2sdk_Kv3SetVector

//go:linkname s2sdk_Kv3SetVector2D __package__/s2sdk._Kv3SetVector2D
var s2sdk_Kv3SetVector2D func(kv uintptr, vec2d plugify.Vector2)
var S2sdk_Kv3SetVector2D = &s2sdk_Kv3SetVector2D

//go:linkname s2sdk_Kv3SetVector4D __package__/s2sdk._Kv3SetVector4D
var s2sdk_Kv3SetVector4D func(kv uintptr, vec4d plugify.Vector4)
var S2sdk_Kv3SetVector4D = &s2sdk_Kv3SetVector4D

//go:linkname s2sdk_Kv3SetQuaternion __package__/s2sdk._Kv3SetQuaternion
var s2sdk_Kv3SetQuaternion func(kv uintptr, quat plugify.Vector4)
var S2sdk_Kv3SetQuaternion = &s2sdk_Kv3SetQuaternion

//go:linkname s2sdk_Kv3SetQAngle __package__/s2sdk._Kv3SetQAngle
var s2sdk_Kv3SetQAngle func(kv uintptr, ang plugify.Vector3)
var S2sdk_Kv3SetQAngle = &s2sdk_Kv3SetQAngle

//go:linkname s2sdk_Kv3SetMatrix3x4 __package__/s2sdk._Kv3SetMatrix3x4
var s2sdk_Kv3SetMatrix3x4 func(kv uintptr, matrix plugify.Matrix4x4)
var S2sdk_Kv3SetMatrix3x4 = &s2sdk_Kv3SetMatrix3x4

//go:linkname s2sdk_Kv3GetArrayElementCount __package__/s2sdk._Kv3GetArrayElementCount
var s2sdk_Kv3GetArrayElementCount func(kv uintptr) int32
var S2sdk_Kv3GetArrayElementCount = &s2sdk_Kv3GetArrayElementCount

//go:linkname s2sdk_Kv3SetArrayElementCount __package__/s2sdk._Kv3SetArrayElementCount
var s2sdk_Kv3SetArrayElementCount func(kv uintptr, count int32, type_ uint8, subtype uint8)
var S2sdk_Kv3SetArrayElementCount = &s2sdk_Kv3SetArrayElementCount

//go:linkname s2sdk_Kv3SetToEmptyKV3Array __package__/s2sdk._Kv3SetToEmptyKV3Array
var s2sdk_Kv3SetToEmptyKV3Array func(kv uintptr)
var S2sdk_Kv3SetToEmptyKV3Array = &s2sdk_Kv3SetToEmptyKV3Array

//go:linkname s2sdk_Kv3GetArrayElement __package__/s2sdk._Kv3GetArrayElement
var s2sdk_Kv3GetArrayElement func(kv uintptr, elem int32) uintptr
var S2sdk_Kv3GetArrayElement = &s2sdk_Kv3GetArrayElement

//go:linkname s2sdk_Kv3ArrayInsertElementBefore __package__/s2sdk._Kv3ArrayInsertElementBefore
var s2sdk_Kv3ArrayInsertElementBefore func(kv uintptr, elem int32) uintptr
var S2sdk_Kv3ArrayInsertElementBefore = &s2sdk_Kv3ArrayInsertElementBefore

//go:linkname s2sdk_Kv3ArrayInsertElementAfter __package__/s2sdk._Kv3ArrayInsertElementAfter
var s2sdk_Kv3ArrayInsertElementAfter func(kv uintptr, elem int32) uintptr
var S2sdk_Kv3ArrayInsertElementAfter = &s2sdk_Kv3ArrayInsertElementAfter

//go:linkname s2sdk_Kv3ArrayAddElementToTail __package__/s2sdk._Kv3ArrayAddElementToTail
var s2sdk_Kv3ArrayAddElementToTail func(kv uintptr) uintptr
var S2sdk_Kv3ArrayAddElementToTail = &s2sdk_Kv3ArrayAddElementToTail

//go:linkname s2sdk_Kv3ArraySwapItems __package__/s2sdk._Kv3ArraySwapItems
var s2sdk_Kv3ArraySwapItems func(kv uintptr, idx1 int32, idx2 int32)
var S2sdk_Kv3ArraySwapItems = &s2sdk_Kv3ArraySwapItems

//go:linkname s2sdk_Kv3ArrayRemoveElement __package__/s2sdk._Kv3ArrayRemoveElement
var s2sdk_Kv3ArrayRemoveElement func(kv uintptr, elem int32)
var S2sdk_Kv3ArrayRemoveElement = &s2sdk_Kv3ArrayRemoveElement

//go:linkname s2sdk_Kv3SetToEmptyTable __package__/s2sdk._Kv3SetToEmptyTable
var s2sdk_Kv3SetToEmptyTable func(kv uintptr)
var S2sdk_Kv3SetToEmptyTable = &s2sdk_Kv3SetToEmptyTable

//go:linkname s2sdk_Kv3GetMemberCount __package__/s2sdk._Kv3GetMemberCount
var s2sdk_Kv3GetMemberCount func(kv uintptr) int32
var S2sdk_Kv3GetMemberCount = &s2sdk_Kv3GetMemberCount

//go:linkname s2sdk_Kv3HasMember __package__/s2sdk._Kv3HasMember
var s2sdk_Kv3HasMember func(kv uintptr, name string) bool
var S2sdk_Kv3HasMember = &s2sdk_Kv3HasMember

//go:linkname s2sdk_Kv3FindMember __package__/s2sdk._Kv3FindMember
var s2sdk_Kv3FindMember func(kv uintptr, name string) uintptr
var S2sdk_Kv3FindMember = &s2sdk_Kv3FindMember

//go:linkname s2sdk_Kv3FindOrCreateMember __package__/s2sdk._Kv3FindOrCreateMember
var s2sdk_Kv3FindOrCreateMember func(kv uintptr, name string) uintptr
var S2sdk_Kv3FindOrCreateMember = &s2sdk_Kv3FindOrCreateMember

//go:linkname s2sdk_Kv3RemoveMember __package__/s2sdk._Kv3RemoveMember
var s2sdk_Kv3RemoveMember func(kv uintptr, name string) bool
var S2sdk_Kv3RemoveMember = &s2sdk_Kv3RemoveMember

//go:linkname s2sdk_Kv3GetMemberName __package__/s2sdk._Kv3GetMemberName
var s2sdk_Kv3GetMemberName func(kv uintptr, index int32) string
var S2sdk_Kv3GetMemberName = &s2sdk_Kv3GetMemberName

//go:linkname s2sdk_Kv3GetMemberByIndex __package__/s2sdk._Kv3GetMemberByIndex
var s2sdk_Kv3GetMemberByIndex func(kv uintptr, index int32) uintptr
var S2sdk_Kv3GetMemberByIndex = &s2sdk_Kv3GetMemberByIndex

//go:linkname s2sdk_Kv3GetMemberBool __package__/s2sdk._Kv3GetMemberBool
var s2sdk_Kv3GetMemberBool func(kv uintptr, name string, defaultValue bool) bool
var S2sdk_Kv3GetMemberBool = &s2sdk_Kv3GetMemberBool

//go:linkname s2sdk_Kv3GetMemberChar __package__/s2sdk._Kv3GetMemberChar
var s2sdk_Kv3GetMemberChar func(kv uintptr, name string, defaultValue int8) int8
var S2sdk_Kv3GetMemberChar = &s2sdk_Kv3GetMemberChar

//go:linkname s2sdk_Kv3GetMemberUChar32 __package__/s2sdk._Kv3GetMemberUChar32
var s2sdk_Kv3GetMemberUChar32 func(kv uintptr, name string, defaultValue uint32) uint32
var S2sdk_Kv3GetMemberUChar32 = &s2sdk_Kv3GetMemberUChar32

//go:linkname s2sdk_Kv3GetMemberInt8 __package__/s2sdk._Kv3GetMemberInt8
var s2sdk_Kv3GetMemberInt8 func(kv uintptr, name string, defaultValue int8) int8
var S2sdk_Kv3GetMemberInt8 = &s2sdk_Kv3GetMemberInt8

//go:linkname s2sdk_Kv3GetMemberUInt8 __package__/s2sdk._Kv3GetMemberUInt8
var s2sdk_Kv3GetMemberUInt8 func(kv uintptr, name string, defaultValue uint8) uint8
var S2sdk_Kv3GetMemberUInt8 = &s2sdk_Kv3GetMemberUInt8

//go:linkname s2sdk_Kv3GetMemberShort __package__/s2sdk._Kv3GetMemberShort
var s2sdk_Kv3GetMemberShort func(kv uintptr, name string, defaultValue int16) int16
var S2sdk_Kv3GetMemberShort = &s2sdk_Kv3GetMemberShort

//go:linkname s2sdk_Kv3GetMemberUShort __package__/s2sdk._Kv3GetMemberUShort
var s2sdk_Kv3GetMemberUShort func(kv uintptr, name string, defaultValue uint16) uint16
var S2sdk_Kv3GetMemberUShort = &s2sdk_Kv3GetMemberUShort

//go:linkname s2sdk_Kv3GetMemberInt __package__/s2sdk._Kv3GetMemberInt
var s2sdk_Kv3GetMemberInt func(kv uintptr, name string, defaultValue int32) int32
var S2sdk_Kv3GetMemberInt = &s2sdk_Kv3GetMemberInt

//go:linkname s2sdk_Kv3GetMemberUInt __package__/s2sdk._Kv3GetMemberUInt
var s2sdk_Kv3GetMemberUInt func(kv uintptr, name string, defaultValue uint32) uint32
var S2sdk_Kv3GetMemberUInt = &s2sdk_Kv3GetMemberUInt

//go:linkname s2sdk_Kv3GetMemberInt64 __package__/s2sdk._Kv3GetMemberInt64
var s2sdk_Kv3GetMemberInt64 func(kv uintptr, name string, defaultValue int64) int64
var S2sdk_Kv3GetMemberInt64 = &s2sdk_Kv3GetMemberInt64

//go:linkname s2sdk_Kv3GetMemberUInt64 __package__/s2sdk._Kv3GetMemberUInt64
var s2sdk_Kv3GetMemberUInt64 func(kv uintptr, name string, defaultValue uint64) uint64
var S2sdk_Kv3GetMemberUInt64 = &s2sdk_Kv3GetMemberUInt64

//go:linkname s2sdk_Kv3GetMemberFloat __package__/s2sdk._Kv3GetMemberFloat
var s2sdk_Kv3GetMemberFloat func(kv uintptr, name string, defaultValue float32) float32
var S2sdk_Kv3GetMemberFloat = &s2sdk_Kv3GetMemberFloat

//go:linkname s2sdk_Kv3GetMemberDouble __package__/s2sdk._Kv3GetMemberDouble
var s2sdk_Kv3GetMemberDouble func(kv uintptr, name string, defaultValue float64) float64
var S2sdk_Kv3GetMemberDouble = &s2sdk_Kv3GetMemberDouble

//go:linkname s2sdk_Kv3GetMemberPointer __package__/s2sdk._Kv3GetMemberPointer
var s2sdk_Kv3GetMemberPointer func(kv uintptr, name string, defaultValue uintptr) uintptr
var S2sdk_Kv3GetMemberPointer = &s2sdk_Kv3GetMemberPointer

//go:linkname s2sdk_Kv3GetMemberStringToken __package__/s2sdk._Kv3GetMemberStringToken
var s2sdk_Kv3GetMemberStringToken func(kv uintptr, name string, defaultValue uint32) uint32
var S2sdk_Kv3GetMemberStringToken = &s2sdk_Kv3GetMemberStringToken

//go:linkname s2sdk_Kv3GetMemberEHandle __package__/s2sdk._Kv3GetMemberEHandle
var s2sdk_Kv3GetMemberEHandle func(kv uintptr, name string, defaultValue int32) int32
var S2sdk_Kv3GetMemberEHandle = &s2sdk_Kv3GetMemberEHandle

//go:linkname s2sdk_Kv3GetMemberString __package__/s2sdk._Kv3GetMemberString
var s2sdk_Kv3GetMemberString func(kv uintptr, name string, defaultValue string) string
var S2sdk_Kv3GetMemberString = &s2sdk_Kv3GetMemberString

//go:linkname s2sdk_Kv3GetMemberColor __package__/s2sdk._Kv3GetMemberColor
var s2sdk_Kv3GetMemberColor func(kv uintptr, name string, defaultValue plugify.Vector4) plugify.Vector4
var S2sdk_Kv3GetMemberColor = &s2sdk_Kv3GetMemberColor

//go:linkname s2sdk_Kv3GetMemberVector __package__/s2sdk._Kv3GetMemberVector
var s2sdk_Kv3GetMemberVector func(kv uintptr, name string, defaultValue plugify.Vector3) plugify.Vector3
var S2sdk_Kv3GetMemberVector = &s2sdk_Kv3GetMemberVector

//go:linkname s2sdk_Kv3GetMemberVector2D __package__/s2sdk._Kv3GetMemberVector2D
var s2sdk_Kv3GetMemberVector2D func(kv uintptr, name string, defaultValue plugify.Vector2) plugify.Vector2
var S2sdk_Kv3GetMemberVector2D = &s2sdk_Kv3GetMemberVector2D

//go:linkname s2sdk_Kv3GetMemberVector4D __package__/s2sdk._Kv3GetMemberVector4D
var s2sdk_Kv3GetMemberVector4D func(kv uintptr, name string, defaultValue plugify.Vector4) plugify.Vector4
var S2sdk_Kv3GetMemberVector4D = &s2sdk_Kv3GetMemberVector4D

//go:linkname s2sdk_Kv3GetMemberQuaternion __package__/s2sdk._Kv3GetMemberQuaternion
var s2sdk_Kv3GetMemberQuaternion func(kv uintptr, name string, defaultValue plugify.Vector4) plugify.Vector4
var S2sdk_Kv3GetMemberQuaternion = &s2sdk_Kv3GetMemberQuaternion

//go:linkname s2sdk_Kv3GetMemberQAngle __package__/s2sdk._Kv3GetMemberQAngle
var s2sdk_Kv3GetMemberQAngle func(kv uintptr, name string, defaultValue plugify.Vector3) plugify.Vector3
var S2sdk_Kv3GetMemberQAngle = &s2sdk_Kv3GetMemberQAngle

//go:linkname s2sdk_Kv3GetMemberMatrix3x4 __package__/s2sdk._Kv3GetMemberMatrix3x4
var s2sdk_Kv3GetMemberMatrix3x4 func(kv uintptr, name string, defaultValue plugify.Matrix4x4) plugify.Matrix4x4
var S2sdk_Kv3GetMemberMatrix3x4 = &s2sdk_Kv3GetMemberMatrix3x4

//go:linkname s2sdk_Kv3SetMemberToNull __package__/s2sdk._Kv3SetMemberToNull
var s2sdk_Kv3SetMemberToNull func(kv uintptr, name string)
var S2sdk_Kv3SetMemberToNull = &s2sdk_Kv3SetMemberToNull

//go:linkname s2sdk_Kv3SetMemberToEmptyArray __package__/s2sdk._Kv3SetMemberToEmptyArray
var s2sdk_Kv3SetMemberToEmptyArray func(kv uintptr, name string)
var S2sdk_Kv3SetMemberToEmptyArray = &s2sdk_Kv3SetMemberToEmptyArray

//go:linkname s2sdk_Kv3SetMemberToEmptyTable __package__/s2sdk._Kv3SetMemberToEmptyTable
var s2sdk_Kv3SetMemberToEmptyTable func(kv uintptr, name string)
var S2sdk_Kv3SetMemberToEmptyTable = &s2sdk_Kv3SetMemberToEmptyTable

//go:linkname s2sdk_Kv3SetMemberToBinaryBlob __package__/s2sdk._Kv3SetMemberToBinaryBlob
var s2sdk_Kv3SetMemberToBinaryBlob func(kv uintptr, name string, blob []uint8)
var S2sdk_Kv3SetMemberToBinaryBlob = &s2sdk_Kv3SetMemberToBinaryBlob

//go:linkname s2sdk_Kv3SetMemberToBinaryBlobExternal __package__/s2sdk._Kv3SetMemberToBinaryBlobExternal
var s2sdk_Kv3SetMemberToBinaryBlobExternal func(kv uintptr, name string, blob []uint8, free_mem bool)
var S2sdk_Kv3SetMemberToBinaryBlobExternal = &s2sdk_Kv3SetMemberToBinaryBlobExternal

//go:linkname s2sdk_Kv3SetMemberToCopyOfValue __package__/s2sdk._Kv3SetMemberToCopyOfValue
var s2sdk_Kv3SetMemberToCopyOfValue func(kv uintptr, name string, other uintptr)
var S2sdk_Kv3SetMemberToCopyOfValue = &s2sdk_Kv3SetMemberToCopyOfValue

//go:linkname s2sdk_Kv3SetMemberBool __package__/s2sdk._Kv3SetMemberBool
var s2sdk_Kv3SetMemberBool func(kv uintptr, name string, value bool)
var S2sdk_Kv3SetMemberBool = &s2sdk_Kv3SetMemberBool

//go:linkname s2sdk_Kv3SetMemberChar __package__/s2sdk._Kv3SetMemberChar
var s2sdk_Kv3SetMemberChar func(kv uintptr, name string, value int8)
var S2sdk_Kv3SetMemberChar = &s2sdk_Kv3SetMemberChar

//go:linkname s2sdk_Kv3SetMemberUChar32 __package__/s2sdk._Kv3SetMemberUChar32
var s2sdk_Kv3SetMemberUChar32 func(kv uintptr, name string, value uint32)
var S2sdk_Kv3SetMemberUChar32 = &s2sdk_Kv3SetMemberUChar32

//go:linkname s2sdk_Kv3SetMemberInt8 __package__/s2sdk._Kv3SetMemberInt8
var s2sdk_Kv3SetMemberInt8 func(kv uintptr, name string, value int8)
var S2sdk_Kv3SetMemberInt8 = &s2sdk_Kv3SetMemberInt8

//go:linkname s2sdk_Kv3SetMemberUInt8 __package__/s2sdk._Kv3SetMemberUInt8
var s2sdk_Kv3SetMemberUInt8 func(kv uintptr, name string, value uint8)
var S2sdk_Kv3SetMemberUInt8 = &s2sdk_Kv3SetMemberUInt8

//go:linkname s2sdk_Kv3SetMemberShort __package__/s2sdk._Kv3SetMemberShort
var s2sdk_Kv3SetMemberShort func(kv uintptr, name string, value int16)
var S2sdk_Kv3SetMemberShort = &s2sdk_Kv3SetMemberShort

//go:linkname s2sdk_Kv3SetMemberUShort __package__/s2sdk._Kv3SetMemberUShort
var s2sdk_Kv3SetMemberUShort func(kv uintptr, name string, value uint16)
var S2sdk_Kv3SetMemberUShort = &s2sdk_Kv3SetMemberUShort

//go:linkname s2sdk_Kv3SetMemberInt __package__/s2sdk._Kv3SetMemberInt
var s2sdk_Kv3SetMemberInt func(kv uintptr, name string, value int32)
var S2sdk_Kv3SetMemberInt = &s2sdk_Kv3SetMemberInt

//go:linkname s2sdk_Kv3SetMemberUInt __package__/s2sdk._Kv3SetMemberUInt
var s2sdk_Kv3SetMemberUInt func(kv uintptr, name string, value uint32)
var S2sdk_Kv3SetMemberUInt = &s2sdk_Kv3SetMemberUInt

//go:linkname s2sdk_Kv3SetMemberInt64 __package__/s2sdk._Kv3SetMemberInt64
var s2sdk_Kv3SetMemberInt64 func(kv uintptr, name string, value int64)
var S2sdk_Kv3SetMemberInt64 = &s2sdk_Kv3SetMemberInt64

//go:linkname s2sdk_Kv3SetMemberUInt64 __package__/s2sdk._Kv3SetMemberUInt64
var s2sdk_Kv3SetMemberUInt64 func(kv uintptr, name string, value uint64)
var S2sdk_Kv3SetMemberUInt64 = &s2sdk_Kv3SetMemberUInt64

//go:linkname s2sdk_Kv3SetMemberFloat __package__/s2sdk._Kv3SetMemberFloat
var s2sdk_Kv3SetMemberFloat func(kv uintptr, name string, value float32)
var S2sdk_Kv3SetMemberFloat = &s2sdk_Kv3SetMemberFloat

//go:linkname s2sdk_Kv3SetMemberDouble __package__/s2sdk._Kv3SetMemberDouble
var s2sdk_Kv3SetMemberDouble func(kv uintptr, name string, value float64)
var S2sdk_Kv3SetMemberDouble = &s2sdk_Kv3SetMemberDouble

//go:linkname s2sdk_Kv3SetMemberPointer __package__/s2sdk._Kv3SetMemberPointer
var s2sdk_Kv3SetMemberPointer func(kv uintptr, name string, ptr uintptr)
var S2sdk_Kv3SetMemberPointer = &s2sdk_Kv3SetMemberPointer

//go:linkname s2sdk_Kv3SetMemberStringToken __package__/s2sdk._Kv3SetMemberStringToken
var s2sdk_Kv3SetMemberStringToken func(kv uintptr, name string, token uint32)
var S2sdk_Kv3SetMemberStringToken = &s2sdk_Kv3SetMemberStringToken

//go:linkname s2sdk_Kv3SetMemberEHandle __package__/s2sdk._Kv3SetMemberEHandle
var s2sdk_Kv3SetMemberEHandle func(kv uintptr, name string, ehandle int32)
var S2sdk_Kv3SetMemberEHandle = &s2sdk_Kv3SetMemberEHandle

//go:linkname s2sdk_Kv3SetMemberString __package__/s2sdk._Kv3SetMemberString
var s2sdk_Kv3SetMemberString func(kv uintptr, name string, str string, subtype uint8)
var S2sdk_Kv3SetMemberString = &s2sdk_Kv3SetMemberString

//go:linkname s2sdk_Kv3SetMemberStringExternal __package__/s2sdk._Kv3SetMemberStringExternal
var s2sdk_Kv3SetMemberStringExternal func(kv uintptr, name string, str string, subtype uint8)
var S2sdk_Kv3SetMemberStringExternal = &s2sdk_Kv3SetMemberStringExternal

//go:linkname s2sdk_Kv3SetMemberColor __package__/s2sdk._Kv3SetMemberColor
var s2sdk_Kv3SetMemberColor func(kv uintptr, name string, color plugify.Vector4)
var S2sdk_Kv3SetMemberColor = &s2sdk_Kv3SetMemberColor

//go:linkname s2sdk_Kv3SetMemberVector __package__/s2sdk._Kv3SetMemberVector
var s2sdk_Kv3SetMemberVector func(kv uintptr, name string, vec plugify.Vector3)
var S2sdk_Kv3SetMemberVector = &s2sdk_Kv3SetMemberVector

//go:linkname s2sdk_Kv3SetMemberVector2D __package__/s2sdk._Kv3SetMemberVector2D
var s2sdk_Kv3SetMemberVector2D func(kv uintptr, name string, vec2d plugify.Vector2)
var S2sdk_Kv3SetMemberVector2D = &s2sdk_Kv3SetMemberVector2D

//go:linkname s2sdk_Kv3SetMemberVector4D __package__/s2sdk._Kv3SetMemberVector4D
var s2sdk_Kv3SetMemberVector4D func(kv uintptr, name string, vec4d plugify.Vector4)
var S2sdk_Kv3SetMemberVector4D = &s2sdk_Kv3SetMemberVector4D

//go:linkname s2sdk_Kv3SetMemberQuaternion __package__/s2sdk._Kv3SetMemberQuaternion
var s2sdk_Kv3SetMemberQuaternion func(kv uintptr, name string, quat plugify.Vector4)
var S2sdk_Kv3SetMemberQuaternion = &s2sdk_Kv3SetMemberQuaternion

//go:linkname s2sdk_Kv3SetMemberQAngle __package__/s2sdk._Kv3SetMemberQAngle
var s2sdk_Kv3SetMemberQAngle func(kv uintptr, name string, ang plugify.Vector3)
var S2sdk_Kv3SetMemberQAngle = &s2sdk_Kv3SetMemberQAngle

//go:linkname s2sdk_Kv3SetMemberMatrix3x4 __package__/s2sdk._Kv3SetMemberMatrix3x4
var s2sdk_Kv3SetMemberMatrix3x4 func(kv uintptr, name string, matrix plugify.Matrix4x4)
var S2sdk_Kv3SetMemberMatrix3x4 = &s2sdk_Kv3SetMemberMatrix3x4

//go:linkname s2sdk_Kv3DebugPrint __package__/s2sdk._Kv3DebugPrint
var s2sdk_Kv3DebugPrint func(kv uintptr)
var S2sdk_Kv3DebugPrint = &s2sdk_Kv3DebugPrint

//go:linkname s2sdk_Kv3LoadFromBuffer __package__/s2sdk._Kv3LoadFromBuffer
var s2sdk_Kv3LoadFromBuffer func(context uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool
var S2sdk_Kv3LoadFromBuffer = &s2sdk_Kv3LoadFromBuffer

//go:linkname s2sdk_Kv3Load __package__/s2sdk._Kv3Load
var s2sdk_Kv3Load func(kv uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool
var S2sdk_Kv3Load = &s2sdk_Kv3Load

//go:linkname s2sdk_Kv3LoadFromText __package__/s2sdk._Kv3LoadFromText
var s2sdk_Kv3LoadFromText func(kv uintptr, error_ *string, input string, kv_name string, flags uint32) bool
var S2sdk_Kv3LoadFromText = &s2sdk_Kv3LoadFromText

//go:linkname s2sdk_Kv3LoadFromFileToContext __package__/s2sdk._Kv3LoadFromFileToContext
var s2sdk_Kv3LoadFromFileToContext func(context uintptr, error_ *string, filename string, path string, flags uint32) bool
var S2sdk_Kv3LoadFromFileToContext = &s2sdk_Kv3LoadFromFileToContext

//go:linkname s2sdk_Kv3LoadFromFile __package__/s2sdk._Kv3LoadFromFile
var s2sdk_Kv3LoadFromFile func(kv uintptr, error_ *string, filename string, path string, flags uint32) bool
var S2sdk_Kv3LoadFromFile = &s2sdk_Kv3LoadFromFile

//go:linkname s2sdk_Kv3LoadFromJSON __package__/s2sdk._Kv3LoadFromJSON
var s2sdk_Kv3LoadFromJSON func(kv uintptr, error_ *string, input string, kv_name string, flags uint32) bool
var S2sdk_Kv3LoadFromJSON = &s2sdk_Kv3LoadFromJSON

//go:linkname s2sdk_Kv3LoadFromJSONFile __package__/s2sdk._Kv3LoadFromJSONFile
var s2sdk_Kv3LoadFromJSONFile func(kv uintptr, error_ *string, path string, filename string, flags uint32) bool
var S2sdk_Kv3LoadFromJSONFile = &s2sdk_Kv3LoadFromJSONFile

//go:linkname s2sdk_Kv3LoadFromKV1File __package__/s2sdk._Kv3LoadFromKV1File
var s2sdk_Kv3LoadFromKV1File func(kv uintptr, error_ *string, path string, filename string, esc_behavior uint8, flags uint32) bool
var S2sdk_Kv3LoadFromKV1File = &s2sdk_Kv3LoadFromKV1File

//go:linkname s2sdk_Kv3LoadFromKV1Text __package__/s2sdk._Kv3LoadFromKV1Text
var s2sdk_Kv3LoadFromKV1Text func(kv uintptr, error_ *string, input string, esc_behavior uint8, kv_name string, unk bool, flags uint32) bool
var S2sdk_Kv3LoadFromKV1Text = &s2sdk_Kv3LoadFromKV1Text

//go:linkname s2sdk_Kv3LoadFromKV1TextTranslated __package__/s2sdk._Kv3LoadFromKV1TextTranslated
var s2sdk_Kv3LoadFromKV1TextTranslated func(kv uintptr, error_ *string, input string, esc_behavior uint8, translation uintptr, unk1 int32, kv_name string, unk2 bool, flags uint32) bool
var S2sdk_Kv3LoadFromKV1TextTranslated = &s2sdk_Kv3LoadFromKV1TextTranslated

//go:linkname s2sdk_Kv3LoadFromKV3OrKV1 __package__/s2sdk._Kv3LoadFromKV3OrKV1
var s2sdk_Kv3LoadFromKV3OrKV1 func(kv uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool
var S2sdk_Kv3LoadFromKV3OrKV1 = &s2sdk_Kv3LoadFromKV3OrKV1

//go:linkname s2sdk_Kv3LoadFromOldSchemaText __package__/s2sdk._Kv3LoadFromOldSchemaText
var s2sdk_Kv3LoadFromOldSchemaText func(kv uintptr, error_ *string, input []uint8, kv_name string, flags uint32) bool
var S2sdk_Kv3LoadFromOldSchemaText = &s2sdk_Kv3LoadFromOldSchemaText

//go:linkname s2sdk_Kv3LoadTextNoHeader __package__/s2sdk._Kv3LoadTextNoHeader
var s2sdk_Kv3LoadTextNoHeader func(kv uintptr, error_ *string, input string, kv_name string, flags uint32) bool
var S2sdk_Kv3LoadTextNoHeader = &s2sdk_Kv3LoadTextNoHeader

//go:linkname s2sdk_Kv3Save __package__/s2sdk._Kv3Save
var s2sdk_Kv3Save func(kv uintptr, error_ *string, output *[]uint8, flags uint32) bool
var S2sdk_Kv3Save = &s2sdk_Kv3Save

//go:linkname s2sdk_Kv3SaveAsJSON __package__/s2sdk._Kv3SaveAsJSON
var s2sdk_Kv3SaveAsJSON func(kv uintptr, error_ *string, output *[]uint8) bool
var S2sdk_Kv3SaveAsJSON = &s2sdk_Kv3SaveAsJSON

//go:linkname s2sdk_Kv3SaveAsJSONString __package__/s2sdk._Kv3SaveAsJSONString
var s2sdk_Kv3SaveAsJSONString func(kv uintptr, error_ *string, output *string) bool
var S2sdk_Kv3SaveAsJSONString = &s2sdk_Kv3SaveAsJSONString

//go:linkname s2sdk_Kv3SaveAsKV1Text __package__/s2sdk._Kv3SaveAsKV1Text
var s2sdk_Kv3SaveAsKV1Text func(kv uintptr, error_ *string, output *[]uint8, esc_behavior uint8) bool
var S2sdk_Kv3SaveAsKV1Text = &s2sdk_Kv3SaveAsKV1Text

//go:linkname s2sdk_Kv3SaveAsKV1TextTranslated __package__/s2sdk._Kv3SaveAsKV1TextTranslated
var s2sdk_Kv3SaveAsKV1TextTranslated func(kv uintptr, error_ *string, output *[]uint8, esc_behavior uint8, translation uintptr, unk int32) bool
var S2sdk_Kv3SaveAsKV1TextTranslated = &s2sdk_Kv3SaveAsKV1TextTranslated

//go:linkname s2sdk_Kv3SaveTextNoHeaderToBuffer __package__/s2sdk._Kv3SaveTextNoHeaderToBuffer
var s2sdk_Kv3SaveTextNoHeaderToBuffer func(kv uintptr, error_ *string, output *[]uint8, flags uint32) bool
var S2sdk_Kv3SaveTextNoHeaderToBuffer = &s2sdk_Kv3SaveTextNoHeaderToBuffer

//go:linkname s2sdk_Kv3SaveTextNoHeader __package__/s2sdk._Kv3SaveTextNoHeader
var s2sdk_Kv3SaveTextNoHeader func(kv uintptr, error_ *string, output *string, flags uint32) bool
var S2sdk_Kv3SaveTextNoHeader = &s2sdk_Kv3SaveTextNoHeader

//go:linkname s2sdk_Kv3SaveTextToString __package__/s2sdk._Kv3SaveTextToString
var s2sdk_Kv3SaveTextToString func(kv uintptr, error_ *string, output *string, flags uint32) bool
var S2sdk_Kv3SaveTextToString = &s2sdk_Kv3SaveTextToString

//go:linkname s2sdk_Kv3SaveToFile __package__/s2sdk._Kv3SaveToFile
var s2sdk_Kv3SaveToFile func(kv uintptr, error_ *string, filename string, path string, flags uint32) bool
var S2sdk_Kv3SaveToFile = &s2sdk_Kv3SaveToFile

//go:linkname s2sdk_DebugBreak __package__/s2sdk._DebugBreak
var s2sdk_DebugBreak func()
var S2sdk_DebugBreak = &s2sdk_DebugBreak

//go:linkname s2sdk_DebugDrawBox __package__/s2sdk._DebugDrawBox
var s2sdk_DebugDrawBox func(center plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, r int32, g int32, b int32, a int32, duration float32)
var S2sdk_DebugDrawBox = &s2sdk_DebugDrawBox

//go:linkname s2sdk_DebugDrawBoxDirection __package__/s2sdk._DebugDrawBoxDirection
var s2sdk_DebugDrawBoxDirection func(center plugify.Vector3, mins plugify.Vector3, maxs plugify.Vector3, forward plugify.Vector3, color plugify.Vector3, alpha float32, duration float32)
var S2sdk_DebugDrawBoxDirection = &s2sdk_DebugDrawBoxDirection

//go:linkname s2sdk_DebugDrawCircle __package__/s2sdk._DebugDrawCircle
var s2sdk_DebugDrawCircle func(center plugify.Vector3, color plugify.Vector3, alpha float32, radius float32, zTest bool, duration float32)
var S2sdk_DebugDrawCircle = &s2sdk_DebugDrawCircle

//go:linkname s2sdk_DebugDrawClear __package__/s2sdk._DebugDrawClear
var s2sdk_DebugDrawClear func()
var S2sdk_DebugDrawClear = &s2sdk_DebugDrawClear

//go:linkname s2sdk_DebugDrawLine __package__/s2sdk._DebugDrawLine
var s2sdk_DebugDrawLine func(origin plugify.Vector3, target plugify.Vector3, r int32, g int32, b int32, zTest bool, duration float32)
var S2sdk_DebugDrawLine = &s2sdk_DebugDrawLine

//go:linkname s2sdk_DebugDrawLine_vCol __package__/s2sdk._DebugDrawLine_vCol
var s2sdk_DebugDrawLine_vCol func(start plugify.Vector3, end plugify.Vector3, color plugify.Vector3, zTest bool, duration float32)
var S2sdk_DebugDrawLine_vCol = &s2sdk_DebugDrawLine_vCol

//go:linkname s2sdk_DebugDrawScreenTextLine __package__/s2sdk._DebugDrawScreenTextLine
var s2sdk_DebugDrawScreenTextLine func(x float32, y float32, lineOffset int32, text string, r int32, g int32, b int32, a int32, duration float32)
var S2sdk_DebugDrawScreenTextLine = &s2sdk_DebugDrawScreenTextLine

//go:linkname s2sdk_DebugDrawSphere __package__/s2sdk._DebugDrawSphere
var s2sdk_DebugDrawSphere func(center plugify.Vector3, color plugify.Vector3, alpha float32, radius float32, zTest bool, duration float32)
var S2sdk_DebugDrawSphere = &s2sdk_DebugDrawSphere

//go:linkname s2sdk_DebugDrawText __package__/s2sdk._DebugDrawText
var s2sdk_DebugDrawText func(origin plugify.Vector3, text string, viewCheck bool, duration float32)
var S2sdk_DebugDrawText = &s2sdk_DebugDrawText

//go:linkname s2sdk_DebugScreenTextPretty __package__/s2sdk._DebugScreenTextPretty
var s2sdk_DebugScreenTextPretty func(x float32, y float32, lineOffset int32, text string, r int32, g int32, b int32, a int32, duration float32, font string, size int32, bold bool)
var S2sdk_DebugScreenTextPretty = &s2sdk_DebugScreenTextPretty

//go:linkname s2sdk_DebugScriptAssert __package__/s2sdk._DebugScriptAssert
var s2sdk_DebugScriptAssert func(assertion bool, message string)
var S2sdk_DebugScriptAssert = &s2sdk_DebugScriptAssert

//go:linkname s2sdk_AnglesDiff __package__/s2sdk._AnglesDiff
var s2sdk_AnglesDiff func(angle1 float32, angle2 float32) float32
var S2sdk_AnglesDiff = &s2sdk_AnglesDiff

//go:linkname s2sdk_AnglesToVector __package__/s2sdk._AnglesToVector
var s2sdk_AnglesToVector func(angles plugify.Vector3) plugify.Vector3
var S2sdk_AnglesToVector = &s2sdk_AnglesToVector

//go:linkname s2sdk_AxisAngleToQuaternion __package__/s2sdk._AxisAngleToQuaternion
var s2sdk_AxisAngleToQuaternion func(axis plugify.Vector3, angle float32) plugify.Vector4
var S2sdk_AxisAngleToQuaternion = &s2sdk_AxisAngleToQuaternion

//go:linkname s2sdk_CalcClosestPointOnEntityOBB __package__/s2sdk._CalcClosestPointOnEntityOBB
var s2sdk_CalcClosestPointOnEntityOBB func(entityHandle int32, position plugify.Vector3) plugify.Vector3
var S2sdk_CalcClosestPointOnEntityOBB = &s2sdk_CalcClosestPointOnEntityOBB

//go:linkname s2sdk_CalcDistanceBetweenEntityOBB __package__/s2sdk._CalcDistanceBetweenEntityOBB
var s2sdk_CalcDistanceBetweenEntityOBB func(entityHandle1 int32, entityHandle2 int32) float32
var S2sdk_CalcDistanceBetweenEntityOBB = &s2sdk_CalcDistanceBetweenEntityOBB

//go:linkname s2sdk_CalcDistanceToLineSegment2D __package__/s2sdk._CalcDistanceToLineSegment2D
var s2sdk_CalcDistanceToLineSegment2D func(p plugify.Vector3, vLineA plugify.Vector3, vLineB plugify.Vector3) float32
var S2sdk_CalcDistanceToLineSegment2D = &s2sdk_CalcDistanceToLineSegment2D

//go:linkname s2sdk_CrossVectors __package__/s2sdk._CrossVectors
var s2sdk_CrossVectors func(v1 plugify.Vector3, v2 plugify.Vector3) plugify.Vector3
var S2sdk_CrossVectors = &s2sdk_CrossVectors

//go:linkname s2sdk_ExponentDecay __package__/s2sdk._ExponentDecay
var s2sdk_ExponentDecay func(decayTo float32, decayTime float32, dt float32) float32
var S2sdk_ExponentDecay = &s2sdk_ExponentDecay

//go:linkname s2sdk_LerpVectors __package__/s2sdk._LerpVectors
var s2sdk_LerpVectors func(start plugify.Vector3, end plugify.Vector3, factor float32) plugify.Vector3
var S2sdk_LerpVectors = &s2sdk_LerpVectors

//go:linkname s2sdk_QSlerp __package__/s2sdk._QSlerp
var s2sdk_QSlerp func(fromAngle plugify.Vector3, toAngle plugify.Vector3, time float32) plugify.Vector3
var S2sdk_QSlerp = &s2sdk_QSlerp

//go:linkname s2sdk_RotateOrientation __package__/s2sdk._RotateOrientation
var s2sdk_RotateOrientation func(a1 plugify.Vector3, a2 plugify.Vector3) plugify.Vector3
var S2sdk_RotateOrientation = &s2sdk_RotateOrientation

//go:linkname s2sdk_RotatePosition __package__/s2sdk._RotatePosition
var s2sdk_RotatePosition func(rotationOrigin plugify.Vector3, rotationAngle plugify.Vector3, vectorToRotate plugify.Vector3) plugify.Vector3
var S2sdk_RotatePosition = &s2sdk_RotatePosition

//go:linkname s2sdk_RotateQuaternionByAxisAngle __package__/s2sdk._RotateQuaternionByAxisAngle
var s2sdk_RotateQuaternionByAxisAngle func(q plugify.Vector4, axis plugify.Vector3, angle float32) plugify.Vector4
var S2sdk_RotateQuaternionByAxisAngle = &s2sdk_RotateQuaternionByAxisAngle

//go:linkname s2sdk_RotationDelta __package__/s2sdk._RotationDelta
var s2sdk_RotationDelta func(src plugify.Vector3, dest plugify.Vector3) plugify.Vector3
var S2sdk_RotationDelta = &s2sdk_RotationDelta

//go:linkname s2sdk_RotationDeltaAsAngularVelocity __package__/s2sdk._RotationDeltaAsAngularVelocity
var s2sdk_RotationDeltaAsAngularVelocity func(a1 plugify.Vector3, a2 plugify.Vector3) plugify.Vector3
var S2sdk_RotationDeltaAsAngularVelocity = &s2sdk_RotationDeltaAsAngularVelocity

//go:linkname s2sdk_SplineQuaternions __package__/s2sdk._SplineQuaternions
var s2sdk_SplineQuaternions func(q0 plugify.Vector4, q1 plugify.Vector4, t float32) plugify.Vector4
var S2sdk_SplineQuaternions = &s2sdk_SplineQuaternions

//go:linkname s2sdk_SplineVectors __package__/s2sdk._SplineVectors
var s2sdk_SplineVectors func(v0 plugify.Vector3, v1 plugify.Vector3, t float32) plugify.Vector3
var S2sdk_SplineVectors = &s2sdk_SplineVectors

//go:linkname s2sdk_VectorToAngles __package__/s2sdk._VectorToAngles
var s2sdk_VectorToAngles func(input plugify.Vector3) plugify.Vector3
var S2sdk_VectorToAngles = &s2sdk_VectorToAngles

//go:linkname s2sdk_RandomFlt __package__/s2sdk._RandomFlt
var s2sdk_RandomFlt func(min float32, max float32) float32
var S2sdk_RandomFlt = &s2sdk_RandomFlt

//go:linkname s2sdk_RandomInt __package__/s2sdk._RandomInt
var s2sdk_RandomInt func(min int32, max int32) int32
var S2sdk_RandomInt = &s2sdk_RandomInt

//go:linkname s2sdk_TraceCollideable __package__/s2sdk._TraceCollideable
var s2sdk_TraceCollideable func(start plugify.Vector3, end plugify.Vector3, entityHandle int32, outPos *plugify.Vector3, outFraction *float64, outHit *bool, outStartSolid *bool, outNormal *plugify.Vector3) bool
var S2sdk_TraceCollideable = &s2sdk_TraceCollideable

//go:linkname s2sdk_TraceCollideable2 __package__/s2sdk._TraceCollideable2
var s2sdk_TraceCollideable2 func(start plugify.Vector3, end plugify.Vector3, entityHandle int32, mins uintptr, maxs uintptr, outPos *plugify.Vector3, outFraction *float64, outHit *bool, outStartSolid *bool, outNormal *plugify.Vector3) bool
var S2sdk_TraceCollideable2 = &s2sdk_TraceCollideable2

//go:linkname s2sdk_TraceHull __package__/s2sdk._TraceHull
var s2sdk_TraceHull func(start plugify.Vector3, end plugify.Vector3, min plugify.Vector3, max plugify.Vector3, mask int32, ignoreHandle int32, outPos *plugify.Vector3, outFraction *float64, outHit *bool, outEntHit *int32, outStartSolid *bool) bool
var S2sdk_TraceHull = &s2sdk_TraceHull

//go:linkname s2sdk_TraceLine __package__/s2sdk._TraceLine
var s2sdk_TraceLine func(startPos plugify.Vector3, endPos plugify.Vector3, mask int32, ignoreHandle int32, outPos *plugify.Vector3, outFraction *float64, outHit *bool, outEntHit *int32, outStartSolid *bool) bool
var S2sdk_TraceLine = &s2sdk_TraceLine

//go:linkname s2sdk_SetTransmitInfoEntity __package__/s2sdk._SetTransmitInfoEntity
var s2sdk_SetTransmitInfoEntity func(info uintptr, entityHandle int32)
var S2sdk_SetTransmitInfoEntity = &s2sdk_SetTransmitInfoEntity

//go:linkname s2sdk_ClearTransmitInfoEntity __package__/s2sdk._ClearTransmitInfoEntity
var s2sdk_ClearTransmitInfoEntity func(info uintptr, entityHandle int32)
var S2sdk_ClearTransmitInfoEntity = &s2sdk_ClearTransmitInfoEntity

//go:linkname s2sdk_IsTransmitInfoEntitySet __package__/s2sdk._IsTransmitInfoEntitySet
var s2sdk_IsTransmitInfoEntitySet func(info uintptr, entityHandle int32) bool
var S2sdk_IsTransmitInfoEntitySet = &s2sdk_IsTransmitInfoEntitySet

//go:linkname s2sdk_SetTransmitInfoEntityAll __package__/s2sdk._SetTransmitInfoEntityAll
var s2sdk_SetTransmitInfoEntityAll func(info uintptr)
var S2sdk_SetTransmitInfoEntityAll = &s2sdk_SetTransmitInfoEntityAll

//go:linkname s2sdk_ClearTransmitInfoEntityAll __package__/s2sdk._ClearTransmitInfoEntityAll
var s2sdk_ClearTransmitInfoEntityAll func(info uintptr)
var S2sdk_ClearTransmitInfoEntityAll = &s2sdk_ClearTransmitInfoEntityAll

//go:linkname s2sdk_SetTransmitInfoNonPlayer __package__/s2sdk._SetTransmitInfoNonPlayer
var s2sdk_SetTransmitInfoNonPlayer func(info uintptr, entityHandle int32)
var S2sdk_SetTransmitInfoNonPlayer = &s2sdk_SetTransmitInfoNonPlayer

//go:linkname s2sdk_ClearTransmitInfoNonPlayer __package__/s2sdk._ClearTransmitInfoNonPlayer
var s2sdk_ClearTransmitInfoNonPlayer func(info uintptr, entityHandle int32)
var S2sdk_ClearTransmitInfoNonPlayer = &s2sdk_ClearTransmitInfoNonPlayer

//go:linkname s2sdk_IsTransmitInfoNonPlayerSet __package__/s2sdk._IsTransmitInfoNonPlayerSet
var s2sdk_IsTransmitInfoNonPlayerSet func(info uintptr, entityHandle int32) bool
var S2sdk_IsTransmitInfoNonPlayerSet = &s2sdk_IsTransmitInfoNonPlayerSet

//go:linkname s2sdk_SetTransmitInfoNonPlayerAll __package__/s2sdk._SetTransmitInfoNonPlayerAll
var s2sdk_SetTransmitInfoNonPlayerAll func(info uintptr)
var S2sdk_SetTransmitInfoNonPlayerAll = &s2sdk_SetTransmitInfoNonPlayerAll

//go:linkname s2sdk_ClearTransmitInfoNonPlayerAll __package__/s2sdk._ClearTransmitInfoNonPlayerAll
var s2sdk_ClearTransmitInfoNonPlayerAll func(info uintptr)
var S2sdk_ClearTransmitInfoNonPlayerAll = &s2sdk_ClearTransmitInfoNonPlayerAll

//go:linkname s2sdk_SetTransmitInfoOutOfPVS __package__/s2sdk._SetTransmitInfoOutOfPVS
var s2sdk_SetTransmitInfoOutOfPVS func(info uintptr, entityHandle int32)
var S2sdk_SetTransmitInfoOutOfPVS = &s2sdk_SetTransmitInfoOutOfPVS

//go:linkname s2sdk_ClearTransmitInfoOutOfPVS __package__/s2sdk._ClearTransmitInfoOutOfPVS
var s2sdk_ClearTransmitInfoOutOfPVS func(info uintptr, entityHandle int32)
var S2sdk_ClearTransmitInfoOutOfPVS = &s2sdk_ClearTransmitInfoOutOfPVS

//go:linkname s2sdk_IsTransmitInfoOutOfPVSSet __package__/s2sdk._IsTransmitInfoOutOfPVSSet
var s2sdk_IsTransmitInfoOutOfPVSSet func(info uintptr, entityHandle int32) bool
var S2sdk_IsTransmitInfoOutOfPVSSet = &s2sdk_IsTransmitInfoOutOfPVSSet

//go:linkname s2sdk_SetTransmitInfoOutOfPVSAll __package__/s2sdk._SetTransmitInfoOutOfPVSAll
var s2sdk_SetTransmitInfoOutOfPVSAll func(info uintptr)
var S2sdk_SetTransmitInfoOutOfPVSAll = &s2sdk_SetTransmitInfoOutOfPVSAll

//go:linkname s2sdk_ClearTransmitInfoOutOfPVSAll __package__/s2sdk._ClearTransmitInfoOutOfPVSAll
var s2sdk_ClearTransmitInfoOutOfPVSAll func(info uintptr)
var S2sdk_ClearTransmitInfoOutOfPVSAll = &s2sdk_ClearTransmitInfoOutOfPVSAll

//go:linkname s2sdk_SetTransmitInfoAlways __package__/s2sdk._SetTransmitInfoAlways
var s2sdk_SetTransmitInfoAlways func(info uintptr, entityHandle int32)
var S2sdk_SetTransmitInfoAlways = &s2sdk_SetTransmitInfoAlways

//go:linkname s2sdk_ClearTransmitInfoAlways __package__/s2sdk._ClearTransmitInfoAlways
var s2sdk_ClearTransmitInfoAlways func(info uintptr, entityHandle int32)
var S2sdk_ClearTransmitInfoAlways = &s2sdk_ClearTransmitInfoAlways

//go:linkname s2sdk_IsTransmitInfoAlwaysSet __package__/s2sdk._IsTransmitInfoAlwaysSet
var s2sdk_IsTransmitInfoAlwaysSet func(info uintptr, entityHandle int32) bool
var S2sdk_IsTransmitInfoAlwaysSet = &s2sdk_IsTransmitInfoAlwaysSet

//go:linkname s2sdk_SetTransmitInfoAlwaysAll __package__/s2sdk._SetTransmitInfoAlwaysAll
var s2sdk_SetTransmitInfoAlwaysAll func(info uintptr)
var S2sdk_SetTransmitInfoAlwaysAll = &s2sdk_SetTransmitInfoAlwaysAll

//go:linkname s2sdk_ClearTransmitInfoAlwaysAll __package__/s2sdk._ClearTransmitInfoAlwaysAll
var s2sdk_ClearTransmitInfoAlwaysAll func(info uintptr)
var S2sdk_ClearTransmitInfoAlwaysAll = &s2sdk_ClearTransmitInfoAlwaysAll

//go:linkname s2sdk_GetTransmitInfoTargetSlotsCount __package__/s2sdk._GetTransmitInfoTargetSlotsCount
var s2sdk_GetTransmitInfoTargetSlotsCount func(info uintptr) int32
var S2sdk_GetTransmitInfoTargetSlotsCount = &s2sdk_GetTransmitInfoTargetSlotsCount

//go:linkname s2sdk_GetTransmitInfoTargetSlot __package__/s2sdk._GetTransmitInfoTargetSlot
var s2sdk_GetTransmitInfoTargetSlot func(info uintptr, index int32) int32
var S2sdk_GetTransmitInfoTargetSlot = &s2sdk_GetTransmitInfoTargetSlot

//go:linkname s2sdk_AddTransmitInfoTargetSlot __package__/s2sdk._AddTransmitInfoTargetSlot
var s2sdk_AddTransmitInfoTargetSlot func(info uintptr, playerSlot int32)
var S2sdk_AddTransmitInfoTargetSlot = &s2sdk_AddTransmitInfoTargetSlot

//go:linkname s2sdk_RemoveTransmitInfoTargetSlot __package__/s2sdk._RemoveTransmitInfoTargetSlot
var s2sdk_RemoveTransmitInfoTargetSlot func(info uintptr, index int32)
var S2sdk_RemoveTransmitInfoTargetSlot = &s2sdk_RemoveTransmitInfoTargetSlot

//go:linkname s2sdk_GetTransmitInfoTargetSlotsAll __package__/s2sdk._GetTransmitInfoTargetSlotsAll
var s2sdk_GetTransmitInfoTargetSlotsAll func(info uintptr) []int32
var S2sdk_GetTransmitInfoTargetSlotsAll = &s2sdk_GetTransmitInfoTargetSlotsAll

//go:linkname s2sdk_RemoveTransmitInfoTargetSlotsAll __package__/s2sdk._RemoveTransmitInfoTargetSlotsAll
var s2sdk_RemoveTransmitInfoTargetSlotsAll func(info uintptr)
var S2sdk_RemoveTransmitInfoTargetSlotsAll = &s2sdk_RemoveTransmitInfoTargetSlotsAll

//go:linkname s2sdk_GetTransmitInfoPlayerSlot __package__/s2sdk._GetTransmitInfoPlayerSlot
var s2sdk_GetTransmitInfoPlayerSlot func(info uintptr) int32
var S2sdk_GetTransmitInfoPlayerSlot = &s2sdk_GetTransmitInfoPlayerSlot

//go:linkname s2sdk_SetTransmitInfoPlayerSlot __package__/s2sdk._SetTransmitInfoPlayerSlot
var s2sdk_SetTransmitInfoPlayerSlot func(info uintptr, playerSlot int32)
var S2sdk_SetTransmitInfoPlayerSlot = &s2sdk_SetTransmitInfoPlayerSlot

//go:linkname s2sdk_GetTransmitInfoFullUpdate __package__/s2sdk._GetTransmitInfoFullUpdate
var s2sdk_GetTransmitInfoFullUpdate func(info uintptr) bool
var S2sdk_GetTransmitInfoFullUpdate = &s2sdk_GetTransmitInfoFullUpdate

//go:linkname s2sdk_SetTransmitInfoFullUpdate __package__/s2sdk._SetTransmitInfoFullUpdate
var s2sdk_SetTransmitInfoFullUpdate func(info uintptr, fullUpdate bool)
var S2sdk_SetTransmitInfoFullUpdate = &s2sdk_SetTransmitInfoFullUpdate

//go:linkname s2sdk_HideTransmitEntities __package__/s2sdk._HideTransmitEntities
var s2sdk_HideTransmitEntities func(playerSlot int32, entHandles []int32)
var S2sdk_HideTransmitEntities = &s2sdk_HideTransmitEntities

//go:linkname s2sdk_ShowTransmitEntities __package__/s2sdk._ShowTransmitEntities
var s2sdk_ShowTransmitEntities func(playerSlot int32, entHandles []int32)
var S2sdk_ShowTransmitEntities = &s2sdk_ShowTransmitEntities

//go:linkname s2sdk_GetHiddenTransmitEntities __package__/s2sdk._GetHiddenTransmitEntities
var s2sdk_GetHiddenTransmitEntities func(playerSlot int32) []int32
var S2sdk_GetHiddenTransmitEntities = &s2sdk_GetHiddenTransmitEntities

//go:linkname s2sdk_HideTransmitEntity __package__/s2sdk._HideTransmitEntity
var s2sdk_HideTransmitEntity func(playerSlot int32, entityHandle int32)
var S2sdk_HideTransmitEntity = &s2sdk_HideTransmitEntity

//go:linkname s2sdk_ShowTransmitEntity __package__/s2sdk._ShowTransmitEntity
var s2sdk_ShowTransmitEntity func(playerSlot int32, entityHandle int32)
var S2sdk_ShowTransmitEntity = &s2sdk_ShowTransmitEntity

//go:linkname s2sdk_HideTransmitEntityFromOtherPlayers __package__/s2sdk._HideTransmitEntityFromOtherPlayers
var s2sdk_HideTransmitEntityFromOtherPlayers func(playerSlot int32, entityHandle int32)
var S2sdk_HideTransmitEntityFromOtherPlayers = &s2sdk_HideTransmitEntityFromOtherPlayers

//go:linkname s2sdk_ShowTransmitEntityToOtherPlayers __package__/s2sdk._ShowTransmitEntityToOtherPlayers
var s2sdk_ShowTransmitEntityToOtherPlayers func(playerSlot int32, entityHandle int32)
var S2sdk_ShowTransmitEntityToOtherPlayers = &s2sdk_ShowTransmitEntityToOtherPlayers

//go:linkname s2sdk_AddBodyImpulseAtPosition __package__/s2sdk._AddBodyImpulseAtPosition
var s2sdk_AddBodyImpulseAtPosition func(entityHandle int32, position plugify.Vector3, impulse plugify.Vector3)
var S2sdk_AddBodyImpulseAtPosition = &s2sdk_AddBodyImpulseAtPosition

//go:linkname s2sdk_AddBodyVelocity __package__/s2sdk._AddBodyVelocity
var s2sdk_AddBodyVelocity func(entityHandle int32, linearVelocity plugify.Vector3, angularVelocity plugify.Vector3)
var S2sdk_AddBodyVelocity = &s2sdk_AddBodyVelocity

//go:linkname s2sdk_DetachBodyFromParent __package__/s2sdk._DetachBodyFromParent
var s2sdk_DetachBodyFromParent func(entityHandle int32)
var S2sdk_DetachBodyFromParent = &s2sdk_DetachBodyFromParent

//go:linkname s2sdk_GetBodySequence __package__/s2sdk._GetBodySequence
var s2sdk_GetBodySequence func(entityHandle int32) int32
var S2sdk_GetBodySequence = &s2sdk_GetBodySequence

//go:linkname s2sdk_IsBodyAttachedToParent __package__/s2sdk._IsBodyAttachedToParent
var s2sdk_IsBodyAttachedToParent func(entityHandle int32) bool
var S2sdk_IsBodyAttachedToParent = &s2sdk_IsBodyAttachedToParent

//go:linkname s2sdk_LookupBodySequence __package__/s2sdk._LookupBodySequence
var s2sdk_LookupBodySequence func(entityHandle int32, name string) int32
var S2sdk_LookupBodySequence = &s2sdk_LookupBodySequence

//go:linkname s2sdk_SetBodySequenceDuration __package__/s2sdk._SetBodySequenceDuration
var s2sdk_SetBodySequenceDuration func(entityHandle int32, sequenceName string) float32
var S2sdk_SetBodySequenceDuration = &s2sdk_SetBodySequenceDuration

//go:linkname s2sdk_SetBodyAngularVelocity __package__/s2sdk._SetBodyAngularVelocity
var s2sdk_SetBodyAngularVelocity func(entityHandle int32, angVelocity plugify.Vector3)
var S2sdk_SetBodyAngularVelocity = &s2sdk_SetBodyAngularVelocity

//go:linkname s2sdk_SetBodyMaterialGroup __package__/s2sdk._SetBodyMaterialGroup
var s2sdk_SetBodyMaterialGroup func(entityHandle int32, materialGroup string)
var S2sdk_SetBodyMaterialGroup = &s2sdk_SetBodyMaterialGroup

//go:linkname s2sdk_SetBodyVelocity __package__/s2sdk._SetBodyVelocity
var s2sdk_SetBodyVelocity func(entityHandle int32, velocity plugify.Vector3)
var S2sdk_SetBodyVelocity = &s2sdk_SetBodyVelocity

//go:linkname s2sdk_EntPointerToPlayerSlot __package__/s2sdk._EntPointerToPlayerSlot
var s2sdk_EntPointerToPlayerSlot func(entity uintptr) int32
var S2sdk_EntPointerToPlayerSlot = &s2sdk_EntPointerToPlayerSlot

//go:linkname s2sdk_PlayerSlotToEntPointer __package__/s2sdk._PlayerSlotToEntPointer
var s2sdk_PlayerSlotToEntPointer func(playerSlot int32) uintptr
var S2sdk_PlayerSlotToEntPointer = &s2sdk_PlayerSlotToEntPointer

//go:linkname s2sdk_PlayerSlotToEntHandle __package__/s2sdk._PlayerSlotToEntHandle
var s2sdk_PlayerSlotToEntHandle func(playerSlot int32) int32
var S2sdk_PlayerSlotToEntHandle = &s2sdk_PlayerSlotToEntHandle

//go:linkname s2sdk_PlayerSlotToClientPtr __package__/s2sdk._PlayerSlotToClientPtr
var s2sdk_PlayerSlotToClientPtr func(playerSlot int32) uintptr
var S2sdk_PlayerSlotToClientPtr = &s2sdk_PlayerSlotToClientPtr

//go:linkname s2sdk_ClientPtrToPlayerSlot __package__/s2sdk._ClientPtrToPlayerSlot
var s2sdk_ClientPtrToPlayerSlot func(client uintptr) int32
var S2sdk_ClientPtrToPlayerSlot = &s2sdk_ClientPtrToPlayerSlot

//go:linkname s2sdk_PlayerSlotToClientIndex __package__/s2sdk._PlayerSlotToClientIndex
var s2sdk_PlayerSlotToClientIndex func(playerSlot int32) int32
var S2sdk_PlayerSlotToClientIndex = &s2sdk_PlayerSlotToClientIndex

//go:linkname s2sdk_ClientIndexToPlayerSlot __package__/s2sdk._ClientIndexToPlayerSlot
var s2sdk_ClientIndexToPlayerSlot func(clientIndex int32) int32
var S2sdk_ClientIndexToPlayerSlot = &s2sdk_ClientIndexToPlayerSlot

//go:linkname s2sdk_PlayerServicesToPlayerSlot __package__/s2sdk._PlayerServicesToPlayerSlot
var s2sdk_PlayerServicesToPlayerSlot func(service uintptr) int32
var S2sdk_PlayerServicesToPlayerSlot = &s2sdk_PlayerServicesToPlayerSlot

//go:linkname s2sdk_GetClientAuthId __package__/s2sdk._GetClientAuthId
var s2sdk_GetClientAuthId func(playerSlot int32) string
var S2sdk_GetClientAuthId = &s2sdk_GetClientAuthId

//go:linkname s2sdk_GetClientAccountId __package__/s2sdk._GetClientAccountId
var s2sdk_GetClientAccountId func(playerSlot int32) uint32
var S2sdk_GetClientAccountId = &s2sdk_GetClientAccountId

//go:linkname s2sdk_GetClientSteamID64 __package__/s2sdk._GetClientSteamID64
var s2sdk_GetClientSteamID64 func(playerSlot int32) uint64
var S2sdk_GetClientSteamID64 = &s2sdk_GetClientSteamID64

//go:linkname s2sdk_GetClientIp __package__/s2sdk._GetClientIp
var s2sdk_GetClientIp func(playerSlot int32) string
var S2sdk_GetClientIp = &s2sdk_GetClientIp

//go:linkname s2sdk_GetClientLanguage __package__/s2sdk._GetClientLanguage
var s2sdk_GetClientLanguage func(playerSlot int32) string
var S2sdk_GetClientLanguage = &s2sdk_GetClientLanguage

//go:linkname s2sdk_GetClientOS __package__/s2sdk._GetClientOS
var s2sdk_GetClientOS func(playerSlot int32) string
var S2sdk_GetClientOS = &s2sdk_GetClientOS

//go:linkname s2sdk_GetClientName __package__/s2sdk._GetClientName
var s2sdk_GetClientName func(playerSlot int32) string
var S2sdk_GetClientName = &s2sdk_GetClientName

//go:linkname s2sdk_GetClientTime __package__/s2sdk._GetClientTime
var s2sdk_GetClientTime func(playerSlot int32) float32
var S2sdk_GetClientTime = &s2sdk_GetClientTime

//go:linkname s2sdk_GetClientLatency __package__/s2sdk._GetClientLatency
var s2sdk_GetClientLatency func(playerSlot int32) float32
var S2sdk_GetClientLatency = &s2sdk_GetClientLatency

//go:linkname s2sdk_GetUserFlagBits __package__/s2sdk._GetUserFlagBits
var s2sdk_GetUserFlagBits func(playerSlot int32) uint64
var S2sdk_GetUserFlagBits = &s2sdk_GetUserFlagBits

//go:linkname s2sdk_SetUserFlagBits __package__/s2sdk._SetUserFlagBits
var s2sdk_SetUserFlagBits func(playerSlot int32, flags uint64)
var S2sdk_SetUserFlagBits = &s2sdk_SetUserFlagBits

//go:linkname s2sdk_AddUserFlags __package__/s2sdk._AddUserFlags
var s2sdk_AddUserFlags func(playerSlot int32, flags uint64)
var S2sdk_AddUserFlags = &s2sdk_AddUserFlags

//go:linkname s2sdk_RemoveUserFlags __package__/s2sdk._RemoveUserFlags
var s2sdk_RemoveUserFlags func(playerSlot int32, flags uint64)
var S2sdk_RemoveUserFlags = &s2sdk_RemoveUserFlags

//go:linkname s2sdk_IsClientAuthorized __package__/s2sdk._IsClientAuthorized
var s2sdk_IsClientAuthorized func(playerSlot int32) bool
var S2sdk_IsClientAuthorized = &s2sdk_IsClientAuthorized

//go:linkname s2sdk_IsClientConnected __package__/s2sdk._IsClientConnected
var s2sdk_IsClientConnected func(playerSlot int32) bool
var S2sdk_IsClientConnected = &s2sdk_IsClientConnected

//go:linkname s2sdk_IsClientInGame __package__/s2sdk._IsClientInGame
var s2sdk_IsClientInGame func(playerSlot int32) bool
var S2sdk_IsClientInGame = &s2sdk_IsClientInGame

//go:linkname s2sdk_IsClientSourceTV __package__/s2sdk._IsClientSourceTV
var s2sdk_IsClientSourceTV func(playerSlot int32) bool
var S2sdk_IsClientSourceTV = &s2sdk_IsClientSourceTV

//go:linkname s2sdk_IsClientAlive __package__/s2sdk._IsClientAlive
var s2sdk_IsClientAlive func(playerSlot int32) bool
var S2sdk_IsClientAlive = &s2sdk_IsClientAlive

//go:linkname s2sdk_IsFakeClient __package__/s2sdk._IsFakeClient
var s2sdk_IsFakeClient func(playerSlot int32) bool
var S2sdk_IsFakeClient = &s2sdk_IsFakeClient

//go:linkname s2sdk_GetClientMoveType __package__/s2sdk._GetClientMoveType
var s2sdk_GetClientMoveType func(playerSlot int32) s2sdk.MoveType
var S2sdk_GetClientMoveType = &s2sdk_GetClientMoveType

//go:linkname s2sdk_SetClientMoveType __package__/s2sdk._SetClientMoveType
var s2sdk_SetClientMoveType func(playerSlot int32, moveType s2sdk.MoveType)
var S2sdk_SetClientMoveType = &s2sdk_SetClientMoveType

//go:linkname s2sdk_GetClientGravity __package__/s2sdk._GetClientGravity
var s2sdk_GetClientGravity func(playerSlot int32) float32
var S2sdk_GetClientGravity = &s2sdk_GetClientGravity

//go:linkname s2sdk_SetClientGravity __package__/s2sdk._SetClientGravity
var s2sdk_SetClientGravity func(playerSlot int32, gravity float32)
var S2sdk_SetClientGravity = &s2sdk_SetClientGravity

//go:linkname s2sdk_GetClientFlags __package__/s2sdk._GetClientFlags
var s2sdk_GetClientFlags func(playerSlot int32) int32
var S2sdk_GetClientFlags = &s2sdk_GetClientFlags

//go:linkname s2sdk_SetClientFlags __package__/s2sdk._SetClientFlags
var s2sdk_SetClientFlags func(playerSlot int32, flags int32)
var S2sdk_SetClientFlags = &s2sdk_SetClientFlags

//go:linkname s2sdk_GetClientRenderColor __package__/s2sdk._GetClientRenderColor
var s2sdk_GetClientRenderColor func(playerSlot int32) plugify.Vector4
var S2sdk_GetClientRenderColor = &s2sdk_GetClientRenderColor

//go:linkname s2sdk_SetClientRenderColor __package__/s2sdk._SetClientRenderColor
var s2sdk_SetClientRenderColor func(playerSlot int32, color plugify.Vector4)
var S2sdk_SetClientRenderColor = &s2sdk_SetClientRenderColor

//go:linkname s2sdk_GetClientRenderMode __package__/s2sdk._GetClientRenderMode
var s2sdk_GetClientRenderMode func(playerSlot int32) s2sdk.RenderMode
var S2sdk_GetClientRenderMode = &s2sdk_GetClientRenderMode

//go:linkname s2sdk_SetClientRenderMode __package__/s2sdk._SetClientRenderMode
var s2sdk_SetClientRenderMode func(playerSlot int32, renderMode s2sdk.RenderMode)
var S2sdk_SetClientRenderMode = &s2sdk_SetClientRenderMode

//go:linkname s2sdk_GetClientMass __package__/s2sdk._GetClientMass
var s2sdk_GetClientMass func(playerSlot int32) int32
var S2sdk_GetClientMass = &s2sdk_GetClientMass

//go:linkname s2sdk_SetClientMass __package__/s2sdk._SetClientMass
var s2sdk_SetClientMass func(playerSlot int32, mass int32)
var S2sdk_SetClientMass = &s2sdk_SetClientMass

//go:linkname s2sdk_GetClientFriction __package__/s2sdk._GetClientFriction
var s2sdk_GetClientFriction func(playerSlot int32) float32
var S2sdk_GetClientFriction = &s2sdk_GetClientFriction

//go:linkname s2sdk_SetClientFriction __package__/s2sdk._SetClientFriction
var s2sdk_SetClientFriction func(playerSlot int32, friction float32)
var S2sdk_SetClientFriction = &s2sdk_SetClientFriction

//go:linkname s2sdk_GetClientHealth __package__/s2sdk._GetClientHealth
var s2sdk_GetClientHealth func(playerSlot int32) int32
var S2sdk_GetClientHealth = &s2sdk_GetClientHealth

//go:linkname s2sdk_SetClientHealth __package__/s2sdk._SetClientHealth
var s2sdk_SetClientHealth func(playerSlot int32, health int32)
var S2sdk_SetClientHealth = &s2sdk_SetClientHealth

//go:linkname s2sdk_GetClientMaxHealth __package__/s2sdk._GetClientMaxHealth
var s2sdk_GetClientMaxHealth func(playerSlot int32) int32
var S2sdk_GetClientMaxHealth = &s2sdk_GetClientMaxHealth

//go:linkname s2sdk_SetClientMaxHealth __package__/s2sdk._SetClientMaxHealth
var s2sdk_SetClientMaxHealth func(playerSlot int32, maxHealth int32)
var S2sdk_SetClientMaxHealth = &s2sdk_SetClientMaxHealth

//go:linkname s2sdk_GetClientTeam __package__/s2sdk._GetClientTeam
var s2sdk_GetClientTeam func(playerSlot int32) s2sdk.CSTeam
var S2sdk_GetClientTeam = &s2sdk_GetClientTeam

//go:linkname s2sdk_SetClientTeam __package__/s2sdk._SetClientTeam
var s2sdk_SetClientTeam func(playerSlot int32, team s2sdk.CSTeam)
var S2sdk_SetClientTeam = &s2sdk_SetClientTeam

//go:linkname s2sdk_GetClientAbsOrigin __package__/s2sdk._GetClientAbsOrigin
var s2sdk_GetClientAbsOrigin func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientAbsOrigin = &s2sdk_GetClientAbsOrigin

//go:linkname s2sdk_SetClientAbsOrigin __package__/s2sdk._SetClientAbsOrigin
var s2sdk_SetClientAbsOrigin func(playerSlot int32, origin plugify.Vector3)
var S2sdk_SetClientAbsOrigin = &s2sdk_SetClientAbsOrigin

//go:linkname s2sdk_GetClientAbsScale __package__/s2sdk._GetClientAbsScale
var s2sdk_GetClientAbsScale func(playerSlot int32) float32
var S2sdk_GetClientAbsScale = &s2sdk_GetClientAbsScale

//go:linkname s2sdk_SetClientAbsScale __package__/s2sdk._SetClientAbsScale
var s2sdk_SetClientAbsScale func(playerSlot int32, scale float32)
var S2sdk_SetClientAbsScale = &s2sdk_SetClientAbsScale

//go:linkname s2sdk_GetClientAbsAngles __package__/s2sdk._GetClientAbsAngles
var s2sdk_GetClientAbsAngles func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientAbsAngles = &s2sdk_GetClientAbsAngles

//go:linkname s2sdk_SetClientAbsAngles __package__/s2sdk._SetClientAbsAngles
var s2sdk_SetClientAbsAngles func(playerSlot int32, angle plugify.Vector3)
var S2sdk_SetClientAbsAngles = &s2sdk_SetClientAbsAngles

//go:linkname s2sdk_GetClientLocalOrigin __package__/s2sdk._GetClientLocalOrigin
var s2sdk_GetClientLocalOrigin func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientLocalOrigin = &s2sdk_GetClientLocalOrigin

//go:linkname s2sdk_SetClientLocalOrigin __package__/s2sdk._SetClientLocalOrigin
var s2sdk_SetClientLocalOrigin func(playerSlot int32, origin plugify.Vector3)
var S2sdk_SetClientLocalOrigin = &s2sdk_SetClientLocalOrigin

//go:linkname s2sdk_GetClientLocalScale __package__/s2sdk._GetClientLocalScale
var s2sdk_GetClientLocalScale func(playerSlot int32) float32
var S2sdk_GetClientLocalScale = &s2sdk_GetClientLocalScale

//go:linkname s2sdk_SetClientLocalScale __package__/s2sdk._SetClientLocalScale
var s2sdk_SetClientLocalScale func(playerSlot int32, scale float32)
var S2sdk_SetClientLocalScale = &s2sdk_SetClientLocalScale

//go:linkname s2sdk_GetClientLocalAngles __package__/s2sdk._GetClientLocalAngles
var s2sdk_GetClientLocalAngles func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientLocalAngles = &s2sdk_GetClientLocalAngles

//go:linkname s2sdk_SetClientLocalAngles __package__/s2sdk._SetClientLocalAngles
var s2sdk_SetClientLocalAngles func(playerSlot int32, angle plugify.Vector3)
var S2sdk_SetClientLocalAngles = &s2sdk_SetClientLocalAngles

//go:linkname s2sdk_GetClientAbsVelocity __package__/s2sdk._GetClientAbsVelocity
var s2sdk_GetClientAbsVelocity func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientAbsVelocity = &s2sdk_GetClientAbsVelocity

//go:linkname s2sdk_SetClientAbsVelocity __package__/s2sdk._SetClientAbsVelocity
var s2sdk_SetClientAbsVelocity func(playerSlot int32, velocity plugify.Vector3)
var S2sdk_SetClientAbsVelocity = &s2sdk_SetClientAbsVelocity

//go:linkname s2sdk_GetClientBaseVelocity __package__/s2sdk._GetClientBaseVelocity
var s2sdk_GetClientBaseVelocity func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientBaseVelocity = &s2sdk_GetClientBaseVelocity

//go:linkname s2sdk_GetClientLocalAngVelocity __package__/s2sdk._GetClientLocalAngVelocity
var s2sdk_GetClientLocalAngVelocity func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientLocalAngVelocity = &s2sdk_GetClientLocalAngVelocity

//go:linkname s2sdk_GetClientAngVelocity __package__/s2sdk._GetClientAngVelocity
var s2sdk_GetClientAngVelocity func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientAngVelocity = &s2sdk_GetClientAngVelocity

//go:linkname s2sdk_SetClientAngVelocity __package__/s2sdk._SetClientAngVelocity
var s2sdk_SetClientAngVelocity func(playerSlot int32, velocity plugify.Vector3)
var S2sdk_SetClientAngVelocity = &s2sdk_SetClientAngVelocity

//go:linkname s2sdk_GetClientLocalVelocity __package__/s2sdk._GetClientLocalVelocity
var s2sdk_GetClientLocalVelocity func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientLocalVelocity = &s2sdk_GetClientLocalVelocity

//go:linkname s2sdk_GetClientAngRotation __package__/s2sdk._GetClientAngRotation
var s2sdk_GetClientAngRotation func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientAngRotation = &s2sdk_GetClientAngRotation

//go:linkname s2sdk_SetClientAngRotation __package__/s2sdk._SetClientAngRotation
var s2sdk_SetClientAngRotation func(playerSlot int32, rotation plugify.Vector3)
var S2sdk_SetClientAngRotation = &s2sdk_SetClientAngRotation

//go:linkname s2sdk_TransformPointClientToWorld __package__/s2sdk._TransformPointClientToWorld
var s2sdk_TransformPointClientToWorld func(playerSlot int32, point plugify.Vector3) plugify.Vector3
var S2sdk_TransformPointClientToWorld = &s2sdk_TransformPointClientToWorld

//go:linkname s2sdk_TransformPointWorldToClient __package__/s2sdk._TransformPointWorldToClient
var s2sdk_TransformPointWorldToClient func(playerSlot int32, point plugify.Vector3) plugify.Vector3
var S2sdk_TransformPointWorldToClient = &s2sdk_TransformPointWorldToClient

//go:linkname s2sdk_GetClientEyePosition __package__/s2sdk._GetClientEyePosition
var s2sdk_GetClientEyePosition func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientEyePosition = &s2sdk_GetClientEyePosition

//go:linkname s2sdk_GetClientEyeAngles __package__/s2sdk._GetClientEyeAngles
var s2sdk_GetClientEyeAngles func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientEyeAngles = &s2sdk_GetClientEyeAngles

//go:linkname s2sdk_SetClientForwardVector __package__/s2sdk._SetClientForwardVector
var s2sdk_SetClientForwardVector func(playerSlot int32, forward plugify.Vector3)
var S2sdk_SetClientForwardVector = &s2sdk_SetClientForwardVector

//go:linkname s2sdk_GetClientForwardVector __package__/s2sdk._GetClientForwardVector
var s2sdk_GetClientForwardVector func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientForwardVector = &s2sdk_GetClientForwardVector

//go:linkname s2sdk_GetClientLeftVector __package__/s2sdk._GetClientLeftVector
var s2sdk_GetClientLeftVector func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientLeftVector = &s2sdk_GetClientLeftVector

//go:linkname s2sdk_GetClientRightVector __package__/s2sdk._GetClientRightVector
var s2sdk_GetClientRightVector func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientRightVector = &s2sdk_GetClientRightVector

//go:linkname s2sdk_GetClientUpVector __package__/s2sdk._GetClientUpVector
var s2sdk_GetClientUpVector func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientUpVector = &s2sdk_GetClientUpVector

//go:linkname s2sdk_GetClientTransform __package__/s2sdk._GetClientTransform
var s2sdk_GetClientTransform func(playerSlot int32) plugify.Matrix4x4
var S2sdk_GetClientTransform = &s2sdk_GetClientTransform

//go:linkname s2sdk_GetClientModel __package__/s2sdk._GetClientModel
var s2sdk_GetClientModel func(playerSlot int32) string
var S2sdk_GetClientModel = &s2sdk_GetClientModel

//go:linkname s2sdk_SetClientModel __package__/s2sdk._SetClientModel
var s2sdk_SetClientModel func(playerSlot int32, model string)
var S2sdk_SetClientModel = &s2sdk_SetClientModel

//go:linkname s2sdk_GetClientWaterLevel __package__/s2sdk._GetClientWaterLevel
var s2sdk_GetClientWaterLevel func(playerSlot int32) float32
var S2sdk_GetClientWaterLevel = &s2sdk_GetClientWaterLevel

//go:linkname s2sdk_GetClientGroundEntity __package__/s2sdk._GetClientGroundEntity
var s2sdk_GetClientGroundEntity func(playerSlot int32) int32
var S2sdk_GetClientGroundEntity = &s2sdk_GetClientGroundEntity

//go:linkname s2sdk_GetClientEffects __package__/s2sdk._GetClientEffects
var s2sdk_GetClientEffects func(playerSlot int32) int32
var S2sdk_GetClientEffects = &s2sdk_GetClientEffects

//go:linkname s2sdk_AddClientEffects __package__/s2sdk._AddClientEffects
var s2sdk_AddClientEffects func(playerSlot int32, effects int32)
var S2sdk_AddClientEffects = &s2sdk_AddClientEffects

//go:linkname s2sdk_RemoveClientEffects __package__/s2sdk._RemoveClientEffects
var s2sdk_RemoveClientEffects func(playerSlot int32, effects int32)
var S2sdk_RemoveClientEffects = &s2sdk_RemoveClientEffects

//go:linkname s2sdk_GetClientBoundingMaxs __package__/s2sdk._GetClientBoundingMaxs
var s2sdk_GetClientBoundingMaxs func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientBoundingMaxs = &s2sdk_GetClientBoundingMaxs

//go:linkname s2sdk_GetClientBoundingMins __package__/s2sdk._GetClientBoundingMins
var s2sdk_GetClientBoundingMins func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientBoundingMins = &s2sdk_GetClientBoundingMins

//go:linkname s2sdk_GetClientCenter __package__/s2sdk._GetClientCenter
var s2sdk_GetClientCenter func(playerSlot int32) plugify.Vector3
var S2sdk_GetClientCenter = &s2sdk_GetClientCenter

//go:linkname s2sdk_TeleportClient __package__/s2sdk._TeleportClient
var s2sdk_TeleportClient func(playerSlot int32, origin plugify.Vector3, angles plugify.Vector3, velocity plugify.Vector3)
var S2sdk_TeleportClient = &s2sdk_TeleportClient

//go:linkname s2sdk_ApplyAbsVelocityImpulseToClient __package__/s2sdk._ApplyAbsVelocityImpulseToClient
var s2sdk_ApplyAbsVelocityImpulseToClient func(playerSlot int32, vecImpulse plugify.Vector3)
var S2sdk_ApplyAbsVelocityImpulseToClient = &s2sdk_ApplyAbsVelocityImpulseToClient

//go:linkname s2sdk_ApplyLocalAngularVelocityImpulseToClient __package__/s2sdk._ApplyLocalAngularVelocityImpulseToClient
var s2sdk_ApplyLocalAngularVelocityImpulseToClient func(playerSlot int32, angImpulse plugify.Vector3)
var S2sdk_ApplyLocalAngularVelocityImpulseToClient = &s2sdk_ApplyLocalAngularVelocityImpulseToClient

//go:linkname s2sdk_AcceptClientInput __package__/s2sdk._AcceptClientInput
var s2sdk_AcceptClientInput func(playerSlot int32, inputName string, activatorHandle int32, callerHandle int32, value any, type_ s2sdk.FieldType, outputId int32)
var S2sdk_AcceptClientInput = &s2sdk_AcceptClientInput

//go:linkname s2sdk_ConnectClientOutput __package__/s2sdk._ConnectClientOutput
var s2sdk_ConnectClientOutput func(playerSlot int32, output string, functionName string)
var S2sdk_ConnectClientOutput = &s2sdk_ConnectClientOutput

//go:linkname s2sdk_DisconnectClientOutput __package__/s2sdk._DisconnectClientOutput
var s2sdk_DisconnectClientOutput func(playerSlot int32, output string, functionName string)
var S2sdk_DisconnectClientOutput = &s2sdk_DisconnectClientOutput

//go:linkname s2sdk_DisconnectClientRedirectedOutput __package__/s2sdk._DisconnectClientRedirectedOutput
var s2sdk_DisconnectClientRedirectedOutput func(playerSlot int32, output string, functionName string, targetHandle int32)
var S2sdk_DisconnectClientRedirectedOutput = &s2sdk_DisconnectClientRedirectedOutput

//go:linkname s2sdk_FireClientOutput __package__/s2sdk._FireClientOutput
var s2sdk_FireClientOutput func(playerSlot int32, outputName string, activatorHandle int32, callerHandle int32, value any, type_ s2sdk.FieldType, delay float32)
var S2sdk_FireClientOutput = &s2sdk_FireClientOutput

//go:linkname s2sdk_RedirectClientOutput __package__/s2sdk._RedirectClientOutput
var s2sdk_RedirectClientOutput func(playerSlot int32, output string, functionName string, targetHandle int32)
var S2sdk_RedirectClientOutput = &s2sdk_RedirectClientOutput

//go:linkname s2sdk_FollowClient __package__/s2sdk._FollowClient
var s2sdk_FollowClient func(playerSlot int32, attachmentHandle int32, boneMerge bool)
var S2sdk_FollowClient = &s2sdk_FollowClient

//go:linkname s2sdk_FollowClientMerge __package__/s2sdk._FollowClientMerge
var s2sdk_FollowClientMerge func(playerSlot int32, attachmentHandle int32, boneOrAttachName string)
var S2sdk_FollowClientMerge = &s2sdk_FollowClientMerge

//go:linkname s2sdk_TakeClientDamage __package__/s2sdk._TakeClientDamage
var s2sdk_TakeClientDamage func(playerSlot int32, inflictorSlot int32, attackerSlot int32, force plugify.Vector3, hitPos plugify.Vector3, damage float32, damageTypes s2sdk.DamageTypes) int32
var S2sdk_TakeClientDamage = &s2sdk_TakeClientDamage

//go:linkname s2sdk_GetClientPawn __package__/s2sdk._GetClientPawn
var s2sdk_GetClientPawn func(playerSlot int32) uintptr
var S2sdk_GetClientPawn = &s2sdk_GetClientPawn

//go:linkname s2sdk_ProcessTargetString __package__/s2sdk._ProcessTargetString
var s2sdk_ProcessTargetString func(caller int32, target string) []int32
var S2sdk_ProcessTargetString = &s2sdk_ProcessTargetString

//go:linkname s2sdk_SwitchClientTeam __package__/s2sdk._SwitchClientTeam
var s2sdk_SwitchClientTeam func(playerSlot int32, team s2sdk.CSTeam)
var S2sdk_SwitchClientTeam = &s2sdk_SwitchClientTeam

//go:linkname s2sdk_ChangeClientTeam __package__/s2sdk._ChangeClientTeam
var s2sdk_ChangeClientTeam func(playerSlot int32, team s2sdk.CSTeam)
var S2sdk_ChangeClientTeam = &s2sdk_ChangeClientTeam

//go:linkname s2sdk_RespawnClient __package__/s2sdk._RespawnClient
var s2sdk_RespawnClient func(playerSlot int32)
var S2sdk_RespawnClient = &s2sdk_RespawnClient

//go:linkname s2sdk_ForcePlayerSuicide __package__/s2sdk._ForcePlayerSuicide
var s2sdk_ForcePlayerSuicide func(playerSlot int32, explode bool, force bool)
var S2sdk_ForcePlayerSuicide = &s2sdk_ForcePlayerSuicide

//go:linkname s2sdk_KickClient __package__/s2sdk._KickClient
var s2sdk_KickClient func(playerSlot int32, reason s2sdk.NetworkDisconnectionReason, message string)
var S2sdk_KickClient = &s2sdk_KickClient

//go:linkname s2sdk_BanClient __package__/s2sdk._BanClient
var s2sdk_BanClient func(playerSlot int32, duration float32, kick bool)
var S2sdk_BanClient = &s2sdk_BanClient

//go:linkname s2sdk_BanIdentity __package__/s2sdk._BanIdentity
var s2sdk_BanIdentity func(steamId uint64, duration float32, kick bool)
var S2sdk_BanIdentity = &s2sdk_BanIdentity

//go:linkname s2sdk_GetClientActiveWeapon __package__/s2sdk._GetClientActiveWeapon
var s2sdk_GetClientActiveWeapon func(playerSlot int32) int32
var S2sdk_GetClientActiveWeapon = &s2sdk_GetClientActiveWeapon

//go:linkname s2sdk_GetClientWeapons __package__/s2sdk._GetClientWeapons
var s2sdk_GetClientWeapons func(playerSlot int32) []int32
var S2sdk_GetClientWeapons = &s2sdk_GetClientWeapons

//go:linkname s2sdk_RemoveWeapons __package__/s2sdk._RemoveWeapons
var s2sdk_RemoveWeapons func(playerSlot int32, removeSuit bool)
var S2sdk_RemoveWeapons = &s2sdk_RemoveWeapons

//go:linkname s2sdk_DropWeapon __package__/s2sdk._DropWeapon
var s2sdk_DropWeapon func(playerSlot int32, weaponHandle int32, target plugify.Vector3, velocity plugify.Vector3)
var S2sdk_DropWeapon = &s2sdk_DropWeapon

//go:linkname s2sdk_SelectWeapon __package__/s2sdk._SelectWeapon
var s2sdk_SelectWeapon func(playerSlot int32, weaponHandle int32)
var S2sdk_SelectWeapon = &s2sdk_SelectWeapon

//go:linkname s2sdk_SwitchWeapon __package__/s2sdk._SwitchWeapon
var s2sdk_SwitchWeapon func(playerSlot int32, weaponHandle int32)
var S2sdk_SwitchWeapon = &s2sdk_SwitchWeapon

//go:linkname s2sdk_RemoveWeapon __package__/s2sdk._RemoveWeapon
var s2sdk_RemoveWeapon func(playerSlot int32, weaponHandle int32)
var S2sdk_RemoveWeapon = &s2sdk_RemoveWeapon

//go:linkname s2sdk_GiveNamedItem __package__/s2sdk._GiveNamedItem
var s2sdk_GiveNamedItem func(playerSlot int32, itemName string) int32
var S2sdk_GiveNamedItem = &s2sdk_GiveNamedItem

//go:linkname s2sdk_GetClientButtons __package__/s2sdk._GetClientButtons
var s2sdk_GetClientButtons func(playerSlot int32, buttonIndex int32) uint64
var S2sdk_GetClientButtons = &s2sdk_GetClientButtons

//go:linkname s2sdk_GetClientArmor __package__/s2sdk._GetClientArmor
var s2sdk_GetClientArmor func(playerSlot int32) int32
var S2sdk_GetClientArmor = &s2sdk_GetClientArmor

//go:linkname s2sdk_SetClientArmor __package__/s2sdk._SetClientArmor
var s2sdk_SetClientArmor func(playerSlot int32, armor int32)
var S2sdk_SetClientArmor = &s2sdk_SetClientArmor

//go:linkname s2sdk_GetClientSpeed __package__/s2sdk._GetClientSpeed
var s2sdk_GetClientSpeed func(playerSlot int32) float32
var S2sdk_GetClientSpeed = &s2sdk_GetClientSpeed

//go:linkname s2sdk_SetClientSpeed __package__/s2sdk._SetClientSpeed
var s2sdk_SetClientSpeed func(playerSlot int32, speed float32)
var S2sdk_SetClientSpeed = &s2sdk_SetClientSpeed

//go:linkname s2sdk_GetClientMoney __package__/s2sdk._GetClientMoney
var s2sdk_GetClientMoney func(playerSlot int32) int32
var S2sdk_GetClientMoney = &s2sdk_GetClientMoney

//go:linkname s2sdk_SetClientMoney __package__/s2sdk._SetClientMoney
var s2sdk_SetClientMoney func(playerSlot int32, money int32)
var S2sdk_SetClientMoney = &s2sdk_SetClientMoney

//go:linkname s2sdk_GetClientKills __package__/s2sdk._GetClientKills
var s2sdk_GetClientKills func(playerSlot int32) int32
var S2sdk_GetClientKills = &s2sdk_GetClientKills

//go:linkname s2sdk_SetClientKills __package__/s2sdk._SetClientKills
var s2sdk_SetClientKills func(playerSlot int32, kills int32)
var S2sdk_SetClientKills = &s2sdk_SetClientKills

//go:linkname s2sdk_GetClientDeaths __package__/s2sdk._GetClientDeaths
var s2sdk_GetClientDeaths func(playerSlot int32) int32
var S2sdk_GetClientDeaths = &s2sdk_GetClientDeaths

//go:linkname s2sdk_SetClientDeaths __package__/s2sdk._SetClientDeaths
var s2sdk_SetClientDeaths func(playerSlot int32, deaths int32)
var S2sdk_SetClientDeaths = &s2sdk_SetClientDeaths

//go:linkname s2sdk_GetClientAssists __package__/s2sdk._GetClientAssists
var s2sdk_GetClientAssists func(playerSlot int32) int32
var S2sdk_GetClientAssists = &s2sdk_GetClientAssists

//go:linkname s2sdk_SetClientAssists __package__/s2sdk._SetClientAssists
var s2sdk_SetClientAssists func(playerSlot int32, assists int32)
var S2sdk_SetClientAssists = &s2sdk_SetClientAssists

//go:linkname s2sdk_GetClientDamage __package__/s2sdk._GetClientDamage
var s2sdk_GetClientDamage func(playerSlot int32) int32
var S2sdk_GetClientDamage = &s2sdk_GetClientDamage

//go:linkname s2sdk_SetClientDamage __package__/s2sdk._SetClientDamage
var s2sdk_SetClientDamage func(playerSlot int32, damage int32)
var S2sdk_SetClientDamage = &s2sdk_SetClientDamage

//go:linkname s2sdk_AddAdminCommand __package__/s2sdk._AddAdminCommand
var s2sdk_AddAdminCommand func(name string, adminFlags int64, description string, flags s2sdk.ConVarFlag, callback s2sdk.ConCommandCallback, type_ s2sdk.HookMode) bool
var S2sdk_AddAdminCommand = &s2sdk_AddAdminCommand

//go:linkname s2sdk_AddConsoleCommand __package__/s2sdk._AddConsoleCommand
var s2sdk_AddConsoleCommand func(name string, description string, flags s2sdk.ConVarFlag, callback s2sdk.ConCommandCallback, type_ s2sdk.HookMode) bool
var S2sdk_AddConsoleCommand = &s2sdk_AddConsoleCommand

//go:linkname s2sdk_RemoveCommand __package__/s2sdk._RemoveCommand
var s2sdk_RemoveCommand func(name string, callback s2sdk.ConCommandCallback) bool
var S2sdk_RemoveCommand = &s2sdk_RemoveCommand

//go:linkname s2sdk_AddCommandListener __package__/s2sdk._AddCommandListener
var s2sdk_AddCommandListener func(name string, callback s2sdk.ConCommandCallback, type_ s2sdk.HookMode) bool
var S2sdk_AddCommandListener = &s2sdk_AddCommandListener

//go:linkname s2sdk_RemoveCommandListener __package__/s2sdk._RemoveCommandListener
var s2sdk_RemoveCommandListener func(name string, callback s2sdk.ConCommandCallback, type_ s2sdk.HookMode) bool
var S2sdk_RemoveCommandListener = &s2sdk_RemoveCommandListener

//go:linkname s2sdk_ServerCommand __package__/s2sdk._ServerCommand
var s2sdk_ServerCommand func(command string)
var S2sdk_ServerCommand = &s2sdk_ServerCommand

//go:linkname s2sdk_ServerCommandEx __package__/s2sdk._ServerCommandEx
var s2sdk_ServerCommandEx func(command string) string
var S2sdk_ServerCommandEx = &s2sdk_ServerCommandEx

//go:linkname s2sdk_ClientCommand __package__/s2sdk._ClientCommand
var s2sdk_ClientCommand func(playerSlot int32, command string)
var S2sdk_ClientCommand = &s2sdk_ClientCommand

//go:linkname s2sdk_FakeClientCommand __package__/s2sdk._FakeClientCommand
var s2sdk_FakeClientCommand func(playerSlot int32, command string)
var S2sdk_FakeClientCommand = &s2sdk_FakeClientCommand

//go:linkname s2sdk_GetAllConCommands __package__/s2sdk._GetAllConCommands
var s2sdk_GetAllConCommands func(flags s2sdk.ConVarFlag) []string
var S2sdk_GetAllConCommands = &s2sdk_GetAllConCommands

//go:linkname s2sdk_GetAllCommands __package__/s2sdk._GetAllCommands
var s2sdk_GetAllCommands func() []string
var S2sdk_GetAllCommands = &s2sdk_GetAllCommands

//go:linkname s2sdk_PrintToServer __package__/s2sdk._PrintToServer
var s2sdk_PrintToServer func(msg string)
var S2sdk_PrintToServer = &s2sdk_PrintToServer

//go:linkname s2sdk_PrintToConsole __package__/s2sdk._PrintToConsole
var s2sdk_PrintToConsole func(playerSlot int32, message string)
var S2sdk_PrintToConsole = &s2sdk_PrintToConsole

//go:linkname s2sdk_PrintToChat __package__/s2sdk._PrintToChat
var s2sdk_PrintToChat func(playerSlot int32, message string)
var S2sdk_PrintToChat = &s2sdk_PrintToChat

//go:linkname s2sdk_PrintCenterText __package__/s2sdk._PrintCenterText
var s2sdk_PrintCenterText func(playerSlot int32, message string)
var S2sdk_PrintCenterText = &s2sdk_PrintCenterText

//go:linkname s2sdk_PrintAlertText __package__/s2sdk._PrintAlertText
var s2sdk_PrintAlertText func(playerSlot int32, message string)
var S2sdk_PrintAlertText = &s2sdk_PrintAlertText

//go:linkname s2sdk_PrintCentreHtml __package__/s2sdk._PrintCentreHtml
var s2sdk_PrintCentreHtml func(playerSlot int32, message string, duration int32)
var S2sdk_PrintCentreHtml = &s2sdk_PrintCentreHtml

//go:linkname s2sdk_PrintToConsoleAll __package__/s2sdk._PrintToConsoleAll
var s2sdk_PrintToConsoleAll func(message string)
var S2sdk_PrintToConsoleAll = &s2sdk_PrintToConsoleAll

//go:linkname s2sdk_PrintToChatAll __package__/s2sdk._PrintToChatAll
var s2sdk_PrintToChatAll func(message string)
var S2sdk_PrintToChatAll = &s2sdk_PrintToChatAll

//go:linkname s2sdk_PrintCenterTextAll __package__/s2sdk._PrintCenterTextAll
var s2sdk_PrintCenterTextAll func(message string)
var S2sdk_PrintCenterTextAll = &s2sdk_PrintCenterTextAll

//go:linkname s2sdk_PrintAlertTextAll __package__/s2sdk._PrintAlertTextAll
var s2sdk_PrintAlertTextAll func(message string)
var S2sdk_PrintAlertTextAll = &s2sdk_PrintAlertTextAll

//go:linkname s2sdk_PrintCentreHtmlAll __package__/s2sdk._PrintCentreHtmlAll
var s2sdk_PrintCentreHtmlAll func(message string, duration int32)
var S2sdk_PrintCentreHtmlAll = &s2sdk_PrintCentreHtmlAll

//go:linkname s2sdk_PrintToChatColored __package__/s2sdk._PrintToChatColored
var s2sdk_PrintToChatColored func(playerSlot int32, message string)
var S2sdk_PrintToChatColored = &s2sdk_PrintToChatColored

//go:linkname s2sdk_PrintToChatColoredAll __package__/s2sdk._PrintToChatColoredAll
var s2sdk_PrintToChatColoredAll func(message string)
var S2sdk_PrintToChatColoredAll = &s2sdk_PrintToChatColoredAll

//go:linkname s2sdk_ReplyToCommand __package__/s2sdk._ReplyToCommand
var s2sdk_ReplyToCommand func(context s2sdk.ConCommandContext, playerSlot int32, message string)
var S2sdk_ReplyToCommand = &s2sdk_ReplyToCommand

//go:linkname s2sdk_CreateConVar __package__/s2sdk._CreateConVar
var s2sdk_CreateConVar func(name string, defaultValue any, description string, flags s2sdk.ConVarFlag) uint64
var S2sdk_CreateConVar = &s2sdk_CreateConVar

//go:linkname s2sdk_CreateConVarBool __package__/s2sdk._CreateConVarBool
var s2sdk_CreateConVarBool func(name string, defaultValue bool, description string, flags s2sdk.ConVarFlag, hasMin bool, min bool, hasMax bool, max bool) uint64
var S2sdk_CreateConVarBool = &s2sdk_CreateConVarBool

//go:linkname s2sdk_CreateConVarInt16 __package__/s2sdk._CreateConVarInt16
var s2sdk_CreateConVarInt16 func(name string, defaultValue int16, description string, flags s2sdk.ConVarFlag, hasMin bool, min int16, hasMax bool, max int16) uint64
var S2sdk_CreateConVarInt16 = &s2sdk_CreateConVarInt16

//go:linkname s2sdk_CreateConVarUInt16 __package__/s2sdk._CreateConVarUInt16
var s2sdk_CreateConVarUInt16 func(name string, defaultValue uint16, description string, flags s2sdk.ConVarFlag, hasMin bool, min uint16, hasMax bool, max uint16) uint64
var S2sdk_CreateConVarUInt16 = &s2sdk_CreateConVarUInt16

//go:linkname s2sdk_CreateConVarInt32 __package__/s2sdk._CreateConVarInt32
var s2sdk_CreateConVarInt32 func(name string, defaultValue int32, description string, flags s2sdk.ConVarFlag, hasMin bool, min int32, hasMax bool, max int32) uint64
var S2sdk_CreateConVarInt32 = &s2sdk_CreateConVarInt32

//go:linkname s2sdk_CreateConVarUInt32 __package__/s2sdk._CreateConVarUInt32
var s2sdk_CreateConVarUInt32 func(name string, defaultValue uint32, description string, flags s2sdk.ConVarFlag, hasMin bool, min uint32, hasMax bool, max uint32) uint64
var S2sdk_CreateConVarUInt32 = &s2sdk_CreateConVarUInt32

//go:linkname s2sdk_CreateConVarInt64 __package__/s2sdk._CreateConVarInt64
var s2sdk_CreateConVarInt64 func(name string, defaultValue int64, description string, flags s2sdk.ConVarFlag, hasMin bool, min int64, hasMax bool, max int64) uint64
var S2sdk_CreateConVarInt64 = &s2sdk_CreateConVarInt64

//go:linkname s2sdk_CreateConVarUInt64 __package__/s2sdk._CreateConVarUInt64
var s2sdk_CreateConVarUInt64 func(name string, defaultValue uint64, description string, flags s2sdk.ConVarFlag, hasMin bool, min uint64, hasMax bool, max uint64) uint64
var S2sdk_CreateConVarUInt64 = &s2sdk_CreateConVarUInt64

//go:linkname s2sdk_CreateConVarFloat __package__/s2sdk._CreateConVarFloat
var s2sdk_CreateConVarFloat func(name string, defaultValue float32, description string, flags s2sdk.ConVarFlag, hasMin bool, min float32, hasMax bool, max float32) uint64
var S2sdk_CreateConVarFloat = &s2sdk_CreateConVarFloat

//go:linkname s2sdk_CreateConVarDouble __package__/s2sdk._CreateConVarDouble
var s2sdk_CreateConVarDouble func(name string, defaultValue float64, description string, flags s2sdk.ConVarFlag, hasMin bool, min float64, hasMax bool, max float64) uint64
var S2sdk_CreateConVarDouble = &s2sdk_CreateConVarDouble

//go:linkname s2sdk_CreateConVarColor __package__/s2sdk._CreateConVarColor
var s2sdk_CreateConVarColor func(name string, defaultValue plugify.Vector4, description string, flags s2sdk.ConVarFlag, hasMin bool, min plugify.Vector4, hasMax bool, max plugify.Vector4) uint64
var S2sdk_CreateConVarColor = &s2sdk_CreateConVarColor

//go:linkname s2sdk_CreateConVarVector2 __package__/s2sdk._CreateConVarVector2
var s2sdk_CreateConVarVector2 func(name string, defaultValue plugify.Vector2, description string, flags s2sdk.ConVarFlag, hasMin bool, min plugify.Vector2, hasMax bool, max plugify.Vector2) uint64
var S2sdk_CreateConVarVector2 = &s2sdk_CreateConVarVector2

//go:linkname s2sdk_CreateConVarVector3 __package__/s2sdk._CreateConVarVector3
var s2sdk_CreateConVarVector3 func(name string, defaultValue plugify.Vector3, description string, flags s2sdk.ConVarFlag, hasMin bool, min plugify.Vector3, hasMax bool, max plugify.Vector3) uint64
var S2sdk_CreateConVarVector3 = &s2sdk_CreateConVarVector3

//go:linkname s2sdk_CreateConVarVector4 __package__/s2sdk._CreateConVarVector4
var s2sdk_CreateConVarVector4 func(name string, defaultValue plugify.Vector4, description string, flags s2sdk.ConVarFlag, hasMin bool, min plugify.Vector4, hasMax bool, max plugify.Vector4) uint64
var S2sdk_CreateConVarVector4 = &s2sdk_CreateConVarVector4

//go:linkname s2sdk_CreateConVarQAngle __package__/s2sdk._CreateConVarQAngle
var s2sdk_CreateConVarQAngle func(name string, defaultValue plugify.Vector3, description string, flags s2sdk.ConVarFlag, hasMin bool, min plugify.Vector3, hasMax bool, max plugify.Vector3) uint64
var S2sdk_CreateConVarQAngle = &s2sdk_CreateConVarQAngle

//go:linkname s2sdk_CreateConVarString __package__/s2sdk._CreateConVarString
var s2sdk_CreateConVarString func(name string, defaultValue string, description string, flags s2sdk.ConVarFlag) uint64
var S2sdk_CreateConVarString = &s2sdk_CreateConVarString

//go:linkname s2sdk_FindConVar __package__/s2sdk._FindConVar
var s2sdk_FindConVar func(name string) uint64
var S2sdk_FindConVar = &s2sdk_FindConVar

//go:linkname s2sdk_FindConVar2 __package__/s2sdk._FindConVar2
var s2sdk_FindConVar2 func(name string, type_ s2sdk.ConVarType) uint64
var S2sdk_FindConVar2 = &s2sdk_FindConVar2

//go:linkname s2sdk_HookConVarChange __package__/s2sdk._HookConVarChange
var s2sdk_HookConVarChange func(conVarHandle uint64, callback s2sdk.ConVarCallback)
var S2sdk_HookConVarChange = &s2sdk_HookConVarChange

//go:linkname s2sdk_UnhookConVarChange __package__/s2sdk._UnhookConVarChange
var s2sdk_UnhookConVarChange func(conVarHandle uint64, callback s2sdk.ConVarCallback)
var S2sdk_UnhookConVarChange = &s2sdk_UnhookConVarChange

//go:linkname s2sdk_IsConVarFlagSet __package__/s2sdk._IsConVarFlagSet
var s2sdk_IsConVarFlagSet func(conVarHandle uint64, flag int64) bool
var S2sdk_IsConVarFlagSet = &s2sdk_IsConVarFlagSet

//go:linkname s2sdk_AddConVarFlags __package__/s2sdk._AddConVarFlags
var s2sdk_AddConVarFlags func(conVarHandle uint64, flags s2sdk.ConVarFlag)
var S2sdk_AddConVarFlags = &s2sdk_AddConVarFlags

//go:linkname s2sdk_RemoveConVarFlags __package__/s2sdk._RemoveConVarFlags
var s2sdk_RemoveConVarFlags func(conVarHandle uint64, flags s2sdk.ConVarFlag)
var S2sdk_RemoveConVarFlags = &s2sdk_RemoveConVarFlags

//go:linkname s2sdk_GetConVarFlags __package__/s2sdk._GetConVarFlags
var s2sdk_GetConVarFlags func(conVarHandle uint64) s2sdk.ConVarFlag
var S2sdk_GetConVarFlags = &s2sdk_GetConVarFlags

//go:linkname s2sdk_GetConVarBounds __package__/s2sdk._GetConVarBounds
var s2sdk_GetConVarBounds func(conVarHandle uint64, max bool) string
var S2sdk_GetConVarBounds = &s2sdk_GetConVarBounds

//go:linkname s2sdk_SetConVarBounds __package__/s2sdk._SetConVarBounds
var s2sdk_SetConVarBounds func(conVarHandle uint64, max bool, value string)
var S2sdk_SetConVarBounds = &s2sdk_SetConVarBounds

//go:linkname s2sdk_GetConVarDefault __package__/s2sdk._GetConVarDefault
var s2sdk_GetConVarDefault func(conVarHandle uint64) string
var S2sdk_GetConVarDefault = &s2sdk_GetConVarDefault

//go:linkname s2sdk_GetConVarValue __package__/s2sdk._GetConVarValue
var s2sdk_GetConVarValue func(conVarHandle uint64) string
var S2sdk_GetConVarValue = &s2sdk_GetConVarValue

//go:linkname s2sdk_GetConVar __package__/s2sdk._GetConVar
var s2sdk_GetConVar func(conVarHandle uint64) any
var S2sdk_GetConVar = &s2sdk_GetConVar

//go:linkname s2sdk_GetConVarBool __package__/s2sdk._GetConVarBool
var s2sdk_GetConVarBool func(conVarHandle uint64) bool
var S2sdk_GetConVarBool = &s2sdk_GetConVarBool

//go:linkname s2sdk_GetConVarInt16 __package__/s2sdk._GetConVarInt16
var s2sdk_GetConVarInt16 func(conVarHandle uint64) int16
var S2sdk_GetConVarInt16 = &s2sdk_GetConVarInt16

//go:linkname s2sdk_GetConVarUInt16 __package__/s2sdk._GetConVarUInt16
var s2sdk_GetConVarUInt16 func(conVarHandle uint64) uint16
var S2sdk_GetConVarUInt16 = &s2sdk_GetConVarUInt16

//go:linkname s2sdk_GetConVarInt32 __package__/s2sdk._GetConVarInt32
var s2sdk_GetConVarInt32 func(conVarHandle uint64) int32
var S2sdk_GetConVarInt32 = &s2sdk_GetConVarInt32

//go:linkname s2sdk_GetConVarUInt32 __package__/s2sdk._GetConVarUInt32
var s2sdk_GetConVarUInt32 func(conVarHandle uint64) uint32
var S2sdk_GetConVarUInt32 = &s2sdk_GetConVarUInt32

//go:linkname s2sdk_GetConVarInt64 __package__/s2sdk._GetConVarInt64
var s2sdk_GetConVarInt64 func(conVarHandle uint64) int64
var S2sdk_GetConVarInt64 = &s2sdk_GetConVarInt64

//go:linkname s2sdk_GetConVarUInt64 __package__/s2sdk._GetConVarUInt64
var s2sdk_GetConVarUInt64 func(conVarHandle uint64) uint64
var S2sdk_GetConVarUInt64 = &s2sdk_GetConVarUInt64

//go:linkname s2sdk_GetConVarFloat __package__/s2sdk._GetConVarFloat
var s2sdk_GetConVarFloat func(conVarHandle uint64) float32
var S2sdk_GetConVarFloat = &s2sdk_GetConVarFloat

//go:linkname s2sdk_GetConVarDouble __package__/s2sdk._GetConVarDouble
var s2sdk_GetConVarDouble func(conVarHandle uint64) float64
var S2sdk_GetConVarDouble = &s2sdk_GetConVarDouble

//go:linkname s2sdk_GetConVarString __package__/s2sdk._GetConVarString
var s2sdk_GetConVarString func(conVarHandle uint64) string
var S2sdk_GetConVarString = &s2sdk_GetConVarString

//go:linkname s2sdk_GetConVarColor __package__/s2sdk._GetConVarColor
var s2sdk_GetConVarColor func(conVarHandle uint64) plugify.Vector4
var S2sdk_GetConVarColor = &s2sdk_GetConVarColor

//go:linkname s2sdk_GetConVarVector2 __package__/s2sdk._GetConVarVector2
var s2sdk_GetConVarVector2 func(conVarHandle uint64) plugify.Vector2
var S2sdk_GetConVarVector2 = &s2sdk_GetConVarVector2

//go:linkname s2sdk_GetConVarVector __package__/s2sdk._GetConVarVector
var s2sdk_GetConVarVector func(conVarHandle uint64) plugify.Vector3
var S2sdk_GetConVarVector = &s2sdk_GetConVarVector

//go:linkname s2sdk_GetConVarVector4 __package__/s2sdk._GetConVarVector4
var s2sdk_GetConVarVector4 func(conVarHandle uint64) plugify.Vector4
var S2sdk_GetConVarVector4 = &s2sdk_GetConVarVector4

//go:linkname s2sdk_GetConVarQAngle __package__/s2sdk._GetConVarQAngle
var s2sdk_GetConVarQAngle func(conVarHandle uint64) plugify.Vector3
var S2sdk_GetConVarQAngle = &s2sdk_GetConVarQAngle

//go:linkname s2sdk_SetConVarValue __package__/s2sdk._SetConVarValue
var s2sdk_SetConVarValue func(conVarHandle uint64, value string, replicate bool, notify bool)
var S2sdk_SetConVarValue = &s2sdk_SetConVarValue

//go:linkname s2sdk_SetConVar __package__/s2sdk._SetConVar
var s2sdk_SetConVar func(conVarHandle uint64, value any, replicate bool, notify bool)
var S2sdk_SetConVar = &s2sdk_SetConVar

//go:linkname s2sdk_SetConVarBool __package__/s2sdk._SetConVarBool
var s2sdk_SetConVarBool func(conVarHandle uint64, value bool, replicate bool, notify bool)
var S2sdk_SetConVarBool = &s2sdk_SetConVarBool

//go:linkname s2sdk_SetConVarInt16 __package__/s2sdk._SetConVarInt16
var s2sdk_SetConVarInt16 func(conVarHandle uint64, value int16, replicate bool, notify bool)
var S2sdk_SetConVarInt16 = &s2sdk_SetConVarInt16

//go:linkname s2sdk_SetConVarUInt16 __package__/s2sdk._SetConVarUInt16
var s2sdk_SetConVarUInt16 func(conVarHandle uint64, value uint16, replicate bool, notify bool)
var S2sdk_SetConVarUInt16 = &s2sdk_SetConVarUInt16

//go:linkname s2sdk_SetConVarInt32 __package__/s2sdk._SetConVarInt32
var s2sdk_SetConVarInt32 func(conVarHandle uint64, value int32, replicate bool, notify bool)
var S2sdk_SetConVarInt32 = &s2sdk_SetConVarInt32

//go:linkname s2sdk_SetConVarUInt32 __package__/s2sdk._SetConVarUInt32
var s2sdk_SetConVarUInt32 func(conVarHandle uint64, value uint32, replicate bool, notify bool)
var S2sdk_SetConVarUInt32 = &s2sdk_SetConVarUInt32

//go:linkname s2sdk_SetConVarInt64 __package__/s2sdk._SetConVarInt64
var s2sdk_SetConVarInt64 func(conVarHandle uint64, value int64, replicate bool, notify bool)
var S2sdk_SetConVarInt64 = &s2sdk_SetConVarInt64

//go:linkname s2sdk_SetConVarUInt64 __package__/s2sdk._SetConVarUInt64
var s2sdk_SetConVarUInt64 func(conVarHandle uint64, value uint64, replicate bool, notify bool)
var S2sdk_SetConVarUInt64 = &s2sdk_SetConVarUInt64

//go:linkname s2sdk_SetConVarFloat __package__/s2sdk._SetConVarFloat
var s2sdk_SetConVarFloat func(conVarHandle uint64, value float32, replicate bool, notify bool)
var S2sdk_SetConVarFloat = &s2sdk_SetConVarFloat

//go:linkname s2sdk_SetConVarDouble __package__/s2sdk._SetConVarDouble
var s2sdk_SetConVarDouble func(conVarHandle uint64, value float64, replicate bool, notify bool)
var S2sdk_SetConVarDouble = &s2sdk_SetConVarDouble

//go:linkname s2sdk_SetConVarString __package__/s2sdk._SetConVarString
var s2sdk_SetConVarString func(conVarHandle uint64, value string, replicate bool, notify bool)
var S2sdk_SetConVarString = &s2sdk_SetConVarString

//go:linkname s2sdk_SetConVarColor __package__/s2sdk._SetConVarColor
var s2sdk_SetConVarColor func(conVarHandle uint64, value plugify.Vector4, replicate bool, notify bool)
var S2sdk_SetConVarColor = &s2sdk_SetConVarColor

//go:linkname s2sdk_SetConVarVector2 __package__/s2sdk._SetConVarVector2
var s2sdk_SetConVarVector2 func(conVarHandle uint64, value plugify.Vector2, replicate bool, notify bool)
var S2sdk_SetConVarVector2 = &s2sdk_SetConVarVector2

//go:linkname s2sdk_SetConVarVector3 __package__/s2sdk._SetConVarVector3
var s2sdk_SetConVarVector3 func(conVarHandle uint64, value plugify.Vector3, replicate bool, notify bool)
var S2sdk_SetConVarVector3 = &s2sdk_SetConVarVector3

//go:linkname s2sdk_SetConVarVector4 __package__/s2sdk._SetConVarVector4
var s2sdk_SetConVarVector4 func(conVarHandle uint64, value plugify.Vector4, replicate bool, notify bool)
var S2sdk_SetConVarVector4 = &s2sdk_SetConVarVector4

//go:linkname s2sdk_SetConVarQAngle __package__/s2sdk._SetConVarQAngle
var s2sdk_SetConVarQAngle func(conVarHandle uint64, value plugify.Vector3, replicate bool, notify bool)
var S2sdk_SetConVarQAngle = &s2sdk_SetConVarQAngle

//go:linkname s2sdk_SendConVarValue __package__/s2sdk._SendConVarValue
var s2sdk_SendConVarValue func(playerSlot int32, conVarHandle uint64, value string)
var S2sdk_SendConVarValue = &s2sdk_SendConVarValue

//go:linkname s2sdk_SendConVarValue2 __package__/s2sdk._SendConVarValue2
var s2sdk_SendConVarValue2 func(conVarHandle uint64, playerSlot int32, value string)
var S2sdk_SendConVarValue2 = &s2sdk_SendConVarValue2

//go:linkname s2sdk_GetClientConVarValue __package__/s2sdk._GetClientConVarValue
var s2sdk_GetClientConVarValue func(playerSlot int32, convarName string) string
var S2sdk_GetClientConVarValue = &s2sdk_GetClientConVarValue

//go:linkname s2sdk_SetFakeClientConVarValue __package__/s2sdk._SetFakeClientConVarValue
var s2sdk_SetFakeClientConVarValue func(playerSlot int32, convarName string, convarValue string)
var S2sdk_SetFakeClientConVarValue = &s2sdk_SetFakeClientConVarValue

//go:linkname s2sdk_QueryClientConVar __package__/s2sdk._QueryClientConVar
var s2sdk_QueryClientConVar func(playerSlot int32, convarName string, callback s2sdk.CvarValueCallback, data []any) int32
var S2sdk_QueryClientConVar = &s2sdk_QueryClientConVar

//go:linkname s2sdk_AutoExecConfig __package__/s2sdk._AutoExecConfig
var s2sdk_AutoExecConfig func(conVarHandles []uint64, autoCreate bool, name string, folder string) bool
var S2sdk_AutoExecConfig = &s2sdk_AutoExecConfig

//go:linkname s2sdk_GetServerLanguage __package__/s2sdk._GetServerLanguage
var s2sdk_GetServerLanguage func() string
var S2sdk_GetServerLanguage = &s2sdk_GetServerLanguage

//go:linkname s2sdk_GetAllConVars __package__/s2sdk._GetAllConVars
var s2sdk_GetAllConVars func() []string
var S2sdk_GetAllConVars = &s2sdk_GetAllConVars

//go:linkname s2sdk_QueryInterface __package__/s2sdk._QueryInterface
var s2sdk_QueryInterface func(module string, name string) uintptr
var S2sdk_QueryInterface = &s2sdk_QueryInterface

//go:linkname s2sdk_GetGameDirectory __package__/s2sdk._GetGameDirectory
var s2sdk_GetGameDirectory func() string
var S2sdk_GetGameDirectory = &s2sdk_GetGameDirectory

//go:linkname s2sdk_ReadFileVPK __package__/s2sdk._ReadFileVPK
var s2sdk_ReadFileVPK func(localFileName string, pathId string) string
var S2sdk_ReadFileVPK = &s2sdk_ReadFileVPK

//go:linkname s2sdk_FindFileAbsoluteList __package__/s2sdk._FindFileAbsoluteList
var s2sdk_FindFileAbsoluteList func(wildcard string, pathId string) []string
var S2sdk_FindFileAbsoluteList = &s2sdk_FindFileAbsoluteList

//go:linkname s2sdk_GetCurrentMap __package__/s2sdk._GetCurrentMap
var s2sdk_GetCurrentMap func() string
var S2sdk_GetCurrentMap = &s2sdk_GetCurrentMap

//go:linkname s2sdk_IsMapValid __package__/s2sdk._IsMapValid
var s2sdk_IsMapValid func(mapname string) bool
var S2sdk_IsMapValid = &s2sdk_IsMapValid

//go:linkname s2sdk_GetGameTime __package__/s2sdk._GetGameTime
var s2sdk_GetGameTime func() float32
var S2sdk_GetGameTime = &s2sdk_GetGameTime

//go:linkname s2sdk_GetGameTickCount __package__/s2sdk._GetGameTickCount
var s2sdk_GetGameTickCount func() int32
var S2sdk_GetGameTickCount = &s2sdk_GetGameTickCount

//go:linkname s2sdk_GetGameFrameTime __package__/s2sdk._GetGameFrameTime
var s2sdk_GetGameFrameTime func() float32
var S2sdk_GetGameFrameTime = &s2sdk_GetGameFrameTime

//go:linkname s2sdk_GetEngineTime __package__/s2sdk._GetEngineTime
var s2sdk_GetEngineTime func() float64
var S2sdk_GetEngineTime = &s2sdk_GetEngineTime

//go:linkname s2sdk_GetMaxClients __package__/s2sdk._GetMaxClients
var s2sdk_GetMaxClients func() int32
var S2sdk_GetMaxClients = &s2sdk_GetMaxClients

//go:linkname s2sdk_Precache __package__/s2sdk._Precache
var s2sdk_Precache func(resource string)
var S2sdk_Precache = &s2sdk_Precache

//go:linkname s2sdk_IsPrecached __package__/s2sdk._IsPrecached
var s2sdk_IsPrecached func(resource string) bool
var S2sdk_IsPrecached = &s2sdk_IsPrecached

//go:linkname s2sdk_GetEconItemSystem __package__/s2sdk._GetEconItemSystem
var s2sdk_GetEconItemSystem func() uintptr
var S2sdk_GetEconItemSystem = &s2sdk_GetEconItemSystem

//go:linkname s2sdk_IsServerPaused __package__/s2sdk._IsServerPaused
var s2sdk_IsServerPaused func() bool
var S2sdk_IsServerPaused = &s2sdk_IsServerPaused

//go:linkname s2sdk_QueueTaskForNextFrame __package__/s2sdk._QueueTaskForNextFrame
var s2sdk_QueueTaskForNextFrame func(callback s2sdk.TaskCallback, userData []any)
var S2sdk_QueueTaskForNextFrame = &s2sdk_QueueTaskForNextFrame

//go:linkname s2sdk_QueueTaskForNextWorldUpdate __package__/s2sdk._QueueTaskForNextWorldUpdate
var s2sdk_QueueTaskForNextWorldUpdate func(callback s2sdk.TaskCallback, userData []any)
var S2sdk_QueueTaskForNextWorldUpdate = &s2sdk_QueueTaskForNextWorldUpdate

//go:linkname s2sdk_GetSoundDuration __package__/s2sdk._GetSoundDuration
var s2sdk_GetSoundDuration func(name string) float32
var S2sdk_GetSoundDuration = &s2sdk_GetSoundDuration

//go:linkname s2sdk_EmitSound __package__/s2sdk._EmitSound
var s2sdk_EmitSound func(entityHandle int32, sound string, pitch int32, volume float32, delay float32)
var S2sdk_EmitSound = &s2sdk_EmitSound

//go:linkname s2sdk_StopSound __package__/s2sdk._StopSound
var s2sdk_StopSound func(entityHandle int32, sound string)
var S2sdk_StopSound = &s2sdk_StopSound

//go:linkname s2sdk_EmitSoundToClient __package__/s2sdk._EmitSoundToClient
var s2sdk_EmitSoundToClient func(entityHandle int32, playersSlot []int32, sound string, volume float32, pitch float32)
var S2sdk_EmitSoundToClient = &s2sdk_EmitSoundToClient

//go:linkname s2sdk_GetPublicAddress __package__/s2sdk._GetPublicAddress
var s2sdk_GetPublicAddress func(onlyBase bool) string
var S2sdk_GetPublicAddress = &s2sdk_GetPublicAddress

//go:linkname s2sdk_GetLocalAddress __package__/s2sdk._GetLocalAddress
var s2sdk_GetLocalAddress func(onlyBase bool) string
var S2sdk_GetLocalAddress = &s2sdk_GetLocalAddress

//go:linkname s2sdk_EntIndexToEntPointer __package__/s2sdk._EntIndexToEntPointer
var s2sdk_EntIndexToEntPointer func(entityIndex int32) uintptr
var S2sdk_EntIndexToEntPointer = &s2sdk_EntIndexToEntPointer

//go:linkname s2sdk_EntPointerToEntIndex __package__/s2sdk._EntPointerToEntIndex
var s2sdk_EntPointerToEntIndex func(entity uintptr) int32
var S2sdk_EntPointerToEntIndex = &s2sdk_EntPointerToEntIndex

//go:linkname s2sdk_EntPointerToEntHandle __package__/s2sdk._EntPointerToEntHandle
var s2sdk_EntPointerToEntHandle func(entity uintptr) int32
var S2sdk_EntPointerToEntHandle = &s2sdk_EntPointerToEntHandle

//go:linkname s2sdk_EntHandleToEntPointer __package__/s2sdk._EntHandleToEntPointer
var s2sdk_EntHandleToEntPointer func(entityHandle int32) uintptr
var S2sdk_EntHandleToEntPointer = &s2sdk_EntHandleToEntPointer

//go:linkname s2sdk_EntIndexToEntHandle __package__/s2sdk._EntIndexToEntHandle
var s2sdk_EntIndexToEntHandle func(entityIndex int32) int32
var S2sdk_EntIndexToEntHandle = &s2sdk_EntIndexToEntHandle

//go:linkname s2sdk_EntHandleToEntIndex __package__/s2sdk._EntHandleToEntIndex
var s2sdk_EntHandleToEntIndex func(entityHandle int32) int32
var S2sdk_EntHandleToEntIndex = &s2sdk_EntHandleToEntIndex

//go:linkname s2sdk_IsValidEntHandle __package__/s2sdk._IsValidEntHandle
var s2sdk_IsValidEntHandle func(entityHandle int32) bool
var S2sdk_IsValidEntHandle = &s2sdk_IsValidEntHandle

//go:linkname s2sdk_IsValidEntPointer __package__/s2sdk._IsValidEntPointer
var s2sdk_IsValidEntPointer func(entity uintptr) bool
var S2sdk_IsValidEntPointer = &s2sdk_IsValidEntPointer

//go:linkname s2sdk_GetFirstActiveEntity __package__/s2sdk._GetFirstActiveEntity
var s2sdk_GetFirstActiveEntity func() int32
var S2sdk_GetFirstActiveEntity = &s2sdk_GetFirstActiveEntity

//go:linkname s2sdk_GetPrevActiveEntity __package__/s2sdk._GetPrevActiveEntity
var s2sdk_GetPrevActiveEntity func(entityHandle int32) int32
var S2sdk_GetPrevActiveEntity = &s2sdk_GetPrevActiveEntity

//go:linkname s2sdk_GetNextActiveEntity __package__/s2sdk._GetNextActiveEntity
var s2sdk_GetNextActiveEntity func(entityHandle int32) int32
var S2sdk_GetNextActiveEntity = &s2sdk_GetNextActiveEntity

//go:linkname s2sdk_HookEntityOutput __package__/s2sdk._HookEntityOutput
var s2sdk_HookEntityOutput func(classname string, output string, callback s2sdk.HookEntityOutputCallback, type_ s2sdk.HookMode) bool
var S2sdk_HookEntityOutput = &s2sdk_HookEntityOutput

//go:linkname s2sdk_UnhookEntityOutput __package__/s2sdk._UnhookEntityOutput
var s2sdk_UnhookEntityOutput func(classname string, output string, callback s2sdk.HookEntityOutputCallback, type_ s2sdk.HookMode) bool
var S2sdk_UnhookEntityOutput = &s2sdk_UnhookEntityOutput

//go:linkname s2sdk_FindEntityByClassname __package__/s2sdk._FindEntityByClassname
var s2sdk_FindEntityByClassname func(startFrom int32, classname string) int32
var S2sdk_FindEntityByClassname = &s2sdk_FindEntityByClassname

//go:linkname s2sdk_FindEntityByClassnameNearest __package__/s2sdk._FindEntityByClassnameNearest
var s2sdk_FindEntityByClassnameNearest func(startFrom int32, classname string, origin plugify.Vector3, maxRadius float32) int32
var S2sdk_FindEntityByClassnameNearest = &s2sdk_FindEntityByClassnameNearest

//go:linkname s2sdk_FindEntityByClassnameWithin __package__/s2sdk._FindEntityByClassnameWithin
var s2sdk_FindEntityByClassnameWithin func(startFrom int32, classname string, origin plugify.Vector3, radius float32) int32
var S2sdk_FindEntityByClassnameWithin = &s2sdk_FindEntityByClassnameWithin

//go:linkname s2sdk_FindEntityByName __package__/s2sdk._FindEntityByName
var s2sdk_FindEntityByName func(startFrom int32, name string) int32
var S2sdk_FindEntityByName = &s2sdk_FindEntityByName

//go:linkname s2sdk_FindEntityByNameNearest __package__/s2sdk._FindEntityByNameNearest
var s2sdk_FindEntityByNameNearest func(name string, origin plugify.Vector3, maxRadius float32) int32
var S2sdk_FindEntityByNameNearest = &s2sdk_FindEntityByNameNearest

//go:linkname s2sdk_FindEntityByNameWithin __package__/s2sdk._FindEntityByNameWithin
var s2sdk_FindEntityByNameWithin func(startFrom int32, name string, origin plugify.Vector3, radius float32) int32
var S2sdk_FindEntityByNameWithin = &s2sdk_FindEntityByNameWithin

//go:linkname s2sdk_FindEntityByTarget __package__/s2sdk._FindEntityByTarget
var s2sdk_FindEntityByTarget func(startFrom int32, name string) int32
var S2sdk_FindEntityByTarget = &s2sdk_FindEntityByTarget

//go:linkname s2sdk_FindEntityInSphere __package__/s2sdk._FindEntityInSphere
var s2sdk_FindEntityInSphere func(startFrom int32, origin plugify.Vector3, radius float32) int32
var S2sdk_FindEntityInSphere = &s2sdk_FindEntityInSphere

//go:linkname s2sdk_SpawnEntityByName __package__/s2sdk._SpawnEntityByName
var s2sdk_SpawnEntityByName func(className string) int32
var S2sdk_SpawnEntityByName = &s2sdk_SpawnEntityByName

//go:linkname s2sdk_CreateEntityByName __package__/s2sdk._CreateEntityByName
var s2sdk_CreateEntityByName func(className string) int32
var S2sdk_CreateEntityByName = &s2sdk_CreateEntityByName

//go:linkname s2sdk_DispatchSpawn __package__/s2sdk._DispatchSpawn
var s2sdk_DispatchSpawn func(entityHandle int32)
var S2sdk_DispatchSpawn = &s2sdk_DispatchSpawn

//go:linkname s2sdk_DispatchSpawn2 __package__/s2sdk._DispatchSpawn2
var s2sdk_DispatchSpawn2 func(entityHandle int32, keys []string, values []any)
var S2sdk_DispatchSpawn2 = &s2sdk_DispatchSpawn2

//go:linkname s2sdk_RemoveEntity __package__/s2sdk._RemoveEntity
var s2sdk_RemoveEntity func(entityHandle int32)
var S2sdk_RemoveEntity = &s2sdk_RemoveEntity

//go:linkname s2sdk_IsEntityPlayerController __package__/s2sdk._IsEntityPlayerController
var s2sdk_IsEntityPlayerController func(entityHandle int32) bool
var S2sdk_IsEntityPlayerController = &s2sdk_IsEntityPlayerController

//go:linkname s2sdk_IsEntityPlayerPawn __package__/s2sdk._IsEntityPlayerPawn
var s2sdk_IsEntityPlayerPawn func(entityHandle int32) bool
var S2sdk_IsEntityPlayerPawn = &s2sdk_IsEntityPlayerPawn

//go:linkname s2sdk_GetEntityClassname __package__/s2sdk._GetEntityClassname
var s2sdk_GetEntityClassname func(entityHandle int32) string
var S2sdk_GetEntityClassname = &s2sdk_GetEntityClassname

//go:linkname s2sdk_GetEntityName __package__/s2sdk._GetEntityName
var s2sdk_GetEntityName func(entityHandle int32) string
var S2sdk_GetEntityName = &s2sdk_GetEntityName

//go:linkname s2sdk_SetEntityName __package__/s2sdk._SetEntityName
var s2sdk_SetEntityName func(entityHandle int32, name string)
var S2sdk_SetEntityName = &s2sdk_SetEntityName

//go:linkname s2sdk_GetEntityMoveType __package__/s2sdk._GetEntityMoveType
var s2sdk_GetEntityMoveType func(entityHandle int32) s2sdk.MoveType
var S2sdk_GetEntityMoveType = &s2sdk_GetEntityMoveType

//go:linkname s2sdk_SetEntityMoveType __package__/s2sdk._SetEntityMoveType
var s2sdk_SetEntityMoveType func(entityHandle int32, moveType s2sdk.MoveType)
var S2sdk_SetEntityMoveType = &s2sdk_SetEntityMoveType

//go:linkname s2sdk_GetEntityGravity __package__/s2sdk._GetEntityGravity
var s2sdk_GetEntityGravity func(entityHandle int32) float32
var S2sdk_GetEntityGravity = &s2sdk_GetEntityGravity

//go:linkname s2sdk_SetEntityGravity __package__/s2sdk._SetEntityGravity
var s2sdk_SetEntityGravity func(entityHandle int32, gravity float32)
var S2sdk_SetEntityGravity = &s2sdk_SetEntityGravity

//go:linkname s2sdk_GetEntityFlags __package__/s2sdk._GetEntityFlags
var s2sdk_GetEntityFlags func(entityHandle int32) int32
var S2sdk_GetEntityFlags = &s2sdk_GetEntityFlags

//go:linkname s2sdk_SetEntityFlags __package__/s2sdk._SetEntityFlags
var s2sdk_SetEntityFlags func(entityHandle int32, flags int32)
var S2sdk_SetEntityFlags = &s2sdk_SetEntityFlags

//go:linkname s2sdk_GetEntityRenderColor __package__/s2sdk._GetEntityRenderColor
var s2sdk_GetEntityRenderColor func(entityHandle int32) plugify.Vector4
var S2sdk_GetEntityRenderColor = &s2sdk_GetEntityRenderColor

//go:linkname s2sdk_SetEntityRenderColor __package__/s2sdk._SetEntityRenderColor
var s2sdk_SetEntityRenderColor func(entityHandle int32, color plugify.Vector4)
var S2sdk_SetEntityRenderColor = &s2sdk_SetEntityRenderColor

//go:linkname s2sdk_GetEntityRenderMode __package__/s2sdk._GetEntityRenderMode
var s2sdk_GetEntityRenderMode func(entityHandle int32) s2sdk.RenderMode
var S2sdk_GetEntityRenderMode = &s2sdk_GetEntityRenderMode

//go:linkname s2sdk_SetEntityRenderMode __package__/s2sdk._SetEntityRenderMode
var s2sdk_SetEntityRenderMode func(entityHandle int32, renderMode s2sdk.RenderMode)
var S2sdk_SetEntityRenderMode = &s2sdk_SetEntityRenderMode

//go:linkname s2sdk_GetEntityMass __package__/s2sdk._GetEntityMass
var s2sdk_GetEntityMass func(entityHandle int32) int32
var S2sdk_GetEntityMass = &s2sdk_GetEntityMass

//go:linkname s2sdk_SetEntityMass __package__/s2sdk._SetEntityMass
var s2sdk_SetEntityMass func(entityHandle int32, mass int32)
var S2sdk_SetEntityMass = &s2sdk_SetEntityMass

//go:linkname s2sdk_GetEntityFriction __package__/s2sdk._GetEntityFriction
var s2sdk_GetEntityFriction func(entityHandle int32) float32
var S2sdk_GetEntityFriction = &s2sdk_GetEntityFriction

//go:linkname s2sdk_SetEntityFriction __package__/s2sdk._SetEntityFriction
var s2sdk_SetEntityFriction func(entityHandle int32, friction float32)
var S2sdk_SetEntityFriction = &s2sdk_SetEntityFriction

//go:linkname s2sdk_GetEntityHealth __package__/s2sdk._GetEntityHealth
var s2sdk_GetEntityHealth func(entityHandle int32) int32
var S2sdk_GetEntityHealth = &s2sdk_GetEntityHealth

//go:linkname s2sdk_SetEntityHealth __package__/s2sdk._SetEntityHealth
var s2sdk_SetEntityHealth func(entityHandle int32, health int32)
var S2sdk_SetEntityHealth = &s2sdk_SetEntityHealth

//go:linkname s2sdk_GetEntityMaxHealth __package__/s2sdk._GetEntityMaxHealth
var s2sdk_GetEntityMaxHealth func(entityHandle int32) int32
var S2sdk_GetEntityMaxHealth = &s2sdk_GetEntityMaxHealth

//go:linkname s2sdk_SetEntityMaxHealth __package__/s2sdk._SetEntityMaxHealth
var s2sdk_SetEntityMaxHealth func(entityHandle int32, maxHealth int32)
var S2sdk_SetEntityMaxHealth = &s2sdk_SetEntityMaxHealth

//go:linkname s2sdk_GetEntityTeam __package__/s2sdk._GetEntityTeam
var s2sdk_GetEntityTeam func(entityHandle int32) s2sdk.CSTeam
var S2sdk_GetEntityTeam = &s2sdk_GetEntityTeam

//go:linkname s2sdk_SetEntityTeam __package__/s2sdk._SetEntityTeam
var s2sdk_SetEntityTeam func(entityHandle int32, team s2sdk.CSTeam)
var S2sdk_SetEntityTeam = &s2sdk_SetEntityTeam

//go:linkname s2sdk_GetEntityOwner __package__/s2sdk._GetEntityOwner
var s2sdk_GetEntityOwner func(entityHandle int32) int32
var S2sdk_GetEntityOwner = &s2sdk_GetEntityOwner

//go:linkname s2sdk_SetEntityOwner __package__/s2sdk._SetEntityOwner
var s2sdk_SetEntityOwner func(entityHandle int32, ownerHandle int32)
var S2sdk_SetEntityOwner = &s2sdk_SetEntityOwner

//go:linkname s2sdk_GetEntityParent __package__/s2sdk._GetEntityParent
var s2sdk_GetEntityParent func(entityHandle int32) int32
var S2sdk_GetEntityParent = &s2sdk_GetEntityParent

//go:linkname s2sdk_SetEntityParent __package__/s2sdk._SetEntityParent
var s2sdk_SetEntityParent func(entityHandle int32, parentHandle int32)
var S2sdk_SetEntityParent = &s2sdk_SetEntityParent

//go:linkname s2sdk_SetEntityParentAttachment __package__/s2sdk._SetEntityParentAttachment
var s2sdk_SetEntityParentAttachment func(entityHandle int32, parentHandle int32, attachmentName string)
var S2sdk_SetEntityParentAttachment = &s2sdk_SetEntityParentAttachment

//go:linkname s2sdk_GetEntityAbsOrigin __package__/s2sdk._GetEntityAbsOrigin
var s2sdk_GetEntityAbsOrigin func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityAbsOrigin = &s2sdk_GetEntityAbsOrigin

//go:linkname s2sdk_SetEntityAbsOrigin __package__/s2sdk._SetEntityAbsOrigin
var s2sdk_SetEntityAbsOrigin func(entityHandle int32, origin plugify.Vector3)
var S2sdk_SetEntityAbsOrigin = &s2sdk_SetEntityAbsOrigin

//go:linkname s2sdk_GetEntityAbsScale __package__/s2sdk._GetEntityAbsScale
var s2sdk_GetEntityAbsScale func(entityHandle int32) float32
var S2sdk_GetEntityAbsScale = &s2sdk_GetEntityAbsScale

//go:linkname s2sdk_SetEntityAbsScale __package__/s2sdk._SetEntityAbsScale
var s2sdk_SetEntityAbsScale func(entityHandle int32, scale float32)
var S2sdk_SetEntityAbsScale = &s2sdk_SetEntityAbsScale

//go:linkname s2sdk_GetEntityAbsAngles __package__/s2sdk._GetEntityAbsAngles
var s2sdk_GetEntityAbsAngles func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityAbsAngles = &s2sdk_GetEntityAbsAngles

//go:linkname s2sdk_SetEntityAbsAngles __package__/s2sdk._SetEntityAbsAngles
var s2sdk_SetEntityAbsAngles func(entityHandle int32, angle plugify.Vector3)
var S2sdk_SetEntityAbsAngles = &s2sdk_SetEntityAbsAngles

//go:linkname s2sdk_GetEntityLocalOrigin __package__/s2sdk._GetEntityLocalOrigin
var s2sdk_GetEntityLocalOrigin func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityLocalOrigin = &s2sdk_GetEntityLocalOrigin

//go:linkname s2sdk_SetEntityLocalOrigin __package__/s2sdk._SetEntityLocalOrigin
var s2sdk_SetEntityLocalOrigin func(entityHandle int32, origin plugify.Vector3)
var S2sdk_SetEntityLocalOrigin = &s2sdk_SetEntityLocalOrigin

//go:linkname s2sdk_GetEntityLocalScale __package__/s2sdk._GetEntityLocalScale
var s2sdk_GetEntityLocalScale func(entityHandle int32) float32
var S2sdk_GetEntityLocalScale = &s2sdk_GetEntityLocalScale

//go:linkname s2sdk_SetEntityLocalScale __package__/s2sdk._SetEntityLocalScale
var s2sdk_SetEntityLocalScale func(entityHandle int32, scale float32)
var S2sdk_SetEntityLocalScale = &s2sdk_SetEntityLocalScale

//go:linkname s2sdk_GetEntityLocalAngles __package__/s2sdk._GetEntityLocalAngles
var s2sdk_GetEntityLocalAngles func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityLocalAngles = &s2sdk_GetEntityLocalAngles

//go:linkname s2sdk_SetEntityLocalAngles __package__/s2sdk._SetEntityLocalAngles
var s2sdk_SetEntityLocalAngles func(entityHandle int32, angle plugify.Vector3)
var S2sdk_SetEntityLocalAngles = &s2sdk_SetEntityLocalAngles

//go:linkname s2sdk_GetEntityAbsVelocity __package__/s2sdk._GetEntityAbsVelocity
var s2sdk_GetEntityAbsVelocity func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityAbsVelocity = &s2sdk_GetEntityAbsVelocity

//go:linkname s2sdk_SetEntityAbsVelocity __package__/s2sdk._SetEntityAbsVelocity
var s2sdk_SetEntityAbsVelocity func(entityHandle int32, velocity plugify.Vector3)
var S2sdk_SetEntityAbsVelocity = &s2sdk_SetEntityAbsVelocity

//go:linkname s2sdk_GetEntityBaseVelocity __package__/s2sdk._GetEntityBaseVelocity
var s2sdk_GetEntityBaseVelocity func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityBaseVelocity = &s2sdk_GetEntityBaseVelocity

//go:linkname s2sdk_GetEntityLocalAngVelocity __package__/s2sdk._GetEntityLocalAngVelocity
var s2sdk_GetEntityLocalAngVelocity func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityLocalAngVelocity = &s2sdk_GetEntityLocalAngVelocity

//go:linkname s2sdk_GetEntityAngVelocity __package__/s2sdk._GetEntityAngVelocity
var s2sdk_GetEntityAngVelocity func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityAngVelocity = &s2sdk_GetEntityAngVelocity

//go:linkname s2sdk_SetEntityAngVelocity __package__/s2sdk._SetEntityAngVelocity
var s2sdk_SetEntityAngVelocity func(entityHandle int32, velocity plugify.Vector3)
var S2sdk_SetEntityAngVelocity = &s2sdk_SetEntityAngVelocity

//go:linkname s2sdk_GetEntityLocalVelocity __package__/s2sdk._GetEntityLocalVelocity
var s2sdk_GetEntityLocalVelocity func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityLocalVelocity = &s2sdk_GetEntityLocalVelocity

//go:linkname s2sdk_GetEntityAngRotation __package__/s2sdk._GetEntityAngRotation
var s2sdk_GetEntityAngRotation func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityAngRotation = &s2sdk_GetEntityAngRotation

//go:linkname s2sdk_SetEntityAngRotation __package__/s2sdk._SetEntityAngRotation
var s2sdk_SetEntityAngRotation func(entityHandle int32, rotation plugify.Vector3)
var S2sdk_SetEntityAngRotation = &s2sdk_SetEntityAngRotation

//go:linkname s2sdk_TransformPointEntityToWorld __package__/s2sdk._TransformPointEntityToWorld
var s2sdk_TransformPointEntityToWorld func(entityHandle int32, point plugify.Vector3) plugify.Vector3
var S2sdk_TransformPointEntityToWorld = &s2sdk_TransformPointEntityToWorld

//go:linkname s2sdk_TransformPointWorldToEntity __package__/s2sdk._TransformPointWorldToEntity
var s2sdk_TransformPointWorldToEntity func(entityHandle int32, point plugify.Vector3) plugify.Vector3
var S2sdk_TransformPointWorldToEntity = &s2sdk_TransformPointWorldToEntity

//go:linkname s2sdk_GetEntityEyePosition __package__/s2sdk._GetEntityEyePosition
var s2sdk_GetEntityEyePosition func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityEyePosition = &s2sdk_GetEntityEyePosition

//go:linkname s2sdk_GetEntityEyeAngles __package__/s2sdk._GetEntityEyeAngles
var s2sdk_GetEntityEyeAngles func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityEyeAngles = &s2sdk_GetEntityEyeAngles

//go:linkname s2sdk_SetEntityForwardVector __package__/s2sdk._SetEntityForwardVector
var s2sdk_SetEntityForwardVector func(entityHandle int32, forward plugify.Vector3)
var S2sdk_SetEntityForwardVector = &s2sdk_SetEntityForwardVector

//go:linkname s2sdk_GetEntityForwardVector __package__/s2sdk._GetEntityForwardVector
var s2sdk_GetEntityForwardVector func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityForwardVector = &s2sdk_GetEntityForwardVector

//go:linkname s2sdk_GetEntityLeftVector __package__/s2sdk._GetEntityLeftVector
var s2sdk_GetEntityLeftVector func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityLeftVector = &s2sdk_GetEntityLeftVector

//go:linkname s2sdk_GetEntityRightVector __package__/s2sdk._GetEntityRightVector
var s2sdk_GetEntityRightVector func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityRightVector = &s2sdk_GetEntityRightVector

//go:linkname s2sdk_GetEntityUpVector __package__/s2sdk._GetEntityUpVector
var s2sdk_GetEntityUpVector func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityUpVector = &s2sdk_GetEntityUpVector

//go:linkname s2sdk_GetEntityTransform __package__/s2sdk._GetEntityTransform
var s2sdk_GetEntityTransform func(entityHandle int32) plugify.Matrix4x4
var S2sdk_GetEntityTransform = &s2sdk_GetEntityTransform

//go:linkname s2sdk_GetEntityModel __package__/s2sdk._GetEntityModel
var s2sdk_GetEntityModel func(entityHandle int32) string
var S2sdk_GetEntityModel = &s2sdk_GetEntityModel

//go:linkname s2sdk_SetEntityModel __package__/s2sdk._SetEntityModel
var s2sdk_SetEntityModel func(entityHandle int32, model string)
var S2sdk_SetEntityModel = &s2sdk_SetEntityModel

//go:linkname s2sdk_GetEntityWaterLevel __package__/s2sdk._GetEntityWaterLevel
var s2sdk_GetEntityWaterLevel func(entityHandle int32) float32
var S2sdk_GetEntityWaterLevel = &s2sdk_GetEntityWaterLevel

//go:linkname s2sdk_GetEntityGroundEntity __package__/s2sdk._GetEntityGroundEntity
var s2sdk_GetEntityGroundEntity func(entityHandle int32) int32
var S2sdk_GetEntityGroundEntity = &s2sdk_GetEntityGroundEntity

//go:linkname s2sdk_GetEntityEffects __package__/s2sdk._GetEntityEffects
var s2sdk_GetEntityEffects func(entityHandle int32) int32
var S2sdk_GetEntityEffects = &s2sdk_GetEntityEffects

//go:linkname s2sdk_AddEntityEffects __package__/s2sdk._AddEntityEffects
var s2sdk_AddEntityEffects func(entityHandle int32, effects int32)
var S2sdk_AddEntityEffects = &s2sdk_AddEntityEffects

//go:linkname s2sdk_RemoveEntityEffects __package__/s2sdk._RemoveEntityEffects
var s2sdk_RemoveEntityEffects func(entityHandle int32, effects int32)
var S2sdk_RemoveEntityEffects = &s2sdk_RemoveEntityEffects

//go:linkname s2sdk_GetEntityBoundingMaxs __package__/s2sdk._GetEntityBoundingMaxs
var s2sdk_GetEntityBoundingMaxs func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityBoundingMaxs = &s2sdk_GetEntityBoundingMaxs

//go:linkname s2sdk_GetEntityBoundingMins __package__/s2sdk._GetEntityBoundingMins
var s2sdk_GetEntityBoundingMins func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityBoundingMins = &s2sdk_GetEntityBoundingMins

//go:linkname s2sdk_GetEntityCenter __package__/s2sdk._GetEntityCenter
var s2sdk_GetEntityCenter func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityCenter = &s2sdk_GetEntityCenter

//go:linkname s2sdk_TeleportEntity __package__/s2sdk._TeleportEntity
var s2sdk_TeleportEntity func(entityHandle int32, origin plugify.Vector3, angles plugify.Vector3, velocity plugify.Vector3)
var S2sdk_TeleportEntity = &s2sdk_TeleportEntity

//go:linkname s2sdk_ApplyAbsVelocityImpulseToEntity __package__/s2sdk._ApplyAbsVelocityImpulseToEntity
var s2sdk_ApplyAbsVelocityImpulseToEntity func(entityHandle int32, vecImpulse plugify.Vector3)
var S2sdk_ApplyAbsVelocityImpulseToEntity = &s2sdk_ApplyAbsVelocityImpulseToEntity

//go:linkname s2sdk_ApplyLocalAngularVelocityImpulseToEntity __package__/s2sdk._ApplyLocalAngularVelocityImpulseToEntity
var s2sdk_ApplyLocalAngularVelocityImpulseToEntity func(entityHandle int32, angImpulse plugify.Vector3)
var S2sdk_ApplyLocalAngularVelocityImpulseToEntity = &s2sdk_ApplyLocalAngularVelocityImpulseToEntity

//go:linkname s2sdk_AcceptEntityInput __package__/s2sdk._AcceptEntityInput
var s2sdk_AcceptEntityInput func(entityHandle int32, inputName string, activatorHandle int32, callerHandle int32, value any, type_ s2sdk.FieldType, outputId int32)
var S2sdk_AcceptEntityInput = &s2sdk_AcceptEntityInput

//go:linkname s2sdk_ConnectEntityOutput __package__/s2sdk._ConnectEntityOutput
var s2sdk_ConnectEntityOutput func(entityHandle int32, output string, functionName string)
var S2sdk_ConnectEntityOutput = &s2sdk_ConnectEntityOutput

//go:linkname s2sdk_DisconnectEntityOutput __package__/s2sdk._DisconnectEntityOutput
var s2sdk_DisconnectEntityOutput func(entityHandle int32, output string, functionName string)
var S2sdk_DisconnectEntityOutput = &s2sdk_DisconnectEntityOutput

//go:linkname s2sdk_DisconnectEntityRedirectedOutput __package__/s2sdk._DisconnectEntityRedirectedOutput
var s2sdk_DisconnectEntityRedirectedOutput func(entityHandle int32, output string, functionName string, targetHandle int32)
var S2sdk_DisconnectEntityRedirectedOutput = &s2sdk_DisconnectEntityRedirectedOutput

//go:linkname s2sdk_FireEntityOutput __package__/s2sdk._FireEntityOutput
var s2sdk_FireEntityOutput func(entityHandle int32, outputName string, activatorHandle int32, callerHandle int32, value any, type_ s2sdk.FieldType, delay float32)
var S2sdk_FireEntityOutput = &s2sdk_FireEntityOutput

//go:linkname s2sdk_RedirectEntityOutput __package__/s2sdk._RedirectEntityOutput
var s2sdk_RedirectEntityOutput func(entityHandle int32, output string, functionName string, targetHandle int32)
var S2sdk_RedirectEntityOutput = &s2sdk_RedirectEntityOutput

//go:linkname s2sdk_FollowEntity __package__/s2sdk._FollowEntity
var s2sdk_FollowEntity func(entityHandle int32, attachmentHandle int32, boneMerge bool)
var S2sdk_FollowEntity = &s2sdk_FollowEntity

//go:linkname s2sdk_FollowEntityMerge __package__/s2sdk._FollowEntityMerge
var s2sdk_FollowEntityMerge func(entityHandle int32, attachmentHandle int32, boneOrAttachName string)
var S2sdk_FollowEntityMerge = &s2sdk_FollowEntityMerge

//go:linkname s2sdk_TakeEntityDamage __package__/s2sdk._TakeEntityDamage
var s2sdk_TakeEntityDamage func(entityHandle int32, inflictorHandle int32, attackerHandle int32, force plugify.Vector3, hitPos plugify.Vector3, damage float32, damageTypes s2sdk.DamageTypes) int32
var S2sdk_TakeEntityDamage = &s2sdk_TakeEntityDamage

//go:linkname s2sdk_GetEntityAttributeFloatValue __package__/s2sdk._GetEntityAttributeFloatValue
var s2sdk_GetEntityAttributeFloatValue func(entityHandle int32, name string, defaultValue float32) float32
var S2sdk_GetEntityAttributeFloatValue = &s2sdk_GetEntityAttributeFloatValue

//go:linkname s2sdk_GetEntityAttributeIntValue __package__/s2sdk._GetEntityAttributeIntValue
var s2sdk_GetEntityAttributeIntValue func(entityHandle int32, name string, defaultValue int32) int32
var S2sdk_GetEntityAttributeIntValue = &s2sdk_GetEntityAttributeIntValue

//go:linkname s2sdk_SetEntityAttributeFloatValue __package__/s2sdk._SetEntityAttributeFloatValue
var s2sdk_SetEntityAttributeFloatValue func(entityHandle int32, name string, value float32)
var S2sdk_SetEntityAttributeFloatValue = &s2sdk_SetEntityAttributeFloatValue

//go:linkname s2sdk_SetEntityAttributeIntValue __package__/s2sdk._SetEntityAttributeIntValue
var s2sdk_SetEntityAttributeIntValue func(entityHandle int32, name string, value int32)
var S2sdk_SetEntityAttributeIntValue = &s2sdk_SetEntityAttributeIntValue

//go:linkname s2sdk_DeleteEntityAttribute __package__/s2sdk._DeleteEntityAttribute
var s2sdk_DeleteEntityAttribute func(entityHandle int32, name string)
var S2sdk_DeleteEntityAttribute = &s2sdk_DeleteEntityAttribute

//go:linkname s2sdk_HasEntityAttribute __package__/s2sdk._HasEntityAttribute
var s2sdk_HasEntityAttribute func(entityHandle int32, name string) bool
var S2sdk_HasEntityAttribute = &s2sdk_HasEntityAttribute

//go:linkname s2sdk_HookEvent __package__/s2sdk._HookEvent
var s2sdk_HookEvent func(name string, callback s2sdk.EventCallback, type_ s2sdk.HookMode) s2sdk.EventHookError
var S2sdk_HookEvent = &s2sdk_HookEvent

//go:linkname s2sdk_UnhookEvent __package__/s2sdk._UnhookEvent
var s2sdk_UnhookEvent func(name string, callback s2sdk.EventCallback, type_ s2sdk.HookMode) s2sdk.EventHookError
var S2sdk_UnhookEvent = &s2sdk_UnhookEvent

//go:linkname s2sdk_CreateEvent __package__/s2sdk._CreateEvent
var s2sdk_CreateEvent func(name string, force bool) uintptr
var S2sdk_CreateEvent = &s2sdk_CreateEvent

//go:linkname s2sdk_FireEvent __package__/s2sdk._FireEvent
var s2sdk_FireEvent func(event uintptr, dontBroadcast bool)
var S2sdk_FireEvent = &s2sdk_FireEvent

//go:linkname s2sdk_FireEventToClient __package__/s2sdk._FireEventToClient
var s2sdk_FireEventToClient func(event uintptr, playerSlot int32)
var S2sdk_FireEventToClient = &s2sdk_FireEventToClient

//go:linkname s2sdk_CancelCreatedEvent __package__/s2sdk._CancelCreatedEvent
var s2sdk_CancelCreatedEvent func(event uintptr)
var S2sdk_CancelCreatedEvent = &s2sdk_CancelCreatedEvent

//go:linkname s2sdk_GetEventBool __package__/s2sdk._GetEventBool
var s2sdk_GetEventBool func(event uintptr, key string) bool
var S2sdk_GetEventBool = &s2sdk_GetEventBool

//go:linkname s2sdk_GetEventFloat __package__/s2sdk._GetEventFloat
var s2sdk_GetEventFloat func(event uintptr, key string) float32
var S2sdk_GetEventFloat = &s2sdk_GetEventFloat

//go:linkname s2sdk_GetEventInt __package__/s2sdk._GetEventInt
var s2sdk_GetEventInt func(event uintptr, key string) int32
var S2sdk_GetEventInt = &s2sdk_GetEventInt

//go:linkname s2sdk_GetEventUInt64 __package__/s2sdk._GetEventUInt64
var s2sdk_GetEventUInt64 func(event uintptr, key string) uint64
var S2sdk_GetEventUInt64 = &s2sdk_GetEventUInt64

//go:linkname s2sdk_GetEventString __package__/s2sdk._GetEventString
var s2sdk_GetEventString func(event uintptr, key string) string
var S2sdk_GetEventString = &s2sdk_GetEventString

//go:linkname s2sdk_GetEventPtr __package__/s2sdk._GetEventPtr
var s2sdk_GetEventPtr func(event uintptr, key string) uintptr
var S2sdk_GetEventPtr = &s2sdk_GetEventPtr

//go:linkname s2sdk_GetEventPlayerController __package__/s2sdk._GetEventPlayerController
var s2sdk_GetEventPlayerController func(event uintptr, key string) uintptr
var S2sdk_GetEventPlayerController = &s2sdk_GetEventPlayerController

//go:linkname s2sdk_GetEventPlayerIndex __package__/s2sdk._GetEventPlayerIndex
var s2sdk_GetEventPlayerIndex func(event uintptr, key string) int32
var S2sdk_GetEventPlayerIndex = &s2sdk_GetEventPlayerIndex

//go:linkname s2sdk_GetEventPlayerSlot __package__/s2sdk._GetEventPlayerSlot
var s2sdk_GetEventPlayerSlot func(event uintptr, key string) int32
var S2sdk_GetEventPlayerSlot = &s2sdk_GetEventPlayerSlot

//go:linkname s2sdk_GetEventPlayerPawn __package__/s2sdk._GetEventPlayerPawn
var s2sdk_GetEventPlayerPawn func(event uintptr, key string) uintptr
var S2sdk_GetEventPlayerPawn = &s2sdk_GetEventPlayerPawn

//go:linkname s2sdk_GetEventEntity __package__/s2sdk._GetEventEntity
var s2sdk_GetEventEntity func(event uintptr, key string) uintptr
var S2sdk_GetEventEntity = &s2sdk_GetEventEntity

//go:linkname s2sdk_GetEventEntityIndex __package__/s2sdk._GetEventEntityIndex
var s2sdk_GetEventEntityIndex func(event uintptr, key string) int32
var S2sdk_GetEventEntityIndex = &s2sdk_GetEventEntityIndex

//go:linkname s2sdk_GetEventEntityHandle __package__/s2sdk._GetEventEntityHandle
var s2sdk_GetEventEntityHandle func(event uintptr, key string) int32
var S2sdk_GetEventEntityHandle = &s2sdk_GetEventEntityHandle

//go:linkname s2sdk_GetEventName __package__/s2sdk._GetEventName
var s2sdk_GetEventName func(event uintptr) string
var S2sdk_GetEventName = &s2sdk_GetEventName

//go:linkname s2sdk_SetEventBool __package__/s2sdk._SetEventBool
var s2sdk_SetEventBool func(event uintptr, key string, value bool)
var S2sdk_SetEventBool = &s2sdk_SetEventBool

//go:linkname s2sdk_SetEventFloat __package__/s2sdk._SetEventFloat
var s2sdk_SetEventFloat func(event uintptr, key string, value float32)
var S2sdk_SetEventFloat = &s2sdk_SetEventFloat

//go:linkname s2sdk_SetEventInt __package__/s2sdk._SetEventInt
var s2sdk_SetEventInt func(event uintptr, key string, value int32)
var S2sdk_SetEventInt = &s2sdk_SetEventInt

//go:linkname s2sdk_SetEventUInt64 __package__/s2sdk._SetEventUInt64
var s2sdk_SetEventUInt64 func(event uintptr, key string, value uint64)
var S2sdk_SetEventUInt64 = &s2sdk_SetEventUInt64

//go:linkname s2sdk_SetEventString __package__/s2sdk._SetEventString
var s2sdk_SetEventString func(event uintptr, key string, value string)
var S2sdk_SetEventString = &s2sdk_SetEventString

//go:linkname s2sdk_SetEventPtr __package__/s2sdk._SetEventPtr
var s2sdk_SetEventPtr func(event uintptr, key string, value uintptr)
var S2sdk_SetEventPtr = &s2sdk_SetEventPtr

//go:linkname s2sdk_SetEventPlayerController __package__/s2sdk._SetEventPlayerController
var s2sdk_SetEventPlayerController func(event uintptr, key string, value uintptr)
var S2sdk_SetEventPlayerController = &s2sdk_SetEventPlayerController

//go:linkname s2sdk_SetEventPlayerIndex __package__/s2sdk._SetEventPlayerIndex
var s2sdk_SetEventPlayerIndex func(event uintptr, key string, value int32)
var S2sdk_SetEventPlayerIndex = &s2sdk_SetEventPlayerIndex

//go:linkname s2sdk_SetEventPlayerSlot __package__/s2sdk._SetEventPlayerSlot
var s2sdk_SetEventPlayerSlot func(event uintptr, key string, value int32)
var S2sdk_SetEventPlayerSlot = &s2sdk_SetEventPlayerSlot

//go:linkname s2sdk_SetEventEntity __package__/s2sdk._SetEventEntity
var s2sdk_SetEventEntity func(event uintptr, key string, value uintptr)
var S2sdk_SetEventEntity = &s2sdk_SetEventEntity

//go:linkname s2sdk_SetEventEntityIndex __package__/s2sdk._SetEventEntityIndex
var s2sdk_SetEventEntityIndex func(event uintptr, key string, value int32)
var S2sdk_SetEventEntityIndex = &s2sdk_SetEventEntityIndex

//go:linkname s2sdk_SetEventEntityHandle __package__/s2sdk._SetEventEntityHandle
var s2sdk_SetEventEntityHandle func(event uintptr, key string, value int32)
var S2sdk_SetEventEntityHandle = &s2sdk_SetEventEntityHandle

//go:linkname s2sdk_SetEventBroadcast __package__/s2sdk._SetEventBroadcast
var s2sdk_SetEventBroadcast func(event uintptr, dontBroadcast bool)
var S2sdk_SetEventBroadcast = &s2sdk_SetEventBroadcast

//go:linkname s2sdk_LoadEventsFromFile __package__/s2sdk._LoadEventsFromFile
var s2sdk_LoadEventsFromFile func(path string, searchAll bool) int32
var S2sdk_LoadEventsFromFile = &s2sdk_LoadEventsFromFile

//go:linkname s2sdk_CloseGameConfigFile __package__/s2sdk._CloseGameConfigFile
var s2sdk_CloseGameConfigFile func(id uint32)
var S2sdk_CloseGameConfigFile = &s2sdk_CloseGameConfigFile

//go:linkname s2sdk_LoadGameConfigFile __package__/s2sdk._LoadGameConfigFile
var s2sdk_LoadGameConfigFile func(paths []string) uint32
var S2sdk_LoadGameConfigFile = &s2sdk_LoadGameConfigFile

//go:linkname s2sdk_GetGameConfigPatch __package__/s2sdk._GetGameConfigPatch
var s2sdk_GetGameConfigPatch func(id uint32, name string) string
var S2sdk_GetGameConfigPatch = &s2sdk_GetGameConfigPatch

//go:linkname s2sdk_GetGameConfigOffset __package__/s2sdk._GetGameConfigOffset
var s2sdk_GetGameConfigOffset func(id uint32, name string) int32
var S2sdk_GetGameConfigOffset = &s2sdk_GetGameConfigOffset

//go:linkname s2sdk_GetGameConfigAddress __package__/s2sdk._GetGameConfigAddress
var s2sdk_GetGameConfigAddress func(id uint32, name string) uintptr
var S2sdk_GetGameConfigAddress = &s2sdk_GetGameConfigAddress

//go:linkname s2sdk_GetGameConfigVTable __package__/s2sdk._GetGameConfigVTable
var s2sdk_GetGameConfigVTable func(id uint32, name string) uintptr
var S2sdk_GetGameConfigVTable = &s2sdk_GetGameConfigVTable

//go:linkname s2sdk_GetGameConfigSignature __package__/s2sdk._GetGameConfigSignature
var s2sdk_GetGameConfigSignature func(id uint32, name string) uintptr
var S2sdk_GetGameConfigSignature = &s2sdk_GetGameConfigSignature

//go:linkname s2sdk_GetGameConfigPatchAll __package__/s2sdk._GetGameConfigPatchAll
var s2sdk_GetGameConfigPatchAll func(name string) string
var S2sdk_GetGameConfigPatchAll = &s2sdk_GetGameConfigPatchAll

//go:linkname s2sdk_GetGameConfigOffsetAll __package__/s2sdk._GetGameConfigOffsetAll
var s2sdk_GetGameConfigOffsetAll func(name string) int32
var S2sdk_GetGameConfigOffsetAll = &s2sdk_GetGameConfigOffsetAll

//go:linkname s2sdk_GetGameConfigAddressAll __package__/s2sdk._GetGameConfigAddressAll
var s2sdk_GetGameConfigAddressAll func(name string) uintptr
var S2sdk_GetGameConfigAddressAll = &s2sdk_GetGameConfigAddressAll

//go:linkname s2sdk_GetGameConfigVTableAll __package__/s2sdk._GetGameConfigVTableAll
var s2sdk_GetGameConfigVTableAll func(name string) uintptr
var S2sdk_GetGameConfigVTableAll = &s2sdk_GetGameConfigVTableAll

//go:linkname s2sdk_GetGameConfigSignatureAll __package__/s2sdk._GetGameConfigSignatureAll
var s2sdk_GetGameConfigSignatureAll func(name string) uintptr
var S2sdk_GetGameConfigSignatureAll = &s2sdk_GetGameConfigSignatureAll

//go:linkname s2sdk_RegisterLoggingChannel __package__/s2sdk._RegisterLoggingChannel
var s2sdk_RegisterLoggingChannel func(name string, flags int32, verbosity s2sdk.LoggingVerbosity, color plugify.Vector4) int32
var S2sdk_RegisterLoggingChannel = &s2sdk_RegisterLoggingChannel

//go:linkname s2sdk_AddLoggerTagToChannel __package__/s2sdk._AddLoggerTagToChannel
var s2sdk_AddLoggerTagToChannel func(channelID int32, tagName string)
var S2sdk_AddLoggerTagToChannel = &s2sdk_AddLoggerTagToChannel

//go:linkname s2sdk_HasLoggerTag __package__/s2sdk._HasLoggerTag
var s2sdk_HasLoggerTag func(channelID int32, tag string) bool
var S2sdk_HasLoggerTag = &s2sdk_HasLoggerTag

//go:linkname s2sdk_IsLoggerChannelEnabledBySeverity __package__/s2sdk._IsLoggerChannelEnabledBySeverity
var s2sdk_IsLoggerChannelEnabledBySeverity func(channelID int32, severity s2sdk.LoggingSeverity) bool
var S2sdk_IsLoggerChannelEnabledBySeverity = &s2sdk_IsLoggerChannelEnabledBySeverity

//go:linkname s2sdk_IsLoggerChannelEnabledByVerbosity __package__/s2sdk._IsLoggerChannelEnabledByVerbosity
var s2sdk_IsLoggerChannelEnabledByVerbosity func(channelID int32, verbosity s2sdk.LoggingVerbosity) bool
var S2sdk_IsLoggerChannelEnabledByVerbosity = &s2sdk_IsLoggerChannelEnabledByVerbosity

//go:linkname s2sdk_GetLoggerChannelVerbosity __package__/s2sdk._GetLoggerChannelVerbosity
var s2sdk_GetLoggerChannelVerbosity func(channelID int32) int32
var S2sdk_GetLoggerChannelVerbosity = &s2sdk_GetLoggerChannelVerbosity

//go:linkname s2sdk_SetLoggerChannelVerbosity __package__/s2sdk._SetLoggerChannelVerbosity
var s2sdk_SetLoggerChannelVerbosity func(channelID int32, verbosity s2sdk.LoggingVerbosity)
var S2sdk_SetLoggerChannelVerbosity = &s2sdk_SetLoggerChannelVerbosity

//go:linkname s2sdk_SetLoggerChannelVerbosityByName __package__/s2sdk._SetLoggerChannelVerbosityByName
var s2sdk_SetLoggerChannelVerbosityByName func(channelID int32, name string, verbosity s2sdk.LoggingVerbosity)
var S2sdk_SetLoggerChannelVerbosityByName = &s2sdk_SetLoggerChannelVerbosityByName

//go:linkname s2sdk_SetLoggerChannelVerbosityByTag __package__/s2sdk._SetLoggerChannelVerbosityByTag
var s2sdk_SetLoggerChannelVerbosityByTag func(channelID int32, tag string, verbosity s2sdk.LoggingVerbosity)
var S2sdk_SetLoggerChannelVerbosityByTag = &s2sdk_SetLoggerChannelVerbosityByTag

//go:linkname s2sdk_GetLoggerChannelColor __package__/s2sdk._GetLoggerChannelColor
var s2sdk_GetLoggerChannelColor func(channelID int32) plugify.Vector4
var S2sdk_GetLoggerChannelColor = &s2sdk_GetLoggerChannelColor

//go:linkname s2sdk_SetLoggerChannelColor __package__/s2sdk._SetLoggerChannelColor
var s2sdk_SetLoggerChannelColor func(channelID int32, color plugify.Vector4)
var S2sdk_SetLoggerChannelColor = &s2sdk_SetLoggerChannelColor

//go:linkname s2sdk_GetLoggerChannelFlags __package__/s2sdk._GetLoggerChannelFlags
var s2sdk_GetLoggerChannelFlags func(channelID int32) int32
var S2sdk_GetLoggerChannelFlags = &s2sdk_GetLoggerChannelFlags

//go:linkname s2sdk_SetLoggerChannelFlags __package__/s2sdk._SetLoggerChannelFlags
var s2sdk_SetLoggerChannelFlags func(channelID int32, eFlags s2sdk.LoggingChannelFlags)
var S2sdk_SetLoggerChannelFlags = &s2sdk_SetLoggerChannelFlags

//go:linkname s2sdk_Log __package__/s2sdk._Log
var s2sdk_Log func(channelID int32, severity s2sdk.LoggingSeverity, message string) int32
var S2sdk_Log = &s2sdk_Log

//go:linkname s2sdk_LogColored __package__/s2sdk._LogColored
var s2sdk_LogColored func(channelID int32, severity s2sdk.LoggingSeverity, color plugify.Vector4, message string) int32
var S2sdk_LogColored = &s2sdk_LogColored

//go:linkname s2sdk_LogFull __package__/s2sdk._LogFull
var s2sdk_LogFull func(channelID int32, severity s2sdk.LoggingSeverity, file string, line int32, function string, message string) int32
var S2sdk_LogFull = &s2sdk_LogFull

//go:linkname s2sdk_LogFullColored __package__/s2sdk._LogFullColored
var s2sdk_LogFullColored func(channelID int32, severity s2sdk.LoggingSeverity, file string, line int32, function string, color plugify.Vector4, message string) int32
var S2sdk_LogFullColored = &s2sdk_LogFullColored

//go:linkname s2sdk_GetEntityAttachmentAngles __package__/s2sdk._GetEntityAttachmentAngles
var s2sdk_GetEntityAttachmentAngles func(entityHandle int32, attachmentIndex int32) plugify.Vector3
var S2sdk_GetEntityAttachmentAngles = &s2sdk_GetEntityAttachmentAngles

//go:linkname s2sdk_GetEntityAttachmentForward __package__/s2sdk._GetEntityAttachmentForward
var s2sdk_GetEntityAttachmentForward func(entityHandle int32, attachmentIndex int32) plugify.Vector3
var S2sdk_GetEntityAttachmentForward = &s2sdk_GetEntityAttachmentForward

//go:linkname s2sdk_GetEntityAttachmentOrigin __package__/s2sdk._GetEntityAttachmentOrigin
var s2sdk_GetEntityAttachmentOrigin func(entityHandle int32, attachmentIndex int32) plugify.Vector3
var S2sdk_GetEntityAttachmentOrigin = &s2sdk_GetEntityAttachmentOrigin

//go:linkname s2sdk_GetEntityMaterialGroupHash __package__/s2sdk._GetEntityMaterialGroupHash
var s2sdk_GetEntityMaterialGroupHash func(entityHandle int32) uint32
var S2sdk_GetEntityMaterialGroupHash = &s2sdk_GetEntityMaterialGroupHash

//go:linkname s2sdk_GetEntityMaterialGroupMask __package__/s2sdk._GetEntityMaterialGroupMask
var s2sdk_GetEntityMaterialGroupMask func(entityHandle int32) uint64
var S2sdk_GetEntityMaterialGroupMask = &s2sdk_GetEntityMaterialGroupMask

//go:linkname s2sdk_GetEntityModelScale __package__/s2sdk._GetEntityModelScale
var s2sdk_GetEntityModelScale func(entityHandle int32) float32
var S2sdk_GetEntityModelScale = &s2sdk_GetEntityModelScale

//go:linkname s2sdk_GetEntityRenderAlpha __package__/s2sdk._GetEntityRenderAlpha
var s2sdk_GetEntityRenderAlpha func(entityHandle int32) int32
var S2sdk_GetEntityRenderAlpha = &s2sdk_GetEntityRenderAlpha

//go:linkname s2sdk_GetEntityRenderColor2 __package__/s2sdk._GetEntityRenderColor2
var s2sdk_GetEntityRenderColor2 func(entityHandle int32) plugify.Vector3
var S2sdk_GetEntityRenderColor2 = &s2sdk_GetEntityRenderColor2

//go:linkname s2sdk_ScriptLookupAttachment __package__/s2sdk._ScriptLookupAttachment
var s2sdk_ScriptLookupAttachment func(entityHandle int32, attachmentName string) int32
var S2sdk_ScriptLookupAttachment = &s2sdk_ScriptLookupAttachment

//go:linkname s2sdk_SetEntityBodygroup __package__/s2sdk._SetEntityBodygroup
var s2sdk_SetEntityBodygroup func(entityHandle int32, group int32, value int32)
var S2sdk_SetEntityBodygroup = &s2sdk_SetEntityBodygroup

//go:linkname s2sdk_SetEntityBodygroupByName __package__/s2sdk._SetEntityBodygroupByName
var s2sdk_SetEntityBodygroupByName func(entityHandle int32, name string, value int32)
var S2sdk_SetEntityBodygroupByName = &s2sdk_SetEntityBodygroupByName

//go:linkname s2sdk_SetEntityLightGroup __package__/s2sdk._SetEntityLightGroup
var s2sdk_SetEntityLightGroup func(entityHandle int32, lightGroup string)
var S2sdk_SetEntityLightGroup = &s2sdk_SetEntityLightGroup

//go:linkname s2sdk_SetEntityMaterialGroup __package__/s2sdk._SetEntityMaterialGroup
var s2sdk_SetEntityMaterialGroup func(entityHandle int32, materialGroup string)
var S2sdk_SetEntityMaterialGroup = &s2sdk_SetEntityMaterialGroup

//go:linkname s2sdk_SetEntityMaterialGroupHash __package__/s2sdk._SetEntityMaterialGroupHash
var s2sdk_SetEntityMaterialGroupHash func(entityHandle int32, hash uint32)
var S2sdk_SetEntityMaterialGroupHash = &s2sdk_SetEntityMaterialGroupHash

//go:linkname s2sdk_SetEntityMaterialGroupMask __package__/s2sdk._SetEntityMaterialGroupMask
var s2sdk_SetEntityMaterialGroupMask func(entityHandle int32, mask uint64)
var S2sdk_SetEntityMaterialGroupMask = &s2sdk_SetEntityMaterialGroupMask

//go:linkname s2sdk_SetEntityModelScale __package__/s2sdk._SetEntityModelScale
var s2sdk_SetEntityModelScale func(entityHandle int32, scale float32)
var S2sdk_SetEntityModelScale = &s2sdk_SetEntityModelScale

//go:linkname s2sdk_SetEntityRenderAlpha __package__/s2sdk._SetEntityRenderAlpha
var s2sdk_SetEntityRenderAlpha func(entityHandle int32, alpha int32)
var S2sdk_SetEntityRenderAlpha = &s2sdk_SetEntityRenderAlpha

//go:linkname s2sdk_SetEntityRenderColor2 __package__/s2sdk._SetEntityRenderColor2
var s2sdk_SetEntityRenderColor2 func(entityHandle int32, r int32, g int32, b int32)
var S2sdk_SetEntityRenderColor2 = &s2sdk_SetEntityRenderColor2

//go:linkname s2sdk_SetEntityRenderMode2 __package__/s2sdk._SetEntityRenderMode2
var s2sdk_SetEntityRenderMode2 func(entityHandle int32, mode int32)
var S2sdk_SetEntityRenderMode2 = &s2sdk_SetEntityRenderMode2

//go:linkname s2sdk_SetEntitySingleMeshGroup __package__/s2sdk._SetEntitySingleMeshGroup
var s2sdk_SetEntitySingleMeshGroup func(entityHandle int32, meshGroupName string)
var S2sdk_SetEntitySingleMeshGroup = &s2sdk_SetEntitySingleMeshGroup

//go:linkname s2sdk_SetEntitySize __package__/s2sdk._SetEntitySize
var s2sdk_SetEntitySize func(entityHandle int32, mins plugify.Vector3, maxs plugify.Vector3)
var S2sdk_SetEntitySize = &s2sdk_SetEntitySize

//go:linkname s2sdk_SetEntitySkin __package__/s2sdk._SetEntitySkin
var s2sdk_SetEntitySkin func(entityHandle int32, skin int32)
var S2sdk_SetEntitySkin = &s2sdk_SetEntitySkin

//go:linkname s2sdk_PanoramaSendYesNoVote __package__/s2sdk._PanoramaSendYesNoVote
var s2sdk_PanoramaSendYesNoVote func(duration float64, caller int32, voteTitle string, detailStr string, votePassTitle string, detailPassStr string, failReason s2sdk.VoteCreateFailed, filter uint64, result s2sdk.YesNoVoteResult, handler s2sdk.YesNoVoteHandler) bool
var S2sdk_PanoramaSendYesNoVote = &s2sdk_PanoramaSendYesNoVote

//go:linkname s2sdk_PanoramaSendYesNoVoteToAll __package__/s2sdk._PanoramaSendYesNoVoteToAll
var s2sdk_PanoramaSendYesNoVoteToAll func(duration float64, caller int32, voteTitle string, detailStr string, votePassTitle string, detailPassStr string, failReason s2sdk.VoteCreateFailed, result s2sdk.YesNoVoteResult, handler s2sdk.YesNoVoteHandler) bool
var S2sdk_PanoramaSendYesNoVoteToAll = &s2sdk_PanoramaSendYesNoVoteToAll

//go:linkname s2sdk_PanoramaRemovePlayerFromVote __package__/s2sdk._PanoramaRemovePlayerFromVote
var s2sdk_PanoramaRemovePlayerFromVote func(playerSlot int32)
var S2sdk_PanoramaRemovePlayerFromVote = &s2sdk_PanoramaRemovePlayerFromVote

//go:linkname s2sdk_PanoramaIsPlayerInVotePool __package__/s2sdk._PanoramaIsPlayerInVotePool
var s2sdk_PanoramaIsPlayerInVotePool func(playerSlot int32) bool
var S2sdk_PanoramaIsPlayerInVotePool = &s2sdk_PanoramaIsPlayerInVotePool

//go:linkname s2sdk_PanoramaRedrawVoteToClient __package__/s2sdk._PanoramaRedrawVoteToClient
var s2sdk_PanoramaRedrawVoteToClient func(playerSlot int32) bool
var S2sdk_PanoramaRedrawVoteToClient = &s2sdk_PanoramaRedrawVoteToClient

//go:linkname s2sdk_PanoramaIsVoteInProgress __package__/s2sdk._PanoramaIsVoteInProgress
var s2sdk_PanoramaIsVoteInProgress func() bool
var S2sdk_PanoramaIsVoteInProgress = &s2sdk_PanoramaIsVoteInProgress

//go:linkname s2sdk_PanoramaEndVote __package__/s2sdk._PanoramaEndVote
var s2sdk_PanoramaEndVote func(reason s2sdk.VoteEndReason)
var S2sdk_PanoramaEndVote = &s2sdk_PanoramaEndVote

//go:linkname s2sdk_GetSchemaOffset __package__/s2sdk._GetSchemaOffset
var s2sdk_GetSchemaOffset func(className string, memberName string) int32
var S2sdk_GetSchemaOffset = &s2sdk_GetSchemaOffset

//go:linkname s2sdk_GetSchemaChainOffset __package__/s2sdk._GetSchemaChainOffset
var s2sdk_GetSchemaChainOffset func(className string) int32
var S2sdk_GetSchemaChainOffset = &s2sdk_GetSchemaChainOffset

//go:linkname s2sdk_IsSchemaFieldNetworked __package__/s2sdk._IsSchemaFieldNetworked
var s2sdk_IsSchemaFieldNetworked func(className string, memberName string) bool
var S2sdk_IsSchemaFieldNetworked = &s2sdk_IsSchemaFieldNetworked

//go:linkname s2sdk_GetSchemaClassSize __package__/s2sdk._GetSchemaClassSize
var s2sdk_GetSchemaClassSize func(className string) int32
var S2sdk_GetSchemaClassSize = &s2sdk_GetSchemaClassSize

//go:linkname s2sdk_GetEntData2 __package__/s2sdk._GetEntData2
var s2sdk_GetEntData2 func(entity uintptr, offset int32, size int32) int64
var S2sdk_GetEntData2 = &s2sdk_GetEntData2

//go:linkname s2sdk_SetEntData2 __package__/s2sdk._SetEntData2
var s2sdk_SetEntData2 func(entity uintptr, offset int32, value int64, size int32, changeState bool, chainOffset int32)
var S2sdk_SetEntData2 = &s2sdk_SetEntData2

//go:linkname s2sdk_GetEntDataFloat2 __package__/s2sdk._GetEntDataFloat2
var s2sdk_GetEntDataFloat2 func(entity uintptr, offset int32, size int32) float64
var S2sdk_GetEntDataFloat2 = &s2sdk_GetEntDataFloat2

//go:linkname s2sdk_SetEntDataFloat2 __package__/s2sdk._SetEntDataFloat2
var s2sdk_SetEntDataFloat2 func(entity uintptr, offset int32, value float64, size int32, changeState bool, chainOffset int32)
var S2sdk_SetEntDataFloat2 = &s2sdk_SetEntDataFloat2

//go:linkname s2sdk_GetEntDataColor2 __package__/s2sdk._GetEntDataColor2
var s2sdk_GetEntDataColor2 func(entity uintptr, offset int32) plugify.Vector4
var S2sdk_GetEntDataColor2 = &s2sdk_GetEntDataColor2

//go:linkname s2sdk_SetEntDataColor2 __package__/s2sdk._SetEntDataColor2
var s2sdk_SetEntDataColor2 func(entity uintptr, offset int32, value plugify.Vector4, changeState bool, chainOffset int32)
var S2sdk_SetEntDataColor2 = &s2sdk_SetEntDataColor2

//go:linkname s2sdk_GetEntDataString2 __package__/s2sdk._GetEntDataString2
var s2sdk_GetEntDataString2 func(entity uintptr, offset int32) string
var S2sdk_GetEntDataString2 = &s2sdk_GetEntDataString2

//go:linkname s2sdk_SetEntDataString2 __package__/s2sdk._SetEntDataString2
var s2sdk_SetEntDataString2 func(entity uintptr, offset int32, value string, changeState bool, chainOffset int32)
var S2sdk_SetEntDataString2 = &s2sdk_SetEntDataString2

//go:linkname s2sdk_GetEntDataCString2 __package__/s2sdk._GetEntDataCString2
var s2sdk_GetEntDataCString2 func(entity uintptr, offset int32, size int32) string
var S2sdk_GetEntDataCString2 = &s2sdk_GetEntDataCString2

//go:linkname s2sdk_SetEntDataCString2 __package__/s2sdk._SetEntDataCString2
var s2sdk_SetEntDataCString2 func(entity uintptr, offset int32, value string, size int32, changeState bool, chainOffset int32)
var S2sdk_SetEntDataCString2 = &s2sdk_SetEntDataCString2

//go:linkname s2sdk_GetEntDataVector3D2 __package__/s2sdk._GetEntDataVector3D2
var s2sdk_GetEntDataVector3D2 func(entity uintptr, offset int32) plugify.Vector3
var S2sdk_GetEntDataVector3D2 = &s2sdk_GetEntDataVector3D2

//go:linkname s2sdk_SetEntDataVector3D2 __package__/s2sdk._SetEntDataVector3D2
var s2sdk_SetEntDataVector3D2 func(entity uintptr, offset int32, value plugify.Vector3, changeState bool, chainOffset int32)
var S2sdk_SetEntDataVector3D2 = &s2sdk_SetEntDataVector3D2

//go:linkname s2sdk_GetEntDataVector4D2 __package__/s2sdk._GetEntDataVector4D2
var s2sdk_GetEntDataVector4D2 func(entity uintptr, offset int32) plugify.Vector4
var S2sdk_GetEntDataVector4D2 = &s2sdk_GetEntDataVector4D2

//go:linkname s2sdk_SetEntDataVector4D2 __package__/s2sdk._SetEntDataVector4D2
var s2sdk_SetEntDataVector4D2 func(entity uintptr, offset int32, value plugify.Vector4, changeState bool, chainOffset int32)
var S2sdk_SetEntDataVector4D2 = &s2sdk_SetEntDataVector4D2

//go:linkname s2sdk_GetEntDataVector2D2 __package__/s2sdk._GetEntDataVector2D2
var s2sdk_GetEntDataVector2D2 func(entity uintptr, offset int32) plugify.Vector2
var S2sdk_GetEntDataVector2D2 = &s2sdk_GetEntDataVector2D2

//go:linkname s2sdk_SetEntDataVector2D2 __package__/s2sdk._SetEntDataVector2D2
var s2sdk_SetEntDataVector2D2 func(entity uintptr, offset int32, value plugify.Vector2, changeState bool, chainOffset int32)
var S2sdk_SetEntDataVector2D2 = &s2sdk_SetEntDataVector2D2

//go:linkname s2sdk_GetEntDataEnt2 __package__/s2sdk._GetEntDataEnt2
var s2sdk_GetEntDataEnt2 func(entity uintptr, offset int32) int32
var S2sdk_GetEntDataEnt2 = &s2sdk_GetEntDataEnt2

//go:linkname s2sdk_SetEntDataEnt2 __package__/s2sdk._SetEntDataEnt2
var s2sdk_SetEntDataEnt2 func(entity uintptr, offset int32, value int32, changeState bool, chainOffset int32)
var S2sdk_SetEntDataEnt2 = &s2sdk_SetEntDataEnt2

//go:linkname s2sdk_ChangeEntityState2 __package__/s2sdk._ChangeEntityState2
var s2sdk_ChangeEntityState2 func(entity uintptr, offset int32, chainOffset int32)
var S2sdk_ChangeEntityState2 = &s2sdk_ChangeEntityState2

//go:linkname s2sdk_GetEntData __package__/s2sdk._GetEntData
var s2sdk_GetEntData func(entityHandle int32, offset int32, size int32) int64
var S2sdk_GetEntData = &s2sdk_GetEntData

//go:linkname s2sdk_SetEntData __package__/s2sdk._SetEntData
var s2sdk_SetEntData func(entityHandle int32, offset int32, value int64, size int32, changeState bool, chainOffset int32)
var S2sdk_SetEntData = &s2sdk_SetEntData

//go:linkname s2sdk_GetEntDataFloat __package__/s2sdk._GetEntDataFloat
var s2sdk_GetEntDataFloat func(entityHandle int32, offset int32, size int32) float64
var S2sdk_GetEntDataFloat = &s2sdk_GetEntDataFloat

//go:linkname s2sdk_SetEntDataFloat __package__/s2sdk._SetEntDataFloat
var s2sdk_SetEntDataFloat func(entityHandle int32, offset int32, value float64, size int32, changeState bool, chainOffset int32)
var S2sdk_SetEntDataFloat = &s2sdk_SetEntDataFloat

//go:linkname s2sdk_GetEntDataColor __package__/s2sdk._GetEntDataColor
var s2sdk_GetEntDataColor func(entityHandle int32, offset int32) plugify.Vector4
var S2sdk_GetEntDataColor = &s2sdk_GetEntDataColor

//go:linkname s2sdk_SetEntDataColor __package__/s2sdk._SetEntDataColor
var s2sdk_SetEntDataColor func(entityHandle int32, offset int32, value plugify.Vector4, changeState bool, chainOffset int32)
var S2sdk_SetEntDataColor = &s2sdk_SetEntDataColor

//go:linkname s2sdk_GetEntDataString __package__/s2sdk._GetEntDataString
var s2sdk_GetEntDataString func(entityHandle int32, offset int32) string
var S2sdk_GetEntDataString = &s2sdk_GetEntDataString

//go:linkname s2sdk_SetEntDataString __package__/s2sdk._SetEntDataString
var s2sdk_SetEntDataString func(entityHandle int32, offset int32, value string, changeState bool, chainOffset int32)
var S2sdk_SetEntDataString = &s2sdk_SetEntDataString

//go:linkname s2sdk_GetEntDataCString __package__/s2sdk._GetEntDataCString
var s2sdk_GetEntDataCString func(entityHandle int32, offset int32, size int32) string
var S2sdk_GetEntDataCString = &s2sdk_GetEntDataCString

//go:linkname s2sdk_SetEntDataCString __package__/s2sdk._SetEntDataCString
var s2sdk_SetEntDataCString func(entityHandle int32, offset int32, value string, size int32, changeState bool, chainOffset int32)
var S2sdk_SetEntDataCString = &s2sdk_SetEntDataCString

//go:linkname s2sdk_GetEntDataVector3D __package__/s2sdk._GetEntDataVector3D
var s2sdk_GetEntDataVector3D func(entityHandle int32, offset int32) plugify.Vector3
var S2sdk_GetEntDataVector3D = &s2sdk_GetEntDataVector3D

//go:linkname s2sdk_SetEntDataVector3D __package__/s2sdk._SetEntDataVector3D
var s2sdk_SetEntDataVector3D func(entityHandle int32, offset int32, value plugify.Vector3, changeState bool, chainOffset int32)
var S2sdk_SetEntDataVector3D = &s2sdk_SetEntDataVector3D

//go:linkname s2sdk_GetEntDataVector4D __package__/s2sdk._GetEntDataVector4D
var s2sdk_GetEntDataVector4D func(entityHandle int32, offset int32) plugify.Vector4
var S2sdk_GetEntDataVector4D = &s2sdk_GetEntDataVector4D

//go:linkname s2sdk_SetEntDataVector4D __package__/s2sdk._SetEntDataVector4D
var s2sdk_SetEntDataVector4D func(entityHandle int32, offset int32, value plugify.Vector4, changeState bool, chainOffset int32)
var S2sdk_SetEntDataVector4D = &s2sdk_SetEntDataVector4D

//go:linkname s2sdk_GetEntDataVector2D __package__/s2sdk._GetEntDataVector2D
var s2sdk_GetEntDataVector2D func(entityHandle int32, offset int32) plugify.Vector2
var S2sdk_GetEntDataVector2D = &s2sdk_GetEntDataVector2D

//go:linkname s2sdk_SetEntDataVector2D __package__/s2sdk._SetEntDataVector2D
var s2sdk_SetEntDataVector2D func(entityHandle int32, offset int32, value plugify.Vector2, changeState bool, chainOffset int32)
var S2sdk_SetEntDataVector2D = &s2sdk_SetEntDataVector2D

//go:linkname s2sdk_GetEntDataEnt __package__/s2sdk._GetEntDataEnt
var s2sdk_GetEntDataEnt func(entityHandle int32, offset int32) int32
var S2sdk_GetEntDataEnt = &s2sdk_GetEntDataEnt

//go:linkname s2sdk_SetEntDataEnt __package__/s2sdk._SetEntDataEnt
var s2sdk_SetEntDataEnt func(entityHandle int32, offset int32, value int32, changeState bool, chainOffset int32)
var S2sdk_SetEntDataEnt = &s2sdk_SetEntDataEnt

//go:linkname s2sdk_ChangeEntityState __package__/s2sdk._ChangeEntityState
var s2sdk_ChangeEntityState func(entityHandle int32, offset int32, chainOffset int32)
var S2sdk_ChangeEntityState = &s2sdk_ChangeEntityState

//go:linkname s2sdk_GetEntSchemaArraySize2 __package__/s2sdk._GetEntSchemaArraySize2
var s2sdk_GetEntSchemaArraySize2 func(entity uintptr, className string, memberName string) int32
var S2sdk_GetEntSchemaArraySize2 = &s2sdk_GetEntSchemaArraySize2

//go:linkname s2sdk_GetEntSchema2 __package__/s2sdk._GetEntSchema2
var s2sdk_GetEntSchema2 func(entity uintptr, className string, memberName string, element int32) int64
var S2sdk_GetEntSchema2 = &s2sdk_GetEntSchema2

//go:linkname s2sdk_SetEntSchema2 __package__/s2sdk._SetEntSchema2
var s2sdk_SetEntSchema2 func(entity uintptr, className string, memberName string, value int64, changeState bool, element int32)
var S2sdk_SetEntSchema2 = &s2sdk_SetEntSchema2

//go:linkname s2sdk_GetEntSchemaFloat2 __package__/s2sdk._GetEntSchemaFloat2
var s2sdk_GetEntSchemaFloat2 func(entity uintptr, className string, memberName string, element int32) float64
var S2sdk_GetEntSchemaFloat2 = &s2sdk_GetEntSchemaFloat2

//go:linkname s2sdk_SetEntSchemaFloat2 __package__/s2sdk._SetEntSchemaFloat2
var s2sdk_SetEntSchemaFloat2 func(entity uintptr, className string, memberName string, value float64, changeState bool, element int32)
var S2sdk_SetEntSchemaFloat2 = &s2sdk_SetEntSchemaFloat2

//go:linkname s2sdk_GetEntSchemaColor2 __package__/s2sdk._GetEntSchemaColor2
var s2sdk_GetEntSchemaColor2 func(entity uintptr, className string, memberName string, element int32) plugify.Vector4
var S2sdk_GetEntSchemaColor2 = &s2sdk_GetEntSchemaColor2

//go:linkname s2sdk_SetEntSchemaColor2 __package__/s2sdk._SetEntSchemaColor2
var s2sdk_SetEntSchemaColor2 func(entity uintptr, className string, memberName string, value plugify.Vector4, changeState bool, element int32)
var S2sdk_SetEntSchemaColor2 = &s2sdk_SetEntSchemaColor2

//go:linkname s2sdk_GetEntSchemaString2 __package__/s2sdk._GetEntSchemaString2
var s2sdk_GetEntSchemaString2 func(entity uintptr, className string, memberName string, element int32) string
var S2sdk_GetEntSchemaString2 = &s2sdk_GetEntSchemaString2

//go:linkname s2sdk_SetEntSchemaString2 __package__/s2sdk._SetEntSchemaString2
var s2sdk_SetEntSchemaString2 func(entity uintptr, className string, memberName string, value string, changeState bool, element int32)
var S2sdk_SetEntSchemaString2 = &s2sdk_SetEntSchemaString2

//go:linkname s2sdk_GetEntSchemaVector3D2 __package__/s2sdk._GetEntSchemaVector3D2
var s2sdk_GetEntSchemaVector3D2 func(entity uintptr, className string, memberName string, element int32) plugify.Vector3
var S2sdk_GetEntSchemaVector3D2 = &s2sdk_GetEntSchemaVector3D2

//go:linkname s2sdk_SetEntSchemaVector3D2 __package__/s2sdk._SetEntSchemaVector3D2
var s2sdk_SetEntSchemaVector3D2 func(entity uintptr, className string, memberName string, value plugify.Vector3, changeState bool, element int32)
var S2sdk_SetEntSchemaVector3D2 = &s2sdk_SetEntSchemaVector3D2

//go:linkname s2sdk_GetEntSchemaVector2D2 __package__/s2sdk._GetEntSchemaVector2D2
var s2sdk_GetEntSchemaVector2D2 func(entity uintptr, className string, memberName string, element int32) plugify.Vector2
var S2sdk_GetEntSchemaVector2D2 = &s2sdk_GetEntSchemaVector2D2

//go:linkname s2sdk_SetEntSchemaVector2D2 __package__/s2sdk._SetEntSchemaVector2D2
var s2sdk_SetEntSchemaVector2D2 func(entity uintptr, className string, memberName string, value plugify.Vector2, changeState bool, element int32)
var S2sdk_SetEntSchemaVector2D2 = &s2sdk_SetEntSchemaVector2D2

//go:linkname s2sdk_GetEntSchemaVector4D2 __package__/s2sdk._GetEntSchemaVector4D2
var s2sdk_GetEntSchemaVector4D2 func(entity uintptr, className string, memberName string, element int32) plugify.Vector4
var S2sdk_GetEntSchemaVector4D2 = &s2sdk_GetEntSchemaVector4D2

//go:linkname s2sdk_SetEntSchemaVector4D2 __package__/s2sdk._SetEntSchemaVector4D2
var s2sdk_SetEntSchemaVector4D2 func(entity uintptr, className string, memberName string, value plugify.Vector4, changeState bool, element int32)
var S2sdk_SetEntSchemaVector4D2 = &s2sdk_SetEntSchemaVector4D2

//go:linkname s2sdk_GetEntSchemaEnt2 __package__/s2sdk._GetEntSchemaEnt2
var s2sdk_GetEntSchemaEnt2 func(entity uintptr, className string, memberName string, element int32) int32
var S2sdk_GetEntSchemaEnt2 = &s2sdk_GetEntSchemaEnt2

//go:linkname s2sdk_SetEntSchemaEnt2 __package__/s2sdk._SetEntSchemaEnt2
var s2sdk_SetEntSchemaEnt2 func(entity uintptr, className string, memberName string, value int32, changeState bool, element int32)
var S2sdk_SetEntSchemaEnt2 = &s2sdk_SetEntSchemaEnt2

//go:linkname s2sdk_PushEntSchemaEnt2 __package__/s2sdk._PushEntSchemaEnt2
var s2sdk_PushEntSchemaEnt2 func(entity uintptr, className string, memberName string, value int32, changeState bool)
var S2sdk_PushEntSchemaEnt2 = &s2sdk_PushEntSchemaEnt2

//go:linkname s2sdk_EraseEntSchemaEnt2 __package__/s2sdk._EraseEntSchemaEnt2
var s2sdk_EraseEntSchemaEnt2 func(entity uintptr, className string, memberName string, element int32, changeState bool)
var S2sdk_EraseEntSchemaEnt2 = &s2sdk_EraseEntSchemaEnt2

//go:linkname s2sdk_NetworkStateChanged2 __package__/s2sdk._NetworkStateChanged2
var s2sdk_NetworkStateChanged2 func(entity uintptr, className string, memberName string)
var S2sdk_NetworkStateChanged2 = &s2sdk_NetworkStateChanged2

//go:linkname s2sdk_GetEntSchemaArraySize __package__/s2sdk._GetEntSchemaArraySize
var s2sdk_GetEntSchemaArraySize func(entityHandle int32, className string, memberName string) int32
var S2sdk_GetEntSchemaArraySize = &s2sdk_GetEntSchemaArraySize

//go:linkname s2sdk_GetEntSchema __package__/s2sdk._GetEntSchema
var s2sdk_GetEntSchema func(entityHandle int32, className string, memberName string, element int32) int64
var S2sdk_GetEntSchema = &s2sdk_GetEntSchema

//go:linkname s2sdk_SetEntSchema __package__/s2sdk._SetEntSchema
var s2sdk_SetEntSchema func(entityHandle int32, className string, memberName string, value int64, changeState bool, element int32)
var S2sdk_SetEntSchema = &s2sdk_SetEntSchema

//go:linkname s2sdk_GetEntSchemaFloat __package__/s2sdk._GetEntSchemaFloat
var s2sdk_GetEntSchemaFloat func(entityHandle int32, className string, memberName string, element int32) float64
var S2sdk_GetEntSchemaFloat = &s2sdk_GetEntSchemaFloat

//go:linkname s2sdk_SetEntSchemaFloat __package__/s2sdk._SetEntSchemaFloat
var s2sdk_SetEntSchemaFloat func(entityHandle int32, className string, memberName string, value float64, changeState bool, element int32)
var S2sdk_SetEntSchemaFloat = &s2sdk_SetEntSchemaFloat

//go:linkname s2sdk_GetEntSchemaColor __package__/s2sdk._GetEntSchemaColor
var s2sdk_GetEntSchemaColor func(entityHandle int32, className string, memberName string, element int32) plugify.Vector4
var S2sdk_GetEntSchemaColor = &s2sdk_GetEntSchemaColor

//go:linkname s2sdk_SetEntSchemaColor __package__/s2sdk._SetEntSchemaColor
var s2sdk_SetEntSchemaColor func(entityHandle int32, className string, memberName string, value plugify.Vector4, changeState bool, element int32)
var S2sdk_SetEntSchemaColor = &s2sdk_SetEntSchemaColor

//go:linkname s2sdk_GetEntSchemaString __package__/s2sdk._GetEntSchemaString
var s2sdk_GetEntSchemaString func(entityHandle int32, className string, memberName string, element int32) string
var S2sdk_GetEntSchemaString = &s2sdk_GetEntSchemaString

//go:linkname s2sdk_SetEntSchemaString __package__/s2sdk._SetEntSchemaString
var s2sdk_SetEntSchemaString func(entityHandle int32, className string, memberName string, value string, changeState bool, element int32)
var S2sdk_SetEntSchemaString = &s2sdk_SetEntSchemaString

//go:linkname s2sdk_GetEntSchemaVector3D __package__/s2sdk._GetEntSchemaVector3D
var s2sdk_GetEntSchemaVector3D func(entityHandle int32, className string, memberName string, element int32) plugify.Vector3
var S2sdk_GetEntSchemaVector3D = &s2sdk_GetEntSchemaVector3D

//go:linkname s2sdk_SetEntSchemaVector3D __package__/s2sdk._SetEntSchemaVector3D
var s2sdk_SetEntSchemaVector3D func(entityHandle int32, className string, memberName string, value plugify.Vector3, changeState bool, element int32)
var S2sdk_SetEntSchemaVector3D = &s2sdk_SetEntSchemaVector3D

//go:linkname s2sdk_GetEntSchemaVector2D __package__/s2sdk._GetEntSchemaVector2D
var s2sdk_GetEntSchemaVector2D func(entityHandle int32, className string, memberName string, element int32) plugify.Vector2
var S2sdk_GetEntSchemaVector2D = &s2sdk_GetEntSchemaVector2D

//go:linkname s2sdk_SetEntSchemaVector2D __package__/s2sdk._SetEntSchemaVector2D
var s2sdk_SetEntSchemaVector2D func(entityHandle int32, className string, memberName string, value plugify.Vector2, changeState bool, element int32)
var S2sdk_SetEntSchemaVector2D = &s2sdk_SetEntSchemaVector2D

//go:linkname s2sdk_GetEntSchemaVector4D __package__/s2sdk._GetEntSchemaVector4D
var s2sdk_GetEntSchemaVector4D func(entityHandle int32, className string, memberName string, element int32) plugify.Vector4
var S2sdk_GetEntSchemaVector4D = &s2sdk_GetEntSchemaVector4D

//go:linkname s2sdk_SetEntSchemaVector4D __package__/s2sdk._SetEntSchemaVector4D
var s2sdk_SetEntSchemaVector4D func(entityHandle int32, className string, memberName string, value plugify.Vector4, changeState bool, element int32)
var S2sdk_SetEntSchemaVector4D = &s2sdk_SetEntSchemaVector4D

//go:linkname s2sdk_GetEntSchemaEnt __package__/s2sdk._GetEntSchemaEnt
var s2sdk_GetEntSchemaEnt func(entityHandle int32, className string, memberName string, element int32) int32
var S2sdk_GetEntSchemaEnt = &s2sdk_GetEntSchemaEnt

//go:linkname s2sdk_SetEntSchemaEnt __package__/s2sdk._SetEntSchemaEnt
var s2sdk_SetEntSchemaEnt func(entityHandle int32, className string, memberName string, value int32, changeState bool, element int32)
var S2sdk_SetEntSchemaEnt = &s2sdk_SetEntSchemaEnt

//go:linkname s2sdk_PushEntSchemaEnt __package__/s2sdk._PushEntSchemaEnt
var s2sdk_PushEntSchemaEnt func(entityHandle int32, className string, memberName string, value int32, changeState bool)
var S2sdk_PushEntSchemaEnt = &s2sdk_PushEntSchemaEnt

//go:linkname s2sdk_EraseEntSchemaEnt __package__/s2sdk._EraseEntSchemaEnt
var s2sdk_EraseEntSchemaEnt func(entityHandle int32, className string, memberName string, element int32, changeState bool)
var S2sdk_EraseEntSchemaEnt = &s2sdk_EraseEntSchemaEnt

//go:linkname s2sdk_NetworkStateChanged __package__/s2sdk._NetworkStateChanged
var s2sdk_NetworkStateChanged func(entityHandle int32, className string, memberName string)
var S2sdk_NetworkStateChanged = &s2sdk_NetworkStateChanged

//go:linkname s2sdk_CreateTimer __package__/s2sdk._CreateTimer
var s2sdk_CreateTimer func(delay float64, callback s2sdk.TimerCallback, flags s2sdk.TimerFlag, userData []any) uint32
var S2sdk_CreateTimer = &s2sdk_CreateTimer

//go:linkname s2sdk_KillsTimer __package__/s2sdk._KillsTimer
var s2sdk_KillsTimer func(timer uint32)
var S2sdk_KillsTimer = &s2sdk_KillsTimer

//go:linkname s2sdk_RescheduleTimer __package__/s2sdk._RescheduleTimer
var s2sdk_RescheduleTimer func(timer uint32, newDaly float64)
var S2sdk_RescheduleTimer = &s2sdk_RescheduleTimer

//go:linkname s2sdk_GetTickInterval __package__/s2sdk._GetTickInterval
var s2sdk_GetTickInterval func() float64
var S2sdk_GetTickInterval = &s2sdk_GetTickInterval

//go:linkname s2sdk_GetTickedTime __package__/s2sdk._GetTickedTime
var s2sdk_GetTickedTime func() float64
var S2sdk_GetTickedTime = &s2sdk_GetTickedTime

//go:linkname s2sdk_OnClientConnect_Register __package__/s2sdk._OnClientConnect_Register
var s2sdk_OnClientConnect_Register func(callback s2sdk.OnClientConnectCallback)
var S2sdk_OnClientConnect_Register = &s2sdk_OnClientConnect_Register

//go:linkname s2sdk_OnClientConnect_Unregister __package__/s2sdk._OnClientConnect_Unregister
var s2sdk_OnClientConnect_Unregister func(callback s2sdk.OnClientConnectCallback)
var S2sdk_OnClientConnect_Unregister = &s2sdk_OnClientConnect_Unregister

//go:linkname s2sdk_OnClientConnect_Post_Register __package__/s2sdk._OnClientConnect_Post_Register
var s2sdk_OnClientConnect_Post_Register func(callback s2sdk.OnClientConnect_PostCallback)
var S2sdk_OnClientConnect_Post_Register = &s2sdk_OnClientConnect_Post_Register

//go:linkname s2sdk_OnClientConnect_Post_Unregister __package__/s2sdk._OnClientConnect_Post_Unregister
var s2sdk_OnClientConnect_Post_Unregister func(callback s2sdk.OnClientConnect_PostCallback)
var S2sdk_OnClientConnect_Post_Unregister = &s2sdk_OnClientConnect_Post_Unregister

//go:linkname s2sdk_OnClientConnected_Register __package__/s2sdk._OnClientConnected_Register
var s2sdk_OnClientConnected_Register func(callback s2sdk.OnClientConnectedCallback)
var S2sdk_OnClientConnected_Register = &s2sdk_OnClientConnected_Register

//go:linkname s2sdk_OnClientConnected_Unregister __package__/s2sdk._OnClientConnected_Unregister
var s2sdk_OnClientConnected_Unregister func(callback s2sdk.OnClientConnectedCallback)
var S2sdk_OnClientConnected_Unregister = &s2sdk_OnClientConnected_Unregister

//go:linkname s2sdk_OnClientPutInServer_Register __package__/s2sdk._OnClientPutInServer_Register
var s2sdk_OnClientPutInServer_Register func(callback s2sdk.OnClientPutInServerCallback)
var S2sdk_OnClientPutInServer_Register = &s2sdk_OnClientPutInServer_Register

//go:linkname s2sdk_OnClientPutInServer_Unregister __package__/s2sdk._OnClientPutInServer_Unregister
var s2sdk_OnClientPutInServer_Unregister func(callback s2sdk.OnClientPutInServerCallback)
var S2sdk_OnClientPutInServer_Unregister = &s2sdk_OnClientPutInServer_Unregister

//go:linkname s2sdk_OnClientDisconnect_Register __package__/s2sdk._OnClientDisconnect_Register
var s2sdk_OnClientDisconnect_Register func(callback s2sdk.OnClientDisconnectCallback)
var S2sdk_OnClientDisconnect_Register = &s2sdk_OnClientDisconnect_Register

//go:linkname s2sdk_OnClientDisconnect_Unregister __package__/s2sdk._OnClientDisconnect_Unregister
var s2sdk_OnClientDisconnect_Unregister func(callback s2sdk.OnClientDisconnectCallback)
var S2sdk_OnClientDisconnect_Unregister = &s2sdk_OnClientDisconnect_Unregister

//go:linkname s2sdk_OnClientDisconnect_Post_Register __package__/s2sdk._OnClientDisconnect_Post_Register
var s2sdk_OnClientDisconnect_Post_Register func(callback s2sdk.OnClientDisconnect_PostCallback)
var S2sdk_OnClientDisconnect_Post_Register = &s2sdk_OnClientDisconnect_Post_Register

//go:linkname s2sdk_OnClientDisconnect_Post_Unregister __package__/s2sdk._OnClientDisconnect_Post_Unregister
var s2sdk_OnClientDisconnect_Post_Unregister func(callback s2sdk.OnClientDisconnect_PostCallback)
var S2sdk_OnClientDisconnect_Post_Unregister = &s2sdk_OnClientDisconnect_Post_Unregister

//go:linkname s2sdk_OnClientActive_Register __package__/s2sdk._OnClientActive_Register
var s2sdk_OnClientActive_Register func(callback s2sdk.OnClientActiveCallback)
var S2sdk_OnClientActive_Register = &s2sdk_OnClientActive_Register

//go:linkname s2sdk_OnClientActive_Unregister __package__/s2sdk._OnClientActive_Unregister
var s2sdk_OnClientActive_Unregister func(callback s2sdk.OnClientActiveCallback)
var S2sdk_OnClientActive_Unregister = &s2sdk_OnClientActive_Unregister

//go:linkname s2sdk_OnClientFullyConnect_Register __package__/s2sdk._OnClientFullyConnect_Register
var s2sdk_OnClientFullyConnect_Register func(callback s2sdk.OnClientFullyConnectCallback)
var S2sdk_OnClientFullyConnect_Register = &s2sdk_OnClientFullyConnect_Register

//go:linkname s2sdk_OnClientFullyConnect_Unregister __package__/s2sdk._OnClientFullyConnect_Unregister
var s2sdk_OnClientFullyConnect_Unregister func(callback s2sdk.OnClientFullyConnectCallback)
var S2sdk_OnClientFullyConnect_Unregister = &s2sdk_OnClientFullyConnect_Unregister

//go:linkname s2sdk_OnClientSettingsChanged_Register __package__/s2sdk._OnClientSettingsChanged_Register
var s2sdk_OnClientSettingsChanged_Register func(callback s2sdk.OnClientSettingsChangedCallback)
var S2sdk_OnClientSettingsChanged_Register = &s2sdk_OnClientSettingsChanged_Register

//go:linkname s2sdk_OnClientSettingsChanged_Unregister __package__/s2sdk._OnClientSettingsChanged_Unregister
var s2sdk_OnClientSettingsChanged_Unregister func(callback s2sdk.OnClientSettingsChangedCallback)
var S2sdk_OnClientSettingsChanged_Unregister = &s2sdk_OnClientSettingsChanged_Unregister

//go:linkname s2sdk_OnClientAuthenticated_Register __package__/s2sdk._OnClientAuthenticated_Register
var s2sdk_OnClientAuthenticated_Register func(callback s2sdk.OnClientAuthenticatedCallback)
var S2sdk_OnClientAuthenticated_Register = &s2sdk_OnClientAuthenticated_Register

//go:linkname s2sdk_OnClientAuthenticated_Unregister __package__/s2sdk._OnClientAuthenticated_Unregister
var s2sdk_OnClientAuthenticated_Unregister func(callback s2sdk.OnClientAuthenticatedCallback)
var S2sdk_OnClientAuthenticated_Unregister = &s2sdk_OnClientAuthenticated_Unregister

//go:linkname s2sdk_OnRoundTerminated_Register __package__/s2sdk._OnRoundTerminated_Register
var s2sdk_OnRoundTerminated_Register func(callback s2sdk.OnRoundTerminatedCallback)
var S2sdk_OnRoundTerminated_Register = &s2sdk_OnRoundTerminated_Register

//go:linkname s2sdk_OnRoundTerminated_Unregister __package__/s2sdk._OnRoundTerminated_Unregister
var s2sdk_OnRoundTerminated_Unregister func(callback s2sdk.OnRoundTerminatedCallback)
var S2sdk_OnRoundTerminated_Unregister = &s2sdk_OnRoundTerminated_Unregister

//go:linkname s2sdk_OnEntityCreated_Register __package__/s2sdk._OnEntityCreated_Register
var s2sdk_OnEntityCreated_Register func(callback s2sdk.OnEntityCreatedCallback)
var S2sdk_OnEntityCreated_Register = &s2sdk_OnEntityCreated_Register

//go:linkname s2sdk_OnEntityCreated_Unregister __package__/s2sdk._OnEntityCreated_Unregister
var s2sdk_OnEntityCreated_Unregister func(callback s2sdk.OnEntityCreatedCallback)
var S2sdk_OnEntityCreated_Unregister = &s2sdk_OnEntityCreated_Unregister

//go:linkname s2sdk_OnEntitySpawned_Register __package__/s2sdk._OnEntitySpawned_Register
var s2sdk_OnEntitySpawned_Register func(callback s2sdk.OnEntitySpawnedCallback)
var S2sdk_OnEntitySpawned_Register = &s2sdk_OnEntitySpawned_Register

//go:linkname s2sdk_OnEntitySpawned_Unregister __package__/s2sdk._OnEntitySpawned_Unregister
var s2sdk_OnEntitySpawned_Unregister func(callback s2sdk.OnEntitySpawnedCallback)
var S2sdk_OnEntitySpawned_Unregister = &s2sdk_OnEntitySpawned_Unregister

//go:linkname s2sdk_OnEntityDeleted_Register __package__/s2sdk._OnEntityDeleted_Register
var s2sdk_OnEntityDeleted_Register func(callback s2sdk.OnEntityDeletedCallback)
var S2sdk_OnEntityDeleted_Register = &s2sdk_OnEntityDeleted_Register

//go:linkname s2sdk_OnEntityDeleted_Unregister __package__/s2sdk._OnEntityDeleted_Unregister
var s2sdk_OnEntityDeleted_Unregister func(callback s2sdk.OnEntityDeletedCallback)
var S2sdk_OnEntityDeleted_Unregister = &s2sdk_OnEntityDeleted_Unregister

//go:linkname s2sdk_OnEntityParentChanged_Register __package__/s2sdk._OnEntityParentChanged_Register
var s2sdk_OnEntityParentChanged_Register func(callback s2sdk.OnEntityParentChangedCallback)
var S2sdk_OnEntityParentChanged_Register = &s2sdk_OnEntityParentChanged_Register

//go:linkname s2sdk_OnEntityParentChanged_Unregister __package__/s2sdk._OnEntityParentChanged_Unregister
var s2sdk_OnEntityParentChanged_Unregister func(callback s2sdk.OnEntityParentChangedCallback)
var S2sdk_OnEntityParentChanged_Unregister = &s2sdk_OnEntityParentChanged_Unregister

//go:linkname s2sdk_OnServerCheckTransmit_Register __package__/s2sdk._OnServerCheckTransmit_Register
var s2sdk_OnServerCheckTransmit_Register func(callback s2sdk.OnServerCheckTransmitCallback)
var S2sdk_OnServerCheckTransmit_Register = &s2sdk_OnServerCheckTransmit_Register

//go:linkname s2sdk_OnServerCheckTransmit_Unregister __package__/s2sdk._OnServerCheckTransmit_Unregister
var s2sdk_OnServerCheckTransmit_Unregister func(callback s2sdk.OnServerCheckTransmitCallback)
var S2sdk_OnServerCheckTransmit_Unregister = &s2sdk_OnServerCheckTransmit_Unregister

//go:linkname s2sdk_OnServerStartup_Register __package__/s2sdk._OnServerStartup_Register
var s2sdk_OnServerStartup_Register func(callback s2sdk.OnServerStartupCallback)
var S2sdk_OnServerStartup_Register = &s2sdk_OnServerStartup_Register

//go:linkname s2sdk_OnServerStartup_Unregister __package__/s2sdk._OnServerStartup_Unregister
var s2sdk_OnServerStartup_Unregister func(callback s2sdk.OnServerStartupCallback)
var S2sdk_OnServerStartup_Unregister = &s2sdk_OnServerStartup_Unregister

//go:linkname s2sdk_OnBuildGameSessionManifest_Register __package__/s2sdk._OnBuildGameSessionManifest_Register
var s2sdk_OnBuildGameSessionManifest_Register func(callback s2sdk.OnBuildGameSessionManifestCallback)
var S2sdk_OnBuildGameSessionManifest_Register = &s2sdk_OnBuildGameSessionManifest_Register

//go:linkname s2sdk_OnBuildGameSessionManifest_Unregister __package__/s2sdk._OnBuildGameSessionManifest_Unregister
var s2sdk_OnBuildGameSessionManifest_Unregister func(callback s2sdk.OnBuildGameSessionManifestCallback)
var S2sdk_OnBuildGameSessionManifest_Unregister = &s2sdk_OnBuildGameSessionManifest_Unregister

//go:linkname s2sdk_OnServerActivate_Register __package__/s2sdk._OnServerActivate_Register
var s2sdk_OnServerActivate_Register func(callback s2sdk.OnServerActivateCallback)
var S2sdk_OnServerActivate_Register = &s2sdk_OnServerActivate_Register

//go:linkname s2sdk_OnServerActivate_Unregister __package__/s2sdk._OnServerActivate_Unregister
var s2sdk_OnServerActivate_Unregister func(callback s2sdk.OnServerActivateCallback)
var S2sdk_OnServerActivate_Unregister = &s2sdk_OnServerActivate_Unregister

//go:linkname s2sdk_OnServerSpawn_Register __package__/s2sdk._OnServerSpawn_Register
var s2sdk_OnServerSpawn_Register func(callback s2sdk.OnServerSpawnCallback)
var S2sdk_OnServerSpawn_Register = &s2sdk_OnServerSpawn_Register

//go:linkname s2sdk_OnServerSpawn_Unregister __package__/s2sdk._OnServerSpawn_Unregister
var s2sdk_OnServerSpawn_Unregister func(callback s2sdk.OnServerSpawnCallback)
var S2sdk_OnServerSpawn_Unregister = &s2sdk_OnServerSpawn_Unregister

//go:linkname s2sdk_OnServerStarted_Register __package__/s2sdk._OnServerStarted_Register
var s2sdk_OnServerStarted_Register func(callback s2sdk.OnServerStartedCallback)
var S2sdk_OnServerStarted_Register = &s2sdk_OnServerStarted_Register

//go:linkname s2sdk_OnServerStarted_Unregister __package__/s2sdk._OnServerStarted_Unregister
var s2sdk_OnServerStarted_Unregister func(callback s2sdk.OnServerStartedCallback)
var S2sdk_OnServerStarted_Unregister = &s2sdk_OnServerStarted_Unregister

//go:linkname s2sdk_OnMapStart_Register __package__/s2sdk._OnMapStart_Register
var s2sdk_OnMapStart_Register func(callback s2sdk.OnMapStartCallback)
var S2sdk_OnMapStart_Register = &s2sdk_OnMapStart_Register

//go:linkname s2sdk_OnMapStart_Unregister __package__/s2sdk._OnMapStart_Unregister
var s2sdk_OnMapStart_Unregister func(callback s2sdk.OnMapStartCallback)
var S2sdk_OnMapStart_Unregister = &s2sdk_OnMapStart_Unregister

//go:linkname s2sdk_OnMapEnd_Register __package__/s2sdk._OnMapEnd_Register
var s2sdk_OnMapEnd_Register func(callback s2sdk.OnMapEndCallback)
var S2sdk_OnMapEnd_Register = &s2sdk_OnMapEnd_Register

//go:linkname s2sdk_OnMapEnd_Unregister __package__/s2sdk._OnMapEnd_Unregister
var s2sdk_OnMapEnd_Unregister func(callback s2sdk.OnMapEndCallback)
var S2sdk_OnMapEnd_Unregister = &s2sdk_OnMapEnd_Unregister

//go:linkname s2sdk_OnGameFrame_Register __package__/s2sdk._OnGameFrame_Register
var s2sdk_OnGameFrame_Register func(callback s2sdk.OnGameFrameCallback)
var S2sdk_OnGameFrame_Register = &s2sdk_OnGameFrame_Register

//go:linkname s2sdk_OnGameFrame_Unregister __package__/s2sdk._OnGameFrame_Unregister
var s2sdk_OnGameFrame_Unregister func(callback s2sdk.OnGameFrameCallback)
var S2sdk_OnGameFrame_Unregister = &s2sdk_OnGameFrame_Unregister

//go:linkname s2sdk_OnUpdateWhenNotInGame_Register __package__/s2sdk._OnUpdateWhenNotInGame_Register
var s2sdk_OnUpdateWhenNotInGame_Register func(callback s2sdk.OnUpdateWhenNotInGameCallback)
var S2sdk_OnUpdateWhenNotInGame_Register = &s2sdk_OnUpdateWhenNotInGame_Register

//go:linkname s2sdk_OnUpdateWhenNotInGame_Unregister __package__/s2sdk._OnUpdateWhenNotInGame_Unregister
var s2sdk_OnUpdateWhenNotInGame_Unregister func(callback s2sdk.OnUpdateWhenNotInGameCallback)
var S2sdk_OnUpdateWhenNotInGame_Unregister = &s2sdk_OnUpdateWhenNotInGame_Unregister

//go:linkname s2sdk_OnPreWorldUpdate_Register __package__/s2sdk._OnPreWorldUpdate_Register
var s2sdk_OnPreWorldUpdate_Register func(callback s2sdk.OnPreWorldUpdateCallback)
var S2sdk_OnPreWorldUpdate_Register = &s2sdk_OnPreWorldUpdate_Register

//go:linkname s2sdk_OnPreWorldUpdate_Unregister __package__/s2sdk._OnPreWorldUpdate_Unregister
var s2sdk_OnPreWorldUpdate_Unregister func(callback s2sdk.OnPreWorldUpdateCallback)
var S2sdk_OnPreWorldUpdate_Unregister = &s2sdk_OnPreWorldUpdate_Unregister

//go:linkname s2sdk_GetGameRulesProxy __package__/s2sdk._GetGameRulesProxy
var s2sdk_GetGameRulesProxy func() uintptr
var S2sdk_GetGameRulesProxy = &s2sdk_GetGameRulesProxy

//go:linkname s2sdk_GetGameRules __package__/s2sdk._GetGameRules
var s2sdk_GetGameRules func() uintptr
var S2sdk_GetGameRules = &s2sdk_GetGameRules

//go:linkname s2sdk_GetGameTeamManager __package__/s2sdk._GetGameTeamManager
var s2sdk_GetGameTeamManager func(team s2sdk.CSTeam) uintptr
var S2sdk_GetGameTeamManager = &s2sdk_GetGameTeamManager

//go:linkname s2sdk_GetGameTeamScore __package__/s2sdk._GetGameTeamScore
var s2sdk_GetGameTeamScore func(team s2sdk.CSTeam) int32
var S2sdk_GetGameTeamScore = &s2sdk_GetGameTeamScore

//go:linkname s2sdk_GetGamePlayerCount __package__/s2sdk._GetGamePlayerCount
var s2sdk_GetGamePlayerCount func(team s2sdk.CSTeam) int32
var S2sdk_GetGamePlayerCount = &s2sdk_GetGamePlayerCount

//go:linkname s2sdk_GetGameTotalRoundsPlayed __package__/s2sdk._GetGameTotalRoundsPlayed
var s2sdk_GetGameTotalRoundsPlayed func() int32
var S2sdk_GetGameTotalRoundsPlayed = &s2sdk_GetGameTotalRoundsPlayed

//go:linkname s2sdk_TerminateRound __package__/s2sdk._TerminateRound
var s2sdk_TerminateRound func(delay float32, reason s2sdk.CSRoundEndReason)
var S2sdk_TerminateRound = &s2sdk_TerminateRound

//go:linkname s2sdk_HookUserMessage __package__/s2sdk._HookUserMessage
var s2sdk_HookUserMessage func(messageId int16, callback s2sdk.UserMessageCallback, mode s2sdk.HookMode) bool
var S2sdk_HookUserMessage = &s2sdk_HookUserMessage

//go:linkname s2sdk_UnhookUserMessage __package__/s2sdk._UnhookUserMessage
var s2sdk_UnhookUserMessage func(messageId int16, callback s2sdk.UserMessageCallback, mode s2sdk.HookMode) bool
var S2sdk_UnhookUserMessage = &s2sdk_UnhookUserMessage

//go:linkname s2sdk_UserMessageCreateFromSerializable __package__/s2sdk._UserMessageCreateFromSerializable
var s2sdk_UserMessageCreateFromSerializable func(msgSerializable uintptr, message uintptr, recipientMask uint64) uintptr
var S2sdk_UserMessageCreateFromSerializable = &s2sdk_UserMessageCreateFromSerializable

//go:linkname s2sdk_UserMessageCreateFromName __package__/s2sdk._UserMessageCreateFromName
var s2sdk_UserMessageCreateFromName func(messageName string) uintptr
var S2sdk_UserMessageCreateFromName = &s2sdk_UserMessageCreateFromName

//go:linkname s2sdk_UserMessageCreateFromId __package__/s2sdk._UserMessageCreateFromId
var s2sdk_UserMessageCreateFromId func(messageId int16) uintptr
var S2sdk_UserMessageCreateFromId = &s2sdk_UserMessageCreateFromId

//go:linkname s2sdk_UserMessageDestroy __package__/s2sdk._UserMessageDestroy
var s2sdk_UserMessageDestroy func(userMessage uintptr)
var S2sdk_UserMessageDestroy = &s2sdk_UserMessageDestroy

//go:linkname s2sdk_UserMessageSend __package__/s2sdk._UserMessageSend
var s2sdk_UserMessageSend func(userMessage uintptr)
var S2sdk_UserMessageSend = &s2sdk_UserMessageSend

//go:linkname s2sdk_UserMessageGetMessageName __package__/s2sdk._UserMessageGetMessageName
var s2sdk_UserMessageGetMessageName func(userMessage uintptr) string
var S2sdk_UserMessageGetMessageName = &s2sdk_UserMessageGetMessageName

//go:linkname s2sdk_UserMessageGetMessageID __package__/s2sdk._UserMessageGetMessageID
var s2sdk_UserMessageGetMessageID func(userMessage uintptr) int16
var S2sdk_UserMessageGetMessageID = &s2sdk_UserMessageGetMessageID

//go:linkname s2sdk_UserMessageHasField __package__/s2sdk._UserMessageHasField
var s2sdk_UserMessageHasField func(userMessage uintptr, fieldName string) bool
var S2sdk_UserMessageHasField = &s2sdk_UserMessageHasField

//go:linkname s2sdk_UserMessageGetProtobufMessage __package__/s2sdk._UserMessageGetProtobufMessage
var s2sdk_UserMessageGetProtobufMessage func(userMessage uintptr) uintptr
var S2sdk_UserMessageGetProtobufMessage = &s2sdk_UserMessageGetProtobufMessage

//go:linkname s2sdk_UserMessageGetSerializableMessage __package__/s2sdk._UserMessageGetSerializableMessage
var s2sdk_UserMessageGetSerializableMessage func(userMessage uintptr) uintptr
var S2sdk_UserMessageGetSerializableMessage = &s2sdk_UserMessageGetSerializableMessage

//go:linkname s2sdk_UserMessageFindMessageIdByName __package__/s2sdk._UserMessageFindMessageIdByName
var s2sdk_UserMessageFindMessageIdByName func(messageName string) int16
var S2sdk_UserMessageFindMessageIdByName = &s2sdk_UserMessageFindMessageIdByName

//go:linkname s2sdk_UserMessageGetRecipientMask __package__/s2sdk._UserMessageGetRecipientMask
var s2sdk_UserMessageGetRecipientMask func(userMessage uintptr) uint64
var S2sdk_UserMessageGetRecipientMask = &s2sdk_UserMessageGetRecipientMask

//go:linkname s2sdk_UserMessageAddRecipient __package__/s2sdk._UserMessageAddRecipient
var s2sdk_UserMessageAddRecipient func(userMessage uintptr, playerSlot int32)
var S2sdk_UserMessageAddRecipient = &s2sdk_UserMessageAddRecipient

//go:linkname s2sdk_UserMessageAddAllPlayers __package__/s2sdk._UserMessageAddAllPlayers
var s2sdk_UserMessageAddAllPlayers func(userMessage uintptr)
var S2sdk_UserMessageAddAllPlayers = &s2sdk_UserMessageAddAllPlayers

//go:linkname s2sdk_UserMessageSetRecipientMask __package__/s2sdk._UserMessageSetRecipientMask
var s2sdk_UserMessageSetRecipientMask func(userMessage uintptr, mask uint64)
var S2sdk_UserMessageSetRecipientMask = &s2sdk_UserMessageSetRecipientMask

//go:linkname s2sdk_UserMessageRemoveAllRecipient __package__/s2sdk._UserMessageRemoveAllRecipient
var s2sdk_UserMessageRemoveAllRecipient func(userMessage uintptr)
var S2sdk_UserMessageRemoveAllRecipient = &s2sdk_UserMessageRemoveAllRecipient

//go:linkname s2sdk_UserMessageGetRepeatedFieldCount __package__/s2sdk._UserMessageGetRepeatedFieldCount
var s2sdk_UserMessageGetRepeatedFieldCount func(userMessage uintptr, fieldName string) int32
var S2sdk_UserMessageGetRepeatedFieldCount = &s2sdk_UserMessageGetRepeatedFieldCount

//go:linkname s2sdk_UserMessageRemoveRepeatedFieldValue __package__/s2sdk._UserMessageRemoveRepeatedFieldValue
var s2sdk_UserMessageRemoveRepeatedFieldValue func(userMessage uintptr, fieldName string, index int32) bool
var S2sdk_UserMessageRemoveRepeatedFieldValue = &s2sdk_UserMessageRemoveRepeatedFieldValue

//go:linkname s2sdk_UserMessageGetDebugString __package__/s2sdk._UserMessageGetDebugString
var s2sdk_UserMessageGetDebugString func(userMessage uintptr) string
var S2sdk_UserMessageGetDebugString = &s2sdk_UserMessageGetDebugString

//go:linkname s2sdk_PbReadEnum __package__/s2sdk._PbReadEnum
var s2sdk_PbReadEnum func(userMessage uintptr, fieldName string, index int32) int32
var S2sdk_PbReadEnum = &s2sdk_PbReadEnum

//go:linkname s2sdk_PbReadInt32 __package__/s2sdk._PbReadInt32
var s2sdk_PbReadInt32 func(userMessage uintptr, fieldName string, index int32) int32
var S2sdk_PbReadInt32 = &s2sdk_PbReadInt32

//go:linkname s2sdk_PbReadInt64 __package__/s2sdk._PbReadInt64
var s2sdk_PbReadInt64 func(userMessage uintptr, fieldName string, index int32) int64
var S2sdk_PbReadInt64 = &s2sdk_PbReadInt64

//go:linkname s2sdk_PbReadUInt32 __package__/s2sdk._PbReadUInt32
var s2sdk_PbReadUInt32 func(userMessage uintptr, fieldName string, index int32) uint32
var S2sdk_PbReadUInt32 = &s2sdk_PbReadUInt32

//go:linkname s2sdk_PbReadUInt64 __package__/s2sdk._PbReadUInt64
var s2sdk_PbReadUInt64 func(userMessage uintptr, fieldName string, index int32) uint64
var S2sdk_PbReadUInt64 = &s2sdk_PbReadUInt64

//go:linkname s2sdk_PbReadFloat __package__/s2sdk._PbReadFloat
var s2sdk_PbReadFloat func(userMessage uintptr, fieldName string, index int32) float32
var S2sdk_PbReadFloat = &s2sdk_PbReadFloat

//go:linkname s2sdk_PbReadDouble __package__/s2sdk._PbReadDouble
var s2sdk_PbReadDouble func(userMessage uintptr, fieldName string, index int32) float64
var S2sdk_PbReadDouble = &s2sdk_PbReadDouble

//go:linkname s2sdk_PbReadBool __package__/s2sdk._PbReadBool
var s2sdk_PbReadBool func(userMessage uintptr, fieldName string, index int32) bool
var S2sdk_PbReadBool = &s2sdk_PbReadBool

//go:linkname s2sdk_PbReadString __package__/s2sdk._PbReadString
var s2sdk_PbReadString func(userMessage uintptr, fieldName string, index int32) string
var S2sdk_PbReadString = &s2sdk_PbReadString

//go:linkname s2sdk_PbReadColor __package__/s2sdk._PbReadColor
var s2sdk_PbReadColor func(userMessage uintptr, fieldName string, index int32) plugify.Vector4
var S2sdk_PbReadColor = &s2sdk_PbReadColor

//go:linkname s2sdk_PbReadVector2 __package__/s2sdk._PbReadVector2
var s2sdk_PbReadVector2 func(userMessage uintptr, fieldName string, index int32) plugify.Vector2
var S2sdk_PbReadVector2 = &s2sdk_PbReadVector2

//go:linkname s2sdk_PbReadVector3 __package__/s2sdk._PbReadVector3
var s2sdk_PbReadVector3 func(userMessage uintptr, fieldName string, index int32) plugify.Vector3
var S2sdk_PbReadVector3 = &s2sdk_PbReadVector3

//go:linkname s2sdk_PbReadVector4 __package__/s2sdk._PbReadVector4
var s2sdk_PbReadVector4 func(userMessage uintptr, fieldName string, index int32) plugify.Vector4
var S2sdk_PbReadVector4 = &s2sdk_PbReadVector4

//go:linkname s2sdk_PbReadQAngle __package__/s2sdk._PbReadQAngle
var s2sdk_PbReadQAngle func(userMessage uintptr, fieldName string, index int32) plugify.Vector3
var S2sdk_PbReadQAngle = &s2sdk_PbReadQAngle

//go:linkname s2sdk_PbReadMessage __package__/s2sdk._PbReadMessage
var s2sdk_PbReadMessage func(userMessage uintptr, fieldName string, index int32) uintptr
var S2sdk_PbReadMessage = &s2sdk_PbReadMessage

//go:linkname s2sdk_PbGetEnum __package__/s2sdk._PbGetEnum
var s2sdk_PbGetEnum func(userMessage uintptr, fieldName string, out *int32) bool
var S2sdk_PbGetEnum = &s2sdk_PbGetEnum

//go:linkname s2sdk_PbSetEnum __package__/s2sdk._PbSetEnum
var s2sdk_PbSetEnum func(userMessage uintptr, fieldName string, value int32) bool
var S2sdk_PbSetEnum = &s2sdk_PbSetEnum

//go:linkname s2sdk_PbGetInt32 __package__/s2sdk._PbGetInt32
var s2sdk_PbGetInt32 func(userMessage uintptr, fieldName string, out *int32) bool
var S2sdk_PbGetInt32 = &s2sdk_PbGetInt32

//go:linkname s2sdk_PbSetInt32 __package__/s2sdk._PbSetInt32
var s2sdk_PbSetInt32 func(userMessage uintptr, fieldName string, value int32) bool
var S2sdk_PbSetInt32 = &s2sdk_PbSetInt32

//go:linkname s2sdk_PbGetInt64 __package__/s2sdk._PbGetInt64
var s2sdk_PbGetInt64 func(userMessage uintptr, fieldName string, out *int64) bool
var S2sdk_PbGetInt64 = &s2sdk_PbGetInt64

//go:linkname s2sdk_PbSetInt64 __package__/s2sdk._PbSetInt64
var s2sdk_PbSetInt64 func(userMessage uintptr, fieldName string, value int64) bool
var S2sdk_PbSetInt64 = &s2sdk_PbSetInt64

//go:linkname s2sdk_PbGetUInt32 __package__/s2sdk._PbGetUInt32
var s2sdk_PbGetUInt32 func(userMessage uintptr, fieldName string, out *uint32) bool
var S2sdk_PbGetUInt32 = &s2sdk_PbGetUInt32

//go:linkname s2sdk_PbSetUInt32 __package__/s2sdk._PbSetUInt32
var s2sdk_PbSetUInt32 func(userMessage uintptr, fieldName string, value uint32) bool
var S2sdk_PbSetUInt32 = &s2sdk_PbSetUInt32

//go:linkname s2sdk_PbGetUInt64 __package__/s2sdk._PbGetUInt64
var s2sdk_PbGetUInt64 func(userMessage uintptr, fieldName string, out *uint64) bool
var S2sdk_PbGetUInt64 = &s2sdk_PbGetUInt64

//go:linkname s2sdk_PbSetUInt64 __package__/s2sdk._PbSetUInt64
var s2sdk_PbSetUInt64 func(userMessage uintptr, fieldName string, value uint64) bool
var S2sdk_PbSetUInt64 = &s2sdk_PbSetUInt64

//go:linkname s2sdk_PbGetBool __package__/s2sdk._PbGetBool
var s2sdk_PbGetBool func(userMessage uintptr, fieldName string, out *bool) bool
var S2sdk_PbGetBool = &s2sdk_PbGetBool

//go:linkname s2sdk_PbSetBool __package__/s2sdk._PbSetBool
var s2sdk_PbSetBool func(userMessage uintptr, fieldName string, value bool) bool
var S2sdk_PbSetBool = &s2sdk_PbSetBool

//go:linkname s2sdk_PbGetFloat __package__/s2sdk._PbGetFloat
var s2sdk_PbGetFloat func(userMessage uintptr, fieldName string, out *float32) bool
var S2sdk_PbGetFloat = &s2sdk_PbGetFloat

//go:linkname s2sdk_PbSetFloat __package__/s2sdk._PbSetFloat
var s2sdk_PbSetFloat func(userMessage uintptr, fieldName string, value float32) bool
var S2sdk_PbSetFloat = &s2sdk_PbSetFloat

//go:linkname s2sdk_PbGetDouble __package__/s2sdk._PbGetDouble
var s2sdk_PbGetDouble func(userMessage uintptr, fieldName string, out *float64) bool
var S2sdk_PbGetDouble = &s2sdk_PbGetDouble

//go:linkname s2sdk_PbSetDouble __package__/s2sdk._PbSetDouble
var s2sdk_PbSetDouble func(userMessage uintptr, fieldName string, value float64) bool
var S2sdk_PbSetDouble = &s2sdk_PbSetDouble

//go:linkname s2sdk_PbGetString __package__/s2sdk._PbGetString
var s2sdk_PbGetString func(userMessage uintptr, fieldName string, out *string) bool
var S2sdk_PbGetString = &s2sdk_PbGetString

//go:linkname s2sdk_PbSetString __package__/s2sdk._PbSetString
var s2sdk_PbSetString func(userMessage uintptr, fieldName string, value string) bool
var S2sdk_PbSetString = &s2sdk_PbSetString

//go:linkname s2sdk_PbGetColor __package__/s2sdk._PbGetColor
var s2sdk_PbGetColor func(userMessage uintptr, fieldName string, out *plugify.Vector4) bool
var S2sdk_PbGetColor = &s2sdk_PbGetColor

//go:linkname s2sdk_PbSetColor __package__/s2sdk._PbSetColor
var s2sdk_PbSetColor func(userMessage uintptr, fieldName string, value plugify.Vector4) bool
var S2sdk_PbSetColor = &s2sdk_PbSetColor

//go:linkname s2sdk_PbGetVector2 __package__/s2sdk._PbGetVector2
var s2sdk_PbGetVector2 func(userMessage uintptr, fieldName string, out *plugify.Vector2) bool
var S2sdk_PbGetVector2 = &s2sdk_PbGetVector2

//go:linkname s2sdk_PbSetVector2 __package__/s2sdk._PbSetVector2
var s2sdk_PbSetVector2 func(userMessage uintptr, fieldName string, value plugify.Vector2) bool
var S2sdk_PbSetVector2 = &s2sdk_PbSetVector2

//go:linkname s2sdk_PbGetVector3 __package__/s2sdk._PbGetVector3
var s2sdk_PbGetVector3 func(userMessage uintptr, fieldName string, out *plugify.Vector3) bool
var S2sdk_PbGetVector3 = &s2sdk_PbGetVector3

//go:linkname s2sdk_PbSetVector3 __package__/s2sdk._PbSetVector3
var s2sdk_PbSetVector3 func(userMessage uintptr, fieldName string, value plugify.Vector3) bool
var S2sdk_PbSetVector3 = &s2sdk_PbSetVector3

//go:linkname s2sdk_PbGetVector4 __package__/s2sdk._PbGetVector4
var s2sdk_PbGetVector4 func(userMessage uintptr, fieldName string, out *plugify.Vector4) bool
var S2sdk_PbGetVector4 = &s2sdk_PbGetVector4

//go:linkname s2sdk_PbSetVector4 __package__/s2sdk._PbSetVector4
var s2sdk_PbSetVector4 func(userMessage uintptr, fieldName string, value plugify.Vector4) bool
var S2sdk_PbSetVector4 = &s2sdk_PbSetVector4

//go:linkname s2sdk_PbGetQAngle __package__/s2sdk._PbGetQAngle
var s2sdk_PbGetQAngle func(userMessage uintptr, fieldName string, out *plugify.Vector3) bool
var S2sdk_PbGetQAngle = &s2sdk_PbGetQAngle

//go:linkname s2sdk_PbSetQAngle __package__/s2sdk._PbSetQAngle
var s2sdk_PbSetQAngle func(userMessage uintptr, fieldName string, value plugify.Vector3) bool
var S2sdk_PbSetQAngle = &s2sdk_PbSetQAngle

//go:linkname s2sdk_PbGetMessage __package__/s2sdk._PbGetMessage
var s2sdk_PbGetMessage func(userMessage uintptr, fieldName string, out *uintptr) bool
var S2sdk_PbGetMessage = &s2sdk_PbGetMessage

//go:linkname s2sdk_PbSetMessage __package__/s2sdk._PbSetMessage
var s2sdk_PbSetMessage func(userMessage uintptr, fieldName string, value uintptr) bool
var S2sdk_PbSetMessage = &s2sdk_PbSetMessage

//go:linkname s2sdk_PbGetRepeatedEnum __package__/s2sdk._PbGetRepeatedEnum
var s2sdk_PbGetRepeatedEnum func(userMessage uintptr, fieldName string, index int32, out *int32) bool
var S2sdk_PbGetRepeatedEnum = &s2sdk_PbGetRepeatedEnum

//go:linkname s2sdk_PbSetRepeatedEnum __package__/s2sdk._PbSetRepeatedEnum
var s2sdk_PbSetRepeatedEnum func(userMessage uintptr, fieldName string, index int32, value int32) bool
var S2sdk_PbSetRepeatedEnum = &s2sdk_PbSetRepeatedEnum

//go:linkname s2sdk_PbAddEnum __package__/s2sdk._PbAddEnum
var s2sdk_PbAddEnum func(userMessage uintptr, fieldName string, value int32) bool
var S2sdk_PbAddEnum = &s2sdk_PbAddEnum

//go:linkname s2sdk_PbGetRepeatedInt32 __package__/s2sdk._PbGetRepeatedInt32
var s2sdk_PbGetRepeatedInt32 func(userMessage uintptr, fieldName string, index int32, out *int32) bool
var S2sdk_PbGetRepeatedInt32 = &s2sdk_PbGetRepeatedInt32

//go:linkname s2sdk_PbSetRepeatedInt32 __package__/s2sdk._PbSetRepeatedInt32
var s2sdk_PbSetRepeatedInt32 func(userMessage uintptr, fieldName string, index int32, value int32) bool
var S2sdk_PbSetRepeatedInt32 = &s2sdk_PbSetRepeatedInt32

//go:linkname s2sdk_PbAddInt32 __package__/s2sdk._PbAddInt32
var s2sdk_PbAddInt32 func(userMessage uintptr, fieldName string, value int32) bool
var S2sdk_PbAddInt32 = &s2sdk_PbAddInt32

//go:linkname s2sdk_PbGetRepeatedInt64 __package__/s2sdk._PbGetRepeatedInt64
var s2sdk_PbGetRepeatedInt64 func(userMessage uintptr, fieldName string, index int32, out *int64) bool
var S2sdk_PbGetRepeatedInt64 = &s2sdk_PbGetRepeatedInt64

//go:linkname s2sdk_PbSetRepeatedInt64 __package__/s2sdk._PbSetRepeatedInt64
var s2sdk_PbSetRepeatedInt64 func(userMessage uintptr, fieldName string, index int32, value int64) bool
var S2sdk_PbSetRepeatedInt64 = &s2sdk_PbSetRepeatedInt64

//go:linkname s2sdk_PbAddInt64 __package__/s2sdk._PbAddInt64
var s2sdk_PbAddInt64 func(userMessage uintptr, fieldName string, value int64) bool
var S2sdk_PbAddInt64 = &s2sdk_PbAddInt64

//go:linkname s2sdk_PbGetRepeatedUInt32 __package__/s2sdk._PbGetRepeatedUInt32
var s2sdk_PbGetRepeatedUInt32 func(userMessage uintptr, fieldName string, index int32, out *uint32) bool
var S2sdk_PbGetRepeatedUInt32 = &s2sdk_PbGetRepeatedUInt32

//go:linkname s2sdk_PbSetRepeatedUInt32 __package__/s2sdk._PbSetRepeatedUInt32
var s2sdk_PbSetRepeatedUInt32 func(userMessage uintptr, fieldName string, index int32, value uint32) bool
var S2sdk_PbSetRepeatedUInt32 = &s2sdk_PbSetRepeatedUInt32

//go:linkname s2sdk_PbAddUInt32 __package__/s2sdk._PbAddUInt32
var s2sdk_PbAddUInt32 func(userMessage uintptr, fieldName string, value uint32) bool
var S2sdk_PbAddUInt32 = &s2sdk_PbAddUInt32

//go:linkname s2sdk_PbGetRepeatedUInt64 __package__/s2sdk._PbGetRepeatedUInt64
var s2sdk_PbGetRepeatedUInt64 func(userMessage uintptr, fieldName string, index int32, out *uint64) bool
var S2sdk_PbGetRepeatedUInt64 = &s2sdk_PbGetRepeatedUInt64

//go:linkname s2sdk_PbSetRepeatedUInt64 __package__/s2sdk._PbSetRepeatedUInt64
var s2sdk_PbSetRepeatedUInt64 func(userMessage uintptr, fieldName string, index int32, value uint64) bool
var S2sdk_PbSetRepeatedUInt64 = &s2sdk_PbSetRepeatedUInt64

//go:linkname s2sdk_PbAddUInt64 __package__/s2sdk._PbAddUInt64
var s2sdk_PbAddUInt64 func(userMessage uintptr, fieldName string, value uint64) bool
var S2sdk_PbAddUInt64 = &s2sdk_PbAddUInt64

//go:linkname s2sdk_PbGetRepeatedBool __package__/s2sdk._PbGetRepeatedBool
var s2sdk_PbGetRepeatedBool func(userMessage uintptr, fieldName string, index int32, out *bool) bool
var S2sdk_PbGetRepeatedBool = &s2sdk_PbGetRepeatedBool

//go:linkname s2sdk_PbSetRepeatedBool __package__/s2sdk._PbSetRepeatedBool
var s2sdk_PbSetRepeatedBool func(userMessage uintptr, fieldName string, index int32, value bool) bool
var S2sdk_PbSetRepeatedBool = &s2sdk_PbSetRepeatedBool

//go:linkname s2sdk_PbAddBool __package__/s2sdk._PbAddBool
var s2sdk_PbAddBool func(userMessage uintptr, fieldName string, value bool) bool
var S2sdk_PbAddBool = &s2sdk_PbAddBool

//go:linkname s2sdk_PbGetRepeatedFloat __package__/s2sdk._PbGetRepeatedFloat
var s2sdk_PbGetRepeatedFloat func(userMessage uintptr, fieldName string, index int32, out *float32) bool
var S2sdk_PbGetRepeatedFloat = &s2sdk_PbGetRepeatedFloat

//go:linkname s2sdk_PbSetRepeatedFloat __package__/s2sdk._PbSetRepeatedFloat
var s2sdk_PbSetRepeatedFloat func(userMessage uintptr, fieldName string, index int32, value float32) bool
var S2sdk_PbSetRepeatedFloat = &s2sdk_PbSetRepeatedFloat

//go:linkname s2sdk_PbAddFloat __package__/s2sdk._PbAddFloat
var s2sdk_PbAddFloat func(userMessage uintptr, fieldName string, value float32) bool
var S2sdk_PbAddFloat = &s2sdk_PbAddFloat

//go:linkname s2sdk_PbGetRepeatedDouble __package__/s2sdk._PbGetRepeatedDouble
var s2sdk_PbGetRepeatedDouble func(userMessage uintptr, fieldName string, index int32, out *float64) bool
var S2sdk_PbGetRepeatedDouble = &s2sdk_PbGetRepeatedDouble

//go:linkname s2sdk_PbSetRepeatedDouble __package__/s2sdk._PbSetRepeatedDouble
var s2sdk_PbSetRepeatedDouble func(userMessage uintptr, fieldName string, index int32, value float64) bool
var S2sdk_PbSetRepeatedDouble = &s2sdk_PbSetRepeatedDouble

//go:linkname s2sdk_PbAddDouble __package__/s2sdk._PbAddDouble
var s2sdk_PbAddDouble func(userMessage uintptr, fieldName string, value float64) bool
var S2sdk_PbAddDouble = &s2sdk_PbAddDouble

//go:linkname s2sdk_PbGetRepeatedString __package__/s2sdk._PbGetRepeatedString
var s2sdk_PbGetRepeatedString func(userMessage uintptr, fieldName string, index int32, out *string) bool
var S2sdk_PbGetRepeatedString = &s2sdk_PbGetRepeatedString

//go:linkname s2sdk_PbSetRepeatedString __package__/s2sdk._PbSetRepeatedString
var s2sdk_PbSetRepeatedString func(userMessage uintptr, fieldName string, index int32, value string) bool
var S2sdk_PbSetRepeatedString = &s2sdk_PbSetRepeatedString

//go:linkname s2sdk_PbAddString __package__/s2sdk._PbAddString
var s2sdk_PbAddString func(userMessage uintptr, fieldName string, value string) bool
var S2sdk_PbAddString = &s2sdk_PbAddString

//go:linkname s2sdk_PbGetRepeatedColor __package__/s2sdk._PbGetRepeatedColor
var s2sdk_PbGetRepeatedColor func(userMessage uintptr, fieldName string, index int32, out *plugify.Vector4) bool
var S2sdk_PbGetRepeatedColor = &s2sdk_PbGetRepeatedColor

//go:linkname s2sdk_PbSetRepeatedColor __package__/s2sdk._PbSetRepeatedColor
var s2sdk_PbSetRepeatedColor func(userMessage uintptr, fieldName string, index int32, value plugify.Vector4) bool
var S2sdk_PbSetRepeatedColor = &s2sdk_PbSetRepeatedColor

//go:linkname s2sdk_PbAddColor __package__/s2sdk._PbAddColor
var s2sdk_PbAddColor func(userMessage uintptr, fieldName string, value plugify.Vector4) bool
var S2sdk_PbAddColor = &s2sdk_PbAddColor

//go:linkname s2sdk_PbGetRepeatedVector2 __package__/s2sdk._PbGetRepeatedVector2
var s2sdk_PbGetRepeatedVector2 func(userMessage uintptr, fieldName string, index int32, out *plugify.Vector2) bool
var S2sdk_PbGetRepeatedVector2 = &s2sdk_PbGetRepeatedVector2

//go:linkname s2sdk_PbSetRepeatedVector2 __package__/s2sdk._PbSetRepeatedVector2
var s2sdk_PbSetRepeatedVector2 func(userMessage uintptr, fieldName string, index int32, value plugify.Vector2) bool
var S2sdk_PbSetRepeatedVector2 = &s2sdk_PbSetRepeatedVector2

//go:linkname s2sdk_PbAddVector2 __package__/s2sdk._PbAddVector2
var s2sdk_PbAddVector2 func(userMessage uintptr, fieldName string, value plugify.Vector2) bool
var S2sdk_PbAddVector2 = &s2sdk_PbAddVector2

//go:linkname s2sdk_PbGetRepeatedVector3 __package__/s2sdk._PbGetRepeatedVector3
var s2sdk_PbGetRepeatedVector3 func(userMessage uintptr, fieldName string, index int32, out *plugify.Vector3) bool
var S2sdk_PbGetRepeatedVector3 = &s2sdk_PbGetRepeatedVector3

//go:linkname s2sdk_PbSetRepeatedVector3 __package__/s2sdk._PbSetRepeatedVector3
var s2sdk_PbSetRepeatedVector3 func(userMessage uintptr, fieldName string, index int32, value plugify.Vector3) bool
var S2sdk_PbSetRepeatedVector3 = &s2sdk_PbSetRepeatedVector3

//go:linkname s2sdk_PbAddVector3 __package__/s2sdk._PbAddVector3
var s2sdk_PbAddVector3 func(userMessage uintptr, fieldName string, value plugify.Vector3) bool
var S2sdk_PbAddVector3 = &s2sdk_PbAddVector3

//go:linkname s2sdk_PbGetRepeatedVector4 __package__/s2sdk._PbGetRepeatedVector4
var s2sdk_PbGetRepeatedVector4 func(userMessage uintptr, fieldName string, index int32, out *plugify.Vector4) bool
var S2sdk_PbGetRepeatedVector4 = &s2sdk_PbGetRepeatedVector4

//go:linkname s2sdk_PbSetRepeatedVector4 __package__/s2sdk._PbSetRepeatedVector4
var s2sdk_PbSetRepeatedVector4 func(userMessage uintptr, fieldName string, index int32, value plugify.Vector4) bool
var S2sdk_PbSetRepeatedVector4 = &s2sdk_PbSetRepeatedVector4

//go:linkname s2sdk_PbAddVector4 __package__/s2sdk._PbAddVector4
var s2sdk_PbAddVector4 func(userMessage uintptr, fieldName string, value plugify.Vector4) bool
var S2sdk_PbAddVector4 = &s2sdk_PbAddVector4

//go:linkname s2sdk_PbGetRepeatedQAngle __package__/s2sdk._PbGetRepeatedQAngle
var s2sdk_PbGetRepeatedQAngle func(userMessage uintptr, fieldName string, index int32, out *plugify.Vector3) bool
var S2sdk_PbGetRepeatedQAngle = &s2sdk_PbGetRepeatedQAngle

//go:linkname s2sdk_PbSetRepeatedQAngle __package__/s2sdk._PbSetRepeatedQAngle
var s2sdk_PbSetRepeatedQAngle func(userMessage uintptr, fieldName string, index int32, value plugify.Vector3) bool
var S2sdk_PbSetRepeatedQAngle = &s2sdk_PbSetRepeatedQAngle

//go:linkname s2sdk_PbAddQAngle __package__/s2sdk._PbAddQAngle
var s2sdk_PbAddQAngle func(userMessage uintptr, fieldName string, value plugify.Vector3) bool
var S2sdk_PbAddQAngle = &s2sdk_PbAddQAngle

//go:linkname s2sdk_PbGetRepeatedMessage __package__/s2sdk._PbGetRepeatedMessage
var s2sdk_PbGetRepeatedMessage func(userMessage uintptr, fieldName string, index int32, out *uintptr) bool
var S2sdk_PbGetRepeatedMessage = &s2sdk_PbGetRepeatedMessage

//go:linkname s2sdk_PbSetRepeatedMessage __package__/s2sdk._PbSetRepeatedMessage
var s2sdk_PbSetRepeatedMessage func(userMessage uintptr, fieldName string, index int32, value uintptr) bool
var S2sdk_PbSetRepeatedMessage = &s2sdk_PbSetRepeatedMessage

//go:linkname s2sdk_PbAddMessage __package__/s2sdk._PbAddMessage
var s2sdk_PbAddMessage func(userMessage uintptr, fieldName string, value uintptr) bool
var S2sdk_PbAddMessage = &s2sdk_PbAddMessage

//go:linkname s2sdk_GetWeaponVDataFromKey __package__/s2sdk._GetWeaponVDataFromKey
var s2sdk_GetWeaponVDataFromKey func(name string) uintptr
var S2sdk_GetWeaponVDataFromKey = &s2sdk_GetWeaponVDataFromKey

//go:linkname s2sdk_GetWeaponVData __package__/s2sdk._GetWeaponVData
var s2sdk_GetWeaponVData func(entityHandle int32) uintptr
var S2sdk_GetWeaponVData = &s2sdk_GetWeaponVData

//go:linkname s2sdk_GetWeaponType __package__/s2sdk._GetWeaponType
var s2sdk_GetWeaponType func(entityHandle int32) s2sdk.CSWeaponType
var S2sdk_GetWeaponType = &s2sdk_GetWeaponType

//go:linkname s2sdk_GetWeaponCategory __package__/s2sdk._GetWeaponCategory
var s2sdk_GetWeaponCategory func(entityHandle int32) s2sdk.CSWeaponCategory
var S2sdk_GetWeaponCategory = &s2sdk_GetWeaponCategory

//go:linkname s2sdk_GetWeaponGearSlot __package__/s2sdk._GetWeaponGearSlot
var s2sdk_GetWeaponGearSlot func(entityHandle int32) s2sdk.GearSlot
var S2sdk_GetWeaponGearSlot = &s2sdk_GetWeaponGearSlot

//go:linkname s2sdk_GetWeaponItemDefinition __package__/s2sdk._GetWeaponItemDefinition
var s2sdk_GetWeaponItemDefinition func(entityHandle int32) s2sdk.WeaponDefIndex
var S2sdk_GetWeaponItemDefinition = &s2sdk_GetWeaponItemDefinition

//go:linkname s2sdk_GetWeaponItemDefinitionByName __package__/s2sdk._GetWeaponItemDefinitionByName
var s2sdk_GetWeaponItemDefinitionByName func(itemName string) s2sdk.WeaponDefIndex
var S2sdk_GetWeaponItemDefinitionByName = &s2sdk_GetWeaponItemDefinitionByName

/*func init() {
	s2sdk.ModuleName = "__package__"
}*/
