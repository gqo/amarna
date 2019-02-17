package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	jsonstruct "../../pkg/JSON-structs"
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

//ValidateUserHandler will handle decoding of JSON packages for user validation
func ValidateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ValidateUser")
	decoder := json.NewDecoder(r.Body)
	var valUsrJSON jsonstruct.ValUser
	valUsrErr := decoder.Decode(&valUsrJSON)
	if valUsrErr != nil {
		panic(valUsrErr)
	}
	log.Println(valUsrJSON)
}

//UpdateKnownLangsHandler will handle decoding of JSON pakcages for known language updating
func UpdateKnownLangsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateKnownLangs")
	decoder := json.NewDecoder(r.Body)
	var updateKnownLangsJSON jsonstruct.UpdateKnownLangs
	updateKnownLangsErr := decoder.Decode(&updateKnownLangsJSON)
	if updateKnownLangsErr != nil {
		panic(updateKnownLangsErr)
	}
	log.Println(updateKnownLangsJSON)
}

//UpdateLearnLangsHandler will handle decoding of JSON pakcages for new or learned language updating
func UpdateLearnLangsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateLearnLangs")
	decoder := json.NewDecoder(r.Body)
	var updateLearnLangsJSON jsonstruct.UpdateLearnLangs
	updateLearnLangsErr := decoder.Decode(&updateLearnLangsJSON)
	if updateLearnLangsErr != nil {
		panic(updateLearnLangsErr)
	}
	log.Println(updateLearnLangsJSON)
}

//InsertPairingHandler will handle decoding of JSON pakcages for pair inserting
func InsertPairingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "InsertPairing")
	decoder := json.NewDecoder(r.Body)
	var InsertPairingJSON jsonstruct.InsertPairing
	InsertPairingErr := decoder.Decode(&InsertPairingJSON)
	if InsertPairingErr != nil {
		panic(InsertPairingErr)
	}
	log.Println(InsertPairingJSON)
}

//GetPairingsHandler will handle decoding of JSON pakcages for pairing retreival
func GetPairingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetPairings")
	decoder := json.NewDecoder(r.Body)
	var GetPairingJSON jsonstruct.GetPairings
	GetPairingErr := decoder.Decode(&GetPairingJSON)
	if GetPairingErr != nil {
		panic(GetPairingErr)
	}
	log.Println(GetPairingJSON)
}

//GetMatchesHandler will handle decoding of JSON pakcages for mathing retreival
func GetMatchesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetMatches")
	decoder := json.NewDecoder(r.Body)
	var GetMatchesJSON jsonstruct.GetMatches
	GetMatchesErr := decoder.Decode(&GetMatchesJSON)
	if GetMatchesErr != nil {
		panic(GetMatchesErr)
	}
	log.Println(GetMatchesJSON)
}

//GetSectionHandler will handle decoding of JSON pakcages for section retreival
func GetSectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetSection")
	decoder := json.NewDecoder(r.Body)
	var GetSectionJSON jsonstruct.GetSection
	GetSectionErr := decoder.Decode(&GetSectionJSON)
	if GetSectionErr != nil {
		panic(GetSectionErr)
	}
	log.Println(GetSectionJSON)
}

//GetNextUncompletedSectionHandler will handle decoding of JSON pakcages for retreival of upcoming and uncompleted sections
func GetNextUncompletedSectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetNextUncompletedSection")
	decoder := json.NewDecoder(r.Body)
	var GetNextUncompletedSectionJSON jsonstruct.GetNextUncompletedSection
	GetNextUncompletedSectionErr := decoder.Decode(&GetNextUncompletedSectionJSON)
	if GetNextUncompletedSectionErr != nil {
		panic(GetNextUncompletedSectionErr)
	}
	log.Println(GetNextUncompletedSectionJSON)
}

func main() {
	// var port = flag.Int("listen", 8081, "the port number") //default port is 8081
	var dsn = flag.String("dsn",
		"admin:basketorangenumberbleacher@tcp(amarna-hacknyu.cfrwuvvgirag.us-east-2.rds.amazonaws.com:3306)/amarna",
		"sets data source name for backend server")

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
