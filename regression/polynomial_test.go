package regression

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// gradient descent
// derivative of cost function
// value of curve

func TestPolynomial_valueOfCurve(t *testing.T) {
	p := &Polynomial{
		Bias: 2,
		Coefficients: [][]float64{
			{3, 2, 1},
			{7, 1, 8},
			{4, 6, 2}},
	}

	features := []float64{4, 8, 1}

	// This polynomial regression had degree equal to 3 and also has 3 variables so the formula of the curve is
	// bias + c00 * x0 + c01 * x1 + c02 * x2 + c10 * x0 ^ 2 + c11 * x1 ^ 2 + c12 * x2 ^ 2 + c20 * x0 ^ 3 + c21 * x1 ^ 3 + c22 * x2 ^ 3
	expectedValue := p.Bias +
		p.Coefficients[0][0]*features[0] + p.Coefficients[0][1]*features[1] + p.Coefficients[0][2]*features[2] +
		p.Coefficients[1][0]*math.Pow(features[0], 2) + p.Coefficients[1][1]*math.Pow(features[1], 2) + p.Coefficients[1][2]*math.Pow(features[2], 2) +
		p.Coefficients[2][0]*math.Pow(features[0], 3) + p.Coefficients[2][1]*math.Pow(features[1], 3) + p.Coefficients[2][2]*math.Pow(features[2], 3)

	assert.Equal(t, expectedValue, p.valueOfCurve(features))
}

func TestPolynomial_derivativeOfCostFunction(t *testing.T) {
	p := &Polynomial{
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

	//The cost function is (1/(2*len(data))) * Sigma(valueOfCurve(features(i)) - y(i))^2
	//so the derivative will be d(a)(b)(1/len(data)) * Sigma(valueOfCurve(feature(i)) - y(i)*c(b)^a)
	//According to the bias and coefficients we get this formula for derivation of cost function
	// 1/3 * [((2 + 3 * 4 + 2 * 8 + 1 * 1 + 7 * 16 + 1 * 64 + 8 * 1 + 4 * 64 + 6 * 512 + 2 * 1) - 2) * appropriate x == 3543 * appropriate x
	//        ((2 + 3 * 5 + 2 * 4 + 1 * 7 + 7 * 25 + 1 * 16 + 8 * 49 + 4 * 125 + 6 * 64 + 2 * 343) - 6) * appropriate x^2 == 2179 * appropriate x^2
	//        ((2 + 3 * 8 + 2 * 1 + 1 * 4 + 7 * 64 + 1 * 1 + 8 * 16 + 4 * 512 + 6 * 1 + 2 * 64) - 3) * appropriate x^3] == 2788 * appropriate x^3

	d_bias, d_coefficients := p.derivativeOfCostFunction(features, values, 0)

	assert.Equal(t, 2836.6666666666665, d_bias)
	assert.Equal(t, 652086.6666666666, d_coefficients[2][1])
}
