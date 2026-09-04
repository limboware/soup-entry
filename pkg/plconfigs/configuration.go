package plconfigs

/*
#include "configuration.h"
#cgo noescape Read
#cgo noescape ReadMultiple
#cgo noescape Make
#cgo noescape Delete
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

// Generated from configs (group: configuration)

var _Read = func(path string) uintptr {
	var __retVal uintptr
	__path := plugify.ConstructString(path)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.Read((*C.String)(unsafe.Pointer(&__path))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__path)
		},
	}.Do()
	return __retVal
}

// Read 
//  @brief Reads a configuration file from the specified path.
//
//  @param path: Path to the configuration file.
//
//  @return Pointer to the Config object, or nullptr if reading failed.
func Read(path string) uintptr {
	return _Read(path)
}

var _ReadMultiple = func(paths []string) uintptr {
	var __retVal uintptr
	__paths := plugify.ConstructVectorString(paths)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.ReadMultiple((*C.Vector)(unsafe.Pointer(&__paths))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__paths)
		},
	}.Do()
	return __retVal
}

// ReadMultiple 
//  @brief Reads multiple configuration files and merges them.
//
//  @param paths: Vector of paths to configuration files.
//
//  @return Pointer to the merged Config object, or nullptr if reading failed.
func ReadMultiple(paths []string) uintptr {
	return _ReadMultiple(paths)
}

var _Make = func() uintptr {
	__retVal := uintptr(C.Make())
	return __retVal
}

// Make 
//  @brief Creates a new empty configuration object.
//
//
//  @return Pointer to the newly created Config object.
func Make() uintptr {
	return _Make()
}

var _Delete = func(config uintptr) {
	__config := C.uintptr_t(config)
	C.Delete(__config)
}

// Delete 
//  @brief Deletes a configuration object and frees its memory.
//
//  @param config: Pointer to the Config object to delete.
func Delete(config uintptr) {
	_Delete(config)
}

