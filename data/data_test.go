package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadCSVData(t *testing.T) {
	headers, data := ReadCSVData("./dataset_test.csv")

	assert.Equal(t, []string{"x", "y"}, headers)
	assert.Equal(t, []string{"0.010", "0.014","0.018","0.022","0.027","0.031","0.035","0.039","0.043","0.047"}, data[headers[0]])
}
