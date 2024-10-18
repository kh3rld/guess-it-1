package variance

// Calculate the variance of a slice of float64 numbers
func Variance(data []float64, m float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var total float64
	for _, value := range data {
		diff := value - m
		total += diff * diff
	}
	return total / float64(len(data))
}
