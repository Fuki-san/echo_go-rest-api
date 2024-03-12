//main package create the executable file that should be run by the firest command
package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)


func main(){
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	//AutoMigrate() create the table if it does not exit
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
y
