package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

const secret = "gelatinous-proabolition-sighted-flea"

func main() {
	r := chi.NewRouter()
	db, err := newDB()
	if err != nil {
		panic(err)
	}
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

	r.Mount("/api/auth", authRouter(db))
	r.Mount("/api/users", userRouter(db))
	r.Mount("/api/teams", teamRouter(db))
	r.Mount("/api/schools", schoolRouter(db))
	r.Mount("/api/permissions", permissionRouter(db))
	r.Mount("/api/roles", roleRouter(db))
	r.Mount("/api/students", studentRouter(db))

	log.Println("Starting on :8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
func withRecover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				log.Printf("PANIC: %+v\n", rvr)
				debug.Stack()
			}

		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// ErrorResponse returns errors to the frontend
type ErrorResponse struct {
	Message string `json:"message"`
}

func withErrorAndDB(db *DB, next func(db *DB, w http.ResponseWriter, r *http.Request) (int, error)) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		code, err := next(db, w, r)
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
