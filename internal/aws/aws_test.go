package myaws

import (
	"testing"
)

func TestInit(t *testing.T) {

	a := "a"
	b := "a"
	if a != b {
		t.Errorf("Init() = ; want ok")
	}
}
