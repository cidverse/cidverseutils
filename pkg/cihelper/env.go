package cihelper

import (
	"strings"
)

func ConvertEnvMapToStringSlice(env map[string]string) []string {
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
