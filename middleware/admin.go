package middleware

import (
	"net/http"
)

// AdminAuthMiddleware checks for the admin password in the Authorization header
func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for Authorization header
		password := r.Header.Get("Authorization")

		// If the password is not correct, return Forbidden status
		if password != "root" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// If authorized, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
