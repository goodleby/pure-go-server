package env

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Load is used to load env variables into a struct
func Load() (Variables, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error while reading .env file: %s", err.Error())
	}

	var vars Variables
	err = envconfig.Process("", &vars)
	if err != nil {
		return vars, err
	}

	return vars, nil
}

// Variables contains all environment variables
type Variables struct {
	ListenAddress string `envconfig:"LISTEN_ADDRESS" default:"localhost:3030"`
}
