package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Option func(*service)

func WithTTL(ttl time.Duration) Option {
	return func(s *service) {
		s.ttl = ttl
	}
}

type Algotithm int

const (
	HS256 Algotithm = iota
)

func WithAlgotithm(a Algotithm) Option {
	return func(s *service) {
		switch a {
		case HS256:
			s.algotithm = jwt.SigningMethodHS256
		default:
			return
		}
	}
}
