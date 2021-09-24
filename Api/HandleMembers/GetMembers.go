package HandleMembers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// MongoDB drivers
)

/*

Write function description here :

*/
func GetMembers(w http.ResponseWriter, r *http.Request) {
  pathParams := mux.Vars(r)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  // decoding the message and displaying
  reqBody, err := ioutil.ReadAll(r.Body)
  if err != nil {
   log.Fatal(err)
  }
  fmt.Printf("Name to be queried : %s\n", reqBody)
  fmt.Printf("Name to be queried : %s\n", pathParams["mem_id"])
  w.Write([]byte(fmt.Sprintf(`{"mem_id":"%s"}`, pathParams["mem_id"])))
}
