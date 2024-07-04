package main

import "fmt"

func main() {
	names := [...]string{ //";=" merupakan short variables declaration
		"Miftahul",
		"Falah",
		"Marsa",
		"Salsabila",
		"Chavila",
		"luhatfim",
		"halaf",
		"Asram",
	}
	slice1 := names[4:6] 
	
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])
	
	slice2 := names[:5] 
	fmt.Println(slice2)
	
	slice3 := names[2:] 
	fmt.Println(slice3)
	
	var slice4 []string = names[:] //ini untuk mendeklarasikan slice secara manual (tidak langsung) dan kita tidak perlu tentukan jumlah tampungan slice nya
	fmt.Println(slice4)

	//FUNCTION SLICE
	days:= [...]string{"Senin","Selasa","Rabu","Kamis","Jumat","Sabtu","Minggu"}
	dayslice1 := days[5:] //Sabtu, Minggu
	dayslice1[0] = "Sabtu Baru"
	dayslice1[1] = "Minggu Baru"
	fmt.Println(days) //[Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru]

	dayslice2 := append(dayslice1, "Libur Baru") //karena kapasitas array nya itu tidak bisa memenuhi untuk menambahkan data, maka otomatis seolah - olah akan terbentuk array baru
	// dan seakan - akan seperti ini days:= [...]string{"Senin","Selasa","Rabu","Kamis","Jumat","Sabtu","Minggu"}
	fmt.Println(dayslice2)
	dayslice2[0] = "Sabtu Lama"
	fmt.Println(dayslice2) //[Ups ,Minggu Baru, Libur Baru]
	fmt.Println(days) //[Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru]

	// KODE PROGRAM MAKE SLICE
	// Jadi kita bisa membuat slice baru dengan otomatis membuat array nya juga

	newSlice := make([]string, 2,5) //make([]string, length, capacity)
	newSlice[0] = "Falah"
	newSlice[1] = "Miftahul"
	// newSlice[2] = "Miftahul" //kita tidak bisa menambahkan data pada array dengan cara seperti ini, maka dari itu kita harus menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	
	newSlice2 := append(newSlice, "Falahhh")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	//KODE PROGRAM COPY SLICE
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) //agar bisa ditampung semua datanya berarti kita perlu bikin slice baru yang ukurannya sama dari fromSlice

	copy(toSlice, fromSlice) //ini dasar code copy slice nya

	fmt.Println(toSlice)

	//NOTES: Jangan sampai ketika kita append kita malah bikin array baru, maka itu akan menghasilkan aplikasi yang lambat karena kita harys membuat array baru secara terus menerus
	//NOTES: Saat membuat array, kita harus berhati - hati, jika salah, maka yang kita buat bukan array melainkan slice

	//CONTOH KODE PERBEDAAN MAKE ARRAY DAN MAKE SLICE
	iniArray := [...]int{1,2,3,4,5} //Ini Array
	iniSlice := []int{1,2,3,4,5} //Ini Slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	//NOTES: Jadi, Slice atau Array? nah kalau kita berada di golang pasti rata - rata menggunakan slice
}