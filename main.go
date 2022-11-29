package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

const (
	SLACK_APP_TOKEN = "xapp-1-A04CZSC6R8C-4433240013302-ff042d01ced97a7d43bc389f29a9d684b51337b962c5fb2f95bae56479dad5a2"
	SLACK_BOT_TOKEN = "xoxb-4014835228961-4439506777730-5JzA0r56Jxj5V54jojvyZz7I"
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
