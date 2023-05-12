package middleware

import (
	"flight-example-api/login"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "unauthenticated", http.StatusForbidden)
			return
		}

		sessToken := c.Value

		usesSession, ok := login.SessionMap[sessToken]
		if !ok {
			http.Error(w, "unauthenticated", http.StatusForbidden)
			return
		}

		if usesSession.IsExpired() {
			http.Error(w, "unauthenticated", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
