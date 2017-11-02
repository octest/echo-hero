package handler

import (
	"bytes"
	"echo-hero/model"
	"echo-hero/template"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	buffer := new(bytes.Buffer)
	template.Index(buffer)
	return c.HTMLBlob(http.StatusOK, buffer.Bytes())
}

func Test2(c echo.Context) error {
	buffer := new(bytes.Buffer)

	account := model.Account{
		ID:   10,
		Name: "user one",
		/*Email: "user@one.cc",*/
	}

	notes := model.Notes{
		model.Note{
			ID:      1,
			Name:    "testing 1",
			Account: account,
		},
		model.Note{
			ID:      2,
			Name:    "testing 2",
			Account: account,
		},
	}

	validate := validator.New()
	err := validate.Struct(account)
	if err != nil {
		c.Error(err)
		/*if _, ok := err.(*validator.InvalidValidationError); ok {
			return c.HTML(http.StatusOK, err.Error())
		}*/
	}

	template.Test2("<i>other</i> message here", notes, buffer)
	return c.HTMLBlob(http.StatusOK, buffer.Bytes())
}
