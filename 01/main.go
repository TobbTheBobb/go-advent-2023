package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

var re = regexp.MustCompile(`\d`)

func main() {
	myInput_Structure := readFile()
	//fmt.Printf("%#v\n",myInput_Structure)
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
	sumOfResults := 0
	for _,s := range structure {
		fmt.Println(s)
		foundNumbers := re.FindAllString(s,-1)
		first := foundNumbers[0]
		last := foundNumbers[len(foundNumbers)-1]
		fmt.Println(foundNumbers)
		fmt.Println("First: "+first+" - Last: "+last)
		resultString := first+last
		fmt.Println("Result: "+resultString)
		result_value,err := strconv.Atoi(resultString)
		check(err)
		sumOfResults += result_value
		fmt.Println("Current sum: "+fmt.Sprint(sumOfResults))
	}
}