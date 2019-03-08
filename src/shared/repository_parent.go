package shared

import mgo "gopkg.in/mgo.v2"

type Repository struct {
	DbMongo *mgo.Database
}

func NewRepository(dbMongo *mgo.Database) *Repository {
	return &Repository{DbMongo: dbMongo}
}
