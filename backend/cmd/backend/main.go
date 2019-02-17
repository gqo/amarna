package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	entity "../../pkg/entities"
	storage "../../pkg/mysql"
)

var db storage.Datastore

var response entity.Response

func main() {
	var dsn = flag.String("dsn", "", "sets data source name for backend server")

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

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ValidateUser", ValidateUserHandler)
	http.HandleFunc("/GetLetters", GetLettersHandler)
	http.HandleFunc("/GetPairings", GetPairingsHandler)
	http.HandleFunc("/SendLetter", SendLetterHandler)
	http.HandleFunc("/GetCurrentLesson", GetCurrentLessonHandler)
	http.HandleFunc("/IncrementLesson", IncrementLessonHandler)
	http.HandleFunc("/RegisterUser", RegisterUserHandler)
	http.HandleFunc("/GetLanguages", GetLanguagesHandler)

	log.Println("Starting Amarna backend...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Hello world!")
}

//ValidateUserHandler will handle decoding of JSON packages for user validation and deliver a result to the frontend
func ValidateUserHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	urlParams := r.URL.Query()
	username := urlParams["username"][0]

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

	urlParams := r.URL.Query()
	lUser := urlParams["leftUsername"][0]
	rUser := urlParams["rightUsername"][0]

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

	urlParams := r.URL.Query()
	username := urlParams["username"][0]

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
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

//SendLetterHandler will handle decoding of JSON pakcages for letter sending and deliver a result to the frontend
func SendLetterHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	decoder := json.NewDecoder(r.Body)
	Data := struct {
		LeftUsername  string `json:"leftUsername"`
		RightUsername string `json:"rightUsername"`
		Body          string `json:"body"`
	}{}
	err := decoder.Decode(&Data)
	if err != nil {
		log.Println(err)
	}

	log.Println("data:", Data)

	lUser := Data.LeftUsername
	rUser := Data.RightUsername
	body := Data.Body

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

//GetCurrentLessonHandler will handle decoding of JSON pakcages for current lesson retreival and deliver a result to the frontend
func GetCurrentLessonHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	urlParams := r.URL.Query()
	lUser := urlParams["leftUsername"][0]
	rUser := urlParams["rightUsername"][0]

	lessonPtr, GetCurrentLessonErr := db.GetCurrentLesson(lUser, rUser)

	if GetCurrentLessonErr != nil {
		response = entity.Response{
			Error: GetCurrentLessonErr,
		}
	} else {
		response = entity.Response{
			Lessons: *lessonPtr,
		}
	}
	respMarsh, marshErr := json.Marshal(response)
	if marshErr != nil {
		fmt.Println("There has been a marshalling error:", marshErr)
	}
	fmt.Fprintf(w, string(respMarsh))
}

//IncrementLessonHandler will handle decoding of JSON packages for lesson incrementation and deliver a result to the frontend
func IncrementLessonHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	urlParams := r.URL.Query()
	lUser := urlParams["leftUsername"][0]
	rUser := urlParams["rightUsername"][0]

	IncrementLessonErr := db.IncrementLesson(lUser, rUser)

	if IncrementLessonErr != nil {
		response = entity.Response{
			Error: IncrementLessonErr,
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

//RegisterUserHandler will handle decoding of JSON pakcages for user registration and deliver a result to the frontend
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	urlParams := r.URL.Query()
	username := urlParams["username"][0]
	knownLang := urlParams["knownLang"][0]
	learnLang := urlParams["learnLang"][0]

	log.Println("User:", username, "\nKnownLang:", knownLang, "\nLearnLang:", learnLang)

	RegisterUserErr := db.RegisterUser(username, knownLang, learnLang)

	if RegisterUserErr != nil {
		response = entity.Response{
			Error: RegisterUserErr,
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
