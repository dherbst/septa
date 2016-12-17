package septabot

import (
// "log"
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
func (api *SeptaAPI) NextToArrive(station string, num int) ([]NextToArriveResult, error) {
	var result []NextToArriveResult

	return result, nil
}
