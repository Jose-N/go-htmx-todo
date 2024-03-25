package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jose-N/go-htmx-todo/internal/store/pgStore"
	"github.com/Jose-N/go-htmx-todo/internal/templates"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) SignUp(c echo.Context) error {
	content := templates.Auth(true, "/signup", "Sign Up")
	title := "Esojist | A Jank Todo App"
	return templates.GenericLayout(content, title).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) SignIn(c echo.Context) error {
	content := templates.Auth(false, "/signin", "Sign In")
	title := "Esojist | A Jank Todo App"
	return templates.GenericLayout(content, title).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) LogIn(c echo.Context) error {
	userStore := pgStore.UserStore{
		Db: h.Db,
	}

	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := userStore.FindByEmailAndPassword(email, password)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return c.String(http.StatusUnauthorized, "Invalid Credentials")
		} else {
			return err
		}
	}

	expiresAt := time.Now().Add(time.Minute * 15)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        expiresAt,
		"authorized": true,
		"email":      user.Email,
		"userId":     user.ID,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("Token String: %s", tokenString))
}
