package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	UserServiceHost string
	UserServicePort int

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load(path string) Config {
	godotenv.Load(path + "/.env")

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		Environment:     conf.GetString("ENVIRONMENT"),
		HTTPPort:        conf.GetString("HTTP_PORT"),
		UserServiceHost: conf.GetString("USER_SERVICE_HOST"),
		UserServicePort: conf.GetInt("USER_SERVICE_PORT"),
		CtxTimeout:      conf.GetInt("CTX_TIMEOUT"),
		LogLevel:        conf.GetString("LOG_LEVEL"),
	}

	return cfg
}
