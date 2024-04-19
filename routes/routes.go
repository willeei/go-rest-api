package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/williamsbarriquero/go-rest-api/controllers"
	"github.com/williamsbarriquero/go-rest-api/middleware"
	"log"
	"net/http"
)

func HandleRequest() {
	const urlPersonalidadeId = "/api/personalidades/{id}"

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc(urlPersonalidadeId, controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc(urlPersonalidadeId, controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc(urlPersonalidadeId, controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
