package main

import "testing"

func initDiceTest() {
	maxDiceCount = 10
	maxDiceSide = 100

	initDiceRegex()
}

func TestCheckMessageForDice(t *testing.T) {
	initDiceTest()

	type testCase struct {
		input  string
		output [3]int
	}

	testCases := []testCase{
		{"Roll 1d20", [3]int{0, 0, 0}},
		{"1d20", [3]int{1, 20, 0}},
		{"d6", [3]int{1, 6, 0}},
		{"5d12 -2", [3]int{5, 12, -2}},
		{"12d10", [3]int{12, 10, 0}},
	}

	var result [3]int
	for _, test := range testCases {
		if result[0], result[1], result[2] = checkMessageForDice(test.input); result != test.output {
			t.Errorf("For %s, %v is expected but %v is returned.", test.input, test.output, result)
		}
	}
}

// TODO: Add tests for random dice
func TestRollDice(t *testing.T) {
	initDiceTest()

	type testCase struct {
		input  [3]int
		output string
	}

	testCases := []testCase{
		{[3]int{0, 0, 0}, "Sorry, I don't have any 0 sided dice."},
		{[3]int{5, 200, 0}, "Sorry, I don't have any 200 sided dice."},
		{[3]int{20, 20, 0}, "Sorry, I don't have 20 dice."},
	}

	for _, test := range testCases {
		if result := rollDice(test.input[0], test.input[1], test.input[2]); result != test.output {
			t.Errorf("For %v, %s is expected but %s is returned.", test.input, test.output, result)
		}
	}
}
