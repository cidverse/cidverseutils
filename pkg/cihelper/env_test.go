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

func TestToEnvName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"foo.bar", "FOO_BAR"},
		{"foo", "FOO"},
		{"FOO.BAR", "FOO_BAR"},
		{"", ""},
		{"FOOBAR", "FOOBAR"},
		{"foo-bar", "FOO_BAR"},
		{"foo_bar", "FOO_BAR"},
	}

	for _, test := range tests {
		result := ToEnvName(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
