package model

type Config struct {
	Username string `yamal: "username"`
	Password string `yamal: "password"`
	Database string `yamal: "database"`
	Host     string `yamal: "host"`
	Port     string `yamal: "port"`
}
