package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hello struct {
	Name string
}

func main() {
	e := echo.New()
	e.GET("/hello", HelloController)
	e.Start(":8000")
}

func HelloController(c echo.Context) error {
	var hello = Hello{"Alterra"}
	return c.JSON(http.StatusOK, hello)
}




