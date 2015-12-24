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

	part1(file)
	part2(file)
}

func part1(file []byte) {
	up := bytes.Replace(file, downlevel, blank, -1)
	down := bytes.Replace(file, upLevel, blank, -1)
	fmt.Printf("Santa went up %d floors and down %d floors and ended up on the %d floor\n", len(up), len(down), len(up)-len(down))

}

func part2(file []byte) {
	floorLevel := 0
	firstTime := true
	for i := 0; i < len(file); i++ {

		if string(file[i]) == "(" {
			floorLevel = floorLevel + 1
		} else {
			floorLevel = floorLevel - 1
		}
		if !firstTime && floorLevel < 0 {
			fmt.Printf("Santa enters the basement at char %d \n", i+1)
			break
		}
		firstTime = false
	}
}
