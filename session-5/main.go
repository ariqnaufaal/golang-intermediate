package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func main() {
	// Echo instance
	r := echo.New()

	// Routes
	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	// Start server
	(r.Start(":9000"))
}
