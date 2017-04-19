package main

type channel struct {
    Name, Title, Group, Group_title string
    Numbers map[string]int
}

type TvShow struct {
    Channel channel
    Title, Desc string
    Start, End int
    Duraction float32
}

func (t TvShow)GetTitle()(string) {
    return t.Title
}

func (t TvShow)GetStart()(int) {
    return t.Start
}

func (t TvShow)GetChannel()(channel) {
    return t.Channel
}

func (c channel)GetName()(string) {
    return c.Name
}
