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
		fmt.Println("ERROR: Read dice.cfg: ", err)
		return err
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
		fmt.Println("ERROR: Parse dice.cfg: ", err)
		return err
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
		fmt.Println("ERROR: Read unit.cfg: ", err)
		return err
	}

	defer cfgFile.Close()

	// Read and parse json file
	bytes, _ := ioutil.ReadAll(cfgFile)
	err = json.Unmarshal(bytes, &units)
	if err != nil {
		fmt.Println("ERROR: Parse unit.cfg: ", err)
		return err
	}

	fmt.Println("Converter is ready.")
	return nil
}
