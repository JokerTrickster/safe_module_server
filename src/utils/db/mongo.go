package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database

	// 컬렉션들
	LogsCollection    *mongo.Collection
	LightsCollection  *mongo.Collection
	SensorsCollection *mongo.Collection
)

const (
	DBName                = "safe_module"
	LogsCollectionName    = "logs"
	LightsCollectionName  = "lights"
	SensorsCollectionName = "sensors"
)

// InitMongoDB initializes MongoDB connection
func InitMongoDB() error {
	// MongoDB 연결 옵션 설정
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB 연결
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// 연결 테스트
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// 전역 변수 설정
	Client = client
	Database = client.Database(DBName)

	// 컬렉션 초기화
	LogsCollection = Database.Collection(LogsCollectionName)
	LightsCollection = Database.Collection(LightsCollectionName)
	SensorsCollection = Database.Collection(SensorsCollectionName)

	fmt.Println("몽고디비 연결 성공!")
	return nil
}

// CloseMongoDB closes MongoDB connection
func CloseMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		return err
	}

	log.Println("MongoDB connection closed.")
	return nil
}
