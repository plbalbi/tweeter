package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

func NewTweet(user string, text string) *Tweet {
	actualTime := time.Now()
	tweet := Tweet{User: user, Text: text, Date: &actualTime}
	return &tweet
}
