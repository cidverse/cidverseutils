package ci

import (
	"strings"
)

func EnvMapToStringSlice(env map[string]string) []string {
	fullEnv := make(map[string]string, len(env))
	for k, v := range env {
		fullEnv[k] = v
	}

	// convert
	var envLines []string
	for k, v := range fullEnv {
		envLines = append(envLines, k+"="+v)
	}

	return envLines
}

func ToEnvName(input string) string {
	input = strings.ToUpper(input)
	input = strings.Replace(input, ".", "_", -1)
	input = strings.Replace(input, "-", "_", -1)

	return input
}

// GetMapValueOrDefault returns a value in a map or a default value
func GetMapValueOrDefault(entity map[string]string, key string, defaultValue string) (val string) {
	value, found := entity[key]

	if found {
		return value
	}

	return defaultValue
}
