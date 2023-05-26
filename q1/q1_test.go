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

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1 map[string]Student, studentData2 map[string]Student) map[string]Student {
	for c1 := range studentData1 {
		for c2 := range studentData2 {
			if c1 == c2 {
				for i1 := range studentData1[c1].Subjects {
					for i2 := range studentData2[c2].Subjects {
						if i1 == i2 {
							studentData1[c1].Subjects[i1] = studentData2[c2].Subjects[i2]
						} else {
							studentData1[c1].Subjects[i2] = studentData2[c2].Subjects[i2]
						}
					}
				}
			}
		}
	}
	return studentData1
}
