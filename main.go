package main

import (
	"fmt"
	"time"

	"padelScraper/config"
	httpClient "padelScraper/http"
)

// Get Endpoint
// https://www.aircourts.com/index.php/api/search_with_club/355?sport=0&date=2022-03-18&start_time=18:00
// start date cannot be 5 days after the current day

// Rackets Pro Saldanha club_id = 355

const (
	dateLayout      = "2006-01-02"
	hourLayout      = "15:04"
	baseTargetUrl   = "https://www.aircourts.com/index.php/api/search_with_club/355?sport=0"
	statusAvailable = "available"
)

func main() {
	var configParameters config.Parameters
	var availableSlots availableSlots

	if err := configParameters.ImportFromFile(); err != nil {
		fmt.Printf("Error during configuration: %s", err)
		return
	}

	startDate, _ := time.Parse(dateLayout, configParameters.StartDate)
	endDate, _ := time.Parse(dateLayout, configParameters.EndDate)

	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {

		year, month, day := d.Date()
		date := fmt.Sprintf("%d-%d-%d", year, int(month), day)

		availableSlot := availableSlot{
			date: date,
		}

		// target url is in the form
		// https://www.aircourts.com/index.php/api/search_with_club/355?sport=0&date=2022-03-18&start_time=18:00
		targetUrl := fmt.Sprintf("%s&date=%s&start_time=07:00", baseTargetUrl, date)
		hc := httpClient.New(targetUrl)
		resp := hc.Get()

		for _, slot := range resp.Results[0].Slots {
			// for each available slot (not booked)
			if slot.Status == statusAvailable && !slot.Locked {
				// check if there are enough slots available after it for (configParameters.MinSlots * 30) minutes of play time
				if slot.Forward >= (configParameters.MinSlots - 1) {
					start, _ := time.Parse(hourLayout, slot.Start)
					end := start.Add(time.Minute * 30 * time.Duration(configParameters.MinSlots))
					endString := end.Format(hourLayout)
					timeframe := fmt.Sprintf("%s - %s", slot.Start, endString)
					fmt.Printf("Timeframe: %s", timeframe)
					//availableSlot.hours = append(availableSlot.hours, timeframe)
				}
			}
			//fmt.Printf("%v", availableSlots)

			availableSlots = append(availableSlots, availableSlot)
		}

	}
	//// iterate over the slots
	//for _, as := range availableSlots {
	//	fmt.Printf("\n\nSlot: %#v\n\n", as)
	//}

}
