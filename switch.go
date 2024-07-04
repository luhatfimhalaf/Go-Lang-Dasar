package main

import (
	"fmt"
)

func main() {
	name := "Falah"

	switch name {
	case "Falah":
		fmt.Println("Hello Falah")
	case "Ahmad":
		fmt.Println("Halo Ahmad")
	default:
		fmt.Println("Halo, Boleh berkenalan")
	}

	// SWITCH DENGAN SHORT STATEMENT
	length := len(name)
	switch length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama kurang panjang")
	}

	// SWITCH TANPA KONDISI
	length1 := len(name)
	switch{
	case length1 > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length1 > 5:
		fmt.Println("Nama Lumayan Panjang")
	case length1 < 4:
		fmt.Println("Nama pendek")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}