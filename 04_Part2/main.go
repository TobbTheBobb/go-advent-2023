package main

import (
	"fmt"
	"os"
	"strings"
)

var sumOfResults = 0

func main() {
	myInput_Structure := readFile()
	fmt.Printf("%#v\n",myInput_Structure)
	sumOfResults = processStructure(myInput_Structure)
	fmt.Println(sumOfResults)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() []string {
	dat, err := os.ReadFile("./input")
	check(err)
	input := string(dat)
	//fmt.Println("Input-File:")
	//fmt.Println(input)
	myInput_Structure := strings.Split(input, "\n")
	return myInput_Structure
}

func processStructure(structure []string) int {
	if len(structure) > 0 {
		for c,s := range structure {
			//fmt.Println(s)
			lists := strings.Split(s, "|")
			list1 := strings.TrimSpace(strings.Split(lists[0],":")[1])
			values2 := strings.Split(lists[1]," ")
			values1 := strings.Split(list1," ")
			m := make(map[string]bool)
			for _,value := range values1 {
				for _,value2 := range values2 {
					if !m[value] && value == value2 && value != ""{
						m[value] = true
					}
				}
			}
			fmt.Println(m)
			if len(m) >= 1 {
				fmt.Println(sumOfResults)
				return sumOfResults+processStructure(structure[c+1:len(m)])
			} else {
				sumOfResults = sumOfResults+1
			}
		}
	} else {
		return 1
	}
	return 1
}