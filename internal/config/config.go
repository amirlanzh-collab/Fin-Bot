package config

import "github.com/spf13/viper"

type Config struct {
	Token string
}

func LoadEnv() *Config {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	return &Config{
		Token: viper.GetString("TOKEN"),
	}
}
