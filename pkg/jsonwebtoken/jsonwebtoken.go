package jsonwebtoken

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
	Role   int64  `json:"role"`
}

const (
	tokenTTL = 12 * time.Hour
)

func GenerateToken(userID string, role int64, signingKey string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
		role,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string, signingKey string) (string, int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", 0, err
	}

	if !token.Valid {
		return "", 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.Role, nil
}
