package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"quickstart/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

//var client *mongo.Client

func GetTasks(w http.ResponseWriter, r *http.Request) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	} 
	defer client.Disconnect(ctx)

	golangmongoDB := client.Database("golang-mongo-api")
	taskcollections := golangmongoDB.Collection("tasks")

	cursor, err := taskcollections.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var task bson.M
		if err = cursor.Decode(&task); err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, task["description"])
		fmt.Fprintf(w, "\n")
	}

	w.Header().Add("content-type", "application/json")

}

func AddTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	newtask := models.Task{}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	json.NewDecoder(r.Body).Decode(&newtask)

	golangmongoDB := client.Database("golang-mongo-api")
	taskss := golangmongoDB.Collection("tasks")

	doc := bson.M{"description": newtask.Description}
	result, err := taskss.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)

}
