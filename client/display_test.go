package client_test

import gowl "github.com/mortdeus/gowl/client"
import "testing"

func TestDisplayConnect(t *testing.T) {
	if err := gowl.Display.Connect(""); err != nil {
		t.Fatal(err)
	}
}
