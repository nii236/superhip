package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/antonholmquist/jason"

	_ "github.com/lib/pq"
	"github.com/nii236/superhip/server/models"

	"github.com/go-chi/chi"
)

func schoolRouter(db *DB) http.Handler {
	r := chi.NewRouter()

	r.Post("/get/list", withErrorAndDB(db, schoolsGetList))
	r.Post("/get", withErrorAndDB(db, schoolsGetOne))
	r.Post("/get/many", withErrorAndDB(db, schoolsGetMany))
	r.Post("/get/many/reference", withErrorAndDB(db, schoolsGetManyReference))

	r.Post("/create", withErrorAndDB(db, schoolsCreate))

	r.Post("/update", withErrorAndDB(db, schoolsUpdate))
	r.Post("/update/many", withErrorAndDB(db, schoolsUpdateMany))

	r.Post("/delete", withErrorAndDB(db, schoolsDelete))
	r.Post("/delete/many", withErrorAndDB(db, schoolsDeleteMany))

	return r
}

func schoolsGetList(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetListRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.SchoolList{}

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

func schoolsGetOne(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetOneRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.School{}

	err := db.Read(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &Response{
			Total:   0,
			Data:    mustMarshal(models.SchoolList{}),
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
		Data:  mustMarshal(models.SchoolList{result}),
	}))
	return 200, nil
}

func schoolsGetMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.SchoolList{}

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

func schoolsGetManyReference(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetManyReferenceRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.SchoolList{}

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

func schoolsUpdate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &UpdateRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	existing := &models.School{}
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

	updated := &models.School{}
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
		Data:  mustMarshal([]*models.School{updated}),
	}))

	return 200, nil
}

func schoolsUpdateMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &UpdateManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	updateTo := &models.School{}
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

	updated := models.SchoolList{}

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

func schoolsCreate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
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
	createWith := &models.School{}
	createWith.Name, err = obj.GetString("name")
	if err != nil {
		return 500, fmt.Errorf("name: %s", err)
	}

	created := &models.School{}
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
		Data:  mustMarshal(models.SchoolList{created}),
	}))

	return 200, nil
}

func schoolsDelete(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &DeleteRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.School{}

	err := db.Delete(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &Response{
			Total:   0,
			Data:    mustMarshal(models.SchoolList{}),
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
		Data:  mustMarshal(models.SchoolList{result}),
	}))
	return 200, nil
}

func schoolsDeleteMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &DeleteManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.SchoolList{}

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
