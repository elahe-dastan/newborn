package data

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Reads CSV data from the path given and returns the headers and data
func ReadCSVData(path string) ([]string, []string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	dataAsBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
	}

	dataAsText := string(dataAsBytes)

	linesOfData := strings.Split(dataAsText, "\n")

	headers := strings.Split(linesOfData[0], ",")

	return headers, linesOfData[1:]
}
