package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	//user routes
	// create
	e.POST("/users", saveUser)
	// read
	e.GET("/user/:id", getUser)
	// update
	e.PATCH("/user/:id", updateUser)
	// delete
	e.DELETE("user/:id", deleteUser)

	//todo routes

	e.Logger.Fatal(e.Start(":8080"))
}

func saveUser(c echo.Context) error {
	user := "You created a new user!"
	return c.String(http.StatusOK, user)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("You are user: %s", id))
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("You are updated: %s", id))
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("You deleted: %s", id))
}
