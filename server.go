package main

import (
	"net/http"

	"github.com/Jose-N/go-htmx-todo/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	h := &handler.Handler{}

	// Index
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	//user routes
	e.POST("/users", h.SaveUser)
	e.GET("/user/:id", h.GetUser)
	e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("user/:id", h.DeleteUser)

	//todo routes

	e.Logger.Fatal(e.Start(":8080"))
}
