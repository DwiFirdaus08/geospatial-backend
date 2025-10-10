// File: api/index.go
package handler

import (
	"net/http"
	"posttest/geospatial-backend/config"
	"posttest/geospatial-backend/router"
)

// Handler adalah fungsi utama yang akan dieksekusi oleh Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Inisialisasi koneksi DB saat fungsi dipanggil
	
	config.ConnectDB()

	// Setup router yang sudah buat sebelumnya
	mux := router.SetupRouter()

	// Serahkan semua request ke router  untuk ditangani
	mux.ServeHTTP(w, r)
}