package main

import (
	"fmt"
	"os"
	"regexp"
	//"strconv"
	"strings"
)

var reRatio = regexp.MustCompile(`(?:\*)+`)
var reDigit = regexp.MustCompile(`\d+`)

func main() {
	myInput_Structure := readFile()
	erg := solveTask1(myInput_Structure)
	fmt.Println(erg)
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
	myInput_Structure := strings.Split(input, "\n")
	return myInput_Structure
}

func checkMatch(str string,left int, right int) bool {
	if left < 0 {
		left = 0
	}
	if right >= len(str) {
		right = len(str)-1
	}
	//fmt.Println(str[left:right+1])
	return reDigit.MatchString(str[left:right+1])
}

func solveTask1(lines []string) int {
	sum := 0
	for l, content := range lines {
		if l > 0 {
			fmt.Println(lines[l-1])
		}
		fmt.Println(content)
		if l < len(lines)-1 {
			fmt.Println(lines[l+1])
		}
		ratios := reRatio.FindAllString(content,-1)
		indexRatios := reRatio.FindAllStringIndex(content,-1)
		for count,ratio := range ratios {
			matched := false
			var leftValue,rightValue string
			left := indexRatios[count][0]-1
			right := indexRatios[count][1]
			if left >= 0 {
				leftValue = content[left:left+1]
				if reDigit.MatchString(leftValue) {
					//fmt.Println(leftValue)
					matched = true
				}
			}
			if right < len(content) && !matched {
				rightValue = content[right:right+1]
				if reDigit.MatchString(rightValue) {
					//fmt.Println(rightValue)
					matched = true
				}
			}
			if l > 0 && !matched {
				contentb := lines[l-1]
				matched = checkMatch(contentb,left,right)
			}
			if l < len(lines)-2 && !matched {
				contentn := lines[l+1]
				matched = checkMatch(contentn,left,right)
			}
			if matched {
				// value,err := strconv.Atoi(ratio)
				// check(err)
				// sum = sum + value
				fmt.Printf("%v MATCHED: %v\n", ratio,matched)
			}
		}
	}
	return sum
}