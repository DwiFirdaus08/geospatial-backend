// File: api/index.go
package handler

import (
	"log"
	"net/http"
	"posttest/geospatial-backend/config"
	"posttest/geospatial-backend/router"
)

// Handler adalah fungsi utama yang akan dieksekusi oleh Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Function Handler started for path:", r.URL.Path)

	// --- LOGGING ---
	log.Println("Attempting to connect to the database...")
	config.ConnectDB()
	log.Println("Database connection function finished.") // Pesan ini mungkin tidak muncul jika ConnectDB panic

	// --- LOGGING ---
	log.Println("Setting up router...")
	mux := router.SetupRouter()
	log.Println("Router setup complete.")

	// Serahkan request ke router
	mux.ServeHTTP(w, r)
	log.Println("ServeHTTP finished.")
}