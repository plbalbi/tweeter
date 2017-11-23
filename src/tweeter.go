package main

import (
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a Tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Write you tweet: ")
			tweet := c.ReadLine()
			c.Print("Write your user: ")
			user := c.ReadLine()
			publishedTweetId, err := service.PublishTweet(domain.NewTweet(user, tweet))
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
			tweet := service.GetTweet()
			c.Println(tweet)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Show all available tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweets := service.GetTweets()
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
			fmt.Println(service.GetTweetById(id))
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanTweets",
		Help: "Clean all tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			service.CleanTweets()
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
			count := service.CountTweetsByUser(userToCount)
			c.Printf("%s has tweeted %d times\n", userToCount, count)
			return
		},
	})

	shell.Run()
}
