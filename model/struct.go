package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type GeoJSONFeatureForFile struct {
    Type       string                 `json:"type"`
    Properties map[string]interface{} `json:"properties"` // Ubah ke map[string]interface{}
    Geometry   Geometry               `json:"geometry"`
}

type GeoJSONFeatureCollectionForFile struct {
    Type     string                  `json:"type"`
    Features []GeoJSONFeatureForFile `json:"features"`
}

type Jalan struct {
    ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Type       string             `json:"type" bson:"type"`
    Properties Properties         `json:"properties" bson:"properties"`
    Geometry   Geometry           `json:"geometry" bson:"geometry"`
}

type Properties struct {
    NamaJalan string `json:"nama_jalan" bson:"nama_jalan"`
}

type Geometry struct {
    Type        string        `json:"type" bson:"type"`
    Coordinates interface{} `json:"coordinates" bson:"coordinates"` // Bisa LineString, Polygon, dll.
}