package main

import "fmt"

type Costumer struct {
	Name, Address string
	Age           int
}

func main() {
	var jessen Costumer
	jessen.Name = "Allouisius Jessen"
	jessen.Address = "Indonesia"
	jessen.Age = 17
	fmt.Println(jessen)
}
