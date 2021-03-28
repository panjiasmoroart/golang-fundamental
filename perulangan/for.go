package main

import "fmt"

func main() {

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Saya sedang belajar Go : ", i)
	// }

	// for dengan gaya while
	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Saya sedang belajar Go : ", i)
	// 	i++
	// }

	// for biasanta digunakan untuk mengambil sebuah collection
	// untuk kasus ini string dulu
	// title := "Golang the best language"

	// for index, letter := range title {
	// fmt.Println("index : ", index, " letter : ", letter)
	// fmt.Println("index : ", index, " letter : ", string(letter))
	// }

	// mau pake range tapi ga butuh indexnya, kita bisa kasih tanda _
	// for _, letter := range title {
	// fmt.Println(" letter : ", string(letter))
	// }

	// quiz, jika indexnya bilangan genap 0, 2, 4 dst cetak title
	title := "Golang the best language"

	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("index : ", index, " letter : ", string(letter))
	// 	}
	// }

	// quiz, cetak hanya jika hurufnya adalah huruf vokal a, i, u, e, o
	for index, letter := range title {
		letterChar := string(letter)

		// if letterChar == "a" || letterChar == "i" || letterChar == "u" || letterChar == "e" || letterChar == "o" {
		// 	fmt.Println("index : ", index, " letter : ", string(letter))
		// }

		// bisa juga pake switch
		switch letterChar {

		case "a", "i", "u", "e", "o":
			fmt.Println("index : ", index, " letter : ", string(letter))
		}
	}
}
