package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/specialagentsoftware/random_quote_service_go/quote_client"
)

func main() {
	/*server entry point instantiates a new echo web server object, and a quote_client object,
	and then binds a get request to the webroot to the formatted string provided by the Output function of the
	quote client*/
	e := echo.New()
	response := quote_client.QcInit()
	e.GET("/", func(c echo.Context) error {
		output := quote_client.Output(response)
		return c.HTML(http.StatusOK, output)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
