package service

import (
	"time"

	"github.com/tweeter/src/domain"
)

var tweet *domain.Tweet

func PublishTweet(t *domain.Tweet) {
	tweet = t
}

func CleanTweet() {
	tweet = nil
}

func GetTweet() *domain.Tweet {
	return tweet
}

func NewTweet(user string, text string) *domain.Tweet {
	actualTime := time.Now()
	tweet := domain.Tweet{User: user, Text: text, Date: &actualTime}
	return &tweet
}
