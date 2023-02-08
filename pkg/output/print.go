package output

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func PrintStruct(t *testing.T, result interface{}) {
	jsonByteArray, jsonErr := json.MarshalIndent(result, "", "\t")
	assert.NoError(t, jsonErr)
	fmt.Println(string(jsonByteArray))
}
