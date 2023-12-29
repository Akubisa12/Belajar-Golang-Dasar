package main

import "fmt"

func main() {

	//mendefinisikan variabel names. Variabel names adalah array dari lima string.
	// Nilai-nilai dalam array adalah "Jessen", "Eko", "Joko", "Budi", dan "Kinan".
	names := [...]string{"Jessen", "Eko", "Joko", "Budi", "Kinan"}
	// membuat slice dari array names. Slice adalah koleksi data yang dapat diubah ukurannya. Slice ini dimulai di indeks 1 dan berakhir di indeks 5.
	slice := names[1:5]
	//mencetak nilai dari elemen pertama dalam slice. Nilai elemen pertama adalah "Eko".
	fmt.Println(slice[0])
	// mencetak nilai dari elemen ketiga dalam slice. Nilai elemen ketiga adalah "Joko"
	fmt.Println(slice[2])
}
