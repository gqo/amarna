package JAYSONstructures

import (
	entity "../entities"
)

//both 'JSON' and 'struct' are keywords so I had to get creative

//ValUser defines struct for JSON packages recieved from frontend
type ValUser struct {
	Username string
}

//UpdateKnownLangs defines struct for JSON packages recieved from frontend
type UpdateKnownLangs struct {
	Username   string
	KnownLangs []string
}

//UpdateLearnLangs defines struct for JSON packages recieved from frontend
type UpdateLearnLangs struct {
	Username   string
	LearnLangs []string
}

//InsertPairing defines struct for JSON packages recieved from frontend
type InsertPairing struct {
	LeftUsername  string
	LeftUserLang  string
	RightUsername string
	RightUserLang string
}

//GetPairings defines struct for JSON packages recieved from frontend
type GetPairings struct {
	Username string
}

//GetMatches defines struct for JSON packages recieved from frontend
type GetMatches struct {
	User entity.User
}

//GetSection defines struct for JSON packages recieved from frontend
type GetSection struct {
	TopicTitle string
	TopicLang  string
}

//GetNextUncompletedSection defines struct for JSON packages recieved from frontend
type GetNextUncompletedSection struct {
	LeftUsername  string
	RightUsername string
	TopicTitle    string
	TopicLang     string
}

//JSONResponse will the the JSON used to respond to the frontend
type JSONResponse struct {
	Error   error
	Boolean *bool
	Section *entity.Section
}
