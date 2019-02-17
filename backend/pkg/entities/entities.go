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

// type Letter struct {
// 	From       string   `json:"from"`
// 	To         string   `json:"to"`
// 	TopicTitle string   `json:"topic_title"`
// 	TopicLang  string   `json:"topic_lang"`
// 	Week       int32    `json:"week"`
// 	Timestamp  JSONTime `json:"timestamp"`
// 	Body       string   `json:"body"`
// }

// Pairing defines pairs of users that have matched
type Pairing struct {
	LeftUser      string `json:"left_user"`
	RightUser     string `json:"right_user"`
	LeftUserLang  string `json:"left_user_lang"`
	RightUserLang string `json:"right_user_lang"`
}

// Topic defines a topic for learning in Amarna
type Topic struct {
	Title  string `json:"title"`
	Length int32  `json:"length"`
}

// Section defines a section of a topic
type Section struct {
	TopicTitle string `json:"topic_title"`
	Week       int32  `json:"week"`
	Desc       string `json:"desc"`
}

// Language defines a language handled by Amarna
type Language struct {
	LangName string `json:"lang_name"`
}

// Letter defines the letters that have been sent between users
type Letter struct {
	Body      string
	Timestamp JSONTime
	From      string
}

// Lesson defines the characters of a lesson
type Lesson struct {
	ID      int64
	Title   string
	Section string
	Dest    string
}

// Prepare will be the JSON sent as a response to the frontend
type Prepare struct {
	Letters   []Letter
	Pairings  []string
	validate  bool
	Error     error
	Lessons   []Lesson
	Languages []string
}
