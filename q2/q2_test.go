package q2

import (
	"github.com/revisao-2/utils"
	"strings"
	"testing"
)

func TestCalculateTotalSalary(t *testing.T) {
	tests := []struct {
		name    string
		emp     *Employee
		want    float64
		wantErr bool
	}{
		{
			name:    "Empregado nil",
			emp:     nil,
			wantErr: true,
		},
		{
			name: "Empregado sem bonus",
			emp: &Employee{
				Name:       "John",
				Title:      "Software Engineer",
				BaseSalary: 1000,
			},
			want: 1000,
		},
		{
			name: "Empregado com bonus abaixo de 1500",
			emp: &Employee{
				Name:       "John",
				Title:      "Software Engineer",
				BaseSalary: 1000,
				Bonuses:    []float64{100, 200, 300},
			},
			want: 1600,
		},
		{
			name: "Empregado com bonus acima de 1500",
			emp: &Employee{
				Name:       "John",
				Title:      "Software Engineer",
				BaseSalary: 1000,
				Bonuses:    []float64{100, 200, 300, 400},
			},
			want: 2000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateTotalSalary(tt.emp)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateTotalSalary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !utils.AssertFloatWithPrecision(tt.want, got, 1e-2) {
				t.Errorf("CalculateTotalSalary() = %v, want %v", got, tt.want)
				return
			}

			if tt.emp != nil && got-tt.emp.BaseSalary > 1500 && !strings.HasPrefix(tt.emp.Title, "Senior") {
				t.Errorf("O título do empregado deveria ser atualizado para Senior %s", tt.emp.Title)
			}
		})
	}
}
