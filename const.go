package main

import "fmt"

func main() {
	const firstName string = "Miftahul" //ga masalah kalau tidak digunakan tapi gabisa diubah isi value variable nya
	const lastName = "Khannedy"

	//error
	// fisrtName = "Tidak Bisa Diubah"
	// lastName = "Tidak Bisa Diubah"

	//Deklarasi Multiple Constant
	//Sama seperti variabel, di Go-Lang juga kita bisa membuat constant secara sekaligus banyak

	const (
		namaPertama  string = "Miftahul"
		namaTerakhir        = "Falah"
	)

	fmt.Println(namaPertama)
	fmt.Println(namaTerakhir)
}