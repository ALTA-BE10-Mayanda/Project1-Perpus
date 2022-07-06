package entity

import (
	"log"

	"gorm.io/gorm"
)
type User struct {gorm.Model

	Id       uint   `gorm:"primaryKey"`
	Nama     string `gorm:"type:varchar(100)"`
	Address  string `gorm:"type:varchar(50)"`
	Hp       string `gorm:"type:varchar(13)"`
	Password string `gorm:"type:varchar(10)"`
	Books []Book `gorm:"foreignKey:User_id"`
}

type AksesUser struct {
	DB *gorm.DB
}

func (au *AksesUser) GetAllUser() []User {
	var daftarUser []User
	// err := au.DB.Raw("Select * from user").Scan(&daftarUser)
	err := au.DB.Find(&daftarUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarUser
}

func (au *AksesUser) AddNewUser(newUser User) User {


	err := au.DB.Create(&newUser).Error
	if err != nil {
		log.Println(err)
		return User{}
	}

	return newUser
}

func (au *AksesUser) GetSpecificUser(id int) User {
	var daftarUser = User{}
	daftarUser.Id = uint(id)
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := au.DB.First(&daftarUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return User{}
	}

	return daftarUser
}

func (su *AksesUser) LoginUser(nama, password string) []User {
	var ListUser = []User{}
	if err := su.DB.Where("Nama = ? and Password = ?", nama, password).First(&ListUser).Error; err != nil {
	  log.Print(err)
	  return nil
	}
	return ListUser
  }

// func (as *UpdaeteUser) UpdateDataUser(IdUser uint) user {
// 	err := 	as.db.Model(&User{}).Where("id = ?", IdUser).Update(User{})
// 	if err != nil {
// 		log.Fatal(err.Error)
// 		return User{}
// 	}
// 	return User{}
// }

func (as *AksesUser) HapusMurid(IdUser int) bool {
	postExc := as.DB.Where("id = ?", IdUser).Delete(&User{})
	// ada masalah ga(?)
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	// berapa data yang berubah (?)
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}

	return true

}