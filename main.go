package main

import (
	"fmt"
	"github.com/elahe-dastan/newborn/data"
	"github.com/elahe-dastan/newborn/knn"
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

	//p := &regression.Polynomial{
	//	Bias: 2,
	//	Coefficients: [][]float64{
	//		{3, 2, 1},
	//		{7, 1, 8},
	//		{4, 6, 2}},
	//}

	//p := regression.New()
	//features := [][]float64{
	//	{1, 4},
	//	{2, 6},
	//	{3, 8},
	//	{4, 2},
	//	{5, 9},
	//	{6, 1},
	//	{7, 4},
	//	{8, 4},
	//	{9, 3},
	//	{10, 6},
	//	{11, 7},
	//	{12, 0},
	//}
	//
	//// 3x1 + 2x2^2
	//values := []float64{35, 78, 137, 20, 177, 20, 53, 56, 45, 102, 131, 36}
	//
	//p.Train(features, values, 2, 0.00028, 4000, 0)
	//fmt.Println(p.Bias)
	//fmt.Println(p.Coefficients)

	headers, content := data.ReadCSVData("./knn/dataset_test.csv")
	currentData := make([][]float64, len(content[headers[0]]))
	for i := range currentData {
		f := make([]float64, len(headers)-1)
		for j := range f {
			f[j], _ = strconv.ParseFloat(content[headers[j]][i], 64)
		}
		currentData[i] = f
	}

	labels := make([]int, len(content[headers[0]]))
	for i := range labels {
		labels[i], _ = strconv.Atoi(content[headers[len(headers)-1]][i])
	}

	newData := []float64{57.0,1.0,4.0,140.0,192.0,0.0,0.0,148.0,0.0,0.4,2.0,0.0,6.0}

	label := knn.KNN(currentData, labels, newData, 3)
	fmt.Println(label)
}
