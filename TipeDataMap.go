package main

import "fmt"

// Map
// Pada Array/Slice untuk mengakses data, kita menggunakan index Number mulai dari 0
// Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan
// jenis tipe data index yang akan kita gunakan
// secara sederhana map adalah tipe data kumpulan key-value ( kata kunci - nilai ),
// dimana kata kuncinya bersifat unik, tidak boleh sama

func main() {
	person := map[string]string{
		"name":    "Jessen",
		"address": "Purwokerto",
	}
	fmt.Println(person)            //map[address:Purwokerto name:Jessen]
	fmt.Println(person["name"])    // Jessen
	fmt.Println(person["address"]) //Purwokerto
}
