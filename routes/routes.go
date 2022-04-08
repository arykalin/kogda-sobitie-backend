package routes

import (
	middlewares "github.com/arykalin/kogda-sobitie-backend/handlers"
	"github.com/arykalin/kogda-sobitie-backend/internal/event_controller"
	"github.com/gorilla/mux"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/event", middlewares.IsAuthorized(controllers.CreateEventEndpoint)).Methods("POST")
	router.HandleFunc("/event", middlewares.IsAuthorized(event_controller.CreateEventEndpoint)).Methods("POST")
	router.HandleFunc("/auth", event_controller.Auths).Methods("GET")
	router.HandleFunc("/events", event_controller.GetEventsEndpoint).Methods("GET")
	router.HandleFunc("/event/{id}", event_controller.GetEventEndpoint).Methods("GET")
	router.HandleFunc("/event/{id}", middlewares.IsAuthorized(event_controller.DeleteEventEndpoint)).Methods("DELETE")
	router.HandleFunc("/event/{id}", middlewares.IsAuthorized(event_controller.UpdateEventEndpoint)).Methods("PUT")
	//router.HandleFunc("/upload", middlewares.IsAuthorized(controllers.UploadFileEndpoint)).Methods("POST")
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./uploaded/"))))
	return router
}
