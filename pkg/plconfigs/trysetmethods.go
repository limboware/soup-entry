package plconfigs

/*
#include "trysetmethods.h"
#cgo noescape TrySetFromBool
#cgo noescape TrySetFromInt
#cgo noescape TrySetFromFloat
#cgo noescape TrySetFromString
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

// Generated from configs (group: trysetmethods)

var _TrySetFromBool = func(config uintptr, key string, value bool) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__value := C.bool(value)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.TrySetFromBool(__config, (*C.String)(unsafe.Pointer(&__key)), __value))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// TrySetFromBool 
//  @brief Attempts to set a configuration value from a boolean, with type conversion if needed.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
//  @param value: The boolean value to set.
//
//  @return True if the value was set successfully, false otherwise.
func TrySetFromBool(config uintptr, key string, value bool) bool {
	return _TrySetFromBool(config, key, value)
}

var _TrySetFromInt = func(config uintptr, key string, value int64) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__value := C.int64_t(value)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.TrySetFromInt(__config, (*C.String)(unsafe.Pointer(&__key)), __value))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// TrySetFromInt 
//  @brief Attempts to set a configuration value from an integer, with type conversion if needed.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
//  @param value: The integer value to set.
//
//  @return True if the value was set successfully, false otherwise.
func TrySetFromInt(config uintptr, key string, value int64) bool {
	return _TrySetFromInt(config, key, value)
}

var _TrySetFromFloat = func(config uintptr, key string, value float64) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__value := C.double(value)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.TrySetFromFloat(__config, (*C.String)(unsafe.Pointer(&__key)), __value))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// TrySetFromFloat 
//  @brief Attempts to set a configuration value from a float, with type conversion if needed.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
//  @param value: The float value to set.
//
//  @return True if the value was set successfully, false otherwise.
func TrySetFromFloat(config uintptr, key string, value float64) bool {
	return _TrySetFromFloat(config, key, value)
}

var _TrySetFromString = func(config uintptr, key string, value string) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.TrySetFromString(__config, (*C.String)(unsafe.Pointer(&__key)), (*C.String)(unsafe.Pointer(&__value))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
			plugify.DestroyString(&__value)
		},
	}.Do()
	return __retVal
}

// TrySetFromString 
//  @brief Attempts to set a configuration value from a string, with type conversion if needed.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
//  @param value: The string value to set.
//
//  @return True if the value was set successfully, false otherwise.
func TrySetFromString(config uintptr, key string, value string) bool {
	return _TrySetFromString(config, key, value)
}

