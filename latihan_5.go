package main

import "fmt"

func main(){
	// string
	var saya string

	saya = "Andi kurniawan"
	fmt.Println(saya)

	var nama = "Faiz Haidar"
	fmt.Println(nama)

	var me = "Haidar Halwi"
	fmt.Println(me)


	// int
	var angka = 89
	fmt.Println(angka)

	var angka2 int

	angka2 = 56
	fmt.Println(angka2)

	// tanpa var
	// penggunaan hanya di awal sisanya boleh tidak pakai :
	negara := "indonesia"
	fmt.Println(negara)


	// cara lebih simple
	var (
		namasaya = "faiz haidar"
		kelas = "1 smp"
	)

	fmt.Println(namasaya)
	fmt.Println(kelas)
}