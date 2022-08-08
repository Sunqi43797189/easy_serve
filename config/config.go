package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Service Service
}

type Service struct {
	Env       string
	ServeType string
}

var C = &Config{}

func InitConf(configFile string) error {
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(conf, C)
	if err != nil {
		return err
	}
	return nil
}
