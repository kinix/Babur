package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var token string

func init() {
	token = os.Getenv("TOKEN")

	// Read dice configs from dice.json
	if err := initDiceConfig(); err != nil {
		// Exit if the config is broken
		os.Exit(1)
	}

	// Read converter configs from units.json
	if err := initUnitConfig(); err != nil {
		// Exit if the config is broken
		os.Exit(1)
	}

	initDiceRegex()
	initUnitRegex()

	// Seed random to avoid same results
	rand.Seed(time.Now().UnixNano())
}

func main() {
	session, err := discordgo.New("Bot " + token)
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

	// Event handler
	session.AddHandler(messageHandler)

	fmt.Println("Babur is ready.")

	// Wait for ctrl+c or termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
