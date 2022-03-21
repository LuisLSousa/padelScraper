package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const configFile = "config/config.json"

type Parameters struct {
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	OnlyIndoor bool   `json:"only_indoor"`
	MinSlots   int    `json:"min_slots"`
}

func (p *Parameters) ImportFromFile() {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error opening the config file: %s", err)
		return
	}
	err = json.Unmarshal(file, p)
	if err != nil {
		fmt.Printf("Error unmarshalling config file: %s", err)
		return
	}

}
