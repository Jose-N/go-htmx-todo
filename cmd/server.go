package main

import (
	"github.com/Jose-N/go-htmx-todo/internal/handler"
	"github.com/Jose-N/go-htmx-todo/internal/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	h := &handler.Handler{}

	// Index
	e.GET("/", func(c echo.Context) error {
		content := templates.Index("World")
		return templates.Layout(content, "Index").Render(c.Request().Context(), c.Response().Writer)
	})

	//user routes
	e.POST("/users", h.SaveUser)
	e.GET("/user/:id", h.GetUser)
	e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("user/:id", h.DeleteUser)

	//todo routes
	e.POST("todos", h.SaveTodo)
	e.GET("/todos", h.GetTodos)
	e.GET("/todo/:id", h.GetTodo)
	e.PATCH("/todo/:id", h.UpdateTodo)
	e.DELETE("/todo/:id", h.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
