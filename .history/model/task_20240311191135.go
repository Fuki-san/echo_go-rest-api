package model

//taskを登録するときの構造体
type Task struct{
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title" gorm:"not null"`
	CreatedAt 
}
