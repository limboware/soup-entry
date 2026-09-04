package plconfigs

/*
#include "removalmethods.h"
#cgo noescape Remove
#cgo noescape RemoveKey
#cgo noescape Clear
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

// Generated from configs (group: removalmethods)

var _Remove = func(config uintptr) int32 {
	var __retVal int32
	__config := C.uintptr_t(config)
	__retVal = int32(C.Remove(__config))
	return __retVal
}

// Remove 
//  @brief Removes the current configuration node.
//
//  @param config: Pointer to the Config object.
//
//  @return The number of nodes removed.
func Remove(config uintptr) int32 {
	return _Remove(config)
}

var _RemoveKey = func(config uintptr, key string) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.RemoveKey(__config, (*C.String)(unsafe.Pointer(&__key))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// RemoveKey 
//  @brief Removes a child node with the specified key from the current configuration node.
//
//  @param config: Pointer to the Config object.
//  @param key: The key of the node to remove.
//
//  @return True if the node was removed successfully, false otherwise.
func RemoveKey(config uintptr, key string) bool {
	return _RemoveKey(config, key)
}

var _Clear = func(config uintptr) {
	__config := C.uintptr_t(config)
	C.Clear(__config)
}

// Clear 
//  @brief Clears all child nodes from the current configuration node.
//
//  @param config: Pointer to the Config object.
func Clear(config uintptr) {
	_Clear(config)
}

