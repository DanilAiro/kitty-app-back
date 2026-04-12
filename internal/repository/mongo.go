package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/DanilAiro/kitty-app-back/internal/models"
)

var DB = connectDB()

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("%s:%s", os.Getenv("MONGO_URI"), os.Getenv("MONGO_PORT"))))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Cannot connect to MongoDB:", err)
	}

	return client
}

func GetUser(user *models.User) *models.User {
    collection := DB.Database(os.Getenv("MONGO_DB")).Collection("users")

    var result *models.User

    err := collection.FindOne(context.TODO(), map[string]interface{}{
        "email": user.Email,
    }).Decode(&result)
    if err != nil {
        return nil
    }

    return result
}

func AddUser(user *models.User) error {
    collection := DB.Database(os.Getenv("MONGO_DB")).Collection("users")

    _, err := collection.InsertOne(context.TODO(), user)

    if err != nil {
        return err
    }

    return nil
}
