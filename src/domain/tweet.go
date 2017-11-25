package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	GetUser() string
	GetText() string
	GetId() int
	IsRetweet() bool
	RetweetedByWhom() string
	SetId(id int) int
	PrintableTweet() string
}

type TextTweet struct {
	User        string
	Text        string
	Date        *time.Time
	Id          int
	IsRetweet   bool
	RetweetedBy string
}

type ImageTweet struct {
	TextTweet
	Image string
}

type QuoteTweet struct {
	TextTweet
	Quote Tweet
}

func NewTextTweet(user, text string) *TextTweet {
	actualTime := time.Now()
	tweet := TextTweet{
		User:      user,
		Text:      text,
		Date:      &actualTime,
		IsRetweet: false,
	}
	return &tweet
}

func NewImageTweet(user, text, URL string) *ImageTweet {
	actualTime := time.Now()
	tweet := ImageTweet{
		TextTweet: TextTweet{
			User:      user,
			Text:      text,
			Date:      &actualTime,
			IsRetweet: false,
		},
		Image: URL,
	}
	tweet.Image = URL
	return &tweet
}

func NewQuoteTweet(user, text string, quote Tweet) *QuoteTweet {
	actualTime := time.Now()
	tweet := QuoteTweet{
		TextTweet: TextTweet{
			User:      user,
			Text:      text,
			Date:      &actualTime,
			IsRetweet: false,
		},
		Quote: quote,
	}
	return &tweet
}

func NewRetweet(originalTweet *Tweet, retweeter string) *Tweet {
	actualTime := time.Now()
	tweet := Tweet{User: originalTweet.User, Text: originalTweet.Text, Date: &actualTime, IsRetweet: true, RetweetedBy: retweeter}
	return &tweet
}

func (tweet *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}

func (tweet *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", tweet.User, tweet.Text, tweet.Image)
}

func (tweet *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \"@%s: %s\"", tweet.User, tweet.Text,
		tweet.Quote.GetUser(), tweet.Quote.GetText())
}

// Setter and getters needed

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *ImageTweet) GetUser() string {
	return tweet.User
}

func (tweet *QuoteTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *ImageTweet) GetText() string {
	return tweet.Text
}

func (tweet *QuoteTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *ImageTweet) GetId() int {
	return tweet.Id
}

func (tweet *QuoteTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id

}
func (tweet *ImageTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *QuoteTweet) SetId(id int) {
	tweet.Id = id
}
