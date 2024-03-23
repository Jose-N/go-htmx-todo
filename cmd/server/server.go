package main

import (
	"log"

	"github.com/Jose-N/go-htmx-todo/internal/handler"
	"github.com/Jose-N/go-htmx-todo/internal/store/pgStore"
	"github.com/Jose-N/go-htmx-todo/internal/templates"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	e := echo.New()
	h := &handler.Handler{}
	db := pgStore.ConnectDB()
	_ = db

	e.Static("/static", "internal/static")

	// Index
	e.GET("/", func(c echo.Context) error {
		content := templates.Index()
		title := "Esojist | A Jank Todo App"
		return templates.Layout(content, title).Render(c.Request().Context(), c.Response().Writer)
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
