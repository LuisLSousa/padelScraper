# Padel Scraper
This is a script to get the available padel fields from [Aircourts](https://www.aircourts.com/) for a given date, and print the results to std out.

Currently, the script only searches for slots in a specific club (the one closest to my office), but it can easily be extended by changing the club id on the path of the GET request.

## Running the script
In the config folder, define the desired start and end dates, and the minimum number of slots of 30 minutes available sequentially (if you want to play for 1h30, set `min_slots = 3`).

After that, simply run
```go run main.go```
