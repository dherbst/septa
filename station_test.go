package septa

import "testing"

// ValidStationTestData holds the test data and expected output of the IsValidStation function.
type ValidStationTestData struct {
	Name           string
	ExpectedResult bool
}

func TestIsValidStation(t *testing.T) {
	data := []ValidStationTestData{
		{"Narberth", true},
		{"Suburban Station", true},
		{"aaaa", false},
		{"", false},
	}

	for i, v := range data {
		result := IsValidStation(v.Name)
		if result != v.ExpectedResult {
			t.Fatalf("Expected test %v to be %v but got %v for %v\n",
				i, v.ExpectedResult, result, v.Name)
		}
	}
}
