package main

import (
	"os"
	"prakerja5/configs"
	"prakerja5/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// loadEnv()
	configs.InitDB() 
	e := echo.New()
	e = routes.InitRoute(e)
	e.Start(":"+getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return port
}

// func loadEnv(){
// 	err := godotenv.Load()
//   	if err != nil {
//     	panic("Failed load file env")
//   	}
// }




