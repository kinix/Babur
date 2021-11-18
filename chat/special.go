// This package can work wrong. It is not well tested. I just wrote that for fun
package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

const imageSearchUrl = "https://www.googleapis.com/customsearch/v1?key=%s&searchType=image&cx=%s&q=%s"
const imageMaxCount = 10 // API returns 10 at max

type imageResponse struct {
	Items []imageItem
}

type imageItem struct {
	Link string
}

// If response begins with %%, Babür can do some calculations to respond
// This function get msg as parameter. It is useless for now, but it can be used in future (maybe image search?)
func (c *ChatHandler) specialResponse(specialType string, msg string) string {
	switch specialType {
	case "randomNumber":
		luckyNumber := rand.Intn(500000)

		if luckyNumber < 2000 {
			return fmt.Sprint(luckyNumber)
		} else if luckyNumber < 10000 {
			return fmt.Sprint(luckyNumber % 1000)
		} else if luckyNumber < 50000 {
			return fmt.Sprint(luckyNumber % 100)
		} else {
			return fmt.Sprint(luckyNumber % 10)
		}
	case "imageSearch":
		// If token or cx is not defined, image search will not work
		if c.googleToken == "" || c.googleCx == "" {
			return ""
		}

		// Prepare text for query
		msg = strings.Trim(msg, " .?-")
		msg = url.QueryEscape(msg)

		// Do the search request
		url := fmt.Sprintf(imageSearchUrl, c.googleToken, c.googleCx, msg)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("ERROR: Image search request: ", err)
			return ""
		}

		defer resp.Body.Close()

		var imageList imageResponse

		// Read the body
		respBytes, err2 := io.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println("ERROR: Image search read: ", err)
			return ""
		}

		// Parse the body
		json.Unmarshal(respBytes, &imageList)

		// If image count is less than expected, set new limit
		imageCount := imageMaxCount
		if imageCount > len(imageList.Items) {
			imageCount = len(imageList.Items)
		}

		// If there is no result, return empty string
		if imageCount == 0 {
			fmt.Println("ERROR: Image search result: No result")
			return ""
		}

		// Pick a random image
		luckyNumber := rand.Intn(imageCount)
		return imageList.Items[luckyNumber].Link
	}

	return ""
}
