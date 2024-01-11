package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBAddress              string
	RedisURI               string
	Port                   string
	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	AccessTokenExpiresIn   time.Duration
	RefreshTokenExpiresIn  time.Duration
	AccessTokenMaxAge      int
	RefreshTokenMaxAge     int
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigType("app")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
