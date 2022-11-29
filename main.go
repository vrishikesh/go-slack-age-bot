package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

// slack access token - xoxe.xoxp-1-Mi0yLTQwMTQ4MzUyMjg5NjEtNDAwMjEwMzY5MzE1NS00NDMyODE5MzU2NDM4LTQ0Mzk0MTc1OTk4MjYtMDBlY2EzYzIxNmIzZjY1NjFlYzFhY2I3OWJkNTc3ZDcyNTk2MzI4MGRjN2EzYTQ4MjNhNDFmYTA4OWNiZWNmNA
// slack refresh token - xoxe-1-My0xLTQwMTQ4MzUyMjg5NjEtNDQzMjgxOTM1NjQzOC00NDYzMTE5NjE0NjcyLWFjNjhlYmY0NTc3ZjEzNDI4ZDIzOTFkMjZhNzM0OWZhNjUyMmRhMThhMzU2ZDA5NTE1MzRhYTA3N2YwOGQ4MGY
// slack socket token - xapp-1-A04CZSC6R8C-4452091575953-c6e1eb8de56ce35dc867ea9b0074e3629febb6cc65e68f66ede3e00ba66c69e5
// . slack bot token - xoxb-4014835228961-4439506777730-nPEtDx1jKN7R3J8xekskCaBZ

const (
	SLACK_APP_TOKEN = "xapp-1-A04CZSC6R8C-4452091575953-c6e1eb8de56ce35dc867ea9b0074e3629febb6cc65e68f66ede3e00ba66c69e5"
	SLACK_BOT_TOKEN = "xoxb-4014835228961-4439506777730-nPEtDx1jKN7R3J8xekskCaBZ"
)

func printCommandEvents(analyticsChan <-chan *slacker.CommandEvent) {
	for event := range analyticsChan {
		log.Println("Command Events")
		log.Println(event.Timestamp)
		log.Println(event.Command)
		log.Println(event.Parameters)
		log.Println(event.Event)
	}
}

func main() {
	bot := slacker.NewClient(SLACK_BOT_TOKEN, SLACK_APP_TOKEN)

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 1990"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Println(err)
				response.Reply(fmt.Sprintf("could not parse year: %s", year))
				return
			}

			currentYear, _, _ := time.Now().Date()
			age := currentYear - yob
			response.Reply(fmt.Sprintf("age is %d", age))
		},
	})

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
