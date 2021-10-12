package main

import "testing"

func TestTruth(t *testing.T) {
	if true != true {
		t.Error("expected true to be true")
	}
}
