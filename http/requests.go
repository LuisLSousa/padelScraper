package httpClient

import (
	"fmt"
	"net/http"
)

type Client struct {
	url string
}

func New(url string) Client {
	return Client{url: url}
}

func (c Client) Get() Response {
	var response Response
	res, err := http.Get(c.url)
	if err != nil {
		fmt.Printf("Error during Get: %s", err)
		return Response{}
	}
	if err = getBody(res.Body, &response); err != nil {
		fmt.Printf("Error decoding response: %s", err)
		return Response{}
	}
	return response
}
