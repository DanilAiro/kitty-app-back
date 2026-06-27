package initializers

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/DanilAiro/kitty-app-back/internal/repository"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

var Logger zerolog.Logger

type MongoLogWriter struct {
	collection *mongo.Collection
}

func NewMongoLogWriter(client *mongo.Client, dbName, collName string) io.Writer {
	return &MongoLogWriter{
		collection: client.Database(dbName).Collection(collName),
	}
}

func CreateLogger() {
	Logger = zerolog.New(NewMongoLogWriter(repository.DB, os.Getenv("MONGO_DB") + "_logs", "system_events")).With().Timestamp().Logger()
}

func (m *MongoLogWriter) Write(p []byte) (n int, err error) {
	var logDoc map[string]interface{}
	if err := json.Unmarshal(p, &logDoc); err != nil {
		return 0, err
	}

	if tStr, ok := logDoc["time"].(string); ok {
		if parsedTime, err := time.Parse(time.RFC3339, tStr); err != nil {
			logDoc["time"] = parsedTime
		}
	}

	go func(doc map[string]interface{})  {
		ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
		defer cancel()
		_, _ = m.collection.InsertOne(ctx, doc)
	} (logDoc)

	return len(p), nil
}