package repository

import (
	"github.com/agungdwiprasetyo/line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/line-chatbot/src/shared"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repoEventMongo struct {
	db         *mgo.Database
	collection *mgo.Collection
}

// NewRepositoryEventMongo construct entry mongo repository
func NewRepositoryEventMongo(repo *shared.Repository) Event {
	repoEvent := new(repoEventMongo)
	repoEvent.db = repo.DbMongo
	repoEvent.collection = repo.DbMongo.C("events")
	return repoEvent
}

func (r *repoEventMongo) Save(data *domain.Event) error {
	data.ID = bson.NewObjectId()
	if err := r.collection.Insert(data); err != nil {
		return err
	}

	return nil
}

func (r *repoEventMongo) Count(filter *shared.Filter) int {
	q := bson.M{}
	query := r.collection.Find(q)
	count, _ := query.Count()
	return count
}

func (r *repoEventMongo) FindAll(filter *shared.Filter) ([]*domain.Event, error) {
	var events []*domain.Event
	q := bson.M{}
	query := r.collection.Find(q).Sort("-timestamp")
	if err := query.Skip(filter.Offset).Limit(filter.Limit).All(&events); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *repoEventMongo) RemoveAll() error {
	_, err := r.collection.RemoveAll(bson.M{})

	return err
}
