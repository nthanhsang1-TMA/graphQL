package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type DB struct {
	MongoClient *mongo.Client
	PgClient    *gorm.DB
}
