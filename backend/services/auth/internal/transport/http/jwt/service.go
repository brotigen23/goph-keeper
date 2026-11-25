package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const defaultTTL = time.Hour * 100

type service struct {
	key       string
	ttl       time.Duration
	algotithm jwt.SigningMethod
}

func New(key string, options ...Option) Service {
	ret := &service{
		key:       key,
		ttl:       defaultTTL,
		algotithm: jwt.SigningMethodHS256,
	}

	for _, f := range options {
		f(ret)
	}
	log.Println(ret)
	return ret
}

func (s *service) Generate(id int) (string, error) {
	claims := JWTClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.ttl)),
		},
	}
	token := jwt.NewWithClaims(s.algotithm, claims)

	tokenString, err := token.SignedString([]byte(s.key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) Parse(string, jwt.Claims) error {
	return nil
}
