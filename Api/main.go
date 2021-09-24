package main

import (
	"fmt"
	"log"
	"net/http"

	// handles url/redirection
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"Api/HandleMembers"
	"Api/RestMethods"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", RestMethods.Get).Methods(http.MethodGet) //65
    r.HandleFunc("/", RestMethods.Post).Methods(http.MethodPost) //71
    r.HandleFunc("/", RestMethods.Put).Methods(http.MethodPut) //77
    r.HandleFunc("/", RestMethods.Patch).Methods(http.MethodPatch) //65
    r.HandleFunc("/", RestMethods.Delete).Methods(http.MethodDelete) // 83
    r.HandleFunc("/", RestMethods.NotFound) // 89

    // For Members:

    // Create Operation:
    r.HandleFunc("/member", HandleMembers.GetMembers).Methods(http.MethodPost)
    // Get Operation:
    r.HandleFunc("/member", HandleMembers.GetMembers).Methods(http.MethodGet)
    // Update Operation:
    r.HandleFunc("/member", HandleMembers.GetMembers).Methods(http.MethodGet)
    // Error 
    r.HandleFunc("/member", RestMethods.NotFound).Methods(http.MethodPatch)

    // To Handle CORS (Cross Origin Resource Sharing)
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS","DELETE", "CREATE"})
    origins := handlers.AllowedOrigins([]string{"*"})

    fmt.Println("Testing-Running on port : 8080")
    log.Fatal(http.ListenAndServe(":8080",handlers.CORS(headers,methods,origins)(r)))
}
