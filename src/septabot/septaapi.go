package septabot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// SeptaAPI calls the Septa API.
type SeptaAPI interface {
	NextToArrive(ctx context.Context, fromStation string, toStation string, num int) ([]NextToArriveResult, error)
}

// NextToArriveResult holds the num results.
type NextToArriveResult struct {
	Train           string `json:"orig_train"`
	Line            string `json:"orig_line"`
	DepartureString string `json:"orig_departure_time"`
	ArrivalString   string `json:"orig_arrival_time"`
	Delay           string `json:"orig_delay"`
	IsDirect        string `json:"isdirect"`
}

// SeptaAPIImpl calls the api
type SeptaAPIImpl struct {
	domain string
}

// NewSeptaAPIImpl returns a pointer to a struct that can call the api
func NewSeptaAPIImpl() *SeptaAPIImpl {
	api := &SeptaAPIImpl{domain: "www3.septa.org"}
	return api
}

func (api *SeptaAPIImpl) NextToArrive(ctx context.Context,
	fromStation string,
	toStation string,
	num int) ([]NextToArriveResult, error) {

	var results []NextToArriveResult
	url := fmt.Sprintf("http://%s/hackathon/NextToArrive/%s/%s/%d", api.domain, fromStation, toStation, num)
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		log.Printf("Error calling %v err=%v\n", url, err)
		return results, err
	}

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response.Body %v\n", err)
		return results, err
	}

	err = json.Unmarshal(result, &results)
	if err != nil {
		log.Printf("Error unmarshalling results %v\n%v", err, string(result))
		return results, err
	}

	return results, nil
}
