package main

import "fmt"

func main() {
	name := "Jessen"
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama suda benar")
	}
}
