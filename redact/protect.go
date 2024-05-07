package redact

import (
	"encoding/base64"
	"slices"
	"strings"
)

var globalRedactor = NewRedactor(nil)
var MASKED = "[MASKED]"

type Redactor struct {
	phrases   []string
	maskValue string
}

func (r *Redactor) ProtectPhrase(phrase string) {
	if phrase == "" {
		return
	}

	if !slices.Contains(r.phrases, phrase) {
		r.phrases = append(r.phrases, phrase)

		// add base64 decoded version, if the phrase is base64 encoded
		if isBase64(phrase) {
			// add base64 decoded version, if the phrase is base64 encoded
			decodedValue, _ := base64.StdEncoding.DecodeString(phrase)
			r.phrases = append(r.phrases, string(decodedValue))
		}

		// add base64 encoded version of the phrase
		phraseBase64 := base64.StdEncoding.EncodeToString([]byte(phrase))
		r.phrases = append(r.phrases, phraseBase64)
	}
}

func (r *Redactor) Redact(input string) string {
	for _, phrase := range r.phrases {
		input = strings.ReplaceAll(input, phrase, MASKED)
	}

	return input
}

func (r *Redactor) ProtectedPhraseCount() int {
	return len(r.phrases)
}

func (r *Redactor) SetMaskValue(maskValue string) {
	r.maskValue = maskValue
}

func (r *Redactor) Reset() {
	r.phrases = nil
	r.maskValue = MASKED
}

func NewRedactor(phrases []string) *Redactor {
	return &Redactor{phrases: phrases, maskValue: MASKED}
}

// ProtectPhrase will cause the provided phrase to be redacted (also base64 encoded values)
func ProtectPhrase(phrase string) {
	globalRedactor.ProtectPhrase(phrase)
}

// Redact redacts all protected phrases in the input string (replace with ***)
func Redact(input string) string {
	return globalRedactor.Redact(input)
}
