package main

import (
	"fmt"
)

func main() {
	checkSaldo()
	language()
	menu()

	fmt.Printf("Mau ngambil saldo berapa? ")
	var v2 int = Bank(100000, 50000)
	fmt.Scan(&v2)

	fmt.Println("Mencetak Struct......")

	if v2 == 50000 {
		fmt.Printf("Sisa saldo anda %v", v2)
	} else {
		fmt.Println("Uang anda tidak cukup")
	}

	fmt.Println("\nSilahkan ambil kembali kartu anda")
}

func Bank(saldo int, saldo1 int) int {
	result := saldo - saldo1
	return (result)
}

func checkSaldo() {
	fmt.Println("Selamat datang di Bank Mandiri")

	fmt.Printf("Silahkan masukkan Pin: ")
	var pin1 int
	fmt.Scan(&pin1)

	fmt.Printf("Correct! %v\n", pin1)
}

func language() {
	fmt.Print("Silahkan pilih Bahasa: (Indonesia / English) ")
	var name string
	fmt.Scan(&name)

	if name == "Indonesia" {
		fmt.Println("Lanjut")
	} else if name == "English" {
		fmt.Println("Next")
	} else {
		fmt.Println("Cannot din the language")
	}
}

func menu() {
	fmt.Println("--------LIST--------")
	var list = []string{
		"1. Check Saldo\n",
		"2. Tarik Tunai\n",
		"3. Exit\n",
	}

	fmt.Println(list)
	fmt.Printf("Choose: ")
	var newList string
	fmt.Scan(&newList)

	if newList == "1" {
		fmt.Println("Saldo anda 100.000")
	}

	fmt.Printf("<= Back to the List")
	var back string
	fmt.Scan(&back)

	if back == "ya" {
		fmt.Println(list)
		fmt.Print("Choose: ")
		var newList string
		fmt.Scan(&newList)
	}

}
