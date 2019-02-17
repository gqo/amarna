package JAYSONstructures

import (
	entity "../entities"
)

//both 'JSON' and 'struct' are keywords so I had to get creative

//ValUser defines struct for JSON packages recieved from frontend
type ValUser struct {
	username string
}

//UpdateKnownLangs defines struct for JSON packages recieved from frontend
type UpdateKnownLangs struct {
	username   string
	knownLangs []string
}

//UpdateLearnLangs defines struct for JSON packages recieved from frontend
type UpdateLearnLangs struct {
	username   string
	learnLangs []string
}

//InsertPairing defines struct for JSON packages recieved from frontend
type InsertPairing struct {
	leftUsername  string
	leftUserLang  string
	rightUsername string
	rightUserLang string
}

//GetPairings defines struct for JSON packages recieved from frontend
type GetPairings struct {
	username string
}

//GetMatches defines struct for JSON packages recieved from frontend
type GetMatches struct {
	user entity.User
}

//GetSection defines struct for JSON packages recieved from frontend
type GetSection struct {
	topicTitle string
	topicLang  string
}

//GetNextUncompletedSection defines struct for JSON packages recieved from frontend
type GetNextUncompletedSection struct {
	leftUsername  string
	rightUsername string
	topicTitle    string
	topicLang     string
}
