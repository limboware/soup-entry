package soupev1

import (
	"fmt"
	"os"
	"time"

	limbov1 "github.com/limboware/limbo"
	configsv2 "github.com/limboware/pkg/configs/v2"
	env2 "github.com/limboware/pkg/env/v2"
	errnov1 "github.com/rejchev/errno"
	"github.com/untrustedmodders/go-plugify"
	"github.com/untrustedmodders/s2sdkv213/s2sdk"
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
	if err := configsv2.GetConfig().Parse(os.Args); errnov1.FAIL(err) {
		return err
	}

	if err := configsv2.Get().ParseEnvi(os.Environ()); errnov1.FAIL(err) {
		return err
	}

	env2.Receive(configsv2.GetConfig().ValueA("env", "prod"))

	// loggerv1.Get().System()

	plugify.Log(fmt.Sprintf("soupe in %s mod", env2.Get().String()), plugify.Info, "test", 0)

	// if err := loggerv1.Get().Init(); errnov1.FAIL(err) {
	// 	return err
	// }

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
	limbov1.SystemRegister[NetSystem](netSystem)
	limbov1.SystemRegister[DestroySystem](destroySystem)

	if err := limbov1.GetWorld().Init(); errnov1.FAIL(err) {
		return err
	}

	flags := s2sdk.ConVarFlag_LinkedConcommand | s2sdk.ConVarFlag_Release | s2sdk.ConVarFlag_ClientCanExecute

	s2sdk.AddConsoleCommand("sm_sticket", "Get ticket to connect", flags, x.onClientCommand, s2sdk.HookMode_Post)

	return errnov1.OK
}

func (x *Server) Update(dt float32) {

	if limbov1.GetWorld().Activate() {
		limbov1.GetWorld().Update(time.Duration(dt * float32(time.Second)))
		limbov1.GetWorld().Deactivate()
	}

	limbov1.Events().PostPublisherRun()
}

func (x *Server) onClientCommand(caller int32, context s2sdk.ConCommandContext, arguments []string) s2sdk.ResultType {

	if caller != -1 {
		if handle := s2sdk.PlayerSlotToEntHandle(caller); s2sdk.IsValidEntHandle(handle) && s2sdk.IsClientInGame(caller) && !s2sdk.IsClientSourceTV(caller) {
			limbov1.Events().PublishAsync("s2sdk.cmd", &UseCommand{
				Caller:  caller,
				Context: context,
				Args:    arguments,
			})
		}
	}

	return s2sdk.ResultType_Handled
}

func (x *Server) Destroy() {

	limbov1.GetWorld().Destroy()
	limbov1.Workers().Close()
	// limbov1.Networks().
}
