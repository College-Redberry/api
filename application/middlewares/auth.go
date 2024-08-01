package middlewares

import (
	"context"
	"net/http"
	"strings"

	"com.redberry.api/infrastructure/jwt"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authorization) < 2 {
			http.Error(w, "token missing", http.StatusUnauthorized)

			return
		}

		claims, err := jwt.Verify(authorization[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)

			return
		}

		ctx := context.WithValue(r.Context(), jwt.ContextKeyValue, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
