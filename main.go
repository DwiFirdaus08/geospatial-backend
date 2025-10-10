package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"posttest/geospatial-backend/config"
	"posttest/geospatial-backend/repository"
	"posttest/geospatial-backend/router"
)

func main() {
	config.ConnectDB()

	importData := flag.Bool("import", false, "run this flag to import geojson data to mongodb")
	flag.Parse()

	if *importData {
		log.Println("Starting data import...")
		if err := repository.ImportGeoJSONData(); err != nil {
			log.Fatalf("Error during data import: %v", err)
		}
		log.Println("Data import finished successfully.")
		return
	}

	router := router.SetupRouter()
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	log.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router)) 
}