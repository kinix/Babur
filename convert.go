package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Unit struct {
	ConvertedUnit string  `json:"unit"`
	Multiplier    float64 `json:"multiplier"`
}

type rawMeasurement struct {
	value float64
	unit  string
}

var units map[string]Unit
var unitRegex *regexp.Regexp

// Init regex once for unit text checks
func initUnitRegex() {
	// Create string include all units (e.g. feet|ft|inch|)
	unitList := ""
	for key := range units {
		unitList += key + "|"
	}

	// Remove the last "|""
	unitList = unitList[:len(unitList)-1]

	// (...) : Groups
	// [0-9]* : Zero or more digits

	// Example values: 2 feet, 3.1 inch, .5 mile
	regexRule := fmt.Sprintf("([0-9]*\\.?[0-9]*) *(%s)", unitList)
	unitRegex = regexp.MustCompile(regexRule)
}

// Check if the message is valid measurement text to convert
// Return measurement value and unit couples if these are available
func checkMessageForConverting(msg string) []rawMeasurement {
	partList := unitRegex.FindAllStringSubmatch(msg, -1)
	result := []rawMeasurement{}

	for _, parts := range partList {
		// parts[0]: Full match
		// parts[1]: Measurement value
		// parts[2]: Measurement unit
		if len(parts) == 3 {
			if val, _ := strconv.ParseFloat(parts[1], 64); val != 0 {
				result = append(result, rawMeasurement{val, parts[2]})
			}
		}
	}

	return result
}

// Convert measurements and generate the output message
func convertUnits(measurements []rawMeasurement) string {
	results := make([]string, len(measurements))

	var convertedVal float64
	var convertedUnit string

	for i, measurement := range measurements {
		// Convert
		convertedVal = measurement.value * units[measurement.unit].Multiplier
		convertedUnit = units[measurement.unit].ConvertedUnit

		// If result value is too big, convert to another unit if another unit is available
		// These units are represented with _ prefix in the config file
		newUnit, exist := units["_"+convertedUnit]
		for exist && convertedVal > 1/newUnit.Multiplier {
			convertedVal *= newUnit.Multiplier
			convertedUnit = newUnit.ConvertedUnit

			newUnit, exist = units["_"+convertedUnit]
		}

		// Format values as x.xx (remove zeros if decimal is zero)
		formattedValue := fmt.Sprintf("%.2f", measurement.value)
		formattedValue = strings.TrimRight(formattedValue, "0")
		formattedValue = strings.TrimRight(formattedValue, ".")

		formattedConvertedValue := fmt.Sprintf("%.2f", convertedVal)
		formattedConvertedValue = strings.TrimRight(formattedConvertedValue, "0")
		formattedConvertedValue = strings.TrimRight(formattedConvertedValue, ".")

		results[i] = fmt.Sprintf("%s %s = %s %s", formattedValue, measurement.unit, formattedConvertedValue, convertedUnit)
	}

	return fmt.Sprintf("```%s```", strings.Join(results, "\n"))
}
