package entity

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	id       uint   `gorm:"primary_key:auto_increment"`
	nama     string `gorm:"type:varchar(100)"`
	address  string `gorm:"type:varchar(50)"`
	hp       string `gorm:"type:varchar(13)"`
	password string `gorm:"type:varchar(10)"`
}

type AksesUser struct {
	DB *gorm.DB
}

func (au *AksesUser) GetAllUser() []User {
	var daftarUser []User
	//err := au.DB.Raw("Select * from user").Scan(&daftarUser)
	err := au.DB.Find(&daftarUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarUser
}

func (au *AksesUser) AddNewUser(newUser User) User {

	if newUser.nama == "Levi" {
		newUser.id = uint(1)
	}
	uid := uuid.New()
	newUser.address = uid.String()
	newUser.hp = uid.String()
	err := au.DB.Create(&newUser).Error
	if err != nil {
		log.Println(err)
		return User{}
	}

	return newUser
}

func (au *AksesUser) GetSpecificUser(id int) User {
	var daftarUser = User{}
	daftarUser.id = uint(id)
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := au.DB.First(&daftarUser)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return User{}
	}

	return daftarUser
}

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