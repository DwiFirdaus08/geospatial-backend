package router

import (
	"posttest/geospatial-backend/handler"
	"github.com/gorilla/mux"
)

// SetupRouter mengkonfigurasi semua rute untuk aplikasi
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// apiRouter := router.PathPrefix("/api").Subrouter()

	// apiRouter.HandleFunc("/jalans", handler.GetAllJalanHandler).Methods("GET")
	// apiRouter.HandleFunc("/jalans", handler.CreateJalanHandler).Methods("POST")
	// apiRouter.HandleFunc("/jalans/{id}", handler.UpdateJalanHandler).Methods("PUT")
	// apiRouter.HandleFunc("/jalans/{id}", handler.DeleteJalanHandler).Methods("DELETE")

	router.HandleFunc("/jalans", handler.GetAllJalanHandler).Methods("GET")
	router.HandleFunc("/jalans", handler.CreateJalanHandler).Methods("POST")
	router.HandleFunc("/jalans/{id}", handler.UpdateJalanHandler).Methods("PUT")
	router.HandleFunc("/jalans/{id}", handler.DeleteJalanHandler).Methods("DELETE")

	return router
}