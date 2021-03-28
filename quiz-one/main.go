package main

import "fmt"

func main() {
	// quiz, jika indexnya bilangan genap 0, 2, 4 dst cetak title
	title := "Golang the best language"

	for index, letter := range title {
		// fmt.Println(index, string(letter))
		if index%2 == 0 {
			fmt.Println("index : ", index, " letter : ", string(letter))
		}
	}
	fmt.Println("=========================")
	// quiz, cetak jika hurufnya adalah huruf vokal a, i, u, e, o
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
