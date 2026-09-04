package plconfigs

/*
#include "getmethods.h"
#cgo noescape GetBool
#cgo noescape GetInt
#cgo noescape GetFloat
#cgo noescape GetString
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

// Generated from configs (group: getmethods)

var _GetBool = func(config uintptr, key string, defaultValue bool) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__defaultValue := C.bool(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.GetBool(__config, (*C.String)(unsafe.Pointer(&__key)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetBool 
//  @brief Retrieves a boolean value from the configuration.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to retrieve.
//  @param defaultValue: The default value to return if the key is not found.
//
//  @return The boolean value if found, otherwise the default value.
func GetBool(config uintptr, key string, defaultValue bool) bool {
	return _GetBool(config, key, defaultValue)
}

var _GetInt = func(config uintptr, key string, defaultValue int64) int64 {
	var __retVal int64
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__defaultValue := C.int64_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = int64(C.GetInt(__config, (*C.String)(unsafe.Pointer(&__key)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetInt 
//  @brief Retrieves an integer value from the configuration.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to retrieve.
//  @param defaultValue: The default value to return if the key is not found.
//
//  @return The integer value if found, otherwise the default value.
func GetInt(config uintptr, key string, defaultValue int64) int64 {
	return _GetInt(config, key, defaultValue)
}

var _GetFloat = func(config uintptr, key string, defaultValue float64) float64 {
	var __retVal float64
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__defaultValue := C.double(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = float64(C.GetFloat(__config, (*C.String)(unsafe.Pointer(&__key)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetFloat 
//  @brief Retrieves a float value from the configuration.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to retrieve.
//  @param defaultValue: The default value to return if the key is not found.
//
//  @return The float value if found, otherwise the default value.
func GetFloat(config uintptr, key string, defaultValue float64) float64 {
	return _GetFloat(config, key, defaultValue)
}

var _GetString = func(config uintptr, key string, defaultValue string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__defaultValue := plugify.ConstructString(defaultValue)
	plugify.Block {
		Try: func() {
			__native := C.GetString(__config, (*C.String)(unsafe.Pointer(&__key)), (*C.String)(unsafe.Pointer(&__defaultValue)))
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
			plugify.DestroyString(&__key)
			plugify.DestroyString(&__defaultValue)
		},
	}.Do()
	return __retVal
}

// GetString 
//  @brief Retrieves a string value from the configuration.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to retrieve.
//  @param defaultValue: The default value to return if the key is not found.
//
//  @return The string value if found, otherwise the default value.
func GetString(config uintptr, key string, defaultValue string) string {
	return _GetString(config, key, defaultValue)
}

