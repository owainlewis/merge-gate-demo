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

func TestClamp(t *testing.T) {
	tests := []struct {
		name           string
		value          int
		minVal, maxVal int
		want           int
	}{
		{name: "below min", value: -5, minVal: 0, maxVal: 10, want: 0},
		{name: "at min", value: 0, minVal: 0, maxVal: 10, want: 0},
		{name: "within bounds", value: 5, minVal: 0, maxVal: 10, want: 5},
		{name: "at max", value: 10, minVal: 0, maxVal: 10, want: 10},
		{name: "above max", value: 15, minVal: 0, maxVal: 10, want: 10},
		{name: "negative range within bounds", value: -3, minVal: -10, maxVal: -1, want: -3},
		{name: "equal bounds value below", value: 2, minVal: 5, maxVal: 5, want: 5},
		{name: "equal bounds value equal", value: 5, minVal: 5, maxVal: 5, want: 5},
		{name: "equal bounds value above", value: 8, minVal: 5, maxVal: 5, want: 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Clamp(tc.value, tc.minVal, tc.maxVal); got != tc.want {
				t.Fatalf("Clamp(%d, %d, %d) = %d, want %d", tc.value, tc.minVal, tc.maxVal, got, tc.want)
			}
		})
	}
}
