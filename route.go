package septa

// Route is a rail route.
type Route struct {
	// ThreeID is a three letter designation for the route
	ThreeID string
	// ID is the api route_id used.
	ID string
	// Name is the route_name used in the api
	Name string
	// Mode is the type of rail, one of Regional Rail, Bus, or there are some other modes.
	Mode string
}

// Routes is a map of the common name to the Route used in the api.
var Routes = map[string]string{
	"Broad Street Line": "rr_route_bsl",
	"bsl":               "rr_route_bsl",
	"paoli":             "rr_route_pao",
	"thorndale":         "rr_route_pao",
}
