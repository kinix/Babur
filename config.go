package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func initDiceConfig() error {
	// Open config file
	cfgFile, err := os.Open("config/dice.json")
	if err != nil {
		return fmt.Errorf("ERROR: Read dice.cfg: %s", err)
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
		return fmt.Errorf("ERROR: Parse dice.cfg: %s", err)
	}

	maxDiceCount = diceCfg.MaxCount
	maxDiceSide = diceCfg.MaxSide

	fmt.Println("Dice are ready.")
	return nil
}

func initUnitConfig() error {
	// Open config file
	cfgFile, err := os.Open("config/units.json")
	if err != nil {
		return fmt.Errorf("ERROR: Read unit.cfg: %s", err)
	}

	defer cfgFile.Close()

	// Read and parse json file
	bytes, _ := ioutil.ReadAll(cfgFile)
	err = json.Unmarshal(bytes, &units)
	if err != nil {
		return fmt.Errorf("ERROR: Parse unit.cfg: %s", err)
	}

	fmt.Println("Converter is ready.")
	return nil
}
