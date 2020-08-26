package main

import (
	"fmt"
	"strconv"

	"github.com/elahe-dastan/newborn/data"
)

func main() {
	headers, content := data.ReadCSVData("./data/dataset_test.csv")
	fmt.Println(headers)

	x := make([]float64, len(content[headers[0]]))
	y := make([]float64, len(content[headers[1]]))

	for i := 0; i < len(content[headers[0]]); i++ {
		x[i], _ = strconv.ParseFloat(content[headers[0]][i], 64)
		y[i], _ = strconv.ParseFloat(content[headers[1]][i], 64)
	}

	data.ScatterPlot(x, y, headers[0], headers[1], "example")
}
