//go:build plugin
// +build plugin
package plconfigs

//TODO: replace "__package__" by your package name
import (
	"unsafe"
	"github.com/untrustedmodders/go-plugify"
)

var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

//go:linkname configs_Read __package__/configs._Read
var configs_Read func(path string) uintptr
var Configs_Read = &configs_Read

//go:linkname configs_ReadMultiple __package__/configs._ReadMultiple
var configs_ReadMultiple func(paths []string) uintptr
var Configs_ReadMultiple = &configs_ReadMultiple

//go:linkname configs_Make __package__/configs._Make
var configs_Make func() uintptr
var Configs_Make = &configs_Make

//go:linkname configs_Delete __package__/configs._Delete
var configs_Delete func(config uintptr)
var Configs_Delete = &configs_Delete

//go:linkname configs_SetError __package__/configs._SetError
var configs_SetError func(error_ string)
var Configs_SetError = &configs_SetError

//go:linkname configs_GetError __package__/configs._GetError
var configs_GetError func() string
var Configs_GetError = &configs_GetError

//go:linkname configs_Merge __package__/configs._Merge
var configs_Merge func(config uintptr, other uintptr)
var Configs_Merge = &configs_Merge

//go:linkname configs_MergeMove __package__/configs._MergeMove
var configs_MergeMove func(config uintptr, other uintptr)
var Configs_MergeMove = &configs_MergeMove

//go:linkname configs_GetType __package__/configs._GetType
var configs_GetType func(config uintptr) int32
var Configs_GetType = &configs_GetType

//go:linkname configs_IsNull __package__/configs._IsNull
var configs_IsNull func(config uintptr) bool
var Configs_IsNull = &configs_IsNull

//go:linkname configs_IsBool __package__/configs._IsBool
var configs_IsBool func(config uintptr) bool
var Configs_IsBool = &configs_IsBool

//go:linkname configs_IsInt __package__/configs._IsInt
var configs_IsInt func(config uintptr) bool
var Configs_IsInt = &configs_IsInt

//go:linkname configs_IsFloat __package__/configs._IsFloat
var configs_IsFloat func(config uintptr) bool
var Configs_IsFloat = &configs_IsFloat

//go:linkname configs_IsString __package__/configs._IsString
var configs_IsString func(config uintptr) bool
var Configs_IsString = &configs_IsString

//go:linkname configs_IsObject __package__/configs._IsObject
var configs_IsObject func(config uintptr) bool
var Configs_IsObject = &configs_IsObject

//go:linkname configs_IsArray __package__/configs._IsArray
var configs_IsArray func(config uintptr) bool
var Configs_IsArray = &configs_IsArray

//go:linkname configs_SetNull __package__/configs._SetNull
var configs_SetNull func(config uintptr, key string)
var Configs_SetNull = &configs_SetNull

//go:linkname configs_SetObject __package__/configs._SetObject
var configs_SetObject func(config uintptr, key string)
var Configs_SetObject = &configs_SetObject

//go:linkname configs_SetArray __package__/configs._SetArray
var configs_SetArray func(config uintptr, key string)
var Configs_SetArray = &configs_SetArray

//go:linkname configs_SetBool __package__/configs._SetBool
var configs_SetBool func(config uintptr, key string, value bool)
var Configs_SetBool = &configs_SetBool

//go:linkname configs_SetInt __package__/configs._SetInt
var configs_SetInt func(config uintptr, key string, value int64)
var Configs_SetInt = &configs_SetInt

//go:linkname configs_SetFloat __package__/configs._SetFloat
var configs_SetFloat func(config uintptr, key string, value float64)
var Configs_SetFloat = &configs_SetFloat

//go:linkname configs_SetString __package__/configs._SetString
var configs_SetString func(config uintptr, key string, value string)
var Configs_SetString = &configs_SetString

//go:linkname configs_TrySetFromBool __package__/configs._TrySetFromBool
var configs_TrySetFromBool func(config uintptr, key string, value bool) bool
var Configs_TrySetFromBool = &configs_TrySetFromBool

//go:linkname configs_TrySetFromInt __package__/configs._TrySetFromInt
var configs_TrySetFromInt func(config uintptr, key string, value int64) bool
var Configs_TrySetFromInt = &configs_TrySetFromInt

//go:linkname configs_TrySetFromFloat __package__/configs._TrySetFromFloat
var configs_TrySetFromFloat func(config uintptr, key string, value float64) bool
var Configs_TrySetFromFloat = &configs_TrySetFromFloat

//go:linkname configs_TrySetFromString __package__/configs._TrySetFromString
var configs_TrySetFromString func(config uintptr, key string, value string) bool
var Configs_TrySetFromString = &configs_TrySetFromString

//go:linkname configs_PushNull __package__/configs._PushNull
var configs_PushNull func(config uintptr)
var Configs_PushNull = &configs_PushNull

//go:linkname configs_PushBool __package__/configs._PushBool
var configs_PushBool func(config uintptr, value bool)
var Configs_PushBool = &configs_PushBool

//go:linkname configs_PushInt __package__/configs._PushInt
var configs_PushInt func(config uintptr, value int64)
var Configs_PushInt = &configs_PushInt

