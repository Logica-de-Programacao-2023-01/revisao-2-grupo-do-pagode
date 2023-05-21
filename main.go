package main

import (
	"fmt"
	"github.com/revisao-2/q1"
)

func main() {
	studentData1 := map[string]q1.Student{
		"John": q1.Student{
			Name: "John",
			Age:  20,
			Subjects: map[string]float64{
				"Math":    8.5,
				"Science": 7.8,
			},
		},
		"Alice": q1.Student{
			Name: "Alice",
			Age:  22,
			Subjects: map[string]float64{
				"History": 9.2,
			},
		},
	}

	studentData2 := map[string]q1.Student{
		"John": q1.Student{
			Name: "John",
			Age:  20,
			Subjects: map[string]float64{
				"Science": 8.0,
				"English": 7.5,
			},
		},
		"Bob": q1.Student{
			Name: "Bob",
			Age:  21,
			Subjects: map[string]float64{
				"Math":    6.7,
				"Physics": 8.9,
			},
		},
	}
	result := q1.MergeStudentData(studentData1, studentData2)
	fmt.Println(result)
}
