package router

import (
	controller "hello/25-MongoApi/Controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
    router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovie).Methods("DELETE")
	return router
}