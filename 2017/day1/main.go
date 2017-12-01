package main

import "fmt"
import "strconv"
func test() {

}

func findSum (s string )(sum int){
	var last rune
	for _, v:= range s {
		// fmt.Printf("(k:%v v:%v)", k, strconv.Atoi(string(v)));
		if last == v { 
			val, _ := strconv.Atoi(string(v))
			sum += val
		}
		last = v
		// fmt.Println(v, sum)
	}

	if last == rune(s[0]) {
		val, _:= strconv.Atoi(string(last))
		sum += val
	}
	return sum
}

func main(){ 
	// m := make([]int, 10 )
	fmt.Println(findSum("1122")	) // 3
	fmt.Println(findSum("1111")) // 4
	fmt.Println(findSum("1234")) // 0
	fmt.Println(findSum("91212129")) 
}