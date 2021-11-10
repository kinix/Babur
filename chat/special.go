// This package can work wrong. It is not well tested. I just wrote that for fun
package chat

import (
	"fmt"
	"math/rand"
)

// If response begins with %%, Babür can do some calculations to respond
// This function get fullMsg as parameter. It is useless for now, but it can be used in future (maybe image search?)
func specialResponse(specialType string, fullMsg string) string {
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
	}

	return ""
}
