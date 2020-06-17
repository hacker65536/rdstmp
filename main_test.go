package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	code := m.Run()
	os.Exit(code)
}
