package HandleIssuances

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
func AddIssuance(w http.ResponseWriter, r *http.Request) {
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
  clientOptions := GetClientOptions()
  client := GetClient(clientOptions)
  collection := GetCollection(client,"Issuance")
  fmt.Println("Connected to MongoDB.")

  byteValue, _ := ioutil.ReadAll(r.Body)
  var users Issuances
  json.Unmarshal(byteValue, &users)
  for i := 0; i < len(users.Issuances); i++ {
    //TODO: Verfiy that the BookID and MemberID exist and then insert the value.
    
    // Inserting Data if Book and Member are Present
    insertResult, err := collection.InsertOne(context.TODO(), users.Issuances[i])
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
