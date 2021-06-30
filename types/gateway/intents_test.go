package gateway

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	if AllIntents.Flags != 32767 {
		t.Error(fmt.Sprintf("expected AllIntents.Flags to be 32797, got - %d", AllIntents.Flags))
	}

	if UnprivilegedIntents.Flags != 32509 {
		t.Error(fmt.Sprintf("expected UnprivilegedIntents.Flags to be 32509, got - %d", UnprivilegedIntents.Flags))
	}
}
