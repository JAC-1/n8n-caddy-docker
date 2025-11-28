package main

import (
	"time"

	"github.com/labstack/echo/v4"
	"go-app/views"
)

func main() {
	e := echo.New()

	// Home route
	e.GET("/", func(c echo.Context) error {
		return views.Home().Render(c.Request().Context(), c.Response().Writer)
	})

	// HTMX click handler
	e.GET("/click", func(c echo.Context) error {
		html := "<p>Hello from HTMX! Current time: " + time.Now().Format("2006-01-02 15:04:05") + "</p>"
		return c.HTML(200, html)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
