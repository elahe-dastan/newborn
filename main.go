package main

import (
	"github.com/elahe-dastan/newborn/data"
	"github.com/elahe-dastan/newborn/regression"
	"strconv"
)

func main() {
	//headers, content := data.ReadCSVData("./data/dataset_test.csv")
	//fmt.Println(headers)
	//
	//x := make([]float64, len(content[headers[0]]))
	//y := make([]float64, len(content[headers[1]]))
	//
	//for i := 0; i < len(content[headers[0]]); i++ {
	//	x[i], _ = strconv.ParseFloat(content[headers[0]][i], 64)
	//	y[i], _ = strconv.ParseFloat(content[headers[1]][i], 64)
	//}
	//
	//data.ScatterPlot(x, y, headers[0], headers[1], "example")

	headers, content := data.ReadCSVData("./regression/dataset_test.csv")
	features := make([][]float64, len(content[headers[0]]))
	for i := range features {
		f := make([]float64, len(headers)-1)
		for j := range f {
			f[j], _ = strconv.ParseFloat(content[headers[j]][i], 64)
		}
		features[i] = f
	}

	values := make([]float64, len(content[headers[0]]))
	for i := range values {
		values[i],_ = strconv.ParseFloat(content[headers[len(headers) - 1]][i], 64)
	}

	p := regression.New()
	p.Train(features, values, 3, 0.5, 1000, 0.1)
}
