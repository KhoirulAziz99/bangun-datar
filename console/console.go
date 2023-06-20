package console

import (
	"fmt"
	"nule-istriku/lingkaran"
)

func Console() {
	menu()
	var selected int
	fmt.Scanln(&selected)
	switch selected {
	case 1:
		fmt.Print("Masukkan jari-jari : ")
		var r int
		fmt.Scanln(&r)
		hasil := lingkaran.LuasLingkaran(r)
		fmt.Println("Jadi luas lingkaran anda adalah : ", hasil)
	}

}

func menu() {
	fmt.Println("=====APLIKASI MENGHITUNG LUAS====")
	fmt.Println("1. Luas Lingkaran")
	fmt.Println("2. Luas Persegi")
	fmt.Println("3. Luas Segitiga")
	fmt.Println("4. Exit")
	fmt.Print("Masukkan opsi kamu : ")

}
