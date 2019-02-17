package nosql

import (
	mongo "github.com/mongodb/mongo-go-driver/mongo"
)

// Datastore is a wrapper for the mongo db client connection
type Datastore struct {
	db *mongo.Client
}

// NewDatastore constructs a new datastore object and returns said object
func NewDatastore(dsn string) (*Datastore, error) {
	client, err := mongo.NewClient(dsn)
	if err != nil {
		return nil, err
	}
	return &Datastore{db: client}, nil
}
