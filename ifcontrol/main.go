package main

import "fmt"

func main() {
	age := 17
	nilai := 80
	score := 65

	// salah satu contoh yg tepat membuat variabel menggunakan var,
	// tidak langusng di di assign nilainya
	var grade string

	if age > 15 {
		fmt.Println("Boleh bermain game")
	}

	if nilai > 85 {
		fmt.Println("Selamat kamu mendapatkan Grade Nilai A")
	} else {
		fmt.Println("Selamat kamu mendapatkan Grade Nilai B")
	}

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		if age > 16 {
			fmt.Println("Umur kamu 17 tapi nilainya kurang bagus tidak boleh main game dulu")
		} else {
			fmt.Println("kurang dari 16")
		}
		grade = "C"

	} else {
		grade = "B"
	}

	fmt.Println(grade)
}

// if
// if else
// else if
// if bersarang atau if didalam if
