package mean

// Calculate the mean of a slice of float64 numbers
func Mean(data []float64) float64 {
	var c float64
	var total float64
	ndata := len(data) - 4
	if ndata < 0 {
		return 0.0
	}
	for _, value := range data[ndata:] {
		total += value
		c++
	}
	return total / c
}
