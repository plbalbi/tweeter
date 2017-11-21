package service

var tweet string

func PublishTweet(t string) {
	tweet = t
}

func GetTweet() string {
	return tweet
}
