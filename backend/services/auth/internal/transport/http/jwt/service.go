package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const defaultTTL = time.Hour

type service struct {
	key       string
	ttl       time.Duration
	algotithm jwt.SigningMethod
}

func New(key string, options ...Option) Service {
	ret := &service{
		key: key,
		ttl: defaultTTL,
	}

	for _, f := range options {
		f(ret)
	}
	return ret
}

func (s *service) Generate(id int) (string, error) {
	claims := JWTClaims{
		ID: id,
	}
	token := jwt.NewWithClaims(s.algotithm, claims)

	tokenString, err := token.SignedString(s.key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) Parse(string, jwt.Claims) error {
	return nil
}
