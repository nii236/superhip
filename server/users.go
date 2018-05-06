package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"

	"github.com/go-chi/chi"
)

func userRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/get/list", withError(usersGetList))
	r.Post("/get", withError(usersGetOne))
	r.Post("/get/many", withError(usersGetMany))
	r.Post("/get/many/reference", withError(usersGetManyReference))

	r.Post("/create", withError(usersCreate))

	r.Post("/update", withError(usersUpdate))
	r.Post("/update/many", withError(usersUpdateMany))

	r.Post("/delete", withError(usersDelete))
	r.Post("/delete/many", withError(usersDeleteMany))

	return r
}

// User is the user model
type User struct {
	ID           uuid.UUID `db:"id" json:"id,omitempty"`
	Email        string    `db:"email" json:"email,omitempty"`
	FirstName    string    `db:"first_name" json:"first_name,omitempty"`
	LastName     string    `db:"last_name" json:"last_name,omitempty"`
	Password     string    `json:"-"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Role         string    `db:"role" json:"role,omitempty"`
}

func mustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func usersGetList(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetListRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return 500, err
	}

	defer r.Body.Close()

	result := []*User{}
	err = conn.Select(&result, userQueries["all"])
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	resp := &Response{
		Total: len(result),
		Data:  mustMarshal(result),
	}

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		return 500, err
	}
	return 200, nil
}

func usersGetOne(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetOneRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return 500, err
	}

	defer r.Body.Close()

	result := &User{}
	log.Println("ID", req.ID)
	err = conn.Get(result, userQueries["get"], req.ID)
	if err != nil && err == sql.ErrNoRows {
		resp := &Response{
			Total:   0,
			Data:    mustMarshal([]*User{}),
			Message: err.Error(),
		}
		err = json.NewEncoder(w).Encode(resp)
		return 200, err
	}
	if err != nil {
		return 500, err
	}
	resp := &Response{
		Total: 1,
		Data:  mustMarshal([]*User{result}),
	}

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		return 500, err
	}
	return 200, nil
}

func usersGetMany(w http.ResponseWriter, r *http.Request) (int, error) {
	id := chi.URLParam(r, "id")
	result := []*User{}
	err := conn.Select(result, userQueries["get"], id)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}
	resp := &Response{
		Total: len(result),
		Data:  mustMarshal(result),
	}

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func usersGetManyReference(w http.ResponseWriter, r *http.Request) (int, error) {
	return 200, nil
}

func usersUpdate(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &UpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return 500, err
	}

	existing := &User{}
	err = conn.Get(existing, userQueries["get"], req.ID)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	existing.FirstName = req.Data["first_name"]
	existing.LastName = req.Data["last_name"]
	existing.Email = req.Data["email"]
	existing.Role = req.Data["role"]

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return 500, err
	}
	hashedB64 := base64.StdEncoding.EncodeToString(hashed)
	if hashedB64 != existing.PasswordHash {
		existing.PasswordHash = hashedB64
	}

	log.Printf("%+v", existing)
	rows, err := conn.NamedQuery(userQueries["update"], existing)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	updated := &User{}

	for rows.Next() {
		err = rows.StructScan(updated)
		if err != nil {
			return 500, err
		}
	}

	json.NewEncoder(w).Encode(&Response{
		Total: 1,
		Data:  mustMarshal([]*User{updated}),
	})

	return 200, nil
}

func usersUpdateMany(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &UpdateManyRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return 500, err
	}
	updatedUsers := []*User{}
	for _, id := range req.IDs {
		existing := &User{}
		err = conn.Get(existing, userQueries["get"], id)
		if err != nil && err == sql.ErrNoRows {
			return 404, err
		}
		if err != nil && err != sql.ErrNoRows {
			return 500, err
		}

		existing.FirstName = req.Data["first_name"]
		existing.LastName = req.Data["last_name"]
		existing.Email = req.Data["email"]
		existing.Role = req.Data["role"]

		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Data["password"]), bcrypt.DefaultCost)
		if err != nil {
			return 500, err
		}
		hashedB64 := base64.StdEncoding.EncodeToString(hashed)
		if hashedB64 != existing.PasswordHash {
			existing.PasswordHash = hashedB64
		}

		log.Printf("%+v", existing)
		rows, err := conn.NamedQuery(userQueries["update"], existing)
		if err != nil {
			return 500, err
		}

		updated := &User{}

		for rows.Next() {
			err = rows.StructScan(updated)
			if err != nil {
				return 500, err
			}
		}

		updatedUsers = append(updatedUsers, updated)

	}

	json.NewEncoder(w).Encode(&Response{
		Total: len(updatedUsers),
		Data:  mustMarshal(updatedUsers),
	})

	return 200, nil
}

func usersCreate(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return 500, err
	}
	log.Println(req.Data)
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return 500, err
	}

	user := &User{
		FirstName:    req.Data["first_name"],
		LastName:     req.Data["last_name"],
		Email:        req.Data["email"],
		PasswordHash: base64.StdEncoding.EncodeToString(hashed),
		Role:         "teacher",
	}

	log.Printf("%+v", user)
	rows, err := conn.NamedQuery(userQueries["create"], user)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	created := &User{}

	for rows.Next() {
		err = rows.StructScan(created)
		if err != nil {
			return 500, err
		}
	}

	json.NewEncoder(w).Encode(&Response{
		Total: 1,
		Data:  mustMarshal([]*User{created}),
	})

	defer r.Body.Close()

	return 200, nil
}

func usersDelete(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &DeleteRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return 500, err
	}
	deleted := &User{}
	err = conn.Get(deleted, userQueries["archive"], req.ID)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	json.NewEncoder(w).Encode(&Response{
		Total: 1,
		Data:  mustMarshal([]*User{deleted}),
	})
	return 200, nil
}

func usersDeleteMany(w http.ResponseWriter, r *http.Request) (int, error) {
	req := &DeleteManyRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return 500, err
	}

	deleted := []*User{}
	for _, id := range req.IDs {
		user := &User{}
		err = conn.Get(user, userQueries["archive"], id)
		if err != nil && err == sql.ErrNoRows {
			return 404, err
		}
		if err != nil && err != sql.ErrNoRows {
			return 500, err
		}
		deleted = append(deleted, user)
	}
	json.NewEncoder(w).Encode(&Response{
		Total: 1,
		Data:  mustMarshal(deleted),
	})
	return 200, nil
}
