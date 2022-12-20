package main

import(
"context"
"encoding/json"
"fmt"
"log"
"os"
"github.com/joho/godotenv"
"github.com/krognol/go-wolfram"
"github.com/shomali11/slacker"
witai "github.com/wit-ai/wit-go/v2"
"github.com/tidwall/gjson"
)

var wolframClient *wolfram.Client

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	godotenv.Load(".env")

	bot := slacker.NewClient(os.Getenv("SLACK_OAUTH_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	wolframClient := &wolfram.Client{AppID: os.Getenv("WOLFRAM_API_ID")}
	go printCommandEvents(bot.CommandEvents())

	bot.Command("question - <message>", &slacker.CommandDefinition{
		Description: "Send a question to wolfram",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")
			msg, _  := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			data, _ := json.MarshalIndent(msg, "", "    ")
			clean := string(data[:])
			value := gjson.Get(clean, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			fmt.Println(value)
			answer := value.String()
			response.Reply("Received!")
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				fmt.Println("There is an error")
			}
			response.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}


}