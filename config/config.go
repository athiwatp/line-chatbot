package config

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

// Config abstraction
type Config struct {
	MongoDB           *mgo.Database
	LineChannelSecret string
	LineChannelToken  string
	HTTPPort          string
}

// Init init config
func Init() *Config {
	conf := new(Config)

	conf.MongoDB = loadMongoConnection()
	conf.LineChannelSecret = os.Getenv("LINE_CHANNEL_SECRET")
	conf.LineChannelToken = os.Getenv("LINE_CHANNEL_TOKEN")
	conf.HTTPPort = os.Getenv("HTTP_PORT")

	return conf
}
