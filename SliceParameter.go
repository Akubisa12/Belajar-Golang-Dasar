package main

import "fmt"

func sumAll2(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	//memanggil fungsi sumAll dengan empat argumen integer (10, 10, 10, 10)
	//dan menyimpan hasil penjumlahannya dalam variabel total.
	total := sumAll2(10, 10, 10, 10)
	// mencetak nilai total ke konsol, yang akan menampilkan 40 dalam kasus ini.
	fmt.Println(total)

	//membuat slice (seperti array dinamis) bernama numbers yang berisi empat angka 10.
	numbers := []int{10, 10, 10, 10}
	//memanggil fungsi sumAll lagi, tetapi kali ini menggunakan numbers...
	//yang menyebarkan elemen-elemen dalam slice numbers sebagai argumen individual ke fungsi.
	//Hasil penjumlahan lagi disimpan dalam variabel total.
	total = sumAll2(numbers...)
	//mencetak baris kosong ke konsol.
	fmt.Println()
}
