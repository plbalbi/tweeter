package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func getPreetyHour(date *time.Time) string {
	return fmt.Sprintf("%d:%d", date.Hour(), date.Minute())
}

func getPreetyDate(date *time.Time) string {
	return fmt.Sprintf("%s %d, %d", date.Month().String(), date.Day(), date.Year())
}

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")
	tweetManager := service.NewTweetManager()

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a Tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write your user: ")
			user := c.ReadLine()
			c.Print("Write you tweet: ")
			tweet := c.ReadLine()
			publishedTweetId, err := tweetManager.PublishTweet(domain.NewTweet(user, tweet))
			if err != nil {
				if err.Error() == "user is required" {
					c.Print("user is required to tweet")
					return
				}
				if err.Error() == "text is required" {
					c.Print("text is required to tweet")
					return
				}
				if err.Error() == "text longer that 140 characters" {
					c.Print("text too long to tweet")
					return
				}
			}
			c.Printf("Tweet sent with id %d\n", publishedTweetId)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweet := tweetManager.GetTweet()
			c.Println(tweet)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Show all available tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweets := tweetManager.GetTweets()
			for index, tweet := range tweets {
				c.Println("Tweet #" + strconv.Itoa(index) + " | User: " + tweet.User + "\nText: " + tweet.Text)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetById",
		Help: "Access a tweet by id",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Which tweet should i bring? [id] : ")
			stringId := c.ReadLine()
			id, _ := strconv.Atoi(stringId)
			fmt.Println(tweetManager.GetTweetById(id))
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanTweets",
		Help: "Clean all tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweetManager.CleanTweets()
			c.Print("All tweets cleaned")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countUserTweets",
		Help: "Count all user tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("User to count tweets from: ")
			userToCount := c.ReadLine()
			count := tweetManager.CountTweetsByUser(userToCount)
			c.Printf("%s has tweeted %d times\n", userToCount, count)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getUserTweets",
		Help: "Get all user tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("User to find tweets from: ")
			userToFindTweetsFrom := c.ReadLine()
			userTweets := tweetManager.GetTweetsByUser(userToFindTweetsFrom)
			c.Print(userTweets)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "follow",
		Help: "follow a user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("you are: ")
			whoAmI := c.ReadLine()
			c.Print("who you wanna follow: ")
			wannaFollow := c.ReadLine()
			err := tweetManager.Follow(whoAmI, wannaFollow)
			if err != nil && err.Error() == "user to follow not found" {
				c.Printf("user %s has not tweeted yet...", wannaFollow)
			}
			c.Print("Followed\n")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTimeline",
		Help: "get timeline a user sees",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("whose timeline: ")
			user := c.ReadLine()
			timeline := tweetManager.Timeline(user)
			c.Printf("\tTimeline from %s\n", user)
			c.Print("---------------------------------\n")
			for _, tweet := range timeline {
				c.Printf("%s tweeted at %s, on %s:\n\t%s\n", tweet.User, getPreetyHour(tweet.Date),
					getPreetyDate(tweet.Date), tweet.Text)
			}
			return
		},
	})

	shell.Run()
}
