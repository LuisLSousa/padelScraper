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

// Note: Rackets Pro Saldanha club_id = 355

type availableSlots []availableSlot

type availableSlot struct {
	club  string
	court string
	date  string
	hours []string
}

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

	for d := startDate; d.Before(endDate) || d == endDate; d = d.AddDate(0, 0, 1) {

		year, month, day := d.Date()
		date := fmt.Sprintf("%d-%d-%d", year, int(month), day)

		// target url is in the form
		// https://www.aircourts.com/index.php/api/search_with_club/355?sport=0&date=2022-03-18&start_time=18:00
		targetUrl := fmt.Sprintf("%s&date=%s&start_time=07:00", baseTargetUrl, date)
		hc := httpClient.New(targetUrl)
		resp := hc.Get()

		availableSlot := availableSlot{
			date: date,
			club: resp.Results[0].ClubName,
		}
		for n, _ := range resp.Results {
			// TODO add the hours to show 07:00 - 12:00 instead of 07:00 - 09:00 09:00 - 11:00, etc.
			availableSlot.court = resp.Results[n].Court
			availableSlot.hours = nil
			skipSlots := 0
			for i, slot := range resp.Results[n].Slots {
				if i < skipSlots {
					continue
				}
				// for each available slot (not booked)
				if slot.Status == statusAvailable && !slot.Locked {
					// check if there are enough slots available after it for (configParameters.MinSlots * 30) minutes of play time
					if slot.Forward >= (configParameters.MinSlots - 1) {
						start, _ := time.Parse(hourLayout, slot.Start)
						end := start.Add(time.Minute * 30 * time.Duration(slot.Forward))
						endFormatted := end.Format(hourLayout)

						timeframe := fmt.Sprintf("%s - %s", slot.Start, endFormatted)

						availableSlot.hours = append(availableSlot.hours, timeframe)
						skipSlots = i + slot.Forward
					}
				}
			}
			availableSlots = append(availableSlots, availableSlot)
		}
	}
	// iterate over the slots
	var date string
	for _, as := range availableSlots {
		// if there are any available slots for this date
		if date != as.date {
			date = as.date
			fmt.Printf("\n\n## Date: %s ##", as.date)
			fmt.Printf("\nClub: %s", as.club)
		}

		if len(as.hours) > 0 {
			fmt.Printf("\nCourt: %s", as.court)
			fmt.Printf("\nTime:")
			for _, hour := range as.hours {
				fmt.Printf(" %s |", hour)
			}
		}
	}

}
