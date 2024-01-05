package config

import (
	_ "embed"
	"log/slog"
	"sync"

	v10env "github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/yaml.v3"
)

// the environment in which the application is running
type Environment string

const (
	Development Environment = "development"
	Staging     Environment = "staging"
	Production  Environment = "production"
	Test        Environment = "test"
)

var (
	env     Environment
	envOnce sync.Once
)

// to initialize the env for which the application is running and
// to fetch it globally in the application for initialization we need to pass an env
// if environment is anything other than development.
// if no env is passed then default env is development
func Env(initEnv ...Environment) Environment {
	envOnce.Do(func() {
		if len(initEnv) > 1 {
			slog.Error("env can only be set once", slog.Any("init env", initEnv))
			panic("env can only be set once")
		} else if len(initEnv) == 0 {
			env = Development
		} else {
			env = initEnv[0]
		}
	})
	return env
}

func IsDevelopment() bool {
	return Env() == Development
}

func IsProduction() bool {
	return Env() == Production
}

func IsTest() bool {
	return Env() == Test
}

func IsStaging() bool {
	return Env() == Staging
}

var (
	secrets    secret
	settings   setting
	configOnce sync.Once

	//go:embed settings.yml
	settingsFile []byte // loading the settings file into the binary so that it can be useed
)

func Secrets() secret {
	configOnce.Do(loadConfigs)
	return secrets
}

func Settings() setting {
	configOnce.Do(loadConfigs)
	return settings
}

func loadConfigs() {
	if err := v10env.Parse(&secrets); err != nil {
		slog.Error("unable to load secrets from .env", slog.Any("error", err))
		panic("unable to load secrets from .env")
	}
	if err := yaml.Unmarshal(settingsFile, &settings); err != nil {
		slog.Error("unable to load settings from settings.yml", slog.Any("error", err))
		panic("unable to load settings from settings.yml")
	}
}
