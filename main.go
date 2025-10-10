package main

import (
	"flag"
	"log"
	"net/http"
	"posttest/geospatial-backend/config"
	"posttest/geospatial-backend/repository"
)

func main() {
	// Koneksi ke Database
	config.ConnectDB()

	// Tambahkan flag untuk proses import data
	importData := flag.Bool("import", false, "run this flag to import geojson data to mongodb")
	flag.Parse()

	if *importData {
		log.Println("Starting data import...")
		if err := repository.ImportGeoJSONData(); err != nil { // Panggilan sudah benar `repository.ImportGeoJSONData()`
			log.Fatalf("Error during data import: %v", err)
		}
		log.Println("Data import finished successfully.")
		return // Hentikan program setelah import selesai
	}

	// Setup Router
	router := route.SetupRouter() // <<< Ini sekarang dipanggil dari package 'route'
	
	// Jalankan Server
	port := ":8080"
	log.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}