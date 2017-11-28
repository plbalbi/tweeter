package service

import (
	"fmt"
	"strings"

	"github.com/tweeter/src/domain"
)

type TweeterPlugin interface {
	Publish(tweet domain.Tweet)
}

type followsMap map[string][]string
type wordCount map[string]int

// userInboxes tiene
// key: 		user / destinatary
// definition: 	*message
type userInboxes map[string][]*domain.Message
type pluginCollection []TweeterPlugin
type tweetsMap map[string][]domain.Tweet

type TweetManager struct {
	tweets         tweetsMap
	follows        followsMap
	hashtagCount   wordCount
	inboxes        userInboxes
	plugins        pluginCollection
	lastFreeId     int
	lastAddedTweet domain.Tweet
}

func NewTweetManager() *TweetManager {
	var tweetManager TweetManager = TweetManager{
		tweets:         make(tweetsMap),
		follows:        make(followsMap),
		hashtagCount:   make(wordCount),
		inboxes:        make(userInboxes),
		lastFreeId:     0,
		lastAddedTweet: nil,
	}
	// initialize empty slice
	return &tweetManager
}

func (tweetManager *TweetManager) PublishTweet(t domain.Tweet) (int, error) {
	// Error checking
	if t.GetUser() == "" {
		return 0, fmt.Errorf("user is required")
	}
	if t.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(t.GetText()) > 140 {
		return 0, fmt.Errorf("text longer that 140 characters")
	}
	// Tweet index tracking
	t.SetId(tweetManager.lastFreeId)
	tweetManager.lastFreeId++
	tweetManager.lastAddedTweet = t
	// Adding the tweet
	tweetManager.tweets[t.GetUser()] = append(tweetManager.tweets[t.GetUser()], t)
	// Updating hashtag count
	for _, word := range strings.Fields(t.GetText()) {
		if _, wordDefined := tweetManager.hashtagCount[word]; !wordDefined {
			tweetManager.hashtagCount[word] = 1
		} else {
			tweetManager.hashtagCount[word]++
		}
	}
	// Llamo al Publish() de cada plugin
	if len(tweetManager.plugins) > 0 {
		// Si hay algún plugin, llamo Publish(tweetPublicado) por cada plugin
		for _, plugin := range tweetManager.plugins {
			plugin.Publish(t)
		}
	}

	return t.GetId(), nil
}

func (tweetManager *TweetManager) CleanTweets() {
	tweetManager.tweets = make(tweetsMap)
	tweetManager.follows = make(followsMap)
	tweetManager.hashtagCount = make(wordCount)
	tweetManager.lastAddedTweet = nil
	tweetManager.lastFreeId = 0
}

// solucion medio mala
// cambiar a iterador o una lista media ensamblada
func (tweetManager *TweetManager) GetTweets() []domain.Tweet {
	allTweets := make([]domain.Tweet, 0)
	for _, tweets := range tweetManager.tweets {
		allTweets = append(allTweets, tweets...)
	}
	return allTweets
}

func (tweetManager *TweetManager) NoTweets() bool {
	return len(tweetManager.tweets) == 0
}

func (tweetManager *TweetManager) GetTweetById(id int) domain.Tweet {
	allTweets := tweetManager.GetTweets()
	for _, tweet := range allTweets {
		if tweet.GetId() == id {
			return tweet
		}
	}
	return nil
}

func (tweetManager *TweetManager) GetTweet() domain.Tweet {
	return tweetManager.lastAddedTweet
}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {
	return len(tweetManager.tweets[user])
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
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

func (tweetManager *TweetManager) Timeline(user string) []domain.Tweet {
	userFollows := tweetManager.follows[user]
	timeline := make([]domain.Tweet, 0)
	for _, follow := range userFollows {
		timeline = append(timeline, tweetManager.GetTweetsByUser(follow)...)
	}
	return timeline
}

func (tweetManager *TweetManager) GetTrendingTopics() []string {
	trendingTopics := make([]string, 2)
	trendingTopics[0] = "undefined"
	trendingTopics[1] = "undefined"
	trendingTopicsCount := make([]int, 2)
	trendingTopicsCount[0] = -1
	trendingTopicsCount[1] = -1
	// Como se que trendingTopicsCount[0] siempre sera el max, si encuentro uno mayor, lo guardo ahi
	// y shifteo el tTC[0] a tTC[1]
	for word, wordCount := range tweetManager.hashtagCount {
		// Check if it's a hashtag
		if word[0] != '#' {
			continue
		}
		// La palabra es el nuevo máximo?
		if wordCount > trendingTopicsCount[0] {
			// Muevo tTC[0] como el segundo mas frecuente
			trendingTopics[1] = trendingTopics[0]
			trendingTopicsCount[1] = trendingTopicsCount[0]
			// Actualizo el maximo actual
			trendingTopics[0] = word
			trendingTopicsCount[0] = wordCount
		}
	}
	return trendingTopics
}

func (tweetManager *TweetManager) GetAllDirectMessages(user string) []*domain.Message {
	return tweetManager.inboxes[user]
}
func (tweetManager *TweetManager) GetUnreadDirectMessages(user string) []*domain.Message {
	unreadMessages := make([]*domain.Message, 0)

	for _, message := range tweetManager.inboxes[user] {
		if !message.Read {
			unreadMessages = append(unreadMessages, message)
		}
	}
	return unreadMessages
}
func (tweetManager *TweetManager) ReadDirectMessage(message *domain.Message) error {
	if message == nil {
		err := fmt.Errorf("message is nil")
		return err
	}
	message.Read = true
	return nil
}

func (tweetManager *TweetManager) SendDirectMessage(message *domain.Message, destinatary string) error {
	// Checking message OK
	if destinatary == "" {
		err := fmt.Errorf("destinatary is required")
		return err
	}
	if message.From == "" {
		err := fmt.Errorf("sender is required")
		return err
	}
	// Sending message
	tweetManager.inboxes[destinatary] = append(tweetManager.inboxes[destinatary], message)
	return nil
}

func (tweetManager *TweetManager) Retweet(originalTweet domain.Tweet, retweeter string) error {
	// Check if tweet is nil
	if originalTweet == nil {
		return fmt.Errorf("cannot retweet nil tweet")
	}
	retweet := domain.NewRetweet(originalTweet, retweeter)
	retweetId, err := tweetManager.PublishTweet(retweet)
	if err != nil {
		return err
	}
	retweet.SetId(retweetId)
	return nil
}
