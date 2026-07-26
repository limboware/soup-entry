package env2

import env1 "github.com/limboware/pkg/env/v1"

var env = env1.UnknownEnv

// Get is ret current env state
func Get() env1.EnvType {
	return env
}

// Set is set u env state
func Set(val env1.EnvType) {
	env = val
}

// Reset is change current env state to Unknown state (eq UnknownEnv)
func Reset() {
	Set(env1.UnknownEnv)
}

// Parse is parsing env string and ret it (like env1.ParseEnv)
func Parse(val string) env1.EnvType {
	return env1.ParseEnv(val)
}

// Receive is parsing (only if current state is Unknown) and save parsed value into package glob var (sets DevEnv on parse fail)
func Receive(val string) env1.EnvType {
	return ReceiveWithDefault(val, env1.DevEnv)
}

// ReceiveWithDefault is parsing (only if current state is Unknown) and save parsed value into package glob var (sets def on parse fail)
func ReceiveWithDefault(val string, def env1.EnvType) env1.EnvType {
	if Get() == env1.UnknownEnv {
		if result := env1.ParseEnv(val); result != env1.UnknownEnv {
			def = result
		}

		Set(def)
	}

	return Get()
}
