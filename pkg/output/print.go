package output

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func PrintStruct(t *testing.T, result interface{}) {
	jsonByteArray, jsonErr := json.MarshalIndent(result, "", "\t")
	assert.NoError(t, jsonErr)
	fmt.Println(string(jsonByteArray))
}
