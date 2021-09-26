package HandleIssuances

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	// MongoDB drivers
)

/*

Write function description here :
Serves GET request for getting Books data
*/
func GetIssuance(w http.ResponseWriter, r *http.Request) {
  pathParams := mux.Vars(r)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  // decoding the message and displaying
  fmt.Printf("ID to be queried : %s\n", pathParams["issuance_id"])

  // QUERYING MONGODB WITH name and returning the results
  // setting mongo variables with Collection : Member
  clientOptions := GetClientOptions()
  client := GetClient(clientOptions)
  collection := GetCollection(client,"Issuance")
  fmt.Println("Connected to MongoDB.")

  // add logic here :
  id, err:= strconv.Atoi(pathParams["issuance_id"])
  filterCursor, err := collection.Find(context.TODO(), bson.M{"issuance_id": id})
  if err != nil {
      log.Fatal(err)
  }
  var Result []bson.M
  if err = filterCursor.All(context.TODO(), &Result); err != nil {
      log.Fatal(err)
  }
  json.NewEncoder(w).Encode(Result)
}


func GetAllIssuances(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  // decoding the message and displaying

  // QUERYING MONGODB WITH name and returning the results
  // setting mongo variables with Collection : Member
  clientOptions := GetClientOptions()
  client := GetClient(clientOptions)
  collection := GetCollection(client,"Issuance")
  fmt.Println("Connected to MongoDB.")

  // add logic here :
  filterCursor, err := collection.Find(context.TODO(), bson.D{{}})
  if err != nil {
      log.Fatal(err)
  }
  var Result []bson.M
  if err = filterCursor.All(context.TODO(), &Result); err != nil {
      log.Fatal(err)
  }
  json.NewEncoder(w).Encode(Result)
}