package handler

import (
	"github.com/Jose-N/go-htmx-todo/internal/templates"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {
	content := templates.Auth()
	title := "Esojist | A Jank Todo App"
	return templates.AuthLayout(content, title).Render(c.Request().Context(), c.Response().Writer)
}
