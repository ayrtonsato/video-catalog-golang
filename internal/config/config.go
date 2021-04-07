package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDR"`
	Port          string `mapstructure:"SERVER_PORT"`
	DBDriver      string `mapstructure:"DB_DRIVE"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBDatabase    string `mapstructure:"DB_DATABASE"`
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
}

func (c *Config) Load() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&c)
	fmt.Println(c)
	if err != nil {
		return err
	}
	return nil
}
