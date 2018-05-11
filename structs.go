package main

import (
	"fmt"
	"strings"
	"time"
)

// CustomTime have this 'YYYY-MM-DDThh:mm:ss' when export/import from json
type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05"

// UnmarshalJSON set *ct to a copy of data
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

// MarshalJSON return ct as JSON encoding
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

// Channel have a information of tv channel
type Channel struct {
	Name, Title, Group, GroupTitle string
	Numbers                        map[string]int
}

// TvShow is a struct to be saved in mongodb
type TvShow struct {
	Channel     Channel
	Title, Desc string
	Start, End  CustomTime
	Duraction   float32
}

// GetTitle return a title of tvshow
func (t TvShow) GetTitle() string {
	return t.Title
}

// GetStart return CustomTime of tvshow start
func (t TvShow) GetStart() CustomTime {
	return t.Start
}

// GetChannel return channel of tvshow
func (t TvShow) GetChannel() Channel {
	return t.Channel
}

// GetName return name of tvshow
func (c Channel) GetName() string {
	return c.Name
}
