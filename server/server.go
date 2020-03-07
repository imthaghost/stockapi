package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/imthaghost/stockapi/controllers"
)

// StockAPI structur
type StockAPI struct {
	e *echo.Echo
}

// NewServer Instance of Echo
func NewServer() *StockAPI {

	return &StockAPI{
		e: echo.New(),
	}
}

// Start server funcitonality
func (s *StockAPI) Start(port string) {
	// logger
	s.e.Use(middleware.Logger())
	// recover
	s.e.Use(middleware.Recover())
	//CORS
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// price endpoint
	s.e.POST("/price", controllers.GetPrice)
	// Start Server
	s.e.Logger.Fatal(s.e.Start(port))
}

// Close server functionality
func (s *StockAPI) Close() {
	s.e.Close()
}
