package components

import (
	"go-htmx-template/model"
	"go-htmx-template/templ_components"
	"go-htmx-template/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserInfo(c echo.Context) error {

	res := model.User{
		Name:  "Wyndham",
		Phone: "8888888",
		Email: "skyscraper@gmail.com",
	}

	return utils.Render(c, http.StatusOK, templ_components.UserInfo(res))
}
