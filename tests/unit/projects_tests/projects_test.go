package projects_tests

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
