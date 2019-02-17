package mysql

import (
	"database/sql"
	"log"

	entity "../entities"

	// This import provides the correct drivers for connecting to a mysql db
	_ "github.com/go-sql-driver/mysql"
)

// Datastore defines the behavior of Amarna database
type Datastore interface {
	// Version 2
	// High priority
	ValidateUser(username string) (bool, error)
	GetLetters(leftUsername, rightUsername string) (interface{}, error)
	GetPairings(username string) ([]string, error)
	// Medium priority
	GetCurrentLesson(leftUsername, rightUser string) (interface{}, error)
	IncrementLesson(leftUsername, rightUsername string) error
	RegisterUser(username, knownLang, learnLang string) error
	// Low priority
	GetLanguages() ([]string, error)

	// Version 1
	// High priority
	// ValidateUser(username string) (bool, error)  // Done
	// GetLanguages() ([]string, error)             // Done

	// // Medium priority
	// GetMatches(user entity.User) ([]entity.User, error)
	// InsertPairing(leftUsername, rightUsername, leftUserLang, rightUserLang string) error // Done                            // Done
	// // Low priority
	// UpdateKnownLangs(username string, knownLangs []string) error // Done
	// UpdateLearnLangs(username string, learnLangs []string) error // Done
	// // Takes a complete user's data
	// GetTopics(language string) ([]string, error) // Done
	// GetSection(topicTitle, topicLang string, week int) (*entity.Section, error)
	// GetNextUncompletedSection(leftUsername, rightUsername, topicTitle, topicLang string) (*entity.Section, error)
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

func (d *datastore) ValidateUser(username string) (bool, error) {
	row := d.db.QueryRow(`
		SELECT username
		FROM User
		WHERE username=?`,
		username)

	var temp string
	err := row.Scan(&temp)

	switch {
	case err == nil:
		return true, nil
	default:
		return false, err
	}
}

func (d *datastore) GetLanguages() ([]string, error) {
	rows, err := d.db.Query(`
		SELECT lang_name
		FROM Language`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []string

	for rows.Next() {
		var Current string

		err = rows.Scan(&Current)
		if err != nil {
			return nil, err
		}

		languages = append(languages, Current)
	}

	return languages, nil
}

func (d *datastore) GetTopics(language string) ([]string, error) {
	rows, err := d.db.Query(`
		SELECT title
		FROM Topic
		WHERE lang_name=?`,
		language)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topics []string

	for rows.Next() {
		var Current string

		err = rows.Scan(&Current)
		if err != nil {
			return nil, err
		}

		topics = append(topics, Current)
	}

	return topics, nil
}

func (d *datastore) insertKnownLang(username, knownLang string) error {
	_, err := d.db.Exec(`
		INSERT INTO KnownLang
		(username, lang_name)
		VALUES
		(?,?)`,
		username, knownLang)
	if err != nil {
		return err
	}

	return nil
}

func (d *datastore) UpdateKnownLangs(username string, knownLangs []string) error {
	_, err := d.db.Exec(`
		DELETE FROM KnownLang
		WHERE username=?`,
		username)
	if err != nil {
		return err
	}

	for i := range knownLangs {
		err = d.insertKnownLang(username, knownLangs[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *datastore) insertLearnLang(username, learnLang string) error {
	_, err := d.db.Exec(`
		INSERT INTO LearnLang
		(username, lang_name)
		VALUES
		(?,?)`,
		username, learnLang)
	if err != nil {
		return err
	}

	return nil
}

func (d *datastore) UpdateLearnLangs(username string, learnLangs []string) error {
	_, err := d.db.Exec(`
		DELETE FROM LearnLang
		WHERE username=?`,
		username)
	if err != nil {
		return err
	}

	for i := range learnLangs {
		err = d.insertLearnLang(username, learnLangs[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *datastore) InsertPairing(leftUsername, rightUsername, leftUserLang, rightUserLang string) error {
	_, err := d.db.Exec(`
		INSERT INTO Pairing
		(leftUSer, rightUser, leftUserLang, rightUserLang)
		VALUES
		(?,?,?,?)`,
		leftUsername, rightUsername, leftUserLang, rightUserLang)
	if err != nil {
		return err
	}

	return nil
}

func (d *datastore) GetPairings(username string) ([]entity.Pairing, error) {
	rows, err := d.db.Query(`
		SELECT DISTINCT username
		FROM Pairing
		WHERE leftUser=?`,
		username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pairings []entity.Pairing

	for rows.Next() {
		var Current entity.Pairing

		err = rows.Scan(&Current.RightUser)
		if err != nil {
			return nil, err
		}

		pairings = append(pairings, Current)
	}

	return pairings, nil
}

func (d *datastore) GetMatches(user entity.User) ([]entity.User, error) {
	return nil, nil
}

func (d *datastore) GetSection(topicTitle, topicLang string, week int) (*entity.Section, error) {
	return nil, nil
}

func (d *datastore) GetNextUncompletedSection(leftUsername, rightUsername, topicTitle, topicLang string) (*entity.Section, error) {
	return nil, nil
}
