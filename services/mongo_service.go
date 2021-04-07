package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"graphQL/graph/model"
	"log"
	"time"
)

func (db *DB) Save(input *model.NewPerson) *model.Person {
	collection := db.MongoClient.Database("gql-sangnt-2021").Collection("persons")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Person{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Name:      input.Name,
		IsGoodBoi: input.IsGoodBoi,
	}
}

func (db *DB) FindByID(ID string) *model.Person {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}

	collection := db.MongoClient.Database("gql-sangnt-2021").Collection("persons")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	person := model.Person{}
	collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&person)
	return &person
}

func (db *DB) FindAll(columns []string, where string) []*model.Person {
	collection := db.MongoClient.Database("gql-sangnt-2021").Collection("persons")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// get columns to sort
	var filter interface{}
	var persons []*model.Person
	if where != "" {
		filter = bson.M{columns[0]: primitive.Regex{Pattern: where}}
	} else {
		filter = bson.D{}
	}

	iter, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer iter.Close(ctx)
	for iter.Next(ctx) {
		var result model.Person
		err := iter.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, &result)
	}
	return persons
}
