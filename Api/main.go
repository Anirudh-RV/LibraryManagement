package main

import (
	"fmt"
	"log"
	"net/http"

	// handles url/redirection
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"Api/HandleJWT"
	"Api/HandleMembers"
)

func main() {
    // Generate the JWT Token for access of the APIs
    validToken, err := HandleJWT.GenerateJWT()
    fmt.Printf("Valid Token: %s\n", validToken)
    if err != nil{
        fmt.Println("Failed to generate token")
    }
    // Handle the requests
    r := mux.NewRouter()

    // For Members:
    // Create Operation:
    r.Handle("/member/", HandleJWT.IsAuthorized(HandleMembers.AddMembers)).Methods(http.MethodPost)
    // Get Operation:
    r.Handle("/member/{mem_id}",HandleJWT.IsAuthorized(HandleMembers.GetMembers)).Methods(http.MethodGet)
    // Update Operation:
    r.HandleFunc("/member/{mem_id}", HandleMembers.UpdateMembersPut).Methods(http.MethodPut)
    // Error 

    // To Handle CORS (Cross Origin Resource Sharing)
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT","DELETE", "PATCH"})
    origins := handlers.AllowedOrigins([]string{"*"})

    fmt.Println("Testing-Running on port : 8080")
    log.Fatal(http.ListenAndServe(":8080",handlers.CORS(headers,methods,origins)(r)))
}
