package plconfigs

/*
#include "navigationmethods.h"
#cgo noescape JumpFirst
#cgo noescape JumpLast
#cgo noescape JumpNext
#cgo noescape JumpPrev
#cgo noescape JumpKey
#cgo noescape JumpN
#cgo noescape JumpBack
#cgo noescape JumpRoot
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

// Generated from configs (group: navigationmethods)

var _JumpFirst = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.JumpFirst(__config))
	return __retVal
}

// JumpFirst 
//  @brief Jumps to the first child node in the current configuration node.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the jump was successful, false otherwise.
func JumpFirst(config uintptr) bool {
	return _JumpFirst(config)
}

var _JumpLast = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.JumpLast(__config))
	return __retVal
}

// JumpLast 
//  @brief Jumps to the last child node in the current configuration node.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the jump was successful, false otherwise.
func JumpLast(config uintptr) bool {
	return _JumpLast(config)
}

var _JumpNext = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.JumpNext(__config))
	return __retVal
}

// JumpNext 
//  @brief Jumps to the next sibling node in the current configuration level.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the jump was successful, false otherwise.
func JumpNext(config uintptr) bool {
	return _JumpNext(config)
}

var _JumpPrev = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.JumpPrev(__config))
	return __retVal
}

// JumpPrev 
//  @brief Jumps to the previous sibling node in the current configuration level.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the jump was successful, false otherwise.
func JumpPrev(config uintptr) bool {
	return _JumpPrev(config)
}

var _JumpKey = func(config uintptr, key string, create bool) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__key := plugify.ConstructString(key)
	__create := C.bool(create)
	plugify.Block {
		Try: func() {
			__retVal = bool(C.JumpKey(__config, (*C.String)(unsafe.Pointer(&__key)), __create))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__key)
		},
	}.Do()
	return __retVal
}

// JumpKey 
//  @brief Jumps to a child node with the specified key, optionally creating it if it doesn't exist.
//
//  @param config: Pointer to the Config object.
//  @param key: The key of the node to jump to.
//  @param create: If true, creates the node if it doesn't exist.
//
//  @return True if the jump was successful, false otherwise.
func JumpKey(config uintptr, key string, create bool) bool {
	return _JumpKey(config, key, create)
}

var _JumpN = func(config uintptr, n int32) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__n := C.int32_t(n)
	__retVal = bool(C.JumpN(__config, __n))
	return __retVal
}

// JumpN 
//  @brief Jumps to the nth child node in the current configuration node.
//
//  @param config: Pointer to the Config object.
//  @param n: The index of the child node to jump to.
//
//  @return True if the jump was successful, false otherwise.
func JumpN(config uintptr, n int32) bool {
	return _JumpN(config, n)
}

var _JumpBack = func(config uintptr) bool {
	var __retVal bool
	__config := C.uintptr_t(config)
	__retVal = bool(C.JumpBack(__config))
	return __retVal
}

// JumpBack 
//  @brief Jumps back to the parent node in the configuration hierarchy.
//
//  @param config: Pointer to the Config object.
//
//  @return True if the jump was successful, false otherwise.
func JumpBack(config uintptr) bool {
	return _JumpBack(config)
}

var _JumpRoot = func(config uintptr) {
	__config := C.uintptr_t(config)
	C.JumpRoot(__config)
}

// JumpRoot 
//  @brief Jumps back to the root node of the configuration.
//
//  @param config: Pointer to the Config object.
func JumpRoot(config uintptr) {
	_JumpRoot(config)
}

