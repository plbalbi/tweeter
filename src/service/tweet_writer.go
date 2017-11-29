package service

import (
	"os"

	"github.com/tweeter/src/domain"
)

type TweetWriter interface {
	/*
		WriteTweet
	*/
	WriteTweet(tweet domain.Tweet, singleTweetQuit *chan bool)
}

type MemoryTweetWriter struct {
	Tweets []domain.Tweet
}

type FileTweetWriter struct {
	file *os.File
}

type ChannelTweetWriter struct {
	tweetWriter TweetWriter
}

func (tr *ChannelTweetWriter) WriteTweet(ts *chan domain.Tweet, quit *chan bool) {
	tweet, open := <-*ts
	for open {
		go tr.tweetWriter.WriteTweet(tweet, quit)
		<-*quit
		tweet, open = <-*ts
	}
	// Desbloquea quit, deberia? No tiene que ser asincrónico con respecto
	// al llamador de 'WriteTweet' del 'ChannelTweetWriter'
}

func (tr *FileTweetWriter) WriteTweet(tweet domain.Tweet, singleTweetQuit *chan bool) {
	// ts is a single tweet channel
	if tr.file != nil {
		byteSlice := []byte(tweet.PrintableTweet())
		tr.file.Write(byteSlice)
	}
	*singleTweetQuit <- true
}

func (tr *MemoryTweetWriter) WriteTweet(tweet domain.Tweet, singleTweetQuit *chan bool) {
	tr.Tweets = append(tr.Tweets, tweet)
	*singleTweetQuit <- true
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	memoryTweetWriter := new(MemoryTweetWriter)
	memoryTweetWriter.Tweets = make([]domain.Tweet, 0)
	return memoryTweetWriter
}

func NewFileTweetWriter() *FileTweetWriter {
	fileTweetWriter := new(FileTweetWriter)
	file, _ := os.OpenFile(
		"tweets.save",
		os.O_RDWR|os.O_APPEND,
		0666,
	)
	fileTweetWriter.file = file
	return fileTweetWriter
}

func NewChannelTweetWriter(tr TweetWriter) *ChannelTweetWriter {
	channelTweetWriter := new(ChannelTweetWriter)
	channelTweetWriter.tweetWriter = tr
	return channelTweetWriter
}
