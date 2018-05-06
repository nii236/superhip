package main

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
)

// SignupRequest creates a new user and response with jwt
type SignupRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func authRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/users", withError(users))
	r.Post("/refresh_token", withError(refreshToken))
	r.Post("/change_password", withError(changePassword))
	r.Post("/change_email", withError(changeEmail))
	r.Post("/forgot_password", withError(forgotPassword))
	r.Post("/forgot_username", withError(forgotUsername))
	r.Post("/check", withError(check))
	return r
}

func users(w http.ResponseWriter, r *http.Request) (int, error) {
	// signup
	// req := &SignupRequest{}
	// err := json.NewDecoder(r.Body).Decode(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 500, err
	// }

	// hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return 500, err
	// }

	// _, err = conn.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", req.Email, base64.StdEncoding.EncodeToString(hashed))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 500, err
	// }

	// var role string
	// var email string
	// err = conn.QueryRow("SELECT role, email FROM users WHERE email = $1 LIMIT 1", req.Email).Scan(&role, &email)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 500, err
	// }

	// t, err := createToken(role, email)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 500, err
	// }

	// w.Write([]byte(t))
	return 200, nil
}

type SigninRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func refreshToken(w http.ResponseWriter, r *http.Request) (int, error) {
	// signin
	// req := &SigninRequest{}
	// err := json.NewDecoder(r.Body).Decode(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 500, err
	// }

	// var role string
	// var email string
	// var password string
	// conn.QueryRow("SELECT role, email, password FROM users WHERE email = $1 LIMIT 1", req.Email).Scan(&role, &email, &password)
	// fmt.Println("email:", email)
	// fmt.Println("role:", role)
	// fmt.Println("password:", password)
	// t, err := createToken(role, email)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 500, err
	// }
	// fmt.Println(t)
	// w.Write([]byte(t))
	return 200, nil
}

func check(w http.ResponseWriter, r *http.Request) (int, error) {
	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	return 500, err
	// }

	// token, err := jwt.Parse(string(b), func(token *jwt.Token) (interface{}, error) {
	// 	// Don't forget to validate the alg is what you expect:
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return []byte(secret), nil
	// })

	// claims := token.Claims.(jwt.MapClaims)
	// json.NewEncoder(w).Encode(claims)
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

func createToken(role string, email string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role":  role,
		"email": email,
		"exp":   time.Now().Add(2 * time.Hour).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(secret))
}
