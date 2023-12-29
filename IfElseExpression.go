package main

import "fmt"

func main() {
	name := "Jessen"
	if name == "Jessen" {
		fmt.Println("Hello Jessen")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hai, siapa nama kamu?")
	}
}
