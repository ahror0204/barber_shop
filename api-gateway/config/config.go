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
	RedisHost  string
	RedisPort  string

	Smtp Smtp

	LogLevel string
	HTTPPort string

	AuthSecretKey string
}

type Smtp struct {
	Sender   string
	Password string
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

		RedisHost: conf.GetString("REDIS_HOST"),
		RedisPort: conf.GetString("REDIS_PORT"),

		Smtp: Smtp{
			Sender: conf.GetString("SMTP_SENDER"),
			Password: conf.GetString("SMTP_PASSWORD"),
		},
		AuthSecretKey: conf.GetString("AUTH_SECRET_KEY"),
	}

	return cfg
}
