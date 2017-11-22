package service_test

import (
	"testing"

	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// init
	var tweet *domain.Tweet

	user := "perro"
	text := "guau"

	tweet = service.NewTweet(user, text)

	// op
	service.PublishTweet(tweet)

	// validate
	publishedTweet := service.GetTweet()

	if publishedTweet.User != user ||
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s : %s\nbut it is %s : %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date cannot be nil")
	}

}

func TestPublishAndCleanTweet(t *testing.T) {
	user := "caballo"
	text := "igihhiih"

	tweet := service.NewTweet(user, text)

	service.PublishTweet(tweet)

	if service.GetTweet().User != user ||
		service.GetTweet().Text != text {
		t.Errorf("Expected tweet is %s : %s\nbut it is %s : %s",
			user, text, tweet.User, tweet.Text)
	}

	if tweet.Date == nil {
		t.Error("Expected date cannot be nil")
	}

	service.CleanTweet()

	if service.GetTweet() != nil {
		t.Error("Tweet expected to be 'nil'")
	}
}
