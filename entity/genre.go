package entity

type Genre struct {
	id       uint    `gorm:"primary_key:auto_increment"`
	nama     string  `gorm:"type:varchar(20)"`
}