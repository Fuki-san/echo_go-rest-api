package main

import (
	"fmt"
	"go-rest-api/db"
)


func main(){
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.
}
