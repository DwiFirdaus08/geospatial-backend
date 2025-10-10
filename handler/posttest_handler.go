package handler

import (
	"encoding/json"
	"net/http"
	"posttest/geospatial-backend/model"
	"posttest/geospatial-backend/repository"

	"github.com/gorilla/mux"
)

// GetAllJalanHandler (GET /api/jalans)
func GetAllJalanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jalans, err := repository.GetAllJalan()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(jalans)
}

// CreateJalanHandler (POST /api/jalans)
func CreateJalanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jalan model.Jalan
	if err := json.NewDecoder(r.Body).Decode(&jalan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertedID, err := repository.CreateJalan(jalan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"inserted_id": insertedID.Hex()})
}

// UpdateJalanHandler (PUT /api/jalans/{id})
func UpdateJalanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var jalan model.Jalan
	if err := json.NewDecoder(r.Body).Decode(&jalan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	modifiedCount, err := repository.UpdateJalan(id, jalan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if modifiedCount == 0 {
		http.Error(w, "No document found to update", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Successfully updated"})
}

// DeleteJalanHandler (DELETE /api/jalans/{id})
func DeleteJalanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	deletedCount, err := repository.DeleteJalan(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if deletedCount == 0 {
		http.Error(w, "No document found to delete", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Successfully deleted"})
}