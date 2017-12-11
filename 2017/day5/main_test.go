package main

import "testing"

func TestValidPassphrase(t *testing.T) {

	tests := map[string]int{
		`0
		3
		0
		1
		-3`: 5,
	}

	for q, expected := range tests {
		if CheckJumps(q) != expected {
			t.Error("Utoh!")
		}
	}
}
