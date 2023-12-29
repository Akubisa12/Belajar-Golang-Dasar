package main

import "fmt"

// Panic function adalah function yang bisa kita gunakan untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, namun defer function tetap diekseskusi

//	Mendefinisikan fungsi bernama endApp yang mencetak pesan "End App" ke konsol.
//
// Fungsi ini berfungsi sebagai cleanup untuk menandakan berakhirnya aplikasi.
func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) { //Mendefinisikan fungsi bernama runApp yang menerima parameter boolean error untuk mensimulasikan terjadinya error.
	defer endApp() // Menjadwalkan eksekusi fungsi endApp secara deferred.
	// Artinya, fungsi endApp akan dieksekusi secara otomatis sebelum fungsi runApp selesai, terlepas dari apakah terjadi error atau tidak.

	//Memeriksa apakah parameter error bernilai true. Jika iya, maka fungsi runApp akan memicu panic dengan pesan "ERROR".
	if error {
		panic("ERROR")
	}
}

func main() {
	endApp()
}

//Perilaku kode dalam dua skenario:
//Skenario 1: Tidak ada error
//Jika parameter error bernilai false, maka fungsi runApp akan berjalan normal tanpa panic.
//Setelah semua perintah selesai dieksekusi, fungsi endApp yang dijadwalkan secara deferred akan dipanggil, mencetak "End App".

//Skenario 2: Terjadi error
//Jika parameter error bernilai true, maka fungsi runApp akan memicu panic dengan pesan "ERROR".
//Panic akan menghentikan eksekusi fungsi runApp secara tiba-tiba.
//Namun, sebelum fungsi runApp sepenuhnya berhenti, fungsi endApp yang dijadwalkan secara deferred akan tetap dieksekusi, mencetak "End App".

//Output kode dalam kedua skenario:
//Skenario 1: "End App"
//Skenario 2: "End App" (diikuti pesan panic "ERROR")

//Kesimpulan:
//Kata kunci defer berguna untuk memastikan eksekusi fungsi cleanup, meskipun terjadi error atau panic dalam fungsi yang sama.
//Hal ini penting untuk menjamin operasi cleanup yang penting, seperti menutup file, melepaskan resource, atau mencatat log error.
