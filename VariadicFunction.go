package main

import "fmt"

// mendefinisikan fungsi bernama sumAll yang
// mengambil sejumlah variabel argumen integer (numbers ...int)
// dan mengembalikan integer (int).
func sumAll(numbers ...int) int { // Loop ini mengulangi setiap integer dalam slice numbers
	//Menginisialisasi variabel total ke 0 untuk menyimpan jumlah.
	total := 0
	// _ digunakan untuk membuang indeks elemen saat ini, karena tidak diperlukan untuk perhitungan.
	for _, number := range numbers {
		total += number // number saat ini ditambahkan ke total.
	}
	return total // Setelah loop, total akhir dikembalikan dari fungsi.
}

func main() {
	//Memanggil fungsi sumAll dengan empat argumen integer (10, 10, 10, 10)
	//dan menyimpan jumlah yang dikembalikan di variabel total.
	total := sumAll(10, 10, 10, 10)
	//Mencetak nilai total ke konsol.
	fmt.Println(total)
}
