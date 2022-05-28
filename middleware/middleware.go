package middleware

import (
	"net/http"

	"github.com/dinner-plans-llc/common/auth"
	"github.com/gorilla/mux"
)

const (
	UID string = "uid"
)

// Middleware ...
type Middleware struct {
	auth auth.IAuth
}

// NewMiddleware
func NewMiddleware(a auth.IAuth) *Middleware {
	return &Middleware{auth: a}
}

func (mi *Middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		token := r.Header.Get("Authentication")

		vars := mux.Vars(r)

		uid, err := mi.auth.VerifyToken(ctx, token)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}

		vars[UID] = uid
		next.ServeHTTP(w, r)
	})
}
