package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=bookstore sslmode=disable")
	fmt.Println("CAME HERE 1")

	if err != nil {
		panic(err)
	}
	fmt.Println("CAME HERE 3")
	db = d
	fmt.Println("CAME HERE 4")
}

func GetDB() *gorm.DB {
	fmt.Println("CAME HERE 5")
	return db
}
