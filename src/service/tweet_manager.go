package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweets [](*domain.Tweet)
var lastUser string
var lastCount int

func InitializeService() {
	// initialize empty slice
	lastUser = ""
	lastCount = 0
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
	if t.User == lastUser {
		lastCount++
	}
	t.Id = len(tweets) - 1
	tweets = append(tweets, t)
	return len(tweets) - 1, nil
}

func CleanTweets() {
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

func CountTweetsByUser(user string) int {
	if lastUser == user {
		return lastCount
	}
	count := 0
	for _, tweet := range GetTweets() {
		if tweet.User == user {
			count++
		}
	}
	return count
}
