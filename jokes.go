package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

type Joke struct {
	Setup string `json:"setup"`
	Punchline string `json:"punchline"`
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func JokeOfTheDay (){
	response, error := http.Get("https://official-joke-api.appspot.com/random_joke")
	if error != nil{
		fmt.Println(error.Error())
		os.Exit(1)
	}
	responseData, error := ioutil.ReadAll(response.Body)
    if error != nil {
        log.Fatal(error)
    }

	botToken := goDotEnvVariable("SLACK_BOT_TOKEN")
	_ = botToken

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

    var responseObject Joke
    json.Unmarshal(responseData, &responseObject)

	// fmt.Println("Joke of the Day is: ")
    // fmt.Println("Question: ",responseObject.Setup)
	// // Calling Sleep method
    // time.Sleep(3 * time.Second)
	// fmt.Println("Response: ",responseObject.Punchline)
	
	// slack bot event of message randompassword, response generated password
	bot.Command("joke", &slacker.CommandDefinition{
		Handler: func(botXtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			SetupQ := responseObject.Setup
			Reply:= responseObject.Punchline
			response.Reply(SetupQ)
			// Calling Sleep method
			time.Sleep(3 * time.Second)
			response.Reply(Reply)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}