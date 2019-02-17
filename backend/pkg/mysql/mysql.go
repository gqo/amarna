package mysql

import (
	"database/sql"
	"log"
	"time"

	entity "../entities"

	// This import provides the correct drivers for connecting to a mysql db
	_ "github.com/go-sql-driver/mysql"
)

// Datastore defines the behavior of Amarna database
type Datastore interface {
	// Version 2
	// High priority
	ValidateUser(username string) (bool, error)
	GetLetters(leftUsername, rightUsername string) ([]entity.Letter, error)
	GetPairings(username string) ([]string, error)
	SendLetter(leftUsername, rightUsername, body string) error
	// Medium priority
	GetCurrentLesson(leftUsername, rightUser string) (*entity.Lesson, error)
	IncrementLesson(leftUsername, rightUsername string) error
	RegisterUser(username, knownLang, learnLang string) error
	// Low priority
	GetLanguages() ([]string, error)
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

func (d *datastore) GetLetters(leftUsername, rightUsername string) ([]entity.Letter, error) {
	rows, err := d.db.Query(`
		SELECT leftUser, ts, body
		FROM Letter
		WHERE (leftUser=? AND rightUser=?)
		OR (rightUser=? AND leftUser=?)
		ORDER BY ts ASC`,
		leftUsername, rightUsername,
		leftUsername, rightUsername)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var letters []entity.Letter

	for rows.Next() {
		var Current entity.Letter

		err = rows.Scan(&Current.From, &Current.Timestamp, &Current.Body)
		if err != nil {
			return nil, err
		}

		letters = append(letters, Current)
	}

	return letters, nil
}

func (d *datastore) GetPairings(username string) ([]string, error) {
	rows, err := d.db.Query(`
		SELECT rightUser
		FROM Pairing
		WHERE leftUser=?`,
		username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pairings []string

	for rows.Next() {
		var Current string

		err = rows.Scan(&Current)
		if err != nil {
			return nil, err
		}

		pairings = append(pairings, Current)
	}

	return pairings, nil
}

func (d *datastore) SendLetter(leftUsername, rightUsername, body string) error {
	referenceID := 1 // Replace later

	_, err := d.db.Exec(`
		INSERT INTO Letter
		(leftUser, rightUser, referenceID, ts, body)
		VALUES
		(?,?,?,?,?)`,
		leftUsername, rightUsername, referenceID, time.Now(), body)
	if err != nil {
		return err
	}

	return nil
}

func (d *datastore) getCurrentLesson(leftUsername, rightUsername string) (*entity.Lesson, error) {
	row := d.db.QueryRow(`
		SELECT leftCount, rightCount
		FROM Pairing
		WHERE leftUser=?
		AND rightUser=?`,
		leftUsername, rightUsername)

	var leftCount, rightCount, referenceID int64

	err := row.Scan(&leftCount, &rightCount)

	switch {
	case err == nil:
		if leftCount < rightCount {
			referenceID = leftCount
		} else {
			referenceID = rightCount
		}
		row := d.db.QueryRow(`
			SELECT title, section, description
			FROM Lesson
			WHERE referenceID=?`,
			referenceID)

		var title, section, description string

		err := row.Scan(&title, &section, &description)

		switch {
		case err == nil:
			return &entity.Lesson{
				ID:      referenceID,
				Title:   title,
				Section: section,
				Desc:    description,
			}, nil
		default:
			return nil, err
		}
	default:
		return nil, err
	}
}

func (d *datastore) GetCurrentLesson(leftUsername, rightUsername string) (*entity.Lesson, error) {
	return d.getCurrentLesson(leftUsername, rightUsername)
}

func (d *datastore) IncrementLesson(leftUsername, rightUsername string) error {
	row := d.db.QueryRow(`
		SELECT leftCount, rightCount
		FROM Pairing
		WHERE leftUser=?
		AND rightUser=?`,
		leftUsername, rightUsername)

	var leftCount, rightCount int32

	err := row.Scan(&leftCount, &rightCount)

	switch {
	case err == nil:
		log.Println("LeftCount:", leftCount, "\nRightCount:", rightCount)
		if leftCount <= rightCount {
			leftCount++
			log.Println("New left count:", leftCount)

			_, err := d.db.Exec(`
				UPDATE Pairing
				SET leftCount=?
				WHERE leftUser=? AND rightUser=?`,
				leftCount,
				leftUsername, rightUsername)
			if err != nil {
				return err
			}

			_, err = d.db.Exec(`
				UPDATE Pairing
				SET rightCount=?
				WHERE rightUser=? AND leftUser=?`,
				leftCount,
				leftUsername, rightUsername)
			if err != nil {
				return err
			}

			return nil
		}
	default:
		return err
	}

	return nil
}

func (d *datastore) RegisterUser(username, knownLang, learnLang string) error {
	_, err := d.db.Exec(`
		INSERT INTO User
		(username, knownLang, learnLang)
		VALUES
		(?,?,?)`,
		username, knownLang, learnLang)
	if err != nil {
		return err
	}

	return nil
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
