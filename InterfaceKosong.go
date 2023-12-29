// Menyatakan bahwa kode ini adalah program utama yang dapat dijalankan secara independen.
package main

//Mengimport package "fmt" yang menyediakan fungsi untuk input dan output, terutama untuk pencetakan ke layar.
import "fmt"

// Mendefinisikan fungsi bernama "Ups()" yang mengembalikan nilai bertipe interface{}. Interface kosong ini dapat menampung nilai dari tipe data apa pun.
func Ups() interface{} {
	//Fungsi ini mengembalikan string "Ups" sebagai nilainya.
	return "Ups"
}

// Fungsi utama yang dijalankan ketika program dimulai
func main() {
	//Memanggil fungsi "Ups()" dan menyimpan nilai kembaliannya (string "Ups") ke dalam variabel "kosong" bertipe interface{}.
	kosong := Ups()
	//Mencetak nilai variabel "kosong" ke layar menggunakan fungsi "Println()" dari package "fmt".
	fmt.Println(kosong)
}

//Key Points:
//Penggunaan interface{} memungkinkan untuk menyimpan nilai dari tipe data apa pun, memberikan fleksibilitas dalam pemrograman.
//Namun, perlu diperhatikan bahwa saat menggunakan interface{}, Anda perlu melakukan type assertion untuk mengakses nilai sebenarnya yang tersimpan di dalamnya.

//Penjelasan Tambahan:
//Dalam kode ini, penggunaan interface{} sebenarnya tidak terlalu memberikan keuntungan karena nilai yang dikembalikan sudah diketahui bertipe string.
//Jika Anda ingin mengembalikan nilai dari berbagai tipe data, penggunaan interface{} bisa lebih bermanfaat.
//Namun, dalam kasus ini, jika Anda yakin nilai yang dikembalikan selalu string, lebih baik langsung mengembalikan string agar kode lebih jelas dan mudah dipahami.
