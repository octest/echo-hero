package template

import (
	"bytes"
)

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
