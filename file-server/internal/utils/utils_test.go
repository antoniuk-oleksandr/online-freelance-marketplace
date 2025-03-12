package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlural(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, ""},
		{2, "s"},
		{0, "s"},
		{10, "s"},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", index+1), func(t *testing.T) {
			actual := Plural(test.input)
			assert.Equal(t, test.expected, actual, "Actual value is not equal to expected")
		})
	}
}

func TestCreateDirIfDoesNotExist(t *testing.T) {
	tempDir := t.TempDir()

	CreateDirIfDoesNotExist(tempDir)
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		t.Errorf("Directory was not created")
	}
}
