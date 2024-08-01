package middlewares

import (
	"net/http"

	"com.redberry.api/infrastructure/jwt"
)

func Permission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(jwt.ContextKeyValue).(jwt.Claims)
		if !ok {
			http.Error(w, "could not parse token", http.StatusUnauthorized)
		}

		if !claims.IsAdmin {
			http.Error(w, "you dont have permission for this action", http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r)
	})
}
