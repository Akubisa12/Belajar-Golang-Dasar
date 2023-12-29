package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80
	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80
	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
