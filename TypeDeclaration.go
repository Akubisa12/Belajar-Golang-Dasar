package main

import "fmt"

func main() {
	type noKTP string

	var ktpJessen noKTP = "111111"
	fmt.Println(ktpJessen)
	fmt.Println(noKTP("22222"))
}
