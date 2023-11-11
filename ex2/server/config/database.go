package config

import (
	"fmt"
	"go_admin/repository"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseMariaDB() (db *gorm.DB) {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/Admin")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&repository.Parent{}, &repository.Icon{}, &repository.Children{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MariaDB")
	return db
}
