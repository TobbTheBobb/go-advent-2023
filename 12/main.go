package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reQuestion = regexp.MustCompile("[?]+")
var reHash = regexp.MustCompile("[#]+")
var reDot = regexp.MustCompile("[.]+")

func main() {
	myInput_Structure := readFile()
	fmt.Printf("%#v\n",myInput_Structure)
	processStructure(myInput_Structure)
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

func processStructure(structure []string) {
	sumOfResults := float64(0)
	for _,s := range structure {
		//fmt.Println(s)
		springs := strings.Split(s," ")[0]
		numbers := strings.Split(strings.Split(s," ")[1],",")
		fmt.Printf("Springs: %v and Numbers: %v\n",springs,numbers)

		q := reQuestion.FindAllStringIndex(springs,-1)
		fmt.Printf("My q: %v\n",q)

		h := reHash.FindAllStringIndex(springs,-1)
		fmt.Printf("My h: %v\n",h)

		d := reDot.FindAllStringIndex(springs,-1)
		fmt.Printf("My d: %v\n",d)

		for i,n := range numbers {
			n,err := strconv.Atoi(n)
			check(err)
			if h[i][1]-h[i][0]<n {
				fmt.Println("Its smaller...")
			}
		}
	}
	fmt.Printf("Ergebnis: %v\n", sumOfResults)
}