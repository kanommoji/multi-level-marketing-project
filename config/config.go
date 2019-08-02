package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"multi-level-marketing-project/model"
	"os"

	"gopkg.in/yaml.v2"
)

func SetupConfig() (model.Config, error) {
	var config model.Config
	environment := "development"
	if os.Getenv("ENV") != "" {
		environment = os.Getenv("ENV")
	}
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
	if os.Getenv("PORT") != "" {
		config.Port = os.Getenv("PORT")
	}
	if os.Getenv("USERNAME") != "" {
		config.Username = os.Getenv("USERNAME")
	}
	if os.Getenv("PASSWORD") != "" {
		config.Password = os.Getenv("PASSWORD")
	}
	if os.Getenv("DATABASE") != "" {
		config.Database = os.Getenv("DATABASE")
	}
	if os.Getenv("HOST") != "" {
		config.Host = os.Getenv("HOST")
	}
	return config, nil
}
