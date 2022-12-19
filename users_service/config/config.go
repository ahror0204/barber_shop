package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort string
	Postgres PostgresConfig

	Redis Redis

	Smtp Smtp

	OrderServiceHost string
	OrderServicePort string

	RPCPort          string
	LogLevel         string

	AuthSecretKey string
}
type Redis struct {
	Addr string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DataBase string
}

type Smtp struct {
	Sender   string
	Password string
}


func Load(path string) Config {
	godotenv.Load(path + "/.env") // load .env file if it exists

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		HTTPPort: conf.GetString("HTTP_PORT"),
		Postgres: PostgresConfig{
			Host:     conf.GetString("POSTGRES_HOST"),
			Port:     conf.GetString("POSTGRES_PORT"),
			User:     conf.GetString("POSTGRES_USER"),
			Password: conf.GetString("POSTGRES_PASSWORD"),
			DataBase: conf.GetString("POSTGRES_DATABASE"),
		},
		Redis: Redis{
			Addr: conf.GetString("REDIS_ADDR"),
		},
		Smtp: Smtp{
			Sender:   conf.GetString("SMTP_SENDER"),
			Password: conf.GetString("SMTP_PASSWORD"),
		},
		OrderServiceHost: conf.GetString("ORDER_SERVICE_HOST"),
		OrderServicePort: conf.GetString("ORDER_SERVICE_PORT"),
		RPCPort: conf.GetString("RPC_PORT"),
		LogLevel: conf.GetString("LOG_LEVEL"),

		AuthSecretKey: conf.GetString("AUTH_SECRET_KEY"),
	}
	return cfg
}
