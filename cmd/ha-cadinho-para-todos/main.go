package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"plramos.win/ptscheduler/internal/datamodel"
	"plramos.win/ptscheduler/internal/sqlc"
	"plramos.win/ptscheduler/internal/webui"
	"plramos.win/ptscheduler/internal/webui/pages"

	_ "modernc.org/sqlite"
)

var ddl = sqlc.Schema
var ctx = context.Background()

func InitDB(path string) (*sql.DB, error) {
	var (
		err         error
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	)
	defer cancel()

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("Opening database: %v", err)
	}
	if db == nil {
		return nil, fmt.Errorf("Opening database: DB connection is nil")
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return db, nil
}

func NewLoginChecker(db *sql.DB) webui.LoginCheckerFunc {
	return func(user, pass [32]byte) bool {
		q := sqlc.New(db)
		pwd, err := q.GetPassword(ctx, user[:])
		if err != nil {
			log.Printf("Failed to login user: %v", err)
			return false
		}
		return bytes.Equal(pass[:], pwd)
	}
}

func NewHomeHander(db *sql.DB, nav []webui.NavBar) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		username, ok := webui.IsSessionExpired(req)
		if !ok {
			http.Redirect(w, req, "/login", 302)
		}

		userhash := sha256.Sum256([]byte(username))

		q := sqlc.New(db)
		name, err := q.RealName(ctx, userhash[:])

		if err != nil || !name.Valid {
			log.Printf("reading full name from username '%s': %v\n", username, err)
		}
		scheduled, err := q.ScheduledTrainning(ctx, userhash[:])
		if err != nil {
			log.Printf("reading trainning sessions for username '%s': %v\n", username, err)
		}

		schDate := make([]time.Time, 0, len(scheduled))
		for _, t := range scheduled {
			if t.Valid {
				parsed, err := time.Parse(time.RFC3339, t.String)
				if err != nil {
					log.Printf("Could not parse date %s from user %s: %v\n", t.String, username, err)
				}
				schDate = append(schDate, parsed)
			}
		}
		avail, err := q.ScheduledAvailability(ctx, userhash[:])
		if err != nil {
			log.Printf("reading trainning sessions for username '%s': %v\n", username, err)
		}

		avails := make([]datamodel.AvailabilityPeriod, 0, len(avail))
		for _, t := range avail {
			var avail datamodel.AvailabilityPeriod
			if t.Startdate.Valid {
				parsed, err := time.Parse(time.RFC3339, t.Startdate.String)
				if err != nil {
					log.Printf("Could not parse date %s from user %s: %v\n", t.Startdate.String, username, err)
					continue
				}
				avail.Start = parsed
			}
			if t.Enddate.Valid {
				parsed, err := time.Parse(time.RFC3339, t.Enddate.String)
				if err != nil {
					log.Printf("Could not parse date %s from user %s: %v\n", t.Enddate.String, username, err)
					continue
				}
				avail.End = parsed
			}
			avails = append(avails, avail)
		}
		webui.NewPage(nav, pages.NewHome(name.String, schDate, avails)).Render(req.Context(), w)
	}
}

func StartServer(db *sql.DB) {
	nav := []webui.NavBar{
		webui.NavBar{Page: "Sobre", Link: "/about"},
		webui.NavBar{Page: "Inicio", Link: "/"},
	}

	mux := http.NewServeMux()

	// Static content
	mux.Handle("/static/", webui.StaticHandler())

	// Pages
	mux.HandleFunc("GET /{$}", NewHomeHander(db, nav))
	mux.Handle("/about", webui.NewPageHandler(nav, pages.NewAbout()))

	// Login
	mux.Handle("POST /login", webui.LoginPost(NewLoginChecker(db)))
	mux.Handle("GET /login", webui.NewPageHandler(nav, pages.NewLogin()))
	mux.HandleFunc("GET /logout", func(w http.ResponseWriter, r *http.Request) {
		webui.Logout(r)
		http.Redirect(w, r, "/login", 302)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))

}
func main() {

	dbPath := os.Getenv("PTSCHD_DB")
	if dbPath == "" {
		dbPath = "./ptsheduler.sqlite"
	}

	isDemo := false
	demo := os.Getenv("MAKEADEMO")
	if demo == "1" {
		isDemo = true
		dbPath = "./demo.sqlite"
		if _, err := os.Stat(dbPath); err == nil {
			os.Remove(dbPath)

		}
	}

	db, err := InitDB(dbPath)
	if err != nil {
		log.Fatalf("Could not init DB: %v", err)
	}

	if isDemo {
		err := AddDemoData(ctx, db)
		if err != nil {
			os.Remove(dbPath)
			log.Fatalf("Failed to add demo data : %v", err)
		}
		defer os.Remove(dbPath)
	}

	fmt.Println("serving at http://localhost:8080")
	StartServer(db)
}
