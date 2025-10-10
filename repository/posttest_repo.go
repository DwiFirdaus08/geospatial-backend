package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"posttest/geospatial-backend/config"
	"posttest/geospatial-backend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fungsi untuk import data dari file, menggunakan struct yang berbeda untuk file
func ImportGeoJSONData() error {
    collection := config.GetJalanCollection()
    filePath := "./data/jalan_kampung.json" 
    fileBytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("failed to read GeoJSON file: %w", err)
    }

    // Gunakan struct khusus untuk membaca dari file GeoJSON
    var featureCollectionFromFile model.GeoJSONFeatureCollectionForFile
    if err := json.Unmarshal(fileBytes, &featureCollectionFromFile); err != nil {
        return fmt.Errorf("failed to unmarshal GeoJSON: %w", err)
    }

    if _, err := collection.DeleteMany(context.TODO(), bson.D{}); err != nil {
        log.Printf("Warning: Could not clear existing data: %v", err)
    } else {
        fmt.Println("Cleared existing data in 'jalans' collection.")
    }

    var documents []interface{}
    for _, featureFromFile := range featureCollectionFromFile.Features {
        // Konversi dari GeoJSONFeatureForFile ke struct model.Jalan
        // Agar bisa disimpan dengan ID yang akan dibuat MongoDB
        jalanToInsert := model.Jalan{
            Type:       featureFromFile.Type,
            Properties: model.Properties{
                // Di sini Anda harus memetakan properti dari file ke struct Properties Anda
                // Contoh:
                NamaJalan: fmt.Sprintf("%v", featureFromFile.Properties["nama_jalan"]),
                // Kondisi: fmt.Sprintf("%v", featureFromFile.Properties["kondisi"]), // Jika ada properti 'kondisi'
            },
            Geometry:   featureFromFile.Geometry,
        }
        documents = append(documents, jalanToInsert)
    }

    if len(documents) > 0 {
        _, err := collection.InsertMany(context.TODO(), documents)
        if err != nil {
            return fmt.Errorf("failed to insert documents: %w", err)
        }
        fmt.Printf("Successfully inserted %d documents.\n", len(documents))
    } else {
        fmt.Println("No documents to insert.")
    }
    return nil
}

// --- FUNGSI CRUD ---

// GetAllJalan (READ)
func GetAllJalan() ([]model.Jalan, error) {
	collection := config.GetJalanCollection()
	var jalans []model.Jalan
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &jalans); err != nil {
		return nil, err
	}
	return jalans, nil
}

// CreateJalan (CREATE)
func CreateJalan(jalan model.Jalan) (*primitive.ObjectID, error) {
	collection := config.GetJalanCollection()
	result, err := collection.InsertOne(context.TODO(), jalan)
	if err != nil {
		return nil, err
	}
	insertedID := result.InsertedID.(primitive.ObjectID)
	return &insertedID, nil
}

// UpdateJalan (UPDATE)
func UpdateJalan(id string, jalan model.Jalan) (int64, error) {
	collection := config.GetJalanCollection()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("invalid id")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"properties": jalan.Properties,
		"geometry":   jalan.Geometry,
		"type":       jalan.Type,
	}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

// DeleteJalan (DELETE)
func DeleteJalan(id string) (int64, error) {
	collection := config.GetJalanCollection()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("invalid id")
	}

	filter := bson.M{"_id": objID}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}