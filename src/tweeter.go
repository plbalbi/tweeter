package main

import (
	"fmt"

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
			err := service.PublishTweet(domain.NewTweet(user, tweet))
			if err != nil {
				if err.Error() == "user is required" {
					fmt.Println("user is required to tweet")
					return
				}
				if err.Error() == "text is required" {
					fmt.Println("text is required to tweet")
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

	shell.Run()

}
