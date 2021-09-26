package HandleIssuances

import (
	"context"
	"fmt"
	"log"
	"net"
	"reflect"
	"unsafe"

	// MongoDB drivers
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Issuance struct {
    Issuance_id  int `json:"issuance_id"`
    Book_id  int `json:"book_id"`
    Issuance_date  string `json:"issuance_date"`
    Issuance_member  int `json:"issuance_member"`
    Issued_by  string `json:"issued_by"`
    Target_return_date  string `json:"target_return_date"`
    Issuance_status  string `json:"issuance_status"`
}

type Issuances struct {
  Issuances []Issuance `json:"issuances"`
}

/*

Write function description here :

*/

func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

/*

Write function description here :

*/
func GetClientOptions() *options.ClientOptions {
  ipAddress := GetLocalIP()
  fmt.Println("Connecting to : "+ipAddress)
  // give the Mongo Atlas API here
  clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
  return clientOptions
}

/*

Write function description here :

*/
func GetClient (clientOptions *options.ClientOptions) *mongo.Client{
  client, err := mongo.Connect(context.TODO(), clientOptions)
  if err != nil {
    log.Fatal(err)
  }
  return client
}

/*

Write function description here :

*/
func GetCollection (client *mongo.Client,collectionname string) *mongo.Collection{
  collection := client.Database("LibraryAdmin").Collection(collectionname)
  return collection
}

/*

Write function description here :

*/
func BytesToString(b []byte) string {
    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
    sh := reflect.StringHeader{bh.Data, bh.Len}
    return *(*string)(unsafe.Pointer(&sh))
}
