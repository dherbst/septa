package septa

import (
	"encoding/json"
	"testing"
)

func TestNextToArriveResult(t *testing.T) {
	response := `[{"orig_train":"3523","orig_line":"Paoli\/Thorndale","orig_departure_time":"10:05AM","orig_arrival_time":"10:22AM","orig_delay":"4 mins","isdirect":"true"},{"orig_train":"1525","orig_line":"Paoli\/Thorndale","orig_departure_time":"10:35AM","orig_arrival_time":"10:52AM","orig_delay":"On time","isdirect":"true"},{"orig_train":"3527","orig_line":"Paoli\/Thorndale","orig_departure_time":"11:05AM","orig_arrival_time":"11:22AM","orig_delay":"11 mins","isdirect":"true"},{"orig_train":"1529","orig_line":"Paoli\/Thorndale","orig_departure_time":"11:35AM","orig_arrival_time":"11:52AM","orig_delay":"On time","isdirect":"true"},{"orig_train":"3531","orig_line":"Paoli\/Thorndale","orig_departure_time":"12:05PM","orig_arrival_time":"12:22PM","orig_delay":"On time","isdirect":"true"},{"orig_train":"1533","orig_line":"Paoli\/Thorndale","orig_departure_time":"12:35PM","orig_arrival_time":"12:52PM","orig_delay":"On time","isdirect":"true"},{"orig_train":"3535","orig_line":"Paoli\/Thorndale","orig_departure_time":" 1:05PM","orig_arrival_time":" 1:22PM","orig_delay":"On time","isdirect":"true"},{"orig_train":"1537","orig_line":"Paoli\/Thorndale","orig_departure_time":" 1:35PM","orig_arrival_time":" 1:52PM","orig_delay":"On time","isdirect":"true"},{"orig_train":"3539","orig_line":"Paoli\/Thorndale","orig_departure_time":" 2:05PM","orig_arrival_time":" 2:22PM","orig_delay":"On time","isdirect":"true"},{"orig_train":"1541","orig_line":"Paoli\/Thorndale","orig_departure_time":" 2:30PM","orig_arrival_time":" 2:47PM","orig_delay":"On time","isdirect":"true"}]`

	var results []NextToArriveResult
	err := json.Unmarshal([]byte(response), &results)
	if err != nil {
		t.Fatalf("Error unmarshalling result %v\n", err)
	}

	if len(results) != 10 {
		t.Fatalf("Expected 10 results, got %v\n", len(results))
	}

	for _, r := range results {
		t.Logf("Train %v Depart %v Late %v\n", r.Train, r.DepartureString, r.Delay)
	}

}
