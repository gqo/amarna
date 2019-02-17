package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
func ValidateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ValidateUser")
}
func UpdateKnownLangsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateKnownLangs")
}
func UpdateLearnLangsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateLearnLangs")
}
func InsertPairingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "InsertPairing")
}
func GetPairingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetPairings")
}
func GetMatchesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetMatches")
}
func GetSectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetSection")
}
func GetNextUncompletedSectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetNextUncompletedSection")
}

func main() {
	// var port = flag.Int("listen", 8081, "the port number") //default port is 8081
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

	// log.Println("Opening client connections on port", *port)
	// l, listenErr := net.Listen("tcp", ":"+strconv.Itoa(*port))
	// if listenErr != nil {
	// 	log.Fatalln("Could not listen to tcp port. Received err:", listenErr)
	// }
	// log.Println("Opened client connections!")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ValidateUser", ValidateUserHandler)
	http.HandleFunc("/UpdateKnownLangs", UpdateKnownLangsHandler)
	http.HandleFunc("/UpdateLearnLangs", UpdateLearnLangsHandler)
	http.HandleFunc("/ InsertPairing", InsertPairingHandler)
	http.HandleFunc("/GetPairings", GetPairingsHandler)
	http.HandleFunc("/GetMatches", GetMatchesHandler)
	http.HandleFunc("/GetSection", GetSectionHandler)
	http.HandleFunc("/GetNextUncompletedSection", GetNextUncompletedSectionHandler)

	log.Println("Starting Amarna backend...")
	// for {
	// 	conn, connErr := l.Accept()
	// 	if connErr != nil {
	// 		log.Println(connErr)
	// 	}

	// 	go handleRequest(conn)

	// }
	log.Fatal(http.ListenAndServe(":8080", nil))
}
