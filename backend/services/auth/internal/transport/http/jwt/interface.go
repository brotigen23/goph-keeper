package jwt

import "github.com/golang-jwt/jwt/v4"

type Service interface {
	Generate(int) (string, error)
	Parse(string, jwt.Claims) error
}
