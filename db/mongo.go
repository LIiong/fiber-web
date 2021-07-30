package db

import (
	"context"
	config2 "fiber-web/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Mg MongoInstance

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func InitMongo() {
	client, err := mongo.Connect(context.TODO(), clientOptions())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(config2.DatabaseConfig.Database)
	Mg = MongoInstance{
		Client: client,
		Db:     database,
	}
}

func clientOptions() *options.ClientOptions {
	sprintf := fmt.Sprintf("mongodb://%s:%s", config2.DatabaseConfig.Host, config2.DatabaseConfig.Port)
	return options.Client().ApplyURI(sprintf)
}

func DestroyMongo() {
	if err := Mg.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
