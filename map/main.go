package main

import "fmt"

func main() {
	// pada dasarnya array itu indexnya, keynya berupa angka,
	// kalau map itu keynya bisa kita tentukan sendiri int, string

	// deklarasi variabel
	var myMap map[string]int
	// untuk menginisialisasi nilanya
	myMap = map[string]int{}

	// cara mengisinya
	myMap["ruby"] = 9
	myMap["go"] = 9
	myMap["JavaScript"] = 8
	// di override
	myMap["go"] = 10

	fmt.Println(myMap)
	// mendapatkan nilai suatu map
	fmt.Println(myMap["JavaScript"])

	// membuat map beserta isinya, style vertikal
	myProgramming := map[string]string{"ruby": "is beautiful", "go": "is super fast"}

	// horizontal
	myDatabase := map[string]string{
		"mysql":      "is Relational database",
		"mongodb":    "is Non Relational database",
		"postgresql": "is Relational database",
	}

	fmt.Println(myProgramming)
	fmt.Println(myDatabase)

	// melakukan iterasi
	for key, value := range myDatabase {
		fmt.Println("key : ", key, " value : ", value)
	}

	// hapus map
	fmt.Println("==========Remove Map==========")
	delete(myDatabase, "postgresql")

	// coba lakukan looping lagi
	for key, value := range myDatabase {
		fmt.Println("key : ", key, " value : ", value)
	}

	// cek nilai map ada nilainya atau tidak
	// value := myDatabase["redis"]
	// fmt.Println(value)

	// dengan cara yg lebih asik
	// map mengembalikan dua nilai yang pertama valuenya, kedua optional
	valueData, isAvailabel := myDatabase["firebase"]
	fmt.Println(valueData)
	fmt.Println(isAvailabel)

}
