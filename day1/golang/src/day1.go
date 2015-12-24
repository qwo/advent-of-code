package main

/// For Avent of Code
/// http://adventofcode.com/day/1/
/// Stanley Zheng

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

var blank = []byte("")
var upLevel = []byte("(")
var downlevel = []byte(")")

// Count the number of floors using string replacement
// Count the difference
func main() {

	file, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("file can't be read")
	}

	up := bytes.Replace(file, downlevel, blank, -1)
	down := bytes.Replace(file, upLevel, blank, -1)
	fmt.Printf("Santa went up %d floors and down %d floors and ended up on the %d floor", len(up), len(down), len(up)-len(down))

}
