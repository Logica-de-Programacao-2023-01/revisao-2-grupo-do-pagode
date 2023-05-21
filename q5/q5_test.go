package q5

import (
	"testing"
)

func TestCalculateTotalSales(t *testing.T) {
	testCases := []struct {
		name     string
		sales    map[string]*Sale
		expected float64
	}{
		{
			name:     "empty sales",
			sales:    map[string]*Sale{},
			expected: 0,
		},
		{
			name: "one sale",
			sales: map[string]*Sale{
				"1": {
					Code:     "1",
					Products: []*Product{{Price: 10}},
				},
			},
			expected: 10,
		},
		{
			name: "two sales",
			sales: map[string]*Sale{
				"1": {
					Code:     "1",
					Products: []*Product{{Price: 10}},
				},
				"2": {
					Code:     "2",
					Products: []*Product{{Price: 20}, {Price: 30}},
				},
			},
			expected: 60,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CalculateTotalSales(tc.sales)

			if result != tc.expected {
				t.Errorf("CalculateTotalSales(%v) = %v, want %v", tc.sales, result, tc.expected)
			}
		})
	}
}
