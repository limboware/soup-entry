package plconfigs

/*
#include "typechecking.h"
#cgo noescape GetType
#cgo noescape IsNull
#cgo noescape IsBool
#cgo noescape IsInt
#cgo noescape IsFloat
#cgo noescape IsString
#cgo noescape IsObject
#cgo noescape IsArray
*/
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
	"github.com/untrustedmodders/go-plugify"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

// Generated from configs (group: typechecking)

var _GetType = func(config uintptr) int32 {
	var __retVal int32
	__config := C.uintptr_t(config)
	__retVal = int32(C.GetType(__config))
	return __retVal
}

// GetType 
//  @brief Returns the type of the current configuration node.
//
//  @param config: Pointer to the Config object.
//
//  @return The NodeType enum value.
func GetType(config uintptr) int32 {
	return _GetType(config)
}

var _IsNull = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.IsNull(__config))
	return __retVal
}

// IsNull 
//  @brief Checks if the current node is null.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the node is null, false otherwise.
func IsNull(config uintptr) bool {
	return _IsNull(config)
}

var _IsBool = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.IsBool(__config))
	return __retVal
}

// IsBool 
//  @brief Checks if the current node is a boolean.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the node is a boolean, false otherwise.
func IsBool(config uintptr) bool {
	return _IsBool(config)
}

var _IsInt = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.IsInt(__config))
	return __retVal
}

// IsInt 
//  @brief Checks if the current node is an integer.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the node is an integer, false otherwise.
func IsInt(config uintptr) bool {
	return _IsInt(config)
}

var _IsFloat = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.IsFloat(__config))
	return __retVal
}

// IsFloat 
//  @brief Checks if the current node is a float.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the node is a float, false otherwise.
func IsFloat(config uintptr) bool {
	return _IsFloat(config)
}

var _IsString = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.IsString(__config))
	return __retVal
}

// IsString 
//  @brief Checks if the current node is a string.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the node is a string, false otherwise.
func IsString(config uintptr) bool {
	return _IsString(config)
}

var _IsObject = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.IsObject(__config))
	return __retVal
}

// IsObject 
//  @brief Checks if the current node is an object.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the node is an object, false otherwise.
func IsObject(config uintptr) bool {
	return _IsObject(config)
}

var _IsArray = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.IsArray(__config))
	return __retVal
}

// IsArray 
//  @brief Checks if the current node is an array.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the node is an array, false otherwise.
func IsArray(config uintptr) bool {
	return _IsArray(config)
}

