package main

import (
	"fmt"
	"go-htmx-template/api"
	"go-htmx-template/pages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))
	e.Use(middleware.Logger())
	e.Static("/dist", "dist")
	// Routes

	api.RegisterRoutes(e)
	pages.RegisterRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
