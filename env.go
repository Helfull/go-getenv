package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var errEnvVarEmpty = errors.New("getenv: environment variable empty")

var loadedVariables map[string]string

// Clears the loaded variables
func ClearDotEnv() {
	loadedVariables = nil
}

// LoadDotEnv loads all environment variables
func LoadDotEnv(fileNames ...string) (map[string]string, error) {

	var env map[string]string

	env, err := godotenv.Read(fileNames...)

	loadedVariables = env

	return env, err
}

// GetEnvStr gets the environment variable of the name
// Environment variables have priority over loaded Variables
// it does return the fallback if the variable name is not found
func GetEnvStr(name string, fallback string) (string, error) {
	var ok bool
	var v string
	v, ok = os.LookupEnv(name)

	if ok {
		return v, nil
	}

	v, ok = loadedVariables[name]

	if ok {
		return v, nil
	}

	return fallback, errEnvVarEmpty
}

// GetEnvInt gets the environment variable of the name
// Environment variables have priority over loaded Variables
// it does return the fallback if the variable name is not found
func GetEnvInt(name string, fallback int) (int, error) {
	s, err := GetEnvStr(name, "")
	if err != nil {
		return fallback, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return fallback, err
	}
	return v, nil
}

// GetEnvBool gets the environment variable of the name
// Environment variables have priority over loaded Variables
// it does return the fallback if the variable name is not found
func GetEnvBool(name string, fallback bool) (bool, error) {
	s, err := GetEnvStr(name, "false")
	if err != nil {
		return fallback, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return fallback, err
	}
	return v, nil
}
