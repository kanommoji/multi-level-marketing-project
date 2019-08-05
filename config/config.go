package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Username string `yamal: "username"`
	Password string `yamal: "password"`
	Database string `yamal: "database"`
	Host     string `yamal: "host"`
	Port     string `yamal: "port"`
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
	if os.Getenv("PORT_DB") != "" {
		config.Port = os.Getenv("PORT_DB")
	}
	if os.Getenv("USERNAME_DB") != "" {
		config.Username = os.Getenv("USERNAME_DB")
	}
	if os.Getenv("PASSWORD_DB") != "" {
		config.Password = os.Getenv("PASSWORD_DB")
	}
	if os.Getenv("DATABASE_NAME") != "" {
		config.Database = os.Getenv("DATABASE_NAME")
	}
	if os.Getenv("HOST_DB") != "" {
		config.Host = os.Getenv("HOST_DB")
	}
	return config, nil
}
