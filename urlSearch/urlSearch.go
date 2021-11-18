package urlSearch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type UrlSearchHandler struct {
	prefix      string
	searchUrl   string
	searchRegex *regexp.Regexp
}

func NewUrlSearchHandler(prefix string, searchUrl string, searchRegex string) (*UrlSearchHandler, error) {
	u := &UrlSearchHandler{}

	u.prefix = prefix
	u.searchUrl = searchUrl
	u.searchRegex = regexp.MustCompile(searchRegex)

	return u, nil
}

func (u *UrlSearchHandler) GetResponse(msg string) string {
	// Does the message start with prefix
	if len(msg) > len(u.prefix) && msg[0:len(u.prefix)] == u.prefix {
		return u.searchDnd(msg[5:])
	}

	return ""
}

func (u *UrlSearchHandler) searchDnd(msg string) string {
	// Prepare the text to search
	link := fmt.Sprintf(u.searchUrl, url.QueryEscape(msg))

	// Search
	resp, err := http.Get(link)
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
	result := u.searchRegex.FindSubmatch(bodyBytes)
	if len(result) == 0 {
		return ""
	}

	return strings.TrimSpace(string(result[1]))
}
