package main

import "fmt"

func main() {

	// dengan slice jadi lebih fleksibel jumlah elemennya, jadi lebih dinamis untuk menampung datanya
	var gammingConsoles []string
	// cara isi elemnnya
	gammingConsoles = append(gammingConsoles, "PlayStation 4")
	gammingConsoles = append(gammingConsoles, "Nintendo Switch")
	gammingConsoles = append(gammingConsoles, "Xbox One")

	// slice dengan langsung nilainya
	gammingNitendo := []string{"Super Mario", "Mortal Kombar", "Bomber Man"}

	fmt.Println(gammingConsoles)
	fmt.Println(gammingNitendo)

	// jika slicenya mau di iterate
	for _, console := range gammingConsoles {
		fmt.Println(console)
	}

}
