package main

import "fmt"

func main() {
	name := "Miftahul"

	// IF EXPRESSION
	if name == "Miftahul" {
		fmt.Println("Haloo Miftahul")
	}
	// ELSE EXPRESSION
	nama := "Jarwo"
	if nama == "Falah"{
		fmt.Println("Haloo Falahh")
	}else{
		fmt.Println("Haii, Boleh Kenalan?")
	}

	//Else If Expression
	namaLengkap := "Miftahul Falah"
	if namaLengkap == "Mario Balotelli"{
		fmt.Println("Halo Bang Balo")
	}else if namaLengkap == "Miftahul Falah"{
		fmt.Println("Haloo mas Falah")
	}else{
		fmt.Println("Halo boleh kenalan ?")
	}

	// If Short Statement
	if length := len(namaLengkap); length > 7{ //; (titik koma) ini yang jadi ciri short statement
		fmt.Println("Nama Terlalu Panjang")
	}else{
		fmt.Println("Nama Sudah Benar")
	}
}