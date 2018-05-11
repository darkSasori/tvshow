package main

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

// TvShowMongoDB save the mongodb uri
var TvShowMongoDB = os.Getenv("TVSHOW_MONGODB")

type mongoConnection struct {
	session     *mgo.Session
	connections int
}

func (mc *mongoConnection) numConnections() int {
	return mc.connections
}

func (mc *mongoConnection) connect() {
	if mc.connections == 0 {
		session, err := mgo.Dial(TvShowMongoDB)
		if err != nil {
			log.Panicln(err)
		}
		mc.session = session
	}
	mc.connections++
}

func (mc *mongoConnection) disconnect() {
	mc.connections--
	if mc.numConnections() <= 0 {
		mc.session.Close()
		mc.session = nil
	}
}

func (mc *mongoConnection) getShows() *mgo.Collection {
	return mc.session.DB("tvshow").C("shows")
}

var mgConnect *mongoConnection

func getConnection() *mongoConnection {
	if mgConnect == nil {
		mgConnect = &mongoConnection{}
	}
	mgConnect.connect()
	return mgConnect
}
