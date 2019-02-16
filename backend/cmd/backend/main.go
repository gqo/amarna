package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"strconv"

	entity "../../pkg/entities"
	nosql "../../pkg/nosql"
)

//Packet will hold JSON information for later decoding
// type Packet struct {
// 	Action   string `json:"action"`
// 	CompName string `json:"compName"`
// }

//Message will hold frontend JSON information
//which will be used for backend parsiing and
//database use
type Message struct {
	MessageType string           `json:"messageType"`
	Users       []entity.User    `json:"users"`
	Letters     []entity.Letter  `json:"letters"`
	Pairings    []entity.Pairing `json:"pairings"`
	Topics      []entity.Topic   `json:"topics"`
	Sections    []entity.Section `json:"sections"`
	Week        int32            `json:"week"`
}

/*
Functionality needed:
GetPairings(user User) (*[]Pairing, error)
InsertPairing(left User, right User) (error)
GetMatches(user User) (*[]User, error)
GetTopics() (*[]Topic, error)
GetLanguages() (*[]Language, error)
GetConversationProgress(left User, right User, topic Topic) (*int32, error)
GetSection(topic Topic, week int32) (*Section, error)
ValidateUser(user User) (*bool, error)
UpdateKnownLangs(user User) (error)
UpdateLearnLangs(user User) (error)
*/
func handleJSON(header Message) {

	switch header.MessageType {
	case "GetPairings":
		for i := range header.Users {
			pairing, getPairError := GetPairings(header.Users[i])
			if getPairError != nil {
				log.Println("An error occured. Error:", getPairError)
			}
		}
	case "InsertPairing":
		for i := range header.Pairings {
			leftUser := header.Pairings[i].LeftUser
			rightUser := header.Pairings[i].RightUser
			insertPairErr := InsertPairing(leftUser, rightUser)
			if insertPairErr != nil {
				log.Println("An error occured. Error:", insertPairErr)
			}
		}
	case "GetMatches":
		for i := range header.Users {
			pairing, getMatchError := GetMatches(header.Users[i])
			if getPairError != nil {
				log.Println("An error occured. Error:", getPairError)
			}
		}
	case "GetTopics":
		allTopics, recvTopicsErr := GetTopics()
		if recvTopicsErr != nil {
			log.Println("An error occured. Error:", recvTopicsErr)
		}
	case "GetLanguages":
		allLanguages, recvLanguagesErr := GetLanguages()
		if recvLanguagesErr != nil {
			log.Println("An error occured. Error:", recvLanguagesErr)
		}
	case "GetSection":
		sectionPtr, getSecErr := GetSection(header.Topics[i], header.Week)
		if getSecErr != nil {
			log.Println("An error occured. Error:", getSecErr)
		}

	case "ValidateUser":
		for i := range header.Users {
			isValidated, validateError := ValidateUser(header.Users[i])
			if getPairError != nil {
				log.Println("An error occured. Error:", getPairError)
			}
		}
	case "UpdateKnownLangs":
		for i := range header.Users {
			updateKnownErr := UpdateKnownLangs(header.Users[i])
			if updateKnownErr != nil {
				log.Println("An error occured. Error:", updateKnownErr)
			}
		}

	case "UpdateLearnLangs":
		for i := range header.Users {
			updateLearnErr := UpdateLearnLangs(header.Users[i])
			if updateLearnErr != nil {
				log.Println("An error occured. Error:", updateLearnErr)
			}
		}
	}
}

func handleRequest(c net.Conn) {
	defer c.Close()
	var header Message
	decode := json.NewDecoder(c)

	decodeErr := decode.Decode(&header)
	if decodeErr != nil {
		log.Println(decodeErr)
	}
	handleJSON(header)
	//fmt.Printf("Action: %s, CompName: %s\n", header.Action, header.CompName)

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
