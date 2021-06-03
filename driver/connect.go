package driver

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //library gorm
)

// Connect ...
func Connect() *gorm.DB {
	con := os.Getenv("URL")
	db, err := gorm.Open("postgres", con)
	if err != nil {
		log.Fatal()
	}

	return db
}
