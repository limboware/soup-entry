package plconfigs

/*
#include "setmethods.h"
#cgo noescape SetNull
#cgo noescape SetObject
#cgo noescape SetArray
#cgo noescape SetBool
#cgo noescape SetInt
#cgo noescape SetFloat
#cgo noescape SetString
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

// Generated from configs (group: setmethods)

var _SetNull = func(config uintptr, key string) {
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	plugify.Block {
		Try: func() {
			C.SetNull(__config, (*C.String)(unsafe.Pointer(&__key)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetNull 
//  @brief Sets a configuration value to null.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
func SetNull(config uintptr, key string) {
	_SetNull(config, key)
}

var _SetObject = func(config uintptr, key string) {
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	plugify.Block {
		Try: func() {
			C.SetObject(__config, (*C.String)(unsafe.Pointer(&__key)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetObject 
//  @brief Sets a configuration value to an empty object.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
func SetObject(config uintptr, key string) {
	_SetObject(config, key)
}

var _SetArray = func(config uintptr, key string) {
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	plugify.Block {
		Try: func() {
			C.SetArray(__config, (*C.String)(unsafe.Pointer(&__key)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetArray 
//  @brief Sets a configuration value to an empty array.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
func SetArray(config uintptr, key string) {
	_SetArray(config, key)
}

var _SetBool = func(config uintptr, key string, value bool) {
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__value := C.bool(value)
	plugify.Block {
		Try: func() {
			C.SetBool(__config, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetBool 
//  @brief Sets a boolean configuration value.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
//  @param value: The boolean value to set.
func SetBool(config uintptr, key string, value bool) {
	_SetBool(config, key, value)
}

var _SetInt = func(config uintptr, key string, value int64) {
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__value := C.int64_t(value)
	plugify.Block {
		Try: func() {
			C.SetInt(__config, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetInt 
//  @brief Sets an integer configuration value.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
//  @param value: The integer value to set.
func SetInt(config uintptr, key string, value int64) {
	_SetInt(config, key, value)
}

var _SetFloat = func(config uintptr, key string, value float64) {
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__value := C.double(value)
	plugify.Block {
		Try: func() {
			C.SetFloat(__config, (*C.String)(unsafe.Pointer(&__key)), __value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
}

// SetFloat 
//  @brief Sets a float configuration value.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
//  @param value: The float value to set.
func SetFloat(config uintptr, key string, value float64) {
	_SetFloat(config, key, value)
}

var _SetString = func(config uintptr, key string, value string) {
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			C.SetString(__config, (*C.String)(unsafe.Pointer(&__key)), (*C.String)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetString 
//  @brief Sets a string configuration value.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to set.
//  @param value: The string value to set.
func SetString(config uintptr, key string, value string) {
	_SetString(config, key, value)
}

