package main

import "fmt"

func main() {
	// numbers := [...]int{1, 2, 3, 4}
	// quiz pertama hitung rata-rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// total := 0
	var total int

	for _, value := range scores {
		// total = total + value
		total += value
	}

	// hitung element array
	countArray := len(scores)

	// hitung rata-rata total/jumlah_elemen_array
	avgScore := (float64(total)) / (float64(countArray))

	fmt.Println("==========Mencari Rata-rata Scores=========")
	fmt.Println(scores)
	fmt.Println("Hasilnya : ", avgScore)

	// quiz kedua masukan scores yang nilainya >= 90 ke array goodScores
	scoresTwo := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int

	fmt.Println("=====masukan scores yang nilainya >= 90 ke array baru=====")
	fmt.Println(scoresTwo)
	for _, value := range scoresTwo {
		if value >= 90 {
			// fmt.Println(value)
			goodScores = append(goodScores, value)
			// fmt.Println(goodScores)
		}
	}

	fmt.Println("Hasilnya : ", goodScores)
}
