package main

import "fmt"

// Mendefinisikan tipe data baru bernama Blacklist yang merupakan fungsi dengan parameter string dan mengembalikan nilai boolean (true atau false).
// Fungsi ini akan digunakan untuk memeriksa apakah sebuah nama termasuk dalam blacklist.
type Blacklist func(string) bool

// Mendefinisikan fungsi registerUser yang menerima dua parameter yaitu :
// 1. name: String yang merepresentasikan nama pengguna yang ingin didaftarkan.
// 2. blacklist: Fungsi dari tipe Blacklist yang akan digunakan untuk memeriksa apakah nama pengguna tersebut termasuk dalam blacklist.
func registerUser(name string, blacklist Blacklist) {
	//Memeriksa apakah fungsi blacklist mengembalikan nilai true untuk nama yang diberikan.
	//Jika true, artinya nama tersebut termasuk dalam blacklist, maka dicetak pesan "You Are Blocked [nama]".
	//Jika false, artinya nama tersebut tidak termasuk dalam blacklist, maka dicetak pesan "Welcome [nama]".
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() { //Fungsi utama yang dieksekusi saat program dijalankan.

	//Mendefinisikan fungsi anonim yang berfungsi sebagai blacklist.
	//Fungsi ini mengembalikan nilai true jika nama yang diberikan adalah "anjing", dan false untuk nama lainnya.
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	//Memanggil fungsi registerUser dengan nama "Jessen" dan fungsi blacklist yang telah didefinisikan sebelumnya.
	//Karena "Jessen" tidak termasuk dalam blacklist, maka akan dicetak pesan "Welcome Jessen".
	registerUser("Jessen", blacklist)

	//Memanggil fungsi registerUser dengan nama "anjing" dan fungsi blacklist yang didefinisikan secara inline.
	//Karena "anjing" termasuk dalam blacklist, maka akan dicetak pesan "You Are Blocked anjing".
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
