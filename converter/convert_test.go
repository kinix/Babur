package converter

import "testing"

func initConvertTest() *ConverterHandler {
	c := &ConverterHandler{}
	c.units = map[string]Unit{
		"feet": {"cm", 30.48},
		"ft":   {"cm", 30.48},
		"_cm":  {"m", 0.01},
		"_m":   {"km", 0.001},
	}

	c.initRegex()
	return c
}

func TestCheckMessageForConverting(t *testing.T) {
	c := initConvertTest()

	type testCase struct {
		input  string
		output []rawMeasurement
	}

	testCases := []testCase{
		{
			"10 feet",
			[]rawMeasurement{
				{10, "feet"},
			},
		},
		{
			"10 felt",
			[]rawMeasurement{},
		},
		{
			"It is between 3.5 feet and 5 ft",
			[]rawMeasurement{
				{3.5, "feet"},
				{5, "ft"},
			},
		},
	}

	for _, test := range testCases {
		results := c.getUnits(test.input)
		for i, result := range results {
			if result != test.output[i] {
				t.Errorf("For %s, %v is expected but %v is returned.", test.input, test.output[i], result)
			}
		}
	}
}

func TestConvertUnits(t *testing.T) {
	c := initConvertTest()

	type testCase struct {
		input  []rawMeasurement
		output string
	}

	testCases := []testCase{
		{
			[]rawMeasurement{
				{10, "feet"},
			},
			"```10 feet = 3.05 m```",
		},
		{
			[]rawMeasurement{
				{1000, "feet"},
				{3, "ft"},
			},
			"```1000 feet = 304.8 m\n3 ft = 91.44 cm```",
		},
	}

	for _, test := range testCases {
		if result := c.convertUnits(test.input); result != test.output {
			t.Errorf("For %v, %s is expected but %s is returned.", test.input, test.output, result)
		}
	}
}

func BenchmarkCheckMessageForConverting(b *testing.B) {
	c := initConvertTest()

	b.Run("Invalid short text", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.getUnits("Convert")
		}
	})

	b.Run("Invalid long text", func(b *testing.B) {
		text := "Something else"
		for len(text) < 2000 {
			text += text
		}

		for i := 0; i < b.N; i++ {
			c.getUnits(text)
		}
	})

	b.Run("Single valid short text", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.getUnits("10 feet")
		}
	})

	b.Run("Single valid long text", func(b *testing.B) {
		text := "Something else"
		for len(text) < 2000 {
			text += text
		}
		text += "10 feet"

		for i := 0; i < b.N; i++ {
			c.getUnits(text)
		}
	})

	b.Run("Multiple valid long text", func(b *testing.B) {
		text := "Something else 10 feet"
		for len(text) < 2000 {
			text += text
		}
		text += "10 feet"

		for i := 0; i < b.N; i++ {
			c.getUnits(text)
		}
	})
}

func BenchmarkConvertUnits(b *testing.B) {
	c := initConvertTest()

	b.Run("Single simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.convertUnits([]rawMeasurement{
				{1, "feet"},
			})
		}
	})

	b.Run("Multiple simple", func(b *testing.B) {
		list := []rawMeasurement{}
		for i := 0; i < 100; i++ {
			list = append(list, rawMeasurement{float64(i), "feet"})
		}

		for i := 0; i < b.N; i++ {
			c.convertUnits(list)
		}
	})

	b.Run("Single big", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.convertUnits([]rawMeasurement{
				{1000000, "feet"},
			})
		}
	})

	b.Run("Multiple big", func(b *testing.B) {
		list := []rawMeasurement{}
		for i := 0; i < 100; i++ {
			list = append(list, rawMeasurement{float64(i) * 1000000, "feet"})
		}

		for i := 0; i < b.N; i++ {
			c.convertUnits(list)
		}
	})
}
