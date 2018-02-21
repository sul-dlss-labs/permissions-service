package config

import (
	"log"
	"os"
	"strconv"
)

// Config is configuration for the application
type Config struct {
	Port int
}

// NewConfig creates a new configuration with values from environment variables
// or defaults
func NewConfig() *Config {
	return &Config{
		Port: getInteger("APP_PORT", 8080),
	}
}

func getString(envVar string, defaultValue string) string {
	var value string
	value = os.Getenv(envVar)
	if value == "" {
		value = defaultValue
		log.Printf("%s: Using default [%s].", envVar, defaultValue)
		return defaultValue
	}
	log.Printf("%s: Found setting [%s].", envVar, value)
	return value
}

func getBool(envVar string, defaultValue bool) bool {
	var value string
	value = os.Getenv(envVar)
	if value == "FALSE" || value == "false" {
		log.Printf("%s: Using default [%s].", envVar, value)
		return false
	}
	log.Printf("%s: Found setting [true].", envVar)
	return true
}

func getInteger(envVar string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(envVar))
	if err != nil || value == 0 {
		log.Printf("%s: Using default [%v].", envVar, defaultValue)
		return defaultValue
	}
	log.Printf("%s: Found setting [%v].", envVar, value)
	return value
}
