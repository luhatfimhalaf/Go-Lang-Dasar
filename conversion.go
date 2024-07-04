package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) //ini akan menjadi negatif karena melebihi kapasitas dan sifatnya numberoverflow (akan balik lagi ke belakang sesuai minimum maksimumnya)
// int16 itu maksimalnya adalah 32767
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	//MENGUBAH 
	var name = "Miftahul Falah"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}