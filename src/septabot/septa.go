package septabot

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// SeptaAPI handles interacting with the Septa API to get information.
type SeptaAPI struct {
	Domain string
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

// NextToArrive returns the num departures for the station.
// Calls http://www3.septa.org/hackathon/NextToArrive/Suburban%20Station/Narberth/10
func (api *SeptaAPI) NextToArrive(fromStation string, toStation string, num int) ([]NextToArriveResult, error) {
	var result []NextToArriveResult
	if fromStation == "" {
		return result, errors.New("missing fromStation")
	}
	if toStation == "" {
		return result, errors.New("missing toStation")
	}
	if num < 1 {
		return result, errors.New("num must be greater than zero")
	}

	url := fmt.Sprintf("http://%s/hackathon/NextToArrive/%s/%s/%d", api.Domain, fromStation, toStation, num)
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Error calling %v err=%v\n", url, err)
		return result, err
	}

	return result, nil
}
