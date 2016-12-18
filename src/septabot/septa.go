package septabot

import (
	"context"
	"errors"
)

// SeptaAPI handles interacting with the Septa API to get information.
type Septa struct {
	api SeptaAPI
}

// NewSepta creates a new Septa object for interacting with the Septa API
func NewSepta(apiImpl SeptaAPI) *Septa {
	septa := &Septa{
		api: apiImpl,
	}
	return septa
}

// NextToArrive returns the num departures for the station.
// Calls http://www3.septa.org/hackathon/NextToArrive/Suburban%20Station/Narberth/10
func (s *Septa) NextToArrive(ctx context.Context, fromStation string, toStation string, num int) ([]NextToArriveResult, error) {
	var result []NextToArriveResult
	var err error

	if fromStation == "" {
		return result, errors.New("missing fromStation")
	}
	if toStation == "" {
		return result, errors.New("missing toStation")
	}
	if num < 1 {
		return result, errors.New("num must be greater than zero")
	}

	result, err = s.api.NextToArrive(ctx, fromStation, toStation, num)

	return result, err
}
