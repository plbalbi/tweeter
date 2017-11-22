package service

var tweet string

func PublishTweet(t string) {
	tweet = t
}

func CleanTweet() {
	tweet = ""
}

func GetTweet() string {
	return tweet
}
