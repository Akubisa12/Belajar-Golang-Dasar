package main

import "fmt"

type filter func(string) string

func sayHelloWithFilter2(name string, filter func(string) string) { //mendefinisikan fungsi bernama sayHelloWithFilter yang mengambil dua argument yaitu name dan filter. name yait string yang mewakili nama untuk diucapkan sedangkan filter adalah fungsi yg mengambil string dan mengembalikan string lain, diguankan untuk memfilter nama sebelum mengucapkan salam
	fmt.Println("Hello ", filter(name))
}

func spamFilter2(name string) string { //mendefinisikan fungsi bernama spamFilter yang mengambil string (name) dan mengembalikan string (string).
	//memeriksa apakah nama adalah "Anjing". Jika iya, mengembalikan "..." untuk menyaring kata tersebut. Jika tidak, mengembalikan nama asli tanpa perubahan.
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	//memanggil fungsi sayHelloWithFilter dengan nama "Jessen" dan fungsi spamFilter sebagai filter.
	// Outputnya akan berupa "Hello Jessen" karena nama tersebut tidak disaring.
	sayHelloWithFilter2("Jessen", spamFilter2)
	// membuat variabel filter dan menetapkan fungsi spamFilter ke dalamnya.
	filter := spamFilter2
	//memanggil fungsi sayHelloWithFilter dengan nama "Anjing" dan fungsi filter (yang mengacu pada spamFilter) sebagai filter. Outputnya akan berupa "Hello ..." karena nama tersebut disaring oleh spamFilter.
	sayHelloWithFilter2("Anjing", filter)

}
