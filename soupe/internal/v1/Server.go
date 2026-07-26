package soupev1

import (
	"fmt"
	"os"
	"reflect"
	"time"

	configsv2 "github.com/limboware/pkg/configs/v2"
	env2 "github.com/limboware/pkg/env/v2"
	limbov1 "github.com/limboware/pkg/limbo/v1"
	errnov1 "github.com/rejchev/errno"
	"github.com/untrustedmodders/go-plugify"
)

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

	if err := limbov1.NetworkHandlers().Init(); errnov1.FAIL(err) {
		return err
	}

	if err := limbov1.Events().Init(); errnov1.FAIL(err) {
		return err
	}

	if err := limbov1.Networks().Init(); errnov1.FAIL(err) {
		return err
	}

	if err := limbov1.Workers().Init(); errnov1.FAIL(err) {
		return err
	}

	limbov1.GetWorld().CreateSystem(reflect.TypeFor[InitSystem]().Name(), new(InitSystem))
	limbov1.GetWorld().CreateSystem(reflect.TypeFor[NetInputSystem]().Name(), new(NetInputSystem))
	limbov1.GetWorld().CreateSystem(reflect.TypeFor[NetSystem]().Name(), new(NetSystem))

	if err := limbov1.GetWorld().Load(); errnov1.FAIL(err) {
		return err
	}

	limbov1.GetWorld().OnAllLoaded()

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
	limbov1.GetWorld().Unload()
	limbov1.Workers().Close()
	// limbov1.Networks().
}