//go:linkname configs_PushFloat __package__/configs._PushFloat
var configs_PushFloat func(config uintptr, value float64)
var Configs_PushFloat = &configs_PushFloat

//go:linkname configs_PushString __package__/configs._PushString
var configs_PushString func(config uintptr, value string)
var Configs_PushString = &configs_PushString

//go:linkname configs_PushObject __package__/configs._PushObject
var configs_PushObject func(config uintptr)
var Configs_PushObject = &configs_PushObject

//go:linkname configs_PushArray __package__/configs._PushArray
var configs_PushArray func(config uintptr)
var Configs_PushArray = &configs_PushArray

//go:linkname configs_GetBool __package__/configs._GetBool
var configs_GetBool func(config uintptr, key string, defaultValue bool) bool
var Configs_GetBool = &configs_GetBool

//go:linkname configs_GetInt __package__/configs._GetInt
var configs_GetInt func(config uintptr, key string, defaultValue int64) int64
var Configs_GetInt = &configs_GetInt

//go:linkname configs_GetFloat __package__/configs._GetFloat
var configs_GetFloat func(config uintptr, key string, defaultValue float64) float64
var Configs_GetFloat = &configs_GetFloat

//go:linkname configs_GetString __package__/configs._GetString
var configs_GetString func(config uintptr, key string, defaultValue string) string
var Configs_GetString = &configs_GetString

//go:linkname configs_GetAsBool __package__/configs._GetAsBool
var configs_GetAsBool func(config uintptr, key string, defaultValue bool) bool
var Configs_GetAsBool = &configs_GetAsBool

//go:linkname configs_GetAsInt __package__/configs._GetAsInt
var configs_GetAsInt func(config uintptr, key string, defaultValue int64) int64
var Configs_GetAsInt = &configs_GetAsInt

//go:linkname configs_GetAsFloat __package__/configs._GetAsFloat
var configs_GetAsFloat func(config uintptr, key string, defaultValue float64) float64
var Configs_GetAsFloat = &configs_GetAsFloat

//go:linkname configs_GetAsString __package__/configs._GetAsString
var configs_GetAsString func(config uintptr, key string, defaultValue string) string
var Configs_GetAsString = &configs_GetAsString

//go:linkname configs_HasKey __package__/configs._HasKey
var configs_HasKey func(config uintptr, key string) bool
var Configs_HasKey = &configs_HasKey

//go:linkname configs_Empty __package__/configs._Empty
var configs_Empty func(config uintptr) bool
var Configs_Empty = &configs_Empty

//go:linkname configs_GetSize __package__/configs._GetSize
var configs_GetSize func(config uintptr) int64
var Configs_GetSize = &configs_GetSize

//go:linkname configs_GetName __package__/configs._GetName
var configs_GetName func(config uintptr) string
var Configs_GetName = &configs_GetName

//go:linkname configs_GetPath __package__/configs._GetPath
var configs_GetPath func(config uintptr) string
var Configs_GetPath = &configs_GetPath

//go:linkname configs_JumpFirst __package__/configs._JumpFirst
var configs_JumpFirst func(config uintptr) bool
var Configs_JumpFirst = &configs_JumpFirst

//go:linkname configs_JumpLast __package__/configs._JumpLast
var configs_JumpLast func(config uintptr) bool
var Configs_JumpLast = &configs_JumpLast

//go:linkname configs_JumpNext __package__/configs._JumpNext
var configs_JumpNext func(config uintptr) bool
var Configs_JumpNext = &configs_JumpNext

//go:linkname configs_JumpPrev __package__/configs._JumpPrev
var configs_JumpPrev func(config uintptr) bool
var Configs_JumpPrev = &configs_JumpPrev

//go:linkname configs_JumpKey __package__/configs._JumpKey
var configs_JumpKey func(config uintptr, key string, create bool) bool
var Configs_JumpKey = &configs_JumpKey

//go:linkname configs_JumpN __package__/configs._JumpN
var configs_JumpN func(config uintptr, n int32) bool
var Configs_JumpN = &configs_JumpN

//go:linkname configs_JumpBack __package__/configs._JumpBack
var configs_JumpBack func(config uintptr) bool
var Configs_JumpBack = &configs_JumpBack

//go:linkname configs_JumpRoot __package__/configs._JumpRoot
var configs_JumpRoot func(config uintptr)
var Configs_JumpRoot = &configs_JumpRoot

//go:linkname configs_Remove __package__/configs._Remove
var configs_Remove func(config uintptr) int32
var Configs_Remove = &configs_Remove

//go:linkname configs_RemoveKey __package__/configs._RemoveKey
var configs_RemoveKey func(config uintptr, key string) bool
var Configs_RemoveKey = &configs_RemoveKey

//go:linkname configs_Clear __package__/configs._Clear
var configs_Clear func(config uintptr)
var Configs_Clear = &configs_Clear

//go:linkname configs_NodeToJsonString __package__/configs._NodeToJsonString
var configs_NodeToJsonString func(config uintptr) string
var Configs_NodeToJsonString = &configs_NodeToJsonString

//go:linkname configs_RootToJsonString __package__/configs._RootToJsonString
var configs_RootToJsonString func(config uintptr) string
var Configs_RootToJsonString = &configs_RootToJsonString

/*func init() {
	configs.ModuleName = "__package__"
}*/
