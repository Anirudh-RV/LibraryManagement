package HandleBooks

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*

Write function description here :
Serves PUT Requests to update an entry.
*/
func UpdateBooksPut(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("PUT QUERY CALLED")
  var ErrorFlag bool
  ErrorFlag = false

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  // Read body
	// Unmarshal
  if r.Body == nil {
      http.Error(w, "Please send a request body", 400)
      return
  }
  // setting mongo variables with Collection : Member
  clientOptions := GetClientOptions()
  client := GetClient(clientOptions)
  collection := GetCollection(client,"Book")
  fmt.Println("Connected to MongoDB.")

  byteValue, _ := ioutil.ReadAll(r.Body)
  var users Books
  json.Unmarshal(byteValue, &users)
  
  //fmt.Println("BODY: " + BytesToString(byteValue))
  for i := 0; i < len(users.Books); i++ {
    filter := bson.M{"book_id": users.Books[i].Book_id}
    update := bson.M{
      "$set": bson.M{
        "book_name":users.Books[i].Book_name,
        "book_cat_id":users.Books[i].Book_cat_id,
        "book_collection_id":users.Books[i].Book_collection_id,
        "book_launch_date":users.Books[i].Book_launch_date,
        "book_publisher":users.Books[i].Book_publisher,
      },
    }
    upsert := false
    after := options.After
    opt := options.FindOneAndUpdateOptions{
      ReturnDocument: &after,
      Upsert:         &upsert,
    }
    updateResult := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt)
    
    if updateResult.Err() != nil {
      fmt.Println("Error Occurred ", updateResult.Err())
      ErrorFlag = true
      break
    }
    fmt.Println("Updated document: ", updateResult)
  }
  
  // To close the connection to MongoDB
  err := client.Disconnect(context.TODO())
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Connection to MongoDB closed.")

  if ErrorFlag == true{
    w.Write([]byte(`{"message": "Error"}`))
  }else{
    w.Write([]byte(`{"message": "Success"}`))
  }
}

func UpdateBooksPatch(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
}
