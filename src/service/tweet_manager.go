package service

import (
	"fmt"

	"github.com/tweeter/src/domain"
)

type followsMap map[string][]string
type tweetsMap map[string][]*domain.Tweet

type TweetManager struct {
	tweets         tweetsMap
	follows        followsMap
	lastFreeId     int
	lastAddedTweet *domain.Tweet
}

func NewTweetManager() *TweetManager {
	var tweetManager TweetManager = TweetManager{
		tweets:         make(tweetsMap),
		follows:        make(followsMap),
		lastFreeId:     0,
		lastAddedTweet: nil,
	}
	// initialize empty slice
	return &tweetManager
}

func (tweetManager *TweetManager) PublishTweet(t *domain.Tweet) (int, error) {
	if t.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if t.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(t.Text) > 140 {
		return 0, fmt.Errorf("text longer that 140 characters")
	}
	t.Id = tweetManager.lastFreeId
	tweetManager.lastFreeId++
	tweetManager.lastAddedTweet = t
	tweetManager.tweets[t.User] = append(tweetManager.tweets[t.User], t)
	return t.Id, nil
}

func (tweetManager *TweetManager) CleanTweets() {
	tweetManager.tweets = make(tweetsMap)
	tweetManager.follows = make(followsMap)
	tweetManager.lastAddedTweet = nil
	tweetManager.lastFreeId = 0
}

// solucion medio mala
// cambiar a iterador o una lista media ensamblada
func (tweetManager *TweetManager) GetTweets() [](*domain.Tweet) {
	allTweets := make([]*domain.Tweet, 0)
	for _, tweets := range tweetManager.tweets {
		allTweets = append(allTweets, tweets...)
	}
	return allTweets
}

func (tweetManager *TweetManager) NoTweets() bool {
	return len(tweetManager.tweets) == 0
}

func (tweetManager *TweetManager) GetTweetById(id int) *domain.Tweet {
	allTweets := tweetManager.GetTweets()
	for _, tweet := range allTweets {
		if tweet.Id == id {
			return tweet
		}
	}
	return nil
}

func (tweetManager *TweetManager) GetTweet() *domain.Tweet {
	return tweetManager.lastAddedTweet
}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {
	return len(tweetManager.tweets[user])
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	return tweetManager.tweets[user]
}

func (tweetManager *TweetManager) Follow(user string, toFollow string) error {
	_, toFollowDefined := tweetManager.tweets[toFollow]
	if !toFollowDefined {
		return fmt.Errorf("user to follow not found")
	}

	tweetManager.follows[user] = append(tweetManager.follows[user], toFollow)
	return nil
}

func (tweetManager *TweetManager) Timeline(user string) []*domain.Tweet {
	userFollows := tweetManager.follows[user]
	timeline := make([]*domain.Tweet, 0)
	for _, follow := range userFollows {
		timeline = append(timeline, tweetManager.GetTweetsByUser(follow)...)
	}
	return timeline
}
