package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Home struct {
	Name string `json:"name" form:"name" query:"name"`
	Date string `json:"date" form:"date" query:"date"`
}

func RegisterRoutes(e *echo.Echo) {

	e.GET("/home", func(c echo.Context) error {
		u := new(Home)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		return c.Render(http.StatusOK, "home", u)
	})

}
