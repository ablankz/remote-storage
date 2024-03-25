// Package config Reads environment variables and creates structures.
//
// The generated structure guarantees the type,
// so when environment variables are used in the application,
// they are called through this structure.
// The structure can also have an initial value or be a required field.
package config

import (
	"fmt"
	"reflect"

	"github.com/caarlos0/env/v10"
)

// Config A structure representing the application settings.
type Config struct {
	GatewayHost               string `env:"GATEWAY_HOST" envDefault:"localhost"`
	GatewayPort               int    `env:"GATEWAY_PORT" envDefault:"50051"`
	GatewayDebuggingPort      int    `env:"GATEWAY_DEBUGGING_PORT" envDefault:"2345"`
	LoggerHost                string `env:"LOGGER_HOST" envDefault:"localhost"`
	LoggerPort                int    `env:"LOGGER_PORT" envDefault:"50052"`
	LoggerDebuggingPort       int    `env:"LOGGER_DEBUGGING_PORT" envDefault:"2346"`
	AuthHost                  string `env:"AUTH_HOST" envDefault:"localhost"`
	AuthPort                  int    `env:"AUTH_PORT" envDefault:"50053"`
	AuthDebuggingPort         int    `env:"AUTH_DEBUGGING_PORT" envDefault:"2347"`
	OrchestratorHost          string `env:"ORCHESTRATOR_HOST" envDefault:"localhost"`
	OrchestratorPort          int    `env:"ORCHESTRATOR_PORT" envDefault:"50054"`
	OrchestratorDebuggingPort int    `env:"ORCHESTRATOR_DEBUGGING_PORT" envDefault:"2348"`
	RepositoryHost            string `env:"REPOSITORY_HOST" envDefault:"localhost"`
	RepositoryPort            int    `env:"REPOSITORY_PORT" envDefault:"50055"`
	RepositoryDebuggingPort   int    `env:"REPOSITORY_DEBUGGING_PORT" envDefault:"2349"`
	StorageHost               string `env:"STORAGE_HOST" envDefault:"localhost"`
	StoragePort               int    `env:"STORAGE_PORT" envDefault:"50056"`
	StorageDebuggingPort      int    `env:"STORAGE_DEBUGGING_PORT" envDefault:"2350"`
	DBHost                    string `env:"DB_HOST" envDefault:"localhost"`
	DBPort                    int    `env:"DB_PORT" envDefault:"5432"`
	DBName                    string `env:"DB_NAME,required"`
	DBUsername                string `env:"DB_USERNAME,required"`
	DBPassword                string `env:"DB_PASSWORD"`
	DBUrl                     string `env:"DB_URL,required"`

	StoragePath string `env:"STORAGE_PATH,required"`

	AppDebug bool `env:"APP_DEBUG"`
	// development, staging, production
	AppEnv EnvironmentMode `env:"APP_ENV" envDefault:"production"`
	// FakeTime Fake time mode setting
	// If a time is specified, fix to that time.
	// If a truthy value is specified, fix to the default time.
	FakeTime FakeTimeMode `env:"FAKE_TIME"`
	LogLevel LogLevel     `env:"LOG_LEVEL,required"`
}

var parseFuncMap = map[reflect.Type]env.ParserFunc{
	reflect.TypeOf(ProductionEnv):  parseEnvironmentMode,
	reflect.TypeOf(FakeTimeMode{}): parseFakeTimeMode,
	reflect.TypeOf(InfoLevel):      parseLogLevel,
}

// Get Get application settings from environment variables.
func Get() (*Config, error) {
	cfg := &Config{}
	if err := env.ParseWithOptions(cfg, env.Options{FuncMap: parseFuncMap}); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return cfg, nil
}
