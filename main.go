package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
	e := echo.New()
	e.GET("/news", NewsController)
	e.GET("/news/:id", DetailNewsController)
	e.POST("/news", AddNewsController)
	e.Start(":8000")
}


func AddNewsController(c echo.Context) error {
	var news News
	c.Bind(&news)

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
	var negara = c.QueryParam("negara")
	var sort = c.QueryParam("sort")

	var data []News
	data = append(data, News{1, negara, "B"})
	data = append(data, News{2, sort, "C"})
	return c.JSON(http.StatusOK, BaseResponse{
		Message: "Success",
		Data: data,
	})
}




