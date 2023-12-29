package main

import "fmt"

func SayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello ", firstName, lastName)
}

func main() {
	SayHelloTo("Allouisius", "Jessen")
}
