package main

/**
import (
    "net/http"
    "fmt"
    proto "github.com/golang/protobuf/proto"
    tvshowpb "github.com/darksasori/tvshow/tvshowpb"
    "log"
    "io/ioutil"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
session, err := mgo.Dial(os.Getenv("TVSHOW_MONGO"))
if err != nil {
    panic(err)
}

func persist(item *tvshowpb.TvShow, collection *mgo.Collection) {
    query := collection.Find(bson.M{
        "title": item.GetTitle(),
        "start": item.GetStart(),
        "channel.name": item.GetChannel().GetName(),
    })
    if count, _ := query.Count(); count > 0 {
        fmt.Println(item.GetTitle(), " Skiped")
        return
    }

    err := collection.Insert(item)
    if err != nil {
        panic(err)
    }
    fmt.Println(item.GetTitle(), " Inserted")
}

func ProtobufHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    item := new(tvshowpb.TvShow)
    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatalln("Error", err)
        return
    }

    if err := proto.Unmarshal(data, item); err != nil {
        log.Fatalln("Error", err)
        return
    }

    go persist(item, session.DB("tvshow").C("shows"))
}
*/

