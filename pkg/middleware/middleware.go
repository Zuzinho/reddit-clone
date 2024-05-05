package middleware

import (
	"context"
	"log"
	"main/pkg/session"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// AppMiddleware - средний слой обработки
type AppMiddleware struct {
	reqMap map[string][]*regexp.Regexp
}

// NewAppMiddleware возвращает экземпляр AppMiddleware
func NewAppMiddleware() *AppMiddleware {
	return &AppMiddleware{

		reqMap: map[string][]*regexp.Regexp{
			http.MethodPost: {
				regexp.MustCompile("/api/posts"),
				regexp.MustCompile("/api/posts"),
				regexp.MustCompile("/api/post/.+"),
			},
			http.MethodGet: {
				regexp.MustCompile("/api/post/.+/(upvote|downvote|unvote)"),
			},
			http.MethodDelete: {
				regexp.MustCompile("/api/post/.+/.+"),
				regexp.MustCompile("/api/post/.+"),
			},
		},
	}
}

// Authorize проверяет авторизацию (валидность токена авторизации)
func (middle *AppMiddleware) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		regs, ok := middle.reqMap[r.Method]
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		req := false
		log.Println(r.URL.Path)
		for _, reg := range regs {
			log.Println(reg.String())

			log.Println(req)
			if reg.MatchString(r.URL.Path) {
				req = true
				break
			}
		}

		if !req {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")

		token = strings.TrimPrefix(token, "Bearer ")

		s, err := session.UnpackToken(token)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if time.Now().After(s.Exp) {
			http.Error(w, "session time wasted", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", s.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RecoverPanic возвращает приложения из panic
func (middle *AppMiddleware) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// PackMiddleware "пакует" обработчик в средний слой
func (middle *AppMiddleware) PackMiddleware(next http.Handler) http.Handler {
	return middle.RecoverPanic(
		middle.Authorize(next))
}
