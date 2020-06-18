package rds

import (
	"testing"
)

func TestGetrds(t *testing.T) {

	got := Getrds("test")
	if got != "ok" {
		t.Errorf("Getrds(\"test\") = %s; want ok", got)
	}
}
