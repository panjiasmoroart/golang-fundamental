package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

// method yg receivernya pointer
func (student *Student) graduateMethod() {
	student.Name = student.Name + " S.KOM"
	// fmt.Println(student.Name)
}

func main() {
	student := Student{1, "Panji Asmoro", 3.61}

	// fmt.Println(student.Name)
	// // call function
	// graduate(&student)

	// fmt.Println(student.Name)

	fmt.Println("==============Method with pointer=============")
	fmt.Println(student.Name)
	// call method
	student.graduateMethod()

	fmt.Println(student.Name)
}

// func graduate(student *Student) {
// 	student.Name = student.Name + " S.KOM"
// 	// fmt.Println(student.Name)
// }
