package goocord

import (
	"fmt"
	"testing"
)

var f Flags = Flags{0}

func TestFlags_Add(t *testing.T) {
	f.Add(1 << 7)
	f.Add(1 << 8)
	if f.Flags != 1<<7|1<<8 {
		t.Error(fmt.Sprintf("expected - %d, got - %d", 1>>7|1>>8, f.Flags))
	}
}

func TestFlags_Has(t *testing.T) {
	if !f.Has(1<<7) || !f.Has(1<<8) {
		t.Error("shit")
	}

	if f.Has(1 << 9) {
		t.Error("expected bitfield not to include 1<<9")
	}
}
