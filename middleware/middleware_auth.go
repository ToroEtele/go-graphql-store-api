package middleware

import (
	"context"
	http "net/http"
	"time"

	model "github.com/ToroEtele/go-graphql-api/cmd/app/domain"
	"github.com/ToroEtele/go-graphql-api/service"
)

type ContextKey string

type authResponseWriter struct {
	http.ResponseWriter
	refreshTokenToResolver string
	refreshTokenFromCookie string
}

func (w *authResponseWriter) Write(b []byte) (int, error) {
	if w.refreshTokenToResolver != w.refreshTokenFromCookie {
		http.SetCookie(w, &http.Cookie{
			Name:     "refreshToken",
			Value:    w.refreshTokenToResolver,
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().Add(7 * 24 * time.Hour),
		})
	}
	return w.ResponseWriter.Write(b)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		arw := authResponseWriter{w, "", ""}
		refreshTokenContextKey := ContextKey("refreshToken")
		userContextKey := ContextKey("user")

		c, _ := r.Cookie("refreshToken")
		if c != nil {
			arw.refreshTokenFromCookie = c.Value
			arw.refreshTokenToResolver = c.Value
		}

		var ctx context.Context

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ctx = context.WithValue(r.Context(), userContextKey, nil)
		} else {
			accessTokenClaims := service.ParseAccessToken(authHeader)
			if accessTokenClaims == nil || accessTokenClaims.Id == "" {
				ctx = context.WithValue(r.Context(), userContextKey, nil)
			} else {
				user := model.CurrentUser{
					ID:      accessTokenClaims.Id,
					Name:    accessTokenClaims.Name,
					Email:   accessTokenClaims.Email,
					IsAdmin: accessTokenClaims.IsAdmin,
				}
				ctx = context.WithValue(r.Context(), userContextKey, &user)
			}
		}
		r = r.WithContext(ctx)

		ctx = context.WithValue(r.Context(), refreshTokenContextKey, &arw.refreshTokenToResolver)
		r = r.WithContext(ctx)

		// executing next
		next.ServeHTTP(&arw, r)

		w.Write([]byte(""))
	})
}
