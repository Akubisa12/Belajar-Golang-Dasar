package main

import "fmt"

func endApp2() { //Mendefinisikan fungsi bernama endApp2 yang bertugas untuk menangani panic dan melakukan cleanup.
	fmt.Println("End App") //Mencetak pesan "End App" ke konsol.
	message := recover()   //Memanggil fungsi recover() untuk menangkap panic yang terjadi.
	// Jika panic tertangkap, nilai panic akan disimpan dalam variabel message. Jika tidak ada panic, message akan bernilai nil.
	fmt.Println("Terjaid Error", message) // Mencetak pesan "Terjadi Error" dan nilai message (jika ada) ke konsol.
}

func runApp2(error bool) { //Mendefinisikan fungsi bernama runApp2 yang menerima parameter boolean error untuk mensimulasikan terjadinya error.
	defer endApp2() // Menjadwalkan eksekusi fungsi endApp2 secara deferred. Jika terjadi panic dalam fungsi runApp2, maka endApp2 akan dipanggil sebelum panic menghentikan program.

	//Memeriksa apakah parameter error bernilai true. Jika iya, maka fungsi runApp2 akan memicu panic dengan pesan "ERROR".
	if error {
		panic("ERROR")
	}
}

func main() { //Fungsi utama yang dieksekusi saat program dijalankan.
	endApp2() //Memanggil fungsi endApp2 secara langsung untuk mendemonstrasikan cara kerjanya.
}

//Kesimpulan:
//Fungsi recover() berguna untuk menangkap panic dan mencegah program berhenti secara tiba-tiba.
//Biasanya, recover() digunakan dalam fungsi deferred untuk melakukan tindakan pemulihan sebelum panic menghentikan program.
//Dalam kode ini, fungsi endApp2 bertindak sebagai fungsi pemulihan untuk mencetak pesan error dan melakukan cleanup sebelum program berhenti.
