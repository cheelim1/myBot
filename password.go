package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-password/password"
	"github.com/shomali11/slacker"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariablePW(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func PasswordGenerator(){
	botToken := goDotEnvVariablePW("SLACK_BOT_TOKEN")
	_ = botToken
	// fmt.Println("SLACK BOT TOKEN", botToken)
	// fmt.Printf("%s uses %s\n", os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	
	password, error := password.Generate(16, 8, 8, false, false)
	if error != nil {
	  log.Fatal(error)
	}
	fmt.Println("This is the random generated password:", password)

	//init bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	//slack bot event of message randompassword, response generated password
	bot.Command("randompassword", &slacker.CommandDefinition{
		Handler: func(botXtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply(password)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}