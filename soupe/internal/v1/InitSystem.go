package soupev1

import (
	"time"
	"unsafe"

	limbov1 "github.com/limboware/limbo"
	errnov1 "github.com/rejchev/errno"
)

var SERVER = limbov1.Entity(0)

const MAX_CLIENTS = 64

type InitSystem struct{}

func initSystem(buff *limbov1.System) {
	v := new(InitSystem)

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
func (x *InitSystem) Activate() bool {
	if !limbov1.Entities().IsAlive(SERVER) {
		// Create SERVER
		SERVER = limbov1.GetWorld().CreateEntity()

		// TODO: Create brokers (after loading config with known size)
		// ...
	}

	return true
}

// Deactivate implements [limbov1.ISystem].
func (x *InitSystem) Deactivate() {}

// Load implements [limbov1.ISystem].
func (x *InitSystem) Load() errnov1.Code {
	return errnov1.OK
}

// OnAllLoaded implements [limbov1.ISystem].
func (x *InitSystem) OnAllLoaded() {}

// Unload implements [limbov1.ISystem].
func (x *InitSystem) Unload() {}

// Update implements [limbov1.ISystem].
func (x *InitSystem) Update(dt time.Duration) {}
