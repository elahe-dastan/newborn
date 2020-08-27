package knn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEuclideanDistance(t *testing.T) {
	first := []float64{3, -5, 2}
	second := []float64{7, 4, 9}

	d := euclideanDistance(first, second)

	assert.Equal(t, 12.083045973594572, d)
}

func TestInsertionSort(t *testing.T) {
	distances := []float64{4.6, 3.6, 1.9, 8.4}
	labels := []int{3, 2, 1, 4}

	insertionsSort(distances, labels)

	assert.Equal(t, []int{1, 2, 3, 4}, labels)
}

func TestMode(t *testing.T) {
	labels := []int{1, 2, 1, 3, 1, 2}

	assert.Equal(t, 1, mode(labels))
}

func TestKNN(t *testing.T) {
	currentData := [][]float64{
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 1},
		{3, 5, 3},
		{8, 5, 1},
	}

	labels := []int{1, 4, 1, 7, 4}

	newData := []float64{0, 0, 0}

	res := knn(currentData, labels, newData, 3)

	assert.Equal(t, 1, res)
}
