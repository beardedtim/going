package env

import (
	"errors"
	"fmt"
	"os"

	logger "mkc-p/modi/log"
)

func GetEnvVar(key string, args ...any) (string, error) {
	value, present := os.LookupEnv(key)

	if present {
		return value, nil
	}

	if len(args) != 0 {
		var fallback string
		fallback = args[0].(string)

		return fallback, nil
	}

	return "", errors.New(fmt.Sprintf(" Cannot reach key \"%s\" and no fallback given as second argumenty", key))
}

func GetEnv(key string, args ...any) string {
	log := logger.CreateLogger("ENV")

	value, err := GetEnvVar(key, args...)

	if err != nil {
		log.Warn("Environemnt not found and no fallback given but using wrapper so swalling error. This may result in runtime issues")
	}

	return value
}
