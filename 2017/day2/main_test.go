package main

import "testing"

type Case struct {
	m   [][]int
	sum int
}

func TestCheckSum(t *testing.T) {
	cases := []Case{
		Case{
			[][]int{
				{5, 1, 9, 5},
				{7, 5, 3},
				{2, 4, 6, 8},
			},
			18,
		},
		Case{
			[][]int{
				{},
				{7, 5, 3},
				{2, 4, 6, 8},
			},
			10,
		}}

	for k, c := range cases {
		if CheckSum(c.m) != c.sum {
			t.Error("No match on case:", k)
		}
	}
}

func TestParsing(arr string) [][]int {

}
