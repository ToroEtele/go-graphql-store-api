package service

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func NewRefreshToken(claims RefreshClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedAccessToken.Claims.(*UserClaims)
}

func ParseRefreshToken(refreshToken string) *RefreshClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedRefreshToken.Claims.(*RefreshClaims)
}

type CookieAccess struct {
	Writer http.ResponseWriter
	// refreshToken string
}

// method to write cookie
func (c *CookieAccess) SetToken(token string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})
}
