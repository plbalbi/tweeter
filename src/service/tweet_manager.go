package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweet *domain.Tweet

func PublishTweet(t *domain.Tweet) error {
	if t.User == "" {
		return fmt.Errorf("user is required")
	}

	tweet = t
	return nil
}

func CleanTweet() {
	tweet = nil
}

func GetTweet() *domain.Tweet {
	return tweet
}
