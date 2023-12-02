package api

import (
	v1 "go-htmx-template/api/v1"
	v2 "go-htmx-template/api/v2"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")

	v1.RegisterRoutes(apiGroup)
	v2.RegisterRoutes(apiGroup)

	apiGroup.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})

}
