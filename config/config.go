package config

import (
	"crypto/rsa"
	"os"

	mgo "gopkg.in/mgo.v2"
)

// Config abstraction
type Config struct {
	MongoDB           *mgo.Database
	LineChannelSecret string
	LineChannelToken  string
	HTTPPort          string
	Key               struct {
		APIKey             string
		Username, Password string
		PublicKey          *rsa.PublicKey
		PrivateKey         *rsa.PrivateKey
	}
}

// Init init config
func Init() *Config {
	conf := new(Config)

	conf.MongoDB = loadMongoConnection()
	conf.LineChannelSecret = os.Getenv("LINE_CHANNEL_SECRET")
	conf.LineChannelToken = os.Getenv("LINE_CHANNEL_TOKEN")
	conf.HTTPPort = os.Getenv("HTTP_PORT")

	conf.Key.APIKey = os.Getenv("API_KEY")
	conf.Key.Username = os.Getenv("USERNAME")
	conf.Key.Password = os.Getenv("PASSWORD")

	return conf
}
