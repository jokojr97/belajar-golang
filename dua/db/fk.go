package db

import "fmt"

type Fake int

func (fk Fake) GetUser() []int {
	fmt.Println("Fake get user")
	return []int{}
}

func (fk Fake) Signing(username string, password string) bool {
	fmt.Println("Fake login")
	return false
}
