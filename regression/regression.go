package regression

import (
	"math"
	"math/rand"
	"time"
)

type Regression struct {
	Bias         float64
	Coefficients [][]float64
}

func New() *Regression {
	return &Regression{}
}

// In the beginning a random value between -10 and 10 is assigned to the coefficients
func Train(features [][]float64, values []float64, degree int, learningRate float64, steps int, lambda float64) {
	if len(features) != len(values) {
		panic("number of data and number of values should be the same")
	}

	rand.Seed(time.Now().UnixNano())
	bias := float64(rand.Intn(20) - 10)

	coefficients := make([][]float64, degree)
	for i := range coefficients {
		c := make([]float64, len(features[0]))
		for j := range c {
			rand.Seed(time.Now().UnixNano())
			c[j] = float64(rand.Intn(20) - 10)
		}
		coefficients[i] = c
	}

	// Gradient Descent
	for i := 0; i < steps; i++ {
		d_bias, d_coefficients := derivativeOfCostFunction(features, values, bias, coefficients, lambda)
		bias -= learningRate * d_bias
		for m := range coefficients {
			for n := range coefficients[0] {
				coefficients[m][n] -= learningRate * d_coefficients[m][n]
			}
		}
	}

}

func derivativeOfCostFunction(features [][]float64, values []float64, bias float64, coefficients [][]float64, lambda float64) (float64, [][]float64) {
	m := len(features)

	sigma := float64(0)
	for i := 0; i < m; i++ {
		sigma += valueOfCurve(features[i], bias, coefficients) - values[i]
	}

	d_bias := float64(1) / float64(m) * sigma

	derivatives := make([][]float64, len(coefficients))

	for k := range derivatives {
		d := make([]float64, len(coefficients[0]))
		for j := range d {
			sigma := float64(0)
			for i := 0; i < m; i++ {
				sigma += (valueOfCurve(features[i], bias, coefficients) - values[i]) * math.Pow(features[i][j], float64(k))
			}
			d[j] = float64(1) / float64(m) * sigma
		}
		derivatives[k] = d
	}

	return d_bias, derivatives
}

func valueOfCurve(features []float64, bias float64, coefficients [][]float64) float64 {
	v := bias

	for i, c := range coefficients {
		for j, f := range features {
			v += c[j] * math.Pow(f, float64(i+1))
		}
	}

	return v
}
