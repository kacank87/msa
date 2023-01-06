package main

import (
	"fmt"
	"sort"
)

func bar(size int, data []int) {
	for a := size + 1; a >= 0; a-- {
		for i := 0; i < len(data); i++ {
			if a >= data[i] {
				fmt.Print(" ")
			} else {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	primes := []int{1, 4, 5, 6, 8, 2}
	size := len(primes)
	varssort := []int{1, 4, 5, 6, 8, 2}
	varreserv := []int{1, 4, 5, 6, 8, 2}
	sort.Ints(varssort)
	sort.Sort(sort.Reverse(sort.IntSlice(varreserv)))
	bar(size, primes)
	fmt.Println(primes)
	fmt.Println("---------------------------------")
	bar(size, varssort)
	fmt.Println(varssort)
	fmt.Println("---------------------------------")
	bar(size, varreserv)
	fmt.Println(varreserv)
}
