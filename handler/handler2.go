package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	return c.String(http.StatusOK, "New test")
}
