package main

import (
	"prakerja5/configs"
	"prakerja5/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDB() 
	e := echo.New()
	e = routes.InitRoute(e)
	e.Start(":8000")
}




