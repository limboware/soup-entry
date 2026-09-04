package plconfigs

/*
#include "jsonexport.h"
#cgo noescape NodeToJsonString
#cgo noescape RootToJsonString
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

// Generated from configs (group: jsonexport)

var _NodeToJsonString = func(config uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__config := C.uintptr_t(config)
	plugify.Block {
		Try: func() {
			__native := C.NodeToJsonString(__config)
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

// NodeToJsonString 
//  @brief Converts the current configuration node to a JSON string.
//
//  @param config: Pointer to the Config object.
//
//  @return The JSON string representation of the current node.
func NodeToJsonString(config uintptr) string {
	return _NodeToJsonString(config)
}

var _RootToJsonString = func(config uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__config := C.uintptr_t(config)
	plugify.Block {
		Try: func() {
			__native := C.RootToJsonString(__config)
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

// RootToJsonString 
//  @brief Converts the entire configuration tree to a JSON string.
//
//  @param config: Pointer to the Config object.
//
//  @return The JSON string representation of the entire configuration tree.
func RootToJsonString(config uintptr) string {
	return _RootToJsonString(config)
}

