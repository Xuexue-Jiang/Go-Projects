package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){

	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3884288083333-00000-2b3FZLxuhynkiunbb6CJbJfW")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03RKPTTTT-0000-9625e676f6222222225ad15ce25fe2d9a2208389b3bbbbbbbbbbbb")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"My yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("Error converting year!")
			}
			age := 2022 - yob
			ageToPrint := fmt.Sprintf("Your age is %d", age)
			response.Reply(ageToPrint)
		},
		
		HideHelp: false,
	})

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {

		log.Fatal(err)
	}
}

