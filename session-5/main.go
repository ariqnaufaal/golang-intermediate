package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

type User struct {
	Name  string `json:"name" form:"name" query:"name" `
	Email string `json:"email" form:"email" query:"email" `
	Age   int    `json:"age" form:"age" query:"age" `
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action home"))
	},
)

var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("from action about"))
		},
	),
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Echo instance
	r := echo.New()
	r.Validator = &CustomValidator{validator: validator.New()}

	// Routes
	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/index", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	r.GET("/html", htmlfunc)

	r.GET("/json", func(ctx echo.Context) error {
		data := M{"Message": "Hello", "Counter": 2}
		return ctx.JSON(http.StatusOK, data)
	})

	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)

		return ctx.String(http.StatusOK, data)
	})

	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf(
			"Hello %s, I have message for you: %s",
			name,
			strings.Replace(message, "/", "", 1),
		)

		return ctx.String(http.StatusOK, data)
	})

	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	r.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}

		if err := c.Validate(u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, true)
	})

	r.Static("/static", "assets")

	r.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok = err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// if castedObject, ok := err.(validator.ValidationErrors); ok {
		// 	for _, err := range castedObject {
		// 		switch err.Tag() {
		// 		case "Required":
		// 			report.Message = fmt.Sprintf("%s is required", err.Field())
		// 		case "email":
		// 			report.Message = fmt.Sprintf("%s is not valid email", err.Field())
		// 		case "gte":
		// 			report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
		// 		case "lte":
		// 			report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
		// 		}

		// 		break
		// 	}
		// }

		// c.Logger().Error(report)
		// c.JSON(report.Code, report)

		errPage := fmt.Sprintf("%d.html", report.Code)
		if err := c.File(err.Page); err != nil {
			c.HTML(report.Code, "Errrroooorrrr")
		}
	}

	// Start server
	(r.Start(":9000"))
}
