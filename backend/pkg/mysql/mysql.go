package mysql

import (
	"database/sql"
	"log"

	// This import provides the correct drivers for connecting to a mysql db
	_ "github.com/go-sql-driver/mysql"
)

// Datastore is a wrapper for the mysql db connection
type Datastore struct {
	db *sql.DB
}

// NewDatastore constructs a new datastore object and returns said object
func NewDatastore(dsn string) (*Datastore, error) {
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

	d := &Datastore{
		db: db,
	}

	return d, nil
}
