package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SaveTodo(c echo.Context) error {
	todo := "You created a new Todo!"
	return c.String(http.StatusOK, todo)
}

func (h *Handler) GetTodos(c echo.Context) error {
	todos := "This is a list of Todos"
	return c.String(http.StatusOK, todos)
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
