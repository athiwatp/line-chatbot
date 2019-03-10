package repository

import (
	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repoProfileMongo struct {
	db         *mgo.Database
	collection *mgo.Collection
}

// NewRepositoryProfileMongo construct entry mongo repository
func NewRepositoryProfileMongo(repo *shared.Repository) Profile {
	repoProfile := new(repoProfileMongo)
	repoProfile.db = repo.DbMongo
	repoProfile.collection = repo.DbMongo.C("profiles")
	return repoProfile
}

func (r *repoProfileMongo) Save(data *domain.Profile) error {
	query := bson.M{"id": data.ID}

	_, err := r.collection.Upsert(query, data)
	return err
}

func (r *repoProfileMongo) Count(filter *shared.Filter) int {
	q := bson.M{}
	query := r.collection.Find(q)
	count, _ := query.Count()
	return count
}

func (r *repoProfileMongo) FindAll(filter *shared.Filter) (res shared.Result) {
	var profiles []*domain.Profile

	q := bson.M{}
	if err := r.collection.Find(q).Skip(filter.Offset).Limit(filter.Limit).All(&profiles); err != nil {
		return shared.Result{Error: err}
	}

	res.Data = profiles
	return
}

func (r *repoProfileMongo) FindByID(id string) (res shared.Result) {
	var profileData domain.Profile
	query := bson.M{"id": id}
	if err := r.collection.Find(query).One(&profileData); err != nil {
		res.Error = err
		return
	}

	res.Data = profileData
	return
}
