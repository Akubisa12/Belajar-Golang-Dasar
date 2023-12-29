package main

import "fmt"

func main() {
	book := make(map[string]string)
	book["tilte"] = "Buku Atomic Habits"
	book["Author"] = "James Clear"
	book["wrong"] = "Ups"
	delete(book, "wrong")
	fmt.Println(book)
}
