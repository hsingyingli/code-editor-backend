package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER              string        `mapstructure:"DB_DRIVER"`
	DB_USERNAME            string        `mapstructure:"DB_USERNAME"`
	DB_PASSWORD            string        `mapstructure:"DB_PASSWORD"`
	DB_URL                 string        `mapstructure:"DB_URL"`
	DB_PORT                string        `mapstructure:"DB_PORT"`
	DB_TABLE               string        `mapstructure:"DB_TABLE"`
	SERVER_ADDR            string        `mapstructure:"SERVER_ADDR"`
	TOKEN_STMMETRIC_KEY    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	ACCESS_TOKEN_DURATION  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	REFRESH_TOKEN_DURATION time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
