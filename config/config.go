package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const (
	configFile = "config/config.json"
	dateLayout = "2006-01-02"
)

type Parameters struct {
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	OnlyIndoor bool   `json:"only_indoor"`
	MinSlots   int    `json:"min_slots"`
}

func (p *Parameters) ImportFromFile() error {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error opening the config file: %s", err)
		return err
	}

	err = json.Unmarshal(file, p)
	if err != nil {
		fmt.Printf("Error unmarshalling config file: %s", err)
		return err
	}

	if err = p.validateFields(); err != nil {
		return err
	}

	return nil
}

func (p Parameters) validateFields() error {
	if p.MinSlots < 2 {
		return fmt.Errorf("min_slots must be at least 2 (number of consecutive 30 minute slots)")
	}

	startDate, err := time.Parse(dateLayout, p.StartDate)
	if err != nil {
		fmt.Printf("Invalid start_date: %s", err)
	}
	endDate, err := time.Parse(dateLayout, p.EndDate)
	if err != nil {
		fmt.Printf("Invalid end_date: %s", err)
	}

	if !startDate.Before(endDate) && startDate != endDate {
		return fmt.Errorf("start_date must be before or equal to end_date")
	}

	if startDate.Before(time.Now()) || endDate.Before(time.Now()) {
		return fmt.Errorf("start_date and end_date cannot be before today")
	}

	if endDate.After(time.Now().AddDate(0, 0, 5)) {
		return fmt.Errorf("end_date cannot be more than five days away from now")
	}
	return nil
}
