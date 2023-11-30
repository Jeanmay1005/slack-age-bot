package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func main(){
	// set up environment and credential tokens
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6283388132241-6267900374341-3JcN5oKkwBtVitCsA1Z7qjIz")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A067J79CSNT-6273233226084-2b03b16eacd297bb89d8f337b5731784d96b0be8e198a42884036f1a34c6933f")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	// prints out info of the event
	go printCommandEvents(bot.CommandEvents())

	// main logistic of chatbot execution
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob Calculator",
		// Examples: "my yob is 1997",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil{
				fmt.Println(err)
			}
			age := 2023 - yob
			r := fmt.Sprintf("Age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)	
	}

}

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}