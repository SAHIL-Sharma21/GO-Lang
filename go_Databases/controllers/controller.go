package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/SAHIL-Sharma21/mongoAPI/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Sahil:Sahil1234@cluster0.5ui1y4n.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongoDB
// init is the special function which will run at the very start of the application
// it will only run at one time!!!
func init() {
	//client options
	// clientoption := options.Client().ApplyURI(connectionString)
	clientoption := options.Client().ApplyURI(connectionString)

	//conect to mongo fb
	client, err := mongo.Connect(context.TODO(), clientoption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance is ready
	fmt.Println("Collection instance is ready!")
}

//MongoDB Helpers -> they are in seperate file

// inset 1
func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("whole inserted data", inserted)
	fmt.Println("Inserted one movie in DB with Id", inserted.InsertedID)
}
