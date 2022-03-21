package httpClient

import (
	"encoding/json"
	"io"
)

type Response struct {
	Results []PadelField `json:"results"`
}

type PadelField struct {
	Slots []Slot `json:"slots"`
}

type Slot struct {
	Date   string `json:"date"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Locked bool   `json:"locked"`
	Status string `json:"status"`
}

func getBody(body io.Reader, resp interface{}) error {
	dec := json.NewDecoder(body)
	if err := dec.Decode(resp); err != nil {
		return err
	}
	return nil
}
