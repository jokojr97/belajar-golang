package main

import (
	"belajar/dua/db"
	"belajar/dua/service"
	"fmt"
)

type myInt int

type myStuct struct {
	ID     int
	Name   string
	Detail string
}

func (ms myStuct) Calculate(param1 int) {
	fmt.Println("calculate " + ms.Name)
}

func (ms myStuct) Calculate2(param1 int, param2 string) {
	fmt.Println("Calculate 2 " + param2)
}

func main() {
	// var my myInt = 10
	// fmt.Printf("%#v \n", my)

	// my := myStuct{
	// 	1,
	// 	"my Name",
	// 	"My Detail",
	// }

	// fmt.Printf("%#v\n", my)

	// my2 := myStuct{
	// 	ID:   2,
	// 	Name: "Joko",
	// }

	// fmt.Printf("%#v\n", my2)

	// // memanggil stuct function
	// my2.Calculate(12)
	// my2.Calculate2(12, "Jokori")

	// ss := satu.NewSs(12, "Name", "detail", false)
	// fmt.Printf("%#v\n", ss.Detail())

	// ss2 := satu.NewSs(12, "Name", "detail", true)
	// fmt.Printf("%#v\n", ss2.Detail())

	// interface
	// var fk db.Fake = 10
	// service := Service{api: fk}

	// service := Service{api: db.Mongo{}}
	// service.api.GetUser()

	var fk db.Fake = 10
	srv := service.New(fk, fk)
	srv.GetUser()
	srv.Login("usernamesaya", "passwordsaya")
}
