package soupev1

import (
	"time"

	limbov1 "github.com/limboware/limbo"
	errnov1 "github.com/rejchev/errno"
)

const SERVER = limbov1.Entity(0)

const MAX_CLIENTS = 64

var _ limbov1.ISystem = (*InitSystem)(nil)

type InitSystem struct{}

// Activate implements [limbov1.ISystem].
func (x *InitSystem) Activate() bool {
	if !limbov1.Entities().IsAlive(SERVER) {

		// Create SERVER
		_ = limbov1.GetWorld().CreateEntity()

		// Create clients like CS2 engine
		for range MAX_CLIENTS {
			_ = limbov1.NewComponentB[Client](limbov1.GetWorld().CreateEntity(), ClientType)
		}

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
