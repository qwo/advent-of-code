package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const secretKey = "bgvyzdsv"
const marker = "00000"

func main() {
	found := false
	i := 0
	for !found {
		val := secretKey + strconv.Itoa(i)
		hash := md5.Sum([]byte(val))
		check := fmt.Sprintf("%x", hash[0:3])
		check = check[0:5]
		if check == marker {
			fmt.Printf("Found at %d", i)
			found = true
		}
		i++
	}
}
