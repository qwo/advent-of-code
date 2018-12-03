package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
		next := strings.Fields(scanner.Text())
		arr = append(arr, next)
	}
	return
}

func partA() {
	count := 0
	for _, i := range read() {
		a, _ := strconv.Atoi(i[0])
		b, _ := strconv.Atoi(i[1])
		c, _ := strconv.Atoi(i[2])
		if a+b > c && b+c > a && a+c > b {
			count++
		} else {
			// fmt.Println("false", i)
		}
	}
	fmt.Println(count)
}

func partB() {
	count := 0
	arr := read()
	for i := 0; i < (len(arr) - 2); i = i + 3 {
		for j := 0; j < 3; j++ {
			a, _ := strconv.Atoi(arr[i][j])
			b, _ := strconv.Atoi(arr[i+1][j])
			c, _ := strconv.Atoi(arr[i+2][j])
			//fmt.Println(a, b, c)
			if a+b > c && b+c > a && a+c > b {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	partA()
	partB()
}
