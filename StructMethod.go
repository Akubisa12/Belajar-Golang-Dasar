package main

import "fmt"

type Costumer3 struct { //Mendefinisikan sebuah struct bernama Costumer3 yang memiliki tiga field
	Name, Address string
	Age           int
}

// Mendefinisikan metode bernama sayHello yang terikat pada struct Costumer3.
// (customer Costumer3): Menunjukkan bahwa metode ini menerima penerima (receiver) bertipe Costumer3 dan akan beroperasi pada instance struct tersebut.
func (customer Costumer3) sayHello() {
	fmt.Println("Hello My Name is", customer.Name) //Mencetak pesan "Hello My Name is" diikuti dengan nama customer yang diambil dari field Name.
}

func main() { //Fungsi utama yang dieksekusi saat program dijalankan.
	rully := Costumer3{Name: "Rully"} //Membuat instance struct Costumer3 bernama rully dan mengisi nilai field Name dengan "Rully".
	rully.sayHello()                  //Memanggil metode sayHello pada instance rully untuk mencetak pesan perkenalan.
}
