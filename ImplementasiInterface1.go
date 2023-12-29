package main // Menyatakan bahwa kode ini adalah program utama yang dapat dijalankan secara independen.

import "fmt" //Mengimport package "fmt" yang menyediakan fungsi untuk input dan output, terutama untuk pencetakan ke layar.

// Mendefinisikan struktur data bernama "Person" dengan satu field, yaitu "Name" bertipe string. Struktur ini mewakili data seseorang dengan namanya.
type Person struct {
	Name string
}

// Mendefinisikan interface bernama "HasName2" yang memiliki satu method, yaitu "GetName()" yang mengembalikan nilai string. Interface ini mewakili objek yang memiliki nama, yang dapat diambil melalui method "GetName()".
type HasName2 interface {
	GetName() string
}

// Mendefinisikan method "GetName()" untuk struct "Person". Method ini mengembalikan nilai field "Name" dari struct Person, sehingga struct Person memenuhi persyaratan interface "HasName2".
func (person Person) GetName() string {
	return person.Name
}

// Mendefinisikan fungsi "sayHello2()" yang menerima argumen berupa objek yang memenuhi interface "HasName2". Fungsi ini mencetak tulisan "Hello" diikuti dengan nama yang didapatkan dari method "GetName()" pada objek tersebut.
func sayHello2(hasName HasName2) {
	fmt.Println("Hello ", hasName.GetName())
}

// Fungsi utama yang dijalankan ketika program dimulai.
func main() {
	//Membuat instance struct Person dengan nama "Jessen".
	person := Person{Name: "Jessen"}
	//Memanggil fungsi "sayHello2()" dengan memberikan instance struct Person sebagai argumen.
	sayHello2(person)
}
