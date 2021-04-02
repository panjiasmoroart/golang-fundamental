package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	fmt.Println(sum(scores))

	fmt.Println("======================")

	result, err := calculate(10, 2, "+")
	// jika ada error, artinya nilnya tidak kosong
	if err != nil {
		fmt.Println("Tidak dapat melakukan operasi matematika")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

// function dengan paramter sebuah array
func sum(datas []int) int {
	total := 0
	for _, value := range datas {
		total = total + value
	}
	return total
}

// mengembalikan 2 return value, hasil dan error
func calculate(number1, number2 int, simbol string) (int, error) {

	var total int
	var errorResult error

	switch simbol {

	case "+":
		total = number1 + number2

	case "-":
		total = number1 - number2

	case "*":
		total = number1 * number2

	case "/":
		total = number1 / number2

	default:
		errorResult = errors.New("error calculate")
	}
	return total, errorResult
}
