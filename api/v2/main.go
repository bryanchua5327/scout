package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(parent *echo.Group) {
	apiGroup := parent.Group("/v2")
	apiGroup.GET("/:path", func(c echo.Context) error {
		path := c.Param("path")
		return c.String(http.StatusOK, path)
	})
}
