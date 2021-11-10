package main

import (
	"babur/chat"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If some one mentioned Babür
	if strings.Contains(m.Message.Content, s.State.User.ID) {
		// Use chat (maybe we can add more language in future. maybe...)
		msg := strings.ReplaceAll(m.Message.Content, "<@!"+s.State.User.ID+">", "")
		if content := chat.ChatHandler(m.Author.ID, msg); content != "" {
			sendMessage(s, m, content)
		}
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
