package service_test

import (
	"testing"

	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

// Auxiliar functions

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.User != user && tweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.User, tweet.Text)
		return false
	}

	if id < 0 {
		t.Error("Expected id cannot be negative")
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

// --------------- TESTS ---------------

func TestPublishedTweetIsSaved(t *testing.T) {

	// init
	service.InitializeService()
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

	service.InitializeService()
	service.PublishTweet(tweet)

	if service.GetTweet().User != user ||
		service.GetTweet().Text != text {
		t.Errorf("Expected tweet is %s : %s\nbut it is %s : %s",
			user, text, tweet.User, tweet.Text)
	}

	if tweet.Date == nil {
		t.Error("Expected date cannot be nil")
	}

	service.CleanTweets()

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
	service.InitializeService()
	_, err = service.PublishTweet(tweet)

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
	service.InitializeService()
	_, err = service.PublishTweet(tweet)

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
	service.InitializeService()
	_, err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("error expected")
	}
	if err != nil && err.Error() != "text longer that 140 characters" {
		t.Error("tweet to long error expected")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, 0, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, 0, user, secondText) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	count := service.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	firstId, _ := service.PublishTweet(tweet)
	secondId, _ := service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	tweets := service.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}
