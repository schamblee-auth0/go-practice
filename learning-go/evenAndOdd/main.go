package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var v string
	for _, integer := range ints {
		if integer%2 == 0 {
			v = "even"
		} else {
			v = "odd"
		}
		fmt.Println(v)
	}
}
