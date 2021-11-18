package dice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

type DiceHandler struct {
	maxDiceCount int
	maxDiceSide  int
	diceRegex    *regexp.Regexp
}

func NewDiceHandler(configFile string) (*DiceHandler, error) {
	// Open config file
	cfgFile, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Read dice.cfg: %s", err)
	}

	defer cfgFile.Close()

	// JSON struct
	var diceCfg struct {
		MaxCount int
		MaxSide  int
	}

	// Read and parse json file
	bytes, _ := ioutil.ReadAll(cfgFile)
	err = json.Unmarshal(bytes, &diceCfg)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Parse dice.cfg: %s", err)
	}

	d := &DiceHandler{}
	d.maxDiceCount = diceCfg.MaxCount
	d.maxDiceSide = diceCfg.MaxSide

	d.initRegex()

	fmt.Println("Dice are ready.")
	return d, nil
}

// Init regex once for dice text checks
func (d *DiceHandler) initRegex() {
	// ^ : Start of the line
	// (...) : Groups
	// [0-9]* : Zero or more digits
	// [0-9]+ : One or more digits
	// [+-] : One of plus or minus symbol
	// (...)? : No match or only one match

	// Example values: 1d20, d10, 2d12-2, 3d6 +5
	d.diceRegex = regexp.MustCompile("^([0-9]*)d([0-9]+) *([+-]([0-9]+))?")
}

func (d *DiceHandler) GetResponse(msg string) string {
	// Does the message have any dice text?
	if dice, side, addition := d.getDice(msg); dice > 0 {
		return d.rollDice(dice, side, addition)
	}

	return ""
}

// Check if the message is valid dice text
// Return count, side and addition if these are available
func (d *DiceHandler) getDice(msg string) (count int, side int, addition int) {
	parts := d.diceRegex.FindStringSubmatch(msg)

	// parts[0]: Full match
	// parts[1]: dice count (set 1 if empty)
	// parts[2]: dice side
	// parts[3]: addition (-5, +3 etc.)
	if len(parts) < 4 {
		return 0, 0, 0
	}

	// dice count (set 1 if empty)
	if string(parts[1]) == "" {
		count = 1
	} else {
		count, _ = strconv.Atoi(parts[1])
	}

	side, _ = strconv.Atoi(parts[2])
	addition, _ = strconv.Atoi(parts[3])

	return
}

// Roll dice and generate the output message
func (d *DiceHandler) rollDice(count int, side int, addition int) string {
	// Check the limits (config/dice.json)
	if count > d.maxDiceCount {
		return fmt.Sprintf("Sorry, I don't have %d dice.", count)
	}

	if side > d.maxDiceSide || side < 1 {
		return fmt.Sprintf("Sorry, I don't have any %d sided dice.", side)
	}

	dice := make([]int, count)
	total := 0

	for i := 0; i < count; i++ {
		// roll it
		dice[i] = rand.Intn(side) + 1
		total += dice[i]
	}

	// If count is 2, it most likely has advantage or disadvantage
	if count == 2 {
		min := dice[0]
		max := dice[1]
		if min > max {
			min, max = max, min
		}

		return fmt.Sprintf("> Total: %d\tDisadvantage: %d\tAdvantage: %d\n```%v```", total+addition, min+addition, max+addition, dice)
	} else {
		return fmt.Sprintf("> Total: %d\n```%v```", total+addition, dice)
	}

}
