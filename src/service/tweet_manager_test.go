package service_test

import (
	"testing"

	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	tweet := "This is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}

func TestPublishAndCleanTweet(t *testing.T) {
	tweet := "Esto es un re tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}

	service.CleanTweet()

	if service.GetTweet() != "" {
		t.Error("Tweet expected to be ''")
	}
}
