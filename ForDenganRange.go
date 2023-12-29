package main

import "fmt"

func main() {
	names := []string{"Jessen", "Andi", "Dika"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
