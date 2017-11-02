package model

import (
	"time"
)

type Note struct {
	ID      int       `json:"id" xml:"id" form:"id" query:"id"`
	Name    string    `json:"name" xml:"name" form:"name" query:"name"`
	Created time.Time `json:"created" xml:"created" form:"created" query:"created"`
	Account Account   `json:"account" xml:"account" form:"account" query:"account"`
}

type Notes []Note
