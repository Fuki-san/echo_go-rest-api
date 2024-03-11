//user.goは、modelふぁう
package model

import "time"

type User struct{
	//`json:"○○"`で、構造体をjsonに渡した際に、
	//フィールド名を○○に変換することができる
	ID unit `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique`
	Password string `json:"password`
	CreatedAt time.Time `json:"created_at`
	UpdatedAt time.Time `json:"updated_at`
}
