package pricing

import "testing"

func TestFormatPrice(t *testing.T) {
	if got := FormatPrice(2900); got != "$29.00" {
		t.Fatalf("FormatPrice(2900) = %q, want $29.00", got)
	}
	if got := FormatPrice(505); got != "$5.05" {
		t.Fatalf("FormatPrice(505) = %q, want $5.05", got)
	}
}

func TestProrateCents(t *testing.T) {
	if got := ProrateCents(3000, 15, 30); got != 1500 {
		t.Fatalf("ProrateCents(3000, 15, 30) = %d, want 1500", got)
	}
	if got := ProrateCents(3000, 0, 30); got != 0 {
		t.Fatalf("ProrateCents(3000, 0, 30) = %d, want 0", got)
	}
}
