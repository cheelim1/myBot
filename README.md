# myBot

myBot is a working Slack bot written in Go. Fork it & use it to build your own cool Slack bot 🤖!

## Some components myBot comes with
- random password generator
- calling external API (random Joke)

### To Get started

1. Ensure you have a slack account & Configure your slack account settings & permissions [https://api.slack.com/]
2. Fork this repository
3. Create your .env file (Add your BOT & APP TOKEN)

Example:
```
# Slack OAuth token
SLACK_BOT_TOKEN=xoxb-...
# Slack Socket Token
SLACK_APP_TOKEN=xapp-...
```
4. To run the application locally
```
go mod init github.com/<github_accountname>/<repository_name>
go get "github.com/shomali11/slacker"
go get "github.com/joho/godotenv"

go build
go run main.go
go run *.go (run multiple files)
```

#### Once added the bot to your Slack app, Interact with your bot
```
@myBot hello
@myBot joke
@myBot randompassword

### External Dependencies
1. Your Slack App Config & Permissions [https://api.slack.com/apps]