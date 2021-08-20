package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	router := echo.New()
	router.GET("/", get)
	router.Logger.Fatal(router.Start("0.0.0.0:8000"))
}
