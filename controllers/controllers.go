package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/williamsbarriquero/go-rest-api/models"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
