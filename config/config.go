package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Redis            RedisConfig `mapstructure:"redis"`
	Token            string      `mapstructure:"token"`
	OpenWeatherToken string      `mapstructure:"open_weather_token"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
}

func (r *RedisConfig) ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       0,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis: %s", err))
	}

	return rdb
}

const envPrefix = "SADBOT"

// New creates a new config with defaults
func New() *Config {

	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("redis.address", "127.0.0.1:6379")
	viper.SetDefault("redis.password", "")

	viper.SetDefault("token", "123456:qwerty")
	viper.SetDefault("open_weather_token", "qwerty")

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		logrus.Fatal(err)
	}
	return &c
}
