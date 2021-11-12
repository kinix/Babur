package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const dndSearchUrl = "http://dnd5e.wikidot.com/search:site/q/%s"

var dndSearchRegex *regexp.Regexp

func initDndRegex() {
	dndSearchRegex = regexp.MustCompile("<div class=\"url\">([^<]+)")
}

func searchDnd(msg string) string {
	// Prepare the text to search
	msg = fmt.Sprintf(dndSearchUrl, url.QueryEscape(msg))

	// Search
	resp, err := http.Get(msg)
	if err != nil {
		fmt.Println("ERROR: DND search request:", msg, err)
		return ""
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR: DND search read:", msg, err)
	}

	// Find the first result
	result := dndSearchRegex.FindSubmatch(bodyBytes)
	if len(result) == 0 {
		return ""
	}

	return strings.TrimSpace(string(result[1]))
}
