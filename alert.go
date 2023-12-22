package septa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// APIBool is a type of bool that can deal with the range of data that the Septa API returns for a bool
// which is any of "N", "Yes", "0", 0, etc.
type APIBool bool

// UnmarshalJSON takes the values in JSON and maps them to the golang types.  This is where N, No, etc are
// mapped to the bool type.
func (item *APIBool) UnmarshalJSON(data []byte) error {
	value := string(data)
	value = strings.ReplaceAll(value, "\"", "")
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
	// Name3 is the three letter designation for the route.
	Name3             string  `json:"route"`
	RouteID           string  `json:"route_id"`
	RouteName         string  `json:"route_name"`
	Sequence          string  `json:"sequence"`
	Mode              string  `json:"mode"`
	IsAdvisory        APIBool `json:"isadvisory"`
	IsDetour          APIBool `json:"isdetour"`
	IsAlert           APIBool `json:"isalert"`
	IsSuspend         APIBool `json:"issuppend"`
	IsElevator        APIBool `json:"iselevator"`
	IsSuspended       APIBool `json:"issuspended"`
	IsStrike          APIBool `json:"isstrike"`
	IsModifiedService APIBool `json:"ismodifiedservice"`
	IsDelays          APIBool `json:"isdelays"`
	CurrentMessage    string  `json:"current_message"`
	Advisory          string  `json:"advisory"`
}

// String prints a formatted version of the AlertResult
func (r AlertResult) String() string {
	if r.CurrentMessage != "" {
		return fmt.Sprintf("%v (%v) %v", r.RouteName, r.RouteID, r.CurrentMessage)
	}
	if r.IsAdvisory {
		return fmt.Sprintf("%v (%v) *Advisory* %v", r.RouteName, r.RouteID, r.Advisory)
	}
	return fmt.Sprintf("%v (%v) no alerts or advisories.", r.RouteName, r.RouteID)
}

// Alerts makes a call to the alerts api and returns the results.  If no matching route is found, alerts for all routes will be displayed.
func (c *Client) Alerts(route string) ([]AlertResult, error) {
	var results []AlertResult

	if route == "" {
		route = "all"
	} else {
		route = strings.ToLower(route)
		route = Routes[route]
	}
	url := fmt.Sprintf("https://%s/api/Alerts/index.php?routes=%s", c.Domain, url.QueryEscape(route))

	// #nosec ignoring G107 here because we have cleaned the input above by escaping the route name.
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error calling alerts err=%v\n", err)
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
