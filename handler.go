package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if _, err := s.ChannelMessageSend(m.ChannelID, "hello world"); err != nil {
		fmt.Println("ERROR: Send message: ", err)
		return
	}
}
