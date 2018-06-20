package repository

import (
	"encoding/json"
	"net/http"
	"html/template"
  	"path"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "../database"
	. "../models"
	. "../config"

)
var db = GoodsDatabase{}
var config = Config{}



type Profile struct {
  Name    string
  Hobbies []string
}

func init() {
	config.Read()

	db.Server = config.Server
	db.Database = config.Database
	db.Connect()
}

// GET method: all goods from Catalog
func AllGoodsEndPoint(w http.ResponseWriter, r *http.Request) {
	goods, err := db.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, goods)
}

// GET method: a good by ID from Catalog
func FindGoodEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	good, err := db.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid good ID")
		return
	}
	respondWithJson(w, http.StatusOK, good)
}

// GET method: a good details in HTML view by ID
func FindGoodHtmlEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	good, err := db.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid good ID")
		return
	}

  fp := path.Join("templates", "index.html")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, good); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

}

// POST method: a new good to catalog from JSON
func CreateGoodEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var good Good
	if err := json.NewDecoder(r.Body).Decode(&good); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	good.ID = bson.NewObjectId()
	if err := db.Insert(good); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, good)
}

// PUT method: update an existing good by ID from JSON
func UpdateGoodEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var good Good
	if err := json.NewDecoder(r.Body).Decode(&good); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := db.Update(good); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE method: an existing good by ID
func DeleteGoodEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var good Good
	if err := json.NewDecoder(r.Body).Decode(&good); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := db.Delete(good); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

//Default error JSON response
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

//Default 200 JSON response
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

