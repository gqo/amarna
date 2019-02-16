package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	mysql "../../pkg/mysql"
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
	var dsn = flag.String("dsn", "", "sets data source name for backend server")

	flag.Parse()

	if *dsn == "" {
		log.Fatalln("You need to provide a DSN using the -dsn flag")
	}

	_, datastoreErr := mysql.NewDatastore(*dsn)
	if datastoreErr != nil {
		log.Fatalln("Could not create datastore. Received err:", datastoreErr)
	}
	log.Println("Connected to the datastore!")

	log.Println("Opening client connections on port", *port)
	l, listenErr := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if listenErr != nil {
		log.Fatalln("Could not listen to tcp port. Received err:", listenErr)
	}
	log.Println("Opened client connections!")

	log.Println("Starting Amarna backend...")
	for {
		conn, connErr := l.Accept()
		if connErr != nil {
			log.Println(connErr)
		}

		go handleRequest(conn)

	}
}
