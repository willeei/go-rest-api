package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/williamsbarriquero/go-rest-api/database"
	"github.com/williamsbarriquero/go-rest-api/models"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "Home page")
	if err != nil {
		return
	}
}

func TodasPersonalidades(w http.ResponseWriter, _ *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		return
	}
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.First(&p, id)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Panic("Nenhuma personalidade encontrada!")
	}
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var p models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Panic("Falha ao realizar o decoder do request!")
	}
	database.DB.Create(&p)
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Panic("Falha ao realizar o encoder do response!")
	}
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.Delete(&p, id)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Panic("Falha ao realizar o encoder do response!")
	}
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Panic("Falha ao realizar o decoder do request!")
	}

	database.DB.Save(&p)
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Panic("Falha ao realizar o encoder do response!")
	}
}
