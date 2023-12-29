package main

import "fmt"

// Mendefinisikan fungsi bernama "NewMap()" yang menerima argumen string "name" dan mengembalikan nilai bertipe map[string]string (map dengan key string dan value string).
func NewMap(name string) map[string]string {
	//Jika argumen "name" kosong, fungsi mengembalikan nil, yang menandakan bahwa map tidak dibuat.
	if name == "" {
		return nil
	} else { //Jika argumen "name" tidak kosong, fungsi mengembalikan map baru yang berisi satu key-value pair, yaitu "name" dengan nilai dari argumen "name".
		return map[string]string{
			"name": name,
		}
	}
}

func main() { //Fungsi utama yang dijalankan ketika program dimulai.
	//Memanggil fungsi "NewMap()" dengan argumen string kosong ("") dan menyimpan nilai kembaliannya ke dalam variabel "data".
	data := NewMap("")
	//Memeriksa apakah nilai "data" adalah nil (menandakan map tidak dibuat).
	//fmt.Println("Data Kosong"): Jika nil, mencetak tulisan "Data Kosong" ke layar.
	if data == nil {
		fmt.Println("Data Kosong")
	} else { //Jika map dibuat, mencetak isi map ke layar.
		fmt.Println(data)
	}
}

//Key Points:
//Fungsi "NewMap()" digunakan untuk membuat map baru dengan satu key-value pair, tetapi hanya jika argumen "name" tidak kosong.
//Konsep mengembalikan nil untuk menandakan bahwa map tidak dibuat adalah praktik yang umum dalam Go.
//Dalam kode ini, fungsi "NewMap()" bisa lebih disederhanakan dengan langsung mengembalikan map kosong jika argumen "name" kosong, alih-alih mengembalikan nil.
