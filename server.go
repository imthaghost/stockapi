package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/imthaghost/stockapi/controllers"
)

func main() {

	e := echo.New()
	// logger
	e.Use(middleware.Logger())
	// recover
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// price endpoint
	e.POST("/price", controllers.GetPrice)
	// Start Server
	e.Logger.Fatal(e.Start(":8000"))

}
