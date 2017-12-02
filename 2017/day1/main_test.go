package main

import "testing"


func TestTimeConsuming(t *testing.T) {
	m:=make(map[string]int){{"1122": 3 }}
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
}