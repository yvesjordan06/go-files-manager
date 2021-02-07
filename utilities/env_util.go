package utilities

import (
	"errors"
	"fmt"
)

// TODO Implements Enviroment variable methods

/// GetEnvOrDefault tries to get a value from the environment
/// if the value is absent, it returns the default value
func GetEnvOrDefault(key string, defaultValue string) string {
	return defaultValue
}

/// GetEnv tries to get a value from the environment
/// if the value is absent, it returns an error
func GetEnv(key string) (string, error) {
	return "", errors.New(fmt.Sprintf("Cannot find an Environment variable with %s key", key))
}

/// GetEnv tries to get a value from the environment
/// if the value is absent, it panics
func GetEnvStrict(key string) string {
	panic("Not implemented")
}
