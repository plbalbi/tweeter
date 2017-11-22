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
	if t.Text == "" {
		return fmt.Errorf("text is required")
	}
	if len(t.Text) > 140 {
		return fmt.Errorf("text longer that 140 characters")
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
