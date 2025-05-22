package db

import "time"

type SensorDTO struct {
	SensorID     string     `json:"sensorID" bson:"sensorID"`
	LightStatus  string     `json:"lightStatus" bson:"lightStatus"`
	FireDetector string     `json:"fireDetector" bson:"fireDetector"`
	Sensors      []Sensor   `json:"sensors" bson:"sensors"`
	CreatedAt    *time.Time `json:"createdAt" bson:"createdAt"`
	DeletedAt    *time.Time `json:"deletedAt" bson:"deletedAt"`
	UpdatedAt    *time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Sensor struct {
	Name   string  `json:"name" bson:"name"`
	Value  float64 `json:"value" bson:"value"`
	Status string  `json:"status" bson:"status"`
	Unit   string  `json:"unit" bson:"unit"`
}

type SensorThreshold struct {
	Name      string     `json:"name" bson:"name"`
	Unit      string     `json:"unit" bson:"unit"`
	Threshold float64    `json:"threshold" bson:"threshold"`
	CreatedAt *time.Time `json:"createdAt" bson:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt" bson:"deletedAt"`
	UpdatedAt *time.Time `json:"updatedAt" bson:"updatedAt"`
}
