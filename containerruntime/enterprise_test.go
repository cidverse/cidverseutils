package containerruntime

import (
	"testing"
)

func TestToJavaNoProxyFormat(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "empty string",
			input:          "",
			expectedOutput: "",
		},
		{
			name:           "single value",
			input:          "localhost",
			expectedOutput: "localhost",
		},
		{
			name:           "multiple values",
			input:          "localhost,127.0.0.1,*.example.com",
			expectedOutput: "localhost|127.0.0.1|*.example.com",
		},
		{
			name:           "no commas",
			input:          "localhost 127.0.0.1 *.example.com",
			expectedOutput: "localhost 127.0.0.1 *.example.com",
		},
		{
			name:           "wildcards with multiple levels",
			input:          "*.example.com,foo.*.example.com",
			expectedOutput: "*.example.com|foo.*.example.com",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output := ToJavaNoProxyFormat(testCase.input)
			if output != testCase.expectedOutput {
				t.Errorf("ToJavaNoProxyFormat() = %q, want %q", output, testCase.expectedOutput)
			}
		})
	}
}
