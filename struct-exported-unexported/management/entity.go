package management

import "fmt"

// jika struct dan property awalanya huruf kecil
// maka tidak akan bisan dipanggil diluar package

// type product struct {
// 	price float64
// }

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

// method DisplayUser
func (user User) DisplayUser() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirtsName, user.LastName, user.Email)
}

// method DisplayGroup
func (group Group) DisplayGroup() {
	fmt.Printf("Group Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")
	fmt.Println("User Name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirtsName)
	}
}
