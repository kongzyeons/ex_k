package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepo interface {
	CreateLog(bearLog BearLog) (err error)
}

type logRepo struct {
	db *mongo.Database
}

func NewLogRepo(db *mongo.Database) LogRepo {
	return logRepo{db}
}

func (obj logRepo) CreateLog(bearLog BearLog) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = obj.db.Collection("log").InsertOne(ctx, bearLog)
	return err
}
