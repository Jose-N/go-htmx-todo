package main

import (
	"log"

	"github.com/Jose-N/go-htmx-todo/internal/handler"
	"github.com/Jose-N/go-htmx-todo/internal/store/pgStore"
	"github.com/Jose-N/go-htmx-todo/internal/templates"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	title := "Esojist | A Jank Todo App"
	db := pgStore.ConnectDB()
	h := &handler.Handler{
		Db: db,
	}
	e := echo.New()
	g := e.Group("/")

	e.Static("/static", "internal/static")
	g.Use(echojwt.WithConfig(echojwt.Config{
		TokenLookup: "cookie:Esojist",
		SigningKey:  []byte("secret"),
	}))

	// Open Routes
	e.GET("/", func(c echo.Context) error {
		content := templates.Index()
		return templates.MainLayout(content, title).Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/auth/login", h.GetLogin)
	e.POST("/auth/login", h.PostLogIn)
	e.GET("/auth/signup", h.GetSignUp)
	e.POST("/auth/signup", h.PostSignUp)

	// Restricted Routes
	g.GET("user", h.GetUser)
	g.PATCH("user", h.UpdateUser)
	g.DELETE("user", h.DeleteUser)

	g.POST("/todo", h.SaveTodo)
	g.GET("todo", h.GetTodos)
	e.GET("/todo/:id", h.GetTodo)
	e.PATCH("/todo/:id", h.UpdateTodo)
	e.DELETE("/todo/:id", h.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
