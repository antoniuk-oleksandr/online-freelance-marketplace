package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	actual   string
	expected string
}

func TestEncodeString(t *testing.T) {
	tests := []Test{
		{"hello", "aGVsbG8="},
		{"world", "d29ybGQ="},
		{"", ""},
		{"a", "YQ=="},
		{"12345", "MTIzNDU="},
	}

	for _, test := range tests {
		actual := EncodeString(test.actual)
		assert.Equal(t, test.expected, actual, "Expected and actual should be equal")
	}
}

func TestDecodeString(t *testing.T) {
	tests := []Test{
		{"aGVsbG8=", "hello"},
		{"d29ybGQ=", "world"},
		{"", ""},
		{"YQ==", "a"},
		{"MTIzNDU=", "12345"},
	}

	for _, test := range tests {
		actual, err := DecodeString(test.actual)
		assert.NoError(t, err, "Decoding should not return an error")
		assert.Equal(t, test.expected, actual, "Expected and actual should be equal")
	}
}
