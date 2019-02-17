package entities

import (
	"fmt"
	"time"
)

// JSONTime is a type alias so we can more easily deal with time in Go -> JSON
type JSONTime time.Time

// MarshalJSON lets you convert time.Time to JSON data
func (t JSONTime) MarshalJSON() ([]byte, error) {
	ts := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2 15:04:05 2006"))
	return []byte(ts), nil
}

// Letter defines the letters that have been sent between users
type Letter struct {
	Body      string   `json:"body"`
	Timestamp JSONTime `json:"timestamp"`
	From      string   `json:"from"`
}

// Lesson defines the characters of a lesson
type Lesson struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Section string `json:"section"`
	Desc    string `json:"description"`
}

// Response will be the JSON sent as a response to the frontend
type Response struct {
	Letters   []Letter `json:"letters"`
	Pairings  []string `json:"pairings"`
	Validate  bool     `json:"validate"`
	Error     error    `json:"error"`
	Lessons   Lesson   `json:"lesson"`
	Languages []string `json:"languages"`
}
