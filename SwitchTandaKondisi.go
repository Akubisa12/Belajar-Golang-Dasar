package main

import "fmt"

func main() {
	name := "Jessen"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama kamu sudah benar")
	}
}
