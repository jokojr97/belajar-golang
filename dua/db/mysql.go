package db

import "fmt"

type Mysql struct {
}

func (m Mysql) GetUser() []int {
	fmt.Println("Mysql get user")
	return []int{}
}
