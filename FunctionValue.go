package main

import "fmt"

func getGoodBye(name string) string { //mendefinisikan fungsi bernama getGoodBye yang mengambil satu argumen string (name) dan mengembalikan string (string).

	//mengembalikan string "Good Bye" diikuti dengan nama yang diberikan sebagai argumen.
	return "Good Bye " + name
}

func main() {
	// membuat variabel goodBye dan menetapkan fungsi getGoodBye ke dalamnya. Ini berarti bahwa goodBye sekarang dapat digunakan untuk memanggil fungsi getGoodBye secara langsung.
	goodBye := getGoodBye

	//memanggil fungsi goodBye dengan argumen "Jessen" dan mencetak hasilnya ke konsol. Output yang dihasilkan akan berupa "Good Bye Jessen".
	fmt.Println(goodBye("Jessen"))
}
