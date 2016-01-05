package main

import "testing"

var list = []struct {
	word     string
	expected bool
}{
	{"ugknbfddgicrmopn", true},
	{"aaa", true},
	{"jchzalrnumimnmhp", false},
	{"haegwjzuvuyypxyu", false},
	{"dvszwmarrgswjxmb", false},
}

func TestNaughtOrNice(t *testing.T) {

	for _, s := range list {
		check := NaughtOrNice(s.word)
		if check != s.expected {
			t.Errorf("Fib(%s): expected %t, actual %t", s.word, s.expected, check)
		}
	}

}
