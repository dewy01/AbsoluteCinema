package middleware

import (
	"absolutecinema/internal/auth"
	"net/http"

	"github.com/google/uuid"
)

func ChainMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func AuthenticationMiddleware(sessionService *auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(auth.CookieName)
			if err != nil {
				http.Error(w, "no session cookie", http.StatusUnauthorized)
				return
			}

			sessionID, err := uuid.Parse(cookie.Value)
			if err != nil {
				http.Error(w, "invalid session id", http.StatusUnauthorized)
				return
			}

			_, err = sessionService.Get(r.Context(), sessionID)
			if err != nil {
				http.Error(w, "session not found or expired", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func AuthorizationMiddleware(sessionService *auth.Service, requiredRole auth.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie(auth.CookieName)
			if err != nil {
				http.Error(w, "no session cookie", http.StatusUnauthorized)
				return
			}

			sessionID, err := uuid.Parse(cookie.Value)
			if err != nil {
				http.Error(w, "invalid session id", http.StatusUnauthorized)
				return
			}

			session, err := sessionService.Get(r.Context(), sessionID)
			if err != nil {
				http.Error(w, "session not found or expired", http.StatusUnauthorized)
				return
			}

			hasAccess := session.Data.Role == requiredRole || session.Data.Role == auth.Admin
			if !hasAccess {
				http.Error(w, "insufficient permissions", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
