package main

import (
	"fmt"

	soupev1 "github.com/limboware/soupe/internal/v1"
	errnov1 "github.com/rejchev/errno"
	"github.com/untrustedmodders/go-plugify"
)

func onPluginStart() error {
	if err := soupev1.Get().Init(); errnov1.FAIL(err) {
		return fmt.Errorf("soupe err on start up (err: %s)", err.String())
	}

	return nil
}

func onPluginUpdate(dt float32) error {
	soupev1.Get().Update(dt)
	return nil
}

func onPluginEnd() error {
	soupev1.Get().Destroy()
	return nil
}

func init() {
	plugify.NewPlugin("soupe", onPluginStart, onPluginUpdate, onPluginEnd)
}

func main() {}
