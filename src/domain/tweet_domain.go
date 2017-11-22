package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
}
