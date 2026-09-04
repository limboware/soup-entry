package plconfigs

/*
#include "arraypushmethods.h"
#cgo noescape PushNull
#cgo noescape PushBool
#cgo noescape PushInt
#cgo noescape PushFloat
#cgo noescape PushString
#cgo noescape PushObject
#cgo noescape PushArray
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

// Generated from configs (group: arraypushmethods)

var _PushNull = func(config uintptr) {
	__config := C.uintptr_t(config)
	C.PushNull(__config)
}

// PushNull 
//  @brief Pushes a null value to the current array node.
//
//  @param config: Pointer to the Config object.
func PushNull(config uintptr) {
	_PushNull(config)
}

var _PushBool = func(config uintptr, value bool) {
	__config := C.uintptr_t(config)
	__value := C.bool(value)
	C.PushBool(__config, __value)
}

// PushBool 
//  @brief Pushes a boolean value to the current array node.
//
//  @param config: Pointer to the Config object.
//  @param value: The boolean value to push.
func PushBool(config uintptr, value bool) {
	_PushBool(config, value)
}

var _PushInt = func(config uintptr, value int64) {
	__config := C.uintptr_t(config)
	__value := C.int64_t(value)
	C.PushInt(__config, __value)
}

// PushInt 
//  @brief Pushes an integer value to the current array node.
//
//  @param config: Pointer to the Config object.
//  @param value: The integer value to push.
func PushInt(config uintptr, value int64) {
	_PushInt(config, value)
}

var _PushFloat = func(config uintptr, value float64) {
	__config := C.uintptr_t(config)
	__value := C.double(value)
	C.PushFloat(__config, __value)
}

// PushFloat 
//  @brief Pushes a float value to the current array node.
//
//  @param config: Pointer to the Config object.
//  @param value: The float value to push.
func PushFloat(config uintptr, value float64) {
	_PushFloat(config, value)
}

var _PushString = func(config uintptr, value string) {
	__config := C.uintptr_t(config)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			C.PushString(__config, (*C.String)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// PushString 
//  @brief Pushes a string value to the current array node.
//
//  @param config: Pointer to the Config object.
//  @param value: The string value to push.
func PushString(config uintptr, value string) {
	_PushString(config, value)
}

var _PushObject = func(config uintptr) {
	__config := C.uintptr_t(config)
	C.PushObject(__config)
}

// PushObject 
//  @brief Pushes an empty object to the current array node.
//
//  @param config: Pointer to the Config object.
func PushObject(config uintptr) {
	_PushObject(config)
}

var _PushArray = func(config uintptr) {
	__config := C.uintptr_t(config)
	C.PushArray(__config)
}

// PushArray 
//  @brief Pushes an empty array to the current array node.
//
//  @param config: Pointer to the Config object.
func PushArray(config uintptr) {
	_PushArray(config)
}

