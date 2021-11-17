// This package can work wrong. It is not well tested. I just wrote that for fun
package chat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func initAnswerConfig() error {
	cfgFile, err := os.Open("config/chat.json")
	if err != nil {
		return fmt.Errorf("ERROR: Read chat.cfg: %s", err)
	}

	defer cfgFile.Close()

	// Read and parse json file
	bytes, _ := ioutil.ReadAll(cfgFile)
	err = json.Unmarshal(bytes, &answerList)
	if err != nil {
		return fmt.Errorf("ERROR: Parse chat.cfg: %s", err)
	}

	return nil
}

func initRegexConfig() error {
	cfgFile, err := os.Open("config/chat_regex.json")
	if err != nil {
		return fmt.Errorf("ERROR: Read chat_regex.cfg: %s", err)
	}

	defer cfgFile.Close()

	var regexConfig map[string]string

	// Read and parse json file
	bytes, _ := ioutil.ReadAll(cfgFile)
	err = json.Unmarshal(bytes, &regexConfig)
	if err != nil {
		return fmt.Errorf("ERROR: Parse chat_regex.cfg: %s", err)
	}

	regexList = map[string]*regexp.Regexp{}
	for key, value := range regexConfig {
		regexList[key] = regexp.MustCompile(value)
	}

	return nil
}

func validateConfig() error {
	for key := range regexList {
		if _, found := answerList[key]; !found {
			return fmt.Errorf("ERROR: Answer key is not found: %s", key)
		}
	}

	if _, found := answerList["_"]; !found {
		return errors.New("ERROR: Answer key is not found: _")
	}

	if _, found := answerList["_?"]; !found {
		return errors.New("ERROR: Answer key is not found: _?")
	}

	return nil
}
