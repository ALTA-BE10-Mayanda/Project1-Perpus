package main

import (
	"fmt"
	"log"
	"perpus/config"
	"perpus/entity"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error: failed to load file env file")
	}

	conn := config.InitDB()
	aksesUser := entity.AksesUser{DB: conn}
	var input int = 0
	for input != 99 {
		fmt.Println("\tSistem peminjaman buku")
		fmt.Println("1. Tambah Data User")
		fmt.Println("2. Lihat Data User")
		fmt.Println("3. Update Data User")
		fmt.Println("4. Hapus Data User")
		fmt.Println("99. Keluar")
		fmt.Print("Masukkan Pilihan menu: ")
		fmt.Scanln(&input)

		switch input {
		// case 1:
		// 	var newUser entity.User
		// 	fmt.Print("Masukkan nama: ")
		// 	fmt.Scanln(&newUser.nama)
		// 	fmt.Print("Masukkan alamat: ")
		// 	fmt.Scanln(&newUser.address)
		// 	fmt.Print("Masukkan nomor hp: ")
		// 	fmt.Scanln(&newUser.hp)
			// res := aksesUser.AddNewUser(newUser)
			// // if res.nama == null {
			// // 	fmt.Println("Tidak bisa input siswa, ada error")
			// // 	break
			// // }
			// // fmt.Println("Berhasl input siswa")
		case 2:
			fmt.Println("Daftar Seluruh User")
			for _, val := range aksesUser.GetAllUser() {
				fmt.Println(val)
			}
		default:
			continue
		}
	}
	fmt.Println("Terima kasih sudah mencoba program saya")
}