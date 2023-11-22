package middleware

import (
	"context"
	http "net/http"
	"time"
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

		// before executing next
		c, _ := r.Cookie("refreshToken")
		if c != nil {
			arw.refreshTokenFromCookie = c.Value
			arw.refreshTokenToResolver = c.Value
		}
		ctx := context.WithValue(r.Context(), refreshTokenContextKey, &arw.refreshTokenToResolver)
		r = r.WithContext(ctx)

		// executing next
		next.ServeHTTP(&arw, r)

		// after executing next
		w.Write([]byte(""))

	})
}
