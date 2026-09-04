package plconfigs

/*
#include "getasmethods.h"
#cgo noescape GetAsBool
#cgo noescape GetAsInt
#cgo noescape GetAsFloat
#cgo noescape GetAsString
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

// Generated from configs (group: getasmethods)

var _GetAsBool = func(config uintptr, key string, defaultValue bool) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__defaultValue := C.bool(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.GetAsBool(__config, (*C.String)(unsafe.Pointer(&__key)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetAsBool 
//  @brief Retrieves a value from the configuration and converts it to boolean.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to retrieve.
//  @param defaultValue: The default value to return if the key is not found.
//
//  @return The boolean value if found and convertible, otherwise the default value.
func GetAsBool(config uintptr, key string, defaultValue bool) bool {
	return _GetAsBool(config, key, defaultValue)
}

var _GetAsInt = func(config uintptr, key string, defaultValue int64) int64 {
	var __retVal int64
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__defaultValue := C.int64_t(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = int64(C.GetAsInt(__config, (*C.String)(unsafe.Pointer(&__key)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetAsInt 
//  @brief Retrieves a value from the configuration and converts it to integer.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to retrieve.
//  @param defaultValue: The default value to return if the key is not found.
//
//  @return The integer value if found and convertible, otherwise the default value.
func GetAsInt(config uintptr, key string, defaultValue int64) int64 {
	return _GetAsInt(config, key, defaultValue)
}

var _GetAsFloat = func(config uintptr, key string, defaultValue float64) float64 {
	var __retVal float64
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__defaultValue := C.double(defaultValue)
	plugify.Block {
		Try: func() {
			__retVal = float64(C.GetAsFloat(__config, (*C.String)(unsafe.Pointer(&__key)), __defaultValue))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// GetAsFloat 
//  @brief Retrieves a value from the configuration and converts it to float.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to retrieve.
//  @param defaultValue: The default value to return if the key is not found.
//
//  @return The float value if found and convertible, otherwise the default value.
func GetAsFloat(config uintptr, key string, defaultValue float64) float64 {
	return _GetAsFloat(config, key, defaultValue)
}

var _GetAsString = func(config uintptr, key string, defaultValue string) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__defaultValue := plugify.ConstructString(defaultValue)
	plugify.Block {
		Try: func() {
			__native := C.GetAsString(__config, (*C.String)(unsafe.Pointer(&__key)), (*C.String)(unsafe.Pointer(&__defaultValue)))
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

// GetAsString 
//  @brief Retrieves a value from the configuration and converts it to string.
//
//  @param config: Pointer to the Config object.
//  @param key: The key to retrieve.
//  @param defaultValue: The default value to return if the key is not found.
//
//  @return The string value if found and convertible, otherwise the default value.
func GetAsString(config uintptr, key string, defaultValue string) string {
	return _GetAsString(config, key, defaultValue)
}

