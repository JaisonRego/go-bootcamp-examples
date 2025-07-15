package netflixhelper

import (
	"context"
	"fmt"

	"github.com/jaisonrego/mongoapi/helper/errorhandler"
	"github.com/jaisonrego/mongoapi/model/netflix"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneMovie(movie netflix.Netflix, collection *mongo.Collection) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	errorhandler.LogError(err)
	fmt.Println("Inserted Record ID is: ", inserted.InsertedID)
}

func UpdateOneMovie(movieID string, collection *mongo.Collection) {
	id, err := primitive.ObjectIDFromHex(movieID)
	errorhandler.LogError(err)
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	errorhandler.LogError(err)
	fmt.Println("Modifed Record count:", result.ModifiedCount)
}

func DeleteOneMovie(movieID string, collection *mongo.Collection) {
	id, err := primitive.ObjectIDFromHex(movieID)
	errorhandler.LogError(err)
	deletedCount, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	errorhandler.LogError(err)
	fmt.Println("Deleted Record Count: ", deletedCount)
}

func DeleteAllMovies(collection *mongo.Collection) {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	errorhandler.LogError(err)
	fmt.Println("Deleted Record count: ", result.DeletedCount)
}

func GetAllMovies(collection *mongo.Collection) []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	errorhandler.LogError(err)

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		errorhandler.LogError(err)
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}
