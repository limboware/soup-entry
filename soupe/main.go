package main

import (
	"fmt"

	limbov1 "github.com/limboware/limbo"
	soupev1 "github.com/limboware/soupe/internal/v1"
	errnov1 "github.com/rejchev/errno"
	"github.com/untrustedmodders/go-plugify"
	s2sdk "github.com/untrustedmodders/s2sdk/213"
)

func onPluginStart() error {
	if err := soupev1.Get().Init(); errnov1.FAIL(err) {
		return fmt.Errorf("soupe err on start up (err: %s)", err.String())
	}

	s2sdk.OnClientAuthenticated_Register(onClientAuthenticated)
	s2sdk.OnClientDisconnect_Register(onClientDisconnect)

	return nil
}

func onClientAuthenticated(playerSlot int32, steamID uint64) {
	limbov1.Events().PublishAsync("clients.connected", &soupev1.ClientConnectedEvent{
		Slot:    playerSlot,
		SteamID: steamID,
	})
}

func onClientDisconnect(playerSlot int32) {
	limbov1.Events().PublishAsync("clients.disconnect", &soupev1.ClientConnectedEvent{
		Slot:    playerSlot,
	})
}

func onPluginUpdate(dt float32) error {
	soupev1.Get().Update(dt)
	return nil
}

func onPluginEnd() error {
	s2sdk.OnClientAuthenticated_Unregister(onClientAuthenticated)
	s2sdk.OnClientDisconnect_Unregister(onClientDisconnect)
	soupev1.Get().Destroy()
	return nil
}

func init() {
	plugify.NewPlugin("soupe", onPluginStart, onPluginUpdate, onPluginEnd)
}

func main() {}
