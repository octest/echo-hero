package handler

import (
	"echo-hero/template"
	"net/http"

	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	return c.HTMLBlob(http.StatusOK, template.Index())
}
