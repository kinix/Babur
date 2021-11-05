package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token string

func init() {
	token = os.Getenv("TOKEN")
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
