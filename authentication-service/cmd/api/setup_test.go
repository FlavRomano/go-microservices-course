package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// set up the environment
	os.Exit(m.Run())
}
