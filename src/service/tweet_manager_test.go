package service_test

import (
	"github.com/tweeter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	tweet := "This is my first tweet"
	service.tweet = "perro"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}
