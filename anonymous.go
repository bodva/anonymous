package main

import (
	"flag"
	"fmt"
	"github.com/bodva/anonymous/httpjson"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var token string

type User struct {
	Status bool   `json:"ok"`
	Error  string `json:"error"`
	Url    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamId string `json:"team_id"`
	UserId string `json:"user_id"`
}

type Message struct {
	Type  string `json:"type"`
	User  string `json:"user"`
	Text  string `json:"text"`
	BotId string `json:"bot_id"`
}

type MessageRespose struct {
	Status  bool    `json:"ok"`
	Error   string  `json:"error"`
	Channel string  `json:"channel"`
	Message Message `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token = os.Getenv("TOKEN")

	fmt.Println("Oh, I remember Member Berries!")
	var method string
	fmt.Println(token)

	action := flag.String("action", "message", "as string ex. \"message\"")
	channel := flag.String("channel", "", "as string")
	asUser := flag.String("as_user", "1", "as int ( 0 from user 1 from @slackbot")
	message := flag.String("message", "", "as string")
	flag.Parse()

	r := httpjson.Response{}
	method = "api.test"
	err = httpjson.GetJson(method, &r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Slack API status response:", r.Status)

	if r.Status {
		u := User{}
		method = "auth.test"

		params := []httpjson.Param{
			{Key: "token", Value: token},
		}
		// fmt.Println(params)
		httpjson.PostJson(method, params, &u)
		fmt.Println("bot info:", u)
		switch *action {
		case "message":
			fmt.Println("Channel:", *channel, "As User:", *asUser, "Message:", *message)
			m := MessageRespose{}
			method = "chat.postMessage"
			params = []httpjson.Param{
				{Key: "token", Value: token},
				{Key: "channel", Value: *channel},
				{Key: "as_user", Value: *asUser},
				{Key: "text", Value: *message},
			}
			httpjson.PostJson(method, params, &m)
			fmt.Println("message response:", m)
		}

	}
}
