package main

import (
	"authentication-service/data"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	// set up the environment
	repo := data.NewPostgresTestRepository(nil) // it's mocked, we don't need a real DB conn pointer.
	testApp.Repository = repo
	os.Exit(m.Run())
}
