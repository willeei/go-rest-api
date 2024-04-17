package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/williamsbarriquero/go-rest-api/models"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "Home page")
	if err != nil {
		return
	}
}

func TodasPersonalidades(w http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(w).Encode(models.Personalidades)
	if err != nil {
		return
	}
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			err := json.NewEncoder(w).Encode(personalidade)
			if err != nil {
				return
			}
		}
	}
}
