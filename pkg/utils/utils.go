package utils

import (
	"fmt"
	"os"
)

// CheckEnviroment function makes sure all enviroment variables are set befire running
func CheckEnviroment() error {
	variables := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASSWORD",
		"DB_PORT",
		"YELP_API",
		"ENVIROMENT",
	}

	for _, v := range variables {
		if value := os.Getenv(v); value == "" {
			return fmt.Errorf("enviroment variable %s is required", v)
		}
	}

	return nil
}