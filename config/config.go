package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Username string `json:"username" yamal: "username"`
	Password string `json:"password" yamal: "password"`
	Database string `json:"database" yamal: "database"`
	Host     string `json:"host" yamal: "host"`
	Port     string `json:"port" yamal: "port"`
}

func (config Config) GetURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
}
func SetupConfig() (Config, error) {
	var config Config
	environment := "development"
	if os.Getenv("ENV") != "" {
		environment = os.Getenv("ENV")
	}
	configFile, err := ioutil.ReadFile(fmt.Sprintf(`config/%s.yml`, environment))
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
