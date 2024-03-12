package main

import (
	"fmt"
	"go-rest-api/db"
	"go"
)


func main(){
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task)
}
