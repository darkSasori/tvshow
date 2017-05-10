package main

import (
    "net/http"
    "fmt"
    "log"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
    "errors"
)

func Persist(item TvShow) {
    connection := GetConnection()
    defer connection.Disconnect()

    query := connection.GetShows().Find(bson.M{
        "title": item.GetTitle(),
        "start": item.GetStart(),
        "channel.name": item.GetChannel().GetName(),
    })
    if count, _ := query.Count(); count > 0 {
        log.Println(item.GetTitle(), " Skiped")
        return
    }

    err := connection.GetShows().Insert(item)
    if err != nil {
        log.Panicln(err)
    }
    log.Println(item.GetTitle(), " Inserted")
}

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
