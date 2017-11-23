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
			_, err := service.PublishTweet(domain.NewTweet(user, tweet))
			if err != nil {
				if err.Error() == "user is required" {
					fmt.Println("user is required to tweet")
					return
				}
				if err.Error() == "text is required" {
					fmt.Println("text is required to tweet")
					return
				}
				if err.Error() == "text longer that 140 characters" {
					fmt.Println("text too long to tweet")
					return
				}
			}
			c.Print("Tweet sent\n")
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

	shell.Run()

}
