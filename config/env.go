package config

import (
	"log/slog"
	"sync"
)

// the environment in which the application is running
type Environment string

const (
	Development Environment = "development"
	Staging Environment = "staging"
	Production Environment = "production"
	Test Environment = "test"
)

var (
	env Environment
	envOnce sync.Once
)

// to initialize the env for which the application is running and 
// to fetch it globally in the application for initialization we need to pass an env 
// if environment is anything other than development.
// if no env is passed then default env is development
func Env(initEnv ...Environment) Environment {
	envOnce.Do(func() {
		if len(initEnv) > 1{
			slog.Error("env can only be set once", slog.Any("init env", initEnv))
			panic("env can only be set once")
		} else if len(initEnv) == 0 {
			env = Development
		} else{
			env = initEnv[0]
		}
	})
	return env
}