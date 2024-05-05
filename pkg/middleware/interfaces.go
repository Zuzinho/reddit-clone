package middleware

import "net/http"

type FuncType func(next http.Handler) http.Handler

type Authorizer interface {
	Authorize(next http.Handler) http.Handler
}

type PanicRecover interface {
	RecoverPanic(next http.Handler) http.Handler
}
