package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadCSVData(t *testing.T) {
	headers, data := ReadCSVData("./dataset_test.csv")

	assert.Equal(t, []string{"x", "y"}, headers)
	assert.Equal(t, 20, len(data))
}
