package db

import "time"

type SensorDTO struct {
	SensorID    string     `json:"sensorID" bson:"sensorID"`
	LightStatus *bool      `json:"lightStatus" bson:"lightStatus"`
	Sensors     []Sensor   `json:"sensors" bson:"sensors"`
	CreatedAt   *time.Time `json:"createdAt" bson:"createdAt"`
	DeletedAt   *time.Time `json:"deletedAt" bson:"deletedAt"`
	UpdatedAt   *time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Sensor struct {
	Name   string  `json:"name" bson:"name"`
	Value  float64 `json:"value" bson:"value"`
	Status string  `json:"status" bson:"status"`
}
