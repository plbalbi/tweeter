package service_test

import (
	"fmt"
	"testing"

	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

var tweetManager *service.TweetManager

func TestMain(m *testing.M) {
	tweetManager = service.NewTweetManager()
	m.Run()
}

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet

	user := "grupoesfera"
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
	Phasellus non purus eget lectus pretium mattis quis nec odio. Cras quis orci metuasds. `

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text longer that 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)

	// Operation
	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		t.Error("not valid tweet")
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		t.Error("not valid tweet")
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

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

func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user, text string) bool {

	if tweet.GetId() != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.GetId())
	}

	if tweet.GetUser() != user && tweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.GetUser(), tweet.GetText())
		return false
	}

	if tweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestTrendingTopicOk(t *testing.T) {
	tweetManager = service.NewTweetManager()
	tweet1 := domain.NewTextTweet("perro", "esto es re loco #dogchow #eukanuba #fritolin")
	tweet2 := domain.NewTextTweet("perro2", "esto es re loco #dogchow #eukanuba")
	tweetManager.PublishTweet(tweet1)
	tweetManager.PublishTweet(tweet2)

	trendingTopics := tweetManager.GetTrendingTopics()

	if len(trendingTopics) != 2 {
		t.Error("trendingTopic count expected to be 2")
		return
	}
	if trendingTopics[0] != "#dogchow" && trendingTopics[1] != "#eukanuba" {
		t.Error("bad trending topic")
		return
	}
	if trendingTopics[0] != "#dogchow" && trendingTopics[1] != "#eukanuba" {
		t.Error("bad trending topic")
		return
	}
}

func TestSendDirectMessageAndGetAllOfThem(t *testing.T) {
	tweetManager = service.NewTweetManager()
	msg1 := domain.NewMessage("perro2", "hola")
	msg2 := domain.NewMessage("perro3", "hola")
	msg3 := domain.NewMessage("perro4", "hola")
	tweetManager.SendDirectMessage(msg1, "perro1")
	tweetManager.SendDirectMessage(msg2, "perro1")
	tweetManager.SendDirectMessage(msg3, "perro1")

	allRecievedMessages := tweetManager.GetAllDirectMessages("perro1")

	if len(allRecievedMessages) != 3 {
		t.Error("Unexpected number of direct messages")
		return
	}

}
func TestGetUnreadMessagesAndRead(t *testing.T) {
	tweetManager = service.NewTweetManager()
	if tweetManager.GetTweet() != nil {
		t.Errorf("tweeterManager not beign cleaned")
	}

	msg1 := domain.NewMessage("perro2", "hola")
	msg2 := domain.NewMessage("perro3", "hola")
	msg3 := domain.NewMessage("perro4", "hola")
	tweetManager.SendDirectMessage(msg1, "perro1")
	tweetManager.SendDirectMessage(msg2, "perro1")
	tweetManager.SendDirectMessage(msg3, "perro1")

	allUnreadMessages := tweetManager.GetUnreadDirectMessages("perro1")

	if len(allUnreadMessages) != 3 {
		t.Error("Unexpected number of unread direct messages")
		return
	}

	tweetManager.ReadDirectMessage(msg1)
	tweetManager.ReadDirectMessage(msg2)

	allUnreadMessages = tweetManager.GetUnreadDirectMessages("perro1")

	if len(allUnreadMessages) != 1 {
		t.Error("Unexpected number of unread direct messages AFTER READ")
		return
	}
}

func TestErrorWhenSendingMessage(t *testing.T) {
	tweetManager = service.NewTweetManager()
	if tweetManager.GetTweet() != nil {
		t.Errorf("tweeterManager not beign cleaned")
	}

	msg1 := domain.NewMessage("gato", "hola")
	err := tweetManager.SendDirectMessage(msg1, "")
	if err != nil && err.Error() != "destinatary is required" {
		t.Errorf("'destinatary required' error expected")
	}
	msg1.From = ""
	err = tweetManager.SendDirectMessage(msg1, "alguien")
	if err != nil && err.Error() != "sender is required" {
		t.Errorf("'sender is required' error expected")
	}
	return
}

func TestRetweetOk(t *testing.T) {
	tweetManager = service.NewTweetManager()
	tweet := domain.NewTextTweet("perro", "esto es re loco #dogchow #eukanuba #fritolin")
	tweetManager.Retweet(tweet, "gato")
	if tweetManager.GetTweet().GetUser() != tweet.GetUser() ||
		tweetManager.GetTweet().GetText() != tweet.GetText() ||
		tweetManager.GetTweet().GetDate() == tweet.GetDate() {
		t.Error("bad retweet generated")
	}
	return
}

func TestRetweetNilTweet(t *testing.T) {
	tweetManager = service.NewTweetManager()
	var tweet domain.Tweet = nil
	err := tweetManager.Retweet(tweet, "gato")
	if err != nil && err.Error() != "cannot retweet nil tweet" {
		t.Error("expected error 'cannot retweet nil tweet'")
	}
	return
}

func TestCanGetPrintableTweet(t *testing.T) {
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	text := tweet.PrintableTweet()

	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is: %s \nBut I got this: %s", expectedText, text)
	}
}

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image", "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := fmt.Sprint(tweet)

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
