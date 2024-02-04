package webui

import (
	"crypto/sha256"
	"crypto/subtle"
	"io"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	"testing"

	"github.com/a-h/templ"
	"plramos.win/ptscheduler/internal/datamodel"
	"plramos.win/ptscheduler/internal/webui/pages"
)

func TestHome(t *testing.T) {
	nav := []NavBar{
		NavBar{Page: "About", Link: "/about"},
	}

	//This works too so I will keep for documentation only
	// 	handler := func(w http.ResponseWriter, r *http.Request) {
	// 		p := NewPage(nav, pages.NewHome())
	// 		p.Render(context.Background(), w)
	// 	}

	handler := templ.Handler(NewPage(nav, pages.NewHome("Pedro Ramos",
		[]time.Time{time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
		[]datamodel.AvailabilityPeriod{
			datamodel.AvailabilityPeriod{Start: time.Now(), End: time.Now().Add(time.Hour)},
		},
	))).ServeHTTP

	req := httptest.NewRequest("GET", "http://example.com/", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("Server did not respond with status code 200. --response--\ncode=%d\nbody=%s\n", resp.StatusCode, body)
	}
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Header.Get("Content-Type"))
	// fmt.Println(string(body))
}

func TestStatic(t *testing.T) {

	req := httptest.NewRequest("GET", "http://example.com/static/output.css", nil)
	w := httptest.NewRecorder()
	StaticHandler().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("Server did not respond with status code 200. --response--\ncode=%d\nbody=%s\n", resp.StatusCode, body)
	}
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Header.Get("Content-Type"))
	// fmt.Println(string(body))
	if resp.StatusCode != 200 {
		t.Errorf("Server did not respond with status code 200 to /static/output.css. --response--\ncode=%d\nbody=%s\n", resp.StatusCode, body)
	}
}

func TestLogin(t *testing.T) {

	NewTestLoginChecker := func(testuser, testpass string) LoginCheckerFunc {
		return func(user, pass [32]byte) bool {
			expectedUsernameHash := sha256.Sum256([]byte(testuser))
			expectedPasswordHash := sha256.Sum256([]byte(testpass))

			// Use the subtle.ConstantTimeCompare() function to check if
			// the provided username and password hashes equal the
			// expected username and password hashes. ConstantTimeCompare
			// will return 1 if the values are equal, or 0 otherwise.
			// Importantly, we should to do the work to evaluate both the
			// username and password before checking the return values to
			// avoid leaking information.
			usernameMatch := (subtle.ConstantTimeCompare(user[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(pass[:], expectedPasswordHash[:]) == 1)

			t.Logf("User=%s sha=%v\n", testuser, expectedUsernameHash)
			t.Logf("Pass=%s sha=%v\n", testuser, expectedPasswordHash)

			t.Logf("rcv user=%v\nrcv pass=%v\n", user, pass)

			return usernameMatch && passwordMatch
		}
	}

	handler := LoginPost(NewTestLoginChecker("pedramos", "123"))
	formdata := url.Values{"username": {"pedramos"}, "password": {"123"}}.Encode()

	formreader := strings.NewReader(formdata)
	req := httptest.NewRequest("POST", "http://example.com/", formreader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 302 {
		t.Errorf("Failed login.\n # response #\ncode=%d\nbody=%s\n", resp.StatusCode, body)
	}
}
