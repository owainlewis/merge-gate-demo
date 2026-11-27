package stats

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
	if got := Add(-4, 4); got != 0 {
		t.Fatalf("Add(-4, 4) = %d, want 0", got)
	}
}

func TestAverage(t *testing.T) {
	if got := Average(nil); got != 0 {
		t.Fatalf("Average(nil) = %v, want 0", got)
	}
	if got := Average([]float64{2, 4, 6}); got != 4 {
		t.Fatalf("Average([2,4,6]) = %v, want 4", got)
	}
}

func TestMultiply(t *testing.T) {
	if got := Multiply(6, 7); got != 42 {
		t.Fatalf("Multiply(6, 7) = %d, want 42", got)
	}
	if got := Multiply(-3, 5); got != -15 {
		t.Fatalf("Multiply(-3, 5) = %d, want -15", got)
	}
}
