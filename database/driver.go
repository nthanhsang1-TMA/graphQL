package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"graphQL/services"
	"log"
	"os"
	"time"
)

func Connect() *services.DB {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	fmt.Println("load env successfully")
	env := os.Getenv("ENV")
	switch env {
	case "mongo":
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		return &services.DB{
			MongoClient: client,
		}
	case "pg":
		dsn := "host=localhost user=postgres password=S@ng29031998 dbname=graphql-sangnt-2021 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		// migrate
		return &services.DB{
			PgClient: db.Debug(),
		}
	default:
		log.Fatal("no env")
		return nil
	}
}
