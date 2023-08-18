package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/kinix/babur/dice"
	"github.com/kinix/babur/llmchat"
)

type Bot struct {
	llm         *llmchat.LLMClient
	dice        *dice.DiceHandler
	Temperature int
	botId       string
}

var babur *Bot

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("ERROR: Load .env file: ", err)
		return
	}

	babur = &Bot{}

	// Seed random to avoid same results
	rand.Seed(time.Now().UnixNano())

	session, err := discordgo.New("Bot " + os.Getenv("BABUR_TOKEN"))
	if err != nil {
		fmt.Println("ERROR: Create bot: ", err)
		return
	}

	defer session.Close()

	// Receive events for channel messages
	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		fmt.Println("ERROR: Open socket: ", err)
		return
	}

	temp, err := strconv.Atoi(os.Getenv("OPENAI_API_TEMPERATURE"))
	if err != nil {
		fmt.Println("ERROR: temperature: ", err)
		temp = 5
	}

	babur.botId = session.State.User.ID
	babur.llm = llmchat.NewClient(os.Getenv("OPENAI_VERSION"), temp)

	session.AddHandler(babur.MessageHandler)

	babur.dice = &dice.DiceHandler{}
	babur.dice.InitRegex()

	fmt.Println("Babur is ready.")

	// Wait for ctrl+c or termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
