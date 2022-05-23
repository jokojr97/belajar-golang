package main

import (
	"belajar/satu"
	"fmt"
)

// variabel global
var i int
var mm map[string]int
var arr []string

func main() {
	fmt.Println("hello world")

	// buat variabel otomatis
	ss := "hola"
	// buat variabel map
	mm = make(map[string]int)
	// cetak variabel dengan format PrintF
	fmt.Printf("lenght: %d; %s\n", len(mm), ss)
	// inisiasi map
	mm["satu"] = 1
	mm["dua"] = 2
	// cetak isi variabel map mm
	fmt.Printf("hello world: %#v\n", mm)
	fmt.Printf("lenght map mm: %d\n", len(mm))

	// deklarasi variabel map metode 2
	mm = map[string]int{"tiga": 3, "empat": 4}
	// cetak isi variabel map
	fmt.Printf("map mm : %#v\n", mm)

	// inisiasi array arr
	arr = make([]string, 0)
	// menambah variabel array arr
	arr = append(arr, "satu")
	arr = append(arr, "dua")
	// cetak arr dengan  format
	fmt.Printf("lenght arr: %#v\n", arr)

	// inisiasi array metode 2
	arr = []string{"empat", "lima"}
	// cetak arr denganformat
	fmt.Printf("lenght arr 2 : %#v\n", arr)
	fmt.Printf("lenght arr 2 : %s\n", arr[0])
	satu.Coba()
}
