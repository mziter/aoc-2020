package six

import "testing"

func TestSomething(t *testing.T) {
	if 1 != 1 {
		t.Error("something is very wrong...")
	}
}
