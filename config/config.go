package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Token string
	Botid string
}

func NewEnv() *Config {
	var config Config

	_ , err := toml.DecodeFile("config.toml",&config)
	if err != nil {
		log.Fatal("cannot find config.toml")
	}
	return &config
}