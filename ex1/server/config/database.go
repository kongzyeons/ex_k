package config

import (
	"context"
	"fmt"
	"go_beer/repository"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseMariaDB() (db *gorm.DB) {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/Beer")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&repository.Beer{})
	fmt.Println("Connected to MariaDB")
	return db
}

func NewDatabaseMongoDB() *mongo.Database {
	connectionDB, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = connectionDB.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return connectionDB.Database("Beer")
}
