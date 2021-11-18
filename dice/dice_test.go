package dice

import "testing"

func initDiceTest() *DiceHandler {
	d := &DiceHandler{}

	d.maxDiceCount = 10
	d.maxDiceSide = 100

	d.initRegex()
	return d
}

func TestCheckMessageForDice(t *testing.T) {
	d := initDiceTest()

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
		if result[0], result[1], result[2] = d.getDice(test.input); result != test.output {
			t.Errorf("For %s, %v is expected but %v is returned.", test.input, test.output, result)
		}
	}
}

// TODO: Add tests for random dice
func TestRollDice(t *testing.T) {
	d := initDiceTest()

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
		if result := d.rollDice(test.input[0], test.input[1], test.input[2]); result != test.output {
			t.Errorf("For %v, %s is expected but %s is returned.", test.input, test.output, result)
		}
	}
}

func BenchmarkCheckMessageForDice(b *testing.B) {
	d := initDiceTest()

	b.Run("Invalid short text", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d.getDice("Roll")
		}
	})

	b.Run("Invalid long text", func(b *testing.B) {
		text := "Something else"
		for len(text) < 2000 {
			text += text
		}

		for i := 0; i < b.N; i++ {
			d.getDice(text)
		}
	})

	b.Run("Valid text", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d.getDice("8d20 +1")
		}
	})
}

func BenchmarkRollDice(b *testing.B) {
	d := initDiceTest()

	b.Run("1d20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d.rollDice(1, 2, 0)
		}
	})

	b.Run("2d20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d.rollDice(2, 20, 0)
		}
	})

	b.Run("8d20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d.rollDice(8, 2, 0)
		}
	})

	b.Run("8d20 +6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d.rollDice(8, 2, 6)
		}
	})
}
