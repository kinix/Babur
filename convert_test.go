package main

import "testing"

func initConvertTest() {
	units = map[string]Unit{
		"feet": {"cm", 30.48},
		"ft":   {"cm", 30.48},
		"_cm":  {"m", 0.01},
		"_m":   {"km", 0.001},
	}

	initUnitRegex()
}

func TestCheckMessageForConverting(t *testing.T) {
	initConvertTest()

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
		results := checkMessageForConverting(test.input)
		for i, result := range results {
			if result != test.output[i] {
				t.Errorf("For %s, %v is expected but %v is returned.", test.input, test.output[i], result)
			}
		}
	}
}

func TestConvertUnits(t *testing.T) {
	initConvertTest()

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
		if result := convertUnits(test.input); result != test.output {
			t.Errorf("For %v, %s is expected but %s is returned.", test.input, test.output, result)
		}
	}
}
