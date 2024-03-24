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

	title := "Esojist | A Jank Todo App"

	db := pgStore.ConnectDB()
	h := &handler.Handler{
		Db: db,
	}

	e := echo.New()

	e.Static("/static", "internal/static")

	// Index
	e.GET("/", func(c echo.Context) error {
		content := templates.Index()
		return templates.MainLayout(content, title).Render(c.Request().Context(), c.Response().Writer)
	})

	//auth routes
	e.GET("/signup", h.SignUp)
	e.POST("/signup", h.SaveUser)

	//user routes
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
