package main

import (
	"fmt"
	"log"
	"perpus/config"
	"perpus/controller"
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
	aksesUser := controller.AksesUser{DB: conn}
	aksesBook := controller.AksesBook{DB: conn}
	aksesGenre := controller.AksesGenre{DB: conn}
	aksesRent := controller.AksesRent{DB: conn}
	var hp string
	var id uint
	var input = false
	var log int 
	for !input && log != 5 {
		fmt.Println("\tSistem peminjaman buku")
		fmt.Println("1. Tambah Data User")
		fmt.Println("2. Login")
		fmt.Println("3. Lihat Daftar Buku")
		fmt.Println("4. Cari Buku")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&log)
		fmt.Print("\n")

		switch log{
		case 1:
			var newUser entity.User
			fmt.Println("\tRegistrasi")
			fmt.Print("Masukkan nama:")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan alamat:")
			fmt.Scanln(&newUser.Address)
			fmt.Print("Masukkan nomor hp:")
			fmt.Scanln(&newUser.Hp)
			fmt.Print("Masukan pasword:")
			fmt.Scanln(&newUser.Password)
			res := aksesUser.Register(newUser)
			if res.Nama == "" {
				fmt.Println("Tidak bisa input user, ada error")
				break
			} else {
				fmt.Println("Berhasil input user")
			}
			fmt.Print("\n")
		case 2:
			var login entity.User
		  fmt.Println("\tLogin")
		  fmt.Print("Nomor Hp:")
		  fmt.Scanln(&login.Hp)
		  fmt.Print("Password:")
		  fmt.Scanln(&login.Password)
		  log := aksesUser.Login(login.Hp, login.Password)
		  if log.Hp == "" && log.Password == ""  {
			fmt.Println("tidak bisa login ")
			break
		  } else {
			fmt.Println("berhasil login")
			hp = log.Hp
			id = log.Id
		  	input = true
			fmt.Println("Id User:", log.Id)
		  }
		  fmt.Print("\n")
		case 3:
			fmt.Println("\tDaftar Buku")
			var allBook = aksesBook.GetAllBook()
			for _, i := range allBook {
				fmt.Println("======================")
				fmt.Printf("Id        : %d \n", i.Id)
				fmt.Printf("Id user   : %d \n", i.User_id)
				fmt.Printf("Id genre  : %d \n", i.Genre_id)
				fmt.Printf("Judul     : %s \n", i.Title)
				fmt.Printf("Isbn      : %s \n", i.Isbn)
				fmt.Printf("Pengarang : %s \n", i.Author)
				fmt.Printf("Penerbit  : %s \n", i.Penerbit)
				fmt.Printf("Jumlah    : %d \n", i.Jumlah)
				fmt.Printf("Deskripsi : %s \n", i.Deskripsi)
			}
			fmt.Print("\n")
		case 4:
			var searchBook entity.Book
			fmt.Println("\tCari Buku")
			fmt.Print("Masukkan Judul Buku yang akan dicari : ")
			fmt.Scanln(&searchBook.Title)
			search := aksesBook.GetBookbyName(searchBook.Title)
			if search.Title == "" {
				fmt.Println("Buku tidak tersedia")
				break
			  } else {
				fmt.Println("Buku ditemukan")
			  }
			fmt.Print("\n")
		}
	}

	for input {
		var pilih int
		fmt.Println("\t    SELAMAT DATANG")
		fmt.Println("\t Silahkan Pilih Menu")
		fmt.Println("1.  Daftar User")
		fmt.Println("2.  Cari User")
		fmt.Println("3.  Update User")
		fmt.Println("4.  Hapus User")
		fmt.Println("5.  Tambah Genre Buku")
		fmt.Println("6.  Daftar Genre Buku")
		fmt.Println("7.  Cari Genre Buku")
		fmt.Println("8.  Hapus Genre Buku")
		fmt.Println("9.  Tambah Buku")
		fmt.Println("10. Edit Buku")
		fmt.Println("11. Hapus Buku")
		fmt.Println("12. Pinjam Buku")
		fmt.Println("13. Daftar Buku Yang Dipinjam")
		fmt.Println("14. Kembalikan Buku")
		fmt.Println("15. Exit")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&pilih)
		fmt.Print("\n")

		switch pilih {
		case 1:
			fmt.Println("Daftar Seluruh User")
			var allUser = aksesUser.GetAllUser()
			for _, v := range allUser {
				fmt.Printf("Id : %d \n", v.Id)
				fmt.Printf("Nama : %s \n", v.Nama)
				fmt.Printf("Alamat : %s \n", v.Address)
				fmt.Printf("HP : %s \n", v.Hp)
			}
			fmt.Print("\n")
		case 2:
			var searchUser entity.User
			fmt.Println("Cari User")
			fmt.Print("Masukkan Nama User :")
		    fmt.Scanln(&searchUser.Nama)
			search := aksesUser.GetUserbyName(searchUser.Nama)
		    if search.Nama == "" {
			fmt.Println("User tidak tersedia")
			break
		  } else {
			fmt.Println("User ditemukan")
		  }
			fmt.Print("\n")
		case 3:
			var update int =  0
			var dataUpdate entity.User
			for update != 5 {
				var pil int
				fmt.Println("\t Pilih Update")
				fmt.Println("1. Nama")
				fmt.Println("2. Alamat")
				fmt.Println("3. Nomor Hp")
				fmt.Println("4. Password")
				fmt.Println("5. Keluar")
				fmt.Print("Pilih Update :")
				fmt.Scanln(&pil)
				
				switch pil {
				case 1:
					var updateNama string
			        fmt.Println("Update User")
			        fmt.Print("Ubah Nama :")
		            fmt.Scanln(&updateNama)
		            if updateNama == "" {
			        fmt.Println("Gagal update nama user")
			        break
		            } else {
					   dataUpdate.Nama = updateNama
					   cond := aksesUser.UpdateDataUser(hp, dataUpdate)
					   if cond  {
						fmt.Println("Nama user telah diupdate")
					   }
		            }
			        fmt.Print("\n")
				case 2:
					var updateAlamat string
			        fmt.Println("Update User")
			        fmt.Print("Ubah Alamat :")
		            fmt.Scanln(&updateAlamat)
		            if updateAlamat == "" {
			        fmt.Println("Gagal update alamat user")
			        break
		            } else {
						dataUpdate.Address = updateAlamat
						cond := aksesUser.UpdateDataUser(hp, dataUpdate)
						if cond  {
						 fmt.Println("Alamat user telah diupdate")
						}
		            }
			        fmt.Print("\n")
				case 3:
					var updateHp string
			        fmt.Println("Update User")
			        fmt.Print("Ubah Nomor HP :")
		            fmt.Scanln(&updateHp)
		            if updateHp == "" {
			        fmt.Println("Gagal update nomor hp user")
			        break
		            } else {
						dataUpdate.Hp = updateHp
						cond := aksesUser.UpdateDataUser(hp, dataUpdate)
						if cond  {
						 fmt.Println("Nomor HP user telah diupdate")
						}
		            }
			        fmt.Print("\n")
				case 4:
					var updatePassword string
			        fmt.Println("Update User")
			        fmt.Print("Ubah Password :")
		            fmt.Scanln(&updatePassword)
		            if updatePassword == "" {
			        fmt.Println("Gagal update password user")
			        break
		            } else {
						dataUpdate.Password = updatePassword
						cond := aksesUser.UpdateDataUser(hp, dataUpdate)
						if cond  {
						 fmt.Println("Password user telah diupdate")
						}
		            }
			        fmt.Print("\n")
				case 5:
					update = 5
				}
			}
				fmt.Print("\n")
		case 4:
			var UserID int
			fmt.Print("Masukkan ID yang akan dihapus :")
			fmt.Scanln(&UserID)
			fmt.Println(aksesUser.DeleteUser(UserID))
			fmt.Print("\n")
		case 5:
			var newGenre entity.Genre
			fmt.Print("Masukkan Nama Genre : ")
			fmt.Scanln(&newGenre.Nama)
			res := aksesGenre.AddGenre(newGenre)
			if res.Nama == "" {
				fmt.Println("Tidak bisa tambah genre, ada error")
				break
			} else {
				fmt.Println("Berhasil tambah genre")
			}
			fmt.Print("\n")
		case 6:
			fmt.Println("Daftar Seluruh Genre Buku")
			var allGenre = aksesGenre.GetAllGenre()
			for _, v := range allGenre {
				fmt.Printf("Id : %d \n", v.Id)
				fmt.Printf("Nama : %s \n", v.Nama)
			}
			fmt.Print("\n")
		case 7:
			var searchGenre entity.Genre
			fmt.Println("Cari Genre Buku")
			fmt.Print("Masukkan Nama Genre Buku :")
		    fmt.Scanln(&searchGenre.Nama)
			search := aksesGenre.GetGenrebyName(searchGenre.Nama)
		    if search.Nama == "" {
			fmt.Println("Genre Buku tidak tersedia")
			break
		  } else {
			fmt.Println("Genre Buku ditemukan")
		  }
			fmt.Print("\n")
		case 8:
			var GenreID int
			fmt.Print("Masukkan ID yang akan dihapus :")
			fmt.Scanln(&GenreID)
			fmt.Println(aksesGenre.DeleteGenre(GenreID))
			fmt.Print("\n")
		case 9:
			var newBook entity.Book
			fmt.Println("\t Tambah Buku")
			fmt.Print("Masukkan id user:")
			fmt.Scanln(&newBook.User_id)
			fmt.Print("Masukkan id genre:")
			fmt.Scanln(&newBook.Genre_id)
			fmt.Print("Masukkan judul:")
			fmt.Scanln(&newBook.Title)
			fmt.Print("Masukkan isbn:")
			fmt.Scanln(&newBook.Isbn)
			fmt.Print("Masukkan nama pengarang:")
			fmt.Scanln(&newBook.Author)
			fmt.Print("Masukan nama penerbit:")
			fmt.Scanln(&newBook.Penerbit)
			fmt.Print("Masukan jumlah:")
			fmt.Scanln(&newBook.Jumlah)
			fmt.Print("Masukan deskripsi:")
			fmt.Scanln(&newBook.Deskripsi)
			res := aksesBook.AddNewBook(newBook)
			if res.Title == "" {
				fmt.Println("Tidak bisa tambah buku, ada error")
				break
			} else {
					fmt.Println("Berhasil tambah buku")
			
			}
			fmt.Print("\n")
		case 10:
			var updateBook entity.Book 
			fmt.Println("Update Buku")
			fmt.Print("Id buku:")
			fmt.Scanln(&id)
			fmt.Print("Ubah genre:")
		    fmt.Scanln(&updateBook.Genre_id)
			fmt.Print("Ubah judul:")
		    fmt.Scanln(&updateBook.Title)
			fmt.Print("Ubah isbn:")
		    fmt.Scanln(&updateBook.Isbn)
			fmt.Print("Ubah nama pengarang:")
		    fmt.Scanln(&updateBook.Author)
			fmt.Print("Ubah nama penerbit:")
		    fmt.Scanln(&updateBook.Penerbit)
			fmt.Print("Ubah jumlah:")
		    fmt.Scanln(&updateBook.Jumlah)
			fmt.Print("Ubah deskripsi:")
		    fmt.Scanln(&updateBook.Deskripsi)
			res := aksesBook.UpdateDataBook(id, updateBook)
		    if !res {
			    fmt.Println("Gagal update data buku")
			    break
		    } else {
				  	fmt.Println("Data buku telah diupdate")
				}
			fmt.Print("\n")
		case 11:
			var BookID uint
			fmt.Print("Masukkan ID yang akan dihapus :")
			fmt.Scanln(&BookID)
			fmt.Println(aksesBook.DeleteBook(BookID))
			fmt.Print("\n")
		case 12:
			var rentBook entity.Rent
			fmt.Println("Pinjam Buku")
			fmt.Print("Id user:")
			fmt.Scanln(&rentBook.User_id)
			fmt.Print("Id buku:")
		    fmt.Scanln(&rentBook.Book_id)
			res := aksesRent.AddNewRent(rentBook)
		    if res.User_id == 0 {
			    fmt.Println("Gagal pinjam buku")
			    break
		    } else {
				  	fmt.Println("Berhasil pinjam buku")
				}
			fmt.Print("\n")
		case 14:
			var IdBook uint
			fmt.Print("Kembalikan buku:")
			fmt.Scanln(&IdBook)
			fmt.Println(aksesRent.ReturnBook(IdBook))
			fmt.Print("\n")
		case 13:
			var listBook entity.Rent
			fmt.Println("Daftar Buku Yang Dipinjam")
			fmt.Print("Id user:")
			fmt.Scanln(&listBook.User_id)
			var allBook = aksesRent.GetBookbyUserID(listBook.User_id)
			for _, b := range allBook {
				fmt.Printf("Id buku: %d \n", b.Book_id)
				fmt.Println("======================")
			}
			fmt.Print("\n")
		case 15:
			fmt.Println("Terima Kasih telah mencoba program ini")
			input = false
		}
	}
	fmt.Println("Terima kasih")
}