package controllers

import (
	"net/http"
	"prakerja5/configs"
	"prakerja5/models"
	"strconv"

	"github.com/labstack/echo/v4"
)


func AddNewsController(c echo.Context) error {
	var news models.News
	c.Bind(&news)

	result := configs.DB.Create(&news)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data: news,
	})
}

func DetailNewsController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id")) 

	var data = models.News{Id: id, Title: "A", Content: "B"}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data: data,
	})
}


func NewsController(c echo.Context) error {
	var data []models.News

	result := configs.DB.Find(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data: data,
	})
}