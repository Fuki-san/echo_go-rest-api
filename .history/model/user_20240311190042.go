package model

import "time"

type User struct{
	ID unit `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique`
	Password string `json:"password`
	CreatedAt time.Time `json`
}
