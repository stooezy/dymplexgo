package util

import (
	"os"
	"strconv"
)

func GetEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
	strVal := GetEnv(key, "")
	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}
	return defaultValue
}

func GetEnvAsBool(key string, defaultValue bool) bool {
	strVal := GetEnv(key, "")
	if val, err := strconv.ParseBool(strVal); err == nil {
		return val
	}
	return defaultValue
}
