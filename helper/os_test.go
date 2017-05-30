package helper

import (
	"os"
	"testing"
)

// TestRemoval tests the Unset method
func TestRemoval(t *testing.T) {
	e := Environ(os.Environ())
	e = append(e, "test=test")
	e.Unset("test")
	for _, val := range e {
		if val == "test=test" {
			t.Log("Found test=test which was expected to be removed")
			t.Fail()
		}
	}
}
