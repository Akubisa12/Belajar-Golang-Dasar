package main

import "fmt"

// Mendefinisikan fungsi bernama "random()" yang mengembalikan interface kosong (interface{}), yang dapat menampung tipe data apa pun. Namun, fungsi ini selalu mengembalikan string "Ok".
func random() interface{} {
	return "Ok"
}

func main() {
	//Memanggil fungsi "random()" dan menyimpan nilai kembaliannya (string "Ok") ke dalam variabel "result".
	result := random()
	//Mencoba untuk menetapkan "result" sebagai string dan menyimpannya di "resultString". Ini akan berhasil.
	resultString := result.(string)
	//Mencetak "resultString" ke konsol, yang akan menghasilkan output "Ok".
	fmt.Println(resultString)
	//Mencoba untuk menetapkan "result" sebagai integer dan menyimpannya di "resultInt". Ini akan gagal, menyebabkan panic runtime.
	resultInt := result.(int)
	//Baris ini tidak akan tercapai karena panic.
	fmt.Println(resultInt)
}

//Poin-poin Penting:
//Penggunaan interface{} memungkinkan fleksibilitas dalam menyimpan berbagai tipe data, tetapi membutuhkan type assertion yang cermat.
//Type assertion yang salah dapat menyebabkan panic runtime.
//Secara umum, lebih aman untuk menghindari penggunaan interface kosong yang tidak perlu dan mengandalkan tipe eksplisit bila memungkinkan.
