package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func loadMongoConnection() *mgo.Database {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	mongoHost := fmt.Sprintf("%s:%s", host, port)
	mongoSession, err := mgo.Dial(mongoHost)
	if err != nil {
		log.Fatal(err)
	}

	db := mongoSession.DB("line_chatbot")

	// Init database collection, set unique index
	go func() {
		coll := db.C("users")
		index := mgo.Index{
			Key:    []string{"id"},
			Unique: true,
		}
		coll.EnsureIndex(index)
	}()

	return db
}

func loadPostgres() *sql.DB {
	return nil
}
