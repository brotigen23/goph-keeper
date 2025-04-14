package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type accessClaims struct {
	jwt.RegisteredClaims
	ID int
}
type refreshClaims struct {
	jwt.RegisteredClaims
}

func createAccessToken(id int, key string, expires time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
		},
		ID: id,
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func createRefreshToken(key string, expires time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
		},
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateTokens(id int, accessKey, refreshKey string, accessExpires, refreshExpires time.Duration) (acessToken, refreshToken string, err error) {
	acessToken, err = createAccessToken(id, accessKey, accessExpires)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = createRefreshToken(refreshKey, refreshExpires)
	if err != nil {
		return "", "", err
	}

	return
}

func GetIDFromJWT(tokenString string, key string) (int, error) {
	claims := &accessClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (any, error) {
			return []byte(key), nil
		})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, ErrTokenIsInvalid
	}
	return claims.ID, nil
}
