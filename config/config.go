package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"multi-level-marketing-project/model"

	"gopkg.in/yaml.v2"
)

func SetupConfig() (model.Config, error) {
	var config model.Config
	environment := "development"
	configFile, err := ioutil.ReadFile(fmt.Sprintf("config/%s.yml", environment))
	if err != nil {
		fmt.Printf("cannot read config file %s", err)
		return config, err
	}
	err = yaml.NewDecoder(bytes.NewReader(configFile)).Decode(&config)
	if err != nil {
		fmt.Printf("cannot decode config file %s", err)
		return config, err
	}
	return config, nil
}
