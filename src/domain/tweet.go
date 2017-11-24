package domain

import (
	"fmt"
	"time"
)

type Tweet struct {
	User        string
	Text        string
	Date        *time.Time
	Id          int
	IsRetweet   bool
	RetweetedBy string
}

func NewTweet(user string, text string) *Tweet {
	actualTime := time.Now()
	tweet := Tweet{User: user, Text: text, Date: &actualTime, IsRetweet: false}
	return &tweet
}

func NewRetweet(originalTweet *Tweet, retweeter string) *Tweet {
	actualTime := time.Now()
	tweet := Tweet{User: originalTweet.User, Text: originalTweet.Text, Date: &actualTime, IsRetweet: true, RetweetedBy: retweeter}
	return &tweet
}

func (tweet *Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}

func (tweet *Tweet) String() string {
	return tweet.PrintableTweet()
}
