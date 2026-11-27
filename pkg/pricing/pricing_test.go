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

func TestCatalogProPrice(t *testing.T) {
	pro, ok := Catalog["pro"]
	if !ok {
		t.Fatalf("Catalog missing pro plan")
	}
	if pro.MonthlyCents != 3900 {
		t.Fatalf("Catalog[\"pro\"].MonthlyCents = %d, want 3900", pro.MonthlyCents)
	}
	if got := FormatPrice(pro.MonthlyCents); got != "$39.00" {
		t.Fatalf("FormatPrice(pro.MonthlyCents) = %q, want $39.00", got)
	}
}

func TestApplyAnnualDiscount(t *testing.T) {
	tests := []struct {
		name         string
		monthlyCents int
		want         int
	}{
		{name: "zero", monthlyCents: 0, want: 0},
		{name: "negative", monthlyCents: -1000, want: 0},
		{name: "pro plan", monthlyCents: 3900, want: 3120},
		{name: "enterprise plan", monthlyCents: 19900, want: 15920},
		{name: "small exact amount", monthlyCents: 500, want: 400},
		{name: "truncates fractional cents", monthlyCents: 101, want: 80},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ApplyAnnualDiscount(tc.monthlyCents); got != tc.want {
				t.Fatalf("ApplyAnnualDiscount(%d) = %d, want %d", tc.monthlyCents, got, tc.want)
			}
		})
	}
}
