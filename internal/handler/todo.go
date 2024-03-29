package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SaveTodo(c echo.Context) error {
	todo := "You created a new Todo!"
	return c.String(http.StatusOK, todo)
}

func (h *Handler) GetTodos(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("failed to cast claims")
	}
	return c.JSON(http.StatusOK, claims)

	//content := templates.Main("jose", "123")
	//return templates.GenericLayout(content, "title").Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) GetTodo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("This is the Todo with id: %s", id))
}

func (h *Handler) UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("Updated Todo with id: %s", id))
}

func (h *Handler) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("Delete Todo with id: %s", id))
}
