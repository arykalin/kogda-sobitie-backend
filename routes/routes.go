package routes

import (
	"net/http"

	"github.com/arykalin/kogda-sobitie-backend/controllers"
	middlewares "github.com/arykalin/kogda-sobitie-backend/handlers"
	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/event", controllers.CreateEventEndpoint).Methods("POST")
	router.HandleFunc("/auth", controllers.Auths).Methods("GET")
	router.HandleFunc("/events", middlewares.IsAuthorized(controllers.GetEventsEndpoint)).Methods("GET")
	router.HandleFunc("/event/{id}", controllers.GetEventEndpoint).Methods("GET")
	router.HandleFunc("/event/{id}", controllers.DeleteEventEndpoint).Methods("DELETE")
	router.HandleFunc("/event/{id}", controllers.UpdateEventEndpoint).Methods("PUT")
	router.HandleFunc("/upload", controllers.UploadFileEndpoint).Methods("POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./uploaded/"))))
	return router
}
