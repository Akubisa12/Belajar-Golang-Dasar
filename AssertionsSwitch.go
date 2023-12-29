package main

import "fmt"

// Mendefinisikan fungsi bernama "random2()" yang mengembalikan interface kosong (interface{}), yang dapat menampung tipe data apa pun. Namun, fungsi ini selalu mengembalikan string "Ok".
func random2() interface{} {
	return "Ok"
}

func main() {
	//Memanggil fungsi "random2()" dan menyimpan nilai kembaliannya (string "Ok") ke dalam variabel "result".
	result := random2()
	// Menggunakan switch type untuk memeriksa tipe data dari variabel "result".
	switch value := result.(type) {
	//Jika tipe data "result" adalah string, cetak tulisan "String" diikuti dengan nilai "result".
	case string:
		fmt.Println("String", value)
	//Jika tipe data "result" adalah integer, cetak tulisan "Int" diikuti dengan nilai "result".
	case int:
		fmt.Println("Int", value)
	//Jika tipe data "result" tidak sesuai dengan case string atau int, cetak tulisan "Unknown".
	default:
		fmt.Println("Unknown")
	}
}
