package entities

import (
	"fmt"
	"time"
)

// User defines the user in the Amarna service
type User struct {
	Username   string   `json:"username"`
	KnownLangs []string `json:"known_langs"`
	LearnLangs []string `json:"learn_langs"`
}

// JSONTime is a type alias so we can more easily deal with time in Go -> JSON
type JSONTime time.Time

// MarshalJSON lets you convert time.Time to JSON data
func (t JSONTime) MarshalJSON() ([]byte, error) {
	ts := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2 15:04:05 2006"))
	return []byte(ts), nil
}

// Letter defines the letters that have been sent between users
type Letter struct {
	From      string   `json:"from"`
	To        string   `json:"to"`
	Timestamp JSONTime `json:"timestamp"`
	Topic     string   `json:"topic"`
	Week      int32    `json:"week"`
}

// Pairing defines pairs of users that have matched
type Pairing struct {
	LeftUser  string `json:"left_user"`
	RightUser string `json:"right_user"`
}

// Topic defines a topic for learning in Amarna
type Topic struct {
	Title  string `json:"title"`
	Length string `json:"length"`
}

// Section defines a section of a topic
type Section struct {
	TopicTitle string `json:"topic_title"`
	Week       int32  `json:"week"`
	Desc       string `json:"desc"`
}

//Language defines a language of a topic
type Language struct {
	LangName string `json:"lang-name"`
}
