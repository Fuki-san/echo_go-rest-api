package db

//gorm is an external package, so we need to install it by go mod tidy
import (
	"log"
	"os"

	"gorm.io/gorm"
	//for using .env file
	"github.com/joho/godotenv"
)

//return db address to connect to the database by gorm package
func NewDB() *gorm.DB{
	//prepare environment variables
	if os.Getenv("GO_ENV") == "dev" {
		//load local .env file if GO_ENV is dev
		err := godotenv.Load()
		if err != nil{
		//if .env
			log.Fatalln(err)
		}
	}
}
