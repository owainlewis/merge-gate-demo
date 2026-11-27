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
