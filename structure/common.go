package structure

import "time"

type TimeRange struct {
	Valid bool
	Start time.Time
	End   time.Time
}
