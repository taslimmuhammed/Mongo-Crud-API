package controller

import (
	"context"
	"encoding/json"
	"fmt"
	model "hello/25-MongoApi/Model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const url = "mongodb://localhost:27017"
const dbName = "Netflix"
const colName = "watchList"

//Most important
var collection *mongo.Collection

//connect wih mongoDb

func init(){
	//client options
	clientOptions:= options.Client().ApplyURI(url)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(),clientOptions)
    if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("succesfully connected to mongoDB")

	collection= client.Database(dbName).Collection(colName)
    //collection instance
	fmt.Println("collection instance is ready")
}

//MongoDB hepers -file

func insertOneMovie(movie model.NetFlix){
	inserted, err := collection.InsertOne(context.Background(), movie)
    if err!=nil{
     log.Fatal(err)
	}

	fmt.Println(inserted)
	fmt.Println("inserted one movie in DB with id: ", inserted.InsertedID)

}

//update 1 record
func updateOneMovie(movieId string){
	id, _:=primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Succefully deleted the movie with delete-count", deleteCount)
}

func deleteAllMovies() int64{
	deleteRes, err := collection.DeleteMany(context.Background(),bson.D{{}}) //giving null value  bson so that everythong will be selected
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Succefully deleted all  movie", deleteRes)
	return deleteRes.DeletedCount
}

func getAllMovies() []primitive.M{
	cur, err := collection.Find(context.Background(),bson.D{{}})
	if err !=nil{
		log.Fatal(err)
	}
	var movies[] primitive.M
    for cur.Next(context.Background()){
		var movie bson.M
		 err:= cur.Decode(&movie)
		 if err !=nil {
			log.Fatal(err)
		 }
		 movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
    return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie (w http.ResponseWriter, r *http.Request ) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.NetFlix
	_= json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)

}