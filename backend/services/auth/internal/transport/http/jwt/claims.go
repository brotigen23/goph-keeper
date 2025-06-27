package jwt

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	ID int
	jwt.RegisteredClaims
}
