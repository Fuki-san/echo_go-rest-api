package model

import "time"
//user data for s
type User struct{
	//`json:"○○"`で、構造体をjsonに渡した際に、
	//フィールド名を○○に変換することができる
	ID uint `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
//user data for response
type UserResponse struct{
	ID uint `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}