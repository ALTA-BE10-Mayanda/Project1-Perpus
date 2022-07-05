package entity

import "gorm.io/gorm"

type Book struct {
	id        uint   `gorm:"primary_key:auto_increment"`
	user_id   uint 
	genre_id  uint
	title     string `gorm:"type:varchar(50)"`
	isbn      string `gorm:"type:varchar(13)"`
	author    string `gorm:"type:varchar(10)"`
	penerbit  string `gorm:"type:varchar(10)"`
	jumlah    uint   `gorm:"type:varchar(10)"`
	deskripsi string `gorm:"type:varchar(10)"`
}

type AksesBook struct {
	DB *gorm.DB
}