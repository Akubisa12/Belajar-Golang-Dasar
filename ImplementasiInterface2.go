// Menyatakan bahwa kode ini adalah program utama yang dapat dijalankan secara independen.
package main

//Mengimport package "fmt" yang menyediakan fungsi untuk input dan output, terutama untuk pencetakan ke layar.
import "fmt"

// Mendefinisikan struktur data bernama "Animal" dengan satu field, yaitu "Name" bertipe string. Struktur ini mewakili data hewan dengan namanya.
type Animal struct {
	Name string
}

// Mendefinisikan interface bernama "HasName3" yang memiliki satu method, yaitu "GetName()" yang mengembalikan nilai string. Interface ini mewakili objek yang memiliki nama, yang dapat diambil melalui method "GetName()".
type HasName3 interface {
	GetName() string
}

// Mendefinisikan fungsi "sayHello3()" yang menerima argumen berupa objek yang memenuhi interface "HasName3". Fungsi ini mencetak tulisan "Hello" diikuti dengan nama yang didapatkan dari method "GetName()" pada objek tersebut.
func sayHello3(hasName HasName3) {
	fmt.Println("Hello ", hasName.GetName())
}

// Mendefinisikan method "GetName()" untuk struct "Animal". Method ini mengembalikan nilai field "Name" dari struct Animal, sehingga struct Animal memenuhi persyaratan interface "HasName3".
func (animal Animal) GetName() string {
	return animal.Name // Langsung kembalikan nilai field Name
}

// Fungsi utama yang dijalankan ketika program dimulai.
func main() {
	//Membuat instance struct Animal dengan nama "Kelinci".
	animal := Animal{Name: "Kelinci"}
	//Memanggil fungsi "sayHello3()" dengan memberikan instance struct Animal sebagai argumen.
	sayHello3(animal)
}
