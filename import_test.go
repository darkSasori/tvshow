package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetJson() *strings.Reader {
	return strings.NewReader("{\"channel\": {\"name\": \"channel1\",\"title\": \"Channel\",\"group\": \"group\",\"group_title\": \"Channel Group\",\"numbers\": {\"NET\": 666}},\"title\": \"GO GO LANG GO\",\"desc\": \"Film About Go Lang\",\"duraction\": 48.0,\"start\": \"2017-04-22T05:12:00\",\"end\": \"2017-04-22T06:00:00\"}")
}

func GetTvShow() (TvShow, error) {
	decode := json.NewDecoder(GetJson())

	var item TvShow
	err := decode.Decode(&item)
	return item, err
}

func ClearDb() {
	session, err := mgo.Dial(os.Getenv("TVSHOW_MONGO"))
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.DB("tvshow").C("shows").Remove(bson.M{})
}

func TestImportHandlerWithoutBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ImportHandler))
	defer server.Close()

	body := strings.NewReader("")
	res, err := http.Post(server.URL, "applicatoin/json", body)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 400 {
		t.Fatalf("Received non-400 response: %d\n", res.StatusCode)
	}
}

func TestImportHandlerMethodNotAllow(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ImportHandler))
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 405 {
		t.Fatalf("Received non-404 response: %d\n", res.StatusCode)
	}
}

func TestImportHandler(t *testing.T) {
	defer ClearDb()
	var buf bytes.Buffer
	log.SetOutput(&buf)

	server := httptest.NewServer(http.HandlerFunc(ImportHandler))
	defer server.Close()

	res, err := http.Post(server.URL, "applicatoin/json", GetJson())
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "Try to import 'GO GO LANG GO'" {
		t.Errorf("Expected 'Try to import 'GO GO LANG GO'', received '%s'", data)
	}

	time.Sleep(10 * time.Millisecond) // Wait for goroutine
}

func TestPersist(t *testing.T) {
	ClearDb()
	item, err := GetTvShow()
	if err != nil {
		t.Fatal(err)
		return
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)

	Persist(item)
	if m, _ := regexp.Match("Inserted", buf.Bytes()); !m {
		t.Errorf("Expected 'Inserted'")
	}
	buf.Reset()

	Persist(item)
	if m, _ := regexp.Match("Skiped", buf.Bytes()); !m {
		t.Errorf("Expected 'Skiped'")
	}

}

func BenchmarkImportHandler(b *testing.B) {
	item, err := GetTvShow()
	if err != nil {
		b.Fatal(err)
		return
	}

	for n := 0; n < b.N; n++ {
		Persist(item)
	}
}
