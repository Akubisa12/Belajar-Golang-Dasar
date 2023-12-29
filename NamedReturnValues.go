package main

import "fmt"

func getCompleteName(firstname, middlename, lastname string) (string, string, string) {
	firstname = "allouisius"
	middlename = "jessen"
	lastname = "muladi"
	return firstname, middlename, lastname
}

func main() {
	firstName, middleName, lastName := getCompleteName("allouisius", "jessen", "muladi") // Provide empty strings as arguments
	fmt.Println(firstName, middleName, lastName)
}
