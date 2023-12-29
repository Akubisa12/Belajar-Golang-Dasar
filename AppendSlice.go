package main

import "fmt"

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) // Senin, Selasa, Rabu, Kamis, Jumaat, Sabtu Baru, Minggu Baru

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) // Ups, Minggu Baru, Libur Baru
	fmt.Println(days)      // Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru
}
