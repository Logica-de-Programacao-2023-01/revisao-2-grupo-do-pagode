package q4

import (
	"reflect"
	"testing"
)

func TestUpdateAverage(t *testing.T) {
	tests := []struct {
		name             string
		students         map[int]*Student
		expectedStudents map[int]*Student
	}{
		{
			name: "case 1",
			students: map[int]*Student{
				1: {
					ID:     1,
					Name:   "John",
					Grades: []float64{10, 5},
				},
				2: {
					ID:     2,
					Name:   "Alice",
					Grades: []float64{10, 20, 30},
				},
			},
			expectedStudents: map[int]*Student{
				1: {
					ID:      1,
					Name:    "John",
					Grades:  []float64{10, 5},
					Average: 7.5,
				},
				2: {
					ID:      2,
					Name:    "Alice",
					Grades:  []float64{10, 20, 30},
					Average: 20,
				},
			},
		},
		{
			name: "case 2",
			students: map[int]*Student{
				1: {
					ID:     1,
					Name:   "John",
					Grades: []float64{10, 5, 15},
				},
				2: {
					ID:     2,
					Name:   "Alice",
					Grades: []float64{10, 20},
				},
			},
			expectedStudents: map[int]*Student{
				1: {
					ID:      1,
					Name:    "John",
					Grades:  []float64{10, 5, 15},
					Average: 10,
				},
				2: {
					ID:      2,
					Name:    "Alice",
					Grades:  []float64{10, 20},
					Average: 15,
				},
			},
		},
		{
			name: "case 3",
			students: map[int]*Student{
				1: {
					ID:     1,
					Name:   "John",
					Grades: []float64{10},
				},
				2: {
					ID:     2,
					Name:   "Alice",
					Grades: []float64{10, 20, 30, 40, 50},
				},
			},
			expectedStudents: map[int]*Student{
				1: {
					ID:      1,
					Name:    "John",
					Grades:  []float64{10},
					Average: 10,
				},
				2: {
					ID:      2,
					Name:    "Alice",
					Grades:  []float64{10, 20, 30, 40, 50},
					Average: 30,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateAverage(tt.students)
			if !reflect.DeepEqual(tt.students, tt.expectedStudents) {
				for i, student := range tt.students {
					if student.Average != tt.expectedStudents[i].Average {
						t.Errorf("student %d got = %v, want %v", i, student.Average, tt.expectedStudents[i].Average)
					}
				}
			}
		})
	}
}
