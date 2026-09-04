package plconfigs

/*
#include "errorhandling.h"
#cgo noescape SetError
#cgo noescape GetError
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

// Generated from configs (group: errorhandling)

var _SetError = func(error_ string) {
	__error_ := plugify.ConstructString(error_)
	plugify.Block {
		Try: func() {
			C.SetError((*C.String)(unsafe.Pointer(&__error_)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__error_)
		},
	}.Do()
}

// SetError 
//  @brief Sets an error message for the configuration system.
//
//  @param error_: The error message to set.
func SetError(error_ string) {
	_SetError(error_)
}

var _GetError = func() string {
	var __retVal string
	var __retVal_native plugify.PlgString
	plugify.Block {
		Try: func() {
			__native := C.GetError()
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

// GetError 
//  @brief Retrieves the last error message from the configuration system.
//
//
//  @return The last error message.
func GetError() string {
	return _GetError()
}

