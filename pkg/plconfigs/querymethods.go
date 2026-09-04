package plconfigs

/*
#include "querymethods.h"
#cgo noescape HasKey
#cgo noescape Empty
#cgo noescape GetSize
#cgo noescape GetName
#cgo noescape GetPath
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

// Generated from configs (group: querymethods)

var _HasKey = func(config uintptr, key string) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.HasKey(__config, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// HasKey 
//  @brief Checks if a key exists in the current configuration node.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to check.
//
//  @return True if the key exists, false otherwise.
func HasKey(config uintptr, key string) bool {
	return _HasKey(config, key)
}

var _Empty = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.Empty(__config))
	return __retVal
}

// Empty 
//  @brief Checks if the current configuration node is empty.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the node is empty, false otherwise.
func Empty(config uintptr) bool {
	return _Empty(config)
}

var _GetSize = func(config uintptr) int64 {
	var __retVal int64
	__config := C.uintptr_t(config)
	__retVal = int64(C.GetSize(__config))
	return __retVal
}

// GetSize 
//  @brief Returns the size of the current configuration node (number of elements in array or object).
//
//  @param config: Pointer to the Config object.
//
//  @return The size of the node.
func GetSize(config uintptr) int64 {
	return _GetSize(config)
}

var _GetName = func(config uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__config := C.uintptr_t(config)
	plugify.Block {
		Try: func() {
			__native := C.GetName(__config)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetName 
//  @brief Returns the name of the current configuration node.
//
//  @param config: Pointer to the Config object.
//
//  @return The name of the node.
func GetName(config uintptr) string {
	return _GetName(config)
}

var _GetPath = func(config uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__config := C.uintptr_t(config)
	plugify.Block {
		Try: func() {
			__native := C.GetPath(__config)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetPath 
//  @brief Returns the full path to the current configuration node.
//
//  @param config: Pointer to the Config object.
//
//  @return The path to the node.
func GetPath(config uintptr) string {
	return _GetPath(config)
}

