package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	nosql "../../pkg/nosql"
)

//Packet will hold JSON information for later decoding
type Packet struct {
	Action   string `json:"action"`
	CompName string `json:"compName"`
}

func handleRequest(c net.Conn) {
	defer c.Close()
	var header Packet
	// we create a decoder that reads directly from the socket
	decode := json.NewDecoder(c)

	decodeErr := decode.Decode(&header)
	if decodeErr != nil {
		log.Println(decodeErr)
	}

	fmt.Printf("Action: %s, CompName: %s\n", header.Action, header.CompName)

}

func main() {
	var port = flag.Int("listen", 8081, "the port number") //default port is 8081
	l, listenErr := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if listenErr != nil {
		log.Panicln(listenErr)
	}

	log.Println("Opening client connections on port", *port)

	dsn := "mongodb://admin:basketorangenumberbleacher@amarna-shard-00-00-gcgag.gcp.mongodb.net:27017,amarna-shard-00-01-gcgag.gcp.mongodb.net:27017,amarna-shard-00-02-gcgag.gcp.mongodb.net:27017/test?ssl=true&replicaSet=Amarna-shard-0&authSource=admin&retryWrites=true"
	_, datastoreErr := nosql.NewDatastore(dsn)
	if datastoreErr != nil {
		log.Fatalln("Could not create datastore. Received err:", datastoreErr)
	}
	log.Println("Connected to the datastore!")

	log.Println("Starting Amarna backend...")

	for {
		conn, connErr := l.Accept()
		if connErr != nil {
			log.Println(connErr)
		}

		go handleRequest(conn)

	}
}
