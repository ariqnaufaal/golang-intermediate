package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func htmlfunc(ctx echo.Context) error {
	data :=
		`
	<html>
	<title>Session 5 - HTTP Payload via Echo Framework</title>
	<head>
		<h1>Hello</h1>
		<h2>This is HTML via Golang.</h2>
	</head>
	<body>
		<img src="https://cdn.dribbble.com/userupload/2624050/file/original-59266f4dea1c2aa43f2064cc0f3b165a.png?resize=400x0">
	</body>
	</html>
	`
	return ctx.HTML(http.StatusOK, data)
}
