[![Build Status](https://cloud.drone.io/api/badges/elahe-dastan/newborn/status.svg)](https://cloud.drone.io/elahe-dastan/newborn)
# Introduction
This repository is really simple and small I tried to put only basic and most frequently ML algorithms in it, I'm going
to explain all the algorithms implemented in this repository, how you can use them and even the math calculation behind
them for those who are interesting.

# Reading Data And Picturing It
One of the most basic capabilities you'll need for sure is to read data from a dataset and plot it if possible, most of 
the times the dataset is a CSVDataset here is a sample of how to use this library
```go
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
```
```
[x y]
```
![](images/example.png)

# Regression
One of the basic things everybody has to learn to study ML is linear regression, logistic regression and
nonlinear regression, regression is used in both predicting a value and classification.<br/>
Let's talk a little deeper about regression, we want to fit a line or curve to a set of data.I'm going to show the math
calculations of what is happening.

[check here if you are interested in the math calculations](https://elahe-dastan.github.io/newborn/)

In this repository I haven't implemented linear regression separately because linear regression is polynomial regression 
with degree equal to 1 and lambda equal to 0

# K Nearest Neighbour
This algorithm is used for classification, to choose a label for a new data the algorithm finds the k nearest data to it
and finds the most repeated label among them, different methods can be used to calculate the distance between the new data 
and the old ones in this repository I use one of the easiest approaches called Euclidean Distance.<br/>
```go
package main

import (
	"fmt"
	"github.com/elahe-dastan/newborn/data"
	"github.com/elahe-dastan/newborn/knn"
	"strconv"
)

func main() {
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
```
```
0
```