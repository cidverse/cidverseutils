package redact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhraseAddition(t *testing.T) {
	globalRedactor.Reset()

	// protect phrase
	ProtectPhrase("bXlzZWNyZXQ=")
	assert.Equal(t, 3, globalRedactor.ProtectedPhraseCount()) // should contain the input text, the base64 decoded text, and the base64 encoded text
	//assert.Equal(t, "bXlzZWNyZXQ=", protectedPhrases[0])
	//assert.Equal(t, "mysecret", protectedPhrases[1])
	//assert.Equal(t, "YlhselpXTnlaWFE9", protectedPhrases[2])
}

func TestRedaction(t *testing.T) {
	globalRedactor.Reset()

	// protect phrase
	ProtectPhrase("mysecret")

	// check redacted
	out := Redact("abc mysecret def")
	assert.Equal(t, "abc [MASKED] def", out)
}

func TestRedactionBase64(t *testing.T) {
	globalRedactor.Reset()

	// protect phrase
	ProtectPhrase("mysecret")

	// check redacted
	out := Redact("abc bXlzZWNyZXQ= def")
	assert.Equal(t, "abc [MASKED] def", out)
}

func TestRedactionBase64Encoded(t *testing.T) {
	globalRedactor.Reset()

	// protect phrase
	ProtectPhrase("bXlzZWNyZXQ=")

	// check redacted
	out := Redact("test mysecret test")
	assert.Equal(t, "test [MASKED] test", out)
}
