package models

type BaseResponse struct {
	Message string 	`json:"message"`
	Data interface{}`json:"data"`
}