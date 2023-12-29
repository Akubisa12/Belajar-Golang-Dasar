package main //Mendeklarasikan bahwa kode ini adalah program utama yang dapat dijalankan secara langsung.

import "fmt" //Mengimpor package fmt untuk fungsi-fungsi input/output, seperti Println untuk mencetak ke konsol.

type Costumer2 struct { //Mendefinisikan sebuah struct bernama Costumer2 yang memiliki tiga field:
	Name, Address string
	Age           int
}

func main() { //Fungsi utama yang dieksekusi saat program dijalankan.
	jessen := Costumer2{ //Membuat instance struct Costumer2 bernama jessen dan mengisi nilai field-nya secara eksplisit
		Name:    "Jessen",
		Address: "Indonesia",
		Age:     17,
	}
	fmt.Println(jessen) //Mencetak isi struct jessen ke konsol.

	//Membuat instance struct Costumer2 lain bernama budi dengan cara yang lebih ringkas, langsung mengisi nilai field sesuai urutan.
	budi := Costumer2{"Budi", "Indonesia", 33}
	//Mencetak isi struct budi ke konsol.
	fmt.Println(budi)
}
