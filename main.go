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
	config.MigrateDB(conn)
	aksesUser := entity.AksesUser{DB: conn}
	var input int = 0
	var ListUser entity.User
	for input != 99 {
		fmt.Println("\tSistem peminjaman buku")
		fmt.Println("1. Tambah Data User")
		fmt.Println("2. Lihat Data User")
		fmt.Println("3. login")
		fmt.Println("4. Update Data User")
		fmt.Println("5. Hapus Data User")
		fmt.Println("99. Keluar")
		fmt.Print("Masukkan Pilihan menu: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var newUser entity.User
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan alamat: ")
			fmt.Scanln(&newUser.Address)
			fmt.Print("Masukkan nomor hp: ")
			fmt.Scanln(&newUser.Hp)
			fmt.Print("masukan pasword")
			fmt.Scanln(&newUser.Password)
			res := aksesUser.AddNewUser(newUser)
			if res.Nama == "" {
				fmt.Println("Tidak bisa input siswa, ada error")
				break
			}
			fmt.Println("Berhasl input siswa")
		case 3:
			var login entity.User
		  fmt.Print("nama: ")
		  fmt.Scanln(&ListUser.Nama)
		  fmt.Print("password: ")
		  fmt.Scanln(&ListUser.Password)
		  res := aksesUser.LoginUser(login.Nama, login.Password)
		  if res[0].ID == 0 {
			fmt.Println("tidak bisa login ")
			break
		  }
		  fmt.Println("berhasil login")
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