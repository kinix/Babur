package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	GetResponse(msg string) string
}

func (babur *Bot) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Log messages
	if len(m.Message.Content) > 120 {
		fmt.Printf("MSG: %s ...\n", m.Message.Content[:120])
	} else {
		fmt.Printf("MSG: %s\n", m.Message.Content)
	}

	// try to roll dice
	if response := babur.dice.GetResponse(m.Message.Content); response != "" {
		sendMessage(s, m, response)
		return
	}

	// llm conversation
	if strings.Contains(m.Message.Content, babur.botId) {
		answer, err := babur.llm.Question(m.Message.Content)

		if err != nil {
			sendMessage(s, m, err.Error())
			return
		}

		if answer == "" {
			return
		}

		sendMessage(s, m, answer)
		return
	}
}

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	// Mention the user in the first line
	msg := fmt.Sprintf("<@%s> \n%s", m.Message.Author.ID, content)

	// Log reponses
	fmt.Printf("RESPONSE: %s\n", msg)

	if _, err := s.ChannelMessageSend(m.ChannelID, msg); err != nil {
		fmt.Println("ERROR: Send message: ", err)
		return
	}
}
