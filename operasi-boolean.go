package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80 //true
	var lulusAbsensi bool = absensi > 80 //false

	var lulus bool = lulusNilaiAkhir && lulusAbsensi //false

	fmt.Println(lulus)
}