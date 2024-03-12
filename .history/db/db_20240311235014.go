package db
//gorm is an external package, so we need to install it by go mod tidy
import (
	"gorm.io/gorm"
	//for using .env file
	"github.com/joho/godotenv"
)


//return db address to connect to the database by gorm package
func NewDB() *gorm.DB{
	if os.
}
