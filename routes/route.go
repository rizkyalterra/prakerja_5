package routes

import (
	"prakerja5/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.GET("/news", controllers.NewsController)
	e.GET("/news/:id", controllers.DetailNewsController)
	e.POST("/news", controllers.AddNewsController)
	return e
}