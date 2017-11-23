package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

var tweets map[string][]*domain.Tweet
var follows map[string][]string
var lastFreeId int
var lastAddedTweet *domain.Tweet

func InitializeService() {
	// initialize empty slice
	lastFreeId = 0
	tweets = make(map[string][]*domain.Tweet)
	follows = make(map[string][]string)
	lastAddedTweet = nil
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
	t.Id = lastFreeId
	lastFreeId++
	lastAddedTweet = t
	tweets[t.User] = append(tweets[t.User], t)
	return t.Id, nil
}

func CleanTweets() {
	tweets = nil
	InitializeService()
}

// solucion medio mala
// cambiar a iterador o una lista media ensamblada
func GetTweets() [](*domain.Tweet) {
	allTweets := make([]*domain.Tweet, 0)
	for _, tweets := range tweets {
		allTweets = append(allTweets, tweets...)
	}
	return allTweets
}

func NoTweets() bool {
	return len(tweets) == 0
}

func GetTweetById(id int) *domain.Tweet {
	allTweets := GetTweets()
	for _, tweet := range allTweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func GetTweet() *domain.Tweet {
	return lastAddedTweet
}

func CountTweetsByUser(user string) int {
	return len(tweets[user])
}

func GetTweetsByUser(user string) []*domain.Tweet {
	return tweets[user]
}

func Follow(user string, toFollow string) error {
	_, toFollowDefined := tweets[toFollow]
	if !toFollowDefined {
		return fmt.Errorf("user %d has not tweeted yet...", toFollow)
	}

	follows[user] = append(follows[user], toFollow)
	return nil
}

func Timeline(user string) []*domain.Tweet {
	userFollows := follows[user]
	timeline := make([]*domain.Tweet, 0)
	for _, follow := range userFollows {
		timeline = append(timeline, GetTweetsByUser(follow)...)
	}
	return timeline
}
