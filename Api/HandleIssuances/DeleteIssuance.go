package HandleIssuances

import (
	"net/http"
)

/*

Write function description here :

*/
func DeleteIssuances(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
}
