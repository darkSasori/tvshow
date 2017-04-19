package main

import (
    "time"
    "strings"
    "fmt"
)

type CustomTime struct {
    time.Time
}

const ctLayout = "2006-01-02T15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
    s := strings.Trim(string(b), "\"")
    if s == "null" {
       ct.Time = time.Time{}
       return
    }
    ct.Time, err = time.Parse(ctLayout, s)
    return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
  if ct.Time.UnixNano() == nilTime {
    return []byte("null"), nil
  }
  return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()
func (ct *CustomTime) IsSet() bool {
    return ct.UnixNano() != nilTime
}

type channel struct {
    Name, Title, Group, Group_title string
    Numbers map[string]int
}

type TvShow struct {
    Channel channel
    Title, Desc string
    Start, End CustomTime
    Duraction float32
}

func (t TvShow)GetTitle()(string) {
    return t.Title
}

func (t TvShow)GetStart()(CustomTime) {
    return t.Start
}

func (t TvShow)GetChannel()(channel) {
    return t.Channel
}

func (c channel)GetName()(string) {
    return c.Name
}
