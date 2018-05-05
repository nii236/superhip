package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx"
)

var conn *pgx.ConnPool

func main() {
	r := chi.NewRouter()
	db()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Post("/auth/users", withError(users))
	r.Post("/auth/refresh_token", withError(refreshToken))
	r.Post("/auth/change_password", withError(changePassword))
	r.Post("/auth/change_email", withError(changeEmail))
	r.Post("/auth/forgot_password", withError(forgotPassword))
	r.Post("/auth/forgot_username", withError(forgotUsername))

	log.Println("Starting on :8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}

func withError(next func(w http.ResponseWriter, r *http.Request) (int, error)) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		code, err := next(w, r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), code)
		}
	}

	return http.HandlerFunc(fn)
}

func db() {
	connConfig := &pgx.ConnConfig{
		Database:  "dev",
		Host:      "localhost",
		Port:      5432,
		User:      "dev",
		Password:  "dev",
		TLSConfig: nil,
	}

	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     *connConfig,
		AfterConnect:   nil,
		MaxConnections: 20,
		AcquireTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}
	conn = connPool
}

func users(w http.ResponseWriter, r *http.Request) (int, error) {
	// signup
	var id int
	conn.QueryRow("SELECT count(id) FROM users").Scan(&id)
	fmt.Println(id)
	t, err := createToken()
	if err != nil {
		fmt.Println(err)
		return 500, err
	}
	fmt.Println(t)
	w.Write([]byte(t))
	return 200, nil
}
func refreshToken(w http.ResponseWriter, r *http.Request) (int, error) {
	// signin
	return 200, nil
}
func changePassword(w http.ResponseWriter, r *http.Request) (int, error) {
	return 200, nil
}
func changeEmail(w http.ResponseWriter, r *http.Request) (int, error) {
	return 200, nil
}
func forgotPassword(w http.ResponseWriter, r *http.Request) (int, error) {
	return 200, nil
}
func forgotUsername(w http.ResponseWriter, r *http.Request) (int, error) {
	return 200, nil
}

func createToken() (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte("proabolition-sighted-flea"))
}
