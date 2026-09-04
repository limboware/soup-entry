package soupev1

import (
	"time"

	limbov1 "github.com/limboware/limbo"
	configsv2 "github.com/limboware/pkg/configs/v2"
	env2 "github.com/limboware/pkg/env/v2"
	loggerv1 "github.com/limboware/pkg/logger/v1"
	errnov1 "github.com/rejchev/errno"
	s2sdk "github.com/untrustedmodders/s2sdk/213"
)

type UseCommand struct {
	Caller  int32
	Context s2sdk.ConCommandContext
	Args    []string
}

type Server struct{}

var srv = Server{}

func Get() *Server {
	return &srv
}

func (x *Server) Init() errnov1.Code {
	if err := loggerv1.Get().Init(); errnov1.FAIL(err) {
		return err
	}

	if err := configsv2.Get().ParsePlug("soupe.json"); errnov1.FAIL(err) {
		return err
	}

	loggerv1.Get().Err("hellow")

	env2.Receive(configsv2.GetConfig().ValueA("env", "prod"))

	if err := loggerv1.Get().Init(); errnov1.FAIL(err) {
		return err
	}

	if err := limbov1.InitNetModule(); errnov1.FAIL(err) {
		return err
	}

	if err := limbov1.Events().Init(); errnov1.FAIL(err) {
		return err
	}

	if err := limbov1.Workers().Init(); errnov1.FAIL(err) {
		return err
	}

	if err := limbov1.Systems().Init(); errnov1.FAIL(err) {
		return err
	}

	if err := Tickets().Init(); errnov1.FAIL(err) {
		return err
	}

	limbov1.SystemRegister[InitSystem](initSystem)
	limbov1.SystemRegister[NetInputSystem](netInputSystem)
	limbov1.SystemRegister[TicketSystem](ticketSystem)
	limbov1.SystemRegister[SlotSystem](slotSystem)
	limbov1.SystemRegister[NetSystem](netSystem)
	limbov1.SystemRegister[DestroySystem](destroySystem)

	if err := limbov1.GetWorld().Init(); errnov1.FAIL(err) {
		return err
	}

	return errnov1.OK
}

func (x *Server) Update(dt float32) {

	if limbov1.GetWorld().Activate() {
		limbov1.GetWorld().Update(time.Duration(dt * float32(time.Second)))
		limbov1.GetWorld().Deactivate()
	}

	limbov1.Events().PostPublisherRun()
}

func (x *Server) Destroy() {

	limbov1.GetWorld().Destroy()
	limbov1.Workers().Close()
	// limbov1.Networks().
}
