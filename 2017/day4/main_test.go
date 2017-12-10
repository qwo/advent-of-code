package main

import "testing"

func TestValidPassphrase(t *testing.T) {

	tests := map[string]bool{
		"aa bb":           true,
		"aa bb cc dd ee":  true,
		"aa bb cc dd aa":  false,
		"aa bb cc dd aaa": true,
	}

	for q, expected := range tests {
		if ValidPassPhrase(q) != expected {
			t.Error("Utoh!")
		}
	}
}
