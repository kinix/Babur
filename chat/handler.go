// This package can work wrong. It is not well tested. I just wrote that for fun
package chat

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type answer struct {
	Weight   int
	Response string
}

var regexList map[string]*regexp.Regexp
var answerList map[string][]answer
var weightList map[string]int

func init() {
	if err := initAnswerConfig(); err != nil {
		os.Exit(1)
	}

	if err := initRegexConfig(); err != nil {
		os.Exit(1)
	}

	weightList = map[string]int{}
	for key, value := range answerList {
		weightList[key] = 0

		for _, singleAnswer := range value {
			weightList[key] += singleAnswer.Weight
		}
	}

	// Read env values for google image search
	googleToken = os.Getenv("GOOGLE_TOKEN")
	googleCx = os.Getenv("GOOGLE_CX")

	fmt.Println("Chat is ready.")
}

// Answer some question in Turkish
func ChatHandler(owner string, msg string) string {
	var found []byte

	// Check the message is match with any of question regex
	for questionType, regex := range regexList {
		if found = regex.Find([]byte(msg)); found != nil {
			// Choose an answer if matched
			return chooseAnswer(questionType, msg)
		}
	}

	// There is no match and it is a question
	if strings.Contains(msg, "?") {
		return chooseAnswer("_?", msg)
	}

	// There is no match and it is not a question
	return chooseAnswer("_", msg)
}

func chooseAnswer(questionType string, msg string) string {
	luckyNumber := rand.Intn(weightList[questionType])

	cursor := 0
	for _, singleAnswer := range answerList[questionType] {
		cursor += singleAnswer.Weight

		if cursor > luckyNumber {
			// If response begins with %%, Babür can do some calculations to respond
			if len(singleAnswer.Response) > 2 && singleAnswer.Response[0:2] == "%%" {
				return specialResponse(singleAnswer.Response[2:], msg)
			} else {
				return singleAnswer.Response
			}
		}
	}

	return ""
}
