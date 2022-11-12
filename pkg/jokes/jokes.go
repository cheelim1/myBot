package jokes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Joke struct {
	Setup string `json:"setup"`
	Punchline string `json:"punchline"`
}

func JokeOfTheDay()( SetupQ, Reply string ){
	response, error := http.Get("https://official-joke-api.appspot.com/random_joke")
	if error != nil{
		fmt.Println(error.Error())
		os.Exit(1)
	}
	responseData, error := ioutil.ReadAll(response.Body)
    if error != nil {
        log.Fatal(error)
    }

    var responseObject Joke
    json.Unmarshal(responseData, &responseObject)

	// fmt.Println("Joke of the Day is: ")
    // fmt.Println("Question: ",responseObject.Setup)
	// // Calling Sleep method
    // time.Sleep(3 * time.Second)
	// fmt.Println("Response: ",responseObject.Punchline)
	
	// slack bot event of message joke
	// bot.Command("joke", &slacker.CommandDefinition{
	// 	Handler: func(botXtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	// 		SetupQ := responseObject.Setup
	// 		Reply:= responseObject.Punchline
	// 		response.Reply(SetupQ)
	// 		// Calling Sleep method
	// 		time.Sleep(3 * time.Second)
	// 		response.Reply(Reply)
	// 	},
	// })
	SetupQ = responseObject.Setup
	Reply = responseObject.Punchline

	return
}

