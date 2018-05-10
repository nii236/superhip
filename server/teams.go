package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/antonholmquist/jason"

	_ "github.com/lib/pq"
	"github.com/nii236/superhip/server/models"

	"github.com/go-chi/chi"
)

func teamRouter(db *DB) http.Handler {
	r := chi.NewRouter()

	r.Post("/get/list", withErrorAndDB(db, teamsGetList))
	r.Post("/get", withErrorAndDB(db, teamsGetOne))
	r.Post("/get/many", withErrorAndDB(db, teamsGetMany))
	r.Post("/get/many/reference", withErrorAndDB(db, teamsGetManyReference))

	r.Post("/create", withErrorAndDB(db, teamsCreate))

	r.Post("/update", withErrorAndDB(db, teamsUpdate))
	r.Post("/update/many", withErrorAndDB(db, teamsUpdateMany))

	r.Post("/delete", withErrorAndDB(db, teamsDelete))
	r.Post("/delete/many", withErrorAndDB(db, teamsDeleteMany))

	return r
}

func teamsGetList(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetListRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.TeamList{}

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

func teamsGetOne(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetOneRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.Team{}

	err := db.Read(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &Response{
			Total:   0,
			Data:    mustMarshal(models.TeamList{}),
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
		Data:  mustMarshal(models.TeamList{result}),
	}))
	return 200, nil
}

func teamsGetMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.TeamList{}

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

func teamsGetManyReference(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &GetManyReferenceRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.TeamList{}

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

func teamsUpdate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &UpdateRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	existing := &models.Team{}
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

	updated := &models.Team{}
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
		Data:  mustMarshal([]*models.Team{updated}),
	}))

	return 200, nil
}

func teamsUpdateMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &UpdateManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	updateTo := &models.Team{}
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

	updated := models.TeamList{}

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

func teamsCreate(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
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
	createWith := &models.Team{}
	createWith.Name, err = obj.GetString("name")
	if err != nil {
		return 500, fmt.Errorf("name: %s", err)
	}
	schoolFK, err := obj.GetString("school_id")
	if err != nil {
		fmt.Printf("%+v", string(req.Data))
		return 500, fmt.Errorf("school_id: %s", err)
	}
	createWith.SchoolID, err = uuid.FromString(schoolFK)
	if err != nil {
		return 500, err
	}
	created := &models.Team{}
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
		Data:  mustMarshal(models.TeamList{created}),
	}))

	return 200, nil
}

func teamsDelete(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &DeleteRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := &models.Team{}

	err := db.Delete(result, req.ID.String())
	if err != nil && err == sql.ErrNoRows {
		resp := &Response{
			Total:   0,
			Data:    mustMarshal(models.TeamList{}),
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
		Data:  mustMarshal(models.TeamList{result}),
	}))
	return 200, nil
}

func teamsDeleteMany(db *DB, w http.ResponseWriter, r *http.Request) (int, error) {
	req := &DeleteManyRequest{}
	mustDecode(r.Body, req)

	defer r.Body.Close()

	result := models.TeamList{}

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
