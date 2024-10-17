package mean

// Calculate the mean of a slice of float64 numbers
func Mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var total float64
	for _, value := range data {
		total += value
	}
	return total / float64(len(data))
}
