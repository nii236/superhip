package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/antonholmquist/jason"
	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
	"github.com/nii236/superhip/server/models"

	"github.com/go-chi/chi"
)

func userRouter(db *DB) http.Handler {
	r := chi.NewRouter()

	r.Post("/get/list", withErrorAndDB(db, usersGetList))
	r.Post("/get", withErrorAndDB(db, usersGetOne))
	r.Post("/get/many", withErrorAndDB(db, usersGetMany))
	r.Post("/get/many/reference", withErrorAndDB(db, usersGetManyReference))

	r.Post("/create", withErrorAndDB(db, usersCreate))

	r.Post("/update", withErrorAndDB(db, usersUpdate))
	r.Post("/update/many", withErrorAndDB(db, usersUpdateMany))

	r.Post("/delete", withErrorAndDB(db, usersDelete))
	r.Post("/delete/many", withErrorAndDB(db, usersDeleteMany))

	return r
}

func usersGetList(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetListRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.UserList{}

	err := db.List(&result)
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
	w.Write(mustMarshal(resp))
	return 200, nil
}

func usersGetOne(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetOneRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.User{}

	err := db.Read(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &Response{
			Total:   0,
			Data:    mustMarshal(models.UserList{}),
			Message: err.Error(),
		}
		err = json.NewEncoder(w).Encode(resp)
		return 200, err
	}
	if err != nil {
		return 500, err
	}

	w.Write(mustMarshal(&Response{
		Total: 1,
		Data:  mustMarshal(models.UserList{result}),
	}))
	return 200, nil
}

func usersGetMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.UserList{}

	IDs := []string{}
	for _, v := range req.IDs {
		IDs = append(IDs, v.String())
	}

	err := db.GetMany(&result, IDs)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&Response{
		Total: len(result),
		Data:  mustMarshal(result),
	}))

	return 200, nil
}

func usersGetManyReference(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetManyReferenceRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.UserList{}

	err := db.Reference(&result, req.Target, req.Column, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&Response{
		Total: len(result),
		Data:  mustMarshal(result),
	}))

	return 200, nil
}

func usersUpdate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &UpdateRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	existing := &models.User{}
	err := db.Read(existing, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	obj, err := jason.NewObjectFromBytes(req.Data)
	if err != nil {
		return 500, err
	}

	schoolFK, err := obj.GetString("school_id")
	if err != nil {
		return 500, fmt.Errorf("school_id: %s", err)
	}

	existing.SchoolID, err = uuid.FromString(schoolFK)
	if err != nil {
		return 500, err
	}

	existing.FirstName, err = obj.GetString("first_name")
	if err != nil {
		return 500, fmt.Errorf("first_name: %s", err)
	}
	existing.LastName, err = obj.GetString("last_name")
	if err != nil {
		return 500, fmt.Errorf("last_name: %s", err)
	}
	existing.Email, err = obj.GetString("email")
	if err != nil {
		return 500, fmt.Errorf("email: %s", err)
	}
	existing.Role, err = obj.GetString("role")
	if err != nil {
		return 500, fmt.Errorf("role: %s", err)
	}
	if err != nil {
		return 500, err
	}
	password, err := obj.GetString("password")
	if err != nil {
		return 500, fmt.Errorf("password: %s", err)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 500, err
	}

	hashedB64 := base64.StdEncoding.EncodeToString(hashed)
	if hashedB64 != existing.PasswordHash {
		existing.PasswordHash = hashedB64
	}

	updated := &models.User{}
	fmt.Printf("%+v\n", existing)
	err = db.Update(updated, existing, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&Response{
		Total: 1,
		Data:  mustMarshal([]*models.User{updated}),
	}))

	return 200, nil
}

func usersUpdateMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &UpdateManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	obj, err := jason.NewObjectFromBytes(req.Data)
	if err != nil {
		return 500, err
	}
	updateTo := &models.User{}

	schoolFK, err := obj.GetString("school_id")
	if err != nil {
		return 500, fmt.Errorf("school_id: %s", err)
	}

	updateTo.SchoolID, err = uuid.FromString(schoolFK)
	if err != nil {
		return 500, err
	}
	updateTo.FirstName, err = obj.GetString("first_name")
	if err != nil {
		return 500, fmt.Errorf("first_name: %s", err)
	}
	updateTo.LastName, err = obj.GetString("last_name")
	if err != nil {
		return 500, fmt.Errorf("last_name: %s", err)
	}
	updateTo.Email, err = obj.GetString("email")
	if err != nil {
		return 500, fmt.Errorf("email: %s", err)
	}
	updateTo.Role, err = obj.GetString("role")
	if err != nil {
		return 500, fmt.Errorf("role: %s", err)
	}
	if err != nil {
		return 500, err
	}

	password, err := obj.GetString("password")
	if err != nil {
		return 500, fmt.Errorf("password: %s", err)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 500, err
	}

	hashedB64 := base64.StdEncoding.EncodeToString(hashed)
	updateTo.PasswordHash = hashedB64

	IDs := []string{}
	for _, v := range req.IDs {
		IDs = append(IDs, v.String())
	}

	updated := models.UserList{}

	fmt.Println(updateTo)
	err = db.UpdateMany(&updated, updateTo, IDs)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&Response{
		Total: len(updated),
		Data:  mustMarshal(updated),
	}))
	return 200, nil
}

func usersCreate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return 500, err
	}

	defer r.Body.Close()

	obj, err := jason.NewObjectFromBytes(req.Data)
	if err != nil {
		return 500, err
	}
	createWith := &models.User{}

	schoolFK, err := obj.GetString("school_id")
	if err != nil {
		return 500, fmt.Errorf("school_id: %s", err)
	}

	createWith.SchoolID, err = uuid.FromString(schoolFK)
	if err != nil {
		return 500, err
	}
	createWith.FirstName, err = obj.GetString("first_name")
	if err != nil {
		return 500, fmt.Errorf("first_name: %s", err)
	}
	createWith.LastName, err = obj.GetString("last_name")
	if err != nil {
		return 500, fmt.Errorf("last_name: %s", err)
	}
	createWith.Email, err = obj.GetString("email")
	if err != nil {
		return 500, fmt.Errorf("email: %s", err)
	}
	createWith.Role, err = obj.GetString("role")
	if err != nil {
		return 500, fmt.Errorf("role: %s", err)
	}
	if err != nil {
		return 500, err
	}

	password, err := obj.GetString("password")
	if err != nil {
		return 500, fmt.Errorf("password: %s", err)
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 500, err
	}
	hashedB64 := base64.StdEncoding.EncodeToString(hashed)
	createWith.PasswordHash = hashedB64

	created := &models.User{}
	log.Printf("%+v", createWith)
	err = db.Create(created, createWith)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&Response{
		Total: 1,
		Data:  mustMarshal(models.UserList{created}),
	}))

	return 200, nil
}

func usersDelete(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &DeleteRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.User{}

	err := db.Delete(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &Response{
			Total:   0,
			Data:    mustMarshal(models.UserList{}),
			Message: err.Error(),
		}
		err = json.NewEncoder(w).Encode(resp)
		return 200, err
	}
	if err != nil {
		return 500, err
	}

	w.Write(mustMarshal(&Response{
		Total: 1,
		Data:  mustMarshal(models.UserList{result}),
	}))
	return 200, nil
}

func usersDeleteMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &DeleteManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.UserList{}

	IDs := []string{}
	for _, v := range req.IDs {
		IDs = append(IDs, v.String())
	}

	err := db.DeleteMany(&result, IDs)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&Response{
		Total: len(result),
		Data:  mustMarshal(result),
	}))

	return 200, nil
}
