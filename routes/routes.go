package routes

import (
	"log"
	"net/http"

	"api-go-rest/controller"
	"api-go-rest/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/", controller.Home).Methods("Get")
	r.HandleFunc("/api/personalities", controller.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controller.GetOne).Methods("Get")
	r.HandleFunc("/api/personalities", controller.Create).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controller.Delete).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controller.Edit).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
