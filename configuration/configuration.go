package configuration

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func Client(collection string) (*mongo.Collection, context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, _ := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)
	var col = client.Database(os.Getenv("DATABASE")).Collection(collection)
	return col, ctx
}
func InsertOne(collection string, obj interface{}) *mongo.InsertOneResult {
	col, ctx := Client(collection)
	res, _ := col.InsertOne(ctx, obj)
	return res
}
