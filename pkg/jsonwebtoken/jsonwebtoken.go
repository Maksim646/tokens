package jsonwebtoken

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId    string `json:"user_id"`
	IP        string `json:"ip"`
	RefreshID string `json:"refresh_id"`
}

func GenerateAccessToken(userID string, ip string, refreshID string, jwtSigningKey string, accessTokenTTL time.Duration) (string, error) {
	claims := &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:    userID,
		IP:        ip,
		RefreshID: refreshID,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenStr, err := accessToken.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", err
	}

	return accessTokenStr, nil
}

func GenerateRefreshToken(userID string, ip string, refreshID string, jwtSigningKey string, refreshTokenLength int) (string, error) {
	refreshTokenStr, err := generateSecureRandomToken(refreshTokenLength)
	if err != nil {
		return "", err
	}
	return refreshTokenStr, nil
}

func ParseToken(accessToken string, signingKey string) (string, string, string, error) {
	return parseJWTToken(accessToken, signingKey, false)
}

func ParseTokenWithOutClaims(accessToken string, signingKey string) (string, string, string, error) {
	return parseJWTToken(accessToken, signingKey, true)
}

func parseJWTToken(accessToken string, signingKey string, skipClaimsValidation bool) (string, string, string, error) {
	parser := &jwt.Parser{
		SkipClaimsValidation: skipClaimsValidation,
	}

	token, err := parser.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", "", "", err
	}

	if !token.Valid {
		return "", "", "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", "", "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.IP, claims.RefreshID, nil
}

func generateSecureRandomToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
