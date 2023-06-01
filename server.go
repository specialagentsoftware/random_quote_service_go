package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/specialagentsoftware/random_quote_service_go/fileclient"
)

func main() {
	response := fileclient.Test()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
