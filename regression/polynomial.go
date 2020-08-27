package regression

import (
	"math"
	"math/rand"
	"time"
)

type Polynomial struct {
	Bias         float64
	Coefficients [][]float64
}

func New() *Polynomial {
	return &Polynomial{}
}

// The model starts to learn
func (p *Polynomial) Train(features [][]float64, values []float64, degree int, learningRate float64, steps int, lambda float64) {
	if len(features) != len(values) {
		panic("number of data and number of values should be the same")
	}

	// length of features[0] shows the number of variables
	p.setRandomVariables(degree, len(features[0]))

	for i := 0; i < steps; i++ {
		p.gradientDescent(features, values, learningRate, lambda)
	}

}

func (p *Polynomial) derivativeOfCostFunction(features [][]float64, values []float64, lambda float64) (float64, [][]float64) {
	m := len(features)

	sigma := float64(0)
	for i := 0; i < m; i++ {
		sigma += p.valueOfCurve(features[i]) - values[i]
	}

	d_bias := float64(1) / float64(m) * sigma

	derivatives := make([][]float64, len(p.Coefficients))

	for k, _ := range derivatives {
		d := make([]float64, len(p.Coefficients[0]))
		for j := range d {
			sigma := float64(0)
			for i := 0; i < m; i++ {
				sigma += (p.valueOfCurve(features[i]) - values[i]) * math.Pow(features[i][j], float64(k+1))
			}
			d[j] = float64(1) / float64(m) * (sigma + lambda * p.Coefficients[k][j])
		}
		derivatives[k] = d
	}

	return d_bias, derivatives
}

func (p *Polynomial) valueOfCurve(features []float64) float64 {
	v := p.Bias

	for i, c := range p.Coefficients {
		for j, f := range features {
			v += c[j] * math.Pow(f, float64(i+1))
		}
	}

	return v
}

func (p *Polynomial) setRandomVariables(degree int, numOfVariables int) {
	// bias is set to a random number between -10 and 10
	rand.Seed(time.Now().UnixNano())
	bias := float64(rand.Intn(20) - 10)

	coefficients := make([][]float64, degree)
	for i := range coefficients {
		c := make([]float64, numOfVariables)
		for j := range c {
			rand.Seed(time.Now().UnixNano())
			c[j] = float64(rand.Intn(20) - 10)
		}
		coefficients[i] = c
	}

	p.Bias = bias
	p.Coefficients = coefficients
}

func (p *Polynomial) gradientDescent(features [][]float64, values []float64, learningRate float64, lambda float64) {
	d_bias, d_coefficients := p.derivativeOfCostFunction(features, values, lambda)
	p.Bias -= learningRate * d_bias

	for m, c := range p.Coefficients {
		for n := range c {
			p.Coefficients[m][n] -= learningRate * d_coefficients[m][n]
		}
	}
}
