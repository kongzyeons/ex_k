package repository

import "time"

type BearLog struct {
	FunctionName string    `json:"function_name" bson:"function_name"`
	Method       string    `json:"method" bson:"method"`
	Message      string    `json:"message" bson:"message"`
	Status       bool      `json:"status" bson:"status"`
	Timestamp    time.Time `json:"timestamp" bson:"timestamp"`
}
