package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	//jsonstruct "../../pkg/JSON-structs"
	entity "../../pkg/entities"
	storage "../../pkg/mysql"
)

var db storage.Datastore

var response entity.Response

func main() {
	// var port = flag.Int("listen", 8081, "the port number") //default port is 8081
	var dsn = flag.String("dsn",
		"admin:basketorangenumberbleacher@tcp(amarna-hacknyu.cfrwuvvgirag.us-east-2.rds.amazonaws.com:3306)/amarna",
		"sets data source name for backend server")

	flag.Parse()

	if *dsn == "" {
		log.Fatalln("You need to provide a DSN using the -dsn flag")
	}

	var datastoreErr error

	db, datastoreErr = storage.NewDatastore(*dsn)

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
	http.HandleFunc("/GetPairings", GetPairingsHandler)
	http.HandleFunc("/GetLetters", GetLettersHandler)

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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Hello world!")
}

//ValidateUserHandler will handle decoding of JSON packages for user validation and deliver a result to the frontend
func ValidateUserHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	username := "gqo"
	isValid, ValidateUserErr := db.ValidateUser(username)

	if ValidateUserErr != nil {
		response = entity.Response{
			Error: ValidateUserErr,
		}
	} else {
		response = entity.Response{
			Validate: isValid,
		}
	}
	respMarsh, marshErr := json.Marshal(response)
	if marshErr != nil {
		fmt.Println("There has been a marshalling error:", marshErr)
	}
	fmt.Fprintf(w, string(respMarsh))
}

//GetLettersHandler will handle decoding of JSON pakcages for letter retreival and deliver a result to the frontend
func GetLettersHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	lUser := "gqo"
	rUser := "amvasquez"

	allLetters, GetLettersErr := db.GetLetters(lUser, rUser)

	if GetLettersErr != nil {
		response = entity.Response{
			Error: GetLettersErr,
		}
	} else {
		response = entity.Response{
			Letters: allLetters,
		}
	}
	respMarsh, marshErr := json.Marshal(response)
	if marshErr != nil {
		fmt.Println("There has been a marshalling error:", marshErr)
	}
	fmt.Fprintf(w, string(respMarsh))
}

//GetPairingsHandler will handle decoding of JSON pakcages for pairing retreival and deliver a result to the frontend
func GetPairingsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	username := "gqo"
	allUserPairs, GetPairingsErr := db.GetPairings(username)

	if GetPairingsErr != nil {
		response = entity.Response{
			Error: GetPairingsErr,
		}
	} else {
		response = entity.Response{
			Pairings: allUserPairs,
		}
	}
	respMarsh, marshErr := json.Marshal(response)
	if marshErr != nil {
		fmt.Println("There has been a marshalling error:", marshErr)
	}
	fmt.Fprintf(w, string(respMarsh))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//SendLetterHandler will handle decoding of JSON pakcages for letter sending and deliver a result to the frontend
func SendLetterHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	lUser := "gqo"
	rUser := "amvasquez"
	body := "THIS IS THE TEXT"
	SendLetterErr := db.SendLetter(lUser, rUser, body)

	if SendLetterErr != nil {
		response = entity.Response{
			Error: SendLetterErr,
		}
	} else {
		response = entity.Response{}
	}
	respMarsh, marshErr := json.Marshal(response)
	if marshErr != nil {
		fmt.Println("There has been a marshalling error:", marshErr)
	}
	fmt.Fprintf(w, string(respMarsh))
}

//GetLanguagesHandler will handle decoding of JSON pakcages for language retreival and deliver a result to the frontend
func GetLanguagesHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	allLanguages, GetLanguagesErr := db.GetLanguages()

	if GetLanguagesErr != nil {
		response = entity.Response{
			Error: GetLanguagesErr,
		}
	} else {
		response = entity.Response{
			Languages: allLanguages,
		}
	}
	respMarsh, marshErr := json.Marshal(response)
	if marshErr != nil {
		fmt.Println("There has been a marshalling error:", marshErr)
	}
	fmt.Fprintf(w, string(respMarsh))
}
