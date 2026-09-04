package plconfigs

// Generated from configs

// noCopy prevents copying via go vet
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// ownership indicates whether the instance owns the underlying handle
type ownership bool

const (
	Owned    ownership = true
	Borrowed ownership = false
)

