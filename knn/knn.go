package knn

import "math"

func knn(currentData [][]float64, labels []int, newData []float64, k int) int {
	if len(currentData) != len(labels) {
		panic("length of current data and labels should be the same")
	}

	distances := make([]float64, len(currentData))

	for i := range distances {
		distance := euclideanDistance(currentData[i], newData)
		distances[i] = distance
	}

	insertionsSort(distances, labels)

	firstK := labels[:k]

	return mode(firstK)
}

func euclideanDistance(first []float64, second []float64) float64 {
	if len(first) != len(second) {
		panic("number of features of data should be the same")
	}

	sumSquaredDistance := float64(0)

	for i := range first {
		sumSquaredDistance += math.Pow(first[i]-second[i], 2)
	}

	return math.Sqrt(sumSquaredDistance)
}

func insertionsSort(distances []float64, labels []int) {
	L := len(distances)

	for i := 0; i < L; i++ {
		j := i
		for j > 0 && distances[j] < distances[j-1] {
			distances[j], distances[j-1] = distances[j-1], distances[j]
			labels[j], labels[j-1] = labels[j-1], labels[j]
			j -= 1
		}
	}
}

func mode(labels []int) int {
	m := make(map[int]int)

	for _, label := range labels {
		m[label]++
	}

	max := 0
	res := 0

	for k, v := range m {
		if v > max {
			max = v
			res = k
		}
	}

	return res
}