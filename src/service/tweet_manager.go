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

func PublishTweet(t *domain.Tweet) (int, error) {
	if t.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if t.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(t.Text) > 140 {
		return 0, fmt.Errorf("text longer that 140 characters")
	}
	tweets = append(tweets, t)
	return len(tweets) - 1, nil
}

func CleanTweet() {
	tweets = nil
	InitializeService()
}

func GetTweets() [](*domain.Tweet) {
	return tweets
}

func NoTweets() bool {
	return len(tweets) == 0
}

func GetTweetById(id int) *domain.Tweet {
	return tweets[id]
}

func GetTweet() *domain.Tweet {
	if NoTweets() {
		return nil
	}
	return tweets[len(tweets)-1]
}
