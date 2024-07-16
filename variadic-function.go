package main

import "fmt"

// Ini memakai slice yang otomatis (varargs)
func sumAll1(numbers ...int) int {
	total := 0
	
	for _, number := range numbers { //jadi numbers ini akan menjadi sebuah slice
		total += number
	}
	return total
}
// Ini memakai slice yang manual
func sumAll2(numbers []int) int {
	total := 0

	for _, number := range numbers { //jadi numbers ini akan menjadi sebuah slice
		total += number
	}
	return total
}

func main() {
	// fmt.Println(sumAll2(10,10,10)) // kalo pakai slice manual, kit gabisa membuat eksekusi function nya seperti ini 
	fmt.Println(sumAll2([]int{10,10,10})) //kalo pakai slice manual, eksekusinya harus seperti ini dan jadinya lumayan ribet
	fmt.Println(sumAll1(10,10,10,10,10,10)) // nah kalo pake varargs itu kita tinggal eksekusi dalam bentuk parameter saja tanpa harus memanggil si slice nya
	fmt.Println(sumAll1(10,10,10,10,10,10,10)) //jadi kita bisa menerima banyak parameter jika kita menggunakan varargs tanpa kita memanggil slicenya terlebih dahulu

	// KASUS SLICE PARAMETER
	// Misalnya, kita sudah terlanjur memiliki slice
	numbers:=[]int{10,10,10,10,10,10,10,10}
	fmt.Println(sumAll1(numbers...))//kita tidak bisa menggunakan "numbers" saja kita harus menambahkan titik - titik setelah "numbers"
}