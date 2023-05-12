package login

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{"enrinal", "123"},
	{"rogu", "123"},
}

type Session struct {
	username string
	expiry   time.Time
}

var SessionMap = map[string]Session{}

func (s *Session) IsExpired() bool {
	return s.expiry.Before(time.Now())
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "not valid", http.StatusForbidden)
			return
		}

		for _, u := range users {
			if u.Username != username && u.Password != pass {
				http.Error(w, "not valid", http.StatusForbidden)
				return
			}
		}

		sessionToken := uuid.New().String()
		expiredAt := time.Now().Add(5 * time.Second)

		SessionMap[sessionToken] = Session{
			username: username,
			expiry:   expiredAt,
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: expiredAt,
		})

		log.Println(SessionMap)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
