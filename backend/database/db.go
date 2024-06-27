package database

import (
	"context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func intiDB(uri,dbName string){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/go-auth"))
	if err != nil {
        log.Fatal(err)
    }
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err!=nil{
		log.Fatal(err)
	}
	DB = client.Database(dbName)
}