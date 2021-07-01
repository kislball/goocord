package types

import "testing"

func TestNewSnowflake(t *testing.T) {
	snowflake, err := NewSnowflake("419524085736013834")
	if err != nil {
		t.Error(err)
	}

	if snowflake.Time.Unix() != 1520092736 {
		t.Error("expected snowflake.Time.Unix() to be 1520092736")
	}
}
