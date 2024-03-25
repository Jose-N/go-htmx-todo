package model

import "github.com/golang-jwt/jwt/v5"

type Auth struct {
	UserId         uint
	Email          string
	StandardClaims jwt.RegisteredClaims
}
