package main

import (
    "os"
    "net/http"
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
)

func Persist(item TvShow) {
    session, err := mgo.Dial(os.Getenv("TVSHOW_MONGO"))
    if err != nil {
        panic(err)
    }
    defer session.Close()
    collection := session.DB("tvshow").C("shows")

    query := collection.Find(bson.M{
        "title": item.GetTitle(),
        "start": item.GetStart(),
        "channel.name": item.GetChannel().GetName(),
    })
    if count, _ := query.Count(); count > 0 {
        fmt.Println(item.GetTitle(), " Skiped")
        return
    }

    err = collection.Insert(item)
    if err != nil {
        panic(err)
    }
    fmt.Println(item.GetTitle(), " Inserted")
}

func ImportHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    decode := json.NewDecoder(r.Body)

    var item TvShow
    err := decode.Decode(&item)
    if err != nil {
        panic(err)
    }

    go Persist(item)
}
