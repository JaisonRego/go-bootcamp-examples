package router

import (
	"github.com/gorilla/mux"
	"github.com/jaisonrego/mongoapi/controller/netflixcontroller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", netflixcontroller.GetMovies).Methods("GET")
	router.HandleFunc("/api/movie", netflixcontroller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", netflixcontroller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", netflixcontroller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", netflixcontroller.DeleteAll).Methods("DELETE")

	return router
}
