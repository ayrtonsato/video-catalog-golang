package setup

import (
	"flag"
	"github.com/spf13/viper"
)

type Config struct {
	ServerMode    string `mapstructure:"SERVER_MODE"`
	ServerAddress string `mapstructure:"SERVER_ADDR"`
	Port          string `mapstructure:"SERVER_PORT"`
	DBDriver      string `mapstructure:"DB_DRIVE"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBDatabase    string `mapstructure:"DB_DATABASE"`
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
}

func (c *Config) Load(path string) error {
	var envFile string
	var p string
	if path == "" {
		p = "."
	} else {
		p = path
	}
	if flag.Lookup("test.v") != nil {
		envFile = ".env.test"
	} else {
		envFile = ".env"
	}
	viper.AddConfigPath(p)
	viper.SetConfigType("env")
	viper.SetConfigName(envFile)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}
	return nil
}
