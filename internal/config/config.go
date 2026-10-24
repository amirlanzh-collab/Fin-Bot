package config

import "github.com/spf13/viper"

type Config struct {
	Token      string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
}

func LoadEnv() *Config {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	return &Config{
		Token:      viper.GetString("TOKEN"),
		DBHost:     viper.GetString("DB_HOST"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBPort:     viper.GetString("DB_PORT"),
		DBSSLMode:  viper.GetString("DB_SSLMODE"),
	}
}
