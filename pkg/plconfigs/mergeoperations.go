package plconfigs

/*
#include "mergeoperations.h"
#cgo noescape Merge
#cgo noescape MergeMove
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

// Generated from configs (group: mergeoperations)

var _Merge = func(config uintptr, other uintptr) {
	__config := C.uintptr_t(config)
	__other := C.uintptr_t(other)
	C.Merge(__config, __other)
}

// Merge 
//  @brief Merges another configuration object into the current one.
//
//  @param config: Pointer to the target Config object.
//  @param other: Pointer to the Config object to merge from.
func Merge(config uintptr, other uintptr) {
	_Merge(config, other)
}

var _MergeMove = func(config uintptr, other uintptr) {
	__config := C.uintptr_t(config)
	__other := C.uintptr_t(other)
	C.MergeMove(__config, __other)
}

// MergeMove 
//  @brief Merges another configuration object into the current one with move semantics.
//
//  @param config: Pointer to the target Config object.
//  @param other: Pointer to the Config object to merge from (will be moved).
func MergeMove(config uintptr, other uintptr) {
	_MergeMove(config, other)
}

