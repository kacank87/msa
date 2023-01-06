package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var source []string
var target []string

func extrack(directory string, value int, tipe string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if tipe == "source" {
			source = append(source, file.Name())
		} else {
			target = append(target, file.Name())
		}

		if file.IsDir() == true {

			if value == 0 {
				fmt.Print("------")
				fmt.Println(file.Name(), file.IsDir())
				extrack(directory+"/"+file.Name(), 1, tipe)
			} else if value == 1 {
				fmt.Print("------------")
				fmt.Println(file.Name(), file.IsDir())
				extrack(directory+"/"+file.Name(), 2, tipe)
			} else if value == 2 {
				fmt.Print("------------------")
				fmt.Println(file.Name(), file.IsDir())
			}

		} else {
			fmt.Println(file.Name(), file.IsDir())
		}
	}
}

func comparasi(source []string, target []string) {
	b := 0
	for i := 0; i < len(source); i++ {
		for a := 0; a < len(target); a++ {
			if source[i] == target[a] {
				//fmt.Println(source[i], "- MATCH")
				b = 1
			}
		}
		if b != 1 {
			fmt.Println(source[i], " NEW")

		}
		b = 0
	}

	c := 0
	for i := 0; i < len(target); i++ {
		for a := 0; a < len(source); a++ {
			if target[i] == source[a] {
				//fmt.Println(source[i], "- MATCH")
				c = 1
			}
		}
		if c != 1 {
			fmt.Println(target[i], " DELETED")

		}
		c = 0
	}

}

func main() {
	extrack("./source", 0, "source")
	extrack("./target", 0, "target")
	fmt.Println("---------------------------")
	comparasi(source, target)
}
