package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antonholmquist/jason"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
	"github.com/nii236/superhip/models"

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
	req := &models.GetListRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.UserList{}
	opts := &ListOptions{
		Offset:         req.Pagination.Page,
		Limit:          req.Pagination.PerPage,
		OrderBy:        req.Sort.Field,
		OrderDirection: req.Sort.Order,
	}
	err := db.List(&result, opts)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	total, err := db.Total("users")
	if err != nil {
		return 500, err
	}
	resp := &models.Response{
		Total: total,
		Data:  mustMarshal(result),
	}
	w.Write(mustMarshal(resp))
	return 200, nil
}

func usersGetOne(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetOneRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.User{}

	err := db.Read(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &models.Response{
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

	w.Write(mustMarshal(&models.Response{
		Total: 1,
		Data:  mustMarshal(models.UserList{result}),
	}))
	return 200, nil
}

func usersGetMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetManyRequest{}
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

	w.Write(mustMarshal(&models.Response{
		Total: len(result),
		Data:  mustMarshal(result),
	}))

	return 200, nil
}

func usersGetManyReference(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetManyReferenceRequest{}
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

	w.Write(mustMarshal(&models.Response{
		Total: len(result),
		Data:  mustMarshal(result),
	}))

	return 200, nil
}

func usersUpdate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.UpdateRequest{}
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
	err = db.Update(updated, existing, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	err = db.DropJoins("schools_users", "user_id", updated.ID.String())
	if err != nil {
		return 500, fmt.Errorf("could not drop joins: %s", err)
	}
	schoolFKs, err := obj.GetStringArray("school_ids")
	if err == nil {
		for _, v := range schoolFKs {
			err = db.MakeJoin("schools_users", "school_id", "user_id", v, updated.ID.String())
			if err != nil {
				return 500, err
			}
		}
	} else {
		fmt.Println("no school ids provided")
	}

	err = db.DropJoins("roles_users", "user_id", updated.ID.String())
	if err != nil {
		return 500, fmt.Errorf("could not drop joins: %s", err)
	}
	roleFKs, err := obj.GetStringArray("role_ids")
	if err == nil {
		for _, v := range roleFKs {
			err = db.MakeJoin("roles_users", "role_id", "user_id", v, updated.ID.String())
			if err != nil {
				return 500, err
			}
		}
	} else {
		fmt.Println("no row ids provided")
	}

	w.Write(mustMarshal(&models.Response{
		Total: 1,
		Data:  mustMarshal([]*models.User{updated}),
	}))

	return 200, nil
}

func usersUpdateMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.UpdateManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	obj, err := jason.NewObjectFromBytes(req.Data)
	if err != nil {
		return 500, err
	}
	updateTo := &models.User{}

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

	err = db.UpdateMany(&updated, updateTo, IDs)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	for _, v := range req.IDs {
		err = db.DropJoins("schools_users", "user_id", v.String())
		if err != nil {
			return 500, fmt.Errorf("could not drop joins: %s", err)
		}
		schoolFKs, err := obj.GetStringArray("school_ids")
		if err == nil {
			for _, fk := range schoolFKs {
				err = db.MakeJoin("schools_users", "school_id", "user_id", fk, v.String())
				if err != nil {
					return 500, err
				}
			}
		} else {
			fmt.Println("no school ids provided")
		}

		err = db.DropJoins("roles_users", "user_id", v.String())
		if err != nil {
			return 500, fmt.Errorf("could not drop joins: %s", err)
		}
		roleFKs, err := obj.GetStringArray("role_ids")
		if err == nil {
			for _, fk := range roleFKs {
				err = db.MakeJoin("roles_users", "role_id", "user_id", fk, v.String())
				if err != nil {
					return 500, err
				}
			}
		} else {
			fmt.Println("no row ids provided")
		}
	}

	w.Write(mustMarshal(&models.Response{
		Total: len(updated),
		Data:  mustMarshal(updated),
	}))
	return 200, nil
}

func usersCreate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.CreateRequest{}
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
	err = db.Create(created, createWith)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	err = db.DropJoins("schools_users", "user_id", created.ID.String())
	if err != nil {
		return 500, fmt.Errorf("could not drop joins: %s", err)
	}

	schoolFKs, err := obj.GetStringArray("school_ids")
	if err == nil {
		for _, v := range schoolFKs {
			err = db.MakeJoin("schools_users", "school_id", "user_id", v, created.ID.String())
			if err != nil {
				return 500, err
			}
		}
	} else {
		fmt.Println("no school ids provided")
	}

	err = db.DropJoins("roles_users", "user_id", created.ID.String())
	if err != nil {
		return 500, fmt.Errorf("could not drop joins: %s", err)
	}
	roleFKs, err := obj.GetStringArray("role_ids")
	if err == nil {
		for _, v := range roleFKs {
			err = db.MakeJoin("roles_users", "role_id", "user_id", v, created.ID.String())
			if err != nil {
				return 500, err
			}
		}
	} else {
		fmt.Println("no row ids provided")
	}

	w.Write(mustMarshal(&models.Response{
		Total: 1,
		Data:  mustMarshal(models.UserList{created}),
	}))

	return 200, nil
}

func usersDelete(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.DeleteRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.User{}

	err := db.Delete(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &models.Response{
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

	w.Write(mustMarshal(&models.Response{
		Total: 1,
		Data:  mustMarshal(models.UserList{result}),
	}))
	return 200, nil
}

func usersDeleteMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.DeleteManyRequest{}
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

	w.Write(mustMarshal(&models.Response{
		Total: len(result),
		Data:  mustMarshal(result),
	}))

	return 200, nil
}
