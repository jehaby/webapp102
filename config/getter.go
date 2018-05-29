package config

import (
	"log"

	"github.com/spf13/viper"
)

func Get(env string) (C, error) {
	viper.SetConfigName(env)
	viper.AddConfigPath("./var/configs/")
	viper.AddConfigPath("./../var/configs")
	viper.AddConfigPath(".")

	res := C{}
	if err := viper.ReadInConfig(); err != nil {
		return res, err
	}
	if err := viper.Unmarshal(&res); err != nil {
		return res, err
	}
	return res, nil
}

func MustGet(env string) C {
	res, err := Get(env)
	if err != nil {
		log.Fatalf("couldn't get config for '%s' env: %v", env, err)
	}
	return res
}
