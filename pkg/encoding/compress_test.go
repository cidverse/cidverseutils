package encoding

import (
	"testing"
)

func TestGZIPBase64EncodeBytes(t *testing.T) {
	tests := []struct {
		input   []byte
		want    string
		wantErr bool
	}{
		{[]byte("hello world"), "H4sIAAAAAAAA/8pIzcnJVyjPL8pJAQQAAP//hRFKDQsAAAA=", false},
		{[]byte("golang"), "H4sIAAAAAAAA/0rPz0nMSwcEAAD//6MlHK8GAAAA", false},
		{[]byte(""), "H4sIAAAAAAAA/wEAAP//AAAAAAAAAAA=", false},
	}
	for _, test := range tests {
		got, err := GZIPBase64EncodeBytes(test.input)
		if (err != nil) != test.wantErr {
			t.Errorf("GZIPBase64EncodeBytes(%q) error = %v, wantErr %v", test.input, err, test.wantErr)
			continue
		}
		if got != test.want {
			t.Errorf("GZIPBase64EncodeBytes(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
