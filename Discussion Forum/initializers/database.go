package initializers

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import the MySQL driver
)

var DB *gorm.DB

func InitDB() {

	var err error
	// Replace with my MySQL database connection parameters.
	dbURI := "root:12345@tcp(localhost:3306)/forum?charset=utf8&parseTime=True&loc=Local" // connection link for the Mysql Workbench

	// Open a connection to the database.
	DB, err = gorm.Open("mysql", dbURI)

	if err != nil {
		log.Fatal(err)
	}

}
