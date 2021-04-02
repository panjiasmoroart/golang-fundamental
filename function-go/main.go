package main

import "fmt"

func main() {
	// panggil function
	printMyResult()
	printWithParams("Go")
	printWithParams("Best")
	printWithParams("Performance")
	// untuk menangkap output kita perlu tampung ke sebuah variable
	sentence := printWithParamsAndReturn("Coding go")
	fmt.Println(sentence)

	resultAdd := add(10, 5)
	fmt.Println(resultAdd)

	// luas, keliling := multipleRetrun(10, 2)
	luas, _ := multipleRetrun(10, 2)
	fmt.Println(luas)
	// fmt.Println(keliling)

	luas2, keliling2 := multipleRetrunPredefined(10, 2)
	fmt.Println(luas2)
	fmt.Println(keliling2)
}

func printMyResult() {
	fmt.Println("Saya sedang belajar Go")
}

func printWithParams(data string) {
	fmt.Println(data)
}

func printWithParamsAndReturn(value string) string {
	newValue := value + " sangat menyenangkan"
	return newValue
}

func add(number, number1 int) int {
	return number + number1
}

func multipleRetrun(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

// predefined return value mendefinisikan varaiabel di parameter
func multipleRetrunPredefined(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}
