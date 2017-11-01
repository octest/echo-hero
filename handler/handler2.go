package handler

import (
	"bytes"
	"echo-hero/template"
	"net/http"

	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	buffer := new(bytes.Buffer)
	template.Index(buffer)
	return c.HTMLBlob(http.StatusOK, buffer.Bytes())
}

func Test2(c echo.Context) error {
	buffer := new(bytes.Buffer)
	template.Test2("<i>other</i> message here", buffer)
	return c.HTMLBlob(http.StatusOK, buffer.Bytes())
}
