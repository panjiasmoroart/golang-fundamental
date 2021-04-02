package main

import "fmt"

type User struct {
	ID        int
	FirtsName string
	LastName  string
	Email     string
	IsActive  bool
}

// struct juga bisa jadi atribute/field dari struct lainnya
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailabel bool
}

// method sebenarynya function, bedanya function yg dimiliki sebuah object
// object itu hasil instance dari sebuah struct
func (user User) displayMethod() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirtsName, user.LastName, user.Email)
}

func main() {
	// cara pertama
	user := User{}
	user.ID = 1
	user.FirtsName = "Panji"
	user.LastName = "Asmoro"
	user.Email = "panjiasmoro@gmail.com"
	user.IsActive = true

	user2 := User{}
	user2.ID = 2
	user2.FirtsName = "Salman"
	user2.LastName = "Alfarisi"
	user2.Email = "salmanalfarisi@gmail.com"
	user2.IsActive = true

	// cara kedua
	user3 := User{
		ID:        3,
		FirtsName: "Muhammad",
		LastName:  "Fathan",
		Email:     "fathan@gmail.com",
		IsActive:  true,
	}

	// cara ketiga harus urut structnya
	user4 := User{
		1,
		"Ibnu",
		"Sinna",
		"ibnu@gmail.com",
		false,
	}

	fmt.Println(user)
	fmt.Println(user2.FirtsName)
	fmt.Println(user3)
	fmt.Println(user4)

	fmt.Println("")
	fmt.Println("==========struct sebagai parameter==========")
	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user2)
	fmt.Println(displayUser1)
	fmt.Println(displayUser2)

	fmt.Println("")
	fmt.Println("==========struct sebagai atribute/field dari struct lainnya===========")
	users := []User{user3, user4}
	group := Group{"Gamer", user3, users, true}
	displayGroup(group)

	fmt.Println("")
	fmt.Println("===========method================")
	// menampilkan sebuah struct menggunakan method
	resultMethod := user.displayMethod()
	fmt.Println(resultMethod)

}

// struct sebagai parameter
func displayUser(user User) string {
	// Sprintf akan mengembalikan sebuah string,harus disimpan kesebuah variabel
	// karena nilai kembalian
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirtsName, user.LastName, user.Email)
	return result
}

func displayGroup(group Group) {
	// kalo printf langsung akan mencetak ke terminal
	// secara default Printf tidak ada karater enter
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	// fmt.Printf(group.Admin.FirtsName)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")
	fmt.Println("User Name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirtsName)
	}
}

// function itu tidak dimiliki siapa siapa, paling melekat kesebuah package
// function showUser(user User)string {
// }
