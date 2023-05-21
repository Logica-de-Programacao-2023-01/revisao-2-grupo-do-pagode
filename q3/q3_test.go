package q3

import (
	"testing"
)

func TestUpdateStock(t *testing.T) {
	tests := []struct {
		name             string
		product          *Product
		expectedQuantity int
		sales            map[string]int
		wantErr          bool
	}{
		{
			name:    "Product nil",
			product: nil,
			wantErr: true,
		},
		{
			name: "Sales nil",
			product: &Product{
				Code:     "123",
				Name:     "Product 1",
				Price:    19.99,
				Quantity: 10,
			},
			sales:            nil,
			expectedQuantity: 10,
			wantErr:          false,
		},
		{
			name: "Product not found",
			product: &Product{
				Code:     "123",
				Name:     "Product 1",
				Price:    19.99,
				Quantity: 10,
			},
			sales: map[string]int{
				"456": 1,
			},
			expectedQuantity: 10,
		},
		{
			name: "Product found",
			product: &Product{
				Code:     "123",
				Name:     "Product 1",
				Price:    19.99,
				Quantity: 10,
			},
			sales: map[string]int{
				"123": 1,
			},
			expectedQuantity: 9,
		},
		{
			name: "Product found, quantity negative",
			product: &Product{
				Code:     "123",
				Name:     "Product 1",
				Price:    19.99,
				Quantity: 10,
			},
			sales: map[string]int{
				"123": 11,
			},
			expectedQuantity: 10,
			wantErr:          true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := UpdateStock(test.product, test.sales)

			if test.wantErr && err == nil {
				t.Errorf("Erro: esperava um erro, mas não houve erro")
			}

			if test.product != nil && test.product.Quantity != test.expectedQuantity {
				t.Errorf("Erro: esperava %d, mas foi %d", test.expectedQuantity, test.product.Quantity)
			}
		})
	}
}
