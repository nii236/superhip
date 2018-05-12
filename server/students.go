package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/antonholmquist/jason"

	_ "github.com/lib/pq"
	"github.com/nii236/superhip/models"

	"github.com/go-chi/chi"
)

func studentRouter(db *DB) http.Handler {
	r := chi.NewRouter()

	r.Post("/get/list", withErrorAndDB(db, studentsGetList))
	r.Post("/get", withErrorAndDB(db, studentsGetOne))
	r.Post("/get/many", withErrorAndDB(db, studentsGetMany))
	r.Post("/get/many/reference", withErrorAndDB(db, studentsGetManyReference))

	r.Post("/create", withErrorAndDB(db, studentsCreate))

	r.Post("/update", withErrorAndDB(db, studentsUpdate))
	r.Post("/update/many", withErrorAndDB(db, studentsUpdateMany))

	r.Post("/delete", withErrorAndDB(db, studentsDelete))
	r.Post("/delete/many", withErrorAndDB(db, studentsDeleteMany))

	return r
}

func studentsGetList(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetListRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.StudentList{}

	err := db.List(&result)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	resp := &models.Response{
		Total: len(result),
		Data:  mustMarshal(result),
	}
	w.Write(mustMarshal(resp))
	return 200, nil
}

func studentsGetOne(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetOneRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.Student{}

	err := db.Read(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &models.Response{
			Total:   0,
			Data:    mustMarshal(models.StudentList{}),
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
		Data:  mustMarshal(models.StudentList{result}),
	}))
	return 200, nil
}

func studentsGetMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.StudentList{}

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

func studentsGetManyReference(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.GetManyReferenceRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.StudentList{}

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

func studentsUpdate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.UpdateRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	existing := &models.Student{}
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
	schoolFK, err := obj.GetString("school_id")
	if err != nil {
		return 500, fmt.Errorf("school_id: %s", err)
	}

	existing.SchoolID, err = uuid.FromString(schoolFK)
	if err != nil {
		return 500, err
	}

	updated := &models.Student{}

	err = db.Update(updated, existing, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&models.Response{
		Total: 1,
		Data:  mustMarshal([]*models.Student{updated}),
	}))

	return 200, nil
}

func studentsUpdateMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.UpdateManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	updateTo := &models.Student{}
	obj, err := jason.NewObjectFromBytes(req.Data)
	if err != nil {
		return 500, err
	}

	updateTo.Name, err = obj.GetString("name")
	if err != nil {
		return 500, fmt.Errorf("name: %s", err)
	}
	schoolFK, err := obj.GetString("school_id")
	if err != nil {
		return 500, fmt.Errorf("school_id: %s", err)
	}
	updateTo.SchoolID, err = uuid.FromString(schoolFK)
	if err != nil {
		return 500, err
	}

	IDs := []string{}
	for _, v := range req.IDs {
		IDs = append(IDs, v.String())
	}

	updated := models.StudentList{}

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

func studentsCreate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
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
	createWith := &models.Student{}
	createWith.Name, err = obj.GetString("name")
	if err != nil {
		return 500, fmt.Errorf("name: %s", err)
	}
	schoolFK, err := obj.GetString("school_id")
	if err != nil {
		return 500, fmt.Errorf("school_id: %s", err)
	}
	createWith.SchoolID, err = uuid.FromString(schoolFK)
	if err != nil {
		return 500, err
	}
	created := &models.Student{}
	err = db.Create(created, createWith)
	if err != nil && err == sql.ErrNoRows {
		return 404, err
	}
	if err != nil && err != sql.ErrNoRows {
		return 500, err
	}

	w.Write(mustMarshal(&models.Response{
		Total: 1,
		Data:  mustMarshal(models.StudentList{created}),
	}))

	return 200, nil
}

func studentsDelete(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.DeleteRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.Student{}

	err := db.Delete(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &models.Response{
			Total:   0,
			Data:    mustMarshal(models.StudentList{}),
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
		Data:  mustMarshal(models.StudentList{result}),
	}))
	return 200, nil
}

func studentsDeleteMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &models.DeleteManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.StudentList{}

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
