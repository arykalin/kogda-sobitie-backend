package routes

import (
	"github.com/arykalin/kogda-sobitie-backend/controllers"
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
	router.HandleFunc("/event", controllers.CreateEventEndpoint).Methods("POST")
	router.HandleFunc("/auth", controllers.Auths).Methods("GET")
	router.HandleFunc("/events", controllers.GetEventsEndpoint).Methods("GET")
	router.HandleFunc("/event/{id}", controllers.GetEventEndpoint).Methods("GET")
	router.HandleFunc("/event/{id}", controllers.DeleteEventEndpoint).Methods("DELETE")
	router.HandleFunc("/event/{id}", controllers.UpdateEventEndpoint).Methods("PUT")
	//router.HandleFunc("/upload", middlewares.IsAuthorized(controllers.UploadFileEndpoint)).Methods("POST")
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./uploaded/"))))
	return router
}
