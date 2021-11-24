package main

import (
	"babur/chat"
	"babur/converter"
	"babur/dice"
	"babur/urlSearch"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	token    string
	handlers []Handler
}

var babur *Bot

func init() {
	babur = &Bot{}
	babur.token = os.Getenv("BABUR_TOKEN")

	// Seed random to avoid same results
	rand.Seed(time.Now().UnixNano())
}

func main() {
	session, err := discordgo.New("Bot " + babur.token)
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

	// Add handlers
	var convertHandler, diceHandler, dndHandler, chatHandler Handler

	if convertHandler, err = converter.NewConverterHandler("config/units.json"); err != nil {
		panic(err)
	}

	if diceHandler, err = dice.NewDiceHandler("config/dice.json"); err != nil {
		panic(err)
	}

	if dndHandler, err = urlSearch.NewUrlSearchHandler("!dnd", "http://dnd5e.wikidot.com/search:site/q/%s", "<div class=\"url\">([^<]+)"); err != nil {
		panic(err)
	}

	if chatHandler, err = chat.NewChatHandler(session.State.User.ID, "config/chat.json", "config/chat_regex.json"); err != nil {
		panic(err)
	}

	babur.handlers = []Handler{convertHandler, diceHandler, dndHandler, chatHandler}

	session.AddHandler(babur.MessageHandler)

	fmt.Println("Babur is ready.")

	// Wait for ctrl+c or termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
