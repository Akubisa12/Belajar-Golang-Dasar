package main

import "fmt"

// Mendefinisikan fungsi rekursif bernama factorialRecursive yang menerima satu parameter
// yaitu value yang merupakan Bilangan integer yang ingin dihitung faktorialnya.
func factorialRecursive(value int) int {
	if value == 1 { //Memeriksi apakah nilai value adalah 1
		//Jika true, maka fungsi mengembalikan 1, karena faktorial dari 1 adalah 1.
		return 1
	} else { //Jika nilai value bukan 1, maka fungsi melakukan perhitungan rekursif
		//Mengembalikan hasil perkalian antara value dengan nilai faktorial dari value-1, yang diperoleh dengan memanggil fungsi factorialRecursive secara rekursif.
		return value * factorialRecursive(value-1)
	}
}

func main() { //Fungsi utama yang dieksekusi saat program dijalankan.
	fmt.Println(factorialRecursive(10)) // Memanggil fungsi factorialRecursive dengan nilai 10 dan mencetak hasilnya ke konsol.
}
