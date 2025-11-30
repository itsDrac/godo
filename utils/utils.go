package utils

import (
	"fmt"
	"os"
)

var lookupEnv = os.LookupEnv

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is not present in the environment, then the defaultValue is returned.
func GetEnv(key, defaultValue string) string {
	if value, exists := lookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetEnvAsInt(name string, defaultVal int) int {
	if valueStr, exists := lookupEnv(name); exists {
		var value int
		_, err := fmt.Sscanf(valueStr, "%d", &value)
		if err == nil {
			return value
		}
	}
	return defaultVal
}

func GetEnvAsBool(name string, defaultVal bool) bool {
	if valueStr, exists := lookupEnv(name); exists {
		switch valueStr {
		case "true", "1":
			return true
		case "false", "0":
			return false
		}
	}
	return defaultVal
}