package septa

var stations = []string{
	"Narberth",
	"Suburban Station",
}

// IsValidStation returns whether the station is a valid station.
func IsValidStation(station string) bool {
	for _, v := range stations {
		if station == v {
			return true
		}
	}
	return false
}
