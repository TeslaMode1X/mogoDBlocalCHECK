package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"myMongoTest/pkg/config/model"
	"myMongoTest/pkg/config/utility"
	"net/http"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "ticketon"
const collectionName = "football_matches"

var collection *mongo.Collection

// this one is very important, because it will be done right after this file was triggered
func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection was established")
	collection = client.Database(dbName).Collection(collectionName)
}

// crud operations under

func getAllTickets() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var tickets []primitive.M

	for cursor.Next(context.Background()) {
		var ticket primitive.M
		err := cursor.Decode(&ticket)
		if err != nil {
			log.Fatal(err)
		}
		tickets = append(tickets, ticket)
	}

	return tickets
}

func createOneMatch(match *model.Match) {
	match.Date = utility.GenerateRandomDate()

	created, err := collection.InsertOne(context.Background(), match)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("One match was appointed", created)
}

func updateOneMatch(matchId string) {
	id, err := primitive.ObjectIDFromHex(matchId)
	if err != nil {
		log.Fatal(err) // Логирует ошибку и завершает программу
	}

	filter := bson.M{"_id": id}                   // Фильтр для поиска документа
	update := bson.M{"$set": bson.M{"was": true}} // Обновление документа

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err) // Логирует ошибку и завершает программу
	}
	fmt.Println("Successfully updated: ", res) // Выводит результат обновления
}

func deleteOneMatch(matchId string) {
	id, err := primitive.ObjectIDFromHex(matchId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("One record deleted!")
}

func deleteAllMovies() {
	delRes, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All records are deleted: ", delRes, "\n", delRes.DeletedCount)
}

func DeleteAllRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	deleteAllMovies()

	json.NewEncoder(w).Encode("All movies are deleted")
}

func DeleteOneMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMatch(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func UpdateOneMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMatch(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func CreateOneMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var match model.Match
	_ = json.NewDecoder(r.Body).Decode(&match)

	createOneMatch(&match)

	json.NewEncoder(w).Encode(match)
}

func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	err := json.NewEncoder(w).Encode(getAllTickets())
	if err != nil {
		log.Fatal(err)
	}
}
