package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Set GOMAXPROCS to the number of CPU cores available
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// fmt.Println(">>>runtime.NumCPU()", runtime.NumCPU())

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	e.Start(":8080")
}
