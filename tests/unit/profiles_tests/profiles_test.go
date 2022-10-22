package profiles_tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup goes here

	//
	code := m.Run()
	// teardown goes here

	//
	os.Exit(code)
}

func TestSumProfile(t *testing.T) {
	total := 5 + 5
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
