package main

import "fmt"

func main() {

	type NoKTP string

	var ktpEko NoKTP = "1111111111"
	fmt.Println(ktpEko)
	fmt.Println(NoKTP("22222222222"))//ini untuk mengonversi string biasa (bukan tipe data NoKTP) menjadi tipe data NoKTP
}