package main

import (
	"log"

	nosql "../../pkg/nosql"
)

func main() {
	dsn := "mongodb://admin:basketorangenumberbleacher@amarna-shard-00-00-gcgag.gcp.mongodb.net:27017,amarna-shard-00-01-gcgag.gcp.mongodb.net:27017,amarna-shard-00-02-gcgag.gcp.mongodb.net:27017/test?ssl=true&replicaSet=Amarna-shard-0&authSource=admin&retryWrites=true"
	_, err := nosql.NewDatastore(dsn)
	if err != nil {
		log.Fatalln("Could not create datastore. Received err:", err)
	}
	log.Println("Connected to the datastore!")

	log.Println("Starting Amarna backend...")
	for {

	}
}
