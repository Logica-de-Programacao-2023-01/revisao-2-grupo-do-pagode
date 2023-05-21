package q1

import (
	"reflect"
	"testing"
)

func TestMergeStudentData(t *testing.T) {
	tests := []struct {
		name         string
		studentData1 map[string]Student
		studentData2 map[string]Student
		expected     map[string]Student
	}{
		{
			name: "Alunos em ambos os conjuntos de dados",
			studentData1: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.8,
					},
				},
				"Alice": {
					Name: "Alice",
					Age:  22,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.0,
					},
				},
			},
			studentData2: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Science": 8.0,
						"English": 7.5,
					},
				},
				"Alice": {
					Name: "Alice",
					Age:  22,
					Subjects: map[string]float64{
						"Math":    9.5,
						"Science": 6.0,
					},
				},
			},
			expected: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 8.0,
						"English": 7.5,
					},
				},
				"Alice": {
					Name: "Alice",
					Age:  22,
					Subjects: map[string]float64{
						"Math":    9.5,
						"Science": 6.0,
					},
				},
			},
		},
		{
			name: "Nenhum aluno em comum",
			studentData1: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.8,
					},
				},
			},
			studentData2: map[string]Student{
				"Alice": {
					Name: "Alice",
					Age:  22,
					Subjects: map[string]float64{
						"Math":    9.5,
						"Science": 6.0,
					},
				},
			},
			expected: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.8,
					},
				},
				"Alice": {
					Name: "Alice",
					Age:  22,
					Subjects: map[string]float64{
						"Math":    9.5,
						"Science": 6.0,
					},
				},
			},
		},
		{
			name: "Alunos com nomes iguais mas matérias diferentes",
			studentData1: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.8,
					},
				},
			},
			studentData2: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"English": 9.5,
						"History": 6.0,
					},
				},
			},
			expected: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.8,
						"English": 9.5,
						"History": 6.0,
					},
				},
			},
		},
		{
			name:         "Mapas vazios",
			studentData1: map[string]Student{},
			studentData2: map[string]Student{},
			expected:     map[string]Student{},
		},
		{
			name:         "Mapa 1 vazio",
			studentData1: map[string]Student{},
			studentData2: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.8,
					},
				},
			},
			expected: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.8,
					},
				},
			},
		},
		{
			name: "Mapa 2 vazio",
			studentData1: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.8,
					},
				},
			},
			studentData2: map[string]Student{},
			expected: map[string]Student{
				"John": {
					Name: "John",
					Age:  20,
					Subjects: map[string]float64{
						"Math":    8.5,
						"Science": 7.8,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MergeStudentData(test.studentData1, test.studentData2)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
