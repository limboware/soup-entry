package soupev1

import (
	"time"
	"unsafe"

	limbov1 "github.com/limboware/pkg/limbo/v1"
	errnov1 "github.com/rejchev/errno"
)

var DestroyType = limbov1.Compotype(0)

type Destroy struct{}

var _ limbov1.ISystem = (*Destroy)(nil)

type DestroySystem struct{}

// Activate implements [limbov1.ISystem].
func (x *Destroy) Activate() bool {
	return true
}

// Deactivate implements [limbov1.ISystem].
func (x *Destroy) Deactivate() {
}

// Load implements [limbov1.ISystem].
func (x *Destroy) Load() errnov1.Code {
	if !limbov1.GetWorld().CreateCompotype(func() unsafe.Pointer { return unsafe.Pointer(new(Destroy)) }, &DestroyType) {
		return errnov1.ECALL
	}

	return errnov1.OK
}

// OnAllLoaded implements [limbov1.ISystem].
func (x *Destroy) OnAllLoaded() {
}

// Unload implements [limbov1.ISystem].
func (x *Destroy) Unload() {
}

// Update implements [limbov1.ISystem].
func (x *Destroy) Update(dt time.Duration) {
}
