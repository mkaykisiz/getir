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
	mongoDB    *mongo.Database
	inMemoryDB *inmemory.DB
}

// Initialize is starting app.
func (app *App) Initialize() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB_URL")))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
	}

	app.mongoDB = client.Database(os.Getenv("MONGO_DB_NAME"))
	app.inMemoryDB = inmemory.New()
}
