package console

import (
	"fmt"
	"nule-istriku/lingkaran"
	"os"
)

func Console() {

	for {
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
		fmt.Print("\n")

	case 2:
		// Fitur luas Persegi

	case 3:
		//  Fitur Luas Segitiga

	case 4:
		os.Exit(0)
	}
	}
	

}

func menu() {
	fmt.Println("==================================")
	fmt.Println("     APLIKASI MENGHITUNG LUAS     ")
	fmt.Println("==================================")
	fmt.Println("1. Luas Lingkaran")
	fmt.Println("2. Luas Persegi")
	fmt.Println("3. Luas Segitiga")
	fmt.Println("4. Exit")
	fmt.Print("Masukkan opsi kamu : ")

}
