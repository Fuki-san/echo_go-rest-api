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
//1. prepare environment variables for rules of connection to the database
//2. load the environment variables from .env file if GO_ENV is dev
//3. prepare the url to connect to the database and get the database address by gorm package
//4. return the database address to connect to the database

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
//close the connection to the database
//*gorm.DB is an big abject that enables abstract database operations
func CloseDB(db *gorm.DB){
	//DB() get the underlying *sql.DB object that manages the connection pool
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil{
		log.Fatalln(err)
	}
}
