package app

import (
	"context"
	"getir/inmemory"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

// App is general http resources.
type App struct {
	mongoDB    *mongo.Collection
	inMemoryDB *inmemory.DB
}

// Initialize is starting app.
func (app *App) Initialize() {
	mongoURL := os.Getenv("MONGO_DB_URL")
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	mongoCollection := os.Getenv("MONGO_COLLECTION_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
	}

	app.mongoDB = client.Database(mongoDBName).Collection(mongoCollection)
	app.inMemoryDB = inmemory.New()
}
