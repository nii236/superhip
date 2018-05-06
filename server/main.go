package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/nleof/goyesql"
)

var conn *sqlx.DB
var userQueries goyesql.Queries
var schoolQueries goyesql.Queries
var teamQueries goyesql.Queries
var studentQueries goyesql.Queries

const secret = "gelatinous-proabolition-sighted-flea"

func init() {
	userQueries = goyesql.MustParseFile("./queries/users.sql")
	schoolQueries = goyesql.MustParseFile("./queries/schools.sql")
	teamQueries = goyesql.MustParseFile("./queries/teams.sql")
	studentQueries = goyesql.MustParseFile("./queries/students.sql")
}

func main() {
	r := chi.NewRouter()
	db()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Content-Range"},
		ExposedHeaders:   []string{"Link", "Content-Range"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/auth", authRouter())
	r.Mount("/users", userRouter())

	log.Println("Starting on :8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func withError(next func(w http.ResponseWriter, r *http.Request) (int, error)) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		code, err := next(w, r)
		if err != nil {
			log.Println(err)
			resp, err := json.Marshal(&ErrorResponse{
				Message: err.Error(),
			})
			if err != nil {
				http.Error(w, err.Error(), code)
				return
			}
			http.Error(w, string(resp), code)
		}
	}

	return http.HandlerFunc(fn)
}

func db() {
	db, err := sqlx.Connect("postgres", "user=dev dbname=superhip password=dev sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	conn = db
}
