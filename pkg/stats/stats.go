package stats

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

// Average returns the arithmetic mean of values, or 0 when values is empty.
func Average(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	var total float64
	for _, v := range values {
		total += v
	}
	return total / float64(len(values))
}

// Clamp restricts value to the inclusive range [minVal, maxVal].
func Clamp(value, minVal, maxVal int) int {
	if value < minVal {
		return minVal
	}
	if value > maxVal {
		return maxVal
	}
	return value
}
