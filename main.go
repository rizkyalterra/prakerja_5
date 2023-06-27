package main

import (
	"prakerja5/configs"
	"prakerja5/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	configs.InitDB() 
	e := echo.New()
	e = routes.InitRoute(e)
	e.Start(":8000")
}

func loadEnv(){
	err := godotenv.Load()
  	if err != nil {
    	panic("Failed load file env")
  	}
}




