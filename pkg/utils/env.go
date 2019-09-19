package utils

import "os"

// GetEnvironment
// Get a value from an environment variable and return value if
// no value exists
func GetEnvironment(key, value string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return value
	}

	return val
}
