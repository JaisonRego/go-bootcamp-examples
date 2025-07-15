package netflixcontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaisonrego/mongoapi/helper/errorhandler"
	"github.com/jaisonrego/mongoapi/helper/netflixhelper"
	"github.com/jaisonrego/mongoapi/model/netflix"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionUrl = "mongodb+srv://jaison:JAISON%40dragneel23@cluster0.eqryn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	errorhandler.LogError(err)
	fmt.Println("Succesfuly connected to mongo DB")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection Instance is ready")
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := netflixhelper.GetAllMovies(collection)
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie netflix.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	errorhandler.LogError(err)

	netflixhelper.InsertOneMovie(movie, collection)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	netflixhelper.UpdateOneMovie(params["id"], collection)
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	netflixhelper.DeleteOneMovie(params["id"], collection)
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	netflixhelper.DeleteAllMovies(collection)
	json.NewEncoder(w).Encode("Sucessfully Deleted all Records")
}
