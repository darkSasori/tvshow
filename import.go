package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// Persist get instance of TvShow and save in mongodb
func Persist(item TvShow) {
	connection := getConnection()
	defer connection.disconnect()

	query := connection.getShows().Find(bson.M{
		"title":        item.GetTitle(),
		"start":        item.GetStart(),
		"channel.name": item.GetChannel().GetName(),
	})
	if count, _ := query.Count(); count > 0 {
		log.Println(item.GetTitle(), " Skiped")
		return
	}

	err := connection.getShows().Insert(item)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(item.GetTitle(), " Inserted")
}

// ImportHandler handle request to call Persist
func ImportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, errors.New("Only POST is suported").Error(), http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	decode := json.NewDecoder(r.Body)

	var item TvShow
	err := decode.Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Try to import '%s'", item.GetTitle())
	go Persist(item)
}
