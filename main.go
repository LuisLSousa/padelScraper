package main

import (
	"fmt"
	"time"

	"padelScraper/config"
)

// Get Endpoint
// https://www.aircourts.com/index.php/api/search_with_club/355?sport=0&date=2022-03-18&start_time=18:00
// start date cannot be 5 days after the current day

// Rackets Pro Saldanha club_id = 355

const (
	dateLayout      = "2006-01-02"
	baseTargetUrl   = "https://www.aircourts.com/index.php/api/search_with_club/355?sport=0"
	statusAvailable = "available"
	statusBooked    = "booked"
)

func main() {
	var configParameters config.Parameters
	configParameters.ImportFromFile()

	startDate, err := time.Parse(dateLayout, configParameters.StartDate)
	if err != nil {
		fmt.Printf("Invalid start_date: %s", err)
	}

	endDate, err := time.Parse(dateLayout, configParameters.EndDate)
	if err != nil {
		fmt.Printf("Invalid end_date: %s", err)
	}

	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		//targetUrl := fmt.Sprintf("%s&date=%s&start_time=07:00", baseTargetUrl, d)
		//hc := httpClient.New(targetUrl)
		//resp := hc.Get()
		fmt.Printf("\nDate: %s\n", d)
		//fmt.Printf("%v", resp.Results[0].Slots)

	}

	// iterate over the slots

}
