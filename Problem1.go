package main

import (
	"fmt"
	"sort"
)

func bar(data []int, tipe string) []int {

	switch tipe {
	case "asc":
		sort.Ints(data)
	case "desc":
		sort.Sort(sort.Reverse(sort.IntSlice(data)))
	}

	size := len(data)
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
	return data
}

func main() {
	primes := []int{1, 4, 5, 6, 8, 2}
	fmt.Println(bar(primes, "default"))
	fmt.Println("---------------------------------")
	fmt.Println(bar(primes, "asc"))
	fmt.Println("---------------------------------")
	fmt.Println(bar(primes, "desc"))
}
