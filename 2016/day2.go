package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func move() {}
func valid(x, y int) bool {
	if math.Abs(float64(x)) <= 1 && math.Abs(float64(y)) <= 1 {
		return true
	}
	return false
}

func read() (arr [][]string) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// bf.Add(scanner.Text())
		next := strings.Split(scanner.Text(), "")
		arr = append(arr, next)
	}
	return
}

func main() {
	lookup := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var x, y int = 0, 0

	for _, i := range read() {
		for j := range i {
			// fmt.Printf(i[j])
			switch i[j] {
			case "L":
				if valid(x+1, y) {
					x++
				}
				// fmt.Println("Left")
			case "R":
				if valid(x-1, y) {
					x--
				}
			case "U":
				if valid(x, y-1) {
					y--
				}
			case "D":
				if valid(x, y+1) {
					y++
				}
			}
		}
		fmt.Println(lookup[x+1][y+1])
		fmt.Println(x, y)
	}
}
