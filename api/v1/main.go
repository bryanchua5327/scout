package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(context echo.Context) error {

	return context.String(http.StatusOK, "hello")
}
func byebye(context echo.Context) error {

	return context.String(http.StatusOK, "bye bye")
}

func getInfo(context echo.Context) error {

	res := map[string]interface{}{
		"Name":  "Wyndham",
		"Phone": "8888888",
		"Email": "skyscraper@gmail.com",
	}
	return context.Render(http.StatusOK, "name_card", res)
}

func RegisterRoutes(parent *echo.Group) {
	apiGroup := parent.Group("/v1")

	apiGroup.GET("/hello", hello)
	apiGroup.GET("/byebye", byebye)
	apiGroup.GET("/get-info", getInfo)

}
