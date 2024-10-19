package mean

// Calculate the mean of a slice of float64 numbers
func Mean(data []float64) float64 {
	var c float64
	var total float64
	for _, value := range data {
		total += value
		c++
	}
	return total / c
}
