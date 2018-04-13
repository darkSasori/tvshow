package main

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var TVSHOW_MONGODB = os.Getenv("TVSHOW_MONGODB")

type MongoConnection struct {
	Session     *mgo.Session
	Connections int
}

func (mc *MongoConnection) NumConnections() int {
	return mc.Connections
}

func (mc *MongoConnection) Connect() {
	if mc.Connections == 0 {
		session, err := mgo.Dial(TVSHOW_MONGODB)
		if err != nil {
			log.Panicln(err)
		}
		mc.Session = session
	}
	mc.Connections++
}

func (mc *MongoConnection) Disconnect() {
	mc.Connections--
	if mc.NumConnections() <= 0 {
		mc.Session.Close()
		mc.Session = nil
	}
}

func (mc *MongoConnection) GetShows() *mgo.Collection {
	return mc.Session.DB("tvshow").C("shows")
}

var mgConnect *MongoConnection

func GetConnection() *MongoConnection {
	if mgConnect == nil {
		mgConnect = &MongoConnection{}
	}
	mgConnect.Connect()
	return mgConnect
}
