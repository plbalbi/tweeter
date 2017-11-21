package service_test

import (
	"github.com/tweeter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet string = "This is my first tweet"

	service.PublishTweet(tweet)

	if service.Tweet != tweet {
		t.Error("Expected tweet is", tweet)
	}
}
