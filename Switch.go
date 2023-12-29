package main

import "fmt"

func main() {
	name := "Jessen"
	switch name {
	case "Jessen":
		fmt.Println("Hello Jessen")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hai, boleh kenalan ?")
	}
}
