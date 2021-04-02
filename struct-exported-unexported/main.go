package main

import (
	"fmt"
	"struct-exported-unexported/management"
)

func main() {
	user1 := management.User{
		ID:        1,
		FirtsName: "Muhammad",
		LastName:  "Fathan",
		Email:     "fathan@gmail.com",
		IsActive:  true,
	}

	user2 := management.User{
		ID:        2,
		FirtsName: "Ibnu",
		LastName:  "Sinna",
		Email:     "ibnusinna@gmail.com",
		IsActive:  true,
	}

	user3 := management.User{
		ID:        1,
		FirtsName: "Muhammad",
		LastName:  "Salah",
		Email:     "muhammadsalah@gmail.com",
		IsActive:  true,
	}

	resultUser1 := user1.DisplayUser()
	fmt.Println(resultUser1)

	resultUser2 := user2.DisplayUser()
	fmt.Println(resultUser2)

	fmt.Println("")
	users := []management.User{user1, user2}
	group := management.Group{"Backend", user3, users, true}
	group.DisplayGroup()
}
