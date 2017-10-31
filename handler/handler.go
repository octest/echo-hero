package handler

import (
	"bytes"
	"net/http"

	"github.com/labstack/echo"
)

/*func Header(c echo.Context) {
c.Response().

}*/

//var buffer *bytes.Buffer

func Header(buff *bytes.Buffer) {
	buff.WriteString(`<html>
<head>
	<title>Title</title>
	<link rel='stylesheet' href='/static/stylesheet.css'>
</head>
<body>`)
}

func Footer(buff *bytes.Buffer) {
	buff.WriteString(`
	<hr size='1'/>
	<div align='center'>&copy; 2017</div>
</body>
</html>`)
}

func Hello(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	//return c.HTML(http.StatusOK, "Hello, <b>World!</b>")

	buffer := new(bytes.Buffer)
	//buffer := &bytes.Buffer{}
	//buffer := bytes.NewBuffer(nil)

	Header(buffer)
	buffer.WriteString("Hello, <b>World!</b> (blob)")
	buffer.WriteString("<br/>Another line [3]")
	//buffer.WriteString("")
	Footer(buffer)
	return c.HTMLBlob(http.StatusOK, buffer.Bytes())
}
