package db

import "fmt"

type Mongo struct {
}

func (m Mongo) GetUser() []int {
	fmt.Println("monggo get user")
	return []int{}
}
