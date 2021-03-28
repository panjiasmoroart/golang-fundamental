package main

import "fmt"

func main() {

	// mendefinisikan varaiabelnya terlebih dahulu
	var languages [5]string
	languages[0] = "Go"
	languages[1] = "Java"
	languages[2] = "Ruby"
	languages[3] = "JavaScript"
	languages[4] = "PHP"

	// bisa juga langsung kita isi element arraynya
	bestFrontend := [3]string{"ReactJS", "Angular", "VueJS"}

	// jika ingin mendefinisikannya secara vertikal harus tambahkan , dibelakanganya
	bestBackend := [6]string{
		"Gin Gonic",
		"Spring Boot",
		"Django",
		"Ruby On Rails",
		"Laravel",
		"ExpressJS",
	}

	// jika ingin lebih fleksibel element arranya
	bestDatabase := [...]string{
		"Oracle",
		"MySQL",
		"PostgreSQL",
		"MongoDB",
		"Redis",
	}

	fmt.Println(languages)
	countArray := len(languages)
	fmt.Println(countArray)

	fmt.Println(bestFrontend)
	fmt.Println(bestBackend)

	fmt.Println(bestDatabase)

	// looping array
	fmt.Println("Looping element array : ")
	for index, backend := range bestBackend {
		fmt.Println("Index : ", index, " Best Backend Framework : ", backend)
	}

}
