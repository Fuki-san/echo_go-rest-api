package db

//gorm is an external package, so we need to install it by go mod tidy
import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
	//for using .env file
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)
// we need to connect to the database by url so we define the url as an environment variable
// environment variables means that we can change the value of the variable without changing the code
// so we can use the same code to connect to different databases by changing the environment variables
//1. prepare environment variables for 
//2. 
//3. 

//return db address to connect to the database by gorm package
func NewDB() *gorm.DB{
	//prepare environment variables
	if os.Getenv("GO_ENV") == "dev" {
		//load local .env file if GO_ENV is dev
		err := godotenv.Load()
		if err != nil{
		//if .env file is not found, log the error and exit
			log.Fatalln(err)
		}
	}
	//url to connect to the database
	// %s is replaced by Getenv's argument like "POSTGRES_USER",...
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println("Connected to the database")
	return db
}
