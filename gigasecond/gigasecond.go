package gigasecond

import (
	"time"
)

// gigasecond value string
const gigasecond string = "1000000000s"

// AddGigasecond adds a gigasecond to the input
func AddGigasecond(t time.Time) time.Time {
	gs, _ := time.ParseDuration(gigasecond)
	return t.Add(gs)
}
