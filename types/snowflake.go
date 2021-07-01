package types

import (
	"strconv"
	"time"
)

// Snowflake represents a Twitter snowflake.
// See https://discord.com/developers/docs/reference#snowflakes for more details.
type Snowflake struct {
	Time      time.Time
	WorkerID  int64
	ProcessID int64
	Increment int64
}

// NewSnowflake parses a Snowflake
func NewSnowflake(snowflake string) (sw *Snowflake, err error) {
	num, err := strconv.ParseInt(snowflake, 10, 64)

	sw = &Snowflake{
		time.Unix(((num>>22)+1420070400000)/1000, 0),
		(num & 0x3E0000) >> 17,
		(num & 0x1F000) >> 12,
		num & 0xFFF,
	}

	return
}
