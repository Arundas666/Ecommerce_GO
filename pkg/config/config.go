package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	BASE_URL   string `mapstructure:"BASE_URL"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	AUTHTOKEN   string `mapstructure:"TWILIO_AUTHTOKEN"`
	ACCOUNTSID  string `mapstructure:"TWILIO_ACCOUNTSID"`
	SERVICESSID string `mapstructure:"TWILIO_SERVICESID"`

	KEY                          string `mapstructure:"KEY"`
	KEY_FOR_ADMIN                string `mapstructure:"KEY_FOR_ADMIN"`
	KEY_FOR_SHIPPING_COORDINATOR string `mapstructure:"KEY_FOR_SHIPPING_COORDINATOR"`
	KEY_ID_FOR_RAYZORPAY         string `mapstructure:"KEY_ID_FOR_RAYZORPAY"`
	SECRET_KEY_ID_FOR_RAYZORPAY  string `mapstructure:"SECRET_KEY_ID_FOR_RAYZORPAY"`
}

var envs = []string{
	"BASE_URL", "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "TWILIO_AUTHTOKEN", "TWILIO_ACCOUNTSID", "TWILIO_SERVICESID", "KEY", "KEY_FOR_ADMIN", "KEY_ID_FOR_RAYZORPAY", "SECRET_KEY_ID_FOR_RAYZORPAY", "KEY_FOR_SHIPPING_COORDINATOR",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	fmt.Println("CONFFF", config)
	return config, nil

}
