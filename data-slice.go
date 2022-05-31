package main

import "fmt"

func main() {
	var month = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//month[5] = "Diubah"
	//fmt.Println(slice1)

	//slice1[0] = "Mei lagi"
	//fmt.Println(month)

	var slice2 = month[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Linux")
	fmt.Println(slice3)
	slice3[1] = "Bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(month)

	//membuat slice yang aman
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Hello"
	newSlice[1] = "World"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	/* Hati - hati saat membuat array, karena nanti yang dibuat malahan slice */
}