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
	
	config.ConnectDB()
	log.Println("Database connection function finished.")

	mux := router.SetupRouter()
	log.Println("Router setup complete.")

	// Buat handler baru yang akan menghapus prefix /api
	// sebelum requestnya sampai ke router mux Anda.
	prefixStripper := http.StripPrefix("/api", mux)

	// Sekarang, jalankan request melalui prefixStripper
	prefixStripper.ServeHTTP(w, r)
	log.Println("ServeHTTP finished.")
}