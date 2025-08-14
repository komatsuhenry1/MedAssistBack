package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Client  *mongo.Client
	MongoDB *mongo.Database
)

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o .env")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	// Aqui você precisa definir a variável MongoDB, com o nome do banco de dados do seu .env, por exemplo:
	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		log.Fatal("MONGO_DB_NAME não está definido no .env")
	}

	MongoDB = Client.Database(dbName)

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func GetMongoDB() *mongo.Database {
	return MongoDB
}
