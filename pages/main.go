package pages

import (
	"net/http"

	"go-htmx-template/pages/components"
	"go-htmx-template/templ_components"
	"go-htmx-template/utils"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/home", HomePage)

	apiGroup := e.Group("/components")
	apiGroup.GET("/userInfo", components.GetUserInfo)

}

func HomePage(c echo.Context) error {
	return utils.Render(c, http.StatusOK, templ_components.HomePage("John"))
}
