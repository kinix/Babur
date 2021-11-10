// This package can work wrong. It is not well tested. I just wrote that for fun
package chat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func initAnswerConfig() error {
	cfgFile, err := os.Open("config/chat.json")
	if err != nil {
		fmt.Println("ERROR: Read chat.cfg: ", err)
		return err
	}

	defer cfgFile.Close()

	// Read and parse json file
	bytes, _ := ioutil.ReadAll(cfgFile)
	err = json.Unmarshal(bytes, &answerList)
	if err != nil {
		fmt.Println("ERROR: Parse chat.cfg: ", err)
		return err
	}

	return nil
}

func initRegexConfig() error {
	cfgFile, err := os.Open("config/chat_regex.json")
	if err != nil {
		fmt.Println("ERROR: Read chat_regex.cfg: ", err)
		return err
	}

	defer cfgFile.Close()

	var regexConfig map[string]string

	// Read and parse json file
	bytes, _ := ioutil.ReadAll(cfgFile)
	err = json.Unmarshal(bytes, &regexConfig)
	if err != nil {
		fmt.Println("ERROR: Parse chat_regex.cfg: ", err)
		return err
	}

	regexList = map[string]*regexp.Regexp{}
	for key, value := range regexConfig {
		regexList[key] = regexp.MustCompile(value)
	}

	return nil
}
