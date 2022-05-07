package controller

import (
	"api-go-rest/database"
	"api-go-rest/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []model.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality model.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var personality model.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Create(&personality)
	json.NewEncoder(w).Encode(personality)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality model.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality model.Personality
	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
}
