package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	jsonstruct "../../pkg/JSON-structs"
	storage "../../pkg/mysql"
)

var db storage.Datastore
var response jsonstruct.JSONResponse

func main() {
	// var port = flag.Int("listen", 8081, "the port number") //default port is 8081
	var dsn = flag.String("dsn",
		"admin:basketorangenumberbleacher@tcp(amarna-hacknyu.cfrwuvvgirag.us-east-2.rds.amazonaws.com:3306)/amarna",
		"sets data source name for backend server")

	flag.Parse()

	if *dsn == "" {
		log.Fatalln("You need to provide a DSN using the -dsn flag")
	}

	var err error

	db, err = storage.NewDatastore(*dsn)

	if err != nil {
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
		log.Println(valUsrErr)
	}
	log.Println(valUsrJSON)
	isValid, ValidateUserErr := db.ValidateUser(valUsrJSON.Username)

	if ValidateUserErr != nil {
		response = jsonstruct.JSONResponse{
			Error:   ValidateUserErr,
			Boolean: nil,
			Section: nil,
		}
	} else {
		response = jsonstruct.JSONResponse{
			Error:   nil,
			Boolean: isValid,
			Section: nil,
		}
	}
	respMarsh, marshErr := json.Marshal(response)
	if marshErr != nil {
		fmt.Println("There has been a marshalling error:", marshErr)
	}
	fmt.Fprintf(w, string(respMarsh))
}

//UpdateKnownLangsHandler will handle decoding of JSON pakcages for known language updating
func UpdateKnownLangsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateKnownLangs")
	decoder := json.NewDecoder(r.Body)
	var updateKnownLangsJSON jsonstruct.UpdateKnownLangs
	updateKnownLangsErr := decoder.Decode(&updateKnownLangsJSON)
	if updateKnownLangsErr != nil {
		log.Println(updateKnownLangsErr)
	}
	log.Println(updateKnownLangsJSON)
	knownUpdateErr := db.UpdateKnownLangs(updateKnownLangsJSON.Username, updateKnownLangsJSON.KnownLangs)
	if knownUpdateErr != nil {
		response = jsonstruct.JSONResponse{
			Error:   knownUpdateErr,
			Boolean: nil,
			Section: nil,
		}
	} else {
		response = jsonstruct.JSONResponse{
			Error:   nil,
			Boolean: nil,
			Section: nil,
		}
	}
	respMarsh, marshErr := json.Marshal(response)
	if marshErr != nil {
		fmt.Println("There has been a marshalling error:", marshErr)
	}
	fmt.Fprintf(w, string(respMarsh))
}

//UpdateLearnLangsHandler will handle decoding of JSON pakcages for new or learned language updating
func UpdateLearnLangsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateLearnLangs")
	decoder := json.NewDecoder(r.Body)
	var updateLearnLangsJSON jsonstruct.UpdateLearnLangs
	updateLearnLangsErr := decoder.Decode(&updateLearnLangsJSON)
	if updateLearnLangsErr != nil {
		log.Println(updateLearnLangsErr)
	}
	log.Println(updateLearnLangsJSON)
	learnUpdateErr := db.UpdateLearnLangs(updateLearnLangsJSON.Username, updateLearnLangsJSON.LearnLangs)
	if learnUpdateErr != nil {
		response = jsonstruct.JSONResponse{
			Error:   learnUpdateErr,
			Boolean: nil,
			Section: nil,
		}
	} else {
		response = jsonstruct.JSONResponse{
			Error:   nil,
			Boolean: nil,
			Section: nil,
		}
	}
	respMarsh, marshErr := json.Marshal(response)
	if marshErr != nil {
		fmt.Println("There has been a marshalling error:", marshErr)
	}
	fmt.Fprintf(w, string(respMarsh))
}

//InsertPairingHandler will handle decoding of JSON pakcages for pair inserting
func InsertPairingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "InsertPairing")
	decoder := json.NewDecoder(r.Body)
	var InsertPairingJSON jsonstruct.InsertPairing
	InsertPairingErr := decoder.Decode(&InsertPairingJSON)
	if InsertPairingErr != nil {
		log.Println(InsertPairingErr)
	}
	log.Println(InsertPairingJSON)
	insertErr := db.InsertPairing(InsertPairingJSON.LeftUsername, InsertPairingJSON.RightUsername, InsertPairingJSON.LeftUserLang, InsertPairingJSON.RightUserLang)

}

//GetPairingsHandler will handle decoding of JSON pakcages for pairing retreival
func GetPairingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetPairings")
	decoder := json.NewDecoder(r.Body)
	var GetPairingJSON jsonstruct.GetPairings
	GetPairingErr := decoder.Decode(&GetPairingJSON)
	if GetPairingErr != nil {
		log.Println(GetPairingErr)
	}
	log.Println(GetPairingJSON)
	db.GetPairings(GetPairingJSON.Username)
}

//GetMatchesHandler will handle decoding of JSON pakcages for mathing retreival
func GetMatchesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetMatches")
	decoder := json.NewDecoder(r.Body)
	var GetMatchesJSON jsonstruct.GetMatches
	GetMatchesErr := decoder.Decode(&GetMatchesJSON)
	if GetMatchesErr != nil {
		log.Println(GetMatchesErr)
	}
	log.Println(GetMatchesJSON)
	db.GetMatches(GetMatchesJSON.User)
}

//GetSectionHandler will handle decoding of JSON pakcages for section retreival
func GetSectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetSection")
	decoder := json.NewDecoder(r.Body)
	var GetSectionJSON jsonstruct.GetSection
	GetSectionErr := decoder.Decode(&GetSectionJSON)
	if GetSectionErr != nil {
		log.Println(GetSectionErr)
	}
	log.Println(GetSectionJSON)
	db.GetSection(GetSectionJSON.TopicTitle, GetSectionJSON.TopicLang)
}

//GetNextUncompletedSectionHandler will handle decoding of JSON pakcages for retreival of upcoming and uncompleted sections
func GetNextUncompletedSectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetNextUncompletedSection")
	decoder := json.NewDecoder(r.Body)
	var GetNextUncompletedSectionJSON jsonstruct.GetNextUncompletedSection
	GetNextUncompletedSectionErr := decoder.Decode(&GetNextUncompletedSectionJSON)
	if GetNextUncompletedSectionErr != nil {
		log.Println(GetNextUncompletedSectionErr)
	}
	log.Println(GetNextUncompletedSectionJSON)
	db.GetNextUncompletedSection(GetNextUncompletedSectionJSON.LeftUsername, GetNextUncompletedSectionJSON.RightUsername, GetNextUncompletedSectionJSON.TopicTitle, GetNextUncompletedSectionJSON.TopicLang)
}
