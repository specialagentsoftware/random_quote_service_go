package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/specialagentsoftware/random_quote_service_go/quote_client"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		response := quote_client.Qc()
		return c.HTML(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
