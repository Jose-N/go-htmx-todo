package util

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderPage(c echo.Context, page templ.Component) error {
	return page.Render(c.Request().Context(), c.Response().Writer)
}
