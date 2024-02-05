package pages

import (
	"net/http"

	"go-htmx-template/model"
	"go-htmx-template/templ_components"
	"go-htmx-template/utils"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", Home)
	e.GET("/home", HomePage)

	// apiGroup := e.Group("/components")
	// apiGroup.GET("/userInfo", components.GetUserInfo)

}

func Home(c echo.Context) error {
	return utils.Render(c, http.StatusOK, templ_components.Home())
}

func HomePage(c echo.Context) error {
	rating := model.Rating{
		Score:        4.0,
		TotalReviews: 10,
	}
	return utils.Render(c, http.StatusOK, templ_components.HomePage(rating))
}
