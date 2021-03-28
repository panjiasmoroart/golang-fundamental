package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Hallo, belajar Golang")
	// karena main.go dan entity.go berada dalam satu package yaitu main
	// jadi kita bisa akses func testAja secara langsung
	// go run main.go entity.go -> cara menjalankannya
	sentence := TestAja()
	fmt.Println(sentence)

	result := calculation.Add(10, 5)
	fmt.Println(result)

	resultMultiply := calculation.Multiply(2, 10)
	fmt.Println(resultMultiply)
}

// Di golang ada 2 type package
// 1. Executable -> langsung di eksekusi, sebuah kode yg ingin lagunsung dieksekusi harus berada didalam pakcage (main)
// 2. Library -> dia package yang ga langsung di eksekusi, tetapi dipanggil bagian file lain

// menggunakan import
// 1. standar Library -> fmt
// 2. package berbeda
// 3. library yg dibuat orang lain
