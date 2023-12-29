package main

import "fmt"

// closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
// harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi

func main() { // Fungsi utama yang dieksekusi saat program dijalankan.
	counter := 0          // Mendeklarasikan variabel counter dengan nilai awal 0.
	increment := func() { // Mendefinisikan fungsi anonim bernama increment yang bertindak sebagai closure.
		fmt.Println("Increment") //  Mencetak pesan "Increment" ke konsol.
		counter++                //  Menambah nilai counter dengan 1. Meskipun counter dideklarasikan di luar fungsi increment, closure dapat mengakses dan memodifikasinya.
	}
	increment()          // Saat fungsi increment dipanggil pertama kali, ia mencetak "Increment" dan nilai counter menjadi 1.
	increment()          // Saat fungsi increment dipanggil kedua kali, ia kembali mencetak "Increment" dan nilai counter menjadi 2.
	fmt.Println(counter) // Terakhir, nilai counter yang sudah menjadi 2 dicetak ke konsol.
}
