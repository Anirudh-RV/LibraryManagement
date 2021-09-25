package HandleBooks

import (
	"context"
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
func GetBooks(w http.ResponseWriter, r *http.Request) {
  pathParams := mux.Vars(r)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  // decoding the message and displaying
  fmt.Printf("ID to be queried : %s\n", pathParams["book_id"])

  // QUERYING MONGODB WITH name and returning the results
  // setting mongo variables with Collection : Member
  clientOptions := GetClientOptions()
  client := GetClient(clientOptions)
  collection := GetCollection(client,"Book")
  fmt.Println("Connected to MongoDB.")

  // add logic here :
  id, err:= strconv.Atoi(pathParams["book_id"])
  filterCursor, err := collection.Find(context.TODO(), bson.M{"book_id": id})
  if err != nil {
      log.Fatal(err)
  }
  var Result []bson.M
  if err = filterCursor.All(context.TODO(), &Result); err != nil {
      log.Fatal(err)
  }
  output := ""
  for _, document := range Result {
    for key, value := range document {
      fmt.Println(key, value)
      if key != "_id"{
        output += fmt.Sprintln("[",key,",",value,"]")
      }
    }
  }   
  w.Write([]byte(fmt.Sprintf(output)))
}
