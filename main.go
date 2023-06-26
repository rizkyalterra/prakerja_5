package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type News struct {
	Id int 			`json:"id"`
	Title string 	`json:"title"`
	Content string 	`json:"content"`
}

type BaseResponse struct {
	Message string 	`json:"message"`
	Data interface{}`json:"data"`
}

func main() {
	initDB()
	e := echo.New()
	e.GET("/news", NewsController)
	e.GET("/news/:id", DetailNewsController)
	// e.POST("/news", AddNewsController)
	e.Start(":8000")
}

func initDB(){
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja5?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database error")
	}
	migration()
}

func migration(){
	DB.AutoMigrate(&News{})
}


func AddNewsController(c echo.Context) error {
	var news News
	c.Bind(&news)

	// validsi

	result := DB.Create(&news)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Message: "Error",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Message: "Success",
		Data: news,
	})
}

func DetailNewsController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id")) 

	var data = News{id, "A", "B"}

	return c.JSON(http.StatusOK, BaseResponse{
		Message: "Success",
		Data: data,
	})
}


func NewsController(c echo.Context) error {
	var data []News

	result := DB.Find(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Message: "Error",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Message: "Success",
		Data: data,
	})
}




