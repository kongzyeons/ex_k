package models

import (
	"mime/multipart"
)

type CreateBeerRequest struct {
	FileImg  *multipart.FileHeader `json:"file_img"`
	Name     string                `json:"name"`
	Category string                `json:"category"`
	Detail   string                `json:"detail"`
}

type GetBeerRequest struct {
	Name string `json:"name"`
}

type GetPaginatBeerRequest struct {
	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_number"`
}

type UpdateBeerRequest struct {
	FileImg  *multipart.FileHeader `json:"file_img"`
	Name     string                `json:"name"`
	Category string                `json:"category"`
	Detail   string                `json:"detail"`
}

type CreateLogRequst struct {
	FunctionName string `json:"function_name" bson:"function_name"`
	Method       string `json:"method" bson:"method"`
	Message      string `json:"message" bson:"message"`
	Status       bool   `json:"status" bson:"status"`
}
