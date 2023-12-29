package main

import "fmt"

// defer adalah function yang bisa kita jadwalkan untuk dieksekusi stlh sebuah function selesai di eksekusi
// defer function akan selalu dieksekusi walaupun terjadi error di function yang diekseskusi

func logging() { //Mendefinisikan fungsi bernama logging yang mencetak pesan "Selesai memanggil function" ke konsol.
	fmt.Println("Selesai memanggil function")
}

func runApplication() { //Mendefinisikan fungsi bernama runApplication yang bertugas menjalankan aplikasi.
	defer logging()                //Menjadwalkan eksekusi fungsi logging secara deferred. Artinya, fungsi logging tidak akan dieksekusi saat ini, melainkan akan ditunda sampai fungsi runApplication selesai.
	fmt.Println("Run Application") //Mencetak pesan "Run Application" ke konsol.
}

func main() { //Fungsi utama yang dieksekusi saat program dijalankan.
	runApplication() //Memanggil fungsi runApplication.
}

//Penjelasan:
//1. Saat fungsi runApplication dipanggil, pertama-tama ia menjadwalkan fungsi logging untuk dieksekusi secara deferred.
//2. Kemudian, fungsi runApplication mencetak pesan "Run Application" ke konsol.
//3. Setelah semua perintah di dalam fungsi runApplication selesai dijalankan, maka fungsi logging yang dijadwalkan secara deferred akan dieksekusi.
//4. Oleh karena itu, pesan "Selesai memanggil function" dicetak setelah pesan "Run Application".

//Kata kunci defer sering digunakan untuk beberapa tujuan, antara lain:
//1. Logika cleanup: Menjamin bahwa operasi tertentu akan dieksekusi, walaupun terjadi error atau panic di tengah fungsi. Misalnya, untuk menutup file yang dibuka, melepaskan lock, atau membebaskan resource yang dialokasikan.
//2. Penundaan operasi: Menunda eksekusi sebuah operasi sampai akhir fungsi untuk meningkatkan kejelasan kode dan menghindari masalah urutan eksekusi.
