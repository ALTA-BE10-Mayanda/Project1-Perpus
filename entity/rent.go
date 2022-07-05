package entity

import "time"

type Rent struct {
	id          uint `gorm:"primary_key:auto_increment"`
	user_id     uint 
	book_id     uint
	created_at  time.Time  `gorm:"autoCreateTime"`
	return_date time.Time
}