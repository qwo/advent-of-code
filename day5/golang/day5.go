package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("../day5.txt")
	count := 0
	if err != nil {
		fmt.Printf("%s", "File not found")
	}
	lines := strings.Split(string(data), "\n")
	for _, word := range lines {

		if NiceWord(word) {
			count++
		} else {
			// fmt.Printf("%s doesn't check out\n", word)
		}
	}
	fmt.Printf("The count is %d", count)

	// Restart the counter
	count = 0
	for _, word := range lines {

		if NiceWordPt2(word) {
			count++
		} else {
			fmt.Printf("%s doesn't check out\n", word)
		}
	}
	fmt.Printf("The count is %d", count)
}

// NaughtyWord determines if the string is naughty or nice
// and follows a set of rules
func NiceWord(w string) (check bool) {
	if doesNotContain(w, []string{"ab", "cd", "pq", "xy"}) != true {
		return
	}
	check = containsThreeVowels(w) && repeatingLetters(w)

	return
}

// NaughtyWord determines if the string is naughty or nice
// and follows a set of rules
func NiceWordPt2(w string) (check bool) {

	check = containsThreeVowels(w) && repeatingLetters(w)

	return
}

func containsThreeVowels(w string) (check bool) {
	vowels := "aeiou"
	count := 0
	for _, letter := range w {
		if strings.ContainsRune(vowels, letter) {
			count++
		}
	}
	if count >= 3 {
		check = true
		// fmt.Printf("%s has three letters\n", w)
	}
	return
}

func repeatingLetters(w string) (check bool) {
	for _, letter := range w {
		if strings.Contains(w, (string(letter) + string(letter))) {
			// fmt.Printf("%s repeats %s\n ", w, (string(letter) + string(letter)))
			check = true
			return
		}
	}
	return
}

func doesNotContain(w string, arr []string) (check bool) {
	for _, seq := range arr {
		if strings.Contains(w, seq) {
			// fmt.Printf("%s contains %s\n", w, seq)
			return
		}
	}
	check = true
	return
}
