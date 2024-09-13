package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/SAHIL-Sharma21/mongoAPI/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}
	//filtering and updating....
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)
}

// delete 1 record from the database
func deleteOneMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("movie deleted successfully!", deleteCount.DeletedCount)
}

// delete all record
func deleteAllMovie() {
	// filter := bson.D{{}} // everything is selected and will be deleted
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted all record and with delete count is", deleteCount.DeletedCount)
}

// get all movies from database
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var movies []primitive.M

	//loop through inside it
	for cursor.Next(context.Background()) {
		var movie bson.M

		err := cursor.Decode(&movie) //if err is there then it will fatal otherwise we will push to movies slice

		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}
