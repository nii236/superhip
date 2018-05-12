package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antonholmquist/jason"

	_ "github.com/lib/pq"
	"github.com/nii236/superhip/models"

	"github.com/go-chi/chi"
)

func roleRouter(db *DB) http.Handler {
	r := chi.NewRouter()

	r.Post("/get/list", withErrorAndDB(db, rolesGetList))
	r.Post("/get", withErrorAndDB(db, rolesGetOne))
	r.Post("/get/many", withErrorAndDB(db, rolesGetMany))
	r.Post("/get/many/reference", withErrorAndDB(db, rolesGetManyReference))

	r.Post("/create", withErrorAndDB(db, rolesCreate))

	r.Post("/update", withErrorAndDB(db, rolesUpdate))
	r.Post("/update/many", withErrorAndDB(db, rolesUpdateMany))

	r.Post("/delete", withErrorAndDB(db, rolesDelete))
	r.Post("/delete/many", withErrorAndDB(db, rolesDeleteMany))

	return r
}

func rolesGetList(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetListRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.RoleList{}
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
	total, err := db.Total("roles")
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

func rolesGetOne(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetOneRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.Role{}

	err := db.Read(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &models.Response{
			Total:   0,
			Data:    mustMarshal(models.RoleList{}),
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
		Data:  mustMarshal(models.RoleList{result}),
	}))
	return 200, nil
}

func rolesGetMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.RoleList{}

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

func rolesGetManyReference(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetManyReferenceRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.RoleList{}

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

func rolesUpdate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.UpdateRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	existing := &models.Role{}
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

	existing.Name, err = obj.GetString("name")
	if err != nil {
		return 500, fmt.Errorf("name: %s", err)
	}

	updated := &models.Role{}
	err = db.Update(updated, existing, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&models.Response{
		Total: 1,
		Data:  mustMarshal([]*models.Role{updated}),
	}))

	return 200, nil
}

func rolesUpdateMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.UpdateManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	updateTo := &models.Role{}
	obj, err := jason.NewObjectFromBytes(req.Data)
	if err != nil {
		return 500, err
	}

	updateTo.Name, err = obj.GetString("name")
	if err != nil {
		return 500, fmt.Errorf("name: %s", err)
	}

	IDs := []string{}
	for _, v := range req.IDs {
		IDs = append(IDs, v.String())
	}

	updated := models.RoleList{}

	err = db.UpdateMany(&updated, updateTo, IDs)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&models.Response{
		Total: len(updated),
		Data:  mustMarshal(updated),
	}))
	return 200, nil
}

func rolesCreate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return 500, fmt.Errorf("could not decode JSON: %s", err)
	}

	defer r.Body.Close()

	obj, err := jason.NewObjectFromBytes(req.Data)
	if err != nil {
		return 500, fmt.Errorf("could not create object from bytes: %s", err)
	}
	createWith := &models.Role{}
	createWith.Name, err = obj.GetString("name")
	if err != nil {
		return 500, fmt.Errorf("name: %s", err)
	}

	created := &models.Role{}
	err = db.Create(created, createWith)
	if err != nil && err == sql.ErrNoRows {
		return 404, fmt.Errorf("model not returned from db: %s", err)
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, fmt.Errorf("could not create model: %s", err)
	}

	w.Write(mustMarshal(&models.Response{
		Total: 1,
		Data:  mustMarshal(models.RoleList{created}),
	}))

	return 200, nil
}

func rolesDelete(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.DeleteRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.Role{}

	err := db.Delete(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &models.Response{
			Total:   0,
			Data:    mustMarshal(models.RoleList{}),
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
		Data:  mustMarshal(models.RoleList{result}),
	}))
	return 200, nil
}

func rolesDeleteMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.DeleteManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.RoleList{}

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
