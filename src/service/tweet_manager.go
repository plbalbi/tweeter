package service

import (
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
