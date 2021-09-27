package config

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB *mongo.Client
)

const (
	host     = "192.168.187.129"
	port     = 5432
	user     = "postgres"
	password = "number1admin!"
	name     = "postgres"
)

func getDBType() string {
	return "postgres"
}

func getPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	return dataBase
}

// func getMongoConnect() {
// 	clientOptions := options.Client().ApplyURI("mongodb://192.168.187.129:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		panic(err)
// 	}

// 	DB = client
// }
