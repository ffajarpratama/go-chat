package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App       App
	JWTConfig JWTConfig
	Postgres  Postgres
}

type App struct {
	Name        string
	Environment string
	Host        string
	Port        int
	URL         string
}

type JWTConfig struct {
	Admin string
	User  string
}

type Postgres struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SSLMode  string
	URI      string
}

func New() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("[config-file-fail-load] \n", err.Error())
	}

	v := viper.GetViper()
	viper.AutomaticEnv()

	return &Config{
		App: App{
			Name:        v.GetString("APP_NAME"),
			Environment: v.GetString("APP_ENV"),
			Host:        v.GetString("APP_HOST"),
			Port:        v.GetInt("APP_PORT"),
			URL:         v.GetString("APP_URL"),
		},
		JWTConfig: JWTConfig{
			Admin: v.GetString("JWT_ADMIN"),
			User:  v.GetString("JWT_USER"),
		},
		Postgres: Postgres{
			Host:     v.GetString("POSTGRES_HOST"),
			Port:     v.GetString("POSTGRES_PORT"),
			Database: v.GetString("POSTGRES_DATABASE"),
			User:     v.GetString("POSTGRES_USER"),
			Password: v.GetString("POSTGRES_PASS"),
			SSLMode:  v.GetString("POSTGRES_SSL_MODE"),
			URI:      v.GetString("POSTGRES_URI"),
		},
	}
}
