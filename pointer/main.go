package main

import "fmt"

func main() {
	// pointer itu pada dasarnya sebuah variabel yg isinya suatu alamat dari suatu nilai
	// yg disimpan itu bukan nilainya tapi alamat memori dikomputer dimana nilai itu disimpan

	numberA := 5
	// referencing, menggunakan &
	// numberB bukan menyimpan angka 5, tapi menyimpan alamat dimana angkan 5 ini disimpan di memori komputer
	numberB := &numberA
	fmt.Println("============sample pointer pertama=============")
	fmt.Println(numberA)
	fmt.Println(numberB)
	// *dereferencing, mengembalikan alamat memori menjadi suatu isi nilai variabel
	fmt.Println(*numberB)

	// numberB ini kan berisi alamat numberA, jadi sebenarnya numberA dan numberB mengacu pada alamat memori yg sama
	// maka ketika salah kita rubah, maka numberA juga ikut berubah nilanya
	*numberB = 10

	fmt.Println(*numberB)
	fmt.Println(numberA)

	fmt.Println("============sample pointer kedua=============")
	var numberC int = 5
	// *int berarti untuk menyimpan alama memori
	var numberD *int = &numberC

	fmt.Println(numberC)
	fmt.Println(numberD)
	fmt.Println(*numberD)

	numberC = 20

	fmt.Println(numberC)
	fmt.Println(numberD)
	fmt.Println(*numberD)

	fmt.Println("============sample pointer ketiga=============")
	number := 5
	fmt.Println("Alamat memory", &number)
	fmt.Println("Nilai awal :", number)

	change(&number, 100)

	fmt.Println("Nilai akhir :", number)
	fmt.Println("Alamat memory", &number)
}

// change value on function
func change(old *int, new int) {
	*old = new
	// fmt.Println("Alamat memory", &old)
	fmt.Println("Alamat memory", old)
	fmt.Println("Di dalam function :", *old)
}
