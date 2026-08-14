package soupev1

import (
	"time"
	"unsafe"

	limbov1 "github.com/limboware/limbo"
	errnov1 "github.com/rejchev/errno"
)

var DestroyType = limbov1.Compotype(0)

type Destroy struct {
	Reason string
}

type DestroySystem struct{}

func destroySystem(buff *limbov1.System) {
	v := new(DestroySystem)

	*buff = limbov1.System{
		Instance:   unsafe.Pointer(v),
		Init:       v.Load,
		Activate:   v.Activate,
		Update:     v.Update,
		Deactivate: v.Deactivate,
		Destroy:    v.Unload,
	}
}

// Activate implements [limbov1.ISystem].
func (x *DestroySystem) Activate() bool {
	return true
}

// Deactivate implements [limbov1.ISystem].
func (x *DestroySystem) Deactivate() {
}

// Load implements [limbov1.ISystem].
func (x *DestroySystem) Load() errnov1.Code {
	if !limbov1.GetWorld().CreateCompotype(func() unsafe.Pointer { return unsafe.Pointer(new(Destroy)) }, &DestroyType) {
		return errnov1.ECALL
	}

	// limbov1.Events().Subscribe("world.loaded", x.onAllLoaded)
	limbov1.Events().Subscribe("entity.destroy", x.onEntityDestroy)
	return errnov1.OK
}

func (x *DestroySystem) onEntityDestroy(_ string, data any) {
	if ent, ok := data.(limbov1.Entity); ok {
		limbov1.DestroyComponent(ent, DestroyType)
	}
}

// Unload implements [limbov1.ISystem].
func (x *DestroySystem) Unload() {
}

// Update implements [limbov1.ISystem].
func (x *DestroySystem) Update(dt time.Duration) {
	limbov1.Components().IterateB(DestroyType, func(e limbov1.Entity, p unsafe.Pointer) {
		limbov1.Entities().Destroy(e)
	})
}
