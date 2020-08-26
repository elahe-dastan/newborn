package data

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gonum.org/v1/plot/plotter"
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

// Gets x and y array, plot it and saves it to a png a file with the name given
func ScatterPlot(x []float64, y []float64, name string) {
	if len(x) != len(y) {
		panic("len of x array and y array should be the same")
	}

	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "ML dataset"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(s)
	p.Legend.Add("scatter", s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, name+".png"); err != nil {
		panic(err)
	}
}
