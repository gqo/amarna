package entities

import "time"

// User defines the user in the Amarna service
type User struct {
	username   string
	knownLangs []string
	learnLangs []string
}

// Letter defines the letters that have been sent between users
type Letter struct {
	from      string
	to        string
	timestamp time.Time
	topic     string
	week      int32
}

// Pairing defines pairs of users that have matched
type Pairing struct {
	leftUser  string
	rightUser string
}

// Topic defines a topic for learning in Amarna
type Topic struct {
	title  string
	length string
}

// Section defines a section of a topic
type Section struct {
	topicTitle string
	week       int32
	desc       string
}
