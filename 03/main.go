package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	//"github.com/fatih/color"
)

var reDigit = regexp.MustCompile(`(?:\d+)+`)
var reNoDot = regexp.MustCompile(`[^.\d\n]`)

func main() {
	myInput_Structure := readFile()
	fmt.Printf("%#v\n", myInput_Structure)
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

func checkMatch(content string,left int, right int) bool {
	if left < 0 {
		left = 0
	}
	if right >= len(content) {
		right = len(content)-1
	}
	return reNoDot.MatchString(content[left:right+1])
}

func solveTask1(lines []string) int {
	sum := 0
	for l, content := range lines {
		fmt.Println(content)
		digits := reDigit.FindAllString(content,-1)
		indexDigits := reDigit.FindAllStringIndex(content,-1)
		for count,digit := range digits {
			matched := false
			var leftValue,rightValue string
			left := indexDigits[count][0]-1
			right := indexDigits[count][1]+1
			if left >= 0 {
				leftValue = content[left:left]
				if reNoDot.MatchString(leftValue) {
					fmt.Println(leftValue)
					matched = true
				}
			}
			if right < len(content) && !matched {
				rightValue = content[right:right+1]
				if reNoDot.MatchString(rightValue) {
					fmt.Println(rightValue)
					matched = true
				}
			}
			if l > 0 && !matched {
				content = lines[l-1]
				matched = checkMatch(content,left,right)
			}
			if l < len(lines) && !matched {
				content = lines[l]
				matched = checkMatch(content,left,right)
			}
			if matched {
				value,err := strconv.Atoi(digit)
				check(err)
				sum = sum + value
				fmt.Printf("digit: %v MATCHED: %v\n", digit,matched)
			}
		}
	}
	return sum
}