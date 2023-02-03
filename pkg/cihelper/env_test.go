package cihelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertEnvMapToStringSlice(t *testing.T) {
	tests := []struct {
		envMap           map[string]string
		expectedEnvLines []string
	}{
		{map[string]string{"key1": "value1", "key2": "value2"}, []string{"key1=value1", "key2=value2"}},
		{map[string]string{}, []string(nil)},
	}

	for _, test := range tests {
		envLines := ConvertEnvMapToStringSlice(test.envMap)

		assert.Len(t, envLines, len(test.expectedEnvLines))
		for _, line := range envLines {
			assert.Contains(t, test.expectedEnvLines, line)
		}
	}
}
