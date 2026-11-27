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

// Multiply returns the product of a and b.
func Multiply(a, b int) int {
	return a * b
}
