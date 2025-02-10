package ci

import (
	"strings"
)

func EnvMapToStringSlice(env map[string]string) []string {
	envLines := make([]string, 0, len(env))
	for k, v := range env {
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
