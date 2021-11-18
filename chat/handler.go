// This package can work wrong. It is not well tested. I just wrote that for fun
package chat

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type ChatHandler struct {
	botMention string

	regexList  map[string]*regexp.Regexp
	answerList map[string][]answer
	weightList map[string]int

	googleToken string
	googleCx    string
}

type answer struct {
	Weight   int
	Response string
}

func NewChatHandler(botId, answerConfigFile, regexConfigFile string) (*ChatHandler, error) {
	c := &ChatHandler{}

	if err := c.initAnswerConfig(answerConfigFile); err != nil {
		return nil, err
	}

	if err := c.initRegexConfig(regexConfigFile); err != nil {
		return nil, err
	}

	if err := c.validateConfig(); err != nil {
		return nil, err
	}

	c.weightList = map[string]int{}
	for key, value := range c.answerList {
		c.weightList[key] = 0

		for _, singleAnswer := range value {
			c.weightList[key] += singleAnswer.Weight
		}
	}

	// Read env values for google image search
	c.googleToken = os.Getenv("GOOGLE_TOKEN")
	c.googleCx = os.Getenv("GOOGLE_CX")

	c.botMention = "<@!" + botId + ">"

	fmt.Println("Chat is ready.")
	return c, nil
}

// Answer some question in Turkish
func (c *ChatHandler) GetResponse(msg string) string {
	// Does the message contain mention to Babur
	if strings.Contains(msg, c.botMention) {
		// Remove mention part
		msg = strings.ReplaceAll(msg, c.botMention, "")
	} else {
		return ""
	}

	var found []byte

	// Check the message is match with any of question regex
	for questionType, regex := range c.regexList {
		if found = regex.Find([]byte(msg)); found != nil {
			// Choose an answer if matched
			return c.chooseAnswer(questionType, msg)
		}
	}

	// There is no match and it is a question
	if strings.Contains(msg, "?") {
		return c.chooseAnswer("_?", msg)
	}

	// There is no match and it is not a question
	return c.chooseAnswer("_", msg)
}

func (c *ChatHandler) chooseAnswer(questionType string, msg string) string {
	luckyNumber := rand.Intn(c.weightList[questionType])

	cursor := 0
	for _, singleAnswer := range c.answerList[questionType] {
		cursor += singleAnswer.Weight

		if cursor > luckyNumber {
			// If response begins with %%, Babür can do some calculations to respond
			if len(singleAnswer.Response) > 2 && singleAnswer.Response[0:2] == "%%" {
				return c.specialResponse(singleAnswer.Response[2:], msg)
			} else {
				return singleAnswer.Response
			}
		}
	}

	return ""
}
