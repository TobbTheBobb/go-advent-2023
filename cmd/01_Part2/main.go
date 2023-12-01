package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

var re = regexp.MustCompile(`oneight|threeight|fiveight|nineight|sevenine|twone|one|two|three|four|five|six|seven|eight|nine|\d`)

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

func convertToInt(s string,f bool) string {
	var value string
	switch s {
	case "twone":
		if (f) {
			value = "2"
		} else {
			value = "1"
		}
	case "oneight":
		if (f) {
			value = "1"
		} else {
			value = "8"
		}
	case "threeight":
		if (f) {
			value = "3"
		} else {
			value = "8"
		}
	case "fiveight":
		if (f) {
			value = "5"
		} else {
			value = "8"
		}
	case "nineight":
		if (f) {
			value = "9"
		} else {
			value = "8"
		}
	case "sevenine":
		if (f) {
			value = "7"
		} else {
			value = "9"
		}
	case "one":
		value = "1"
	case "two":
		value = "2"
	case "three":
		value = "3"
	case "four":
		value = "4"
	case "five":
		value = "5"
	case "six":
		value = "6"
	case "seven":
		value = "7"
	case "eight":
		value = "8"
	case "nine":
		value = "9"
	default:
		value = s
	}
	return value
}

func processStructure(structure []string) {
	var sumOfResults int
	for _,s := range structure {
		fmt.Print(s+"  -  ")
		foundNumbers := re.FindAllString(s,-1)
		first := foundNumbers[0]
		last := foundNumbers[len(foundNumbers)-1]
		fmt.Print(foundNumbers)
		fmt.Print(" - First: "+first+" - Last: "+last+" - ")
		firstValue := convertToInt(first,true)
		lastValue := convertToInt(last,false)
		resultString := firstValue+lastValue
		fmt.Print(resultString)
		resultValue,err := strconv.Atoi(resultString)
		check(err)
		sumOfResults += resultValue
		fmt.Print(" - Current sum: "+fmt.Sprint(sumOfResults)+"\n")
	}
}