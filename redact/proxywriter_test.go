package redact

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProtectedWriter(t *testing.T) {
	globalRedactor.Reset()
	ProtectPhrase("mySecret")

	writer := NewProtectedWriter(nil, nil, &sync.Mutex{}, globalRedactor)
	_, _ = writer.Write([]byte("this contains a secret: mySecret"))
	assert.Equal(t, "this contains a secret: [MASKED]", lastProxyWrite)
}
