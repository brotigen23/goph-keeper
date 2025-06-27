package http

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	ID int
	jwt.Claims
}
