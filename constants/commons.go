package constants

import "time"

type commons struct {
	DateFormat      string
	DateShortFormat string
}

var Commons = commons{
	DateFormat:      time.RFC3339,
	DateShortFormat: "2006-01-02",
}
