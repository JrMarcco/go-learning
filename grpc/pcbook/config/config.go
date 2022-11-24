package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	db dbCfg `yaml:"db"`
}

type dbCfg struct {
	Driver  string `yaml:"driver"`
	Source  string `yaml:"source"`
	MaxIdle int    `yaml:"maxIdle"`
	MaxOpen int    `yaml:"maxOpen"`
}

func load(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}

var cfg Config

func init() {
	config, err := load("..")
	if err != nil {
		log.Fatalln(err)
	}

	cfg = config
}

func DbCfg() dbCfg {
	return cfg.db
}
