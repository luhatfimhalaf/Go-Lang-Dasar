package main

import "fmt"

func main() {
	var names [3]string //"names" itu nama tipe data array nya terus kalau "[3]" itu total data yang dapat ditampung oleh array nya terus kalau string ya tipe datanya
	// kalau di line 6 itu kita tidak bisa menggunakan [...] karena dianggap akan 0 datanya 
	names[0] = "Miftahul" //"names[0]" itu merupakan indeksnya 
	names[1] = "Falah"
	names[2] = "Chavila"
	// names[3] = "Salsabila" //Kalau kaya gini bakalan error karena si array tidak bisa menampung datanya bikez karena tidak muat

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//MEMBUAT ARRAY SECARA LANGSUNG 
	var values = [3]int{
		80,
		79, // ini misal kalau dikosongkan pada larik ketiga maka hasil outputnya akan 0
		90, // nah kalau kita tambahkan enter itu kita harus akhiri dengan koma
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//FUNCTION ARRAY
	var values1 = [...]int{ //kita bisa menggunakan [...] kalau misal kita belum bisa menentukan panjang array nya
		90,
		80,
		96,
	}

	fmt.Println(values1)
	fmt.Println(len(values1)) // untuk menghitung panjang array 
	values1[0] = 100 // untuk mengganti value dari indeks yang dituju pada array
	fmt.Println(values1)

	//NOTES: Di golang itu tidak ada kegiatan menghapus data array kalau kita ingin menghapus ya kita replace dengan data kosong saja
}