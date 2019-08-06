package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

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
		fmt.Printf("Cannot read config file %s", err)
		return config, err
	}
	err = yaml.NewDecoder(bytes.NewReader(configFile)).Decode(&config)
	if err != nil {
		fmt.Printf("Cannot decode config file %s", err)
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

func TimeNow() time.Time {
	dateLayout := "20060102"
	dateText := os.Getenv("TIME")
	if dateText != "" {
		timeNow, err := time.Parse(dateLayout, dateText)
		if err != nil {
			log.Printf("Error pasing date from env , %s", err)
		}
		return timeNow
	}
	return time.Now()
}
