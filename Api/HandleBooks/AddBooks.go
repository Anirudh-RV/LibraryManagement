package HandleBooks

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*

Write function description here :
Serves POST requests for adding values to Book collection.
*/
func AddBooks(w http.ResponseWriter, r *http.Request) {
  var ErrorFlag bool
  ErrorFlag = false

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  // Read body
  fmt.Printf("POST QUERY CALLED")
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
    insertResult, err := collection.InsertOne(context.TODO(), users.Books[i])
    if err != nil {
      fmt.Println("Error Occurred ", err)
      ErrorFlag = true
      break
    }
    fmt.Println("Inserted document: ", insertResult.InsertedID)
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
