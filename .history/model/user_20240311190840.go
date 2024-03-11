package model

import "time"

type User struct{
	//`json:"○○"`で、構造体をjsonに渡した際に、
	//フィールド名を○○に変換することができる
	ID uint `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
//ユーザー登録時のリクエストに対して、
//クライアントに返すレスポンスの構造体
type UserResponse struct{
	ID uint `json:"id"`
	Email string `json:"email"`
}
