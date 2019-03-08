package repository

import (
	"time"

	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repoEntryMongo struct {
	db         *mgo.Database
	collection *mgo.Collection
}

// NewRepositoryEntryMongo construct entry mongo repository
func NewRepositoryEntryMongo(repo *shared.Repository) Entry {
	repoEntry := new(repoEntryMongo)
	repoEntry.db = repo.DbMongo
	repoEntry.collection = repo.DbMongo.C("entries")
	return repoEntry
}

func (r *repoEntryMongo) Create(data *domain.Entry) error {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	data.ID = bson.NewObjectId()
	data.CreatedAt = time.Now().In(loc)
	if err := r.collection.Insert(data); err != nil {
		return err
	}

	return nil
}
