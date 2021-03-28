package main

import "fmt"

func main() {
	// slice & map
	students := []map[string]string{
		{"name": "Panji", "score": "A"},
		{"name": "Yumna", "score": "B"},
		{"name": "Tarisa", "score": "C"},
	}

	// fmt.Println(students)
	for _, student := range students {
		fmt.Println(student["name"], " - ", student["score"])
	}
}
