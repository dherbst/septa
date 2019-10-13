package septa

import "time"

// Client is used to make calls to the septa website.
type Client struct {

	// Domain is the api domain.
	Domain string

	// Timeout is the number of seconds before the call times out.
	Timeout time.Duration
}
