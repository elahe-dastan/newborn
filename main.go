package main

import (
	"fmt"
	"github.com/elahe-dastan/newborn/regression"
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

	//headers, content := data.ReadCSVData("./regression/dataset_test.csv")
	//features := make([][]float64, len(content[headers[0]]))
	//for i := range features {
	//	f := make([]float64, len(headers)-1)
	//	for j := range f {
	//		f[j], _ = strconv.ParseFloat(content[headers[j]][i], 64)
	//	}
	//	features[i] = f
	//}
	//
	//values := make([]float64, len(content[headers[0]]))
	//for i := range values {
	//	values[i],_ = strconv.ParseFloat(content[headers[len(headers) - 1]][i], 64)
	//}

	//p := regression.New()
	p := &regression.Polynomial{
		Bias: 2,
		Coefficients: [][]float64{
			{3, 2, 1},
			{7, 1, 8},
			{4, 6, 2}},
	}

	features := [][]float64{
		{4, 8, 1},
		{5, 4, 7},
		{8, 1, 4}}

	values := []float64{2, 6, 3}

	p.Train(features, values, 3, 0.5, 1000, 2)
	fmt.Println(p.Bias)
	fmt.Println(p.Coefficients)
}
