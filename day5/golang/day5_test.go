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

var list2 = []struct {
	word     string
	expected bool
}{
	{"qjhvhtzxzqqjkmpb", true},
	{"xxyxx", true},
	{"uurcxstgmygtbstg", false},
	{"ieodomkazucvgmuy", false},
}

func TestNaughtOrNice(t *testing.T) {

	for _, s := range list {
		check := NiceWord(s.word)
		if check != s.expected {
			t.Errorf("Fib(%s): expected %t, actual %t", s.word, s.expected, check)
		}
	}
}

func TestNaughtOrNice2(t *testing.T) {

	for _, s := range list2 {
		check := NiceWordPt2(s.word)
		if check != s.expected {
			t.Errorf("Fib(%s): expected %t, actual %t", s.word, s.expected, check)
		}
	}
}
