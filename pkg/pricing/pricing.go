package pricing

import "fmt"

// Plan represents a subscription tier billed monthly in whole cents.
type Plan struct {
	ID           string
	Name         string
	MonthlyCents int
}

var Catalog = map[string]Plan{
	"free":       {ID: "free", Name: "Free", MonthlyCents: 0},
	"pro":        {ID: "pro", Name: "Pro", MonthlyCents: 3900},
	"enterprise": {ID: "enterprise", Name: "Enterprise", MonthlyCents: 19900},
}

// FormatPrice renders whole cents as a USD price string.
func FormatPrice(cents int) string {
	return fmt.Sprintf("$%d.%02d", cents/100, cents%100)
}

// ApplyAnnualDiscount returns the monthly price in cents after applying a 20%
// discount for annual billing.
func ApplyAnnualDiscount(monthlyCents int) int {
	if monthlyCents <= 0 {
		return 0
	}
	return (monthlyCents * 80) / 100
}

// ProrateCents computes the remaining charge for a partial billing period.
func ProrateCents(monthlyCents, remainingDays, totalDays int) int {
	if monthlyCents <= 0 || remainingDays <= 0 || totalDays <= 0 {
		return 0
	}
	if remainingDays >= totalDays {
		return monthlyCents
	}
	return (monthlyCents * remainingDays) / totalDays
}
