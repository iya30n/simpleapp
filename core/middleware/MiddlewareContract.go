package middleware

import (
	"net/http"
)

type MiddlewareContract interface {
	Handle(w http.ResponseWriter, r *http.Request) bool
}