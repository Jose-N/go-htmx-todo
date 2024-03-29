package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Jose-N/go-htmx-todo/internal/model"
	"github.com/Jose-N/go-htmx-todo/internal/store/pgStore"
	"github.com/Jose-N/go-htmx-todo/internal/templates"
	"github.com/Jose-N/go-htmx-todo/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) GetLogin(c echo.Context) error {
	content := templates.Auth(true, "/auth/login", "Sign Up")
	title := "Esojist | A Jank Todo App"
	return util.RenderPage(c, templates.GenericLayout(content, title))
}

func (h *Handler) PostLogIn(c echo.Context) error {
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
	token := generateJWTClaims(expiresAt.Unix(), *user)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	c.SetCookie(generateCookie(tokenString))

	return c.String(http.StatusOK, fmt.Sprintf("Token String: %s", tokenString))
}

func (h *Handler) GetSignUp(c echo.Context) error {
	content := templates.Auth(false, "/auth/signup", "Sign In")
	title := "Esojist | A Jank Todo App"
	return util.RenderPage(c, templates.GenericLayout(content, title))
}

func (h *Handler) PostSignUp(c echo.Context) error {
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

func generateJWTClaims(unixExpireTime int64, user model.User) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        unixExpireTime,
		"authorized": true,
		"email":      user.Email,
		"userId":     user.ID,
	})

	return token
}

func generateCookie(tokenString string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Path = "/"
	cookie.Name = "Esojist"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = false

	return cookie
}
