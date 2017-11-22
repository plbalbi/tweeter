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

	tweet = domain.NewTweet(user, text)

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

	tweet := domain.NewTweet(user, text)

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

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string
	text := "this is a tweet"
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string = "hola"
	text := ""
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is user required")
	}
}

func TestTweetTooLongIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string = "hola"
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent varius feugiat ex, ac finibus ex elementum egestas. Donec et massa posuere.`
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("error expected")
	}
	if err != nil && err.Error() != "text longer that 140 characters" {
		t.Error("tweet to long error expected")
	}
}
