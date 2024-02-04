package webui

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	username string
	expiry   time.Time
}

func (s Session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func IsSessionExpired(req *http.Request) (string, bool) {
	cookie, err := req.Cookie("session_token")
	if err != nil {
		return "", false
	}
	token := cookie.Value
	if v, ok := sessions.Load(token); ok {
		s, ok := v.(Session)
		if ok && !s.isExpired() {
			return s.username, true
		}
	}
	return "", false
}

var sessions sync.Map

// username|pass
type LoginCheckerFunc func([32]byte, [32]byte) bool

func Logout(req *http.Request) error {
	cookie, err := req.Cookie("session_token")
	if err != nil {
		return fmt.Errorf("fetching session token: %v", err)
	}
	token := cookie.Value
	sessions.Delete(token)
	return nil
}

func LoginPost(isValidLoginFn LoginCheckerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ok := true
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Failed to parse form: %v\n", err)
			ok = false
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if ok {
			// Calculate SHA-256 hashes for the provided and expected
			// usernames and passwords.
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))

			// If the username and password are correct, then call
			// the next handler in the chain. Make sure to return
			// afterwards, so that none of the code below is run.
			if isValidLoginFn(usernameHash, passwordHash) {
				// Create a new random session token
				// we use the "github.com/google/uuid" library to generate UUIDs
				sessionToken := uuid.NewString()
				expiresAt := time.Now().Add(120 * time.Second)

				// Set the token in the session map, along with the session information
				sessions.Store(sessionToken, Session{
					username: username,
					expiry:   expiresAt,
				})

				// Finally, we set the client cookie for "session_token" as the session token we just generated
				// we also set an expiry time of 120 seconds
				http.SetCookie(w, &http.Cookie{
					Name:    "session_token",
					Value:   sessionToken,
					Expires: expiresAt,
				})
				http.Redirect(w, r, "/", 302)
				return
			}
		}
		http.Error(w, "Failed to Login", http.StatusUnauthorized)
	})
}
