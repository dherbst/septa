package septa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Client is used to make calls to the septa website.
type Client struct {

	// Domain is the api domain.
	Domain string

	// Timeout is the number of seconds before the call times out.
	Timeout time.Duration
}

// NewClient creates a Client with the given domain to make api calls.
func NewClient(domain string) *Client {
	if domain == "" {
		domain = "www3.septa.org"
	}
	c := &Client{Domain: domain}
	return c
}

// NextToArriveResult is a line in the NTA response.
type NextToArriveResult struct {
	Train           string `json:"orig_train"`
	Line            string `json:"orig_line"`
	DepartureString string `json:"orig_departure_time"`
	ArrivalString   string `json:"orig_arrival_time"`
	Delay           string `json:"orig_delay"`
	IsDirect        string `json:"isdirect"`
}

// NextToArrive makes a call to the next to arrive api and returns the results
func (c *Client) NextToArrive(from string, to string, num int) ([]NextToArriveResult, error) {
	url := fmt.Sprintf("https://%s/hackathon/NextToArrive/%s/%s/%d",
		c.Domain,
		url.PathEscape(from),
		url.PathEscape(to),
		num)
	log.Printf("url=%v\n", url)

	var results []NextToArriveResult
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error calling NextToArrive err=%v\n", err)
		return results, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response.Body %v\n", err)
		return results, err
	}
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Printf("Error unmarshalling results %v\n%v", err, string(body))
		return results, err
	}

	return results, nil

}
