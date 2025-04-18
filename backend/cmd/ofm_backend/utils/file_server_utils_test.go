package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Images []string
}

type NestedTestData struct {
	Data TestData
}

type ComplexData struct {
	NestedImages [][]string
}

func TestAddServerURLToFiles(t *testing.T) {
	host := "http://localhost:8030"
	os.Setenv("FILE_SERVER_HOST", host)

	defer os.Unsetenv("FILE_SERVER_HOST")

	data := TestData{
		Images: []string{"image1.jpg", "image2.jpg"},
	}

	expectedData := TestData{
		Images: []string{
			fmt.Sprintf("%s/%s", host, "image1.jpg"),
			fmt.Sprintf("%s/%s", host, "image2.jpg"),
		},
	}

	actualData := AddServerURLToFiles(&data)

	assert.Equal(t, expectedData, *actualData, "Nested file URLs were not updated correctly")
}

func TestAddServerURLToFilesNestedData(t *testing.T) {
	host := "http://localhost:8030"
	os.Setenv("FILE_SERVER_HOST", host)

	defer os.Unsetenv("FILE_SERVER_HOST")

	data := NestedTestData{
		Data: TestData{
			Images: []string{"image1.jpg", "image2.jpg"},
		},
	}

	expectedData := NestedTestData{
		Data: TestData{
			Images: []string{
				fmt.Sprintf("%s/%s", host, "image1.jpg"),
				fmt.Sprintf("%s/%s", host, "image2.jpg"),
			},
		},
	}

	actualData := AddServerURLToFiles(&data)

	assert.Equal(t, expectedData, *actualData, "Nested file URLs were not updated correctly")
}

func TestAddServerURLToFilesNonFileString(t *testing.T) {
	host := "http://localhost:8030"
	os.Setenv("FILE_SERVER_HOST", host)

	defer os.Unsetenv("FILE_SERVER_HOST")

	data := TestData{
		Images: []string{"not_a_file", "another_string"},
	}

	expectedData := TestData{
		Images: []string{"not_a_file", "another_string"},
	}

	actualData := AddServerURLToFiles(&data)

	assert.Equal(t, expectedData, *actualData, "Non-file strings should not be modified")
}

func TestAddServerURLToFilesEmptySlice(t *testing.T) {
	host := "http://localhost:8030"
	os.Setenv("FILE_SERVER_HOST", host)

	defer os.Unsetenv("FILE_SERVER_HOST")

	data := TestData{
		Images: []string{},
	}

	expectedData := TestData{
		Images: []string{},
	}

	actualData := AddServerURLToFiles(&data)

	assert.Equal(t, expectedData, *actualData, "Empty slices should be handled correctly")
}

func TestAddServerURLToFilesNestedSlices(t *testing.T) {
	host := "http://localhost:8030"
	os.Setenv("FILE_SERVER_HOST", host)

	defer os.Unsetenv("FILE_SERVER_HOST")

	data := ComplexData{
		NestedImages: [][]string{
			{"image1.jpg", "image2.jpg"},
			{"image3.jpg", "not_a_file"},
		},
	}

	expectedData := ComplexData{
		NestedImages: [][]string{
			{
				fmt.Sprintf("%s/%s", host, "image1.jpg"),
				fmt.Sprintf("%s/%s", host, "image2.jpg"),
			},
			{
				fmt.Sprintf("%s/%s", host, "image3.jpg"),
				"not_a_file",
			},
		},
	}

	actualData := AddServerURLToFiles(&data)

	assert.Equal(t, expectedData, *actualData, "Deeply nested slices should be handled correctly")
}

func TestAddServerURLToFilesMissingEnvVars(t *testing.T) {
	os.Unsetenv("FILE_SERVER_HOST")

	data := TestData{
		Images: []string{"image1.jpg", "image2.jpg"},
	}

	expectedData := data

	actualData := AddServerURLToFiles(&data)

	assert.Equal(t, expectedData, *actualData, "Function should fail gracefully when environment variables are missing")
}

func TestAddServerURLToFilesNilInput(t *testing.T) {
	var data *TestData

	actualData := AddServerURLToFiles(data)

	assert.Nil(t, actualData, "Nil input should return nil")
}
