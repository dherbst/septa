package septa

import "errors"

// SeptaAPIBool is a type of bool that can deal with the range of data that the Septa API returns for a bool
// which is any of "N", "Yes", "0", 0, etc.
type SeptaAPIBool bool

func (item *SeptaAPIBool) UnmarshalJSON(data []byte) error {
	value := string(data)
	if value == "Yes" || value == "Y" || value == "true" {
		*item = true
	} else if value == "No" || value == "N" || value == "false" {
		*item = false
	} else {
		return errors.New("Unexpected bool value " + value)
	}
	return nil
}

// AlertResult is a line returned from the alert API.
type AlertResult struct {
	RouteID    string       `json:"route_id"`
	RouteName  string       `json:"route_name"`
	Mode       string       `json:"mode"`
	IsAdvisory SeptaAPIBool `json:"isadvisory"`
}

// Alerts makes a call to the alerts api and returns the results
func (c *Client) Alerts(route string) ([]AlertResult, error) {
	var results []AlertResult

	return results, nil
}
