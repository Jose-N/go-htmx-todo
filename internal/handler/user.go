package handler

import (
	"golang.org/x/crypto/bcrypt"

	"fmt"
	"net/http"
	"strconv"

	"github.com/Jose-N/go-htmx-todo/internal/store/pgStore"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SaveUser(c echo.Context) error {
	userStore := pgStore.UserStore{
		Db: h.Db,
	}

	firstName := c.FormValue("first-name")
	lastName := c.FormValue("last-name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	id, err := userStore.CreateUser(firstName, lastName, email, string(pass))
	if err != nil {
		return err
	}

	stringId := strconv.FormatUint(uint64(*id), 10)
	return c.String(http.StatusOK, stringId)
}

func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("You are user: %s", id))
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("You are updated: %s", id))
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("You deleted: %s", id))
}
