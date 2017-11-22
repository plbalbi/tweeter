package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweets [](*domain.Tweet)

func InitializeService() {
	// initialize empty slice
	tweets = make([](*domain.Tweet), 0)
}

func PublishTweet(t *domain.Tweet) error {
	if t.User == "" {
		return fmt.Errorf("user is required")
	}
	if t.Text == "" {
		return fmt.Errorf("text is required")
	}
	if len(t.Text) > 140 {
		return fmt.Errorf("text longer that 140 characters")
	}

	tweets = append(tweets, t)
	return nil
}

func CleanTweet() {
	tweets = nil
}

func GetTweets() [](*domain.Tweet) {
	return tweets
}

func NoTweets() bool {
	return len(tweets) == 0
}

func GetTweet() *domain.Tweet {
	if NoTweets() {
		return nil
	}
	return tweets[len(tweets)-1]
}
