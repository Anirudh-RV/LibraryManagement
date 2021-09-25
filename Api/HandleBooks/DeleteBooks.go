package HandleBooks

import (
	"net/http"
)

/*

Write function description here :

*/
func DeleteBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
}
