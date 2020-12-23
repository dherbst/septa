package septa

import (
	"testing"
)

func TestSeptaAPIBool(t *testing.T) {
	testData := map[string]bool{
		"Yes": true,
		"Y":   true,
		"No":  false,
		"N":   false,
		"err": false,
	}

	var value SeptaAPIBool = false

	for input, _ := range testData {
		err := value.UnmarshalJSON([]byte(input))
		if input == "err" && err == nil {
			t.Fatalf("Expected error for unknown value")
		}
		if bool(value) != testData[input] {
			t.Fatalf("Expected %v for %v got %v\n", testData[input], input, value)
		}
	}
}
