package RestMethods

import (
	"net/http"
)

/*
Write function description here :
If you are trying to read a record, you should be using “GET.”
*/
func Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "GET called"}`))
}

/*
Write function description here :
This means that if you want to create a new record you should be using “POST.” 


*/
func Post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "POST called"}`))
}

/*
Write function description here :
To update a record utilizing “PUT” or “PATCH.” 
PUT is used to perform a complete override of the record (if fields are not included, they are removed) 
*/
func Put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "PUT called"}`))
}

/*
PATCH is used to update a record based on the information provided (or patch it with what you provide)
*/
func Patch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "PUT called"}`))
}

/*
Write function description here :
To delete a record, using “DELETE.”
*/
func Delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Delete called"}`))
}

/*
Write function description here :

*/
func NotFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "not found"}`))
}
