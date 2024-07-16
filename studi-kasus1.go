// Soal Nomor 1
/*
// Buatlah sebuah program kalkulator sederhana yang dapat melakukan operasi penjumlahan, pengurangan, perkalian, dan pembagian. Program ini harus meminta input dua angka dan operasi yang diinginkan dari pengguna, kemudian menampilkan hasilnya.
*/

package main

import "fmt"

func main()  {
	var num1,num2 float64
	var operation string
	
	fmt.Println("Masukkan Angka pertama : ")
	fmt.Scanln(&num1)
	fmt.Println("Masukkan Angka kedua : ")
	fmt.Scanln(&num2)
	fmt.Println("Masukkan Operasi (+,-,*,/) : ")
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Println("Hasil -> ",num1,"+",num2,"=",num1+num2)
	case "-":
		fmt.Println("Hasil -> ",num1,"-",num2,"=",num1-num2)
	case "/":
		if num2 != 0{
			fmt.Println("Hasil -> ",num1,"/",num2,"=",num1/num2)
		}else{
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan")
		}
	case "*":
		fmt.Println("Hasil -> ",num1,"*",num2,"=",num1*num2)
	default:
		fmt.Println("Operasi tidak valid")
	}

}