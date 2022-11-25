package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort string
	Postgres PostgresConfig

	OrderServiceHost string
	OrderServicePort string

	RPCPort          string
	LogLevel         string


}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DataBase string
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
		OrderServiceHost: conf.GetString("ORDER_SERVICE_HOST"),
		OrderServicePort: conf.GetString("ORDER_SERVICE_PORT"),
		RPCPort: conf.GetString("RPC_PORT"),
		LogLevel: conf.GetString("LOG_LEVEL"),
	}
	return cfg
}
