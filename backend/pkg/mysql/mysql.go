package mysql

import (
	"database/sql"
	"log"

	entity "../entities"

	// This import provides the correct drivers for connecting to a mysql db
	_ "github.com/go-sql-driver/mysql"
)

/*
Functionality needed:
GetPairings(user User) (*[]Pairing, error)
InsertPairing(left User, right User) (error)
GetMatches(user User) (*[]User, error)
GetTopics() (*[]Topic, error)
GetConversationProgress(left User, right User, topic Topic) (*int32, error)
GetSection(topic Topic, week int32) (*Section, error)
ValidateUser(user User) (*bool, error)
UpdateKnownLangs(user User) (error)
UpdateLearnLangs(user User) (error)
*/

// Datastore defines the behavior of Amarna database
type Datastore interface {
	ValidateUser(username string) (*bool, error)
	UpdateKnownLangs(username string, knownLangs []string) error
	UpdateLearnLangs(username string, learnLangs []string) error
	InsertPairing(leftUsername, rightUsername, leftUserLang, rightUserLang string) error
	GetPairings(username string) (*[]entity.Pairing, error)
	// Takes a complete user's data
	GetMatches(user entity.User) (*[]entity.User, error)
	GetSection(topicTitle, topicLang string) (*entity.Section, error)
	GetNextUncompletedSection(leftUsername, rightUsername, topicTitle, topicLang string) (*entity.Section, error)
}

// datastore is a wrapper for the mysql db connection
type datastore struct {
	db *sql.DB
}

// NewDatastore constructs a new datastore object and returns said object
func NewDatastore(dsn string) (Datastore, error) {
	log.Println("Creating new mysql datastore...")
	configParams := "?parseTime=true"
	fullDSN := dsn + configParams

	log.Println(fullDSN)

	db, err := sql.Open("mysql", fullDSN)
	if err != nil {
		log.Println("Could not open mysql connection")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Could not ping mysql connection")
		return nil, err
	}

	log.Println("Created new mysql datastore!")

	d := &datastore{
		db: db,
	}

	return d, nil
}

func (d *datastore) ValidateUser(username string) (*bool, error) {
	return nil, nil
}

func (d *datastore) UpdateKnownLangs(username string, knownLangs []string) error {
	return nil
}

func (d *datastore) UpdateLearnLangs(username string, knownLangs []string) error {
	return nil
}

func (d *datastore) InsertPairing(leftUsername, rightUsername, leftUserLang, rightUserLang string) error {
	return nil
}

func (d *datastore) GetPairings(username string) (*[]entity.Pairing, error) {
	return nil, nil
}

func (d *datastore) GetMatches(user entity.User) (*[]entity.User, error) {
	return nil, nil
}

func (d *datastore) GetSection(topicTitle, topicLang string) (*entity.Section, error) {
	return nil, nil
}

func (d *datastore) GetNextUncompletedSection(leftUsername, rightUsername, topicTitle, topicLang string) (*entity.Section, error) {
	return nil, nil
}
