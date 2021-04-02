package main

import "fmt"

type User struct {
	ID        int
	FirtsName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailabel bool
}

// method
func (group Group) displayMethod() {
	fmt.Printf("Group Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")
	fmt.Println("User Name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirtsName)
	}
}

func main() {
	user1 := User{
		ID:        1,
		FirtsName: "Muhammad",
		LastName:  "Fathan",
		Email:     "fathan@gmail.com",
		IsActive:  true,
	}

	user2 := User{
		ID:        2,
		FirtsName: "Ibnu",
		LastName:  "Sinna",
		Email:     "ibnusinna@gmail.com",
		IsActive:  true,
	}

	users := []User{user1, user2}
	group := Group{"Backend", user1, users, true}

	// call function
	// displayGroup(group)

	// call method
	group.displayMethod()
}

// quiz rubah sebuah function displayGroup menjadi sebuah method
// func displayGroup(group Group) {
// 	fmt.Printf("Group Name : %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count : %d", len(group.Users))
// 	fmt.Println("")
// 	fmt.Println("User Name : ")
// 	for _, user := range group.Users {
// 		fmt.Println(user.FirtsName)
// 	}
// }
