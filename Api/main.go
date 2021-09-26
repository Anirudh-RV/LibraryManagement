package main

import (
	"fmt"
	"log"
	"net/http"

	// handles url/redirection
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"Api/HandleBooks"
	"Api/HandleIssuances"
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
    // POST Operation:
    r.Handle("/member", HandleJWT.IsAuthorized(HandleMembers.AddMembers)).Methods(http.MethodPost)
    // Get Operation:
    r.Handle("/member",HandleJWT.IsAuthorized(HandleMembers.GetAllMembers)).Methods(http.MethodGet)
    r.Handle("/member/{mem_id}",HandleJWT.IsAuthorized(HandleMembers.GetMembers)).Methods(http.MethodGet)
    // PUT Operation:
    r.Handle("/member", HandleJWT.IsAuthorized(HandleMembers.UpdateMembersPut)).Methods(http.MethodPut)

    // For Books:
    // POST Operation:
    r.Handle("/book", HandleJWT.IsAuthorized(HandleBooks.AddBooks)).Methods(http.MethodPost)
    // Get Operation:
    r.Handle("/book",HandleJWT.IsAuthorized(HandleBooks.GetAllBooks)).Methods(http.MethodGet)
    r.Handle("/book/{book_id}",HandleJWT.IsAuthorized(HandleBooks.GetBooks)).Methods(http.MethodGet)
    // PUT Operation:
    r.Handle("/book", HandleJWT.IsAuthorized(HandleBooks.UpdateBooksPut)).Methods(http.MethodPut)

    // For Issuances:
    // POST Operation:
    r.Handle("/issuance", HandleJWT.IsAuthorized(HandleIssuances.AddIssuance)).Methods(http.MethodPost)
    // Get Operation:
    r.Handle("/issuance",HandleJWT.IsAuthorized(HandleIssuances.GetAllIssuances)).Methods(http.MethodGet)
    r.Handle("/issuance/{issuance_id}",HandleJWT.IsAuthorized(HandleIssuances.GetIssuance)).Methods(http.MethodGet)
    // PUT Operation:
    r.Handle("/issuance", HandleJWT.IsAuthorized(HandleIssuances.UpdateIssuancePut)).Methods(http.MethodPut)

    // To Handle CORS (Cross Origin Resource Sharing)
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT","DELETE", "PATCH"})
    origins := handlers.AllowedOrigins([]string{"*"})

    fmt.Println("Testing-Running on port : 8080")
    log.Fatal(http.ListenAndServe(":8080",handlers.CORS(headers,methods,origins)(r)))
}
