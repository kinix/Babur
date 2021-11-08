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

	// Does the message have any dice text?
	if dice, side, addition := checkMessageForDice(m.Message.Content); dice > 0 {
		content := rollDice(dice, side, addition)
		sendMessage(s, m, content)
	}

	if measurements := checkMessageForConverting(m.Message.Content); len(measurements) > 0 {
		content := convertUnits(measurements)
		sendMessage(s, m, content)
	}
}

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	// Mention the user in the first line
	msg := fmt.Sprintf("<@%s> \n%s", m.Message.Author.ID, content)

	if _, err := s.ChannelMessageSend(m.ChannelID, msg); err != nil {
		fmt.Println("ERROR: Send message: ", err)
		return
	}
}
