package HandleMembers

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

type Member struct {
    Mem_id  int `json:"mem_id"`
    Mem_name  string `json:"mem_name"`
    Mem_phone  string `json:"mem_phone"`
    Mem_email  string `json:"mem_email"`
}

type Members struct {
  Members []Member `json:"members"`
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
